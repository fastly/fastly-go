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

// LoggingUseTLS Whether to use TLS.
type LoggingUseTLS int32

// List of logging_use_tls
const (
	LOGGINGUSETLS_no_tls LoggingUseTLS = 0
	LOGGINGUSETLS_use_tls LoggingUseTLS = 1
)

// AllowedLoggingUseTLSEnumValues All allowed values of LoggingUseTLS enum
var AllowedLoggingUseTLSEnumValues = []LoggingUseTLS{
	0,
	1,
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *LoggingUseTLS) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoggingUseTLS(value)
	for _, existing := range AllowedLoggingUseTLSEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoggingUseTLS", value)
}

// NewLoggingUseTLSFromValue returns a pointer to a valid LoggingUseTLS
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoggingUseTLSFromValue(v int32) (*LoggingUseTLS, error) {
	ev := LoggingUseTLS(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for LoggingUseTLS: valid values are %v", v, AllowedLoggingUseTLSEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoggingUseTLS) IsValid() bool {
	for _, existing := range AllowedLoggingUseTLSEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to logging_use_tls value
func (v LoggingUseTLS) Ptr() *LoggingUseTLS {
	return &v
}

// NullableLoggingUseTLS is a helper abstraction for handling nullable loggingusetls types. 
type NullableLoggingUseTLS struct {
	value *LoggingUseTLS
	isSet bool
}

// Get returns the value.
func (v NullableLoggingUseTLS) Get() *LoggingUseTLS {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingUseTLS) Set(val *LoggingUseTLS) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingUseTLS) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingUseTLS) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingUseTLS returns a pointer to a new instance of NullableLoggingUseTLS.
func NewNullableLoggingUseTLS(val *LoggingUseTLS) *NullableLoggingUseTLS {
	return &NullableLoggingUseTLS{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingUseTLS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingUseTLS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
