// This file was auto-generated by Fern from our API Definition.

package client

import (
	core "github.com/merge-api/merge-go-client/core"
	accountdetails "github.com/merge-api/merge-go-client/filestorage/accountdetails"
	accounttoken "github.com/merge-api/merge-go-client/filestorage/accounttoken"
	asyncpassthrough "github.com/merge-api/merge-go-client/filestorage/asyncpassthrough"
	audittrail "github.com/merge-api/merge-go-client/filestorage/audittrail"
	availableactions "github.com/merge-api/merge-go-client/filestorage/availableactions"
	deleteaccount "github.com/merge-api/merge-go-client/filestorage/deleteaccount"
	drives "github.com/merge-api/merge-go-client/filestorage/drives"
	files "github.com/merge-api/merge-go-client/filestorage/files"
	folders "github.com/merge-api/merge-go-client/filestorage/folders"
	forceresync "github.com/merge-api/merge-go-client/filestorage/forceresync"
	generatekey "github.com/merge-api/merge-go-client/filestorage/generatekey"
	groups "github.com/merge-api/merge-go-client/filestorage/groups"
	issues "github.com/merge-api/merge-go-client/filestorage/issues"
	linkedaccounts "github.com/merge-api/merge-go-client/filestorage/linkedaccounts"
	linktoken "github.com/merge-api/merge-go-client/filestorage/linktoken"
	passthrough "github.com/merge-api/merge-go-client/filestorage/passthrough"
	regeneratekey "github.com/merge-api/merge-go-client/filestorage/regeneratekey"
	selectivesync "github.com/merge-api/merge-go-client/filestorage/selectivesync"
	syncstatus "github.com/merge-api/merge-go-client/filestorage/syncstatus"
	users "github.com/merge-api/merge-go-client/filestorage/users"
	webhookreceivers "github.com/merge-api/merge-go-client/filestorage/webhookreceivers"
	http "net/http"
)

type Client interface {
	AccountDetails() accountdetails.Client
	AccountToken() accounttoken.Client
	AsyncPassthrough() asyncpassthrough.Client
	AuditTrail() audittrail.Client
	AvailableActions() availableactions.Client
	DeleteAccount() deleteaccount.Client
	Drives() drives.Client
	Files() files.Client
	Folders() folders.Client
	GenerateKey() generatekey.Client
	Groups() groups.Client
	Issues() issues.Client
	LinkToken() linktoken.Client
	LinkedAccounts() linkedaccounts.Client
	Passthrough() passthrough.Client
	RegenerateKey() regeneratekey.Client
	SelectiveSync() selectivesync.Client
	SyncStatus() syncstatus.Client
	ForceResync() forceresync.Client
	Users() users.Client
	WebhookReceivers() webhookreceivers.Client
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:                options.BaseURL,
		httpClient:             options.HTTPClient,
		header:                 options.ToHeader(),
		accountDetailsClient:   accountdetails.NewClient(opts...),
		accountTokenClient:     accounttoken.NewClient(opts...),
		asyncPassthroughClient: asyncpassthrough.NewClient(opts...),
		auditTrailClient:       audittrail.NewClient(opts...),
		availableActionsClient: availableactions.NewClient(opts...),
		deleteAccountClient:    deleteaccount.NewClient(opts...),
		drivesClient:           drives.NewClient(opts...),
		filesClient:            files.NewClient(opts...),
		foldersClient:          folders.NewClient(opts...),
		generateKeyClient:      generatekey.NewClient(opts...),
		groupsClient:           groups.NewClient(opts...),
		issuesClient:           issues.NewClient(opts...),
		linkTokenClient:        linktoken.NewClient(opts...),
		linkedAccountsClient:   linkedaccounts.NewClient(opts...),
		passthroughClient:      passthrough.NewClient(opts...),
		regenerateKeyClient:    regeneratekey.NewClient(opts...),
		selectiveSyncClient:    selectivesync.NewClient(opts...),
		syncStatusClient:       syncstatus.NewClient(opts...),
		forceResyncClient:      forceresync.NewClient(opts...),
		usersClient:            users.NewClient(opts...),
		webhookReceiversClient: webhookreceivers.NewClient(opts...),
	}
}

type client struct {
	baseURL                string
	httpClient             core.HTTPClient
	header                 http.Header
	accountDetailsClient   accountdetails.Client
	accountTokenClient     accounttoken.Client
	asyncPassthroughClient asyncpassthrough.Client
	auditTrailClient       audittrail.Client
	availableActionsClient availableactions.Client
	deleteAccountClient    deleteaccount.Client
	drivesClient           drives.Client
	filesClient            files.Client
	foldersClient          folders.Client
	generateKeyClient      generatekey.Client
	groupsClient           groups.Client
	issuesClient           issues.Client
	linkTokenClient        linktoken.Client
	linkedAccountsClient   linkedaccounts.Client
	passthroughClient      passthrough.Client
	regenerateKeyClient    regeneratekey.Client
	selectiveSyncClient    selectivesync.Client
	syncStatusClient       syncstatus.Client
	forceResyncClient      forceresync.Client
	usersClient            users.Client
	webhookReceiversClient webhookreceivers.Client
}

func (c *client) AccountDetails() accountdetails.Client {
	return c.accountDetailsClient
}

func (c *client) AccountToken() accounttoken.Client {
	return c.accountTokenClient
}

func (c *client) AsyncPassthrough() asyncpassthrough.Client {
	return c.asyncPassthroughClient
}

func (c *client) AuditTrail() audittrail.Client {
	return c.auditTrailClient
}

func (c *client) AvailableActions() availableactions.Client {
	return c.availableActionsClient
}

func (c *client) DeleteAccount() deleteaccount.Client {
	return c.deleteAccountClient
}

func (c *client) Drives() drives.Client {
	return c.drivesClient
}

func (c *client) Files() files.Client {
	return c.filesClient
}

func (c *client) Folders() folders.Client {
	return c.foldersClient
}

func (c *client) GenerateKey() generatekey.Client {
	return c.generateKeyClient
}

func (c *client) Groups() groups.Client {
	return c.groupsClient
}

func (c *client) Issues() issues.Client {
	return c.issuesClient
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

func (c *client) Users() users.Client {
	return c.usersClient
}

func (c *client) WebhookReceivers() webhookreceivers.Client {
	return c.webhookReceiversClient
}
