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

// ReadOnlyId struct for ReadOnlyId
type ReadOnlyId struct {
}

// NewReadOnlyId instantiates a new ReadOnlyId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadOnlyId() *ReadOnlyId {
	this := ReadOnlyId{}
	return &this
}

// NewReadOnlyIdWithDefaults instantiates a new ReadOnlyId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadOnlyIdWithDefaults() *ReadOnlyId {
	this := ReadOnlyId{}
	return &this
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ReadOnlyId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	return json.Marshal(toSerialize)
}

// NullableReadOnlyId is a helper abstraction for handling nullable readonlyid types.
type NullableReadOnlyId struct {
	value *ReadOnlyId
	isSet bool
}

// Get returns the value.
func (v NullableReadOnlyId) Get() *ReadOnlyId {
	return v.value
}

// Set modifies the value.
func (v *NullableReadOnlyId) Set(val *ReadOnlyId) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableReadOnlyId) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableReadOnlyId) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableReadOnlyId returns a pointer to a new instance of NullableReadOnlyId.
func NewNullableReadOnlyId(val *ReadOnlyId) *NullableReadOnlyId {
	return &NullableReadOnlyId{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableReadOnlyId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableReadOnlyId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
