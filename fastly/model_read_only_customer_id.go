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

// ReadOnlyCustomerId struct for ReadOnlyCustomerId
type ReadOnlyCustomerId struct {
}

// NewReadOnlyCustomerId instantiates a new ReadOnlyCustomerId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadOnlyCustomerId() *ReadOnlyCustomerId {
	this := ReadOnlyCustomerId{}
	return &this
}

// NewReadOnlyCustomerIdWithDefaults instantiates a new ReadOnlyCustomerId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadOnlyCustomerIdWithDefaults() *ReadOnlyCustomerId {
	this := ReadOnlyCustomerId{}
	return &this
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ReadOnlyCustomerId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	return json.Marshal(toSerialize)
}

// NullableReadOnlyCustomerId is a helper abstraction for handling nullable readonlycustomerid types.
type NullableReadOnlyCustomerId struct {
	value *ReadOnlyCustomerId
	isSet bool
}

// Get returns the value.
func (v NullableReadOnlyCustomerId) Get() *ReadOnlyCustomerId {
	return v.value
}

// Set modifies the value.
func (v *NullableReadOnlyCustomerId) Set(val *ReadOnlyCustomerId) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableReadOnlyCustomerId) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableReadOnlyCustomerId) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableReadOnlyCustomerId returns a pointer to a new instance of NullableReadOnlyCustomerId.
func NewNullableReadOnlyCustomerId(val *ReadOnlyCustomerId) *NullableReadOnlyCustomerId {
	return &NullableReadOnlyCustomerId{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableReadOnlyCustomerId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableReadOnlyCustomerId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
