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

// NetworkProtectRuleIPProtocol The IP protocol to match. Omit to match any protocol.
type NetworkProtectRuleIPProtocol string

// List of network_protect_rule_ip_protocol
const (
	NETWORKPROTECTRULEIPPROTOCOL_TCP    NetworkProtectRuleIPProtocol = "tcp"
	NETWORKPROTECTRULEIPPROTOCOL_UDP    NetworkProtectRuleIPProtocol = "udp"
	NETWORKPROTECTRULEIPPROTOCOL_ICMP   NetworkProtectRuleIPProtocol = "icmp"
	NETWORKPROTECTRULEIPPROTOCOL_ICMPV6 NetworkProtectRuleIPProtocol = "icmpv6"
)

// AllowedNetworkProtectRuleIPProtocolEnumValues All allowed values of NetworkProtectRuleIPProtocol enum
var AllowedNetworkProtectRuleIPProtocolEnumValues = []NetworkProtectRuleIPProtocol{
	"tcp",
	"udp",
	"icmp",
	"icmpv6",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NetworkProtectRuleIPProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkProtectRuleIPProtocol(value)
	for _, existing := range AllowedNetworkProtectRuleIPProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkProtectRuleIPProtocol", value)
}

// NewNetworkProtectRuleIPProtocolFromValue returns a pointer to a valid NetworkProtectRuleIPProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkProtectRuleIPProtocolFromValue(v string) (*NetworkProtectRuleIPProtocol, error) {
	ev := NetworkProtectRuleIPProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkProtectRuleIPProtocol: valid values are %v", v, AllowedNetworkProtectRuleIPProtocolEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkProtectRuleIPProtocol) IsValid() bool {
	for _, existing := range AllowedNetworkProtectRuleIPProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to network_protect_rule_ip_protocol value
func (v NetworkProtectRuleIPProtocol) Ptr() *NetworkProtectRuleIPProtocol {
	return &v
}

// NullableNetworkProtectRuleIPProtocol is a helper abstraction for handling nullable networkprotectruleipprotocol types.
type NullableNetworkProtectRuleIPProtocol struct {
	value *NetworkProtectRuleIPProtocol
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleIPProtocol) Get() *NetworkProtectRuleIPProtocol {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleIPProtocol) Set(val *NetworkProtectRuleIPProtocol) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleIPProtocol) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleIPProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleIPProtocol returns a pointer to a new instance of NullableNetworkProtectRuleIPProtocol.
func NewNullableNetworkProtectRuleIPProtocol(val *NetworkProtectRuleIPProtocol) *NullableNetworkProtectRuleIPProtocol {
	return &NullableNetworkProtectRuleIPProtocol{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleIPProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleIPProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
