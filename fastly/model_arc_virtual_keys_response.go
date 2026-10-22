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

// ArcVirtualKeysResponse struct for ArcVirtualKeysResponse
type ArcVirtualKeysResponse struct {
	// The virtual keys returned by the request.
	Data                 []ArcVirtualKeyListItemResponse `json:"data"`
	Meta                 ArcVirtualKeysMeta              `json:"meta"`
	AdditionalProperties map[string]any
}

type _ArcVirtualKeysResponse ArcVirtualKeysResponse

// NewArcVirtualKeysResponse instantiates a new ArcVirtualKeysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcVirtualKeysResponse(data []ArcVirtualKeyListItemResponse, meta ArcVirtualKeysMeta) *ArcVirtualKeysResponse {
	this := ArcVirtualKeysResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewArcVirtualKeysResponseWithDefaults instantiates a new ArcVirtualKeysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcVirtualKeysResponseWithDefaults() *ArcVirtualKeysResponse {
	this := ArcVirtualKeysResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ArcVirtualKeysResponse) GetData() []ArcVirtualKeyListItemResponse {
	if o == nil {
		var ret []ArcVirtualKeyListItemResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeysResponse) GetDataOk() ([]ArcVirtualKeyListItemResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ArcVirtualKeysResponse) SetData(v []ArcVirtualKeyListItemResponse) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *ArcVirtualKeysResponse) GetMeta() ArcVirtualKeysMeta {
	if o == nil {
		var ret ArcVirtualKeysMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeysResponse) GetMetaOk() (*ArcVirtualKeysMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ArcVirtualKeysResponse) SetMeta(v ArcVirtualKeysMeta) {
	o.Meta = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcVirtualKeysResponse) MarshalJSON() ([]byte, error) {
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
func (o *ArcVirtualKeysResponse) UnmarshalJSON(bytes []byte) (err error) {
	varArcVirtualKeysResponse := _ArcVirtualKeysResponse{}

	if err = json.Unmarshal(bytes, &varArcVirtualKeysResponse); err == nil {
		*o = ArcVirtualKeysResponse(varArcVirtualKeysResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcVirtualKeysResponse is a helper abstraction for handling nullable arcvirtualkeysresponse types.
type NullableArcVirtualKeysResponse struct {
	value *ArcVirtualKeysResponse
	isSet bool
}

// Get returns the value.
func (v NullableArcVirtualKeysResponse) Get() *ArcVirtualKeysResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableArcVirtualKeysResponse) Set(val *ArcVirtualKeysResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcVirtualKeysResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcVirtualKeysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcVirtualKeysResponse returns a pointer to a new instance of NullableArcVirtualKeysResponse.
func NewNullableArcVirtualKeysResponse(val *ArcVirtualKeysResponse) *NullableArcVirtualKeysResponse {
	return &NullableArcVirtualKeysResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcVirtualKeysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcVirtualKeysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
