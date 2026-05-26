package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/merge-api/merge-go-client/v2/ticketing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExpandURLFormat_Omitted(t *testing.T) {
	c, reqs := captureClient(t, func(_ *http.Request) *http.Response { return jsonListResponse() })
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{PageSize: intPtr(1)})
	require.NoError(t, err)
	got := reqs()
	assert.Empty(t, got[0].URL.Query()["expand"], "expand should be absent when not set")
}

func TestExpandURLFormat_SingleItem(t *testing.T) {
	c, reqs := captureClient(t, func(_ *http.Request) *http.Response { return jsonListResponse() })
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{
		Expand: []*ticketing.TicketsListRequestExpandItem{ticketing.TicketsListRequestExpandItemAccount.Ptr()},
	})
	require.NoError(t, err)
	got := reqs()
	assert.Equal(t, "account", got[0].URL.Query().Get("expand"))
}

func TestExpandURLFormat_MultipleItemsJoinedWithComma(t *testing.T) {
	c, reqs := captureClient(t, func(_ *http.Request) *http.Response { return jsonListResponse() })
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{
		Expand: []*ticketing.TicketsListRequestExpandItem{
			ticketing.TicketsListRequestExpandItemAssignees.Ptr(),
			ticketing.TicketsListRequestExpandItemAccount.Ptr(),
			ticketing.TicketsListRequestExpandItemCreator.Ptr(),
		},
	})
	require.NoError(t, err)
	got := reqs()
	// Multiple expand values must be comma-joined into a single query parameter.
	assert.Equal(t, "assignees,account,creator", got[0].URL.Query().Get("expand"))
	assert.Len(t, got[0].URL.Query()["expand"], 1, "expand should appear exactly once in the URL")
}

func TestExpandURLFormat_CoexistsWithOtherParams(t *testing.T) {
	c, reqs := captureClient(t, func(_ *http.Request) *http.Response { return jsonListResponse() })
	cursor := "cursor-abc"
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{
		Expand:   []*ticketing.TicketsListRequestExpandItem{ticketing.TicketsListRequestExpandItemAccount.Ptr()},
		PageSize: intPtr(50),
		Cursor:   &cursor,
	})
	require.NoError(t, err)
	got := reqs()
	params := got[0].URL.Query()
	assert.Equal(t, "account", params.Get("expand"))
	assert.Equal(t, "50", params.Get("page_size"))
	assert.Equal(t, "cursor-abc", params.Get("cursor"))
}

func TestExpandURLFormat_EmptySliceOmitted(t *testing.T) {
	c, reqs := captureClient(t, func(_ *http.Request) *http.Response { return jsonListResponse() })
	_, err := c.Ticketing.Tickets.List(context.Background(), &ticketing.TicketsListRequest{
		Expand: []*ticketing.TicketsListRequestExpandItem{},
	})
	require.NoError(t, err)
	got := reqs()
	assert.Empty(t, got[0].URL.Query()["expand"], "empty expand slice should produce no expand param")
}

func TestExpandURLFormat_CrossCategoryExpand(t *testing.T) {
	c, reqs := captureClient(t, func(_ *http.Request) *http.Response { return jsonListResponse() })
	_, err := c.Ticketing.Comments.List(context.Background(), &ticketing.CommentsListRequest{
		Expand: []*ticketing.CommentsListRequestExpandItem{
			ticketing.CommentsListRequestExpandItemUser.Ptr(),
			ticketing.CommentsListRequestExpandItemTicket.Ptr(),
		},
	})
	require.NoError(t, err)
	got := reqs()
	assert.Equal(t, "user,ticket", got[0].URL.Query().Get("expand"))
}
