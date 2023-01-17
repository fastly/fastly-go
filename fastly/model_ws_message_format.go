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

// WsMessageFormat Payload format for delivering to subscribers of WebSocket messages. One of `content` or `content-bin` must be specified.
type WsMessageFormat struct {
	// The content of a WebSocket `TEXT` message.
	Content *string `json:"content,omitempty"`
	// The base64-encoded content of a WebSocket `BINARY` message.
	ContentBin *string `json:"content-bin,omitempty"`
	AdditionalProperties map[string]any
}

type _WsMessageFormat WsMessageFormat

// NewWsMessageFormat instantiates a new WsMessageFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWsMessageFormat() *WsMessageFormat {
	this := WsMessageFormat{}
	return &this
}

// NewWsMessageFormatWithDefaults instantiates a new WsMessageFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWsMessageFormatWithDefaults() *WsMessageFormat {
	this := WsMessageFormat{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *WsMessageFormat) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsMessageFormat) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *WsMessageFormat) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *WsMessageFormat) SetContent(v string) {
	o.Content = &v
}

// GetContentBin returns the ContentBin field value if set, zero value otherwise.
func (o *WsMessageFormat) GetContentBin() string {
	if o == nil || o.ContentBin == nil {
		var ret string
		return ret
	}
	return *o.ContentBin
}

// GetContentBinOk returns a tuple with the ContentBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WsMessageFormat) GetContentBinOk() (*string, bool) {
	if o == nil || o.ContentBin == nil {
		return nil, false
	}
	return o.ContentBin, true
}

// HasContentBin returns a boolean if a field has been set.
func (o *WsMessageFormat) HasContentBin() bool {
	if o != nil && o.ContentBin != nil {
		return true
	}

	return false
}

// SetContentBin gets a reference to the given string and assigns it to the ContentBin field.
func (o *WsMessageFormat) SetContentBin(v string) {
	o.ContentBin = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WsMessageFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.ContentBin != nil {
		toSerialize["content-bin"] = o.ContentBin
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *WsMessageFormat) UnmarshalJSON(bytes []byte) (err error) {
	varWsMessageFormat := _WsMessageFormat{}

	if err = json.Unmarshal(bytes, &varWsMessageFormat); err == nil {
		*o = WsMessageFormat(varWsMessageFormat)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		delete(additionalProperties, "content-bin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWsMessageFormat is a helper abstraction for handling nullable wsmessageformat types. 
type NullableWsMessageFormat struct {
	value *WsMessageFormat
	isSet bool
}

// Get returns the value.
func (v NullableWsMessageFormat) Get() *WsMessageFormat {
	return v.value
}

// Set modifies the value.
func (v *NullableWsMessageFormat) Set(val *WsMessageFormat) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWsMessageFormat) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWsMessageFormat) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWsMessageFormat returns a pointer to a new instance of NullableWsMessageFormat.
func NewNullableWsMessageFormat(val *WsMessageFormat) *NullableWsMessageFormat {
	return &NullableWsMessageFormat{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWsMessageFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWsMessageFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
