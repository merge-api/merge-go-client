package integration

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/merge-api/merge-go-client/v2/ticketing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDatetimeFilters_UTCSerializesAsZ(t *testing.T) {
	c, reqs := captureClient(t, func(_ *http.Request) *http.Response { return jsonListResponse() })
	when := time.Date(2026, 1, 15, 9, 30, 0, 0, time.UTC)
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{CreatedAfter: &when})
	require.NoError(t, err)
	got := reqs()
	assert.Equal(t, "2026-01-15T09:30:00Z", got[0].URL.Query().Get("created_after"))
}

func TestDatetimeFilters_NegativeOffsetPreserved(t *testing.T) {
	c, reqs := captureClient(t, func(_ *http.Request) *http.Response { return jsonListResponse() })
	loc := time.FixedZone("UTC-5", -5*60*60)
	when := time.Date(2026, 1, 15, 9, 30, 0, 0, loc)
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{CreatedAfter: &when})
	require.NoError(t, err)
	got := reqs()
	assert.Equal(t, "2026-01-15T09:30:00-05:00", got[0].URL.Query().Get("created_after"))
}

func TestDatetimeFilters_FilterPairBothPresent(t *testing.T) {
	c, reqs := captureClient(t, func(_ *http.Request) *http.Response { return jsonListResponse() })
	after := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	before := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{
		RemoteUpdatedAfter:  &after,
		RemoteUpdatedBefore: &before,
	})
	require.NoError(t, err)
	got := reqs()
	params := got[0].URL.Query()
	assert.Equal(t, "2026-01-01T00:00:00Z", params.Get("remote_updated_after"))
	assert.Equal(t, "2026-02-01T00:00:00Z", params.Get("remote_updated_before"))
}

func TestDatetimeFilters_NilDateOmittedFromURL(t *testing.T) {
	c, reqs := captureClient(t, func(_ *http.Request) *http.Response { return jsonListResponse() })
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{CreatedAfter: nil})
	require.NoError(t, err)
	got := reqs()
	assert.Empty(t, got[0].URL.Query().Get("created_after"), "nil datetime should be omitted from URL")
}
