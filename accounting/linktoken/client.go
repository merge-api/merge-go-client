// Code generated by Fern. DO NOT EDIT.

package linktoken

import (
	context "context"
	accounting "github.com/merge-api/merge-go-client/v2/accounting"
	core "github.com/merge-api/merge-go-client/v2/core"
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

// Creates a link token to be used when linking a new end user.
func (c *Client) Create(
	ctx context.Context,
	request *accounting.EndUserDetailsRequest,
	opts ...option.RequestOption,
) (*accounting.LinkToken, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"",
	)
	endpointURL := baseURL + "/accounting/v1/link-token"
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")

	var response *accounting.LinkToken
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
