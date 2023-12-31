// This file was auto-generated by Fern from our API Definition.

package asyncpassthrough

import (
	context "context"
	fmt "fmt"
	core "github.com/merge-api/merge-go-client/core"
	crm "github.com/merge-api/merge-go-client/crm"
	http "net/http"
)

type Client interface {
	Create(ctx context.Context, request *crm.DataPassthroughRequest) (*crm.AsyncPassthroughReciept, error)
	Retrieve(ctx context.Context, asyncPassthroughReceiptId string) (*crm.RemoteResponse, error)
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

// Asynchronously pull data from an endpoint not currently supported by Merge.
func (c *client) Create(ctx context.Context, request *crm.DataPassthroughRequest) (*crm.AsyncPassthroughReciept, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/crm/v1/async-passthrough"

	var response *crm.AsyncPassthroughReciept
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

// Retrieves data from earlier async-passthrough POST request
func (c *client) Retrieve(ctx context.Context, asyncPassthroughReceiptId string) (*crm.RemoteResponse, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"api/crm/v1/async-passthrough/%v", asyncPassthroughReceiptId)

	var response *crm.RemoteResponse
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
