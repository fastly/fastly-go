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

// LoggingFormatVersion The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`. 
type LoggingFormatVersion int32

// List of logging_format_version
const (
	LOGGINGFORMATVERSION_v1 LoggingFormatVersion = 1
	LOGGINGFORMATVERSION_v2 LoggingFormatVersion = 2
)

// AllowedLoggingFormatVersionEnumValues All allowed values of LoggingFormatVersion enum
var AllowedLoggingFormatVersionEnumValues = []LoggingFormatVersion{
	1,
	2,
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *LoggingFormatVersion) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoggingFormatVersion(value)
	for _, existing := range AllowedLoggingFormatVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoggingFormatVersion", value)
}

// NewLoggingFormatVersionFromValue returns a pointer to a valid LoggingFormatVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoggingFormatVersionFromValue(v int32) (*LoggingFormatVersion, error) {
	ev := LoggingFormatVersion(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for LoggingFormatVersion: valid values are %v", v, AllowedLoggingFormatVersionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoggingFormatVersion) IsValid() bool {
	for _, existing := range AllowedLoggingFormatVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to logging_format_version value
func (v LoggingFormatVersion) Ptr() *LoggingFormatVersion {
	return &v
}

// NullableLoggingFormatVersion is a helper abstraction for handling nullable loggingformatversion types. 
type NullableLoggingFormatVersion struct {
	value *LoggingFormatVersion
	isSet bool
}

// Get returns the value.
func (v NullableLoggingFormatVersion) Get() *LoggingFormatVersion {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingFormatVersion) Set(val *LoggingFormatVersion) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingFormatVersion) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingFormatVersion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingFormatVersion returns a pointer to a new instance of NullableLoggingFormatVersion.
func NewNullableLoggingFormatVersion(val *LoggingFormatVersion) *NullableLoggingFormatVersion {
	return &NullableLoggingFormatVersion{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingFormatVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingFormatVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
