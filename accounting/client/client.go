// This file was auto-generated by Fern from our API Definition.

package client

import (
	accountdetails "github.com/merge-api/merge-go-client/accounting/accountdetails"
	accountingperiods "github.com/merge-api/merge-go-client/accounting/accountingperiods"
	accounts "github.com/merge-api/merge-go-client/accounting/accounts"
	accounttoken "github.com/merge-api/merge-go-client/accounting/accounttoken"
	addresses "github.com/merge-api/merge-go-client/accounting/addresses"
	asyncpassthrough "github.com/merge-api/merge-go-client/accounting/asyncpassthrough"
	attachments "github.com/merge-api/merge-go-client/accounting/attachments"
	audittrail "github.com/merge-api/merge-go-client/accounting/audittrail"
	availableactions "github.com/merge-api/merge-go-client/accounting/availableactions"
	balancesheets "github.com/merge-api/merge-go-client/accounting/balancesheets"
	cashflowstatements "github.com/merge-api/merge-go-client/accounting/cashflowstatements"
	companyinfo "github.com/merge-api/merge-go-client/accounting/companyinfo"
	contacts "github.com/merge-api/merge-go-client/accounting/contacts"
	creditnotes "github.com/merge-api/merge-go-client/accounting/creditnotes"
	deleteaccount "github.com/merge-api/merge-go-client/accounting/deleteaccount"
	expenses "github.com/merge-api/merge-go-client/accounting/expenses"
	forceresync "github.com/merge-api/merge-go-client/accounting/forceresync"
	generatekey "github.com/merge-api/merge-go-client/accounting/generatekey"
	incomestatements "github.com/merge-api/merge-go-client/accounting/incomestatements"
	invoices "github.com/merge-api/merge-go-client/accounting/invoices"
	issues "github.com/merge-api/merge-go-client/accounting/issues"
	items "github.com/merge-api/merge-go-client/accounting/items"
	journalentries "github.com/merge-api/merge-go-client/accounting/journalentries"
	linkedaccounts "github.com/merge-api/merge-go-client/accounting/linkedaccounts"
	linktoken "github.com/merge-api/merge-go-client/accounting/linktoken"
	passthrough "github.com/merge-api/merge-go-client/accounting/passthrough"
	payments "github.com/merge-api/merge-go-client/accounting/payments"
	phonenumbers "github.com/merge-api/merge-go-client/accounting/phonenumbers"
	purchaseorders "github.com/merge-api/merge-go-client/accounting/purchaseorders"
	regeneratekey "github.com/merge-api/merge-go-client/accounting/regeneratekey"
	selectivesync "github.com/merge-api/merge-go-client/accounting/selectivesync"
	syncstatus "github.com/merge-api/merge-go-client/accounting/syncstatus"
	taxrates "github.com/merge-api/merge-go-client/accounting/taxrates"
	trackingcategories "github.com/merge-api/merge-go-client/accounting/trackingcategories"
	transactions "github.com/merge-api/merge-go-client/accounting/transactions"
	vendorcredits "github.com/merge-api/merge-go-client/accounting/vendorcredits"
	webhookreceivers "github.com/merge-api/merge-go-client/accounting/webhookreceivers"
	core "github.com/merge-api/merge-go-client/core"
	http "net/http"
)

type Client interface {
	AccountDetails() accountdetails.Client
	AccountToken() accounttoken.Client
	AccountingPeriods() accountingperiods.Client
	Accounts() accounts.Client
	Addresses() addresses.Client
	AsyncPassthrough() asyncpassthrough.Client
	Attachments() attachments.Client
	AuditTrail() audittrail.Client
	AvailableActions() availableactions.Client
	BalanceSheets() balancesheets.Client
	CashFlowStatements() cashflowstatements.Client
	CompanyInfo() companyinfo.Client
	Contacts() contacts.Client
	CreditNotes() creditnotes.Client
	DeleteAccount() deleteaccount.Client
	Expenses() expenses.Client
	GenerateKey() generatekey.Client
	IncomeStatements() incomestatements.Client
	Invoices() invoices.Client
	Issues() issues.Client
	Items() items.Client
	JournalEntries() journalentries.Client
	LinkToken() linktoken.Client
	LinkedAccounts() linkedaccounts.Client
	Passthrough() passthrough.Client
	Payments() payments.Client
	PhoneNumbers() phonenumbers.Client
	PurchaseOrders() purchaseorders.Client
	RegenerateKey() regeneratekey.Client
	SelectiveSync() selectivesync.Client
	SyncStatus() syncstatus.Client
	ForceResync() forceresync.Client
	TaxRates() taxrates.Client
	TrackingCategories() trackingcategories.Client
	Transactions() transactions.Client
	VendorCredits() vendorcredits.Client
	WebhookReceivers() webhookreceivers.Client
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:                  options.BaseURL,
		httpClient:               options.HTTPClient,
		header:                   options.ToHeader(),
		accountDetailsClient:     accountdetails.NewClient(opts...),
		accountTokenClient:       accounttoken.NewClient(opts...),
		accountingPeriodsClient:  accountingperiods.NewClient(opts...),
		accountsClient:           accounts.NewClient(opts...),
		addressesClient:          addresses.NewClient(opts...),
		asyncPassthroughClient:   asyncpassthrough.NewClient(opts...),
		attachmentsClient:        attachments.NewClient(opts...),
		auditTrailClient:         audittrail.NewClient(opts...),
		availableActionsClient:   availableactions.NewClient(opts...),
		balanceSheetsClient:      balancesheets.NewClient(opts...),
		cashFlowStatementsClient: cashflowstatements.NewClient(opts...),
		companyInfoClient:        companyinfo.NewClient(opts...),
		contactsClient:           contacts.NewClient(opts...),
		creditNotesClient:        creditnotes.NewClient(opts...),
		deleteAccountClient:      deleteaccount.NewClient(opts...),
		expensesClient:           expenses.NewClient(opts...),
		generateKeyClient:        generatekey.NewClient(opts...),
		incomeStatementsClient:   incomestatements.NewClient(opts...),
		invoicesClient:           invoices.NewClient(opts...),
		issuesClient:             issues.NewClient(opts...),
		itemsClient:              items.NewClient(opts...),
		journalEntriesClient:     journalentries.NewClient(opts...),
		linkTokenClient:          linktoken.NewClient(opts...),
		linkedAccountsClient:     linkedaccounts.NewClient(opts...),
		passthroughClient:        passthrough.NewClient(opts...),
		paymentsClient:           payments.NewClient(opts...),
		phoneNumbersClient:       phonenumbers.NewClient(opts...),
		purchaseOrdersClient:     purchaseorders.NewClient(opts...),
		regenerateKeyClient:      regeneratekey.NewClient(opts...),
		selectiveSyncClient:      selectivesync.NewClient(opts...),
		syncStatusClient:         syncstatus.NewClient(opts...),
		forceResyncClient:        forceresync.NewClient(opts...),
		taxRatesClient:           taxrates.NewClient(opts...),
		trackingCategoriesClient: trackingcategories.NewClient(opts...),
		transactionsClient:       transactions.NewClient(opts...),
		vendorCreditsClient:      vendorcredits.NewClient(opts...),
		webhookReceiversClient:   webhookreceivers.NewClient(opts...),
	}
}

type client struct {
	baseURL                  string
	httpClient               core.HTTPClient
	header                   http.Header
	accountDetailsClient     accountdetails.Client
	accountTokenClient       accounttoken.Client
	accountingPeriodsClient  accountingperiods.Client
	accountsClient           accounts.Client
	addressesClient          addresses.Client
	asyncPassthroughClient   asyncpassthrough.Client
	attachmentsClient        attachments.Client
	auditTrailClient         audittrail.Client
	availableActionsClient   availableactions.Client
	balanceSheetsClient      balancesheets.Client
	cashFlowStatementsClient cashflowstatements.Client
	companyInfoClient        companyinfo.Client
	contactsClient           contacts.Client
	creditNotesClient        creditnotes.Client
	deleteAccountClient      deleteaccount.Client
	expensesClient           expenses.Client
	generateKeyClient        generatekey.Client
	incomeStatementsClient   incomestatements.Client
	invoicesClient           invoices.Client
	issuesClient             issues.Client
	itemsClient              items.Client
	journalEntriesClient     journalentries.Client
	linkTokenClient          linktoken.Client
	linkedAccountsClient     linkedaccounts.Client
	passthroughClient        passthrough.Client
	paymentsClient           payments.Client
	phoneNumbersClient       phonenumbers.Client
	purchaseOrdersClient     purchaseorders.Client
	regenerateKeyClient      regeneratekey.Client
	selectiveSyncClient      selectivesync.Client
	syncStatusClient         syncstatus.Client
	forceResyncClient        forceresync.Client
	taxRatesClient           taxrates.Client
	trackingCategoriesClient trackingcategories.Client
	transactionsClient       transactions.Client
	vendorCreditsClient      vendorcredits.Client
	webhookReceiversClient   webhookreceivers.Client
}

func (c *client) AccountDetails() accountdetails.Client {
	return c.accountDetailsClient
}

func (c *client) AccountToken() accounttoken.Client {
	return c.accountTokenClient
}

func (c *client) AccountingPeriods() accountingperiods.Client {
	return c.accountingPeriodsClient
}

func (c *client) Accounts() accounts.Client {
	return c.accountsClient
}

func (c *client) Addresses() addresses.Client {
	return c.addressesClient
}

func (c *client) AsyncPassthrough() asyncpassthrough.Client {
	return c.asyncPassthroughClient
}

func (c *client) Attachments() attachments.Client {
	return c.attachmentsClient
}

func (c *client) AuditTrail() audittrail.Client {
	return c.auditTrailClient
}

func (c *client) AvailableActions() availableactions.Client {
	return c.availableActionsClient
}

func (c *client) BalanceSheets() balancesheets.Client {
	return c.balanceSheetsClient
}

func (c *client) CashFlowStatements() cashflowstatements.Client {
	return c.cashFlowStatementsClient
}

func (c *client) CompanyInfo() companyinfo.Client {
	return c.companyInfoClient
}

func (c *client) Contacts() contacts.Client {
	return c.contactsClient
}

func (c *client) CreditNotes() creditnotes.Client {
	return c.creditNotesClient
}

func (c *client) DeleteAccount() deleteaccount.Client {
	return c.deleteAccountClient
}

func (c *client) Expenses() expenses.Client {
	return c.expensesClient
}

func (c *client) GenerateKey() generatekey.Client {
	return c.generateKeyClient
}

func (c *client) IncomeStatements() incomestatements.Client {
	return c.incomeStatementsClient
}

func (c *client) Invoices() invoices.Client {
	return c.invoicesClient
}

func (c *client) Issues() issues.Client {
	return c.issuesClient
}

func (c *client) Items() items.Client {
	return c.itemsClient
}

func (c *client) JournalEntries() journalentries.Client {
	return c.journalEntriesClient
}

func (c *client) LinkToken() linktoken.Client {
	return c.linkTokenClient
}

func (c *client) LinkedAccounts() linkedaccounts.Client {
	return c.linkedAccountsClient
}

func (c *client) Passthrough() passthrough.Client {
	return c.passthroughClient
}

func (c *client) Payments() payments.Client {
	return c.paymentsClient
}

func (c *client) PhoneNumbers() phonenumbers.Client {
	return c.phoneNumbersClient
}

func (c *client) PurchaseOrders() purchaseorders.Client {
	return c.purchaseOrdersClient
}

func (c *client) RegenerateKey() regeneratekey.Client {
	return c.regenerateKeyClient
}

func (c *client) SelectiveSync() selectivesync.Client {
	return c.selectiveSyncClient
}

func (c *client) SyncStatus() syncstatus.Client {
	return c.syncStatusClient
}

func (c *client) ForceResync() forceresync.Client {
	return c.forceResyncClient
}

func (c *client) TaxRates() taxrates.Client {
	return c.taxRatesClient
}

func (c *client) TrackingCategories() trackingcategories.Client {
	return c.trackingCategoriesClient
}

func (c *client) Transactions() transactions.Client {
	return c.transactionsClient
}

func (c *client) VendorCredits() vendorcredits.Client {
	return c.vendorCreditsClient
}

func (c *client) WebhookReceivers() webhookreceivers.Client {
	return c.webhookReceiversClient
}
