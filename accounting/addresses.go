// This file was auto-generated by Fern from our API Definition.

package accounting

type AddressesRetrieveRequest struct {
	// Whether to include the original data Merge fetched from the third-party to produce these models.
	IncludeRemoteData *bool `json:"-"`
	// Deprecated. Use show_enum_origins.
	RemoteFields *string `json:"-"`
	// Which fields should be returned in non-normalized form.
	ShowEnumOrigins *string `json:"-"`
}
