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

// TypeUser Resource type
type TypeUser string

// List of type_user
const (
	TYPEUSER_USER TypeUser = "user"
)

// AllowedTypeUserEnumValues All allowed values of TypeUser enum
var AllowedTypeUserEnumValues = []TypeUser{
	"user",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeUser) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeUser(value)
	for _, existing := range AllowedTypeUserEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeUser", value)
}

// NewTypeUserFromValue returns a pointer to a valid TypeUser
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeUserFromValue(v string) (*TypeUser, error) {
	ev := TypeUser(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeUser: valid values are %v", v, AllowedTypeUserEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeUser) IsValid() bool {
	for _, existing := range AllowedTypeUserEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to type_user value
func (v TypeUser) Ptr() *TypeUser {
	return &v
}

// NullableTypeUser is a helper abstraction for handling nullable typeuser types.
type NullableTypeUser struct {
	value *TypeUser
	isSet bool
}

// Get returns the value.
func (v NullableTypeUser) Get() *TypeUser {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeUser) Set(val *TypeUser) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeUser) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeUser) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeUser returns a pointer to a new instance of NullableTypeUser.
func NewNullableTypeUser(val *TypeUser) *NullableTypeUser {
	return &NullableTypeUser{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
