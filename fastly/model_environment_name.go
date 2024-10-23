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

// EnvironmentName the model 'EnvironmentName'
type EnvironmentName string

// List of environment_name
const (
	ENVIRONMENTNAME_STAGING EnvironmentName = "staging"
)

// AllowedEnvironmentNameEnumValues All allowed values of EnvironmentName enum
var AllowedEnvironmentNameEnumValues = []EnvironmentName{
	"staging",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *EnvironmentName) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnvironmentName(value)
	for _, existing := range AllowedEnvironmentNameEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnvironmentName", value)
}

// NewEnvironmentNameFromValue returns a pointer to a valid EnvironmentName
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnvironmentNameFromValue(v string) (*EnvironmentName, error) {
	ev := EnvironmentName(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for EnvironmentName: valid values are %v", v, AllowedEnvironmentNameEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnvironmentName) IsValid() bool {
	for _, existing := range AllowedEnvironmentNameEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to environment_name value
func (v EnvironmentName) Ptr() *EnvironmentName {
	return &v
}

// NullableEnvironmentName is a helper abstraction for handling nullable environmentname types.
type NullableEnvironmentName struct {
	value *EnvironmentName
	isSet bool
}

// Get returns the value.
func (v NullableEnvironmentName) Get() *EnvironmentName {
	return v.value
}

// Set modifies the value.
func (v *NullableEnvironmentName) Set(val *EnvironmentName) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableEnvironmentName) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableEnvironmentName) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableEnvironmentName returns a pointer to a new instance of NullableEnvironmentName.
func NewNullableEnvironmentName(val *EnvironmentName) *NullableEnvironmentName {
	return &NullableEnvironmentName{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableEnvironmentName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableEnvironmentName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
