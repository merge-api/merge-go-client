// This file was auto-generated by Fern from our API Definition.

package ticketing

type EndUserDetailsRequest struct {
	// Your end user's email address. This is purely for identification purposes - setting this value will not cause any emails to be sent.
	EndUserEmailAddress string `json:"end_user_email_address"`
	// Your end user's organization.
	EndUserOrganizationName string `json:"end_user_organization_name"`
	// This unique identifier typically represents the ID for your end user in your product's database. This value must be distinct from other Linked Accounts' unique identifiers.
	EndUserOriginId string `json:"end_user_origin_id"`
	// The integration categories to show in Merge Link.
	Categories []CategoriesEnum `json:"categories,omitempty"`
	// The slug of a specific pre-selected integration for this linking flow token. For examples of slugs, see https://docs.merge.dev/guides/merge-link/single-integration/.
	Integration *string `json:"integration,omitempty"`
	// An integer number of minutes between [30, 720 or 10080 if for a Magic Link URL] for how long this token is valid. Defaults to 30.
	LinkExpiryMins *int `json:"link_expiry_mins,omitempty"`
	// Whether to generate a Magic Link URL. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
	ShouldCreateMagicLinkUrl *bool `json:"should_create_magic_link_url,omitempty"`
	// Whether to generate a Magic Link URL on the Admin Needed screen during the linking flow. Defaults to false. For more information on Magic Link, see https://merge.dev/blog/integrations-fast-say-hello-to-magic-link.
	HideAdminMagicLink *bool `json:"hide_admin_magic_link,omitempty"`
	// An array of objects to specify the models and fields that will be disabled for a given Linked Account. Each object uses model_id, enabled_actions, and disabled_fields to specify the model, method, and fields that are scoped for a given Linked Account.
	CommonModels []*CommonModelScopesBodyRequest `json:"common_models,omitempty"`
	// When creating a Link Token, you can set permissions for Common Models that will apply to the account that is going to be linked. Any model or field not specified in link token payload will default to existing settings.
	CategoryCommonModelScopes map[string][]*IndividualCommonModelScopeDeserializerRequest `json:"category_common_model_scopes,omitempty"`
	// The following subset of IETF language tags can be used to configure localization.
	//
	// * `en` - en
	// * `de` - de
	Language *LanguageEnum `json:"language,omitempty"`
	// The boolean that indicates whether initial, periodic, and force syncs will be disabled.
	AreSyncsDisabled *bool `json:"are_syncs_disabled,omitempty"`
	// A JSON object containing integration-specific configuration options.
	IntegrationSpecificConfig map[string]interface{} `json:"integration_specific_config,omitempty"`
}
