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

// LoggingHttpsAdditional struct for LoggingHttpsAdditional
type LoggingHttpsAdditional struct {
	// The URL to send logs to. Must use HTTPS. Required.
	Url *string `json:"url,omitempty"`
	// The maximum number of logs sent in one request. Defaults `0` (10k).
	RequestMaxEntries *int32 `json:"request_max_entries,omitempty"`
	// The maximum number of bytes sent in one request. Defaults `0` (100MB).
	RequestMaxBytes *int32 `json:"request_max_bytes,omitempty"`
	// Content type of the header sent with the request.
	ContentType NullableString `json:"content_type,omitempty"`
	// Name of the custom header sent with the request.
	HeaderName  NullableString      `json:"header_name,omitempty"`
	MessageType *LoggingMessageType `json:"message_type,omitempty"`
	// Value of the custom header sent with the request.
	HeaderValue NullableString `json:"header_value,omitempty"`
	// HTTP method used for request.
	Method *string `json:"method,omitempty"`
	// Enforces valid JSON formatting for log entries.
	JsonFormat *string `json:"json_format,omitempty"`
	// A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
	Format *string `json:"format,omitempty"`
	// How frequently, in seconds, batches of log data are sent to the HTTPS endpoint. A value of `0` sends logs at the same interval as the default, which is `5` seconds.
	Period               *int32 `json:"period,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingHttpsAdditional LoggingHttpsAdditional

// NewLoggingHttpsAdditional instantiates a new LoggingHttpsAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingHttpsAdditional() *LoggingHttpsAdditional {
	this := LoggingHttpsAdditional{}
	var requestMaxEntries int32 = 0
	this.RequestMaxEntries = &requestMaxEntries
	var requestMaxBytes int32 = 0
	this.RequestMaxBytes = &requestMaxBytes
	var contentType string = "null"
	this.ContentType = *NewNullableString(&contentType)
	var headerName string = "null"
	this.HeaderName = *NewNullableString(&headerName)
	var messageType LoggingMessageType = LOGGINGMESSAGETYPE_CLASSIC
	this.MessageType = &messageType
	var headerValue string = "null"
	this.HeaderValue = *NewNullableString(&headerValue)
	var method string = "POST"
	this.Method = &method
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	var period int32 = 5
	this.Period = &period
	return &this
}

// NewLoggingHttpsAdditionalWithDefaults instantiates a new LoggingHttpsAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingHttpsAdditionalWithDefaults() *LoggingHttpsAdditional {
	this := LoggingHttpsAdditional{}
	var requestMaxEntries int32 = 0
	this.RequestMaxEntries = &requestMaxEntries
	var requestMaxBytes int32 = 0
	this.RequestMaxBytes = &requestMaxBytes
	var contentType string = "null"
	this.ContentType = *NewNullableString(&contentType)
	var headerName string = "null"
	this.HeaderName = *NewNullableString(&headerName)
	var messageType LoggingMessageType = LOGGINGMESSAGETYPE_CLASSIC
	this.MessageType = &messageType
	var headerValue string = "null"
	this.HeaderValue = *NewNullableString(&headerValue)
	var method string = "POST"
	this.Method = &method
	var format string = "%h %l %u %t \"%r\" %&gt;s %b"
	this.Format = &format
	var period int32 = 5
	this.Period = &period
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *LoggingHttpsAdditional) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsAdditional) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *LoggingHttpsAdditional) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *LoggingHttpsAdditional) SetUrl(v string) {
	o.Url = &v
}

// GetRequestMaxEntries returns the RequestMaxEntries field value if set, zero value otherwise.
func (o *LoggingHttpsAdditional) GetRequestMaxEntries() int32 {
	if o == nil || o.RequestMaxEntries == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxEntries
}

// GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsAdditional) GetRequestMaxEntriesOk() (*int32, bool) {
	if o == nil || o.RequestMaxEntries == nil {
		return nil, false
	}
	return o.RequestMaxEntries, true
}

// HasRequestMaxEntries returns a boolean if a field has been set.
func (o *LoggingHttpsAdditional) HasRequestMaxEntries() bool {
	if o != nil && o.RequestMaxEntries != nil {
		return true
	}

	return false
}

// SetRequestMaxEntries gets a reference to the given int32 and assigns it to the RequestMaxEntries field.
func (o *LoggingHttpsAdditional) SetRequestMaxEntries(v int32) {
	o.RequestMaxEntries = &v
}

// GetRequestMaxBytes returns the RequestMaxBytes field value if set, zero value otherwise.
func (o *LoggingHttpsAdditional) GetRequestMaxBytes() int32 {
	if o == nil || o.RequestMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxBytes
}

// GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsAdditional) GetRequestMaxBytesOk() (*int32, bool) {
	if o == nil || o.RequestMaxBytes == nil {
		return nil, false
	}
	return o.RequestMaxBytes, true
}

// HasRequestMaxBytes returns a boolean if a field has been set.
func (o *LoggingHttpsAdditional) HasRequestMaxBytes() bool {
	if o != nil && o.RequestMaxBytes != nil {
		return true
	}

	return false
}

// SetRequestMaxBytes gets a reference to the given int32 and assigns it to the RequestMaxBytes field.
func (o *LoggingHttpsAdditional) SetRequestMaxBytes(v int32) {
	o.RequestMaxBytes = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsAdditional) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsAdditional) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *LoggingHttpsAdditional) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *LoggingHttpsAdditional) SetContentType(v string) {
	o.ContentType.Set(&v)
}

// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *LoggingHttpsAdditional) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *LoggingHttpsAdditional) UnsetContentType() {
	o.ContentType.Unset()
}

// GetHeaderName returns the HeaderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsAdditional) GetHeaderName() string {
	if o == nil || o.HeaderName.Get() == nil {
		var ret string
		return ret
	}
	return *o.HeaderName.Get()
}

// GetHeaderNameOk returns a tuple with the HeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsAdditional) GetHeaderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderName.Get(), o.HeaderName.IsSet()
}

// HasHeaderName returns a boolean if a field has been set.
func (o *LoggingHttpsAdditional) HasHeaderName() bool {
	if o != nil && o.HeaderName.IsSet() {
		return true
	}

	return false
}

// SetHeaderName gets a reference to the given NullableString and assigns it to the HeaderName field.
func (o *LoggingHttpsAdditional) SetHeaderName(v string) {
	o.HeaderName.Set(&v)
}

// SetHeaderNameNil sets the value for HeaderName to be an explicit nil
func (o *LoggingHttpsAdditional) SetHeaderNameNil() {
	o.HeaderName.Set(nil)
}

// UnsetHeaderName ensures that no value is present for HeaderName, not even an explicit nil
func (o *LoggingHttpsAdditional) UnsetHeaderName() {
	o.HeaderName.Unset()
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LoggingHttpsAdditional) GetMessageType() LoggingMessageType {
	if o == nil || o.MessageType == nil {
		var ret LoggingMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsAdditional) GetMessageTypeOk() (*LoggingMessageType, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LoggingHttpsAdditional) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given LoggingMessageType and assigns it to the MessageType field.
func (o *LoggingHttpsAdditional) SetMessageType(v LoggingMessageType) {
	o.MessageType = &v
}

// GetHeaderValue returns the HeaderValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHttpsAdditional) GetHeaderValue() string {
	if o == nil || o.HeaderValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.HeaderValue.Get()
}

// GetHeaderValueOk returns a tuple with the HeaderValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHttpsAdditional) GetHeaderValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderValue.Get(), o.HeaderValue.IsSet()
}

// HasHeaderValue returns a boolean if a field has been set.
func (o *LoggingHttpsAdditional) HasHeaderValue() bool {
	if o != nil && o.HeaderValue.IsSet() {
		return true
	}

	return false
}

// SetHeaderValue gets a reference to the given NullableString and assigns it to the HeaderValue field.
func (o *LoggingHttpsAdditional) SetHeaderValue(v string) {
	o.HeaderValue.Set(&v)
}

// SetHeaderValueNil sets the value for HeaderValue to be an explicit nil
func (o *LoggingHttpsAdditional) SetHeaderValueNil() {
	o.HeaderValue.Set(nil)
}

// UnsetHeaderValue ensures that no value is present for HeaderValue, not even an explicit nil
func (o *LoggingHttpsAdditional) UnsetHeaderValue() {
	o.HeaderValue.Unset()
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LoggingHttpsAdditional) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsAdditional) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LoggingHttpsAdditional) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *LoggingHttpsAdditional) SetMethod(v string) {
	o.Method = &v
}

// GetJsonFormat returns the JsonFormat field value if set, zero value otherwise.
func (o *LoggingHttpsAdditional) GetJsonFormat() string {
	if o == nil || o.JsonFormat == nil {
		var ret string
		return ret
	}
	return *o.JsonFormat
}

// GetJsonFormatOk returns a tuple with the JsonFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsAdditional) GetJsonFormatOk() (*string, bool) {
	if o == nil || o.JsonFormat == nil {
		return nil, false
	}
	return o.JsonFormat, true
}

// HasJsonFormat returns a boolean if a field has been set.
func (o *LoggingHttpsAdditional) HasJsonFormat() bool {
	if o != nil && o.JsonFormat != nil {
		return true
	}

	return false
}

// SetJsonFormat gets a reference to the given string and assigns it to the JsonFormat field.
func (o *LoggingHttpsAdditional) SetJsonFormat(v string) {
	o.JsonFormat = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingHttpsAdditional) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsAdditional) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingHttpsAdditional) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingHttpsAdditional) SetFormat(v string) {
	o.Format = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *LoggingHttpsAdditional) GetPeriod() int32 {
	if o == nil || o.Period == nil {
		var ret int32
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHttpsAdditional) GetPeriodOk() (*int32, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *LoggingHttpsAdditional) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int32 and assigns it to the Period field.
func (o *LoggingHttpsAdditional) SetPeriod(v int32) {
	o.Period = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingHttpsAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.RequestMaxEntries != nil {
		toSerialize["request_max_entries"] = o.RequestMaxEntries
	}
	if o.RequestMaxBytes != nil {
		toSerialize["request_max_bytes"] = o.RequestMaxBytes
	}
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	if o.HeaderName.IsSet() {
		toSerialize["header_name"] = o.HeaderName.Get()
	}
	if o.MessageType != nil {
		toSerialize["message_type"] = o.MessageType
	}
	if o.HeaderValue.IsSet() {
		toSerialize["header_value"] = o.HeaderValue.Get()
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.JsonFormat != nil {
		toSerialize["json_format"] = o.JsonFormat
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingHttpsAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingHttpsAdditional := _LoggingHttpsAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingHttpsAdditional); err == nil {
		*o = LoggingHttpsAdditional(varLoggingHttpsAdditional)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "request_max_entries")
		delete(additionalProperties, "request_max_bytes")
		delete(additionalProperties, "content_type")
		delete(additionalProperties, "header_name")
		delete(additionalProperties, "message_type")
		delete(additionalProperties, "header_value")
		delete(additionalProperties, "method")
		delete(additionalProperties, "json_format")
		delete(additionalProperties, "format")
		delete(additionalProperties, "period")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingHttpsAdditional is a helper abstraction for handling nullable logginghttpsadditional types.
type NullableLoggingHttpsAdditional struct {
	value *LoggingHttpsAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingHttpsAdditional) Get() *LoggingHttpsAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingHttpsAdditional) Set(val *LoggingHttpsAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingHttpsAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingHttpsAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingHttpsAdditional returns a pointer to a new instance of NullableLoggingHttpsAdditional.
func NewNullableLoggingHttpsAdditional(val *LoggingHttpsAdditional) *NullableLoggingHttpsAdditional {
	return &NullableLoggingHttpsAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingHttpsAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingHttpsAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
