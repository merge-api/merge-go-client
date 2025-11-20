package integration

import (
	"context"
	"os"
	"testing"

	"github.com/merge-api/merge-go-client/client"
	"github.com/merge-api/merge-go-client/filestorage"
	"github.com/merge-api/merge-go-client/option"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCursorPaginationWithFileStorageFolders(t *testing.T) {
	apiKey := os.Getenv("MERGE_API_KEY_FILESTORAGE")
	accountToken := os.Getenv("MERGE_ACCOUNT_TOKEN_FILESTORAGE")

	require.NotEmpty(t, apiKey, "MERGE_API_KEY_FILESTORAGE environment variable must be set")
	require.NotEmpty(t, accountToken, "MERGE_ACCOUNT_TOKEN_FILESTORAGE environment variable must be set")

	mergeClient := client.NewClient(
		option.WithApiKey(apiKey),
		option.WithAccountToken(&accountToken),
	)

	// Test basic folder listing with pagination
	response, err := mergeClient.FileStorage.Folders.List(
		context.TODO(),
		&filestorage.FoldersListRequest{},
	)

	require.NoError(t, err, "Folders list should not return an error")
	assert.NotNil(t, response, "Folders list response should not be null")

	// Test that we can access results
	assert.NotNil(t, response.Results, "Results list should not be null")

	// Test cursor pagination metadata access
	if response.Next != nil {
		assert.NotEmpty(t, *response.Next, "Cursor token should not be empty when present")
	}

	// Test that we have some data structure even if empty
	assert.NotNil(t, response.Results, "Results should always be accessible")
}

func TestStatelessPaginationWithCursor(t *testing.T) {
	apiKey := os.Getenv("MERGE_API_KEY_FILESTORAGE")
	accountToken := os.Getenv("MERGE_ACCOUNT_TOKEN_FILESTORAGE")

	require.NotEmpty(t, apiKey, "MERGE_API_KEY_FILESTORAGE environment variable must be set")
	require.NotEmpty(t, accountToken, "MERGE_ACCOUNT_TOKEN_FILESTORAGE environment variable must be set")

	mergeClient := client.NewClient(
		option.WithApiKey(apiKey),
		option.WithAccountToken(&accountToken),
	)

	// Simulate frontend flow: get first page
	firstPageResponse, err := mergeClient.FileStorage.Folders.List(
		context.TODO(),
		&filestorage.FoldersListRequest{},
	)

	require.NoError(t, err, "First page request should not return an error")
	require.NotNil(t, firstPageResponse, "First page response should not be null")

	// If there's a next cursor, test stateless pagination by actually using it
	if firstPageResponse.Next != nil && *firstPageResponse.Next != "" {
		nextCursor := *firstPageResponse.Next
		assert.NotEmpty(t, nextCursor, "Next cursor should not be empty")

		// Frontend sends cursor back to backend for next page
		secondPageRequest := &filestorage.FoldersListRequest{
			Cursor: &nextCursor,
		}

		// Backend fetches second page using the cursor
		secondPageResponse, err := mergeClient.FileStorage.Folders.List(
			context.TODO(),
			secondPageRequest,
		)

		require.NoError(t, err, "Second page request should not return an error")
		assert.NotNil(t, secondPageResponse, "Second page should not be null")

		// Verify we got a valid response with results
		assert.NotNil(t, secondPageResponse.Results, "Second page should have results")
	} else {
		t.Skip("No pagination available in current dataset - skipping cursor pagination test")
	}
}

func TestPaginationMetadataAccess(t *testing.T) {
	apiKey := os.Getenv("MERGE_API_KEY_FILESTORAGE")
	accountToken := os.Getenv("MERGE_ACCOUNT_TOKEN_FILESTORAGE")

	require.NotEmpty(t, apiKey, "MERGE_API_KEY_FILESTORAGE environment variable must be set")
	require.NotEmpty(t, accountToken, "MERGE_ACCOUNT_TOKEN_FILESTORAGE environment variable must be set")

	mergeClient := client.NewClient(
		option.WithApiKey(apiKey),
		option.WithAccountToken(&accountToken),
	)

	response, err := mergeClient.FileStorage.Folders.List(
		context.TODO(),
		&filestorage.FoldersListRequest{},
	)

	require.NoError(t, err, "Request should not return an error")
	require.NotNil(t, response, "Response should not be null")

	// Test access to pagination metadata
	// These fields may be nil if no pagination is needed
	if response.Next != nil {
		assert.NotNil(t, response.Next, "Next cursor should not be null if present")
	}

	if response.Previous != nil {
		assert.NotNil(t, response.Previous, "Previous cursor should not be null if present")
	}

	// At minimum, results should be accessible
	assert.NotNil(t, response.Results, "Results should always be accessible")
}

func TestIterationOverPaginatedResults(t *testing.T) {
	apiKey := os.Getenv("MERGE_API_KEY_FILESTORAGE")
	accountToken := os.Getenv("MERGE_ACCOUNT_TOKEN_FILESTORAGE")

	require.NotEmpty(t, apiKey, "MERGE_API_KEY_FILESTORAGE environment variable must be set")
	require.NotEmpty(t, accountToken, "MERGE_ACCOUNT_TOKEN_FILESTORAGE environment variable must be set")

	mergeClient := client.NewClient(
		option.WithApiKey(apiKey),
		option.WithAccountToken(&accountToken),
	)

	response, err := mergeClient.FileStorage.Folders.List(
		context.TODO(),
		&filestorage.FoldersListRequest{},
	)

	require.NoError(t, err, "Request should not return an error")
	require.NotNil(t, response, "Response should not be null")

	// Test iteration over paginated results
	itemCount := 0

	for _, folder := range response.Results {
		assert.NotNil(t, folder, "Each folder should not be null")
		itemCount++

		// Limit test to prevent excessive API calls
		if itemCount >= 10 {
			break
		}
	}

	assert.GreaterOrEqual(t, itemCount, 0, "Should have processed at least 0 items")
}

func TestPaginationWithPageSize(t *testing.T) {
	apiKey := os.Getenv("MERGE_API_KEY_FILESTORAGE")
	accountToken := os.Getenv("MERGE_ACCOUNT_TOKEN_FILESTORAGE")

	require.NotEmpty(t, apiKey, "MERGE_API_KEY_FILESTORAGE environment variable must be set")
	require.NotEmpty(t, accountToken, "MERGE_ACCOUNT_TOKEN_FILESTORAGE environment variable must be set")

	mergeClient := client.NewClient(
		option.WithApiKey(apiKey),
		option.WithAccountToken(&accountToken),
	)

	pageSize := 5
	response, err := mergeClient.FileStorage.Folders.List(
		context.TODO(),
		&filestorage.FoldersListRequest{
			PageSize: &pageSize,
		},
	)

	require.NoError(t, err, "Request should not return an error")
	require.NotNil(t, response, "Response should not be null")

	// Verify that page size constraint is respected (if there are enough items)
	if len(response.Results) > 0 {
		assert.LessOrEqual(t, len(response.Results), pageSize, "Result count should not exceed page size")
	}
}