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

// LoggingMessageType How the message should be formatted.
type LoggingMessageType string

// List of logging_message_type
const (
	LOGGINGMESSAGETYPE_CLASSIC LoggingMessageType = "classic"
	LOGGINGMESSAGETYPE_LOGGLY  LoggingMessageType = "loggly"
	LOGGINGMESSAGETYPE_LOGPLEX LoggingMessageType = "logplex"
	LOGGINGMESSAGETYPE_BLANK   LoggingMessageType = "blank"
)

// AllowedLoggingMessageTypeEnumValues All allowed values of LoggingMessageType enum
var AllowedLoggingMessageTypeEnumValues = []LoggingMessageType{
	"classic",
	"loggly",
	"logplex",
	"blank",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *LoggingMessageType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LoggingMessageType(value)
	for _, existing := range AllowedLoggingMessageTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LoggingMessageType", value)
}

// NewLoggingMessageTypeFromValue returns a pointer to a valid LoggingMessageType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLoggingMessageTypeFromValue(v string) (*LoggingMessageType, error) {
	ev := LoggingMessageType(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for LoggingMessageType: valid values are %v", v, AllowedLoggingMessageTypeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LoggingMessageType) IsValid() bool {
	for _, existing := range AllowedLoggingMessageTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to logging_message_type value
func (v LoggingMessageType) Ptr() *LoggingMessageType {
	return &v
}

// NullableLoggingMessageType is a helper abstraction for handling nullable loggingmessagetype types.
type NullableLoggingMessageType struct {
	value *LoggingMessageType
	isSet bool
}

// Get returns the value.
func (v NullableLoggingMessageType) Get() *LoggingMessageType {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingMessageType) Set(val *LoggingMessageType) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingMessageType) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingMessageType) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingMessageType returns a pointer to a new instance of NullableLoggingMessageType.
func NewNullableLoggingMessageType(val *LoggingMessageType) *NullableLoggingMessageType {
	return &NullableLoggingMessageType{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingMessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingMessageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
