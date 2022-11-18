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

// TypeTLSBulkCertificate Resource type
type TypeTLSBulkCertificate string

// List of resourceTypetls_bulk_certificate
const (
	TYPETLSBULKCERTIFICATE_TLS_BULK_CERTIFICATE TypeTLSBulkCertificate = "tls_bulk_certificate"
)

// AllowedTypeTLSBulkCertificateEnumValues All allowed values of TypeTLSBulkCertificate enum
var AllowedTypeTLSBulkCertificateEnumValues = []TypeTLSBulkCertificate{
	"tls_bulk_certificate",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeTLSBulkCertificate) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTLSBulkCertificate(value)
	for _, existing := range AllowedTypeTLSBulkCertificateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTLSBulkCertificate", value)
}

// NewTypeTLSBulkCertificateFromValue returns a pointer to a valid TypeTLSBulkCertificate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTLSBulkCertificateFromValue(v string) (*TypeTLSBulkCertificate, error) {
	ev := TypeTLSBulkCertificate(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeTLSBulkCertificate: valid values are %v", v, AllowedTypeTLSBulkCertificateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTLSBulkCertificate) IsValid() bool {
	for _, existing := range AllowedTypeTLSBulkCertificateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypetls_bulk_certificate value
func (v TypeTLSBulkCertificate) Ptr() *TypeTLSBulkCertificate {
	return &v
}

// NullableTypeTLSBulkCertificate is a helper abstraction for handling nullable typetlsbulkcertificate types. 
type NullableTypeTLSBulkCertificate struct {
	value *TypeTLSBulkCertificate
	isSet bool
}

// Get returns the value.
func (v NullableTypeTLSBulkCertificate) Get() *TypeTLSBulkCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTLSBulkCertificate) Set(val *TypeTLSBulkCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTLSBulkCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTLSBulkCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTLSBulkCertificate returns a pointer to a new instance of NullableTypeTLSBulkCertificate.
func NewNullableTypeTLSBulkCertificate(val *TypeTLSBulkCertificate) *NullableTypeTLSBulkCertificate {
	return &NullableTypeTLSBulkCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTLSBulkCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeTLSBulkCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
