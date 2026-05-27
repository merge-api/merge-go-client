package integration

import (
	"context"
	"os"
	"testing"

	"github.com/merge-api/merge-go-client/v2/client"
	"github.com/merge-api/merge-go-client/v2/option"
	"github.com/merge-api/merge-go-client/v2/ticketing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTicketingClient(t *testing.T) *client.Client {
	t.Helper()
	apiKey := os.Getenv("SDK_TESTING_KEY")
	accountToken := os.Getenv("SDK_TESTING_TICKETING_ACCOUNT_TOKEN")
	if apiKey == "" || accountToken == "" {
		t.Fatal("SDK_TESTING_KEY and SDK_TESTING_TICKETING_ACCOUNT_TOKEN environment variables must be set")
	}
	return client.NewClient(
		option.WithApiKey(apiKey),
		option.WithAccountToken(&accountToken),
	)
}

// --- Tickets ---

func TestTicketing_Tickets_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Tickets_ListExpandAssignees(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{
		Expand: []*ticketing.TicketsListRequestExpandItem{ticketing.TicketsListRequestExpandItemAssignees.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Tickets_ListExpandCreator(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{
		Expand: []*ticketing.TicketsListRequestExpandItem{ticketing.TicketsListRequestExpandItemCreator.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Tickets_ListExpandCollections(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{
		Expand: []*ticketing.TicketsListRequestExpandItem{ticketing.TicketsListRequestExpandItemCollections.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Tickets_ListExpandAttachments(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{
		Expand: []*ticketing.TicketsListRequestExpandItem{ticketing.TicketsListRequestExpandItemAttachments.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Tickets_ListExpandAccount(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{
		Expand: []*ticketing.TicketsListRequestExpandItem{ticketing.TicketsListRequestExpandItemAccount.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Tickets_ListExpandContact(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{
		Expand: []*ticketing.TicketsListRequestExpandItem{ticketing.TicketsListRequestExpandItemContact.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Tickets_ListExpandParentTicket(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{
		Expand: []*ticketing.TicketsListRequestExpandItem{ticketing.TicketsListRequestExpandItemParentTicket.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

// --- Collections ---

func TestTicketing_Collections_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Collections.List(context.Background(), &ticketing.CollectionsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Collections_ListExpandParentCollection(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Collections.List(context.Background(), &ticketing.CollectionsListRequest{
		Expand: []*ticketing.CollectionsListRequestExpandItem{ticketing.CollectionsListRequestExpandItemParentCollection.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Collections_ListExpandPermissions(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Collections.List(context.Background(), &ticketing.CollectionsListRequest{
		Expand: []*ticketing.CollectionsListRequestExpandItem{ticketing.CollectionsListRequestExpandItemPermissions.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

// --- Comments ---

func TestTicketing_Comments_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Comments.List(context.Background(), &ticketing.CommentsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Comments_Retrieve(t *testing.T) {
	c := newTicketingClient(t)
	listResponse, err := c.Ticketing.Comments.List(context.Background(), &ticketing.CommentsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no comments available to test retrieve")
	}
	commentID := listResponse.Results[0].GetId()
	require.NotNil(t, commentID, "comment ID should not be nil")
	comment, err := c.Ticketing.Comments.Retrieve(context.Background(), *commentID, &ticketing.CommentsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, comment)
	assert.Equal(t, commentID, comment.GetId())
}

func TestTicketing_Comments_ListExpandUser(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Comments.List(context.Background(), &ticketing.CommentsListRequest{
		Expand: []*ticketing.CommentsListRequestExpandItem{ticketing.CommentsListRequestExpandItemUser.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Comments_ListExpandTicket(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Comments.List(context.Background(), &ticketing.CommentsListRequest{
		Expand: []*ticketing.CommentsListRequestExpandItem{ticketing.CommentsListRequestExpandItemTicket.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Comments_ListExpandContact(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Comments.List(context.Background(), &ticketing.CommentsListRequest{
		Expand: []*ticketing.CommentsListRequestExpandItem{ticketing.CommentsListRequestExpandItemContact.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

// --- Projects ---

func TestTicketing_Projects_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Projects.List(context.Background(), &ticketing.ProjectsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Projects_Retrieve(t *testing.T) {
	c := newTicketingClient(t)
	listResponse, err := c.Ticketing.Projects.List(context.Background(), &ticketing.ProjectsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no projects available to test retrieve")
	}
	projectID := listResponse.Results[0].GetId()
	require.NotNil(t, projectID, "project ID should not be nil")
	project, err := c.Ticketing.Projects.Retrieve(context.Background(), *projectID, &ticketing.ProjectsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, project)
	assert.Equal(t, projectID, project.GetId())
}

// --- Contacts ---

func TestTicketing_Contacts_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Contacts.List(context.Background(), &ticketing.ContactsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Contacts_Retrieve(t *testing.T) {
	c := newTicketingClient(t)
	listResponse, err := c.Ticketing.Contacts.List(context.Background(), &ticketing.ContactsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no contacts available to test retrieve")
	}
	contactID := listResponse.Results[0].GetId()
	require.NotNil(t, contactID, "contact ID should not be nil")
	contact, err := c.Ticketing.Contacts.Retrieve(context.Background(), *contactID, &ticketing.ContactsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, contact)
	assert.Equal(t, contactID, contact.GetId())
}

// --- Users ---

func TestTicketing_Users_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Users.List(context.Background(), &ticketing.UsersListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Users_Retrieve(t *testing.T) {
	c := newTicketingClient(t)
	listResponse, err := c.Ticketing.Users.List(context.Background(), &ticketing.UsersListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no users available to test retrieve")
	}
	userID := listResponse.Results[0].GetId()
	require.NotNil(t, userID, "user ID should not be nil")
	user, err := c.Ticketing.Users.Retrieve(context.Background(), *userID, &ticketing.UsersRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userID, user.GetId())
}

func TestTicketing_Users_ListExpandRoles(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Users.List(context.Background(), &ticketing.UsersListRequest{
		Expand: []*ticketing.UsersListRequestExpandItem{ticketing.UsersListRequestExpandItemRoles.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Users_ListExpandTeams(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Users.List(context.Background(), &ticketing.UsersListRequest{
		Expand: []*ticketing.UsersListRequestExpandItem{ticketing.UsersListRequestExpandItemTeams.Ptr()},
	})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

// --- Teams ---

func TestTicketing_Teams_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Teams.List(context.Background(), &ticketing.TeamsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Teams_Retrieve(t *testing.T) {
	c := newTicketingClient(t)
	listResponse, err := c.Ticketing.Teams.List(context.Background(), &ticketing.TeamsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no teams available to test retrieve")
	}
	teamID := listResponse.Results[0].GetId()
	require.NotNil(t, teamID, "team ID should not be nil")
	team, err := c.Ticketing.Teams.Retrieve(context.Background(), *teamID, &ticketing.TeamsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, team)
	assert.Equal(t, teamID, team.GetId())
}

// --- Roles ---

func TestTicketing_Roles_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Roles.List(context.Background(), &ticketing.RolesListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Roles_Retrieve(t *testing.T) {
	c := newTicketingClient(t)
	listResponse, err := c.Ticketing.Roles.List(context.Background(), &ticketing.RolesListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no roles available to test retrieve")
	}
	roleID := listResponse.Results[0].GetId()
	require.NotNil(t, roleID, "role ID should not be nil")
	role, err := c.Ticketing.Roles.Retrieve(context.Background(), *roleID, &ticketing.RolesRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, role)
	assert.Equal(t, roleID, role.GetId())
}

// --- Tags ---

func TestTicketing_Tags_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Tags.List(context.Background(), &ticketing.TagsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Tags_Retrieve(t *testing.T) {
	c := newTicketingClient(t)
	listResponse, err := c.Ticketing.Tags.List(context.Background(), &ticketing.TagsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no tags available to test retrieve")
	}
	tagID := listResponse.Results[0].GetId()
	require.NotNil(t, tagID, "tag ID should not be nil")
	tag, err := c.Ticketing.Tags.Retrieve(context.Background(), *tagID, &ticketing.TagsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, tag)
	assert.Equal(t, tagID, tag.GetId())
}

// --- Attachments ---

func TestTicketing_Attachments_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Attachments.List(context.Background(), &ticketing.AttachmentsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Attachments_Retrieve(t *testing.T) {
	c := newTicketingClient(t)
	listResponse, err := c.Ticketing.Attachments.List(context.Background(), &ticketing.AttachmentsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no attachments available to test retrieve")
	}
	attachmentID := listResponse.Results[0].GetId()
	require.NotNil(t, attachmentID, "attachment ID should not be nil")
	attachment, err := c.Ticketing.Attachments.Retrieve(context.Background(), *attachmentID, &ticketing.AttachmentsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, attachment)
	assert.Equal(t, attachmentID, attachment.GetId())
}

// --- Accounts ---

func TestTicketing_Accounts_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Accounts.List(context.Background(), &ticketing.AccountsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Accounts_Retrieve(t *testing.T) {
	c := newTicketingClient(t)
	listResponse, err := c.Ticketing.Accounts.List(context.Background(), &ticketing.AccountsListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no accounts available to test retrieve")
	}
	accountID := listResponse.Results[0].GetId()
	require.NotNil(t, accountID, "account ID should not be nil")
	account, err := c.Ticketing.Accounts.Retrieve(context.Background(), *accountID, &ticketing.AccountsRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, accountID, account.GetId())
}

// --- System resources ---

func TestTicketing_SyncStatus_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.SyncStatus.List(context.Background(), &ticketing.SyncStatusListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_AccountDetails_Retrieve(t *testing.T) {
	c := newTicketingClient(t)
	accountDetails, err := c.Ticketing.AccountDetails.Retrieve(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, accountDetails)
	assert.NotNil(t, accountDetails.GetId())
}

func TestTicketing_AuditTrail_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.AuditTrail.List(context.Background(), &ticketing.AuditTrailListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_AvailableActions_Retrieve(t *testing.T) {
	c := newTicketingClient(t)
	availableActions, err := c.Ticketing.AvailableActions.Retrieve(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, availableActions)
}

func TestTicketing_LinkedAccounts_List(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.LinkedAccounts.List(context.Background(), &ticketing.LinkedAccountsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestTicketing_Scopes_DefaultScopesRetrieve(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Scopes.DefaultScopesRetrieve(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, response)
}

func TestTicketing_Scopes_LinkedAccountScopesRetrieve(t *testing.T) {
	c := newTicketingClient(t)
	response, err := c.Ticketing.Scopes.LinkedAccountScopesRetrieve(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, response)
}
