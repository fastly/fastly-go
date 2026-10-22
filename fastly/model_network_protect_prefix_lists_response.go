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

// NetworkProtectPrefixListsResponse A paginated collection of Network Protect prefix lists.
type NetworkProtectPrefixListsResponse struct {
	// Prefix lists returned on the current page.
	Data                 []NetworkProtectPrefixListResponse `json:"data"`
	Meta                 NetworkProtectListMeta             `json:"meta"`
	AdditionalProperties map[string]any
}

type _NetworkProtectPrefixListsResponse NetworkProtectPrefixListsResponse

// NewNetworkProtectPrefixListsResponse instantiates a new NetworkProtectPrefixListsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectPrefixListsResponse(data []NetworkProtectPrefixListResponse, meta NetworkProtectListMeta) *NetworkProtectPrefixListsResponse {
	this := NetworkProtectPrefixListsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewNetworkProtectPrefixListsResponseWithDefaults instantiates a new NetworkProtectPrefixListsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectPrefixListsResponseWithDefaults() *NetworkProtectPrefixListsResponse {
	this := NetworkProtectPrefixListsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *NetworkProtectPrefixListsResponse) GetData() []NetworkProtectPrefixListResponse {
	if o == nil {
		var ret []NetworkProtectPrefixListResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListsResponse) GetDataOk() ([]NetworkProtectPrefixListResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *NetworkProtectPrefixListsResponse) SetData(v []NetworkProtectPrefixListResponse) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *NetworkProtectPrefixListsResponse) GetMeta() NetworkProtectListMeta {
	if o == nil {
		var ret NetworkProtectListMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListsResponse) GetMetaOk() (*NetworkProtectListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *NetworkProtectPrefixListsResponse) SetMeta(v NetworkProtectListMeta) {
	o.Meta = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectPrefixListsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectPrefixListsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectPrefixListsResponse := _NetworkProtectPrefixListsResponse{}

	if err = json.Unmarshal(bytes, &varNetworkProtectPrefixListsResponse); err == nil {
		*o = NetworkProtectPrefixListsResponse(varNetworkProtectPrefixListsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectPrefixListsResponse is a helper abstraction for handling nullable networkprotectprefixlistsresponse types.
type NullableNetworkProtectPrefixListsResponse struct {
	value *NetworkProtectPrefixListsResponse
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectPrefixListsResponse) Get() *NetworkProtectPrefixListsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectPrefixListsResponse) Set(val *NetworkProtectPrefixListsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectPrefixListsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectPrefixListsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectPrefixListsResponse returns a pointer to a new instance of NullableNetworkProtectPrefixListsResponse.
func NewNullableNetworkProtectPrefixListsResponse(val *NetworkProtectPrefixListsResponse) *NullableNetworkProtectPrefixListsResponse {
	return &NullableNetworkProtectPrefixListsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectPrefixListsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectPrefixListsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
