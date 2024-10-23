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

// RoleUser The permissions role assigned to the user. Can be `user`, `billing`, `engineer`, or `superuser`.
type RoleUser string

// List of role_user
const (
	ROLEUSER_USER      RoleUser = "user"
	ROLEUSER_BILLING   RoleUser = "billing"
	ROLEUSER_ENGINEER  RoleUser = "engineer"
	ROLEUSER_SUPERUSER RoleUser = "superuser"
)

// AllowedRoleUserEnumValues All allowed values of RoleUser enum
var AllowedRoleUserEnumValues = []RoleUser{
	"user",
	"billing",
	"engineer",
	"superuser",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *RoleUser) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoleUser(value)
	for _, existing := range AllowedRoleUserEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoleUser", value)
}

// NewRoleUserFromValue returns a pointer to a valid RoleUser
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleUserFromValue(v string) (*RoleUser, error) {
	ev := RoleUser(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RoleUser: valid values are %v", v, AllowedRoleUserEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoleUser) IsValid() bool {
	for _, existing := range AllowedRoleUserEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to role_user value
func (v RoleUser) Ptr() *RoleUser {
	return &v
}

// NullableRoleUser is a helper abstraction for handling nullable roleuser types.
type NullableRoleUser struct {
	value *RoleUser
	isSet bool
}

// Get returns the value.
func (v NullableRoleUser) Get() *RoleUser {
	return v.value
}

// Set modifies the value.
func (v *NullableRoleUser) Set(val *RoleUser) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRoleUser) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRoleUser) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRoleUser returns a pointer to a new instance of NullableRoleUser.
func NewNullableRoleUser(val *RoleUser) *NullableRoleUser {
	return &NullableRoleUser{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRoleUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRoleUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
