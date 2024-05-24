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

// HTTPStreamFormat Payload format for delivering to subscribers of HTTP streaming response bodies (`stream` hold mode). One of `content` or `content-bin` must be specified.
type HTTPStreamFormat struct {
	// A fragment of body data as a string.
	Content *string `json:"content,omitempty"`
	// A fragment of body data as a base64-encoded binary blob.
	ContentBin *string `json:"content-bin,omitempty"`
	AdditionalProperties map[string]any
}

type _HTTPStreamFormat HTTPStreamFormat

// NewHTTPStreamFormat instantiates a new HTTPStreamFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHTTPStreamFormat() *HTTPStreamFormat {
	this := HTTPStreamFormat{}
	return &this
}

// NewHTTPStreamFormatWithDefaults instantiates a new HTTPStreamFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHTTPStreamFormatWithDefaults() *HTTPStreamFormat {
	this := HTTPStreamFormat{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *HTTPStreamFormat) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPStreamFormat) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *HTTPStreamFormat) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *HTTPStreamFormat) SetContent(v string) {
	o.Content = &v
}

// GetContentBin returns the ContentBin field value if set, zero value otherwise.
func (o *HTTPStreamFormat) GetContentBin() string {
	if o == nil || o.ContentBin == nil {
		var ret string
		return ret
	}
	return *o.ContentBin
}

// GetContentBinOk returns a tuple with the ContentBin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HTTPStreamFormat) GetContentBinOk() (*string, bool) {
	if o == nil || o.ContentBin == nil {
		return nil, false
	}
	return o.ContentBin, true
}

// HasContentBin returns a boolean if a field has been set.
func (o *HTTPStreamFormat) HasContentBin() bool {
	if o != nil && o.ContentBin != nil {
		return true
	}

	return false
}

// SetContentBin gets a reference to the given string and assigns it to the ContentBin field.
func (o *HTTPStreamFormat) SetContentBin(v string) {
	o.ContentBin = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HTTPStreamFormat) MarshalJSON() ([]byte, error) {
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
func (o *HTTPStreamFormat) UnmarshalJSON(bytes []byte) (err error) {
	varHTTPStreamFormat := _HTTPStreamFormat{}

	if err = json.Unmarshal(bytes, &varHTTPStreamFormat); err == nil {
		*o = HTTPStreamFormat(varHTTPStreamFormat)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		delete(additionalProperties, "content-bin")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHTTPStreamFormat is a helper abstraction for handling nullable httpstreamformat types. 
type NullableHTTPStreamFormat struct {
	value *HTTPStreamFormat
	isSet bool
}

// Get returns the value.
func (v NullableHTTPStreamFormat) Get() *HTTPStreamFormat {
	return v.value
}

// Set modifies the value.
func (v *NullableHTTPStreamFormat) Set(val *HTTPStreamFormat) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHTTPStreamFormat) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHTTPStreamFormat) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHTTPStreamFormat returns a pointer to a new instance of NullableHTTPStreamFormat.
func NewNullableHTTPStreamFormat(val *HTTPStreamFormat) *NullableHTTPStreamFormat {
	return &NullableHTTPStreamFormat{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHTTPStreamFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHTTPStreamFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
