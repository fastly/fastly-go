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

// TypeTLSActivation Resource type.
type TypeTLSActivation string

// List of resourceTypetls_activation
const (
	TYPETLSACTIVATION_TLS_ACTIVATION TypeTLSActivation = "tls_activation"
)

// AllowedTypeTLSActivationEnumValues All allowed values of TypeTLSActivation enum
var AllowedTypeTLSActivationEnumValues = []TypeTLSActivation{
	"tls_activation",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeTLSActivation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTLSActivation(value)
	for _, existing := range AllowedTypeTLSActivationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTLSActivation", value)
}

// NewTypeTLSActivationFromValue returns a pointer to a valid TypeTLSActivation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTLSActivationFromValue(v string) (*TypeTLSActivation, error) {
	ev := TypeTLSActivation(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeTLSActivation: valid values are %v", v, AllowedTypeTLSActivationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTLSActivation) IsValid() bool {
	for _, existing := range AllowedTypeTLSActivationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypetls_activation value
func (v TypeTLSActivation) Ptr() *TypeTLSActivation {
	return &v
}

// NullableTypeTLSActivation is a helper abstraction for handling nullable typetlsactivation types. 
type NullableTypeTLSActivation struct {
	value *TypeTLSActivation
	isSet bool
}

// Get returns the value.
func (v NullableTypeTLSActivation) Get() *TypeTLSActivation {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTLSActivation) Set(val *TypeTLSActivation) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTLSActivation) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTLSActivation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTLSActivation returns a pointer to a new instance of NullableTypeTLSActivation.
func NewNullableTypeTLSActivation(val *TypeTLSActivation) *NullableTypeTLSActivation {
	return &NullableTypeTLSActivation{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTLSActivation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeTLSActivation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
