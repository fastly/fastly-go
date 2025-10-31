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

// TypeTlsCertificate Resource type
type TypeTlsCertificate string

// List of type_tls_certificate
const (
	TYPETLSCERTIFICATE_TLS_CERTIFICATE TypeTlsCertificate = "tls_certificate"
)

// AllowedTypeTlsCertificateEnumValues All allowed values of TypeTlsCertificate enum
var AllowedTypeTlsCertificateEnumValues = []TypeTlsCertificate{
	"tls_certificate",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeTlsCertificate) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTlsCertificate(value)
	for _, existing := range AllowedTypeTlsCertificateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTlsCertificate", value)
}

// NewTypeTlsCertificateFromValue returns a pointer to a valid TypeTlsCertificate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTlsCertificateFromValue(v string) (*TypeTlsCertificate, error) {
	ev := TypeTlsCertificate(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeTlsCertificate: valid values are %v", v, AllowedTypeTlsCertificateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTlsCertificate) IsValid() bool {
	for _, existing := range AllowedTypeTlsCertificateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to type_tls_certificate value
func (v TypeTlsCertificate) Ptr() *TypeTlsCertificate {
	return &v
}

// NullableTypeTlsCertificate is a helper abstraction for handling nullable typetlscertificate types.
type NullableTypeTlsCertificate struct {
	value *TypeTlsCertificate
	isSet bool
}

// Get returns the value.
func (v NullableTypeTlsCertificate) Get() *TypeTlsCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTlsCertificate) Set(val *TypeTlsCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTlsCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTlsCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTlsCertificate returns a pointer to a new instance of NullableTypeTlsCertificate.
func NewNullableTypeTlsCertificate(val *TypeTlsCertificate) *NullableTypeTlsCertificate {
	return &NullableTypeTlsCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTlsCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeTlsCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
