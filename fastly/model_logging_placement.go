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

// LoggingPlacement Where in the generated VCL the logging call should be placed. If not set, endpoints with `format_version` of 2 are placed in `vcl_log` and those with `format_version` of 1 are placed in `vcl_deliver`. 
type LoggingPlacement string

// List of logging_placement
const (
	LOGGINGPLACEMENT_NONE LoggingPlacement = "none"
	LOGGINGPLACEMENT_WAF_DEBUG LoggingPlacement = "waf_debug"
	LOGGINGPLACEMENT_NULL LoggingPlacement = "null"
)

// AllowedLoggingPlacementEnumValues All allowed values of LoggingPlacement enum
var AllowedLoggingPlacementEnumValues = []LoggingPlacement{
	"none",
	"waf_debug",
	"null",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *LoggingPlacement) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoggingPlacement(value)
	for _, existing := range AllowedLoggingPlacementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoggingPlacement", value)
}

// NewLoggingPlacementFromValue returns a pointer to a valid LoggingPlacement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoggingPlacementFromValue(v string) (*LoggingPlacement, error) {
	ev := LoggingPlacement(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for LoggingPlacement: valid values are %v", v, AllowedLoggingPlacementEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoggingPlacement) IsValid() bool {
	for _, existing := range AllowedLoggingPlacementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to logging_placement value
func (v LoggingPlacement) Ptr() *LoggingPlacement {
	return &v
}

// NullableLoggingPlacement is a helper abstraction for handling nullable loggingplacement types. 
type NullableLoggingPlacement struct {
	value *LoggingPlacement
	isSet bool
}

// Get returns the value.
func (v NullableLoggingPlacement) Get() *LoggingPlacement {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingPlacement) Set(val *LoggingPlacement) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingPlacement) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingPlacement) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingPlacement returns a pointer to a new instance of NullableLoggingPlacement.
func NewNullableLoggingPlacement(val *LoggingPlacement) *NullableLoggingPlacement {
	return &NullableLoggingPlacement{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingPlacement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingPlacement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
