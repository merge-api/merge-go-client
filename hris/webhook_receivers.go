// This file was auto-generated by Fern from our API Definition.

package hris

type WebhookReceiverRequest struct {
	Event    string  `json:"event"`
	IsActive bool    `json:"is_active"`
	Key      *string `json:"key,omitempty"`
}
