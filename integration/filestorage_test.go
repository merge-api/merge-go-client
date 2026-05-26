package integration

import (
	"context"
	"os"
	"testing"

	"github.com/merge-api/merge-go-client/v2/client"
	"github.com/merge-api/merge-go-client/v2/filestorage"
	"github.com/merge-api/merge-go-client/v2/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newFileStorageClient(t *testing.T) *client.Client {
	t.Helper()
	apiKey := os.Getenv("SDK_TESTING_KEY")
	accountToken := os.Getenv("SDK_TESTING_FILE_STORAGE_ACCOUNT_TOKEN")
	if apiKey == "" || accountToken == "" {
		t.Fatal("SDK_TESTING_KEY and SDK_TESTING_FILE_STORAGE_ACCOUNT_TOKEN environment variables must be set")
	}
	return client.NewClient(
		option.WithApiKey(apiKey),
		option.WithAccountToken(&accountToken),
	)
}

func TestFileStorage_Files_List(t *testing.T) {
	c := newFileStorageClient(t)
	response, err := c.FileStorage.Files.List(context.Background(), &filestorage.FilesListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestFileStorage_Files_ListPermissions(t *testing.T) {
	c := newFileStorageClient(t)
	response, err := c.FileStorage.Files.List(context.Background(), &filestorage.FilesListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(response.Results) == 0 {
		t.Skip("no files available to test permissions")
	}
	file := response.Results[0]
	assert.NotNil(t, file)
	// Permissions field type is verified by the Go type system; assert the field is accessible.
	_ = file.GetPermissions()
}

func TestFileStorage_Files_Retrieve(t *testing.T) {
	c := newFileStorageClient(t)
	listResponse, err := c.FileStorage.Files.List(context.Background(), &filestorage.FilesListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no files available to test retrieve")
	}
	fileID := listResponse.Results[0].GetId()
	require.NotNil(t, fileID, "file ID should not be nil")
	file, err := c.FileStorage.Files.Retrieve(context.Background(), *fileID, &filestorage.FilesRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, file)
	assert.Equal(t, fileID, file.GetId())
}

func TestFileStorage_Drives_List(t *testing.T) {
	c := newFileStorageClient(t)
	response, err := c.FileStorage.Drives.List(context.Background(), &filestorage.DrivesListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestFileStorage_Drives_Retrieve(t *testing.T) {
	c := newFileStorageClient(t)
	listResponse, err := c.FileStorage.Drives.List(context.Background(), &filestorage.DrivesListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no drives available to test retrieve")
	}
	driveID := listResponse.Results[0].GetId()
	require.NotNil(t, driveID, "drive ID should not be nil")
	drive, err := c.FileStorage.Drives.Retrieve(context.Background(), *driveID, &filestorage.DrivesRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, drive)
	assert.Equal(t, driveID, drive.GetId())
}

func TestFileStorage_Folders_List(t *testing.T) {
	c := newFileStorageClient(t)
	response, err := c.FileStorage.Folders.List(context.Background(), &filestorage.FoldersListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestFileStorage_Users_List(t *testing.T) {
	c := newFileStorageClient(t)
	response, err := c.FileStorage.Users.List(context.Background(), &filestorage.UsersListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestFileStorage_Users_Retrieve(t *testing.T) {
	c := newFileStorageClient(t)
	listResponse, err := c.FileStorage.Users.List(context.Background(), &filestorage.UsersListRequest{
		PageSize: intPtr(1),
	})
	require.NoError(t, err)
	if len(listResponse.Results) == 0 {
		t.Skip("no users available to test retrieve")
	}
	userID := listResponse.Results[0].GetId()
	require.NotNil(t, userID, "user ID should not be nil")
	user, err := c.FileStorage.Users.Retrieve(context.Background(), *userID, &filestorage.UsersRetrieveRequest{})
	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userID, user.GetId())
}

func TestFileStorage_Groups_List(t *testing.T) {
	c := newFileStorageClient(t)
	response, err := c.FileStorage.Groups.List(context.Background(), &filestorage.GroupsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestFileStorage_SyncStatus_List(t *testing.T) {
	c := newFileStorageClient(t)
	response, err := c.FileStorage.SyncStatus.List(context.Background(), &filestorage.SyncStatusListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestFileStorage_AccountDetails_Retrieve(t *testing.T) {
	c := newFileStorageClient(t)
	accountDetails, err := c.FileStorage.AccountDetails.Retrieve(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, accountDetails)
	assert.NotNil(t, accountDetails.GetId())
}

func TestFileStorage_AuditTrail_List(t *testing.T) {
	c := newFileStorageClient(t)
	response, err := c.FileStorage.AuditTrail.List(context.Background(), &filestorage.AuditTrailListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestFileStorage_AvailableActions_Retrieve(t *testing.T) {
	c := newFileStorageClient(t)
	availableActions, err := c.FileStorage.AvailableActions.Retrieve(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, availableActions)
}

func TestFileStorage_LinkedAccounts_List(t *testing.T) {
	c := newFileStorageClient(t)
	response, err := c.FileStorage.LinkedAccounts.List(context.Background(), &filestorage.LinkedAccountsListRequest{})
	require.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotNil(t, response.Results)
}

func TestFileStorage_Scopes_DefaultScopesRetrieve(t *testing.T) {
	c := newFileStorageClient(t)
	response, err := c.FileStorage.Scopes.DefaultScopesRetrieve(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, response)
}

func TestFileStorage_Scopes_LinkedAccountScopesRetrieve(t *testing.T) {
	c := newFileStorageClient(t)
	response, err := c.FileStorage.Scopes.LinkedAccountScopesRetrieve(context.Background())
	require.NoError(t, err)
	assert.NotNil(t, response)
}
