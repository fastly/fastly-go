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

// ArcSessionsResponse struct for ArcSessionsResponse
type ArcSessionsResponse struct {
	// The sessions returned by the request.
	Data                 []ArcSession    `json:"data"`
	Meta                 ArcSessionsMeta `json:"meta"`
	AdditionalProperties map[string]any
}

type _ArcSessionsResponse ArcSessionsResponse

// NewArcSessionsResponse instantiates a new ArcSessionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcSessionsResponse(data []ArcSession, meta ArcSessionsMeta) *ArcSessionsResponse {
	this := ArcSessionsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewArcSessionsResponseWithDefaults instantiates a new ArcSessionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcSessionsResponseWithDefaults() *ArcSessionsResponse {
	this := ArcSessionsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ArcSessionsResponse) GetData() []ArcSession {
	if o == nil {
		var ret []ArcSession
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ArcSessionsResponse) GetDataOk() ([]ArcSession, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ArcSessionsResponse) SetData(v []ArcSession) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *ArcSessionsResponse) GetMeta() ArcSessionsMeta {
	if o == nil {
		var ret ArcSessionsMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ArcSessionsResponse) GetMetaOk() (*ArcSessionsMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ArcSessionsResponse) SetMeta(v ArcSessionsMeta) {
	o.Meta = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcSessionsResponse) MarshalJSON() ([]byte, error) {
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
func (o *ArcSessionsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varArcSessionsResponse := _ArcSessionsResponse{}

	if err = json.Unmarshal(bytes, &varArcSessionsResponse); err == nil {
		*o = ArcSessionsResponse(varArcSessionsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcSessionsResponse is a helper abstraction for handling nullable arcsessionsresponse types.
type NullableArcSessionsResponse struct {
	value *ArcSessionsResponse
	isSet bool
}

// Get returns the value.
func (v NullableArcSessionsResponse) Get() *ArcSessionsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableArcSessionsResponse) Set(val *ArcSessionsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcSessionsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcSessionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcSessionsResponse returns a pointer to a new instance of NullableArcSessionsResponse.
func NewNullableArcSessionsResponse(val *ArcSessionsResponse) *NullableArcSessionsResponse {
	return &NullableArcSessionsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcSessionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcSessionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
