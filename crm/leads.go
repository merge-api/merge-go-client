// This file was auto-generated by Fern from our API Definition.

package crm

import (
	fmt "fmt"
	time "time"
)

type LeadEndpointRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool        `json:"-"`
	Model    *LeadRequest `json:"model,omitempty"`
}

type LeadsListRequest struct {
	// If provided, will only return leads with this account.
	ConvertedAccountId *string `json:"-"`
	// If provided, will only return leads with this contact.
	ConvertedContactId *string `json:"-"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-"`
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// If provided, will only return contacts matching the email addresses; multiple email_addresses can be separated by commas.
	EmailAddresses *string `json:"-"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *LeadsListRequestExpand `json:"-"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
	IncludeRemoteFields *bool `json:"-"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-"`
	// If provided, will only return leads with this owner.
	OwnerId *string `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
	// If provided, will only return contacts matching the phone numbers; multiple phone numbers can be separated by commas.
	PhoneNumbers *string `json:"-"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-"`
}

type LeadsRemoteFieldClassesListRequest struct {
	// The pagination cursor value.
	Cursor *string `json:"-"`
	// Whether to include data that was marked as deleted by third party webhooks.
	IncludeDeletedData *bool `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
	IncludeRemoteFields *bool `json:"-"`
	// If provided, will only return remote field classes with this is_common_model_field value
	IsCommonModelField *bool `json:"-"`
	// Number of results to return per page.
	PageSize *int `json:"-"`
}

type LeadsRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand *LeadsRetrieveRequestExpand `json:"-"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
	IncludeRemoteFields *bool `json:"-"`
}

type LeadsListRequestExpand string

const (
	LeadsListRequestExpandConvertedAccount                      LeadsListRequestExpand = "converted_account"
	LeadsListRequestExpandConvertedContact                      LeadsListRequestExpand = "converted_contact"
	LeadsListRequestExpandConvertedContactConvertedAccount      LeadsListRequestExpand = "converted_contact,converted_account"
	LeadsListRequestExpandOwner                                 LeadsListRequestExpand = "owner"
	LeadsListRequestExpandOwnerConvertedAccount                 LeadsListRequestExpand = "owner,converted_account"
	LeadsListRequestExpandOwnerConvertedContact                 LeadsListRequestExpand = "owner,converted_contact"
	LeadsListRequestExpandOwnerConvertedContactConvertedAccount LeadsListRequestExpand = "owner,converted_contact,converted_account"
)

func NewLeadsListRequestExpandFromString(s string) (LeadsListRequestExpand, error) {
	switch s {
	case "converted_account":
		return LeadsListRequestExpandConvertedAccount, nil
	case "converted_contact":
		return LeadsListRequestExpandConvertedContact, nil
	case "converted_contact,converted_account":
		return LeadsListRequestExpandConvertedContactConvertedAccount, nil
	case "owner":
		return LeadsListRequestExpandOwner, nil
	case "owner,converted_account":
		return LeadsListRequestExpandOwnerConvertedAccount, nil
	case "owner,converted_contact":
		return LeadsListRequestExpandOwnerConvertedContact, nil
	case "owner,converted_contact,converted_account":
		return LeadsListRequestExpandOwnerConvertedContactConvertedAccount, nil
	}
	var t LeadsListRequestExpand
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LeadsListRequestExpand) Ptr() *LeadsListRequestExpand {
	return &l
}

type LeadsRetrieveRequestExpand string

const (
	LeadsRetrieveRequestExpandConvertedAccount                      LeadsRetrieveRequestExpand = "converted_account"
	LeadsRetrieveRequestExpandConvertedContact                      LeadsRetrieveRequestExpand = "converted_contact"
	LeadsRetrieveRequestExpandConvertedContactConvertedAccount      LeadsRetrieveRequestExpand = "converted_contact,converted_account"
	LeadsRetrieveRequestExpandOwner                                 LeadsRetrieveRequestExpand = "owner"
	LeadsRetrieveRequestExpandOwnerConvertedAccount                 LeadsRetrieveRequestExpand = "owner,converted_account"
	LeadsRetrieveRequestExpandOwnerConvertedContact                 LeadsRetrieveRequestExpand = "owner,converted_contact"
	LeadsRetrieveRequestExpandOwnerConvertedContactConvertedAccount LeadsRetrieveRequestExpand = "owner,converted_contact,converted_account"
)

func NewLeadsRetrieveRequestExpandFromString(s string) (LeadsRetrieveRequestExpand, error) {
	switch s {
	case "converted_account":
		return LeadsRetrieveRequestExpandConvertedAccount, nil
	case "converted_contact":
		return LeadsRetrieveRequestExpandConvertedContact, nil
	case "converted_contact,converted_account":
		return LeadsRetrieveRequestExpandConvertedContactConvertedAccount, nil
	case "owner":
		return LeadsRetrieveRequestExpandOwner, nil
	case "owner,converted_account":
		return LeadsRetrieveRequestExpandOwnerConvertedAccount, nil
	case "owner,converted_contact":
		return LeadsRetrieveRequestExpandOwnerConvertedContact, nil
	case "owner,converted_contact,converted_account":
		return LeadsRetrieveRequestExpandOwnerConvertedContactConvertedAccount, nil
	}
	var t LeadsRetrieveRequestExpand
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (l LeadsRetrieveRequestExpand) Ptr() *LeadsRetrieveRequestExpand {
	return &l
}
