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

// TypeTlsPrivateKey Resource type
type TypeTlsPrivateKey string

// List of type_tls_private_key
const (
	TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY TypeTlsPrivateKey = "tls_private_key"
)

// AllowedTypeTlsPrivateKeyEnumValues All allowed values of TypeTlsPrivateKey enum
var AllowedTypeTlsPrivateKeyEnumValues = []TypeTlsPrivateKey{
	"tls_private_key",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeTlsPrivateKey) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTlsPrivateKey(value)
	for _, existing := range AllowedTypeTlsPrivateKeyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTlsPrivateKey", value)
}

// NewTypeTlsPrivateKeyFromValue returns a pointer to a valid TypeTlsPrivateKey
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTlsPrivateKeyFromValue(v string) (*TypeTlsPrivateKey, error) {
	ev := TypeTlsPrivateKey(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeTlsPrivateKey: valid values are %v", v, AllowedTypeTlsPrivateKeyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTlsPrivateKey) IsValid() bool {
	for _, existing := range AllowedTypeTlsPrivateKeyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to type_tls_private_key value
func (v TypeTlsPrivateKey) Ptr() *TypeTlsPrivateKey {
	return &v
}

// NullableTypeTlsPrivateKey is a helper abstraction for handling nullable typetlsprivatekey types.
type NullableTypeTlsPrivateKey struct {
	value *TypeTlsPrivateKey
	isSet bool
}

// Get returns the value.
func (v NullableTypeTlsPrivateKey) Get() *TypeTlsPrivateKey {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTlsPrivateKey) Set(val *TypeTlsPrivateKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTlsPrivateKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTlsPrivateKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTlsPrivateKey returns a pointer to a new instance of NullableTypeTlsPrivateKey.
func NewNullableTypeTlsPrivateKey(val *TypeTlsPrivateKey) *NullableTypeTlsPrivateKey {
	return &NullableTypeTlsPrivateKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTlsPrivateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeTlsPrivateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
