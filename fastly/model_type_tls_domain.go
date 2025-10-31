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

// TypeTlsDomain Resource type
type TypeTlsDomain string

// List of type_tls_domain
const (
	TYPETLSDOMAIN_TLS_DOMAIN TypeTlsDomain = "tls_domain"
)

// AllowedTypeTlsDomainEnumValues All allowed values of TypeTlsDomain enum
var AllowedTypeTlsDomainEnumValues = []TypeTlsDomain{
	"tls_domain",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeTlsDomain) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTlsDomain(value)
	for _, existing := range AllowedTypeTlsDomainEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTlsDomain", value)
}

// NewTypeTlsDomainFromValue returns a pointer to a valid TypeTlsDomain
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTlsDomainFromValue(v string) (*TypeTlsDomain, error) {
	ev := TypeTlsDomain(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeTlsDomain: valid values are %v", v, AllowedTypeTlsDomainEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTlsDomain) IsValid() bool {
	for _, existing := range AllowedTypeTlsDomainEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to type_tls_domain value
func (v TypeTlsDomain) Ptr() *TypeTlsDomain {
	return &v
}

// NullableTypeTlsDomain is a helper abstraction for handling nullable typetlsdomain types.
type NullableTypeTlsDomain struct {
	value *TypeTlsDomain
	isSet bool
}

// Get returns the value.
func (v NullableTypeTlsDomain) Get() *TypeTlsDomain {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTlsDomain) Set(val *TypeTlsDomain) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTlsDomain) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTlsDomain) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTlsDomain returns a pointer to a new instance of NullableTypeTlsDomain.
func NewNullableTypeTlsDomain(val *TypeTlsDomain) *NullableTypeTlsDomain {
	return &NullableTypeTlsDomain{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTlsDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeTlsDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
