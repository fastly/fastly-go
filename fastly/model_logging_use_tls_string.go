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

// LoggingUseTLSString Whether to use TLS.
type LoggingUseTLSString string

// List of logging_use_tls_string
const (
	LOGGINGUSETLSSTRING_no_tls LoggingUseTLSString = "0"
	LOGGINGUSETLSSTRING_use_tls LoggingUseTLSString = "1"
)

// AllowedLoggingUseTLSStringEnumValues All allowed values of LoggingUseTLSString enum
var AllowedLoggingUseTLSStringEnumValues = []LoggingUseTLSString{
	"0",
	"1",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *LoggingUseTLSString) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoggingUseTLSString(value)
	for _, existing := range AllowedLoggingUseTLSStringEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoggingUseTLSString", value)
}

// NewLoggingUseTLSStringFromValue returns a pointer to a valid LoggingUseTLSString
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoggingUseTLSStringFromValue(v string) (*LoggingUseTLSString, error) {
	ev := LoggingUseTLSString(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for LoggingUseTLSString: valid values are %v", v, AllowedLoggingUseTLSStringEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoggingUseTLSString) IsValid() bool {
	for _, existing := range AllowedLoggingUseTLSStringEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to logging_use_tls_string value
func (v LoggingUseTLSString) Ptr() *LoggingUseTLSString {
	return &v
}

// NullableLoggingUseTLSString is a helper abstraction for handling nullable loggingusetlsstring types. 
type NullableLoggingUseTLSString struct {
	value *LoggingUseTLSString
	isSet bool
}

// Get returns the value.
func (v NullableLoggingUseTLSString) Get() *LoggingUseTLSString {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingUseTLSString) Set(val *LoggingUseTLSString) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingUseTLSString) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingUseTLSString) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingUseTLSString returns a pointer to a new instance of NullableLoggingUseTLSString.
func NewNullableLoggingUseTLSString(val *LoggingUseTLSString) *NullableLoggingUseTLSString {
	return &NullableLoggingUseTLSString{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingUseTLSString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingUseTLSString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
