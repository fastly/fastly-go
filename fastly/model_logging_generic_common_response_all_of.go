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

// LoggingGenericCommonResponseAllOf struct for LoggingGenericCommonResponseAllOf
type LoggingGenericCommonResponseAllOf struct {
	// How the message should be formatted.
	MessageType *string `json:"message_type,omitempty"`
	// A timestamp format
	TimestampFormat NullableString `json:"timestamp_format,omitempty"`
	// The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	CompressionCodec *string `json:"compression_codec,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingGenericCommonResponseAllOf LoggingGenericCommonResponseAllOf

// NewLoggingGenericCommonResponseAllOf instantiates a new LoggingGenericCommonResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingGenericCommonResponseAllOf() *LoggingGenericCommonResponseAllOf {
	this := LoggingGenericCommonResponseAllOf{}
	var messageType string = "classic"
	this.MessageType = &messageType
	return &this
}

// NewLoggingGenericCommonResponseAllOfWithDefaults instantiates a new LoggingGenericCommonResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingGenericCommonResponseAllOfWithDefaults() *LoggingGenericCommonResponseAllOf {
	this := LoggingGenericCommonResponseAllOf{}
	var messageType string = "classic"
	this.MessageType = &messageType
	return &this
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LoggingGenericCommonResponseAllOf) GetMessageType() string {
	if o == nil || o.MessageType == nil {
		var ret string
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGenericCommonResponseAllOf) GetMessageTypeOk() (*string, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LoggingGenericCommonResponseAllOf) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given string and assigns it to the MessageType field.
func (o *LoggingGenericCommonResponseAllOf) SetMessageType(v string) {
	o.MessageType = &v
}

// GetTimestampFormat returns the TimestampFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingGenericCommonResponseAllOf) GetTimestampFormat() string {
	if o == nil || o.TimestampFormat.Get() == nil {
		var ret string
		return ret
	}
	return *o.TimestampFormat.Get()
}

// GetTimestampFormatOk returns a tuple with the TimestampFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingGenericCommonResponseAllOf) GetTimestampFormatOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TimestampFormat.Get(), o.TimestampFormat.IsSet()
}

// HasTimestampFormat returns a boolean if a field has been set.
func (o *LoggingGenericCommonResponseAllOf) HasTimestampFormat() bool {
	if o != nil && o.TimestampFormat.IsSet() {
		return true
	}

	return false
}

// SetTimestampFormat gets a reference to the given NullableString and assigns it to the TimestampFormat field.
func (o *LoggingGenericCommonResponseAllOf) SetTimestampFormat(v string) {
	o.TimestampFormat.Set(&v)
}
// SetTimestampFormatNil sets the value for TimestampFormat to be an explicit nil
func (o *LoggingGenericCommonResponseAllOf) SetTimestampFormatNil() {
	o.TimestampFormat.Set(nil)
}

// UnsetTimestampFormat ensures that no value is present for TimestampFormat, not even an explicit nil
func (o *LoggingGenericCommonResponseAllOf) UnsetTimestampFormat() {
	o.TimestampFormat.Unset()
}

// GetCompressionCodec returns the CompressionCodec field value if set, zero value otherwise.
func (o *LoggingGenericCommonResponseAllOf) GetCompressionCodec() string {
	if o == nil || o.CompressionCodec == nil {
		var ret string
		return ret
	}
	return *o.CompressionCodec
}

// GetCompressionCodecOk returns a tuple with the CompressionCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGenericCommonResponseAllOf) GetCompressionCodecOk() (*string, bool) {
	if o == nil || o.CompressionCodec == nil {
		return nil, false
	}
	return o.CompressionCodec, true
}

// HasCompressionCodec returns a boolean if a field has been set.
func (o *LoggingGenericCommonResponseAllOf) HasCompressionCodec() bool {
	if o != nil && o.CompressionCodec != nil {
		return true
	}

	return false
}

// SetCompressionCodec gets a reference to the given string and assigns it to the CompressionCodec field.
func (o *LoggingGenericCommonResponseAllOf) SetCompressionCodec(v string) {
	o.CompressionCodec = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingGenericCommonResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.MessageType != nil {
		toSerialize["message_type"] = o.MessageType
	}
	if o.TimestampFormat.IsSet() {
		toSerialize["timestamp_format"] = o.TimestampFormat.Get()
	}
	if o.CompressionCodec != nil {
		toSerialize["compression_codec"] = o.CompressionCodec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingGenericCommonResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingGenericCommonResponseAllOf := _LoggingGenericCommonResponseAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingGenericCommonResponseAllOf); err == nil {
		*o = LoggingGenericCommonResponseAllOf(varLoggingGenericCommonResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message_type")
		delete(additionalProperties, "timestamp_format")
		delete(additionalProperties, "compression_codec")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingGenericCommonResponseAllOf is a helper abstraction for handling nullable logginggenericcommonresponseallof types. 
type NullableLoggingGenericCommonResponseAllOf struct {
	value *LoggingGenericCommonResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingGenericCommonResponseAllOf) Get() *LoggingGenericCommonResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingGenericCommonResponseAllOf) Set(val *LoggingGenericCommonResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingGenericCommonResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingGenericCommonResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingGenericCommonResponseAllOf returns a pointer to a new instance of NullableLoggingGenericCommonResponseAllOf.
func NewNullableLoggingGenericCommonResponseAllOf(val *LoggingGenericCommonResponseAllOf) *NullableLoggingGenericCommonResponseAllOf {
	return &NullableLoggingGenericCommonResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingGenericCommonResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingGenericCommonResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
