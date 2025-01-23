// This file was auto-generated by Fern from our API Definition.

package crm

import (
	fmt "fmt"
)

type LinkedAccountsListRequest struct {
	// Options: `accounting`, `ats`, `crm`, `filestorage`, `hris`, `mktg`, `ticketing`
	//
	// - `hris` - hris
	// - `ats` - ats
	// - `accounting` - accounting
	// - `ticketing` - ticketing
	// - `crm` - crm
	// - `mktg` - mktg
	// - `filestorage` - filestorage
	Category *LinkedAccountsListRequestCategory `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// If provided, will only return linked accounts associated with the given email address.
	EndUserEmailAddress *string `json:"-"`
	// If provided, will only return linked accounts associated with the given organization name.
	EndUserOrganizationName *string `json:"-"`
	// If provided, will only return linked accounts associated with the given origin ID.
	EndUserOriginId *string `json:"-"`
	// Comma-separated list of EndUser origin IDs, making it possible to specify multiple EndUsers at once.
	EndUserOriginIds *string `json:"-"`
	Id               *string `json:"-"`
	// Comma-separated list of LinkedAccount IDs, making it possible to specify multiple LinkedAccounts at once.
	Ids *string `json:"-"`
	// If `true`, will include complete production duplicates of the account specified by the `id` query parameter in the response. `id` must be for a complete production linked account.
	IncludeDuplicates *bool `json:"-"`
	// If provided, will only return linked accounts associated with the given integration name.
	IntegrationName *string `json:"-"`
	// If included, will only include test linked accounts. If not included, will only include non-test linked accounts.
	IsTestAccount *string `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// Filter by status. Options: `COMPLETE`, `IDLE`, `INCOMPLETE`, `RELINK_NEEDED`
	Status *string `json:"-"`
}

type LinkedAccountsListRequestCategory string

const (
	LinkedAccountsListRequestCategoryAccounting  LinkedAccountsListRequestCategory = "accounting"
	LinkedAccountsListRequestCategoryAts         LinkedAccountsListRequestCategory = "ats"
	LinkedAccountsListRequestCategoryCrm         LinkedAccountsListRequestCategory = "crm"
	LinkedAccountsListRequestCategoryFilestorage LinkedAccountsListRequestCategory = "filestorage"
	LinkedAccountsListRequestCategoryHris        LinkedAccountsListRequestCategory = "hris"
	LinkedAccountsListRequestCategoryMktg        LinkedAccountsListRequestCategory = "mktg"
	LinkedAccountsListRequestCategoryTicketing   LinkedAccountsListRequestCategory = "ticketing"
)

func NewLinkedAccountsListRequestCategoryFromString(s string) (LinkedAccountsListRequestCategory, error) {
	switch s {
	case "accounting":
		return LinkedAccountsListRequestCategoryAccounting, nil
	case "ats":
		return LinkedAccountsListRequestCategoryAts, nil
	case "crm":
		return LinkedAccountsListRequestCategoryCrm, nil
	case "filestorage":
		return LinkedAccountsListRequestCategoryFilestorage, nil
	case "hris":
		return LinkedAccountsListRequestCategoryHris, nil
	case "mktg":
		return LinkedAccountsListRequestCategoryMktg, nil
	case "ticketing":
		return LinkedAccountsListRequestCategoryTicketing, nil
	}
	var t LinkedAccountsListRequestCategory
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LinkedAccountsListRequestCategory) Ptr() *LinkedAccountsListRequestCategory {
	return &l
}
