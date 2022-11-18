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

// TypeWafActiveRule Resource type.
type TypeWafActiveRule string

// List of resourceTypewaf_active_rule
const (
	TYPEWAFACTIVERULE_WAF_ACTIVE_RULE TypeWafActiveRule = "waf_active_rule"
)

// AllowedTypeWafActiveRuleEnumValues All allowed values of TypeWafActiveRule enum
var AllowedTypeWafActiveRuleEnumValues = []TypeWafActiveRule{
	"waf_active_rule",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeWafActiveRule) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeWafActiveRule(value)
	for _, existing := range AllowedTypeWafActiveRuleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeWafActiveRule", value)
}

// NewTypeWafActiveRuleFromValue returns a pointer to a valid TypeWafActiveRule
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeWafActiveRuleFromValue(v string) (*TypeWafActiveRule, error) {
	ev := TypeWafActiveRule(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeWafActiveRule: valid values are %v", v, AllowedTypeWafActiveRuleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeWafActiveRule) IsValid() bool {
	for _, existing := range AllowedTypeWafActiveRuleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypewaf_active_rule value
func (v TypeWafActiveRule) Ptr() *TypeWafActiveRule {
	return &v
}

// NullableTypeWafActiveRule is a helper abstraction for handling nullable typewafactiverule types. 
type NullableTypeWafActiveRule struct {
	value *TypeWafActiveRule
	isSet bool
}

// Get returns the value.
func (v NullableTypeWafActiveRule) Get() *TypeWafActiveRule {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeWafActiveRule) Set(val *TypeWafActiveRule) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeWafActiveRule) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeWafActiveRule) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeWafActiveRule returns a pointer to a new instance of NullableTypeWafActiveRule.
func NewNullableTypeWafActiveRule(val *TypeWafActiveRule) *NullableTypeWafActiveRule {
	return &NullableTypeWafActiveRule{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeWafActiveRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeWafActiveRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
