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
)

// LoggingFtpResponseAllOf struct for LoggingFtpResponseAllOf
type LoggingFtpResponseAllOf struct {
	// The port number.
	Port                 *string `json:"port,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingFtpResponseAllOf LoggingFtpResponseAllOf

// NewLoggingFtpResponseAllOf instantiates a new LoggingFtpResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingFtpResponseAllOf() *LoggingFtpResponseAllOf {
	this := LoggingFtpResponseAllOf{}
	var port string = "21"
	this.Port = &port
	return &this
}

// NewLoggingFtpResponseAllOfWithDefaults instantiates a new LoggingFtpResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingFtpResponseAllOfWithDefaults() *LoggingFtpResponseAllOf {
	this := LoggingFtpResponseAllOf{}
	var port string = "21"
	this.Port = &port
	return &this
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *LoggingFtpResponseAllOf) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingFtpResponseAllOf) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *LoggingFtpResponseAllOf) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *LoggingFtpResponseAllOf) SetPort(v string) {
	o.Port = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingFtpResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingFtpResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingFtpResponseAllOf := _LoggingFtpResponseAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingFtpResponseAllOf); err == nil {
		*o = LoggingFtpResponseAllOf(varLoggingFtpResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "port")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingFtpResponseAllOf is a helper abstraction for handling nullable loggingftpresponseallof types.
type NullableLoggingFtpResponseAllOf struct {
	value *LoggingFtpResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingFtpResponseAllOf) Get() *LoggingFtpResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingFtpResponseAllOf) Set(val *LoggingFtpResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingFtpResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingFtpResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingFtpResponseAllOf returns a pointer to a new instance of NullableLoggingFtpResponseAllOf.
func NewNullableLoggingFtpResponseAllOf(val *LoggingFtpResponseAllOf) *NullableLoggingFtpResponseAllOf {
	return &NullableLoggingFtpResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingFtpResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingFtpResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
