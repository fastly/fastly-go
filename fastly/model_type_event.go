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
	"fmt"
)

// TypeEvent Resource type
type TypeEvent string

// List of type_event
const (
	TYPEEVENT_EVENT TypeEvent = "event"
)

// AllowedTypeEventEnumValues All allowed values of TypeEvent enum
var AllowedTypeEventEnumValues = []TypeEvent{
	"event",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeEvent) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeEvent(value)
	for _, existing := range AllowedTypeEventEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeEvent", value)
}

// NewTypeEventFromValue returns a pointer to a valid TypeEvent
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeEventFromValue(v string) (*TypeEvent, error) {
	ev := TypeEvent(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeEvent: valid values are %v", v, AllowedTypeEventEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeEvent) IsValid() bool {
	for _, existing := range AllowedTypeEventEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to type_event value
func (v TypeEvent) Ptr() *TypeEvent {
	return &v
}

// NullableTypeEvent is a helper abstraction for handling nullable typeevent types.
type NullableTypeEvent struct {
	value *TypeEvent
	isSet bool
}

// Get returns the value.
func (v NullableTypeEvent) Get() *TypeEvent {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeEvent) Set(val *TypeEvent) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeEvent) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeEvent) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeEvent returns a pointer to a new instance of NullableTypeEvent.
func NewNullableTypeEvent(val *TypeEvent) *NullableTypeEvent {
	return &NullableTypeEvent{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
