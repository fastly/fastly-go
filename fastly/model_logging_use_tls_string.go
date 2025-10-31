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

// LoggingUseTlsString Whether to use TLS.
type LoggingUseTlsString string

// List of logging_use_tls_string
const (
	LOGGINGUSETLSSTRING_no_tls  LoggingUseTlsString = "0"
	LOGGINGUSETLSSTRING_use_tls LoggingUseTlsString = "1"
)

// AllowedLoggingUseTlsStringEnumValues All allowed values of LoggingUseTlsString enum
var AllowedLoggingUseTlsStringEnumValues = []LoggingUseTlsString{
	"0",
	"1",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *LoggingUseTlsString) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoggingUseTlsString(value)
	for _, existing := range AllowedLoggingUseTlsStringEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoggingUseTlsString", value)
}

// NewLoggingUseTlsStringFromValue returns a pointer to a valid LoggingUseTlsString
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoggingUseTlsStringFromValue(v string) (*LoggingUseTlsString, error) {
	ev := LoggingUseTlsString(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LoggingUseTlsString: valid values are %v", v, AllowedLoggingUseTlsStringEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoggingUseTlsString) IsValid() bool {
	for _, existing := range AllowedLoggingUseTlsStringEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to logging_use_tls_string value
func (v LoggingUseTlsString) Ptr() *LoggingUseTlsString {
	return &v
}

// NullableLoggingUseTlsString is a helper abstraction for handling nullable loggingusetlsstring types.
type NullableLoggingUseTlsString struct {
	value *LoggingUseTlsString
	isSet bool
}

// Get returns the value.
func (v NullableLoggingUseTlsString) Get() *LoggingUseTlsString {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingUseTlsString) Set(val *LoggingUseTlsString) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingUseTlsString) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingUseTlsString) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingUseTlsString returns a pointer to a new instance of NullableLoggingUseTlsString.
func NewNullableLoggingUseTlsString(val *LoggingUseTlsString) *NullableLoggingUseTlsString {
	return &NullableLoggingUseTlsString{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingUseTlsString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingUseTlsString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
