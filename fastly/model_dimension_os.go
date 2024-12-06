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

// DimensionOs struct for DimensionOs
type DimensionOs struct {
	// The client's operating system for this dimension.
	Os                   *string `json:"os,omitempty"`
	AdditionalProperties map[string]any
}

type _DimensionOs DimensionOs

// NewDimensionOs instantiates a new DimensionOs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionOs() *DimensionOs {
	this := DimensionOs{}
	return &this
}

// NewDimensionOsWithDefaults instantiates a new DimensionOs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionOsWithDefaults() *DimensionOs {
	this := DimensionOs{}
	return &this
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *DimensionOs) GetOs() string {
	if o == nil || o.Os == nil {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionOs) GetOsOk() (*string, bool) {
	if o == nil || o.Os == nil {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *DimensionOs) HasOs() bool {
	if o != nil && o.Os != nil {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *DimensionOs) SetOs(v string) {
	o.Os = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DimensionOs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Os != nil {
		toSerialize["os"] = o.Os
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DimensionOs) UnmarshalJSON(bytes []byte) (err error) {
	varDimensionOs := _DimensionOs{}

	if err = json.Unmarshal(bytes, &varDimensionOs); err == nil {
		*o = DimensionOs(varDimensionOs)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "os")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDimensionOs is a helper abstraction for handling nullable dimensionos types.
type NullableDimensionOs struct {
	value *DimensionOs
	isSet bool
}

// Get returns the value.
func (v NullableDimensionOs) Get() *DimensionOs {
	return v.value
}

// Set modifies the value.
func (v *NullableDimensionOs) Set(val *DimensionOs) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDimensionOs) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDimensionOs) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDimensionOs returns a pointer to a new instance of NullableDimensionOs.
func NewNullableDimensionOs(val *DimensionOs) *NullableDimensionOs {
	return &NullableDimensionOs{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDimensionOs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDimensionOs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
