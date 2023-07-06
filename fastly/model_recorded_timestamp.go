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

// RecordedTimestamp The Unix timestamp at which this record's data was generated.
type RecordedTimestamp struct {
	AdditionalProperties map[string]any
}

type _RecordedTimestamp RecordedTimestamp

// NewRecordedTimestamp instantiates a new RecordedTimestamp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecordedTimestamp() *RecordedTimestamp {
	this := RecordedTimestamp{}
	return &this
}

// NewRecordedTimestampWithDefaults instantiates a new RecordedTimestamp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordedTimestampWithDefaults() *RecordedTimestamp {
	this := RecordedTimestamp{}
	return &this
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RecordedTimestamp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RecordedTimestamp) UnmarshalJSON(bytes []byte) (err error) {
	varRecordedTimestamp := _RecordedTimestamp{}

	if err = json.Unmarshal(bytes, &varRecordedTimestamp); err == nil {
		*o = RecordedTimestamp(varRecordedTimestamp)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRecordedTimestamp is a helper abstraction for handling nullable recordedtimestamp types. 
type NullableRecordedTimestamp struct {
	value *RecordedTimestamp
	isSet bool
}

// Get returns the value.
func (v NullableRecordedTimestamp) Get() *RecordedTimestamp {
	return v.value
}

// Set modifies the value.
func (v *NullableRecordedTimestamp) Set(val *RecordedTimestamp) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRecordedTimestamp) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRecordedTimestamp) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRecordedTimestamp returns a pointer to a new instance of NullableRecordedTimestamp.
func NewNullableRecordedTimestamp(val *RecordedTimestamp) *NullableRecordedTimestamp {
	return &NullableRecordedTimestamp{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRecordedTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRecordedTimestamp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
