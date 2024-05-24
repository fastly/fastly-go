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

// TypeTLSConfiguration Resource type
type TypeTLSConfiguration string

// List of resourceTypetls_configuration
const (
	TYPETLSCONFIGURATION_TLS_CONFIGURATION TypeTLSConfiguration = "tls_configuration"
)

// AllowedTypeTLSConfigurationEnumValues All allowed values of TypeTLSConfiguration enum
var AllowedTypeTLSConfigurationEnumValues = []TypeTLSConfiguration{
	"tls_configuration",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeTLSConfiguration) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTLSConfiguration(value)
	for _, existing := range AllowedTypeTLSConfigurationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTLSConfiguration", value)
}

// NewTypeTLSConfigurationFromValue returns a pointer to a valid TypeTLSConfiguration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTLSConfigurationFromValue(v string) (*TypeTLSConfiguration, error) {
	ev := TypeTLSConfiguration(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeTLSConfiguration: valid values are %v", v, AllowedTypeTLSConfigurationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTLSConfiguration) IsValid() bool {
	for _, existing := range AllowedTypeTLSConfigurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypetls_configuration value
func (v TypeTLSConfiguration) Ptr() *TypeTLSConfiguration {
	return &v
}

// NullableTypeTLSConfiguration is a helper abstraction for handling nullable typetlsconfiguration types. 
type NullableTypeTLSConfiguration struct {
	value *TypeTLSConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableTypeTLSConfiguration) Get() *TypeTLSConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTLSConfiguration) Set(val *TypeTLSConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTLSConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTLSConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTLSConfiguration returns a pointer to a new instance of NullableTypeTLSConfiguration.
func NewNullableTypeTLSConfiguration(val *TypeTLSConfiguration) *NullableTypeTLSConfiguration {
	return &NullableTypeTLSConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTLSConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeTLSConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
