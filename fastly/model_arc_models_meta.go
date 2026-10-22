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

// ArcModelsMeta struct for ArcModelsMeta
type ArcModelsMeta struct {
	// The count of models returned.
	Total                *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcModelsMeta ArcModelsMeta

// NewArcModelsMeta instantiates a new ArcModelsMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcModelsMeta() *ArcModelsMeta {
	this := ArcModelsMeta{}
	return &this
}

// NewArcModelsMetaWithDefaults instantiates a new ArcModelsMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcModelsMetaWithDefaults() *ArcModelsMeta {
	this := ArcModelsMeta{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ArcModelsMeta) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcModelsMeta) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ArcModelsMeta) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ArcModelsMeta) SetTotal(v int32) {
	o.Total = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcModelsMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcModelsMeta) UnmarshalJSON(bytes []byte) (err error) {
	varArcModelsMeta := _ArcModelsMeta{}

	if err = json.Unmarshal(bytes, &varArcModelsMeta); err == nil {
		*o = ArcModelsMeta(varArcModelsMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcModelsMeta is a helper abstraction for handling nullable arcmodelsmeta types.
type NullableArcModelsMeta struct {
	value *ArcModelsMeta
	isSet bool
}

// Get returns the value.
func (v NullableArcModelsMeta) Get() *ArcModelsMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableArcModelsMeta) Set(val *ArcModelsMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcModelsMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcModelsMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcModelsMeta returns a pointer to a new instance of NullableArcModelsMeta.
func NewNullableArcModelsMeta(val *ArcModelsMeta) *NullableArcModelsMeta {
	return &NullableArcModelsMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcModelsMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcModelsMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
