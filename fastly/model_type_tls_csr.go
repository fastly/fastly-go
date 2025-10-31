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

// TypeTlsCsr CSR Resource Type
type TypeTlsCsr string

// List of type_tls_csr
const (
	TYPETLSCSR_CSR TypeTlsCsr = "csr"
)

// AllowedTypeTlsCsrEnumValues All allowed values of TypeTlsCsr enum
var AllowedTypeTlsCsrEnumValues = []TypeTlsCsr{
	"csr",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeTlsCsr) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTlsCsr(value)
	for _, existing := range AllowedTypeTlsCsrEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTlsCsr", value)
}

// NewTypeTlsCsrFromValue returns a pointer to a valid TypeTlsCsr
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTlsCsrFromValue(v string) (*TypeTlsCsr, error) {
	ev := TypeTlsCsr(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeTlsCsr: valid values are %v", v, AllowedTypeTlsCsrEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTlsCsr) IsValid() bool {
	for _, existing := range AllowedTypeTlsCsrEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to type_tls_csr value
func (v TypeTlsCsr) Ptr() *TypeTlsCsr {
	return &v
}

// NullableTypeTlsCsr is a helper abstraction for handling nullable typetlscsr types.
type NullableTypeTlsCsr struct {
	value *TypeTlsCsr
	isSet bool
}

// Get returns the value.
func (v NullableTypeTlsCsr) Get() *TypeTlsCsr {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTlsCsr) Set(val *TypeTlsCsr) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTlsCsr) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTlsCsr) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTlsCsr returns a pointer to a new instance of NullableTypeTlsCsr.
func NewNullableTypeTlsCsr(val *TypeTlsCsr) *NullableTypeTlsCsr {
	return &NullableTypeTlsCsr{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTlsCsr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeTlsCsr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
