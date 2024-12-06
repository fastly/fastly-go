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

// ValuesMisses struct for ValuesMisses
type ValuesMisses struct {
	// The miss rate for requests to the URL in the current dimension.
	MissRate             *float32 `json:"miss_rate,omitempty"`
	AdditionalProperties map[string]any
}

type _ValuesMisses ValuesMisses

// NewValuesMisses instantiates a new ValuesMisses object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuesMisses() *ValuesMisses {
	this := ValuesMisses{}
	return &this
}

// NewValuesMissesWithDefaults instantiates a new ValuesMisses object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuesMissesWithDefaults() *ValuesMisses {
	this := ValuesMisses{}
	return &this
}

// GetMissRate returns the MissRate field value if set, zero value otherwise.
func (o *ValuesMisses) GetMissRate() float32 {
	if o == nil || o.MissRate == nil {
		var ret float32
		return ret
	}
	return *o.MissRate
}

// GetMissRateOk returns a tuple with the MissRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesMisses) GetMissRateOk() (*float32, bool) {
	if o == nil || o.MissRate == nil {
		return nil, false
	}
	return o.MissRate, true
}

// HasMissRate returns a boolean if a field has been set.
func (o *ValuesMisses) HasMissRate() bool {
	if o != nil && o.MissRate != nil {
		return true
	}

	return false
}

// SetMissRate gets a reference to the given float32 and assigns it to the MissRate field.
func (o *ValuesMisses) SetMissRate(v float32) {
	o.MissRate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValuesMisses) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.MissRate != nil {
		toSerialize["miss_rate"] = o.MissRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ValuesMisses) UnmarshalJSON(bytes []byte) (err error) {
	varValuesMisses := _ValuesMisses{}

	if err = json.Unmarshal(bytes, &varValuesMisses); err == nil {
		*o = ValuesMisses(varValuesMisses)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "miss_rate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValuesMisses is a helper abstraction for handling nullable valuesmisses types.
type NullableValuesMisses struct {
	value *ValuesMisses
	isSet bool
}

// Get returns the value.
func (v NullableValuesMisses) Get() *ValuesMisses {
	return v.value
}

// Set modifies the value.
func (v *NullableValuesMisses) Set(val *ValuesMisses) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValuesMisses) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValuesMisses) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValuesMisses returns a pointer to a new instance of NullableValuesMisses.
func NewNullableValuesMisses(val *ValuesMisses) *NullableValuesMisses {
	return &NullableValuesMisses{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValuesMisses) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableValuesMisses) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
