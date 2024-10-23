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

// Permission The permission the user has in relation to the service.
type Permission string

// List of permission
const (
	PERMISSION_FULL         Permission = "full"
	PERMISSION_READ_ONLY    Permission = "read_only"
	PERMISSION_PURGE_SELECT Permission = "purge_select"
	PERMISSION_PURGE_ALL    Permission = "purge_all"
)

// AllowedPermissionEnumValues All allowed values of Permission enum
var AllowedPermissionEnumValues = []Permission{
	"full",
	"read_only",
	"purge_select",
	"purge_all",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *Permission) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Permission(value)
	for _, existing := range AllowedPermissionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Permission", value)
}

// NewPermissionFromValue returns a pointer to a valid Permission
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPermissionFromValue(v string) (*Permission, error) {
	ev := Permission(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for Permission: valid values are %v", v, AllowedPermissionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Permission) IsValid() bool {
	for _, existing := range AllowedPermissionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to permission value
func (v Permission) Ptr() *Permission {
	return &v
}

// NullablePermission is a helper abstraction for handling nullable permission types.
type NullablePermission struct {
	value *Permission
	isSet bool
}

// Get returns the value.
func (v NullablePermission) Get() *Permission {
	return v.value
}

// Set modifies the value.
func (v *NullablePermission) Set(val *Permission) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePermission) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePermission) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePermission returns a pointer to a new instance of NullablePermission.
func NewNullablePermission(val *Permission) *NullablePermission {
	return &NullablePermission{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
