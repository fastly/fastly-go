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

// TypeTLSCertificate Resource type
type TypeTLSCertificate string

// List of resourceTypetls_certificate
const (
	TYPETLSCERTIFICATE_TLS_CERTIFICATE TypeTLSCertificate = "tls_certificate"
)

// AllowedTypeTLSCertificateEnumValues All allowed values of TypeTLSCertificate enum
var AllowedTypeTLSCertificateEnumValues = []TypeTLSCertificate{
	"tls_certificate",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeTLSCertificate) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTLSCertificate(value)
	for _, existing := range AllowedTypeTLSCertificateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTLSCertificate", value)
}

// NewTypeTLSCertificateFromValue returns a pointer to a valid TypeTLSCertificate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTLSCertificateFromValue(v string) (*TypeTLSCertificate, error) {
	ev := TypeTLSCertificate(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeTLSCertificate: valid values are %v", v, AllowedTypeTLSCertificateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTLSCertificate) IsValid() bool {
	for _, existing := range AllowedTypeTLSCertificateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypetls_certificate value
func (v TypeTLSCertificate) Ptr() *TypeTLSCertificate {
	return &v
}

// NullableTypeTLSCertificate is a helper abstraction for handling nullable typetlscertificate types. 
type NullableTypeTLSCertificate struct {
	value *TypeTLSCertificate
	isSet bool
}

// Get returns the value.
func (v NullableTypeTLSCertificate) Get() *TypeTLSCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTLSCertificate) Set(val *TypeTLSCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTLSCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTLSCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTLSCertificate returns a pointer to a new instance of NullableTypeTLSCertificate.
func NewNullableTypeTLSCertificate(val *TypeTLSCertificate) *NullableTypeTLSCertificate {
	return &NullableTypeTLSCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTLSCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeTLSCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
