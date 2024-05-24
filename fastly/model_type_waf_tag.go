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

// TypeWafTag Resource type.
type TypeWafTag string

// List of resourceTypewaf_tag
const (
	TYPEWAFTAG_WAF_TAG TypeWafTag = "waf_tag"
)

// AllowedTypeWafTagEnumValues All allowed values of TypeWafTag enum
var AllowedTypeWafTagEnumValues = []TypeWafTag{
	"waf_tag",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeWafTag) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeWafTag(value)
	for _, existing := range AllowedTypeWafTagEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeWafTag", value)
}

// NewTypeWafTagFromValue returns a pointer to a valid TypeWafTag
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeWafTagFromValue(v string) (*TypeWafTag, error) {
	ev := TypeWafTag(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeWafTag: valid values are %v", v, AllowedTypeWafTagEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeWafTag) IsValid() bool {
	for _, existing := range AllowedTypeWafTagEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypewaf_tag value
func (v TypeWafTag) Ptr() *TypeWafTag {
	return &v
}

// NullableTypeWafTag is a helper abstraction for handling nullable typewaftag types. 
type NullableTypeWafTag struct {
	value *TypeWafTag
	isSet bool
}

// Get returns the value.
func (v NullableTypeWafTag) Get() *TypeWafTag {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeWafTag) Set(val *TypeWafTag) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeWafTag) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeWafTag) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeWafTag returns a pointer to a new instance of NullableTypeWafTag.
func NewNullableTypeWafTag(val *TypeWafTag) *NullableTypeWafTag {
	return &NullableTypeWafTag{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeWafTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeWafTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
