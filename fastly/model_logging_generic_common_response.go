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

// LoggingGenericCommonResponse struct for LoggingGenericCommonResponse
type LoggingGenericCommonResponse struct {
	// How the message should be formatted.
	MessageType *string `json:"message_type,omitempty"`
	// A timestamp format
	TimestampFormat NullableString `json:"timestamp_format,omitempty"`
	// The codec used for compressing your logs. Valid values are `zstd`, `snappy`, and `gzip`. Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	CompressionCodec *string `json:"compression_codec,omitempty"`
	// How frequently log files are finalized so they can be available for reading (in seconds).
	Period *string `json:"period,omitempty"`
	// The level of gzip encoding when sending logs (default `0`, no compression). Specifying both `compression_codec` and `gzip_level` in the same API request will result in an error.
	GzipLevel            *string `json:"gzip_level,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingGenericCommonResponse LoggingGenericCommonResponse

// NewLoggingGenericCommonResponse instantiates a new LoggingGenericCommonResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingGenericCommonResponse() *LoggingGenericCommonResponse {
	this := LoggingGenericCommonResponse{}
	var messageType string = "classic"
	this.MessageType = &messageType
	var period string = "3600"
	this.Period = &period
	var gzipLevel string = "0"
	this.GzipLevel = &gzipLevel
	return &this
}

// NewLoggingGenericCommonResponseWithDefaults instantiates a new LoggingGenericCommonResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingGenericCommonResponseWithDefaults() *LoggingGenericCommonResponse {
	this := LoggingGenericCommonResponse{}
	var messageType string = "classic"
	this.MessageType = &messageType
	var period string = "3600"
	this.Period = &period
	var gzipLevel string = "0"
	this.GzipLevel = &gzipLevel
	return &this
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LoggingGenericCommonResponse) GetMessageType() string {
	if o == nil || o.MessageType == nil {
		var ret string
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGenericCommonResponse) GetMessageTypeOk() (*string, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LoggingGenericCommonResponse) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given string and assigns it to the MessageType field.
func (o *LoggingGenericCommonResponse) SetMessageType(v string) {
	o.MessageType = &v
}

// GetTimestampFormat returns the TimestampFormat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingGenericCommonResponse) GetTimestampFormat() string {
	if o == nil || o.TimestampFormat.Get() == nil {
		var ret string
		return ret
	}
	return *o.TimestampFormat.Get()
}

// GetTimestampFormatOk returns a tuple with the TimestampFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingGenericCommonResponse) GetTimestampFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimestampFormat.Get(), o.TimestampFormat.IsSet()
}

// HasTimestampFormat returns a boolean if a field has been set.
func (o *LoggingGenericCommonResponse) HasTimestampFormat() bool {
	if o != nil && o.TimestampFormat.IsSet() {
		return true
	}

	return false
}

// SetTimestampFormat gets a reference to the given NullableString and assigns it to the TimestampFormat field.
func (o *LoggingGenericCommonResponse) SetTimestampFormat(v string) {
	o.TimestampFormat.Set(&v)
}

// SetTimestampFormatNil sets the value for TimestampFormat to be an explicit nil
func (o *LoggingGenericCommonResponse) SetTimestampFormatNil() {
	o.TimestampFormat.Set(nil)
}

// UnsetTimestampFormat ensures that no value is present for TimestampFormat, not even an explicit nil
func (o *LoggingGenericCommonResponse) UnsetTimestampFormat() {
	o.TimestampFormat.Unset()
}

// GetCompressionCodec returns the CompressionCodec field value if set, zero value otherwise.
func (o *LoggingGenericCommonResponse) GetCompressionCodec() string {
	if o == nil || o.CompressionCodec == nil {
		var ret string
		return ret
	}
	return *o.CompressionCodec
}

// GetCompressionCodecOk returns a tuple with the CompressionCodec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGenericCommonResponse) GetCompressionCodecOk() (*string, bool) {
	if o == nil || o.CompressionCodec == nil {
		return nil, false
	}
	return o.CompressionCodec, true
}

// HasCompressionCodec returns a boolean if a field has been set.
func (o *LoggingGenericCommonResponse) HasCompressionCodec() bool {
	if o != nil && o.CompressionCodec != nil {
		return true
	}

	return false
}

// SetCompressionCodec gets a reference to the given string and assigns it to the CompressionCodec field.
func (o *LoggingGenericCommonResponse) SetCompressionCodec(v string) {
	o.CompressionCodec = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *LoggingGenericCommonResponse) GetPeriod() string {
	if o == nil || o.Period == nil {
		var ret string
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGenericCommonResponse) GetPeriodOk() (*string, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *LoggingGenericCommonResponse) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given string and assigns it to the Period field.
func (o *LoggingGenericCommonResponse) SetPeriod(v string) {
	o.Period = &v
}

// GetGzipLevel returns the GzipLevel field value if set, zero value otherwise.
func (o *LoggingGenericCommonResponse) GetGzipLevel() string {
	if o == nil || o.GzipLevel == nil {
		var ret string
		return ret
	}
	return *o.GzipLevel
}

// GetGzipLevelOk returns a tuple with the GzipLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingGenericCommonResponse) GetGzipLevelOk() (*string, bool) {
	if o == nil || o.GzipLevel == nil {
		return nil, false
	}
	return o.GzipLevel, true
}

// HasGzipLevel returns a boolean if a field has been set.
func (o *LoggingGenericCommonResponse) HasGzipLevel() bool {
	if o != nil && o.GzipLevel != nil {
		return true
	}

	return false
}

// SetGzipLevel gets a reference to the given string and assigns it to the GzipLevel field.
func (o *LoggingGenericCommonResponse) SetGzipLevel(v string) {
	o.GzipLevel = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingGenericCommonResponse) MarshalJSON() ([]byte, error) {
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
func (o *LoggingGenericCommonResponse) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingGenericCommonResponse := _LoggingGenericCommonResponse{}

	if err = json.Unmarshal(bytes, &varLoggingGenericCommonResponse); err == nil {
		*o = LoggingGenericCommonResponse(varLoggingGenericCommonResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message_type")
		delete(additionalProperties, "timestamp_format")
		delete(additionalProperties, "compression_codec")
		delete(additionalProperties, "period")
		delete(additionalProperties, "gzip_level")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingGenericCommonResponse is a helper abstraction for handling nullable logginggenericcommonresponse types.
type NullableLoggingGenericCommonResponse struct {
	value *LoggingGenericCommonResponse
	isSet bool
}

// Get returns the value.
func (v NullableLoggingGenericCommonResponse) Get() *LoggingGenericCommonResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingGenericCommonResponse) Set(val *LoggingGenericCommonResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingGenericCommonResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingGenericCommonResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingGenericCommonResponse returns a pointer to a new instance of NullableLoggingGenericCommonResponse.
func NewNullableLoggingGenericCommonResponse(val *LoggingGenericCommonResponse) *NullableLoggingGenericCommonResponse {
	return &NullableLoggingGenericCommonResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingGenericCommonResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingGenericCommonResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
