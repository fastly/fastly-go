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

// LoggingFormatVersionInteger struct for LoggingFormatVersionInteger
type LoggingFormatVersionInteger struct {
	// The version of the custom logging format used for the configured endpoint. The logging call gets placed by default in `vcl_log` if `format_version` is set to `2` and in `vcl_deliver` if `format_version` is set to `1`. 
	FormatVersion *int32 `json:"format_version,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingFormatVersionInteger LoggingFormatVersionInteger

// NewLoggingFormatVersionInteger instantiates a new LoggingFormatVersionInteger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingFormatVersionInteger() *LoggingFormatVersionInteger {
	this := LoggingFormatVersionInteger{}
	var formatVersion int32 = 2
	this.FormatVersion = &formatVersion
	return &this
}

// NewLoggingFormatVersionIntegerWithDefaults instantiates a new LoggingFormatVersionInteger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingFormatVersionIntegerWithDefaults() *LoggingFormatVersionInteger {
	this := LoggingFormatVersionInteger{}
	var formatVersion int32 = 2
	this.FormatVersion = &formatVersion
	return &this
}

// GetFormatVersion returns the FormatVersion field value if set, zero value otherwise.
func (o *LoggingFormatVersionInteger) GetFormatVersion() int32 {
	if o == nil || o.FormatVersion == nil {
		var ret int32
		return ret
	}
	return *o.FormatVersion
}

// GetFormatVersionOk returns a tuple with the FormatVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingFormatVersionInteger) GetFormatVersionOk() (*int32, bool) {
	if o == nil || o.FormatVersion == nil {
		return nil, false
	}
	return o.FormatVersion, true
}

// HasFormatVersion returns a boolean if a field has been set.
func (o *LoggingFormatVersionInteger) HasFormatVersion() bool {
	if o != nil && o.FormatVersion != nil {
		return true
	}

	return false
}

// SetFormatVersion gets a reference to the given int32 and assigns it to the FormatVersion field.
func (o *LoggingFormatVersionInteger) SetFormatVersion(v int32) {
	o.FormatVersion = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingFormatVersionInteger) MarshalJSON() ([]byte, error) {
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
func (o *LoggingFormatVersionInteger) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingFormatVersionInteger := _LoggingFormatVersionInteger{}

	if err = json.Unmarshal(bytes, &varLoggingFormatVersionInteger); err == nil {
		*o = LoggingFormatVersionInteger(varLoggingFormatVersionInteger)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "format_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingFormatVersionInteger is a helper abstraction for handling nullable loggingformatversioninteger types. 
type NullableLoggingFormatVersionInteger struct {
	value *LoggingFormatVersionInteger
	isSet bool
}

// Get returns the value.
func (v NullableLoggingFormatVersionInteger) Get() *LoggingFormatVersionInteger {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingFormatVersionInteger) Set(val *LoggingFormatVersionInteger) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingFormatVersionInteger) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingFormatVersionInteger) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingFormatVersionInteger returns a pointer to a new instance of NullableLoggingFormatVersionInteger.
func NewNullableLoggingFormatVersionInteger(val *LoggingFormatVersionInteger) *NullableLoggingFormatVersionInteger {
	return &NullableLoggingFormatVersionInteger{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingFormatVersionInteger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingFormatVersionInteger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
