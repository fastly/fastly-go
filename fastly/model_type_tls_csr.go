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

// TypeTLSCsr CSR Resource Type
type TypeTLSCsr string

// List of resourceTypetls_csr
const (
	TYPETLSCSR_CSR TypeTLSCsr = "csr"
)

// AllowedTypeTLSCsrEnumValues All allowed values of TypeTLSCsr enum
var AllowedTypeTLSCsrEnumValues = []TypeTLSCsr{
	"csr",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeTLSCsr) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTLSCsr(value)
	for _, existing := range AllowedTypeTLSCsrEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTLSCsr", value)
}

// NewTypeTLSCsrFromValue returns a pointer to a valid TypeTLSCsr
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTLSCsrFromValue(v string) (*TypeTLSCsr, error) {
	ev := TypeTLSCsr(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeTLSCsr: valid values are %v", v, AllowedTypeTLSCsrEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTLSCsr) IsValid() bool {
	for _, existing := range AllowedTypeTLSCsrEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypetls_csr value
func (v TypeTLSCsr) Ptr() *TypeTLSCsr {
	return &v
}

// NullableTypeTLSCsr is a helper abstraction for handling nullable typetlscsr types. 
type NullableTypeTLSCsr struct {
	value *TypeTLSCsr
	isSet bool
}

// Get returns the value.
func (v NullableTypeTLSCsr) Get() *TypeTLSCsr {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTLSCsr) Set(val *TypeTLSCsr) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTLSCsr) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTLSCsr) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTLSCsr returns a pointer to a new instance of NullableTypeTLSCsr.
func NewNullableTypeTLSCsr(val *TypeTLSCsr) *NullableTypeTLSCsr {
	return &NullableTypeTLSCsr{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTLSCsr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeTLSCsr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
