package integration

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/merge-api/merge-go-client/v2/client"
	"github.com/merge-api/merge-go-client/v2/option"
)

func intPtr(v int) *int    { return &v }
func strPtr(v string) *string { return &v }

// roundTripFunc is a RoundTripper backed by a plain function.
type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

// jsonListResponse returns a minimal paginated list response with no results.
func jsonListResponse() *http.Response {
	return &http.Response{
		StatusCode: 200,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(`{"next":null,"previous":null,"results":[]}`)),
	}
}

// jsonResponse builds an HTTP response with the given status, body, and optional headers.
func jsonResponse(statusCode int, body string, headers map[string]string) *http.Response {
	h := make(http.Header)
	for k, v := range headers {
		h.Set(k, v)
	}
	return &http.Response{
		StatusCode: statusCode,
		Header:     h,
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}

// captureClient builds a Merge client that intercepts all HTTP calls via respond.
// Extra opts are applied after the default api-key / account-token / http-client options,
// so callers can override the base URL for environment-routing tests.
func captureClient(t *testing.T, respond func(*http.Request) *http.Response, opts ...option.RequestOption) (*client.Client, func() []*http.Request) {
	t.Helper()
	var mu sync.Mutex
	var reqs []*http.Request

	transport := roundTripFunc(func(req *http.Request) (*http.Response, error) {
		mu.Lock()
		reqs = append(reqs, req)
		mu.Unlock()
		return respond(req), nil
	})

	token := "test-token"
	baseOpts := []option.RequestOption{
		option.WithApiKey("test-key"),
		option.WithAccountToken(&token),
		option.WithHTTPClient(&http.Client{Transport: transport}),
	}
	baseOpts = append(baseOpts, opts...)

	return client.NewClient(baseOpts...), func() []*http.Request {
		mu.Lock()
		defer mu.Unlock()
		out := make([]*http.Request, len(reqs))
		copy(out, reqs)
		return out
	}
}
