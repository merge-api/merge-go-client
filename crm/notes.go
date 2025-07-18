// Code generated by Fern. DO NOT EDIT.

package crm

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/merge-api/merge-go-client/v2/internal"
	time "time"
)

type NoteEndpointRequest struct {
	// Whether to include debug fields (such as log file links) in the response.
	IsDebugMode *bool `json:"-" url:"is_debug_mode,omitempty"`
	// Whether or not third-party updates should be run asynchronously.
	RunAsync *bool        `json:"-" url:"run_async,omitempty"`
	Model    *NoteRequest `json:"model,omitempty" url:"-"`
}

type NotesListRequest struct {
	// If provided, will only return notes with this account.
	AccountId *string `json:"-" url:"account_id,omitempty"`
	// If provided, will only return notes with this contact.
	ContactId *string `json:"-" url:"contact_id,omitempty"`
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-" url:"created_after,omitempty"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-" url:"created_before,omitempty"`
	// The pagination cursor value.
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand []*NotesListRequestExpandItem `json:"-" url:"expand,omitempty"`
	// Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
	IncludeDeletedData *bool `json:"-" url:"include_deleted_data,omitempty"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-" url:"include_remote_data,omitempty"`
	// Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
	IncludeRemoteFields *bool `json:"-" url:"include_remote_fields,omitempty"`
	// Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
	IncludeShellData *bool `json:"-" url:"include_shell_data,omitempty"`
	// If provided, only objects synced by Merge after this date time will be returned.
	ModifiedAfter *time.Time `json:"-" url:"modified_after,omitempty"`
	// If provided, only objects synced by Merge before this date time will be returned.
	ModifiedBefore *time.Time `json:"-" url:"modified_before,omitempty"`
	// If provided, will only return notes with this opportunity.
	OpportunityId *string `json:"-" url:"opportunity_id,omitempty"`
	// If provided, will only return notes with this owner.
	OwnerId *string `json:"-" url:"owner_id,omitempty"`
	// Number of results to return per page.
	PageSize *int `json:"-" url:"page_size,omitempty"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-" url:"remote_id,omitempty"`
}

type NotesRemoteFieldClassesListRequest struct {
	// The pagination cursor value.
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
	IncludeDeletedData *bool `json:"-" url:"include_deleted_data,omitempty"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-" url:"include_remote_data,omitempty"`
	// Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
	IncludeRemoteFields *bool `json:"-" url:"include_remote_fields,omitempty"`
	// Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
	IncludeShellData *bool `json:"-" url:"include_shell_data,omitempty"`
	// If provided, will only return remote field classes with this is_common_model_field value
	IsCommonModelField *bool `json:"-" url:"is_common_model_field,omitempty"`
	// If provided, will only return remote fields classes with this is_custom value
	IsCustom *bool `json:"-" url:"is_custom,omitempty"`
	// Number of results to return per page.
	PageSize *int `json:"-" url:"page_size,omitempty"`
}

type NotesRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand []*NotesRetrieveRequestExpandItem `json:"-" url:"expand,omitempty"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-" url:"include_remote_data,omitempty"`
	// Whether to include all remote fields, including fields that Merge did not map to common models, in a normalized format.
	IncludeRemoteFields *bool `json:"-" url:"include_remote_fields,omitempty"`
	// Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
	IncludeShellData *bool `json:"-" url:"include_shell_data,omitempty"`
}

type NotesListRequestExpandItem string

const (
	NotesListRequestExpandItemAccount     NotesListRequestExpandItem = "account"
	NotesListRequestExpandItemContact     NotesListRequestExpandItem = "contact"
	NotesListRequestExpandItemOpportunity NotesListRequestExpandItem = "opportunity"
	NotesListRequestExpandItemOwner       NotesListRequestExpandItem = "owner"
)

func NewNotesListRequestExpandItemFromString(s string) (NotesListRequestExpandItem, error) {
	switch s {
	case "account":
		return NotesListRequestExpandItemAccount, nil
	case "contact":
		return NotesListRequestExpandItemContact, nil
	case "opportunity":
		return NotesListRequestExpandItemOpportunity, nil
	case "owner":
		return NotesListRequestExpandItemOwner, nil
	}
	var t NotesListRequestExpandItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (n NotesListRequestExpandItem) Ptr() *NotesListRequestExpandItem {
	return &n
}

type NotesRetrieveRequestExpandItem string

const (
	NotesRetrieveRequestExpandItemAccount     NotesRetrieveRequestExpandItem = "account"
	NotesRetrieveRequestExpandItemContact     NotesRetrieveRequestExpandItem = "contact"
	NotesRetrieveRequestExpandItemOpportunity NotesRetrieveRequestExpandItem = "opportunity"
	NotesRetrieveRequestExpandItemOwner       NotesRetrieveRequestExpandItem = "owner"
)

func NewNotesRetrieveRequestExpandItemFromString(s string) (NotesRetrieveRequestExpandItem, error) {
	switch s {
	case "account":
		return NotesRetrieveRequestExpandItemAccount, nil
	case "contact":
		return NotesRetrieveRequestExpandItemContact, nil
	case "opportunity":
		return NotesRetrieveRequestExpandItemOpportunity, nil
	case "owner":
		return NotesRetrieveRequestExpandItemOwner, nil
	}
	var t NotesRetrieveRequestExpandItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (n NotesRetrieveRequestExpandItem) Ptr() *NotesRetrieveRequestExpandItem {
	return &n
}

// # The Note Object
// ### Description
// The `Note` object is used to represent a note on another object.
// ### Usage Example
// TODO
type Note struct {
	Id *string `json:"id,omitempty" url:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId *string `json:"remote_id,omitempty" url:"remote_id,omitempty"`
	// The datetime that this object was created by Merge.
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The datetime that this object was modified by Merge.
	ModifiedAt *time.Time `json:"modified_at,omitempty" url:"modified_at,omitempty"`
	// The note's owner.
	Owner *NoteOwner `json:"owner,omitempty" url:"owner,omitempty"`
	// The note's content.
	Content *string `json:"content,omitempty" url:"content,omitempty"`
	// The note's contact.
	Contact *NoteContact `json:"contact,omitempty" url:"contact,omitempty"`
	// The note's account.
	Account *NoteAccount `json:"account,omitempty" url:"account,omitempty"`
	// The note's opportunity.
	Opportunity *NoteOpportunity `json:"opportunity,omitempty" url:"opportunity,omitempty"`
	// When the third party's lead was updated.
	RemoteUpdatedAt *time.Time `json:"remote_updated_at,omitempty" url:"remote_updated_at,omitempty"`
	// When the third party's lead was created.
	RemoteCreatedAt *time.Time `json:"remote_created_at,omitempty" url:"remote_created_at,omitempty"`
	// Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
	RemoteWasDeleted *bool                  `json:"remote_was_deleted,omitempty" url:"remote_was_deleted,omitempty"`
	FieldMappings    map[string]interface{} `json:"field_mappings,omitempty" url:"field_mappings,omitempty"`
	RemoteData       []*RemoteData          `json:"remote_data,omitempty" url:"remote_data,omitempty"`
	RemoteFields     []*RemoteField         `json:"remote_fields,omitempty" url:"remote_fields,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (n *Note) GetId() *string {
	if n == nil {
		return nil
	}
	return n.Id
}

func (n *Note) GetRemoteId() *string {
	if n == nil {
		return nil
	}
	return n.RemoteId
}

func (n *Note) GetCreatedAt() *time.Time {
	if n == nil {
		return nil
	}
	return n.CreatedAt
}

func (n *Note) GetModifiedAt() *time.Time {
	if n == nil {
		return nil
	}
	return n.ModifiedAt
}

func (n *Note) GetOwner() *NoteOwner {
	if n == nil {
		return nil
	}
	return n.Owner
}

func (n *Note) GetContent() *string {
	if n == nil {
		return nil
	}
	return n.Content
}

func (n *Note) GetContact() *NoteContact {
	if n == nil {
		return nil
	}
	return n.Contact
}

func (n *Note) GetAccount() *NoteAccount {
	if n == nil {
		return nil
	}
	return n.Account
}

func (n *Note) GetOpportunity() *NoteOpportunity {
	if n == nil {
		return nil
	}
	return n.Opportunity
}

func (n *Note) GetRemoteUpdatedAt() *time.Time {
	if n == nil {
		return nil
	}
	return n.RemoteUpdatedAt
}

func (n *Note) GetRemoteCreatedAt() *time.Time {
	if n == nil {
		return nil
	}
	return n.RemoteCreatedAt
}

func (n *Note) GetRemoteWasDeleted() *bool {
	if n == nil {
		return nil
	}
	return n.RemoteWasDeleted
}

func (n *Note) GetFieldMappings() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.FieldMappings
}

func (n *Note) GetRemoteData() []*RemoteData {
	if n == nil {
		return nil
	}
	return n.RemoteData
}

func (n *Note) GetRemoteFields() []*RemoteField {
	if n == nil {
		return nil
	}
	return n.RemoteFields
}

func (n *Note) GetExtraProperties() map[string]interface{} {
	return n.extraProperties
}

func (n *Note) UnmarshalJSON(data []byte) error {
	type embed Note
	var unmarshaler = struct {
		embed
		CreatedAt       *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt      *internal.DateTime `json:"modified_at,omitempty"`
		RemoteUpdatedAt *internal.DateTime `json:"remote_updated_at,omitempty"`
		RemoteCreatedAt *internal.DateTime `json:"remote_created_at,omitempty"`
	}{
		embed: embed(*n),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*n = Note(unmarshaler.embed)
	n.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	n.ModifiedAt = unmarshaler.ModifiedAt.TimePtr()
	n.RemoteUpdatedAt = unmarshaler.RemoteUpdatedAt.TimePtr()
	n.RemoteCreatedAt = unmarshaler.RemoteCreatedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	n.rawJSON = json.RawMessage(data)
	return nil
}

func (n *Note) MarshalJSON() ([]byte, error) {
	type embed Note
	var marshaler = struct {
		embed
		CreatedAt       *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt      *internal.DateTime `json:"modified_at,omitempty"`
		RemoteUpdatedAt *internal.DateTime `json:"remote_updated_at,omitempty"`
		RemoteCreatedAt *internal.DateTime `json:"remote_created_at,omitempty"`
	}{
		embed:           embed(*n),
		CreatedAt:       internal.NewOptionalDateTime(n.CreatedAt),
		ModifiedAt:      internal.NewOptionalDateTime(n.ModifiedAt),
		RemoteUpdatedAt: internal.NewOptionalDateTime(n.RemoteUpdatedAt),
		RemoteCreatedAt: internal.NewOptionalDateTime(n.RemoteCreatedAt),
	}
	return json.Marshal(marshaler)
}

func (n *Note) String() string {
	if len(n.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(n.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}

// The note's account.
type NoteAccount struct {
	String  string
	Account *Account

	typ string
}

func (n *NoteAccount) GetString() string {
	if n == nil {
		return ""
	}
	return n.String
}

func (n *NoteAccount) GetAccount() *Account {
	if n == nil {
		return nil
	}
	return n.Account
}

func (n *NoteAccount) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		n.typ = "String"
		n.String = valueString
		return nil
	}
	valueAccount := new(Account)
	if err := json.Unmarshal(data, &valueAccount); err == nil {
		n.typ = "Account"
		n.Account = valueAccount
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, n)
}

func (n NoteAccount) MarshalJSON() ([]byte, error) {
	if n.typ == "String" || n.String != "" {
		return json.Marshal(n.String)
	}
	if n.typ == "Account" || n.Account != nil {
		return json.Marshal(n.Account)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", n)
}

type NoteAccountVisitor interface {
	VisitString(string) error
	VisitAccount(*Account) error
}

func (n *NoteAccount) Accept(visitor NoteAccountVisitor) error {
	if n.typ == "String" || n.String != "" {
		return visitor.VisitString(n.String)
	}
	if n.typ == "Account" || n.Account != nil {
		return visitor.VisitAccount(n.Account)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", n)
}

// The note's contact.
type NoteContact struct {
	String  string
	Contact *Contact

	typ string
}

func (n *NoteContact) GetString() string {
	if n == nil {
		return ""
	}
	return n.String
}

func (n *NoteContact) GetContact() *Contact {
	if n == nil {
		return nil
	}
	return n.Contact
}

func (n *NoteContact) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		n.typ = "String"
		n.String = valueString
		return nil
	}
	valueContact := new(Contact)
	if err := json.Unmarshal(data, &valueContact); err == nil {
		n.typ = "Contact"
		n.Contact = valueContact
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, n)
}

func (n NoteContact) MarshalJSON() ([]byte, error) {
	if n.typ == "String" || n.String != "" {
		return json.Marshal(n.String)
	}
	if n.typ == "Contact" || n.Contact != nil {
		return json.Marshal(n.Contact)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", n)
}

type NoteContactVisitor interface {
	VisitString(string) error
	VisitContact(*Contact) error
}

func (n *NoteContact) Accept(visitor NoteContactVisitor) error {
	if n.typ == "String" || n.String != "" {
		return visitor.VisitString(n.String)
	}
	if n.typ == "Contact" || n.Contact != nil {
		return visitor.VisitContact(n.Contact)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", n)
}

// The note's opportunity.
type NoteOpportunity struct {
	String      string
	Opportunity *Opportunity

	typ string
}

func (n *NoteOpportunity) GetString() string {
	if n == nil {
		return ""
	}
	return n.String
}

func (n *NoteOpportunity) GetOpportunity() *Opportunity {
	if n == nil {
		return nil
	}
	return n.Opportunity
}

func (n *NoteOpportunity) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		n.typ = "String"
		n.String = valueString
		return nil
	}
	valueOpportunity := new(Opportunity)
	if err := json.Unmarshal(data, &valueOpportunity); err == nil {
		n.typ = "Opportunity"
		n.Opportunity = valueOpportunity
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, n)
}

func (n NoteOpportunity) MarshalJSON() ([]byte, error) {
	if n.typ == "String" || n.String != "" {
		return json.Marshal(n.String)
	}
	if n.typ == "Opportunity" || n.Opportunity != nil {
		return json.Marshal(n.Opportunity)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", n)
}

type NoteOpportunityVisitor interface {
	VisitString(string) error
	VisitOpportunity(*Opportunity) error
}

func (n *NoteOpportunity) Accept(visitor NoteOpportunityVisitor) error {
	if n.typ == "String" || n.String != "" {
		return visitor.VisitString(n.String)
	}
	if n.typ == "Opportunity" || n.Opportunity != nil {
		return visitor.VisitOpportunity(n.Opportunity)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", n)
}

// The note's owner.
type NoteOwner struct {
	String string
	User   *User

	typ string
}

func (n *NoteOwner) GetString() string {
	if n == nil {
		return ""
	}
	return n.String
}

func (n *NoteOwner) GetUser() *User {
	if n == nil {
		return nil
	}
	return n.User
}

func (n *NoteOwner) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		n.typ = "String"
		n.String = valueString
		return nil
	}
	valueUser := new(User)
	if err := json.Unmarshal(data, &valueUser); err == nil {
		n.typ = "User"
		n.User = valueUser
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, n)
}

func (n NoteOwner) MarshalJSON() ([]byte, error) {
	if n.typ == "String" || n.String != "" {
		return json.Marshal(n.String)
	}
	if n.typ == "User" || n.User != nil {
		return json.Marshal(n.User)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", n)
}

type NoteOwnerVisitor interface {
	VisitString(string) error
	VisitUser(*User) error
}

func (n *NoteOwner) Accept(visitor NoteOwnerVisitor) error {
	if n.typ == "String" || n.String != "" {
		return visitor.VisitString(n.String)
	}
	if n.typ == "User" || n.User != nil {
		return visitor.VisitUser(n.User)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", n)
}

// # The Note Object
// ### Description
// The `Note` object is used to represent a note on another object.
// ### Usage Example
// TODO
type NoteRequest struct {
	// The note's owner.
	Owner *NoteRequestOwner `json:"owner,omitempty" url:"owner,omitempty"`
	// The note's content.
	Content *string `json:"content,omitempty" url:"content,omitempty"`
	// The note's contact.
	Contact *NoteRequestContact `json:"contact,omitempty" url:"contact,omitempty"`
	// The note's account.
	Account *NoteRequestAccount `json:"account,omitempty" url:"account,omitempty"`
	// The note's opportunity.
	Opportunity         *NoteRequestOpportunity `json:"opportunity,omitempty" url:"opportunity,omitempty"`
	IntegrationParams   map[string]interface{}  `json:"integration_params,omitempty" url:"integration_params,omitempty"`
	LinkedAccountParams map[string]interface{}  `json:"linked_account_params,omitempty" url:"linked_account_params,omitempty"`
	RemoteFields        []*RemoteFieldRequest   `json:"remote_fields,omitempty" url:"remote_fields,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (n *NoteRequest) GetOwner() *NoteRequestOwner {
	if n == nil {
		return nil
	}
	return n.Owner
}

func (n *NoteRequest) GetContent() *string {
	if n == nil {
		return nil
	}
	return n.Content
}

func (n *NoteRequest) GetContact() *NoteRequestContact {
	if n == nil {
		return nil
	}
	return n.Contact
}

func (n *NoteRequest) GetAccount() *NoteRequestAccount {
	if n == nil {
		return nil
	}
	return n.Account
}

func (n *NoteRequest) GetOpportunity() *NoteRequestOpportunity {
	if n == nil {
		return nil
	}
	return n.Opportunity
}

func (n *NoteRequest) GetIntegrationParams() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.IntegrationParams
}

func (n *NoteRequest) GetLinkedAccountParams() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.LinkedAccountParams
}

func (n *NoteRequest) GetRemoteFields() []*RemoteFieldRequest {
	if n == nil {
		return nil
	}
	return n.RemoteFields
}

func (n *NoteRequest) GetExtraProperties() map[string]interface{} {
	return n.extraProperties
}

func (n *NoteRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler NoteRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NoteRequest(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	n.rawJSON = json.RawMessage(data)
	return nil
}

func (n *NoteRequest) String() string {
	if len(n.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(n.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}

// The note's account.
type NoteRequestAccount struct {
	String  string
	Account *Account

	typ string
}

func (n *NoteRequestAccount) GetString() string {
	if n == nil {
		return ""
	}
	return n.String
}

func (n *NoteRequestAccount) GetAccount() *Account {
	if n == nil {
		return nil
	}
	return n.Account
}

func (n *NoteRequestAccount) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		n.typ = "String"
		n.String = valueString
		return nil
	}
	valueAccount := new(Account)
	if err := json.Unmarshal(data, &valueAccount); err == nil {
		n.typ = "Account"
		n.Account = valueAccount
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, n)
}

func (n NoteRequestAccount) MarshalJSON() ([]byte, error) {
	if n.typ == "String" || n.String != "" {
		return json.Marshal(n.String)
	}
	if n.typ == "Account" || n.Account != nil {
		return json.Marshal(n.Account)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", n)
}

type NoteRequestAccountVisitor interface {
	VisitString(string) error
	VisitAccount(*Account) error
}

func (n *NoteRequestAccount) Accept(visitor NoteRequestAccountVisitor) error {
	if n.typ == "String" || n.String != "" {
		return visitor.VisitString(n.String)
	}
	if n.typ == "Account" || n.Account != nil {
		return visitor.VisitAccount(n.Account)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", n)
}

// The note's contact.
type NoteRequestContact struct {
	String  string
	Contact *Contact

	typ string
}

func (n *NoteRequestContact) GetString() string {
	if n == nil {
		return ""
	}
	return n.String
}

func (n *NoteRequestContact) GetContact() *Contact {
	if n == nil {
		return nil
	}
	return n.Contact
}

func (n *NoteRequestContact) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		n.typ = "String"
		n.String = valueString
		return nil
	}
	valueContact := new(Contact)
	if err := json.Unmarshal(data, &valueContact); err == nil {
		n.typ = "Contact"
		n.Contact = valueContact
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, n)
}

func (n NoteRequestContact) MarshalJSON() ([]byte, error) {
	if n.typ == "String" || n.String != "" {
		return json.Marshal(n.String)
	}
	if n.typ == "Contact" || n.Contact != nil {
		return json.Marshal(n.Contact)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", n)
}

type NoteRequestContactVisitor interface {
	VisitString(string) error
	VisitContact(*Contact) error
}

func (n *NoteRequestContact) Accept(visitor NoteRequestContactVisitor) error {
	if n.typ == "String" || n.String != "" {
		return visitor.VisitString(n.String)
	}
	if n.typ == "Contact" || n.Contact != nil {
		return visitor.VisitContact(n.Contact)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", n)
}

// The note's opportunity.
type NoteRequestOpportunity struct {
	String      string
	Opportunity *Opportunity

	typ string
}

func (n *NoteRequestOpportunity) GetString() string {
	if n == nil {
		return ""
	}
	return n.String
}

func (n *NoteRequestOpportunity) GetOpportunity() *Opportunity {
	if n == nil {
		return nil
	}
	return n.Opportunity
}

func (n *NoteRequestOpportunity) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		n.typ = "String"
		n.String = valueString
		return nil
	}
	valueOpportunity := new(Opportunity)
	if err := json.Unmarshal(data, &valueOpportunity); err == nil {
		n.typ = "Opportunity"
		n.Opportunity = valueOpportunity
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, n)
}

func (n NoteRequestOpportunity) MarshalJSON() ([]byte, error) {
	if n.typ == "String" || n.String != "" {
		return json.Marshal(n.String)
	}
	if n.typ == "Opportunity" || n.Opportunity != nil {
		return json.Marshal(n.Opportunity)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", n)
}

type NoteRequestOpportunityVisitor interface {
	VisitString(string) error
	VisitOpportunity(*Opportunity) error
}

func (n *NoteRequestOpportunity) Accept(visitor NoteRequestOpportunityVisitor) error {
	if n.typ == "String" || n.String != "" {
		return visitor.VisitString(n.String)
	}
	if n.typ == "Opportunity" || n.Opportunity != nil {
		return visitor.VisitOpportunity(n.Opportunity)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", n)
}

// The note's owner.
type NoteRequestOwner struct {
	String string
	User   *User

	typ string
}

func (n *NoteRequestOwner) GetString() string {
	if n == nil {
		return ""
	}
	return n.String
}

func (n *NoteRequestOwner) GetUser() *User {
	if n == nil {
		return nil
	}
	return n.User
}

func (n *NoteRequestOwner) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		n.typ = "String"
		n.String = valueString
		return nil
	}
	valueUser := new(User)
	if err := json.Unmarshal(data, &valueUser); err == nil {
		n.typ = "User"
		n.User = valueUser
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, n)
}

func (n NoteRequestOwner) MarshalJSON() ([]byte, error) {
	if n.typ == "String" || n.String != "" {
		return json.Marshal(n.String)
	}
	if n.typ == "User" || n.User != nil {
		return json.Marshal(n.User)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", n)
}

type NoteRequestOwnerVisitor interface {
	VisitString(string) error
	VisitUser(*User) error
}

func (n *NoteRequestOwner) Accept(visitor NoteRequestOwnerVisitor) error {
	if n.typ == "String" || n.String != "" {
		return visitor.VisitString(n.String)
	}
	if n.typ == "User" || n.User != nil {
		return visitor.VisitUser(n.User)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", n)
}

type NoteResponse struct {
	Model    *Note                       `json:"model" url:"model"`
	Warnings []*WarningValidationProblem `json:"warnings" url:"warnings"`
	Errors   []*ErrorValidationProblem   `json:"errors" url:"errors"`
	Logs     []*DebugModeLog             `json:"logs,omitempty" url:"logs,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (n *NoteResponse) GetModel() *Note {
	if n == nil {
		return nil
	}
	return n.Model
}

func (n *NoteResponse) GetWarnings() []*WarningValidationProblem {
	if n == nil {
		return nil
	}
	return n.Warnings
}

func (n *NoteResponse) GetErrors() []*ErrorValidationProblem {
	if n == nil {
		return nil
	}
	return n.Errors
}

func (n *NoteResponse) GetLogs() []*DebugModeLog {
	if n == nil {
		return nil
	}
	return n.Logs
}

func (n *NoteResponse) GetExtraProperties() map[string]interface{} {
	return n.extraProperties
}

func (n *NoteResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler NoteResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*n = NoteResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	n.rawJSON = json.RawMessage(data)
	return nil
}

func (n *NoteResponse) String() string {
	if len(n.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(n.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(n); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", n)
}

type PaginatedNoteList struct {
	Next     *string `json:"next,omitempty" url:"next,omitempty"`
	Previous *string `json:"previous,omitempty" url:"previous,omitempty"`
	Results  []*Note `json:"results,omitempty" url:"results,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PaginatedNoteList) GetNext() *string {
	if p == nil {
		return nil
	}
	return p.Next
}

func (p *PaginatedNoteList) GetPrevious() *string {
	if p == nil {
		return nil
	}
	return p.Previous
}

func (p *PaginatedNoteList) GetResults() []*Note {
	if p == nil {
		return nil
	}
	return p.Results
}

func (p *PaginatedNoteList) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaginatedNoteList) UnmarshalJSON(data []byte) error {
	type unmarshaler PaginatedNoteList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PaginatedNoteList(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaginatedNoteList) String() string {
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
