// This file was auto-generated by Fern from our API Definition.

package ticketing

type LinkedAccountsListRequest struct {
	// Options: ('hris', 'ats', 'accounting', 'ticketing', 'crm', 'mktg', 'filestorage')
	//
	// * `hris` - hris
	// * `ats` - ats
	// * `accounting` - accounting
	// * `ticketing` - ticketing
	// * `crm` - crm
	// * `mktg` - mktg
	// * `filestorage` - filestorage
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
	// Filter by status. Options: `COMPLETE`, `INCOMPLETE`, `RELINK_NEEDED`
	Status *string `json:"-"`
}
