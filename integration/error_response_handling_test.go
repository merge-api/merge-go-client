package integration

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/merge-api/merge-go-client/v2/core"
	"github.com/merge-api/merge-go-client/v2/ticketing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrorHandling_404RaisesAPIErrorWithStatusCode(t *testing.T) {
	c, _ := captureClient(t, func(_ *http.Request) *http.Response {
		return jsonResponse(404, `{"detail":"Not found."}`, map[string]string{"Content-Type": "application/json"})
	})
	_, err := c.Ticketing.Tickets.Retrieve(context.Background(), "nonexistent-id", &ticketing.TicketsRetrieveRequest{})
	require.Error(t, err)
	var apiErr *core.APIError
	require.True(t, errors.As(err, &apiErr), "error should be *core.APIError")
	assert.Equal(t, 404, apiErr.StatusCode)
}

func TestErrorHandling_429PreservesRetryAfterHeader(t *testing.T) {
	c, _ := captureClient(t, func(_ *http.Request) *http.Response {
		return jsonResponse(429, `{"detail":"Request was throttled."}`, map[string]string{
			"Content-Type": "application/json",
			"Retry-After":  "60",
		})
	})
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{})
	require.Error(t, err)
	var apiErr *core.APIError
	require.True(t, errors.As(err, &apiErr), "error should be *core.APIError")
	assert.Equal(t, 429, apiErr.StatusCode)
	assert.Equal(t, "60", apiErr.Header.Get("Retry-After"))
}

func TestErrorHandling_NonJSONBodyDoesNotCrash(t *testing.T) {
	c, _ := captureClient(t, func(_ *http.Request) *http.Response {
		return jsonResponse(502, "Bad Gateway", map[string]string{"Content-Type": "text/html"})
	})
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{})
	require.Error(t, err)
	var apiErr *core.APIError
	require.True(t, errors.As(err, &apiErr), "error should be *core.APIError")
	assert.Equal(t, 502, apiErr.StatusCode)
}

func TestErrorHandling_200DoesNotError(t *testing.T) {
	c, _ := captureClient(t, func(_ *http.Request) *http.Response {
		return jsonListResponse()
	})
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{})
	require.NoError(t, err)
}
