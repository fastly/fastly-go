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

// RealtimeEntryRecorded The Unix timestamp at which this record's data was generated.
type RealtimeEntryRecorded struct {
}

// NewRealtimeEntryRecorded instantiates a new RealtimeEntryRecorded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRealtimeEntryRecorded() *RealtimeEntryRecorded {
	this := RealtimeEntryRecorded{}
	return &this
}

// NewRealtimeEntryRecordedWithDefaults instantiates a new RealtimeEntryRecorded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRealtimeEntryRecordedWithDefaults() *RealtimeEntryRecorded {
	this := RealtimeEntryRecorded{}
	return &this
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RealtimeEntryRecorded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	return json.Marshal(toSerialize)
}

// NullableRealtimeEntryRecorded is a helper abstraction for handling nullable realtimeentryrecorded types.
type NullableRealtimeEntryRecorded struct {
	value *RealtimeEntryRecorded
	isSet bool
}

// Get returns the value.
func (v NullableRealtimeEntryRecorded) Get() *RealtimeEntryRecorded {
	return v.value
}

// Set modifies the value.
func (v *NullableRealtimeEntryRecorded) Set(val *RealtimeEntryRecorded) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRealtimeEntryRecorded) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRealtimeEntryRecorded) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRealtimeEntryRecorded returns a pointer to a new instance of NullableRealtimeEntryRecorded.
func NewNullableRealtimeEntryRecorded(val *RealtimeEntryRecorded) *NullableRealtimeEntryRecorded {
	return &NullableRealtimeEntryRecorded{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRealtimeEntryRecorded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRealtimeEntryRecorded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
