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

// TypeTlsConfiguration Resource type
type TypeTlsConfiguration string

// List of type_tls_configuration
const (
	TYPETLSCONFIGURATION_TLS_CONFIGURATION TypeTlsConfiguration = "tls_configuration"
)

// AllowedTypeTlsConfigurationEnumValues All allowed values of TypeTlsConfiguration enum
var AllowedTypeTlsConfigurationEnumValues = []TypeTlsConfiguration{
	"tls_configuration",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeTlsConfiguration) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTlsConfiguration(value)
	for _, existing := range AllowedTypeTlsConfigurationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTlsConfiguration", value)
}

// NewTypeTlsConfigurationFromValue returns a pointer to a valid TypeTlsConfiguration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTlsConfigurationFromValue(v string) (*TypeTlsConfiguration, error) {
	ev := TypeTlsConfiguration(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeTlsConfiguration: valid values are %v", v, AllowedTypeTlsConfigurationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTlsConfiguration) IsValid() bool {
	for _, existing := range AllowedTypeTlsConfigurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to type_tls_configuration value
func (v TypeTlsConfiguration) Ptr() *TypeTlsConfiguration {
	return &v
}

// NullableTypeTlsConfiguration is a helper abstraction for handling nullable typetlsconfiguration types.
type NullableTypeTlsConfiguration struct {
	value *TypeTlsConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableTypeTlsConfiguration) Get() *TypeTlsConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTlsConfiguration) Set(val *TypeTlsConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTlsConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTlsConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTlsConfiguration returns a pointer to a new instance of NullableTypeTlsConfiguration.
func NewNullableTypeTlsConfiguration(val *TypeTlsConfiguration) *NullableTypeTlsConfiguration {
	return &NullableTypeTlsConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTlsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeTlsConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
