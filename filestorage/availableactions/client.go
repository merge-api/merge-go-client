// Code generated by Fern. DO NOT EDIT.

package availableactions

import (
	context "context"
	core "github.com/merge-api/merge-go-client/v2/core"
	filestorage "github.com/merge-api/merge-go-client/v2/filestorage"
	internal "github.com/merge-api/merge-go-client/v2/internal"
	option "github.com/merge-api/merge-go-client/v2/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// Returns a list of models and actions available for an account.
func (c *Client) Retrieve(
	ctx context.Context,
	opts ...option.RequestOption,
) (*filestorage.AvailableActions, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"",
	)
	endpointURL := baseURL + "/filestorage/v1/available-actions"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *filestorage.AvailableActions
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
