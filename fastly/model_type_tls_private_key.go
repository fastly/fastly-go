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

// TypeTLSPrivateKey Resource type
type TypeTLSPrivateKey string

// List of resourceTypetls_private_key
const (
	TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY TypeTLSPrivateKey = "tls_private_key"
)

// AllowedTypeTLSPrivateKeyEnumValues All allowed values of TypeTLSPrivateKey enum
var AllowedTypeTLSPrivateKeyEnumValues = []TypeTLSPrivateKey{
	"tls_private_key",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeTLSPrivateKey) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTLSPrivateKey(value)
	for _, existing := range AllowedTypeTLSPrivateKeyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTLSPrivateKey", value)
}

// NewTypeTLSPrivateKeyFromValue returns a pointer to a valid TypeTLSPrivateKey
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTLSPrivateKeyFromValue(v string) (*TypeTLSPrivateKey, error) {
	ev := TypeTLSPrivateKey(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeTLSPrivateKey: valid values are %v", v, AllowedTypeTLSPrivateKeyEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTLSPrivateKey) IsValid() bool {
	for _, existing := range AllowedTypeTLSPrivateKeyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypetls_private_key value
func (v TypeTLSPrivateKey) Ptr() *TypeTLSPrivateKey {
	return &v
}

// NullableTypeTLSPrivateKey is a helper abstraction for handling nullable typetlsprivatekey types. 
type NullableTypeTLSPrivateKey struct {
	value *TypeTLSPrivateKey
	isSet bool
}

// Get returns the value.
func (v NullableTypeTLSPrivateKey) Get() *TypeTLSPrivateKey {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTLSPrivateKey) Set(val *TypeTLSPrivateKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTLSPrivateKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTLSPrivateKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTLSPrivateKey returns a pointer to a new instance of NullableTypeTLSPrivateKey.
func NewNullableTypeTLSPrivateKey(val *TypeTLSPrivateKey) *NullableTypeTLSPrivateKey {
	return &NullableTypeTLSPrivateKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTLSPrivateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeTLSPrivateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
