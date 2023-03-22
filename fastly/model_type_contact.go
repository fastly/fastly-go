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

// TypeContact Resource type
type TypeContact string

// List of resourceTypecontact
const (
	TYPECONTACT_CONTACT TypeContact = "contact"
)

// AllowedTypeContactEnumValues All allowed values of TypeContact enum
var AllowedTypeContactEnumValues = []TypeContact{
	"contact",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeContact) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeContact(value)
	for _, existing := range AllowedTypeContactEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeContact", value)
}

// NewTypeContactFromValue returns a pointer to a valid TypeContact
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeContactFromValue(v string) (*TypeContact, error) {
	ev := TypeContact(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeContact: valid values are %v", v, AllowedTypeContactEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeContact) IsValid() bool {
	for _, existing := range AllowedTypeContactEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypecontact value
func (v TypeContact) Ptr() *TypeContact {
	return &v
}

// NullableTypeContact is a helper abstraction for handling nullable typecontact types. 
type NullableTypeContact struct {
	value *TypeContact
	isSet bool
}

// Get returns the value.
func (v NullableTypeContact) Get() *TypeContact {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeContact) Set(val *TypeContact) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeContact) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeContact) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeContact returns a pointer to a new instance of NullableTypeContact.
func NewNullableTypeContact(val *TypeContact) *NullableTypeContact {
	return &NullableTypeContact{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeContact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeContact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
