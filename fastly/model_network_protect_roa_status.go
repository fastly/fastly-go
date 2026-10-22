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

// NetworkProtectRoaStatus Whether a Route Origin Authorization permits Fastly to announce this prefix. Fastly refreshes this on a schedule; reading a network never triggers a check, so this can be older than the moment you read it. A status of `pending` means no check has completed; `valid` authorizes Fastly at this prefix length; `invalid` means a covering ROA does not authorize Fastly; and `not_found` means no ROA covers the prefix. The controller may omit an unrecognized status, and omission is not an authorization result. A valid ROA alone does not make a network ready for advertisement.
type NetworkProtectRoaStatus string

// List of network_protect_roa_status
const (
	NETWORKPROTECTROASTATUS_PENDING   NetworkProtectRoaStatus = "pending"
	NETWORKPROTECTROASTATUS_VALID     NetworkProtectRoaStatus = "valid"
	NETWORKPROTECTROASTATUS_INVALID   NetworkProtectRoaStatus = "invalid"
	NETWORKPROTECTROASTATUS_NOT_FOUND NetworkProtectRoaStatus = "not_found"
)

// AllowedNetworkProtectRoaStatusEnumValues All allowed values of NetworkProtectRoaStatus enum
var AllowedNetworkProtectRoaStatusEnumValues = []NetworkProtectRoaStatus{
	"pending",
	"valid",
	"invalid",
	"not_found",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NetworkProtectRoaStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkProtectRoaStatus(value)
	for _, existing := range AllowedNetworkProtectRoaStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkProtectRoaStatus", value)
}

// NewNetworkProtectRoaStatusFromValue returns a pointer to a valid NetworkProtectRoaStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkProtectRoaStatusFromValue(v string) (*NetworkProtectRoaStatus, error) {
	ev := NetworkProtectRoaStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkProtectRoaStatus: valid values are %v", v, AllowedNetworkProtectRoaStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkProtectRoaStatus) IsValid() bool {
	for _, existing := range AllowedNetworkProtectRoaStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to network_protect_roa_status value
func (v NetworkProtectRoaStatus) Ptr() *NetworkProtectRoaStatus {
	return &v
}

// NullableNetworkProtectRoaStatus is a helper abstraction for handling nullable networkprotectroastatus types.
type NullableNetworkProtectRoaStatus struct {
	value *NetworkProtectRoaStatus
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRoaStatus) Get() *NetworkProtectRoaStatus {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRoaStatus) Set(val *NetworkProtectRoaStatus) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRoaStatus) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRoaStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRoaStatus returns a pointer to a new instance of NullableNetworkProtectRoaStatus.
func NewNullableNetworkProtectRoaStatus(val *NetworkProtectRoaStatus) *NullableNetworkProtectRoaStatus {
	return &NullableNetworkProtectRoaStatus{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRoaStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRoaStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
