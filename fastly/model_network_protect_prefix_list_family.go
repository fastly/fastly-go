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

// NetworkProtectPrefixListFamily The address family; it cannot change after creation.
type NetworkProtectPrefixListFamily string

// List of network_protect_prefix_list_family
const (
	NETWORKPROTECTPREFIXLISTFAMILY_V4 NetworkProtectPrefixListFamily = "v4"
	NETWORKPROTECTPREFIXLISTFAMILY_V6 NetworkProtectPrefixListFamily = "v6"
)

// AllowedNetworkProtectPrefixListFamilyEnumValues All allowed values of NetworkProtectPrefixListFamily enum
var AllowedNetworkProtectPrefixListFamilyEnumValues = []NetworkProtectPrefixListFamily{
	"v4",
	"v6",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NetworkProtectPrefixListFamily) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkProtectPrefixListFamily(value)
	for _, existing := range AllowedNetworkProtectPrefixListFamilyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkProtectPrefixListFamily", value)
}

// NewNetworkProtectPrefixListFamilyFromValue returns a pointer to a valid NetworkProtectPrefixListFamily
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkProtectPrefixListFamilyFromValue(v string) (*NetworkProtectPrefixListFamily, error) {
	ev := NetworkProtectPrefixListFamily(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkProtectPrefixListFamily: valid values are %v", v, AllowedNetworkProtectPrefixListFamilyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkProtectPrefixListFamily) IsValid() bool {
	for _, existing := range AllowedNetworkProtectPrefixListFamilyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to network_protect_prefix_list_family value
func (v NetworkProtectPrefixListFamily) Ptr() *NetworkProtectPrefixListFamily {
	return &v
}

// NullableNetworkProtectPrefixListFamily is a helper abstraction for handling nullable networkprotectprefixlistfamily types.
type NullableNetworkProtectPrefixListFamily struct {
	value *NetworkProtectPrefixListFamily
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectPrefixListFamily) Get() *NetworkProtectPrefixListFamily {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectPrefixListFamily) Set(val *NetworkProtectPrefixListFamily) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectPrefixListFamily) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectPrefixListFamily) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectPrefixListFamily returns a pointer to a new instance of NullableNetworkProtectPrefixListFamily.
func NewNullableNetworkProtectPrefixListFamily(val *NetworkProtectPrefixListFamily) *NullableNetworkProtectPrefixListFamily {
	return &NullableNetworkProtectPrefixListFamily{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectPrefixListFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectPrefixListFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
