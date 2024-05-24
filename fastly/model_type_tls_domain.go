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

// TypeTLSDomain Resource type
type TypeTLSDomain string

// List of resourceTypetls_domain
const (
	TYPETLSDOMAIN_TLS_DOMAIN TypeTLSDomain = "tls_domain"
)

// AllowedTypeTLSDomainEnumValues All allowed values of TypeTLSDomain enum
var AllowedTypeTLSDomainEnumValues = []TypeTLSDomain{
	"tls_domain",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeTLSDomain) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTLSDomain(value)
	for _, existing := range AllowedTypeTLSDomainEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTLSDomain", value)
}

// NewTypeTLSDomainFromValue returns a pointer to a valid TypeTLSDomain
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTLSDomainFromValue(v string) (*TypeTLSDomain, error) {
	ev := TypeTLSDomain(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeTLSDomain: valid values are %v", v, AllowedTypeTLSDomainEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTLSDomain) IsValid() bool {
	for _, existing := range AllowedTypeTLSDomainEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypetls_domain value
func (v TypeTLSDomain) Ptr() *TypeTLSDomain {
	return &v
}

// NullableTypeTLSDomain is a helper abstraction for handling nullable typetlsdomain types. 
type NullableTypeTLSDomain struct {
	value *TypeTLSDomain
	isSet bool
}

// Get returns the value.
func (v NullableTypeTLSDomain) Get() *TypeTLSDomain {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTLSDomain) Set(val *TypeTLSDomain) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTLSDomain) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTLSDomain) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTLSDomain returns a pointer to a new instance of NullableTypeTLSDomain.
func NewNullableTypeTLSDomain(val *TypeTLSDomain) *NullableTypeTLSDomain {
	return &NullableTypeTLSDomain{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTLSDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeTLSDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
