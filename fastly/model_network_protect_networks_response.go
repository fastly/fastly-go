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

// NetworkProtectNetworksResponse A paginated collection of Network Protect networks.
type NetworkProtectNetworksResponse struct {
	// Networks returned on the current page.
	Data                 []NetworkProtectNetworkResponse `json:"data"`
	Meta                 NetworkProtectListMeta          `json:"meta"`
	AdditionalProperties map[string]any
}

type _NetworkProtectNetworksResponse NetworkProtectNetworksResponse

// NewNetworkProtectNetworksResponse instantiates a new NetworkProtectNetworksResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectNetworksResponse(data []NetworkProtectNetworkResponse, meta NetworkProtectListMeta) *NetworkProtectNetworksResponse {
	this := NetworkProtectNetworksResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewNetworkProtectNetworksResponseWithDefaults instantiates a new NetworkProtectNetworksResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectNetworksResponseWithDefaults() *NetworkProtectNetworksResponse {
	this := NetworkProtectNetworksResponse{}
	return &this
}

// GetData returns the Data field value
func (o *NetworkProtectNetworksResponse) GetData() []NetworkProtectNetworkResponse {
	if o == nil {
		var ret []NetworkProtectNetworkResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworksResponse) GetDataOk() ([]NetworkProtectNetworkResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *NetworkProtectNetworksResponse) SetData(v []NetworkProtectNetworkResponse) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *NetworkProtectNetworksResponse) GetMeta() NetworkProtectListMeta {
	if o == nil {
		var ret NetworkProtectListMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworksResponse) GetMetaOk() (*NetworkProtectListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *NetworkProtectNetworksResponse) SetMeta(v NetworkProtectListMeta) {
	o.Meta = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectNetworksResponse) MarshalJSON() ([]byte, error) {
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
func (o *NetworkProtectNetworksResponse) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectNetworksResponse := _NetworkProtectNetworksResponse{}

	if err = json.Unmarshal(bytes, &varNetworkProtectNetworksResponse); err == nil {
		*o = NetworkProtectNetworksResponse(varNetworkProtectNetworksResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectNetworksResponse is a helper abstraction for handling nullable networkprotectnetworksresponse types.
type NullableNetworkProtectNetworksResponse struct {
	value *NetworkProtectNetworksResponse
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectNetworksResponse) Get() *NetworkProtectNetworksResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectNetworksResponse) Set(val *NetworkProtectNetworksResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectNetworksResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectNetworksResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectNetworksResponse returns a pointer to a new instance of NullableNetworkProtectNetworksResponse.
func NewNullableNetworkProtectNetworksResponse(val *NetworkProtectNetworksResponse) *NullableNetworkProtectNetworksResponse {
	return &NullableNetworkProtectNetworksResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectNetworksResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectNetworksResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
