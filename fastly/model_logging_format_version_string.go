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
)

// LoggingFormatVersionString struct for LoggingFormatVersionString
type LoggingFormatVersionString struct {
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`. 
	FormatVersion *string `json:"format_version,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingFormatVersionString LoggingFormatVersionString

// NewLoggingFormatVersionString instantiates a new LoggingFormatVersionString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingFormatVersionString() *LoggingFormatVersionString {
	this := LoggingFormatVersionString{}
	var formatVersion string = "2"
	this.FormatVersion = &formatVersion
	return &this
}

// NewLoggingFormatVersionStringWithDefaults instantiates a new LoggingFormatVersionString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingFormatVersionStringWithDefaults() *LoggingFormatVersionString {
	this := LoggingFormatVersionString{}
	var formatVersion string = "2"
	this.FormatVersion = &formatVersion
	return &this
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingFormatVersionString) GetFormatVersion() string {
	if o == nil || o.FormatVersion == nil {
		var ret string
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingFormatVersionString) GetFormatVersionOk() (*string, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingFormatVersionString) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given string and assigns it to the FormatVersion field.
func (o *LoggingFormatVersionString) SetFormatVersion(v string) {
	o.FormatVersion = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingFormatVersionString) MarshalJSON() ([]byte, error) {
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
func (o *LoggingFormatVersionString) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingFormatVersionString := _LoggingFormatVersionString{}

	if err = json.Unmarshal(bytes, &varLoggingFormatVersionString); err == nil {
		*o = LoggingFormatVersionString(varLoggingFormatVersionString)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "format_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingFormatVersionString is a helper abstraction for handling nullable loggingformatversionstring types. 
type NullableLoggingFormatVersionString struct {
	value *LoggingFormatVersionString
	isSet bool
}

// Get returns the value.
func (v NullableLoggingFormatVersionString) Get() *LoggingFormatVersionString {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingFormatVersionString) Set(val *LoggingFormatVersionString) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingFormatVersionString) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingFormatVersionString) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingFormatVersionString returns a pointer to a new instance of NullableLoggingFormatVersionString.
func NewNullableLoggingFormatVersionString(val *LoggingFormatVersionString) *NullableLoggingFormatVersionString {
	return &NullableLoggingFormatVersionString{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingFormatVersionString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingFormatVersionString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
