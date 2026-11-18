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

// TypeMutualAuthentication Resource type
type TypeMutualAuthentication string

// List of resourceTypemutual_authentication
const (
	TYPEMUTUALAUTHENTICATION_MUTUAL_AUTHENTICATION TypeMutualAuthentication = "mutual_authentication"
)

// AllowedTypeMutualAuthenticationEnumValues All allowed values of TypeMutualAuthentication enum
var AllowedTypeMutualAuthenticationEnumValues = []TypeMutualAuthentication{
	"mutual_authentication",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeMutualAuthentication) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeMutualAuthentication(value)
	for _, existing := range AllowedTypeMutualAuthenticationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeMutualAuthentication", value)
}

// NewTypeMutualAuthenticationFromValue returns a pointer to a valid TypeMutualAuthentication
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeMutualAuthenticationFromValue(v string) (*TypeMutualAuthentication, error) {
	ev := TypeMutualAuthentication(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeMutualAuthentication: valid values are %v", v, AllowedTypeMutualAuthenticationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeMutualAuthentication) IsValid() bool {
	for _, existing := range AllowedTypeMutualAuthenticationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypemutual_authentication value
func (v TypeMutualAuthentication) Ptr() *TypeMutualAuthentication {
	return &v
}

// NullableTypeMutualAuthentication is a helper abstraction for handling nullable typemutualauthentication types.
type NullableTypeMutualAuthentication struct {
	value *TypeMutualAuthentication
	isSet bool
}

// Get returns the value.
func (v NullableTypeMutualAuthentication) Get() *TypeMutualAuthentication {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeMutualAuthentication) Set(val *TypeMutualAuthentication) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeMutualAuthentication) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeMutualAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeMutualAuthentication returns a pointer to a new instance of NullableTypeMutualAuthentication.
func NewNullableTypeMutualAuthentication(val *TypeMutualAuthentication) *NullableTypeMutualAuthentication {
	return &NullableTypeMutualAuthentication{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeMutualAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeMutualAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
