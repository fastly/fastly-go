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

// TypeWafRuleRevision Resource type.
type TypeWafRuleRevision string

// List of resourceTypewaf_rule_revision
const (
	TYPEWAFRULEREVISION_WAF_RULE_REVISION TypeWafRuleRevision = "waf_rule_revision"
)

// AllowedTypeWafRuleRevisionEnumValues All allowed values of TypeWafRuleRevision enum
var AllowedTypeWafRuleRevisionEnumValues = []TypeWafRuleRevision{
	"waf_rule_revision",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeWafRuleRevision) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeWafRuleRevision(value)
	for _, existing := range AllowedTypeWafRuleRevisionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeWafRuleRevision", value)
}

// NewTypeWafRuleRevisionFromValue returns a pointer to a valid TypeWafRuleRevision
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeWafRuleRevisionFromValue(v string) (*TypeWafRuleRevision, error) {
	ev := TypeWafRuleRevision(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeWafRuleRevision: valid values are %v", v, AllowedTypeWafRuleRevisionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeWafRuleRevision) IsValid() bool {
	for _, existing := range AllowedTypeWafRuleRevisionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypewaf_rule_revision value
func (v TypeWafRuleRevision) Ptr() *TypeWafRuleRevision {
	return &v
}

// NullableTypeWafRuleRevision is a helper abstraction for handling nullable typewafrulerevision types. 
type NullableTypeWafRuleRevision struct {
	value *TypeWafRuleRevision
	isSet bool
}

// Get returns the value.
func (v NullableTypeWafRuleRevision) Get() *TypeWafRuleRevision {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeWafRuleRevision) Set(val *TypeWafRuleRevision) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeWafRuleRevision) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeWafRuleRevision) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeWafRuleRevision returns a pointer to a new instance of NullableTypeWafRuleRevision.
func NewNullableTypeWafRuleRevision(val *TypeWafRuleRevision) *NullableTypeWafRuleRevision {
	return &NullableTypeWafRuleRevision{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeWafRuleRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeWafRuleRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
