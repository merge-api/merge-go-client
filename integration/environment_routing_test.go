package integration

import (
	"context"
	"net/http"
	"strings"
	"testing"

	merge "github.com/merge-api/merge-go-client/v2"
	"github.com/merge-api/merge-go-client/v2/option"
	"github.com/merge-api/merge-go-client/v2/ticketing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func listOnce(t *testing.T, opts ...option.RequestOption) *http.Request {
	t.Helper()
	c, reqs := captureClient(t, func(_ *http.Request) *http.Response {
		return jsonListResponse()
	}, opts...)
	pageSize := 1
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{PageSize: &pageSize})
	require.NoError(t, err)
	got := reqs()
	require.Len(t, got, 1)
	return got[0]
}

func TestEnvironmentRouting_DefaultIsProduction(t *testing.T) {
	req := listOnce(t)
	assert.True(t, strings.HasPrefix(req.URL.String(), merge.Environments.Production),
		"expected URL to start with %s, got %s", merge.Environments.Production, req.URL.String())
}

func TestEnvironmentRouting_Sandbox(t *testing.T) {
	req := listOnce(t, option.WithBaseURL(merge.Environments.Sandbox))
	assert.True(t, strings.HasPrefix(req.URL.String(), merge.Environments.Sandbox),
		"expected URL to start with %s, got %s", merge.Environments.Sandbox, req.URL.String())
}

func TestEnvironmentRouting_ProductionEU(t *testing.T) {
	req := listOnce(t, option.WithBaseURL(merge.Environments.ProductionEu))
	assert.True(t, strings.HasPrefix(req.URL.String(), merge.Environments.ProductionEu),
		"expected URL to start with %s, got %s", merge.Environments.ProductionEu, req.URL.String())
}

func TestEnvironmentRouting_CustomBaseURLOverrides(t *testing.T) {
	customBase := "https://proxy.example.com/merge"
	req := listOnce(t, option.WithBaseURL(customBase))
	assert.True(t, strings.HasPrefix(req.URL.String(), customBase),
		"expected URL to start with %s, got %s", customBase, req.URL.String())
}

func TestEnvironmentRouting_RequestMethodAndPath(t *testing.T) {
	req := listOnce(t)
	assert.Equal(t, "GET", req.Method)
	assert.Contains(t, req.URL.Path, "/ticketing/v1/tickets")
}
