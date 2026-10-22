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

// BotManagementWorkspacesResponse struct for BotManagementWorkspacesResponse
type BotManagementWorkspacesResponse struct {
	Meta BotManagementMeta `json:"meta"`
	// The workspaces returned by the request.
	Data                 []BotManagementWorkspaceResponse `json:"data"`
	AdditionalProperties map[string]any
}

type _BotManagementWorkspacesResponse BotManagementWorkspacesResponse

// NewBotManagementWorkspacesResponse instantiates a new BotManagementWorkspacesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementWorkspacesResponse(meta BotManagementMeta, data []BotManagementWorkspaceResponse) *BotManagementWorkspacesResponse {
	this := BotManagementWorkspacesResponse{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewBotManagementWorkspacesResponseWithDefaults instantiates a new BotManagementWorkspacesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementWorkspacesResponseWithDefaults() *BotManagementWorkspacesResponse {
	this := BotManagementWorkspacesResponse{}
	return &this
}

// GetMeta returns the Meta field value
func (o *BotManagementWorkspacesResponse) GetMeta() BotManagementMeta {
	if o == nil {
		var ret BotManagementMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspacesResponse) GetMetaOk() (*BotManagementMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *BotManagementWorkspacesResponse) SetMeta(v BotManagementMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *BotManagementWorkspacesResponse) GetData() []BotManagementWorkspaceResponse {
	if o == nil {
		var ret []BotManagementWorkspaceResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspacesResponse) GetDataOk() ([]BotManagementWorkspaceResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BotManagementWorkspacesResponse) SetData(v []BotManagementWorkspaceResponse) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementWorkspacesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["meta"] = o.Meta
	}
	if true {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementWorkspacesResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementWorkspacesResponse := _BotManagementWorkspacesResponse{}

	if err = json.Unmarshal(bytes, &varBotManagementWorkspacesResponse); err == nil {
		*o = BotManagementWorkspacesResponse(varBotManagementWorkspacesResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementWorkspacesResponse is a helper abstraction for handling nullable botmanagementworkspacesresponse types.
type NullableBotManagementWorkspacesResponse struct {
	value *BotManagementWorkspacesResponse
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementWorkspacesResponse) Get() *BotManagementWorkspacesResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementWorkspacesResponse) Set(val *BotManagementWorkspacesResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementWorkspacesResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementWorkspacesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementWorkspacesResponse returns a pointer to a new instance of NullableBotManagementWorkspacesResponse.
func NewNullableBotManagementWorkspacesResponse(val *BotManagementWorkspacesResponse) *NullableBotManagementWorkspacesResponse {
	return &NullableBotManagementWorkspacesResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementWorkspacesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementWorkspacesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
