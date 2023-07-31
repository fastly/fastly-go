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

// SubsequentRequestTimestamp Value to use for subsequent requests.
type SubsequentRequestTimestamp struct {
}

// NewSubsequentRequestTimestamp instantiates a new SubsequentRequestTimestamp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubsequentRequestTimestamp() *SubsequentRequestTimestamp {
	this := SubsequentRequestTimestamp{}
	return &this
}

// NewSubsequentRequestTimestampWithDefaults instantiates a new SubsequentRequestTimestamp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubsequentRequestTimestampWithDefaults() *SubsequentRequestTimestamp {
	this := SubsequentRequestTimestamp{}
	return &this
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SubsequentRequestTimestamp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	return json.Marshal(toSerialize)
}

// NullableSubsequentRequestTimestamp is a helper abstraction for handling nullable subsequentrequesttimestamp types. 
type NullableSubsequentRequestTimestamp struct {
	value *SubsequentRequestTimestamp
	isSet bool
}

// Get returns the value.
func (v NullableSubsequentRequestTimestamp) Get() *SubsequentRequestTimestamp {
	return v.value
}

// Set modifies the value.
func (v *NullableSubsequentRequestTimestamp) Set(val *SubsequentRequestTimestamp) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSubsequentRequestTimestamp) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSubsequentRequestTimestamp) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSubsequentRequestTimestamp returns a pointer to a new instance of NullableSubsequentRequestTimestamp.
func NewNullableSubsequentRequestTimestamp(val *SubsequentRequestTimestamp) *NullableSubsequentRequestTimestamp {
	return &NullableSubsequentRequestTimestamp{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSubsequentRequestTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableSubsequentRequestTimestamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
