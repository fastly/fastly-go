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

// NetworkProtectLoaStatus Fastly operator status for the Letter of Authorization. A status of `pending` means review is incomplete; `approved` means operators accepted it; and `rejected` means they declined it. An approved LOA alone does not make a network ready for advertisement.
type NetworkProtectLoaStatus string

// List of network_protect_loa_status
const (
	NETWORKPROTECTLOASTATUS_PENDING  NetworkProtectLoaStatus = "pending"
	NETWORKPROTECTLOASTATUS_APPROVED NetworkProtectLoaStatus = "approved"
	NETWORKPROTECTLOASTATUS_REJECTED NetworkProtectLoaStatus = "rejected"
)

// AllowedNetworkProtectLoaStatusEnumValues All allowed values of NetworkProtectLoaStatus enum
var AllowedNetworkProtectLoaStatusEnumValues = []NetworkProtectLoaStatus{
	"pending",
	"approved",
	"rejected",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NetworkProtectLoaStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkProtectLoaStatus(value)
	for _, existing := range AllowedNetworkProtectLoaStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkProtectLoaStatus", value)
}

// NewNetworkProtectLoaStatusFromValue returns a pointer to a valid NetworkProtectLoaStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkProtectLoaStatusFromValue(v string) (*NetworkProtectLoaStatus, error) {
	ev := NetworkProtectLoaStatus(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkProtectLoaStatus: valid values are %v", v, AllowedNetworkProtectLoaStatusEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkProtectLoaStatus) IsValid() bool {
	for _, existing := range AllowedNetworkProtectLoaStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to network_protect_loa_status value
func (v NetworkProtectLoaStatus) Ptr() *NetworkProtectLoaStatus {
	return &v
}

// NullableNetworkProtectLoaStatus is a helper abstraction for handling nullable networkprotectloastatus types.
type NullableNetworkProtectLoaStatus struct {
	value *NetworkProtectLoaStatus
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectLoaStatus) Get() *NetworkProtectLoaStatus {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectLoaStatus) Set(val *NetworkProtectLoaStatus) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectLoaStatus) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectLoaStatus) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectLoaStatus returns a pointer to a new instance of NullableNetworkProtectLoaStatus.
func NewNullableNetworkProtectLoaStatus(val *NetworkProtectLoaStatus) *NullableNetworkProtectLoaStatus {
	return &NullableNetworkProtectLoaStatus{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectLoaStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectLoaStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
