package integration

import (
	"context"
	"os"
	"testing"

	"github.com/merge-api/merge-go-client/v2/client"
	"github.com/merge-api/merge-go-client/v2/hris"
	"github.com/merge-api/merge-go-client/v2/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newHRISClient(t *testing.T) *client.Client {
	t.Helper()
	apiKey := os.Getenv("SDK_TESTING_KEY")
	accountToken := os.Getenv("SDK_TESTING_HRIS_ACCOUNT_TOKEN")
	if apiKey == "" || accountToken == "" {
		t.Fatal("SDK_TESTING_KEY and SDK_TESTING_HRIS_ACCOUNT_TOKEN environment variables must be set")
	}
	return client.NewClient(
		option.WithApiKey(apiKey),
		option.WithAccountToken(&accountToken),
	)
}

// --- Employees ---

func TestHRIS_Employees_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.Employees.List(context.Background(), &hris.EmployeesListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_Employees_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.Employees.List(context.Background(), &hris.EmployeesListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no employees available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	employee, err := c.Hris.Employees.Retrieve(context.Background(), *id, &hris.EmployeesRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, employee)
	assert.Equal(t, id, employee.GetId())
}

// --- Employments ---

func TestHRIS_Employments_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.Employments.List(context.Background(), &hris.EmploymentsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_Employments_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.Employments.List(context.Background(), &hris.EmploymentsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no employments available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	employment, err := c.Hris.Employments.Retrieve(context.Background(), *id, &hris.EmploymentsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, employment)
	assert.Equal(t, id, employment.GetId())
}

// --- Teams ---

func TestHRIS_Teams_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.Teams.List(context.Background(), &hris.TeamsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_Teams_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.Teams.List(context.Background(), &hris.TeamsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no teams available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	team, err := c.Hris.Teams.Retrieve(context.Background(), *id, &hris.TeamsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, team)
	assert.Equal(t, id, team.GetId())
}

// --- Groups ---

func TestHRIS_Groups_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.Groups.List(context.Background(), &hris.GroupsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_Groups_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.Groups.List(context.Background(), &hris.GroupsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no groups available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	group, err := c.Hris.Groups.Retrieve(context.Background(), *id, &hris.GroupsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, group)
	assert.Equal(t, id, group.GetId())
}

// --- Locations ---

func TestHRIS_Locations_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.Locations.List(context.Background(), &hris.LocationsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_Locations_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.Locations.List(context.Background(), &hris.LocationsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no locations available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	location, err := c.Hris.Locations.Retrieve(context.Background(), *id, &hris.LocationsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, location)
	assert.Equal(t, id, location.GetId())
}

// --- TimeOff ---

func TestHRIS_TimeOff_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.TimeOff.List(context.Background(), &hris.TimeOffListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_TimeOff_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.TimeOff.List(context.Background(), &hris.TimeOffListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no time-off records available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	timeOff, err := c.Hris.TimeOff.Retrieve(context.Background(), *id, &hris.TimeOffRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, timeOff)
	assert.Equal(t, id, timeOff.GetId())
}

// --- TimeOffBalances ---

func TestHRIS_TimeOffBalances_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.TimeOffBalances.List(context.Background(), &hris.TimeOffBalancesListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_TimeOffBalances_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.TimeOffBalances.List(context.Background(), &hris.TimeOffBalancesListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no time-off balances available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	balance, err := c.Hris.TimeOffBalances.Retrieve(context.Background(), *id, &hris.TimeOffBalancesRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, balance)
	assert.Equal(t, id, balance.GetId())
}

// --- Benefits ---

func TestHRIS_Benefits_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.Benefits.List(context.Background(), &hris.BenefitsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_Benefits_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.Benefits.List(context.Background(), &hris.BenefitsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no benefits available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	benefit, err := c.Hris.Benefits.Retrieve(context.Background(), *id, &hris.BenefitsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, benefit)
	assert.Equal(t, id, benefit.GetId())
}

// --- Companies ---

func TestHRIS_Companies_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.Companies.List(context.Background(), &hris.CompaniesListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_Companies_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.Companies.List(context.Background(), &hris.CompaniesListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no companies available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	company, err := c.Hris.Companies.Retrieve(context.Background(), *id, &hris.CompaniesRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, company)
	assert.Equal(t, id, company.GetId())
}

// --- Dependents ---

func TestHRIS_Dependents_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.Dependents.List(context.Background(), &hris.DependentsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_Dependents_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.Dependents.List(context.Background(), &hris.DependentsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no dependents available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	dependent, err := c.Hris.Dependents.Retrieve(context.Background(), *id, &hris.DependentsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, dependent)
	assert.Equal(t, id, dependent.GetId())
}

// --- BankInfo ---

func TestHRIS_BankInfo_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.BankInfo.List(context.Background(), &hris.BankInfoListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_BankInfo_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.BankInfo.List(context.Background(), &hris.BankInfoListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no bank info records available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	bankInfo, err := c.Hris.BankInfo.Retrieve(context.Background(), *id, &hris.BankInfoRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, bankInfo)
	assert.Equal(t, id, bankInfo.GetId())
}

// --- PayGroups ---

func TestHRIS_PayGroups_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.PayGroups.List(context.Background(), &hris.PayGroupsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_PayGroups_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.PayGroups.List(context.Background(), &hris.PayGroupsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no pay groups available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	payGroup, err := c.Hris.PayGroups.Retrieve(context.Background(), *id, &hris.PayGroupsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, payGroup)
	assert.Equal(t, id, payGroup.GetId())
}

// --- PayrollRuns ---

func TestHRIS_PayrollRuns_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.PayrollRuns.List(context.Background(), &hris.PayrollRunsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_PayrollRuns_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.PayrollRuns.List(context.Background(), &hris.PayrollRunsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no payroll runs available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	payrollRun, err := c.Hris.PayrollRuns.Retrieve(context.Background(), *id, &hris.PayrollRunsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, payrollRun)
	assert.Equal(t, id, payrollRun.GetId())
}

// --- EmployeePayrollRuns ---

func TestHRIS_EmployeePayrollRuns_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.EmployeePayrollRuns.List(context.Background(), &hris.EmployeePayrollRunsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_EmployeePayrollRuns_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.EmployeePayrollRuns.List(context.Background(), &hris.EmployeePayrollRunsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no employee payroll runs available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	if id == nil {
		t.Skip("employee payroll run has no id, skipping retrieve")
	}
	run, err := c.Hris.EmployeePayrollRuns.Retrieve(context.Background(), *id, &hris.EmployeePayrollRunsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, run)
	assert.Equal(t, id, run.GetId())
}

// --- TimesheetEntries ---

func TestHRIS_TimesheetEntries_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.TimesheetEntries.List(context.Background(), &hris.TimesheetEntriesListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_TimesheetEntries_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.TimesheetEntries.List(context.Background(), &hris.TimesheetEntriesListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no timesheet entries available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	entry, err := c.Hris.TimesheetEntries.Retrieve(context.Background(), *id, &hris.TimesheetEntriesRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, entry)
	assert.Equal(t, id, entry.GetId())
}

// --- EmployerBenefits ---

func TestHRIS_EmployerBenefits_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.EmployerBenefits.List(context.Background(), &hris.EmployerBenefitsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_EmployerBenefits_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	listResponse, err := c.Hris.EmployerBenefits.List(context.Background(), &hris.EmployerBenefitsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no employer benefits available to test retrieve")
	}
	id := listResponse.Results[0].GetId()
	require.NotNil(t, id)
	benefit, err := c.Hris.EmployerBenefits.Retrieve(context.Background(), *id, &hris.EmployerBenefitsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, benefit)
	assert.Equal(t, id, benefit.GetId())
}

// --- System resources ---

func TestHRIS_SyncStatus_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.SyncStatus.List(context.Background(), &hris.SyncStatusListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_AccountDetails_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	accountDetails, err := c.Hris.AccountDetails.Retrieve(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, accountDetails)
	assert.NotNil(t, accountDetails.GetId())
}

func TestHRIS_AuditTrail_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.AuditTrail.List(context.Background(), &hris.AuditTrailListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_AvailableActions_Retrieve(t *testing.T) {
	c := newHRISClient(t)
	availableActions, err := c.Hris.AvailableActions.Retrieve(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, availableActions)
}

func TestHRIS_LinkedAccounts_List(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.LinkedAccounts.List(context.Background(), &hris.LinkedAccountsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestHRIS_Scopes_DefaultScopesRetrieve(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.Scopes.DefaultScopesRetrieve(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, response)
}

func TestHRIS_Scopes_LinkedAccountScopesRetrieve(t *testing.T) {
	c := newHRISClient(t)
	response, err := c.Hris.Scopes.LinkedAccountScopesRetrieve(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, response)
}
