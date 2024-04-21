// This file was auto-generated by Fern from our API Definition.

package attachments

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	core "github.com/merge-api/merge-go-client/core"
	ticketing "github.com/merge-api/merge-go-client/ticketing"
	io "io"
	http "net/http"
	url "net/url"
	time "time"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

// Returns a list of `Attachment` objects.
func (c *Client) List(ctx context.Context, request *ticketing.AttachmentsListRequest) (*ticketing.PaginatedAttachmentList, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "ticketing/v1/attachments"

	queryParams := make(url.Values)
	if request.CreatedAfter != nil {
		queryParams.Add("created_after", fmt.Sprintf("%v", request.CreatedAfter.Format(time.RFC3339)))
	}
	if request.CreatedBefore != nil {
		queryParams.Add("created_before", fmt.Sprintf("%v", request.CreatedBefore.Format(time.RFC3339)))
	}
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.Expand != nil {
		queryParams.Add("expand", fmt.Sprintf("%v", request.Expand))
	}
	if request.IncludeDeletedData != nil {
		queryParams.Add("include_deleted_data", fmt.Sprintf("%v", *request.IncludeDeletedData))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if request.ModifiedAfter != nil {
		queryParams.Add("modified_after", fmt.Sprintf("%v", request.ModifiedAfter.Format(time.RFC3339)))
	}
	if request.ModifiedBefore != nil {
		queryParams.Add("modified_before", fmt.Sprintf("%v", request.ModifiedBefore.Format(time.RFC3339)))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if request.RemoteCreatedAfter != nil {
		queryParams.Add("remote_created_after", fmt.Sprintf("%v", request.RemoteCreatedAfter.Format(time.RFC3339)))
	}
	if request.RemoteId != nil {
		queryParams.Add("remote_id", fmt.Sprintf("%v", *request.RemoteId))
	}
	if request.TicketId != nil {
		queryParams.Add("ticket_id", fmt.Sprintf("%v", *request.TicketId))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *ticketing.PaginatedAttachmentList
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Creates an `Attachment` object with the given values.
func (c *Client) Create(ctx context.Context, request *ticketing.TicketingAttachmentEndpointRequest) (*ticketing.TicketingAttachmentResponse, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "ticketing/v1/attachments"

	queryParams := make(url.Values)
	if request.IsDebugMode != nil {
		queryParams.Add("is_debug_mode", fmt.Sprintf("%v", *request.IsDebugMode))
	}
	if request.RunAsync != nil {
		queryParams.Add("run_async", fmt.Sprintf("%v", *request.RunAsync))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *ticketing.TicketingAttachmentResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns an `Attachment` object with the given `id`.
func (c *Client) Retrieve(ctx context.Context, id string, request *ticketing.AttachmentsRetrieveRequest) (*ticketing.Attachment, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"ticketing/v1/attachments/%v", id)

	queryParams := make(url.Values)
	if request.Expand != nil {
		queryParams.Add("expand", fmt.Sprintf("%v", request.Expand))
	}
	if request.IncludeRemoteData != nil {
		queryParams.Add("include_remote_data", fmt.Sprintf("%v", *request.IncludeRemoteData))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *ticketing.Attachment
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns an `Attachment` object with the given `id`.
func (c *Client) DownloadRetrieve(ctx context.Context, id string, request *ticketing.AttachmentsDownloadRetrieveRequest) (io.Reader, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"ticketing/v1/attachments/%v/download", id)

	queryParams := make(url.Values)
	if request.MimeType != nil {
		queryParams.Add("mime_type", fmt.Sprintf("%v", *request.MimeType))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	response := bytes.NewBuffer(nil)
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Returns metadata for `TicketingAttachment` POSTs.
func (c *Client) MetaPostRetrieve(ctx context.Context) (*ticketing.MetaResponse, error) {
	baseURL := "https://api.merge.dev/api"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "ticketing/v1/attachments/meta/post"

	var response *ticketing.MetaResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  c.header,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
