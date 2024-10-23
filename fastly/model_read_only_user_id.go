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

// ReadOnlyUserID struct for ReadOnlyUserID
type ReadOnlyUserID struct {
}

// NewReadOnlyUserID instantiates a new ReadOnlyUserID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadOnlyUserID() *ReadOnlyUserID {
	this := ReadOnlyUserID{}
	return &this
}

// NewReadOnlyUserIDWithDefaults instantiates a new ReadOnlyUserID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadOnlyUserIDWithDefaults() *ReadOnlyUserID {
	this := ReadOnlyUserID{}
	return &this
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ReadOnlyUserID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	return json.Marshal(toSerialize)
}

// NullableReadOnlyUserID is a helper abstraction for handling nullable readonlyuserid types.
type NullableReadOnlyUserID struct {
	value *ReadOnlyUserID
	isSet bool
}

// Get returns the value.
func (v NullableReadOnlyUserID) Get() *ReadOnlyUserID {
	return v.value
}

// Set modifies the value.
func (v *NullableReadOnlyUserID) Set(val *ReadOnlyUserID) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableReadOnlyUserID) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableReadOnlyUserID) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableReadOnlyUserID returns a pointer to a new instance of NullableReadOnlyUserID.
func NewNullableReadOnlyUserID(val *ReadOnlyUserID) *NullableReadOnlyUserID {
	return &NullableReadOnlyUserID{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableReadOnlyUserID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableReadOnlyUserID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
