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

// BotManagementWorkspaceBotsResponse struct for BotManagementWorkspaceBotsResponse
type BotManagementWorkspaceBotsResponse struct {
	Meta BotManagementMeta `json:"meta"`
	// The bots returned by the request.
	Data                 []BotManagementWorkspaceBotResponse `json:"data"`
	AdditionalProperties map[string]any
}

type _BotManagementWorkspaceBotsResponse BotManagementWorkspaceBotsResponse

// NewBotManagementWorkspaceBotsResponse instantiates a new BotManagementWorkspaceBotsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementWorkspaceBotsResponse(meta BotManagementMeta, data []BotManagementWorkspaceBotResponse) *BotManagementWorkspaceBotsResponse {
	this := BotManagementWorkspaceBotsResponse{}
	this.Meta = meta
	this.Data = data
	return &this
}

// NewBotManagementWorkspaceBotsResponseWithDefaults instantiates a new BotManagementWorkspaceBotsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementWorkspaceBotsResponseWithDefaults() *BotManagementWorkspaceBotsResponse {
	this := BotManagementWorkspaceBotsResponse{}
	return &this
}

// GetMeta returns the Meta field value
func (o *BotManagementWorkspaceBotsResponse) GetMeta() BotManagementMeta {
	if o == nil {
		var ret BotManagementMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceBotsResponse) GetMetaOk() (*BotManagementMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *BotManagementWorkspaceBotsResponse) SetMeta(v BotManagementMeta) {
	o.Meta = v
}

// GetData returns the Data field value
func (o *BotManagementWorkspaceBotsResponse) GetData() []BotManagementWorkspaceBotResponse {
	if o == nil {
		var ret []BotManagementWorkspaceBotResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceBotsResponse) GetDataOk() ([]BotManagementWorkspaceBotResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BotManagementWorkspaceBotsResponse) SetData(v []BotManagementWorkspaceBotResponse) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementWorkspaceBotsResponse) MarshalJSON() ([]byte, error) {
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
func (o *BotManagementWorkspaceBotsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementWorkspaceBotsResponse := _BotManagementWorkspaceBotsResponse{}

	if err = json.Unmarshal(bytes, &varBotManagementWorkspaceBotsResponse); err == nil {
		*o = BotManagementWorkspaceBotsResponse(varBotManagementWorkspaceBotsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementWorkspaceBotsResponse is a helper abstraction for handling nullable botmanagementworkspacebotsresponse types.
type NullableBotManagementWorkspaceBotsResponse struct {
	value *BotManagementWorkspaceBotsResponse
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementWorkspaceBotsResponse) Get() *BotManagementWorkspaceBotsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementWorkspaceBotsResponse) Set(val *BotManagementWorkspaceBotsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementWorkspaceBotsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementWorkspaceBotsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementWorkspaceBotsResponse returns a pointer to a new instance of NullableBotManagementWorkspaceBotsResponse.
func NewNullableBotManagementWorkspaceBotsResponse(val *BotManagementWorkspaceBotsResponse) *NullableBotManagementWorkspaceBotsResponse {
	return &NullableBotManagementWorkspaceBotsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementWorkspaceBotsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementWorkspaceBotsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
