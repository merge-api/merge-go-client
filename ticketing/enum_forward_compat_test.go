package ticketing_test

import (
	json "encoding/json"
	testing "testing"

	ticketing "github.com/merge-api/merge-go-client/v2/ticketing"
	assert "github.com/stretchr/testify/assert"
	require "github.com/stretchr/testify/require"
)

func TestEnumForwardCompatibility(t *testing.T) {
	t.Run("known value deserializes correctly", func(t *testing.T) {
		data := []byte(`"LIST"`)
		var e ticketing.CollectionTypeEnum
		require.NoError(t, json.Unmarshal(data, &e))
		assert.Equal(t, ticketing.CollectionTypeEnumList, e)
	})

	t.Run("unknown value is preserved without error", func(t *testing.T) {
		data := []byte(`"FUTURE_UNKNOWN_VALUE"`)
		var e ticketing.CollectionTypeEnum
		require.NoError(t, json.Unmarshal(data, &e))
		assert.Equal(t, ticketing.CollectionTypeEnum("FUTURE_UNKNOWN_VALUE"), e)
	})

	t.Run("unknown value roundtrips through marshal", func(t *testing.T) {
		original := ticketing.CollectionTypeEnum("FUTURE_UNKNOWN_VALUE")
		data, err := json.Marshal(original)
		require.NoError(t, err)

		var roundtripped ticketing.CollectionTypeEnum
		require.NoError(t, json.Unmarshal(data, &roundtripped))
		assert.Equal(t, original, roundtripped)
	})

	t.Run("unknown value preserved in struct field", func(t *testing.T) {
		data := []byte(`{"collection_type":"FUTURE_UNKNOWN_VALUE","name":"test"}`)
		var c ticketing.Collection
		require.NoError(t, json.Unmarshal(data, &c))
		require.NotNil(t, c.CollectionType)
		assert.Equal(t, ticketing.CollectionTypeEnum("FUTURE_UNKNOWN_VALUE"), *c.CollectionType)
	})
}
