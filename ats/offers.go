// Code generated by Fern. DO NOT EDIT.

package ats

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/merge-api/merge-go-client/v2/internal"
	time "time"
)

type OffersListRequest struct {
	// If provided, will only return offers for this application.
	ApplicationId *string `json:"-" url:"application_id,omitempty"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-" url:"created_after,omitempty"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-" url:"created_before,omitempty"`
	// If provided, will only return offers created by this user.
	CreatorId *string `json:"-" url:"creator_id,omitempty"`
	// The pagination cursor value.
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand []*OffersListRequestExpandItem `json:"-" url:"expand,omitempty"`
	// Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
	IncludeDeletedData *bool `json:"-" url:"include_deleted_data,omitempty"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-" url:"include_remote_data,omitempty"`
	// Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
	IncludeShellData *bool `json:"-" url:"include_shell_data,omitempty"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-" url:"modified_after,omitempty"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-" url:"modified_before,omitempty"`
	// Number of results to return per page.
	PageSize *int `json:"-" url:"page_size,omitempty"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *string `json:"-" url:"remote_fields,omitempty"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-" url:"remote_id,omitempty"`
	// A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
	ShowEnumOrigins *string `json:"-" url:"show_enum_origins,omitempty"`
}

type OffersRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand []*OffersRetrieveRequestExpandItem `json:"-" url:"expand,omitempty"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-" url:"include_remote_data,omitempty"`
	// Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
	IncludeShellData *bool `json:"-" url:"include_shell_data,omitempty"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *string `json:"-" url:"remote_fields,omitempty"`
	// A comma separated list of enum field names for which you'd like the original values to be returned, instead of Merge's normalized enum values. [Learn more](https://help.merge.dev/en/articles/8950958-show_enum_origins-query-parameter)
	ShowEnumOrigins *string `json:"-" url:"show_enum_origins,omitempty"`
}

type OffersListRequestExpandItem string

const (
	OffersListRequestExpandItemApplication OffersListRequestExpandItem = "application"
	OffersListRequestExpandItemCreator     OffersListRequestExpandItem = "creator"
)

func NewOffersListRequestExpandItemFromString(s string) (OffersListRequestExpandItem, error) {
	switch s {
	case "application":
		return OffersListRequestExpandItemApplication, nil
	case "creator":
		return OffersListRequestExpandItemCreator, nil
	}
	var t OffersListRequestExpandItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (o OffersListRequestExpandItem) Ptr() *OffersListRequestExpandItem {
	return &o
}

type OffersRetrieveRequestExpandItem string

const (
	OffersRetrieveRequestExpandItemApplication OffersRetrieveRequestExpandItem = "application"
	OffersRetrieveRequestExpandItemCreator     OffersRetrieveRequestExpandItem = "creator"
)

func NewOffersRetrieveRequestExpandItemFromString(s string) (OffersRetrieveRequestExpandItem, error) {
	switch s {
	case "application":
		return OffersRetrieveRequestExpandItemApplication, nil
	case "creator":
		return OffersRetrieveRequestExpandItemCreator, nil
	}
	var t OffersRetrieveRequestExpandItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (o OffersRetrieveRequestExpandItem) Ptr() *OffersRetrieveRequestExpandItem {
	return &o
}

type PaginatedOfferList struct {
	Next     *string  `json:"next,omitempty" url:"next,omitempty"`
	Previous *string  `json:"previous,omitempty" url:"previous,omitempty"`
	Results  []*Offer `json:"results,omitempty" url:"results,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PaginatedOfferList) GetNext() *string {
	if p == nil {
		return nil
	}
	return p.Next
}

func (p *PaginatedOfferList) GetPrevious() *string {
	if p == nil {
		return nil
	}
	return p.Previous
}

func (p *PaginatedOfferList) GetResults() []*Offer {
	if p == nil {
		return nil
	}
	return p.Results
}

func (p *PaginatedOfferList) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaginatedOfferList) UnmarshalJSON(data []byte) error {
	type unmarshaler PaginatedOfferList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PaginatedOfferList(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaginatedOfferList) String() string {
	if len(p.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(p.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}
