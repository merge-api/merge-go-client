// This file was auto-generated by Fern from our API Definition.

package hris

import (
	time "time"
)

type PayrollRunsListRequest struct {
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// If provided, will only return payroll runs ended after this datetime.
	EndedAfter *time.Time `json:"-"`
	// If provided, will only return payroll runs ended before this datetime.
	EndedBefore *time.Time `json:"-"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *PayrollRunsListRequestRemoteFields `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// If provided, will only return PayrollRun's with this status. Options: ('REGULAR', 'OFF_CYCLE', 'CORRECTION', 'TERMINATION', 'SIGN_ON_BONUS')
	//
	// * `REGULAR` - REGULAR
	// * `OFF_CYCLE` - OFF_CYCLE
	// * `CORRECTION` - CORRECTION
	// * `TERMINATION` - TERMINATION
	// * `SIGN_ON_BONUS` - SIGN_ON_BONUS
	RunType *PayrollRunsListRequestRunType `json:"-"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *PayrollRunsListRequestShowEnumOrigins `json:"-"`
	// If provided, will only return payroll runs started after this datetime.
	StartedAfter *time.Time `json:"-"`
	// If provided, will only return payroll runs started before this datetime.
	StartedBefore *time.Time `json:"-"`
}

type PayrollRunsRetrieveRequest struct {
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *PayrollRunsRetrieveRequestRemoteFields `json:"-"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *PayrollRunsRetrieveRequestShowEnumOrigins `json:"-"`
}
