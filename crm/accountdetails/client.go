// This file was auto-generated by Fern from our API Definition.

package accountdetails

import (
	context "context"
	core "github.com/merge-api/merge-go-client/core"
	crm "github.com/merge-api/merge-go-client/crm"
	http "net/http"
)

type Client interface {
	Retrieve(ctx context.Context) (*crm.AccountDetails, error)
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

// Get details for a linked account.
func (c *client) Retrieve(ctx context.Context) (*crm.AccountDetails, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/crm/v1/account-details"

	var response *crm.AccountDetails
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		false,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}