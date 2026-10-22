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

// NetworkProtectRuleIpProtocol The IP protocol to match. Omit to match any protocol.
type NetworkProtectRuleIpProtocol string

// List of network_protect_rule_ip_protocol
const (
	NETWORKPROTECTRULEIPPROTOCOL_TCP    NetworkProtectRuleIpProtocol = "tcp"
	NETWORKPROTECTRULEIPPROTOCOL_UDP    NetworkProtectRuleIpProtocol = "udp"
	NETWORKPROTECTRULEIPPROTOCOL_ICMP   NetworkProtectRuleIpProtocol = "icmp"
	NETWORKPROTECTRULEIPPROTOCOL_ICMPV6 NetworkProtectRuleIpProtocol = "icmpv6"
)

// AllowedNetworkProtectRuleIpProtocolEnumValues All allowed values of NetworkProtectRuleIpProtocol enum
var AllowedNetworkProtectRuleIpProtocolEnumValues = []NetworkProtectRuleIpProtocol{
	"tcp",
	"udp",
	"icmp",
	"icmpv6",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NetworkProtectRuleIpProtocol) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkProtectRuleIpProtocol(value)
	for _, existing := range AllowedNetworkProtectRuleIpProtocolEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkProtectRuleIpProtocol", value)
}

// NewNetworkProtectRuleIpProtocolFromValue returns a pointer to a valid NetworkProtectRuleIpProtocol
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkProtectRuleIpProtocolFromValue(v string) (*NetworkProtectRuleIpProtocol, error) {
	ev := NetworkProtectRuleIpProtocol(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkProtectRuleIpProtocol: valid values are %v", v, AllowedNetworkProtectRuleIpProtocolEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkProtectRuleIpProtocol) IsValid() bool {
	for _, existing := range AllowedNetworkProtectRuleIpProtocolEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to network_protect_rule_ip_protocol value
func (v NetworkProtectRuleIpProtocol) Ptr() *NetworkProtectRuleIpProtocol {
	return &v
}

// NullableNetworkProtectRuleIpProtocol is a helper abstraction for handling nullable networkprotectruleipprotocol types.
type NullableNetworkProtectRuleIpProtocol struct {
	value *NetworkProtectRuleIpProtocol
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleIpProtocol) Get() *NetworkProtectRuleIpProtocol {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleIpProtocol) Set(val *NetworkProtectRuleIpProtocol) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleIpProtocol) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleIpProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleIpProtocol returns a pointer to a new instance of NullableNetworkProtectRuleIpProtocol.
func NewNullableNetworkProtectRuleIpProtocol(val *NetworkProtectRuleIpProtocol) *NullableNetworkProtectRuleIpProtocol {
	return &NullableNetworkProtectRuleIpProtocol{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleIpProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleIpProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
