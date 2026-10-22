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

// ArcProvidersResponse struct for ArcProvidersResponse
type ArcProvidersResponse struct {
	// The providers returned by the request.
	Data                 []ArcProvider    `json:"data"`
	Meta                 ArcProvidersMeta `json:"meta"`
	AdditionalProperties map[string]any
}

type _ArcProvidersResponse ArcProvidersResponse

// NewArcProvidersResponse instantiates a new ArcProvidersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcProvidersResponse(data []ArcProvider, meta ArcProvidersMeta) *ArcProvidersResponse {
	this := ArcProvidersResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewArcProvidersResponseWithDefaults instantiates a new ArcProvidersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcProvidersResponseWithDefaults() *ArcProvidersResponse {
	this := ArcProvidersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ArcProvidersResponse) GetData() []ArcProvider {
	if o == nil {
		var ret []ArcProvider
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ArcProvidersResponse) GetDataOk() ([]ArcProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ArcProvidersResponse) SetData(v []ArcProvider) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *ArcProvidersResponse) GetMeta() ArcProvidersMeta {
	if o == nil {
		var ret ArcProvidersMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ArcProvidersResponse) GetMetaOk() (*ArcProvidersMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ArcProvidersResponse) SetMeta(v ArcProvidersMeta) {
	o.Meta = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcProvidersResponse) MarshalJSON() ([]byte, error) {
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
func (o *ArcProvidersResponse) UnmarshalJSON(bytes []byte) (err error) {
	varArcProvidersResponse := _ArcProvidersResponse{}

	if err = json.Unmarshal(bytes, &varArcProvidersResponse); err == nil {
		*o = ArcProvidersResponse(varArcProvidersResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcProvidersResponse is a helper abstraction for handling nullable arcprovidersresponse types.
type NullableArcProvidersResponse struct {
	value *ArcProvidersResponse
	isSet bool
}

// Get returns the value.
func (v NullableArcProvidersResponse) Get() *ArcProvidersResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableArcProvidersResponse) Set(val *ArcProvidersResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcProvidersResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcProvidersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcProvidersResponse returns a pointer to a new instance of NullableArcProvidersResponse.
func NewNullableArcProvidersResponse(val *ArcProvidersResponse) *NullableArcProvidersResponse {
	return &NullableArcProvidersResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcProvidersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcProvidersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
