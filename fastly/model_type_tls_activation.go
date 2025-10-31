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

// TypeTlsActivation Resource type.
type TypeTlsActivation string

// List of type_tls_activation
const (
	TYPETLSACTIVATION_TLS_ACTIVATION TypeTlsActivation = "tls_activation"
)

// AllowedTypeTlsActivationEnumValues All allowed values of TypeTlsActivation enum
var AllowedTypeTlsActivationEnumValues = []TypeTlsActivation{
	"tls_activation",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeTlsActivation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTlsActivation(value)
	for _, existing := range AllowedTypeTlsActivationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTlsActivation", value)
}

// NewTypeTlsActivationFromValue returns a pointer to a valid TypeTlsActivation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTlsActivationFromValue(v string) (*TypeTlsActivation, error) {
	ev := TypeTlsActivation(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeTlsActivation: valid values are %v", v, AllowedTypeTlsActivationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTlsActivation) IsValid() bool {
	for _, existing := range AllowedTypeTlsActivationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to type_tls_activation value
func (v TypeTlsActivation) Ptr() *TypeTlsActivation {
	return &v
}

// NullableTypeTlsActivation is a helper abstraction for handling nullable typetlsactivation types.
type NullableTypeTlsActivation struct {
	value *TypeTlsActivation
	isSet bool
}

// Get returns the value.
func (v NullableTypeTlsActivation) Get() *TypeTlsActivation {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTlsActivation) Set(val *TypeTlsActivation) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTlsActivation) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTlsActivation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTlsActivation returns a pointer to a new instance of NullableTypeTlsActivation.
func NewNullableTypeTlsActivation(val *TypeTlsActivation) *NullableTypeTlsActivation {
	return &NullableTypeTlsActivation{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTlsActivation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeTlsActivation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
