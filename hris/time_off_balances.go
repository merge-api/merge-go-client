// This file was auto-generated by Fern from our API Definition.

package hris

import (
	time "time"
)

type TimeOffBalancesListRequest struct {
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// If provided, will only return time off balances for this employee.
	EmployeeId *string `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *string `json:"-"`
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
	// If provided, will only return TimeOffBalance with this policy type. Options: ('VACATION', 'SICK', 'PERSONAL', 'JURY_DUTY', 'VOLUNTEER', 'BEREAVEMENT')
	//
	// * `VACATION` - VACATION
	// * `SICK` - SICK
	// * `PERSONAL` - PERSONAL
	// * `JURY_DUTY` - JURY_DUTY
	// * `VOLUNTEER` - VOLUNTEER
	// * `BEREAVEMENT` - BEREAVEMENT
	PolicyType *TimeOffBalancesListRequestPolicyType `json:"-"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *string `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *string `json:"-"`
}

type TimeOffBalancesRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *string `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *string `json:"-"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *string `json:"-"`
}
