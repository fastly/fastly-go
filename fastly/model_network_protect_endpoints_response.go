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

// NetworkProtectEndpointsResponse A paginated collection of Network Protect endpoints.
type NetworkProtectEndpointsResponse struct {
	// Endpoints returned on the current page.
	Data                 []NetworkProtectEndpointResponse `json:"data"`
	Meta                 NetworkProtectListMeta           `json:"meta"`
	AdditionalProperties map[string]any
}

type _NetworkProtectEndpointsResponse NetworkProtectEndpointsResponse

// NewNetworkProtectEndpointsResponse instantiates a new NetworkProtectEndpointsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectEndpointsResponse(data []NetworkProtectEndpointResponse, meta NetworkProtectListMeta) *NetworkProtectEndpointsResponse {
	this := NetworkProtectEndpointsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewNetworkProtectEndpointsResponseWithDefaults instantiates a new NetworkProtectEndpointsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectEndpointsResponseWithDefaults() *NetworkProtectEndpointsResponse {
	this := NetworkProtectEndpointsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *NetworkProtectEndpointsResponse) GetData() []NetworkProtectEndpointResponse {
	if o == nil {
		var ret []NetworkProtectEndpointResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpointsResponse) GetDataOk() ([]NetworkProtectEndpointResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *NetworkProtectEndpointsResponse) SetData(v []NetworkProtectEndpointResponse) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *NetworkProtectEndpointsResponse) GetMeta() NetworkProtectListMeta {
	if o == nil {
		var ret NetworkProtectListMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpointsResponse) GetMetaOk() (*NetworkProtectListMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *NetworkProtectEndpointsResponse) SetMeta(v NetworkProtectListMeta) {
	o.Meta = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectEndpointsResponse) MarshalJSON() ([]byte, error) {
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
func (o *NetworkProtectEndpointsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectEndpointsResponse := _NetworkProtectEndpointsResponse{}

	if err = json.Unmarshal(bytes, &varNetworkProtectEndpointsResponse); err == nil {
		*o = NetworkProtectEndpointsResponse(varNetworkProtectEndpointsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectEndpointsResponse is a helper abstraction for handling nullable networkprotectendpointsresponse types.
type NullableNetworkProtectEndpointsResponse struct {
	value *NetworkProtectEndpointsResponse
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectEndpointsResponse) Get() *NetworkProtectEndpointsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectEndpointsResponse) Set(val *NetworkProtectEndpointsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectEndpointsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectEndpointsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectEndpointsResponse returns a pointer to a new instance of NullableNetworkProtectEndpointsResponse.
func NewNullableNetworkProtectEndpointsResponse(val *NetworkProtectEndpointsResponse) *NullableNetworkProtectEndpointsResponse {
	return &NullableNetworkProtectEndpointsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectEndpointsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectEndpointsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
