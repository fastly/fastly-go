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

// TypeWafRule Resource type.
type TypeWafRule string

// List of resourceTypewaf_rule
const (
	TYPEWAFRULE_WAF_RULE TypeWafRule = "waf_rule"
)

// AllowedTypeWafRuleEnumValues All allowed values of TypeWafRule enum
var AllowedTypeWafRuleEnumValues = []TypeWafRule{
	"waf_rule",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeWafRule) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeWafRule(value)
	for _, existing := range AllowedTypeWafRuleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeWafRule", value)
}

// NewTypeWafRuleFromValue returns a pointer to a valid TypeWafRule
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeWafRuleFromValue(v string) (*TypeWafRule, error) {
	ev := TypeWafRule(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeWafRule: valid values are %v", v, AllowedTypeWafRuleEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeWafRule) IsValid() bool {
	for _, existing := range AllowedTypeWafRuleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypewaf_rule value
func (v TypeWafRule) Ptr() *TypeWafRule {
	return &v
}

// NullableTypeWafRule is a helper abstraction for handling nullable typewafrule types.
type NullableTypeWafRule struct {
	value *TypeWafRule
	isSet bool
}

// Get returns the value.
func (v NullableTypeWafRule) Get() *TypeWafRule {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeWafRule) Set(val *TypeWafRule) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeWafRule) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeWafRule) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeWafRule returns a pointer to a new instance of NullableTypeWafRule.
func NewNullableTypeWafRule(val *TypeWafRule) *NullableTypeWafRule {
	return &NullableTypeWafRule{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeWafRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeWafRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
