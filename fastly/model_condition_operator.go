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

// ConditionOperator The comparison operator used to evaluate the condition.
type ConditionOperator string

// List of condition_operator
const (
	CONDITIONOPERATOR_condition_operator_equals      ConditionOperator = "equals"
	CONDITIONOPERATOR_condition_operator_starts_with ConditionOperator = "starts_with"
	CONDITIONOPERATOR_condition_operator_ends_with   ConditionOperator = "ends_with"
	CONDITIONOPERATOR_condition_operator_contains    ConditionOperator = "contains"
)

// AllowedConditionOperatorEnumValues All allowed values of ConditionOperator enum
var AllowedConditionOperatorEnumValues = []ConditionOperator{
	"equals",
	"starts_with",
	"ends_with",
	"contains",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *ConditionOperator) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConditionOperator(value)
	for _, existing := range AllowedConditionOperatorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConditionOperator", value)
}

// NewConditionOperatorFromValue returns a pointer to a valid ConditionOperator
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConditionOperatorFromValue(v string) (*ConditionOperator, error) {
	ev := ConditionOperator(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for ConditionOperator: valid values are %v", v, AllowedConditionOperatorEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConditionOperator) IsValid() bool {
	for _, existing := range AllowedConditionOperatorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to condition_operator value
func (v ConditionOperator) Ptr() *ConditionOperator {
	return &v
}

// NullableConditionOperator is a helper abstraction for handling nullable conditionoperator types.
type NullableConditionOperator struct {
	value *ConditionOperator
	isSet bool
}

// Get returns the value.
func (v NullableConditionOperator) Get() *ConditionOperator {
	return v.value
}

// Set modifies the value.
func (v *NullableConditionOperator) Set(val *ConditionOperator) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableConditionOperator) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableConditionOperator) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableConditionOperator returns a pointer to a new instance of NullableConditionOperator.
func NewNullableConditionOperator(val *ConditionOperator) *NullableConditionOperator {
	return &NullableConditionOperator{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableConditionOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableConditionOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
