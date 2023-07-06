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

// OriginInspectorRealtimeEntryRecorded The Unix timestamp at which this record's data was generated.
type OriginInspectorRealtimeEntryRecorded struct {
	AdditionalProperties map[string]any
}

type _OriginInspectorRealtimeEntryRecorded OriginInspectorRealtimeEntryRecorded

// NewOriginInspectorRealtimeEntryRecorded instantiates a new OriginInspectorRealtimeEntryRecorded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspectorRealtimeEntryRecorded() *OriginInspectorRealtimeEntryRecorded {
	this := OriginInspectorRealtimeEntryRecorded{}
	return &this
}

// NewOriginInspectorRealtimeEntryRecordedWithDefaults instantiates a new OriginInspectorRealtimeEntryRecorded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorRealtimeEntryRecordedWithDefaults() *OriginInspectorRealtimeEntryRecorded {
	this := OriginInspectorRealtimeEntryRecorded{}
	return &this
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspectorRealtimeEntryRecorded) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *OriginInspectorRealtimeEntryRecorded) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspectorRealtimeEntryRecorded := _OriginInspectorRealtimeEntryRecorded{}

	if err = json.Unmarshal(bytes, &varOriginInspectorRealtimeEntryRecorded); err == nil {
		*o = OriginInspectorRealtimeEntryRecorded(varOriginInspectorRealtimeEntryRecorded)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOriginInspectorRealtimeEntryRecorded is a helper abstraction for handling nullable origininspectorrealtimeentryrecorded types. 
type NullableOriginInspectorRealtimeEntryRecorded struct {
	value *OriginInspectorRealtimeEntryRecorded
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspectorRealtimeEntryRecorded) Get() *OriginInspectorRealtimeEntryRecorded {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspectorRealtimeEntryRecorded) Set(val *OriginInspectorRealtimeEntryRecorded) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspectorRealtimeEntryRecorded) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspectorRealtimeEntryRecorded) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspectorRealtimeEntryRecorded returns a pointer to a new instance of NullableOriginInspectorRealtimeEntryRecorded.
func NewNullableOriginInspectorRealtimeEntryRecorded(val *OriginInspectorRealtimeEntryRecorded) *NullableOriginInspectorRealtimeEntryRecorded {
	return &NullableOriginInspectorRealtimeEntryRecorded{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspectorRealtimeEntryRecorded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableOriginInspectorRealtimeEntryRecorded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
