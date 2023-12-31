// This file was auto-generated by Fern from our API Definition.

package audittrail

import (
	context "context"
	fmt "fmt"
	core "github.com/merge-api/merge-go-client/core"
	filestorage "github.com/merge-api/merge-go-client/filestorage"
	http "net/http"
	url "net/url"
)

type Client interface {
	List(ctx context.Context, request *filestorage.AuditTrailListRequest) (*filestorage.PaginatedAuditLogEventList, error)
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

// Gets a list of audit trail events.
func (c *client) List(ctx context.Context, request *filestorage.AuditTrailListRequest) (*filestorage.PaginatedAuditLogEventList, error) {
	baseURL := "https://api.merge.dev"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "api/filestorage/v1/audit-trail"

	queryParams := make(url.Values)
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if request.EndDate != nil {
		queryParams.Add("end_date", fmt.Sprintf("%v", *request.EndDate))
	}
	if request.EventType != nil {
		queryParams.Add("event_type", fmt.Sprintf("%v", *request.EventType))
	}
	if request.PageSize != nil {
		queryParams.Add("page_size", fmt.Sprintf("%v", *request.PageSize))
	}
	if request.StartDate != nil {
		queryParams.Add("start_date", fmt.Sprintf("%v", *request.StartDate))
	}
	if request.UserEmail != nil {
		queryParams.Add("user_email", fmt.Sprintf("%v", *request.UserEmail))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	var response *filestorage.PaginatedAuditLogEventList
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
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
