// Code generated by Fern. DO NOT EDIT.

package hris

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/merge-api/merge-go-client/v2/internal"
	time "time"
)

type EmployeePayrollRunsListRequest struct {
	// If provided, will only return objects created after this datetime.
	CreatedAfter *time.Time `json:"-" url:"created_after,omitempty"`
	// If provided, will only return objects created before this datetime.
	CreatedBefore *time.Time `json:"-" url:"created_before,omitempty"`
	// The pagination cursor value.
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// If provided, will only return employee payroll runs for this employee.
	EmployeeId *string `json:"-" url:"employee_id,omitempty"`
	// If provided, will only return employee payroll runs ended after this datetime.
	EndedAfter *time.Time `json:"-" url:"ended_after,omitempty"`
	// If provided, will only return employee payroll runs ended before this datetime.
	EndedBefore *time.Time `json:"-" url:"ended_before,omitempty"`
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand []*EmployeePayrollRunsListRequestExpandItem `json:"-" url:"expand,omitempty"`
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
	// If provided, will only return employee payroll runs for this employee.
	PayrollRunId *string `json:"-" url:"payroll_run_id,omitempty"`
	// The API provider's ID for the given object.
	RemoteId *string `json:"-" url:"remote_id,omitempty"`
	// If provided, will only return employee payroll runs started after this datetime.
	StartedAfter *time.Time `json:"-" url:"started_after,omitempty"`
	// If provided, will only return employee payroll runs started before this datetime.
	StartedBefore *time.Time `json:"-" url:"started_before,omitempty"`
}

type EmployeePayrollRunsRetrieveRequest struct {
	// Which relations should be returned in expanded form. Multiple relation names should be comma separated without spaces.
	Expand []*EmployeePayrollRunsRetrieveRequestExpandItem `json:"-" url:"expand,omitempty"`
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-" url:"include_remote_data,omitempty"`
	// Whether to include shell records. Shell records are empty records (they may contain some metadata but all other fields are null).
	IncludeShellData *bool `json:"-" url:"include_shell_data,omitempty"`
}

type EmployeePayrollRunsListRequestExpandItem string

const (
	EmployeePayrollRunsListRequestExpandItemEmployee   EmployeePayrollRunsListRequestExpandItem = "employee"
	EmployeePayrollRunsListRequestExpandItemPayrollRun EmployeePayrollRunsListRequestExpandItem = "payroll_run"
)

func NewEmployeePayrollRunsListRequestExpandItemFromString(s string) (EmployeePayrollRunsListRequestExpandItem, error) {
	switch s {
	case "employee":
		return EmployeePayrollRunsListRequestExpandItemEmployee, nil
	case "payroll_run":
		return EmployeePayrollRunsListRequestExpandItemPayrollRun, nil
	}
	var t EmployeePayrollRunsListRequestExpandItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EmployeePayrollRunsListRequestExpandItem) Ptr() *EmployeePayrollRunsListRequestExpandItem {
	return &e
}

type EmployeePayrollRunsRetrieveRequestExpandItem string

const (
	EmployeePayrollRunsRetrieveRequestExpandItemEmployee   EmployeePayrollRunsRetrieveRequestExpandItem = "employee"
	EmployeePayrollRunsRetrieveRequestExpandItemPayrollRun EmployeePayrollRunsRetrieveRequestExpandItem = "payroll_run"
)

func NewEmployeePayrollRunsRetrieveRequestExpandItemFromString(s string) (EmployeePayrollRunsRetrieveRequestExpandItem, error) {
	switch s {
	case "employee":
		return EmployeePayrollRunsRetrieveRequestExpandItemEmployee, nil
	case "payroll_run":
		return EmployeePayrollRunsRetrieveRequestExpandItemPayrollRun, nil
	}
	var t EmployeePayrollRunsRetrieveRequestExpandItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EmployeePayrollRunsRetrieveRequestExpandItem) Ptr() *EmployeePayrollRunsRetrieveRequestExpandItem {
	return &e
}

// # The Deduction Object
// ### Description
// The `Deduction` object is used to represent an array of the wages withheld from total earnings for the purpose of paying taxes.
//
// ### Usage Example
// Fetch from the `LIST Deductions` endpoint and filter by `ID` to show all deductions.
type Deduction struct {
	Id *string `json:"id,omitempty" url:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId *string `json:"remote_id,omitempty" url:"remote_id,omitempty"`
	// The datetime that this object was created by Merge.
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The datetime that this object was modified by Merge.
	ModifiedAt         *time.Time `json:"modified_at,omitempty" url:"modified_at,omitempty"`
	EmployeePayrollRun *string    `json:"employee_payroll_run,omitempty" url:"employee_payroll_run,omitempty"`
	// The deduction's name.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The amount of money that is withheld from an employee's gross pay by the employee.
	EmployeeDeduction *float64 `json:"employee_deduction,omitempty" url:"employee_deduction,omitempty"`
	// The amount of money that is withheld on behalf of an employee by the company.
	CompanyDeduction *float64 `json:"company_deduction,omitempty" url:"company_deduction,omitempty"`
	// Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
	RemoteWasDeleted *bool                  `json:"remote_was_deleted,omitempty" url:"remote_was_deleted,omitempty"`
	FieldMappings    map[string]interface{} `json:"field_mappings,omitempty" url:"field_mappings,omitempty"`
	RemoteData       []*RemoteData          `json:"remote_data,omitempty" url:"remote_data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (d *Deduction) GetId() *string {
	if d == nil {
		return nil
	}
	return d.Id
}

func (d *Deduction) GetRemoteId() *string {
	if d == nil {
		return nil
	}
	return d.RemoteId
}

func (d *Deduction) GetCreatedAt() *time.Time {
	if d == nil {
		return nil
	}
	return d.CreatedAt
}

func (d *Deduction) GetModifiedAt() *time.Time {
	if d == nil {
		return nil
	}
	return d.ModifiedAt
}

func (d *Deduction) GetEmployeePayrollRun() *string {
	if d == nil {
		return nil
	}
	return d.EmployeePayrollRun
}

func (d *Deduction) GetName() *string {
	if d == nil {
		return nil
	}
	return d.Name
}

func (d *Deduction) GetEmployeeDeduction() *float64 {
	if d == nil {
		return nil
	}
	return d.EmployeeDeduction
}

func (d *Deduction) GetCompanyDeduction() *float64 {
	if d == nil {
		return nil
	}
	return d.CompanyDeduction
}

func (d *Deduction) GetRemoteWasDeleted() *bool {
	if d == nil {
		return nil
	}
	return d.RemoteWasDeleted
}

func (d *Deduction) GetFieldMappings() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.FieldMappings
}

func (d *Deduction) GetRemoteData() []*RemoteData {
	if d == nil {
		return nil
	}
	return d.RemoteData
}

func (d *Deduction) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *Deduction) UnmarshalJSON(data []byte) error {
	type embed Deduction
	var unmarshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
	}{
		embed: embed(*d),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*d = Deduction(unmarshaler.embed)
	d.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	d.ModifiedAt = unmarshaler.ModifiedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	d.rawJSON = json.RawMessage(data)
	return nil
}

func (d *Deduction) MarshalJSON() ([]byte, error) {
	type embed Deduction
	var marshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
	}{
		embed:      embed(*d),
		CreatedAt:  internal.NewOptionalDateTime(d.CreatedAt),
		ModifiedAt: internal.NewOptionalDateTime(d.ModifiedAt),
	}
	return json.Marshal(marshaler)
}

func (d *Deduction) String() string {
	if len(d.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(d.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

// # The Earning Object
// ### Description
// The `Earning` object is used to represent an array of different compensations that an employee receives within specific wage categories.
//
// ### Usage Example
// Fetch from the `LIST Earnings` endpoint and filter by `ID` to show all earnings.
type Earning struct {
	Id *string `json:"id,omitempty" url:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId *string `json:"remote_id,omitempty" url:"remote_id,omitempty"`
	// The datetime that this object was created by Merge.
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The datetime that this object was modified by Merge.
	ModifiedAt         *time.Time `json:"modified_at,omitempty" url:"modified_at,omitempty"`
	EmployeePayrollRun *string    `json:"employee_payroll_run,omitempty" url:"employee_payroll_run,omitempty"`
	// The amount earned.
	Amount *float64 `json:"amount,omitempty" url:"amount,omitempty"`
	// The type of earning.
	//
	// * `SALARY` - SALARY
	// * `REIMBURSEMENT` - REIMBURSEMENT
	// * `OVERTIME` - OVERTIME
	// * `BONUS` - BONUS
	Type *EarningType `json:"type,omitempty" url:"type,omitempty"`
	// Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
	RemoteWasDeleted *bool                  `json:"remote_was_deleted,omitempty" url:"remote_was_deleted,omitempty"`
	FieldMappings    map[string]interface{} `json:"field_mappings,omitempty" url:"field_mappings,omitempty"`
	RemoteData       []*RemoteData          `json:"remote_data,omitempty" url:"remote_data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *Earning) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *Earning) GetRemoteId() *string {
	if e == nil {
		return nil
	}
	return e.RemoteId
}

func (e *Earning) GetCreatedAt() *time.Time {
	if e == nil {
		return nil
	}
	return e.CreatedAt
}

func (e *Earning) GetModifiedAt() *time.Time {
	if e == nil {
		return nil
	}
	return e.ModifiedAt
}

func (e *Earning) GetEmployeePayrollRun() *string {
	if e == nil {
		return nil
	}
	return e.EmployeePayrollRun
}

func (e *Earning) GetAmount() *float64 {
	if e == nil {
		return nil
	}
	return e.Amount
}

func (e *Earning) GetType() *EarningType {
	if e == nil {
		return nil
	}
	return e.Type
}

func (e *Earning) GetRemoteWasDeleted() *bool {
	if e == nil {
		return nil
	}
	return e.RemoteWasDeleted
}

func (e *Earning) GetFieldMappings() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.FieldMappings
}

func (e *Earning) GetRemoteData() []*RemoteData {
	if e == nil {
		return nil
	}
	return e.RemoteData
}

func (e *Earning) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *Earning) UnmarshalJSON(data []byte) error {
	type embed Earning
	var unmarshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
	}{
		embed: embed(*e),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*e = Earning(unmarshaler.embed)
	e.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	e.ModifiedAt = unmarshaler.ModifiedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *Earning) MarshalJSON() ([]byte, error) {
	type embed Earning
	var marshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
	}{
		embed:      embed(*e),
		CreatedAt:  internal.NewOptionalDateTime(e.CreatedAt),
		ModifiedAt: internal.NewOptionalDateTime(e.ModifiedAt),
	}
	return json.Marshal(marshaler)
}

func (e *Earning) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

// The type of earning.
//
// * `SALARY` - SALARY
// * `REIMBURSEMENT` - REIMBURSEMENT
// * `OVERTIME` - OVERTIME
// * `BONUS` - BONUS
type EarningType struct {
	EarningTypeEnum EarningTypeEnum
	String          string

	typ string
}

func (e *EarningType) GetEarningTypeEnum() EarningTypeEnum {
	if e == nil {
		return ""
	}
	return e.EarningTypeEnum
}

func (e *EarningType) GetString() string {
	if e == nil {
		return ""
	}
	return e.String
}

func (e *EarningType) UnmarshalJSON(data []byte) error {
	var valueEarningTypeEnum EarningTypeEnum
	if err := json.Unmarshal(data, &valueEarningTypeEnum); err == nil {
		e.typ = "EarningTypeEnum"
		e.EarningTypeEnum = valueEarningTypeEnum
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		e.typ = "String"
		e.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, e)
}

func (e EarningType) MarshalJSON() ([]byte, error) {
	if e.typ == "EarningTypeEnum" || e.EarningTypeEnum != "" {
		return json.Marshal(e.EarningTypeEnum)
	}
	if e.typ == "String" || e.String != "" {
		return json.Marshal(e.String)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", e)
}

type EarningTypeVisitor interface {
	VisitEarningTypeEnum(EarningTypeEnum) error
	VisitString(string) error
}

func (e *EarningType) Accept(visitor EarningTypeVisitor) error {
	if e.typ == "EarningTypeEnum" || e.EarningTypeEnum != "" {
		return visitor.VisitEarningTypeEnum(e.EarningTypeEnum)
	}
	if e.typ == "String" || e.String != "" {
		return visitor.VisitString(e.String)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", e)
}

// * `SALARY` - SALARY
// * `REIMBURSEMENT` - REIMBURSEMENT
// * `OVERTIME` - OVERTIME
// * `BONUS` - BONUS
type EarningTypeEnum string

const (
	EarningTypeEnumSalary        EarningTypeEnum = "SALARY"
	EarningTypeEnumReimbursement EarningTypeEnum = "REIMBURSEMENT"
	EarningTypeEnumOvertime      EarningTypeEnum = "OVERTIME"
	EarningTypeEnumBonus         EarningTypeEnum = "BONUS"
)

func NewEarningTypeEnumFromString(s string) (EarningTypeEnum, error) {
	switch s {
	case "SALARY":
		return EarningTypeEnumSalary, nil
	case "REIMBURSEMENT":
		return EarningTypeEnumReimbursement, nil
	case "OVERTIME":
		return EarningTypeEnumOvertime, nil
	case "BONUS":
		return EarningTypeEnumBonus, nil
	}
	var t EarningTypeEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EarningTypeEnum) Ptr() *EarningTypeEnum {
	return &e
}

// # The EmployeePayrollRun Object
// ### Description
// The `EmployeePayrollRun` object is used to represent an employee's pay statement for a specific payroll run.
//
// ### Usage Example
// Fetch from the `LIST EmployeePayrollRun` endpoint and filter by `ID` to show all employee payroll runs.
type EmployeePayrollRun struct {
	Id *string `json:"id,omitempty" url:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId *string `json:"remote_id,omitempty" url:"remote_id,omitempty"`
	// The datetime that this object was created by Merge.
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The datetime that this object was modified by Merge.
	ModifiedAt *time.Time `json:"modified_at,omitempty" url:"modified_at,omitempty"`
	// The employee whose payroll is being run.
	Employee *EmployeePayrollRunEmployee `json:"employee,omitempty" url:"employee,omitempty"`
	// The payroll being run.
	PayrollRun *EmployeePayrollRunPayrollRun `json:"payroll_run,omitempty" url:"payroll_run,omitempty"`
	// The total earnings throughout a given period for an employee before any deductions are made.
	GrossPay *float64 `json:"gross_pay,omitempty" url:"gross_pay,omitempty"`
	// The take-home pay throughout a given period for an employee after deductions are made.
	NetPay *float64 `json:"net_pay,omitempty" url:"net_pay,omitempty"`
	// The day and time the payroll run started.
	StartDate *time.Time `json:"start_date,omitempty" url:"start_date,omitempty"`
	// The day and time the payroll run ended.
	EndDate *time.Time `json:"end_date,omitempty" url:"end_date,omitempty"`
	// The day and time the payroll run was checked.
	CheckDate  *time.Time   `json:"check_date,omitempty" url:"check_date,omitempty"`
	Earnings   []*Earning   `json:"earnings,omitempty" url:"earnings,omitempty"`
	Deductions []*Deduction `json:"deductions,omitempty" url:"deductions,omitempty"`
	Taxes      []*Tax       `json:"taxes,omitempty" url:"taxes,omitempty"`
	// Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
	RemoteWasDeleted *bool                  `json:"remote_was_deleted,omitempty" url:"remote_was_deleted,omitempty"`
	FieldMappings    map[string]interface{} `json:"field_mappings,omitempty" url:"field_mappings,omitempty"`
	RemoteData       []*RemoteData          `json:"remote_data,omitempty" url:"remote_data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (e *EmployeePayrollRun) GetId() *string {
	if e == nil {
		return nil
	}
	return e.Id
}

func (e *EmployeePayrollRun) GetRemoteId() *string {
	if e == nil {
		return nil
	}
	return e.RemoteId
}

func (e *EmployeePayrollRun) GetCreatedAt() *time.Time {
	if e == nil {
		return nil
	}
	return e.CreatedAt
}

func (e *EmployeePayrollRun) GetModifiedAt() *time.Time {
	if e == nil {
		return nil
	}
	return e.ModifiedAt
}

func (e *EmployeePayrollRun) GetEmployee() *EmployeePayrollRunEmployee {
	if e == nil {
		return nil
	}
	return e.Employee
}

func (e *EmployeePayrollRun) GetPayrollRun() *EmployeePayrollRunPayrollRun {
	if e == nil {
		return nil
	}
	return e.PayrollRun
}

func (e *EmployeePayrollRun) GetGrossPay() *float64 {
	if e == nil {
		return nil
	}
	return e.GrossPay
}

func (e *EmployeePayrollRun) GetNetPay() *float64 {
	if e == nil {
		return nil
	}
	return e.NetPay
}

func (e *EmployeePayrollRun) GetStartDate() *time.Time {
	if e == nil {
		return nil
	}
	return e.StartDate
}

func (e *EmployeePayrollRun) GetEndDate() *time.Time {
	if e == nil {
		return nil
	}
	return e.EndDate
}

func (e *EmployeePayrollRun) GetCheckDate() *time.Time {
	if e == nil {
		return nil
	}
	return e.CheckDate
}

func (e *EmployeePayrollRun) GetEarnings() []*Earning {
	if e == nil {
		return nil
	}
	return e.Earnings
}

func (e *EmployeePayrollRun) GetDeductions() []*Deduction {
	if e == nil {
		return nil
	}
	return e.Deductions
}

func (e *EmployeePayrollRun) GetTaxes() []*Tax {
	if e == nil {
		return nil
	}
	return e.Taxes
}

func (e *EmployeePayrollRun) GetRemoteWasDeleted() *bool {
	if e == nil {
		return nil
	}
	return e.RemoteWasDeleted
}

func (e *EmployeePayrollRun) GetFieldMappings() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.FieldMappings
}

func (e *EmployeePayrollRun) GetRemoteData() []*RemoteData {
	if e == nil {
		return nil
	}
	return e.RemoteData
}

func (e *EmployeePayrollRun) GetExtraProperties() map[string]interface{} {
	return e.extraProperties
}

func (e *EmployeePayrollRun) UnmarshalJSON(data []byte) error {
	type embed EmployeePayrollRun
	var unmarshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
		StartDate  *internal.DateTime `json:"start_date,omitempty"`
		EndDate    *internal.DateTime `json:"end_date,omitempty"`
		CheckDate  *internal.DateTime `json:"check_date,omitempty"`
	}{
		embed: embed(*e),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*e = EmployeePayrollRun(unmarshaler.embed)
	e.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	e.ModifiedAt = unmarshaler.ModifiedAt.TimePtr()
	e.StartDate = unmarshaler.StartDate.TimePtr()
	e.EndDate = unmarshaler.EndDate.TimePtr()
	e.CheckDate = unmarshaler.CheckDate.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	e.rawJSON = json.RawMessage(data)
	return nil
}

func (e *EmployeePayrollRun) MarshalJSON() ([]byte, error) {
	type embed EmployeePayrollRun
	var marshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
		StartDate  *internal.DateTime `json:"start_date,omitempty"`
		EndDate    *internal.DateTime `json:"end_date,omitempty"`
		CheckDate  *internal.DateTime `json:"check_date,omitempty"`
	}{
		embed:      embed(*e),
		CreatedAt:  internal.NewOptionalDateTime(e.CreatedAt),
		ModifiedAt: internal.NewOptionalDateTime(e.ModifiedAt),
		StartDate:  internal.NewOptionalDateTime(e.StartDate),
		EndDate:    internal.NewOptionalDateTime(e.EndDate),
		CheckDate:  internal.NewOptionalDateTime(e.CheckDate),
	}
	return json.Marshal(marshaler)
}

func (e *EmployeePayrollRun) String() string {
	if len(e.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(e.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(e); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", e)
}

// The employee whose payroll is being run.
type EmployeePayrollRunEmployee struct {
	String   string
	Employee *Employee

	typ string
}

func (e *EmployeePayrollRunEmployee) GetString() string {
	if e == nil {
		return ""
	}
	return e.String
}

func (e *EmployeePayrollRunEmployee) GetEmployee() *Employee {
	if e == nil {
		return nil
	}
	return e.Employee
}

func (e *EmployeePayrollRunEmployee) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		e.typ = "String"
		e.String = valueString
		return nil
	}
	valueEmployee := new(Employee)
	if err := json.Unmarshal(data, &valueEmployee); err == nil {
		e.typ = "Employee"
		e.Employee = valueEmployee
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, e)
}

func (e EmployeePayrollRunEmployee) MarshalJSON() ([]byte, error) {
	if e.typ == "String" || e.String != "" {
		return json.Marshal(e.String)
	}
	if e.typ == "Employee" || e.Employee != nil {
		return json.Marshal(e.Employee)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", e)
}

type EmployeePayrollRunEmployeeVisitor interface {
	VisitString(string) error
	VisitEmployee(*Employee) error
}

func (e *EmployeePayrollRunEmployee) Accept(visitor EmployeePayrollRunEmployeeVisitor) error {
	if e.typ == "String" || e.String != "" {
		return visitor.VisitString(e.String)
	}
	if e.typ == "Employee" || e.Employee != nil {
		return visitor.VisitEmployee(e.Employee)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", e)
}

// The payroll being run.
type EmployeePayrollRunPayrollRun struct {
	String     string
	PayrollRun *PayrollRun

	typ string
}

func (e *EmployeePayrollRunPayrollRun) GetString() string {
	if e == nil {
		return ""
	}
	return e.String
}

func (e *EmployeePayrollRunPayrollRun) GetPayrollRun() *PayrollRun {
	if e == nil {
		return nil
	}
	return e.PayrollRun
}

func (e *EmployeePayrollRunPayrollRun) UnmarshalJSON(data []byte) error {
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		e.typ = "String"
		e.String = valueString
		return nil
	}
	valuePayrollRun := new(PayrollRun)
	if err := json.Unmarshal(data, &valuePayrollRun); err == nil {
		e.typ = "PayrollRun"
		e.PayrollRun = valuePayrollRun
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, e)
}

func (e EmployeePayrollRunPayrollRun) MarshalJSON() ([]byte, error) {
	if e.typ == "String" || e.String != "" {
		return json.Marshal(e.String)
	}
	if e.typ == "PayrollRun" || e.PayrollRun != nil {
		return json.Marshal(e.PayrollRun)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", e)
}

type EmployeePayrollRunPayrollRunVisitor interface {
	VisitString(string) error
	VisitPayrollRun(*PayrollRun) error
}

func (e *EmployeePayrollRunPayrollRun) Accept(visitor EmployeePayrollRunPayrollRunVisitor) error {
	if e.typ == "String" || e.String != "" {
		return visitor.VisitString(e.String)
	}
	if e.typ == "PayrollRun" || e.PayrollRun != nil {
		return visitor.VisitPayrollRun(e.PayrollRun)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", e)
}

type PaginatedEmployeePayrollRunList struct {
	Next     *string               `json:"next,omitempty" url:"next,omitempty"`
	Previous *string               `json:"previous,omitempty" url:"previous,omitempty"`
	Results  []*EmployeePayrollRun `json:"results,omitempty" url:"results,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (p *PaginatedEmployeePayrollRunList) GetNext() *string {
	if p == nil {
		return nil
	}
	return p.Next
}

func (p *PaginatedEmployeePayrollRunList) GetPrevious() *string {
	if p == nil {
		return nil
	}
	return p.Previous
}

func (p *PaginatedEmployeePayrollRunList) GetResults() []*EmployeePayrollRun {
	if p == nil {
		return nil
	}
	return p.Results
}

func (p *PaginatedEmployeePayrollRunList) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PaginatedEmployeePayrollRunList) UnmarshalJSON(data []byte) error {
	type unmarshaler PaginatedEmployeePayrollRunList
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PaginatedEmployeePayrollRunList(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	p.rawJSON = json.RawMessage(data)
	return nil
}

func (p *PaginatedEmployeePayrollRunList) String() string {
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

// # The Tax Object
// ### Description
// The `Tax` object is used to represent an array of the tax deductions for a given employee's payroll run.
//
// ### Usage Example
// Fetch from the `LIST Taxes` endpoint and filter by `ID` to show all taxes.
type Tax struct {
	Id *string `json:"id,omitempty" url:"id,omitempty"`
	// The third-party API ID of the matching object.
	RemoteId *string `json:"remote_id,omitempty" url:"remote_id,omitempty"`
	// The datetime that this object was created by Merge.
	CreatedAt *time.Time `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The datetime that this object was modified by Merge.
	ModifiedAt         *time.Time `json:"modified_at,omitempty" url:"modified_at,omitempty"`
	EmployeePayrollRun *string    `json:"employee_payroll_run,omitempty" url:"employee_payroll_run,omitempty"`
	// The tax's name.
	Name *string `json:"name,omitempty" url:"name,omitempty"`
	// The tax amount.
	Amount *float64 `json:"amount,omitempty" url:"amount,omitempty"`
	// Whether or not the employer is responsible for paying the tax.
	EmployerTax *bool `json:"employer_tax,omitempty" url:"employer_tax,omitempty"`
	// Indicates whether or not this object has been deleted in the third party platform. Full coverage deletion detection is a premium add-on. Native deletion detection is offered for free with limited coverage. [Learn more](https://docs.merge.dev/integrations/hris/supported-features/).
	RemoteWasDeleted *bool                  `json:"remote_was_deleted,omitempty" url:"remote_was_deleted,omitempty"`
	FieldMappings    map[string]interface{} `json:"field_mappings,omitempty" url:"field_mappings,omitempty"`
	RemoteData       []*RemoteData          `json:"remote_data,omitempty" url:"remote_data,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (t *Tax) GetId() *string {
	if t == nil {
		return nil
	}
	return t.Id
}

func (t *Tax) GetRemoteId() *string {
	if t == nil {
		return nil
	}
	return t.RemoteId
}

func (t *Tax) GetCreatedAt() *time.Time {
	if t == nil {
		return nil
	}
	return t.CreatedAt
}

func (t *Tax) GetModifiedAt() *time.Time {
	if t == nil {
		return nil
	}
	return t.ModifiedAt
}

func (t *Tax) GetEmployeePayrollRun() *string {
	if t == nil {
		return nil
	}
	return t.EmployeePayrollRun
}

func (t *Tax) GetName() *string {
	if t == nil {
		return nil
	}
	return t.Name
}

func (t *Tax) GetAmount() *float64 {
	if t == nil {
		return nil
	}
	return t.Amount
}

func (t *Tax) GetEmployerTax() *bool {
	if t == nil {
		return nil
	}
	return t.EmployerTax
}

func (t *Tax) GetRemoteWasDeleted() *bool {
	if t == nil {
		return nil
	}
	return t.RemoteWasDeleted
}

func (t *Tax) GetFieldMappings() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.FieldMappings
}

func (t *Tax) GetRemoteData() []*RemoteData {
	if t == nil {
		return nil
	}
	return t.RemoteData
}

func (t *Tax) GetExtraProperties() map[string]interface{} {
	return t.extraProperties
}

func (t *Tax) UnmarshalJSON(data []byte) error {
	type embed Tax
	var unmarshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
	}{
		embed: embed(*t),
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	*t = Tax(unmarshaler.embed)
	t.CreatedAt = unmarshaler.CreatedAt.TimePtr()
	t.ModifiedAt = unmarshaler.ModifiedAt.TimePtr()
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	t.rawJSON = json.RawMessage(data)
	return nil
}

func (t *Tax) MarshalJSON() ([]byte, error) {
	type embed Tax
	var marshaler = struct {
		embed
		CreatedAt  *internal.DateTime `json:"created_at,omitempty"`
		ModifiedAt *internal.DateTime `json:"modified_at,omitempty"`
	}{
		embed:      embed(*t),
		CreatedAt:  internal.NewOptionalDateTime(t.CreatedAt),
		ModifiedAt: internal.NewOptionalDateTime(t.ModifiedAt),
	}
	return json.Marshal(marshaler)
}

func (t *Tax) String() string {
	if len(t.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(t.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}
