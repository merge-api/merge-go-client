// This file was auto-generated by Fern from our API Definition.

package linktoken

import (
	context "context"
	accounting "github.com/merge-api/merge-go-client/accounting"
	core "github.com/merge-api/merge-go-client/core"
	http "net/http"
)

type Client interface {
	Create(ctx context.Context, request *accounting.EndUserDetailsRequest) (*accounting.LinkToken, error)
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type client struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

// Creates a link token to be used when linking a new end user.
func (c *client) Create(ctx context.Context, request *accounting.EndUserDetailsRequest) (*accounting.LinkToken, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/accounting/v1/link-token"

	var response *accounting.LinkToken
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}
