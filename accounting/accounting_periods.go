// This file was auto-generated by Fern from our API Definition.

package accounting

type AccountingPeriodsListRequest struct {
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
}

type AccountingPeriodsRetrieveRequest struct {
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
}
