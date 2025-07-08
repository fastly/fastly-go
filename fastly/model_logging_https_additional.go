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

// LoggingHTTPSAdditional struct for LoggingHTTPSAdditional
type LoggingHTTPSAdditional struct {
	// The URL to send logs to. Must use HTTPS. Required.
	URL *string `json:"url,omitempty"`
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
	JSONFormat *string `json:"json_format,omitempty"`
	// A Fastly [log format string](https://www.fastly.com/documentation/guides/integrations/streaming-logs/custom-log-formats/).
	Format               *string `json:"format,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingHTTPSAdditional LoggingHTTPSAdditional

// NewLoggingHTTPSAdditional instantiates a new LoggingHTTPSAdditional object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingHTTPSAdditional() *LoggingHTTPSAdditional {
	this := LoggingHTTPSAdditional{}
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
	return &this
}

// NewLoggingHTTPSAdditionalWithDefaults instantiates a new LoggingHTTPSAdditional object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingHTTPSAdditionalWithDefaults() *LoggingHTTPSAdditional {
	this := LoggingHTTPSAdditional{}
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
	return &this
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *LoggingHTTPSAdditional) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSAdditional) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *LoggingHTTPSAdditional) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *LoggingHTTPSAdditional) SetURL(v string) {
	o.URL = &v
}

// GetRequestMaxEntries returns the RequestMaxEntries field value if set, zero value otherwise.
func (o *LoggingHTTPSAdditional) GetRequestMaxEntries() int32 {
	if o == nil || o.RequestMaxEntries == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxEntries
}

// GetRequestMaxEntriesOk returns a tuple with the RequestMaxEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSAdditional) GetRequestMaxEntriesOk() (*int32, bool) {
	if o == nil || o.RequestMaxEntries == nil {
		return nil, false
	}
	return o.RequestMaxEntries, true
}

// HasRequestMaxEntries returns a boolean if a field has been set.
func (o *LoggingHTTPSAdditional) HasRequestMaxEntries() bool {
	if o != nil && o.RequestMaxEntries != nil {
		return true
	}

	return false
}

// SetRequestMaxEntries gets a reference to the given int32 and assigns it to the RequestMaxEntries field.
func (o *LoggingHTTPSAdditional) SetRequestMaxEntries(v int32) {
	o.RequestMaxEntries = &v
}

// GetRequestMaxBytes returns the RequestMaxBytes field value if set, zero value otherwise.
func (o *LoggingHTTPSAdditional) GetRequestMaxBytes() int32 {
	if o == nil || o.RequestMaxBytes == nil {
		var ret int32
		return ret
	}
	return *o.RequestMaxBytes
}

// GetRequestMaxBytesOk returns a tuple with the RequestMaxBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSAdditional) GetRequestMaxBytesOk() (*int32, bool) {
	if o == nil || o.RequestMaxBytes == nil {
		return nil, false
	}
	return o.RequestMaxBytes, true
}

// HasRequestMaxBytes returns a boolean if a field has been set.
func (o *LoggingHTTPSAdditional) HasRequestMaxBytes() bool {
	if o != nil && o.RequestMaxBytes != nil {
		return true
	}

	return false
}

// SetRequestMaxBytes gets a reference to the given int32 and assigns it to the RequestMaxBytes field.
func (o *LoggingHTTPSAdditional) SetRequestMaxBytes(v int32) {
	o.RequestMaxBytes = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSAdditional) GetContentType() string {
	if o == nil || o.ContentType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSAdditional) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *LoggingHTTPSAdditional) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *LoggingHTTPSAdditional) SetContentType(v string) {
	o.ContentType.Set(&v)
}

// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *LoggingHTTPSAdditional) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *LoggingHTTPSAdditional) UnsetContentType() {
	o.ContentType.Unset()
}

// GetHeaderName returns the HeaderName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSAdditional) GetHeaderName() string {
	if o == nil || o.HeaderName.Get() == nil {
		var ret string
		return ret
	}
	return *o.HeaderName.Get()
}

// GetHeaderNameOk returns a tuple with the HeaderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSAdditional) GetHeaderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderName.Get(), o.HeaderName.IsSet()
}

// HasHeaderName returns a boolean if a field has been set.
func (o *LoggingHTTPSAdditional) HasHeaderName() bool {
	if o != nil && o.HeaderName.IsSet() {
		return true
	}

	return false
}

// SetHeaderName gets a reference to the given NullableString and assigns it to the HeaderName field.
func (o *LoggingHTTPSAdditional) SetHeaderName(v string) {
	o.HeaderName.Set(&v)
}

// SetHeaderNameNil sets the value for HeaderName to be an explicit nil
func (o *LoggingHTTPSAdditional) SetHeaderNameNil() {
	o.HeaderName.Set(nil)
}

// UnsetHeaderName ensures that no value is present for HeaderName, not even an explicit nil
func (o *LoggingHTTPSAdditional) UnsetHeaderName() {
	o.HeaderName.Unset()
}

// GetMessageType returns the MessageType field value if set, zero value otherwise.
func (o *LoggingHTTPSAdditional) GetMessageType() LoggingMessageType {
	if o == nil || o.MessageType == nil {
		var ret LoggingMessageType
		return ret
	}
	return *o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSAdditional) GetMessageTypeOk() (*LoggingMessageType, bool) {
	if o == nil || o.MessageType == nil {
		return nil, false
	}
	return o.MessageType, true
}

// HasMessageType returns a boolean if a field has been set.
func (o *LoggingHTTPSAdditional) HasMessageType() bool {
	if o != nil && o.MessageType != nil {
		return true
	}

	return false
}

// SetMessageType gets a reference to the given LoggingMessageType and assigns it to the MessageType field.
func (o *LoggingHTTPSAdditional) SetMessageType(v LoggingMessageType) {
	o.MessageType = &v
}

// GetHeaderValue returns the HeaderValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingHTTPSAdditional) GetHeaderValue() string {
	if o == nil || o.HeaderValue.Get() == nil {
		var ret string
		return ret
	}
	return *o.HeaderValue.Get()
}

// GetHeaderValueOk returns a tuple with the HeaderValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingHTTPSAdditional) GetHeaderValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.HeaderValue.Get(), o.HeaderValue.IsSet()
}

// HasHeaderValue returns a boolean if a field has been set.
func (o *LoggingHTTPSAdditional) HasHeaderValue() bool {
	if o != nil && o.HeaderValue.IsSet() {
		return true
	}

	return false
}

// SetHeaderValue gets a reference to the given NullableString and assigns it to the HeaderValue field.
func (o *LoggingHTTPSAdditional) SetHeaderValue(v string) {
	o.HeaderValue.Set(&v)
}

// SetHeaderValueNil sets the value for HeaderValue to be an explicit nil
func (o *LoggingHTTPSAdditional) SetHeaderValueNil() {
	o.HeaderValue.Set(nil)
}

// UnsetHeaderValue ensures that no value is present for HeaderValue, not even an explicit nil
func (o *LoggingHTTPSAdditional) UnsetHeaderValue() {
	o.HeaderValue.Unset()
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *LoggingHTTPSAdditional) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSAdditional) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *LoggingHTTPSAdditional) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *LoggingHTTPSAdditional) SetMethod(v string) {
	o.Method = &v
}

// GetJSONFormat returns the JSONFormat field value if set, zero value otherwise.
func (o *LoggingHTTPSAdditional) GetJSONFormat() string {
	if o == nil || o.JSONFormat == nil {
		var ret string
		return ret
	}
	return *o.JSONFormat
}

// GetJSONFormatOk returns a tuple with the JSONFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSAdditional) GetJSONFormatOk() (*string, bool) {
	if o == nil || o.JSONFormat == nil {
		return nil, false
	}
	return o.JSONFormat, true
}

// HasJSONFormat returns a boolean if a field has been set.
func (o *LoggingHTTPSAdditional) HasJSONFormat() bool {
	if o != nil && o.JSONFormat != nil {
		return true
	}

	return false
}

// SetJSONFormat gets a reference to the given string and assigns it to the JSONFormat field.
func (o *LoggingHTTPSAdditional) SetJSONFormat(v string) {
	o.JSONFormat = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingHTTPSAdditional) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingHTTPSAdditional) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingHTTPSAdditional) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingHTTPSAdditional) SetFormat(v string) {
	o.Format = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingHTTPSAdditional) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.URL != nil {
		toSerialize["url"] = o.URL
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
	if o.JSONFormat != nil {
		toSerialize["json_format"] = o.JSONFormat
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingHTTPSAdditional) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingHTTPSAdditional := _LoggingHTTPSAdditional{}

	if err = json.Unmarshal(bytes, &varLoggingHTTPSAdditional); err == nil {
		*o = LoggingHTTPSAdditional(varLoggingHTTPSAdditional)
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
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingHTTPSAdditional is a helper abstraction for handling nullable logginghttpsadditional types.
type NullableLoggingHTTPSAdditional struct {
	value *LoggingHTTPSAdditional
	isSet bool
}

// Get returns the value.
func (v NullableLoggingHTTPSAdditional) Get() *LoggingHTTPSAdditional {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingHTTPSAdditional) Set(val *LoggingHTTPSAdditional) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingHTTPSAdditional) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingHTTPSAdditional) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingHTTPSAdditional returns a pointer to a new instance of NullableLoggingHTTPSAdditional.
func NewNullableLoggingHTTPSAdditional(val *LoggingHTTPSAdditional) *NullableLoggingHTTPSAdditional {
	return &NullableLoggingHTTPSAdditional{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingHTTPSAdditional) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingHTTPSAdditional) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
