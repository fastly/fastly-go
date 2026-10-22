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

// ArcModelsResponse struct for ArcModelsResponse
type ArcModelsResponse struct {
	// The models returned by the request.
	Data                 []ArcModel    `json:"data"`
	Meta                 ArcModelsMeta `json:"meta"`
	AdditionalProperties map[string]any
}

type _ArcModelsResponse ArcModelsResponse

// NewArcModelsResponse instantiates a new ArcModelsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcModelsResponse(data []ArcModel, meta ArcModelsMeta) *ArcModelsResponse {
	this := ArcModelsResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewArcModelsResponseWithDefaults instantiates a new ArcModelsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcModelsResponseWithDefaults() *ArcModelsResponse {
	this := ArcModelsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ArcModelsResponse) GetData() []ArcModel {
	if o == nil {
		var ret []ArcModel
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ArcModelsResponse) GetDataOk() ([]ArcModel, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ArcModelsResponse) SetData(v []ArcModel) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *ArcModelsResponse) GetMeta() ArcModelsMeta {
	if o == nil {
		var ret ArcModelsMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ArcModelsResponse) GetMetaOk() (*ArcModelsMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ArcModelsResponse) SetMeta(v ArcModelsMeta) {
	o.Meta = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcModelsResponse) MarshalJSON() ([]byte, error) {
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
func (o *ArcModelsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varArcModelsResponse := _ArcModelsResponse{}

	if err = json.Unmarshal(bytes, &varArcModelsResponse); err == nil {
		*o = ArcModelsResponse(varArcModelsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcModelsResponse is a helper abstraction for handling nullable arcmodelsresponse types.
type NullableArcModelsResponse struct {
	value *ArcModelsResponse
	isSet bool
}

// Get returns the value.
func (v NullableArcModelsResponse) Get() *ArcModelsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableArcModelsResponse) Set(val *ArcModelsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcModelsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcModelsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcModelsResponse returns a pointer to a new instance of NullableArcModelsResponse.
func NewNullableArcModelsResponse(val *ArcModelsResponse) *NullableArcModelsResponse {
	return &NullableArcModelsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcModelsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcModelsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
