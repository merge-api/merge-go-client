// Code generated by Fern. DO NOT EDIT.

package generalledgertransactions

import (
	context "context"
	fmt "fmt"
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

// Returns a list of `GeneralLedgerTransaction` objects.
func (c *Client) List(
	ctx context.Context,
	request *accounting.GeneralLedgerTransactionsListRequest,
	opts ...option.RequestOption,
) (*core.Page[*accounting.GeneralLedgerTransaction], error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"",
	)
	endpointURL := baseURL + "/accounting/v1/general-ledger-transactions"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	prepareCall := func(pageRequest *internal.PageRequest[*string]) *internal.CallParams {
		if pageRequest.Cursor != nil {
			queryParams.Set("cursor", fmt.Sprintf("%v", *pageRequest.Cursor))
		}
		nextURL := endpointURL
		if len(queryParams) > 0 {
			nextURL += "?" + queryParams.Encode()
		}
		return &internal.CallParams{
			URL:             nextURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        pageRequest.Response,
		}
	}
	readPageResponse := func(response *accounting.PaginatedGeneralLedgerTransactionList) *internal.PageResponse[*string, *accounting.GeneralLedgerTransaction] {
		var zeroValue *string
		next := response.Next
		results := response.Results
		return &internal.PageResponse[*string, *accounting.GeneralLedgerTransaction]{
			Next:    next,
			Results: results,
			Done:    next == zeroValue,
		}
	}
	pager := internal.NewCursorPager(
		c.caller,
		prepareCall,
		readPageResponse,
	)
	return pager.GetPage(ctx, request.Cursor)
}

// Returns a `GeneralLedgerTransaction` object with the given `id`.
func (c *Client) Retrieve(
	ctx context.Context,
	id string,
	request *accounting.GeneralLedgerTransactionsRetrieveRequest,
	opts ...option.RequestOption,
) (*accounting.GeneralLedgerTransaction, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/accounting/v1/general-ledger-transactions/%v",
		id,
	)
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)

	var response *accounting.GeneralLedgerTransaction
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
