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

// LoggingGenericCommonResponseAllOf1 struct for LoggingGenericCommonResponseAllOf1
type LoggingGenericCommonResponseAllOf1 struct {
	// How frequently log files are finalized so they can be available for reading (in seconds).
	Period *string `json:"period,omitempty"`
	// The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	GzipLevel            *string `json:"gzip_level,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingGenericCommonResponseAllOf1 LoggingGenericCommonResponseAllOf1

// NewLoggingGenericCommonResponseAllOf1 instantiates a new LoggingGenericCommonResponseAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingGenericCommonResponseAllOf1() *LoggingGenericCommonResponseAllOf1 {
	this := LoggingGenericCommonResponseAllOf1{}
	var period string = "3600"
	this.Period = &period
	var gzipLevel string = "0"
	this.GzipLevel = &gzipLevel
	return &this
}

// NewLoggingGenericCommonResponseAllOf1WithDefaults instantiates a new LoggingGenericCommonResponseAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingGenericCommonResponseAllOf1WithDefaults() *LoggingGenericCommonResponseAllOf1 {
	this := LoggingGenericCommonResponseAllOf1{}
	var period string = "3600"
	this.Period = &period
	var gzipLevel string = "0"
	this.GzipLevel = &gzipLevel
	return &this
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *LoggingGenericCommonResponseAllOf1) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGenericCommonResponseAllOf1) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *LoggingGenericCommonResponseAllOf1) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *LoggingGenericCommonResponseAllOf1) SetPeriod(v string) {
	o.Period = &v
}

// GetGzipLevel returns the GzipLevel field value if set, zero value otherwise.
func (o *LoggingGenericCommonResponseAllOf1) GetGzipLevel() string {
	if o == nil || o.GzipLevel == nil {
		var ret string
		return ret
	}
	return *o.GzipLevel
}

// GetGzipLevelOk returns a tuple with the GzipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGenericCommonResponseAllOf1) GetGzipLevelOk() (*string, bool) {
	if o == nil || o.GzipLevel == nil {
		return nil, false
	}
	return o.GzipLevel, true
}

// HasGzipLevel returns a boolean if a field has been set.
func (o *LoggingGenericCommonResponseAllOf1) HasGzipLevel() bool {
	if o != nil && o.GzipLevel != nil {
		return true
	}

	return false
}

// SetGzipLevel gets a reference to the given string and assigns it to the GzipLevel field.
func (o *LoggingGenericCommonResponseAllOf1) SetGzipLevel(v string) {
	o.GzipLevel = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingGenericCommonResponseAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.GzipLevel != nil {
		toSerialize["gzip_level"] = o.GzipLevel
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingGenericCommonResponseAllOf1) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingGenericCommonResponseAllOf1 := _LoggingGenericCommonResponseAllOf1{}

	if err = json.Unmarshal(bytes, &varLoggingGenericCommonResponseAllOf1); err == nil {
		*o = LoggingGenericCommonResponseAllOf1(varLoggingGenericCommonResponseAllOf1)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "period")
		delete(additionalProperties, "gzip_level")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingGenericCommonResponseAllOf1 is a helper abstraction for handling nullable logginggenericcommonresponseallof1 types.
type NullableLoggingGenericCommonResponseAllOf1 struct {
	value *LoggingGenericCommonResponseAllOf1
	isSet bool
}

// Get returns the value.
func (v NullableLoggingGenericCommonResponseAllOf1) Get() *LoggingGenericCommonResponseAllOf1 {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingGenericCommonResponseAllOf1) Set(val *LoggingGenericCommonResponseAllOf1) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingGenericCommonResponseAllOf1) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingGenericCommonResponseAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingGenericCommonResponseAllOf1 returns a pointer to a new instance of NullableLoggingGenericCommonResponseAllOf1.
func NewNullableLoggingGenericCommonResponseAllOf1(val *LoggingGenericCommonResponseAllOf1) *NullableLoggingGenericCommonResponseAllOf1 {
	return &NullableLoggingGenericCommonResponseAllOf1{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingGenericCommonResponseAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingGenericCommonResponseAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
