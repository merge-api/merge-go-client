// This file was auto-generated by Fern from our API Definition.

package hris

import (
	fmt "fmt"
	time "time"
)

type TimesheetEntryEndpointRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool                  `json:"-"`
	Model    *TimesheetEntryRequest `json:"model,omitempty"`
}

type TimesheetEntriesListRequest struct {
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// If provided, will only return timesheet entries for this employee.
	EmployeeId *string `json:"-"`
	// If provided, will only return timesheet entries ended after this datetime.
	EndedAfter *time.Time `json:"-"`
	// If provided, will only return timesheet entries ended before this datetime.
	EndedBefore *time.Time `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *string `json:"-"`
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
	// Overrides the default ordering for this endpoint. Possible values include: start_time, -start_time.
	OrderBy *TimesheetEntriesListRequestOrderBy `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
	// If provided, will only return timesheet entries started after this datetime.
	StartedAfter *time.Time `json:"-"`
	// If provided, will only return timesheet entries started before this datetime.
	StartedBefore *time.Time `json:"-"`
}

type TimesheetEntriesRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *string `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
}

type TimesheetEntriesListRequestOrderBy string

const (
	TimesheetEntriesListRequestOrderByStartTimeDescending TimesheetEntriesListRequestOrderBy = "-start_time"
	TimesheetEntriesListRequestOrderByStartTimeAscending  TimesheetEntriesListRequestOrderBy = "start_time"
)

func NewTimesheetEntriesListRequestOrderByFromString(s string) (TimesheetEntriesListRequestOrderBy, error) {
	switch s {
	case "-start_time":
		return TimesheetEntriesListRequestOrderByStartTimeDescending, nil
	case "start_time":
		return TimesheetEntriesListRequestOrderByStartTimeAscending, nil
	}
	var t TimesheetEntriesListRequestOrderBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (t TimesheetEntriesListRequestOrderBy) Ptr() *TimesheetEntriesListRequestOrderBy {
	return &t
}
