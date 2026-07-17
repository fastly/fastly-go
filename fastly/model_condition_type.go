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

// ConditionType The type of condition.
type ConditionType string

// List of condition_type
const (
	CONDITIONTYPE_condition_type_header ConditionType = "header"
)

// AllowedConditionTypeEnumValues All allowed values of ConditionType enum
var AllowedConditionTypeEnumValues = []ConditionType{
	"header",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *ConditionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConditionType(value)
	for _, existing := range AllowedConditionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConditionType", value)
}

// NewConditionTypeFromValue returns a pointer to a valid ConditionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConditionTypeFromValue(v string) (*ConditionType, error) {
	ev := ConditionType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConditionType: valid values are %v", v, AllowedConditionTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConditionType) IsValid() bool {
	for _, existing := range AllowedConditionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to condition_type value
func (v ConditionType) Ptr() *ConditionType {
	return &v
}

// NullableConditionType is a helper abstraction for handling nullable conditiontype types.
type NullableConditionType struct {
	value *ConditionType
	isSet bool
}

// Get returns the value.
func (v NullableConditionType) Get() *ConditionType {
	return v.value
}

// Set modifies the value.
func (v *NullableConditionType) Set(val *ConditionType) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableConditionType) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableConditionType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableConditionType returns a pointer to a new instance of NullableConditionType.
func NewNullableConditionType(val *ConditionType) *NullableConditionType {
	return &NullableConditionType{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableConditionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableConditionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
