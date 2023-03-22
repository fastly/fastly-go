// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
	"fmt"
)

// TypeWafFirewallVersion Resource type.
type TypeWafFirewallVersion string

// List of resourceTypewaf_firewall_version
const (
	TYPEWAFFIREWALLVERSION_WAF_FIREWALL_VERSION TypeWafFirewallVersion = "waf_firewall_version"
)

// AllowedTypeWafFirewallVersionEnumValues All allowed values of TypeWafFirewallVersion enum
var AllowedTypeWafFirewallVersionEnumValues = []TypeWafFirewallVersion{
	"waf_firewall_version",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeWafFirewallVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeWafFirewallVersion(value)
	for _, existing := range AllowedTypeWafFirewallVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeWafFirewallVersion", value)
}

// NewTypeWafFirewallVersionFromValue returns a pointer to a valid TypeWafFirewallVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeWafFirewallVersionFromValue(v string) (*TypeWafFirewallVersion, error) {
	ev := TypeWafFirewallVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeWafFirewallVersion: valid values are %v", v, AllowedTypeWafFirewallVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeWafFirewallVersion) IsValid() bool {
	for _, existing := range AllowedTypeWafFirewallVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypewaf_firewall_version value
func (v TypeWafFirewallVersion) Ptr() *TypeWafFirewallVersion {
	return &v
}

// NullableTypeWafFirewallVersion is a helper abstraction for handling nullable typewaffirewallversion types. 
type NullableTypeWafFirewallVersion struct {
	value *TypeWafFirewallVersion
	isSet bool
}

// Get returns the value.
func (v NullableTypeWafFirewallVersion) Get() *TypeWafFirewallVersion {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeWafFirewallVersion) Set(val *TypeWafFirewallVersion) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeWafFirewallVersion) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeWafFirewallVersion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeWafFirewallVersion returns a pointer to a new instance of NullableTypeWafFirewallVersion.
func NewNullableTypeWafFirewallVersion(val *TypeWafFirewallVersion) *NullableTypeWafFirewallVersion {
	return &NullableTypeWafFirewallVersion{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeWafFirewallVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeWafFirewallVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
