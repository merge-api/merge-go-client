// Code generated by Fern. DO NOT EDIT.

package accounting

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/merge-api/merge-go-client/v2/internal"
)

type AccountingPeriodsListRequest struct {
	// The pagination cursor value.
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
	IncludeDeletedData *bool `json:"-" url:"include_deleted_data,omitempty"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-" url:"include_remote_data,omitempty"`
	// Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
	IncludeShellData *bool `json:"-" url:"include_shell_data,omitempty"`
	// Number of results to return per page.
	PageSize *int `json:"-" url:"page_size,omitempty"`
}

type AccountingPeriodsRetrieveRequest struct {
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-" url:"include_remote_data,omitempty"`
	// Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
	IncludeShellData *bool `json:"-" url:"include_shell_data,omitempty"`
}

type PaginatedAccountingPeriodList struct {
	Next     *string             `json:"next,omitempty" url:"next,omitempty"`
	Previous *string             `json:"previous,omitempty" url:"previous,omitempty"`
	Results  []*AccountingPeriod `json:"results,omitempty" url:"results,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PaginatedAccountingPeriodList) GetNext() *string {
	if p == nil {
		return nil
	}
	return p.Next
}

func (p *PaginatedAccountingPeriodList) GetPrevious() *string {
	if p == nil {
		return nil
	}
	return p.Previous
}

func (p *PaginatedAccountingPeriodList) GetResults() []*AccountingPeriod {
	if p == nil {
		return nil
	}
	return p.Results
}

func (p *PaginatedAccountingPeriodList) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaginatedAccountingPeriodList) UnmarshalJSON(data []byte) error {
	type unmarshaler PaginatedAccountingPeriodList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PaginatedAccountingPeriodList(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaginatedAccountingPeriodList) String() string {
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
