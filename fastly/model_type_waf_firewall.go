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

// TypeWafFirewall Resource type.
type TypeWafFirewall string

// List of resourceTypewaf_firewall
const (
	TYPEWAFFIREWALL_WAF_FIREWALL TypeWafFirewall = "waf_firewall"
)

// AllowedTypeWafFirewallEnumValues All allowed values of TypeWafFirewall enum
var AllowedTypeWafFirewallEnumValues = []TypeWafFirewall{
	"waf_firewall",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeWafFirewall) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeWafFirewall(value)
	for _, existing := range AllowedTypeWafFirewallEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeWafFirewall", value)
}

// NewTypeWafFirewallFromValue returns a pointer to a valid TypeWafFirewall
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeWafFirewallFromValue(v string) (*TypeWafFirewall, error) {
	ev := TypeWafFirewall(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeWafFirewall: valid values are %v", v, AllowedTypeWafFirewallEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeWafFirewall) IsValid() bool {
	for _, existing := range AllowedTypeWafFirewallEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypewaf_firewall value
func (v TypeWafFirewall) Ptr() *TypeWafFirewall {
	return &v
}

// NullableTypeWafFirewall is a helper abstraction for handling nullable typewaffirewall types. 
type NullableTypeWafFirewall struct {
	value *TypeWafFirewall
	isSet bool
}

// Get returns the value.
func (v NullableTypeWafFirewall) Get() *TypeWafFirewall {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeWafFirewall) Set(val *TypeWafFirewall) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeWafFirewall) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeWafFirewall) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeWafFirewall returns a pointer to a new instance of NullableTypeWafFirewall.
func NewNullableTypeWafFirewall(val *TypeWafFirewall) *NullableTypeWafFirewall {
	return &NullableTypeWafFirewall{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeWafFirewall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeWafFirewall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
