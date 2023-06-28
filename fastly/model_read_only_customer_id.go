// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// ReadOnlyCustomerID struct for ReadOnlyCustomerID
type ReadOnlyCustomerID struct {
}

// NewReadOnlyCustomerID instantiates a new ReadOnlyCustomerID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadOnlyCustomerID() *ReadOnlyCustomerID {
	this := ReadOnlyCustomerID{}
	return &this
}

// NewReadOnlyCustomerIDWithDefaults instantiates a new ReadOnlyCustomerID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadOnlyCustomerIDWithDefaults() *ReadOnlyCustomerID {
	this := ReadOnlyCustomerID{}
	return &this
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ReadOnlyCustomerID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	return json.Marshal(toSerialize)
}

// NullableReadOnlyCustomerID is a helper abstraction for handling nullable readonlycustomerid types. 
type NullableReadOnlyCustomerID struct {
	value *ReadOnlyCustomerID
	isSet bool
}

// Get returns the value.
func (v NullableReadOnlyCustomerID) Get() *ReadOnlyCustomerID {
	return v.value
}

// Set modifies the value.
func (v *NullableReadOnlyCustomerID) Set(val *ReadOnlyCustomerID) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableReadOnlyCustomerID) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableReadOnlyCustomerID) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableReadOnlyCustomerID returns a pointer to a new instance of NullableReadOnlyCustomerID.
func NewNullableReadOnlyCustomerID(val *ReadOnlyCustomerID) *NullableReadOnlyCustomerID {
	return &NullableReadOnlyCustomerID{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableReadOnlyCustomerID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableReadOnlyCustomerID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
