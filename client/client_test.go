// Code generated by Fern. DO NOT EDIT.

package client

import (
	option "github.com/merge-api/merge-go-client/v2/option"
	merge "github.com/merge-api/merge-go-client/v2"
	assert "github.com/stretchr/testify/assert"
	http "net/http"
	testing "testing"
	time "time"
)

func TestNewClient(t *testing.T) {
	t.Run("default", func(t *testing.T) {
		c := NewClient()
		assert.Equal(t, merge.Environments.Production, c.baseURL)
	})

	t.Run("base url", func(t *testing.T) {
		c := NewClient(
			option.WithBaseURL("test.co"),
		)
		assert.Equal(t, "test.co", c.baseURL)
	})

	t.Run("http client", func(t *testing.T) {
		httpClient := &http.Client{
			Timeout: 5 * time.Second,
		}
		c := NewClient(
			option.WithHTTPClient(httpClient),
		)
		assert.Equal(t, merge.Environments.Production, c.baseURL)
	})

	t.Run("http header", func(t *testing.T) {
		header := make(http.Header)
		header.Set("X-API-Tenancy", "test")
		c := NewClient(
			option.WithHTTPHeader(header),
		)
		assert.Equal(t, merge.Environments.Production, c.baseURL)
		assert.Equal(t, "test", c.header.Get("X-API-Tenancy"))
	})
}
