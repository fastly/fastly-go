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

// TypeTlsBulkCertificate Resource type
type TypeTlsBulkCertificate string

// List of type_tls_bulk_certificate
const (
	TYPETLSBULKCERTIFICATE_TLS_BULK_CERTIFICATE TypeTlsBulkCertificate = "tls_bulk_certificate"
)

// AllowedTypeTlsBulkCertificateEnumValues All allowed values of TypeTlsBulkCertificate enum
var AllowedTypeTlsBulkCertificateEnumValues = []TypeTlsBulkCertificate{
	"tls_bulk_certificate",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeTlsBulkCertificate) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTlsBulkCertificate(value)
	for _, existing := range AllowedTypeTlsBulkCertificateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTlsBulkCertificate", value)
}

// NewTypeTlsBulkCertificateFromValue returns a pointer to a valid TypeTlsBulkCertificate
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTlsBulkCertificateFromValue(v string) (*TypeTlsBulkCertificate, error) {
	ev := TypeTlsBulkCertificate(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeTlsBulkCertificate: valid values are %v", v, AllowedTypeTlsBulkCertificateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTlsBulkCertificate) IsValid() bool {
	for _, existing := range AllowedTypeTlsBulkCertificateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to type_tls_bulk_certificate value
func (v TypeTlsBulkCertificate) Ptr() *TypeTlsBulkCertificate {
	return &v
}

// NullableTypeTlsBulkCertificate is a helper abstraction for handling nullable typetlsbulkcertificate types.
type NullableTypeTlsBulkCertificate struct {
	value *TypeTlsBulkCertificate
	isSet bool
}

// Get returns the value.
func (v NullableTypeTlsBulkCertificate) Get() *TypeTlsBulkCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTlsBulkCertificate) Set(val *TypeTlsBulkCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTlsBulkCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTlsBulkCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTlsBulkCertificate returns a pointer to a new instance of NullableTypeTlsBulkCertificate.
func NewNullableTypeTlsBulkCertificate(val *TypeTlsBulkCertificate) *NullableTypeTlsBulkCertificate {
	return &NullableTypeTlsBulkCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTlsBulkCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeTlsBulkCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
