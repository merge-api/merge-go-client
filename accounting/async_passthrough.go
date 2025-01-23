// This file was auto-generated by Fern from our API Definition.

package accounting

import (
	json "encoding/json"
	fmt "fmt"
)

type AsyncPassthroughRetrieveResponse struct {
	typeName       string
	RemoteResponse *RemoteResponse
	String         string
}

func NewAsyncPassthroughRetrieveResponseFromRemoteResponse(value *RemoteResponse) *AsyncPassthroughRetrieveResponse {
	return &AsyncPassthroughRetrieveResponse{typeName: "remoteResponse", RemoteResponse: value}
}

func NewAsyncPassthroughRetrieveResponseFromString(value string) *AsyncPassthroughRetrieveResponse {
	return &AsyncPassthroughRetrieveResponse{typeName: "string", String: value}
}

func (a *AsyncPassthroughRetrieveResponse) UnmarshalJSON(data []byte) error {
	valueRemoteResponse := new(RemoteResponse)
	if err := json.Unmarshal(data, &valueRemoteResponse); err == nil {
		a.typeName = "remoteResponse"
		a.RemoteResponse = valueRemoteResponse
		return nil
	}
	var valueString string
	if err := json.Unmarshal(data, &valueString); err == nil {
		a.typeName = "string"
		a.String = valueString
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, a)
}

func (a AsyncPassthroughRetrieveResponse) MarshalJSON() ([]byte, error) {
	switch a.typeName {
	default:
		return nil, fmt.Errorf("invalid type %s in %T", a.typeName, a)
	case "remoteResponse":
		return json.Marshal(a.RemoteResponse)
	case "string":
		return json.Marshal(a.String)
	}
}

type AsyncPassthroughRetrieveResponseVisitor interface {
	VisitRemoteResponse(*RemoteResponse) error
	VisitString(string) error
}

func (a *AsyncPassthroughRetrieveResponse) Accept(visitor AsyncPassthroughRetrieveResponseVisitor) error {
	switch a.typeName {
	default:
		return fmt.Errorf("invalid type %s in %T", a.typeName, a)
	case "remoteResponse":
		return visitor.VisitRemoteResponse(a.RemoteResponse)
	case "string":
		return visitor.VisitString(a.String)
	}
}
