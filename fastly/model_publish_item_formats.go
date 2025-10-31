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

// PublishItemFormats Transport-specific message payload representations to be used for delivery. At least one format (`http-response`, `http-stream`, and/or `ws-message`) must be specified. Messages are only delivered to subscribers interested in the provided formats. For example, the `ws-message` format will only be sent to WebSocket clients.
type PublishItemFormats struct {
	HttpResponse         *HttpResponseFormat `json:"http-response,omitempty"`
	HttpStream           *HttpStreamFormat   `json:"http-stream,omitempty"`
	WsMessage            *WsMessageFormat    `json:"ws-message,omitempty"`
	AdditionalProperties map[string]any
}

type _PublishItemFormats PublishItemFormats

// NewPublishItemFormats instantiates a new PublishItemFormats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishItemFormats() *PublishItemFormats {
	this := PublishItemFormats{}
	return &this
}

// NewPublishItemFormatsWithDefaults instantiates a new PublishItemFormats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishItemFormatsWithDefaults() *PublishItemFormats {
	this := PublishItemFormats{}
	return &this
}

// GetHttpResponse returns the HttpResponse field value if set, zero value otherwise.
func (o *PublishItemFormats) GetHttpResponse() HttpResponseFormat {
	if o == nil || o.HttpResponse == nil {
		var ret HttpResponseFormat
		return ret
	}
	return *o.HttpResponse
}

// GetHttpResponseOk returns a tuple with the HttpResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishItemFormats) GetHttpResponseOk() (*HttpResponseFormat, bool) {
	if o == nil || o.HttpResponse == nil {
		return nil, false
	}
	return o.HttpResponse, true
}

// HasHttpResponse returns a boolean if a field has been set.
func (o *PublishItemFormats) HasHttpResponse() bool {
	if o != nil && o.HttpResponse != nil {
		return true
	}

	return false
}

// SetHttpResponse gets a reference to the given HttpResponseFormat and assigns it to the HttpResponse field.
func (o *PublishItemFormats) SetHttpResponse(v HttpResponseFormat) {
	o.HttpResponse = &v
}

// GetHttpStream returns the HttpStream field value if set, zero value otherwise.
func (o *PublishItemFormats) GetHttpStream() HttpStreamFormat {
	if o == nil || o.HttpStream == nil {
		var ret HttpStreamFormat
		return ret
	}
	return *o.HttpStream
}

// GetHttpStreamOk returns a tuple with the HttpStream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishItemFormats) GetHttpStreamOk() (*HttpStreamFormat, bool) {
	if o == nil || o.HttpStream == nil {
		return nil, false
	}
	return o.HttpStream, true
}

// HasHttpStream returns a boolean if a field has been set.
func (o *PublishItemFormats) HasHttpStream() bool {
	if o != nil && o.HttpStream != nil {
		return true
	}

	return false
}

// SetHttpStream gets a reference to the given HttpStreamFormat and assigns it to the HttpStream field.
func (o *PublishItemFormats) SetHttpStream(v HttpStreamFormat) {
	o.HttpStream = &v
}

// GetWsMessage returns the WsMessage field value if set, zero value otherwise.
func (o *PublishItemFormats) GetWsMessage() WsMessageFormat {
	if o == nil || o.WsMessage == nil {
		var ret WsMessageFormat
		return ret
	}
	return *o.WsMessage
}

// GetWsMessageOk returns a tuple with the WsMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishItemFormats) GetWsMessageOk() (*WsMessageFormat, bool) {
	if o == nil || o.WsMessage == nil {
		return nil, false
	}
	return o.WsMessage, true
}

// HasWsMessage returns a boolean if a field has been set.
func (o *PublishItemFormats) HasWsMessage() bool {
	if o != nil && o.WsMessage != nil {
		return true
	}

	return false
}

// SetWsMessage gets a reference to the given WsMessageFormat and assigns it to the WsMessage field.
func (o *PublishItemFormats) SetWsMessage(v WsMessageFormat) {
	o.WsMessage = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PublishItemFormats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.HttpResponse != nil {
		toSerialize["http-response"] = o.HttpResponse
	}
	if o.HttpStream != nil {
		toSerialize["http-stream"] = o.HttpStream
	}
	if o.WsMessage != nil {
		toSerialize["ws-message"] = o.WsMessage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PublishItemFormats) UnmarshalJSON(bytes []byte) (err error) {
	varPublishItemFormats := _PublishItemFormats{}

	if err = json.Unmarshal(bytes, &varPublishItemFormats); err == nil {
		*o = PublishItemFormats(varPublishItemFormats)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "http-response")
		delete(additionalProperties, "http-stream")
		delete(additionalProperties, "ws-message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePublishItemFormats is a helper abstraction for handling nullable publishitemformats types.
type NullablePublishItemFormats struct {
	value *PublishItemFormats
	isSet bool
}

// Get returns the value.
func (v NullablePublishItemFormats) Get() *PublishItemFormats {
	return v.value
}

// Set modifies the value.
func (v *NullablePublishItemFormats) Set(val *PublishItemFormats) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePublishItemFormats) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePublishItemFormats) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePublishItemFormats returns a pointer to a new instance of NullablePublishItemFormats.
func NewNullablePublishItemFormats(val *PublishItemFormats) *NullablePublishItemFormats {
	return &NullablePublishItemFormats{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePublishItemFormats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePublishItemFormats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
