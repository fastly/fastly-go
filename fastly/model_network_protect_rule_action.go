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

// NetworkProtectRuleAction Action for traffic matching this rule.
type NetworkProtectRuleAction string

// List of network_protect_rule_action
const (
	NETWORKPROTECTRULEACTION_ALLOW NetworkProtectRuleAction = "allow"
	NETWORKPROTECTRULEACTION_BLOCK NetworkProtectRuleAction = "block"
)

// AllowedNetworkProtectRuleActionEnumValues All allowed values of NetworkProtectRuleAction enum
var AllowedNetworkProtectRuleActionEnumValues = []NetworkProtectRuleAction{
	"allow",
	"block",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NetworkProtectRuleAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NetworkProtectRuleAction(value)
	for _, existing := range AllowedNetworkProtectRuleActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NetworkProtectRuleAction", value)
}

// NewNetworkProtectRuleActionFromValue returns a pointer to a valid NetworkProtectRuleAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNetworkProtectRuleActionFromValue(v string) (*NetworkProtectRuleAction, error) {
	ev := NetworkProtectRuleAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for NetworkProtectRuleAction: valid values are %v", v, AllowedNetworkProtectRuleActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NetworkProtectRuleAction) IsValid() bool {
	for _, existing := range AllowedNetworkProtectRuleActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to network_protect_rule_action value
func (v NetworkProtectRuleAction) Ptr() *NetworkProtectRuleAction {
	return &v
}

// NullableNetworkProtectRuleAction is a helper abstraction for handling nullable networkprotectruleaction types.
type NullableNetworkProtectRuleAction struct {
	value *NetworkProtectRuleAction
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleAction) Get() *NetworkProtectRuleAction {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleAction) Set(val *NetworkProtectRuleAction) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleAction) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleAction) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleAction returns a pointer to a new instance of NullableNetworkProtectRuleAction.
func NewNullableNetworkProtectRuleAction(val *NetworkProtectRuleAction) *NullableNetworkProtectRuleAction {
	return &NullableNetworkProtectRuleAction{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
