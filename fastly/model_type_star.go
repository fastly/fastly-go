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

// TypeStar Resource type
type TypeStar string

// List of resourceTypestar
const (
	TYPESTAR_STAR TypeStar = "star"
)

// AllowedTypeStarEnumValues All allowed values of TypeStar enum
var AllowedTypeStarEnumValues = []TypeStar{
	"star",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeStar) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeStar(value)
	for _, existing := range AllowedTypeStarEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeStar", value)
}

// NewTypeStarFromValue returns a pointer to a valid TypeStar
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeStarFromValue(v string) (*TypeStar, error) {
	ev := TypeStar(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeStar: valid values are %v", v, AllowedTypeStarEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeStar) IsValid() bool {
	for _, existing := range AllowedTypeStarEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypestar value
func (v TypeStar) Ptr() *TypeStar {
	return &v
}

// NullableTypeStar is a helper abstraction for handling nullable typestar types. 
type NullableTypeStar struct {
	value *TypeStar
	isSet bool
}

// Get returns the value.
func (v NullableTypeStar) Get() *TypeStar {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeStar) Set(val *TypeStar) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeStar) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeStar) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeStar returns a pointer to a new instance of NullableTypeStar.
func NewNullableTypeStar(val *TypeStar) *NullableTypeStar {
	return &NullableTypeStar{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeStar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeStar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
