// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://www.fastly.com/documentation/reference/api/)

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.

import (
	"encoding/json"
)

// AutomationTokenWorkspacesResponse struct for AutomationTokenWorkspacesResponse
type AutomationTokenWorkspacesResponse struct {
	// Workspaces the token can access. Empty if the token has no workspace-specific permissions or if `limit_workspaces` is `false` (unrestricted access).
	Data []AutomationTokenWorkspace    `json:"data"`
	Meta AutomationTokenWorkspacesMeta `json:"meta"`
	// Indicates whether workspace access is restricted: - `true`: Token can only access workspaces listed in `data` - `false`: Token has unrestricted access to all workspaces
	LimitWorkspaces      bool `json:"limit_workspaces"`
	AdditionalProperties map[string]any
}

type _AutomationTokenWorkspacesResponse AutomationTokenWorkspacesResponse

// NewAutomationTokenWorkspacesResponse instantiates a new AutomationTokenWorkspacesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenWorkspacesResponse(data []AutomationTokenWorkspace, meta AutomationTokenWorkspacesMeta, limitWorkspaces bool) *AutomationTokenWorkspacesResponse {
	this := AutomationTokenWorkspacesResponse{}
	this.Data = data
	this.Meta = meta
	this.LimitWorkspaces = limitWorkspaces
	return &this
}

// NewAutomationTokenWorkspacesResponseWithDefaults instantiates a new AutomationTokenWorkspacesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenWorkspacesResponseWithDefaults() *AutomationTokenWorkspacesResponse {
	this := AutomationTokenWorkspacesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AutomationTokenWorkspacesResponse) GetData() []AutomationTokenWorkspace {
	if o == nil {
		var ret []AutomationTokenWorkspace
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AutomationTokenWorkspacesResponse) GetDataOk() ([]AutomationTokenWorkspace, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AutomationTokenWorkspacesResponse) SetData(v []AutomationTokenWorkspace) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *AutomationTokenWorkspacesResponse) GetMeta() AutomationTokenWorkspacesMeta {
	if o == nil {
		var ret AutomationTokenWorkspacesMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *AutomationTokenWorkspacesResponse) GetMetaOk() (*AutomationTokenWorkspacesMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *AutomationTokenWorkspacesResponse) SetMeta(v AutomationTokenWorkspacesMeta) {
	o.Meta = v
}

// GetLimitWorkspaces returns the LimitWorkspaces field value
func (o *AutomationTokenWorkspacesResponse) GetLimitWorkspaces() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LimitWorkspaces
}

// GetLimitWorkspacesOk returns a tuple with the LimitWorkspaces field value
// and a boolean to check if the value has been set.
func (o *AutomationTokenWorkspacesResponse) GetLimitWorkspacesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LimitWorkspaces, true
}

// SetLimitWorkspaces sets field value
func (o *AutomationTokenWorkspacesResponse) SetLimitWorkspaces(v bool) {
	o.LimitWorkspaces = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenWorkspacesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}
	if true {
		toSerialize["limit_workspaces"] = o.LimitWorkspaces
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AutomationTokenWorkspacesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenWorkspacesResponse := _AutomationTokenWorkspacesResponse{}

	if err = json.Unmarshal(bytes, &varAutomationTokenWorkspacesResponse); err == nil {
		*o = AutomationTokenWorkspacesResponse(varAutomationTokenWorkspacesResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		delete(additionalProperties, "limit_workspaces")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenWorkspacesResponse is a helper abstraction for handling nullable automationtokenworkspacesresponse types.
type NullableAutomationTokenWorkspacesResponse struct {
	value *AutomationTokenWorkspacesResponse
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenWorkspacesResponse) Get() *AutomationTokenWorkspacesResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenWorkspacesResponse) Set(val *AutomationTokenWorkspacesResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenWorkspacesResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenWorkspacesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenWorkspacesResponse returns a pointer to a new instance of NullableAutomationTokenWorkspacesResponse.
func NewNullableAutomationTokenWorkspacesResponse(val *AutomationTokenWorkspacesResponse) *NullableAutomationTokenWorkspacesResponse {
	return &NullableAutomationTokenWorkspacesResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenWorkspacesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAutomationTokenWorkspacesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
