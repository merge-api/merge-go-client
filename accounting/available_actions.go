// Code generated by Fern. DO NOT EDIT.

package accounting

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/merge-api/merge-go-client/v2/internal"
)

// # The AvailableActions Object
// ### Description
// The `Activity` object is used to see all available model/operation combinations for an integration.
//
// ### Usage Example
// Fetch all the actions available for the `Zenefits` integration.
type AvailableActions struct {
	Integration              *AccountIntegration `json:"integration" url:"integration"`
	PassthroughAvailable     bool                `json:"passthrough_available" url:"passthrough_available"`
	AvailableModelOperations []*ModelOperation   `json:"available_model_operations,omitempty" url:"available_model_operations,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (a *AvailableActions) GetIntegration() *AccountIntegration {
	if a == nil {
		return nil
	}
	return a.Integration
}

func (a *AvailableActions) GetPassthroughAvailable() bool {
	if a == nil {
		return false
	}
	return a.PassthroughAvailable
}

func (a *AvailableActions) GetAvailableModelOperations() []*ModelOperation {
	if a == nil {
		return nil
	}
	return a.AvailableModelOperations
}

func (a *AvailableActions) GetExtraProperties() map[string]interface{} {
	return a.extraProperties
}

func (a *AvailableActions) UnmarshalJSON(data []byte) error {
	type unmarshaler AvailableActions
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AvailableActions(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	a.rawJSON = json.RawMessage(data)
	return nil
}

func (a *AvailableActions) String() string {
	if len(a.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(a.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(a); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", a)
}
