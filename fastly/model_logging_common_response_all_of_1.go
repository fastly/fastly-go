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

// LoggingCommonResponseAllOf1 struct for LoggingCommonResponseAllOf1
type LoggingCommonResponseAllOf1 struct {
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`.
	FormatVersion        *string `json:"format_version,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingCommonResponseAllOf1 LoggingCommonResponseAllOf1

// NewLoggingCommonResponseAllOf1 instantiates a new LoggingCommonResponseAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingCommonResponseAllOf1() *LoggingCommonResponseAllOf1 {
	this := LoggingCommonResponseAllOf1{}
	var formatVersion string = "2"
	this.FormatVersion = &formatVersion
	return &this
}

// NewLoggingCommonResponseAllOf1WithDefaults instantiates a new LoggingCommonResponseAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingCommonResponseAllOf1WithDefaults() *LoggingCommonResponseAllOf1 {
	this := LoggingCommonResponseAllOf1{}
	var formatVersion string = "2"
	this.FormatVersion = &formatVersion
	return &this
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingCommonResponseAllOf1) GetFormatVersion() string {
	if o == nil || o.FormatVersion == nil {
		var ret string
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingCommonResponseAllOf1) GetFormatVersionOk() (*string, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingCommonResponseAllOf1) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given string and assigns it to the FormatVersion field.
func (o *LoggingCommonResponseAllOf1) SetFormatVersion(v string) {
	o.FormatVersion = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingCommonResponseAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.FormatVersion != nil {
		toSerialize["format_version"] = o.FormatVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingCommonResponseAllOf1) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingCommonResponseAllOf1 := _LoggingCommonResponseAllOf1{}

	if err = json.Unmarshal(bytes, &varLoggingCommonResponseAllOf1); err == nil {
		*o = LoggingCommonResponseAllOf1(varLoggingCommonResponseAllOf1)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "format_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingCommonResponseAllOf1 is a helper abstraction for handling nullable loggingcommonresponseallof1 types.
type NullableLoggingCommonResponseAllOf1 struct {
	value *LoggingCommonResponseAllOf1
	isSet bool
}

// Get returns the value.
func (v NullableLoggingCommonResponseAllOf1) Get() *LoggingCommonResponseAllOf1 {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingCommonResponseAllOf1) Set(val *LoggingCommonResponseAllOf1) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingCommonResponseAllOf1) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingCommonResponseAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingCommonResponseAllOf1 returns a pointer to a new instance of NullableLoggingCommonResponseAllOf1.
func NewNullableLoggingCommonResponseAllOf1(val *LoggingCommonResponseAllOf1) *NullableLoggingCommonResponseAllOf1 {
	return &NullableLoggingCommonResponseAllOf1{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingCommonResponseAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingCommonResponseAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
