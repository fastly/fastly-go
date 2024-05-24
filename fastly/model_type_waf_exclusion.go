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

// TypeWafExclusion Resource type.
type TypeWafExclusion string

// List of resourceTypewaf_exclusion
const (
	TYPEWAFEXCLUSION_WAF_EXCLUSION TypeWafExclusion = "waf_exclusion"
)

// AllowedTypeWafExclusionEnumValues All allowed values of TypeWafExclusion enum
var AllowedTypeWafExclusionEnumValues = []TypeWafExclusion{
	"waf_exclusion",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeWafExclusion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeWafExclusion(value)
	for _, existing := range AllowedTypeWafExclusionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeWafExclusion", value)
}

// NewTypeWafExclusionFromValue returns a pointer to a valid TypeWafExclusion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeWafExclusionFromValue(v string) (*TypeWafExclusion, error) {
	ev := TypeWafExclusion(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeWafExclusion: valid values are %v", v, AllowedTypeWafExclusionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeWafExclusion) IsValid() bool {
	for _, existing := range AllowedTypeWafExclusionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypewaf_exclusion value
func (v TypeWafExclusion) Ptr() *TypeWafExclusion {
	return &v
}

// NullableTypeWafExclusion is a helper abstraction for handling nullable typewafexclusion types. 
type NullableTypeWafExclusion struct {
	value *TypeWafExclusion
	isSet bool
}

// Get returns the value.
func (v NullableTypeWafExclusion) Get() *TypeWafExclusion {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeWafExclusion) Set(val *TypeWafExclusion) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeWafExclusion) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeWafExclusion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeWafExclusion returns a pointer to a new instance of NullableTypeWafExclusion.
func NewNullableTypeWafExclusion(val *TypeWafExclusion) *NullableTypeWafExclusion {
	return &NullableTypeWafExclusion{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeWafExclusion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeWafExclusion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
