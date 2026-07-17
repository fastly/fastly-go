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

// ActionType The type of action to take when a rule matches.
type ActionType string

// List of action_type
const (
	ACTIONTYPE_action_type_service ActionType = "service"
)

// AllowedActionTypeEnumValues All allowed values of ActionType enum
var AllowedActionTypeEnumValues = []ActionType{
	"service",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *ActionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActionType(value)
	for _, existing := range AllowedActionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActionType", value)
}

// NewActionTypeFromValue returns a pointer to a valid ActionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActionTypeFromValue(v string) (*ActionType, error) {
	ev := ActionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ActionType: valid values are %v", v, AllowedActionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActionType) IsValid() bool {
	for _, existing := range AllowedActionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to action_type value
func (v ActionType) Ptr() *ActionType {
	return &v
}

// NullableActionType is a helper abstraction for handling nullable actiontype types.
type NullableActionType struct {
	value *ActionType
	isSet bool
}

// Get returns the value.
func (v NullableActionType) Get() *ActionType {
	return v.value
}

// Set modifies the value.
func (v *NullableActionType) Set(val *ActionType) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableActionType) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableActionType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableActionType returns a pointer to a new instance of NullableActionType.
func NewNullableActionType(val *ActionType) *NullableActionType {
	return &NullableActionType{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableActionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableActionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
