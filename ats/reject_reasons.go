// This file was auto-generated by Fern from our API Definition.

package ats

import (
	time "time"
)

type RejectReasonsListRequest struct {
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
	IncludeShellData *bool `json:"-"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
}

type RejectReasonsRetrieveRequest struct {
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
}
