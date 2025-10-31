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

// TlsCsrResponseAttributes struct for TlsCsrResponseAttributes
type TlsCsrResponseAttributes struct {
	// The PEM encoded CSR.
	Content              *string `json:"content,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsCsrResponseAttributes TlsCsrResponseAttributes

// NewTlsCsrResponseAttributes instantiates a new TlsCsrResponseAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsCsrResponseAttributes() *TlsCsrResponseAttributes {
	this := TlsCsrResponseAttributes{}
	return &this
}

// NewTlsCsrResponseAttributesWithDefaults instantiates a new TlsCsrResponseAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsCsrResponseAttributesWithDefaults() *TlsCsrResponseAttributes {
	this := TlsCsrResponseAttributes{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *TlsCsrResponseAttributes) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCsrResponseAttributes) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *TlsCsrResponseAttributes) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *TlsCsrResponseAttributes) SetContent(v string) {
	o.Content = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsCsrResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsCsrResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTlsCsrResponseAttributes := _TlsCsrResponseAttributes{}

	if err = json.Unmarshal(bytes, &varTlsCsrResponseAttributes); err == nil {
		*o = TlsCsrResponseAttributes(varTlsCsrResponseAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsCsrResponseAttributes is a helper abstraction for handling nullable tlscsrresponseattributes types.
type NullableTlsCsrResponseAttributes struct {
	value *TlsCsrResponseAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTlsCsrResponseAttributes) Get() *TlsCsrResponseAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsCsrResponseAttributes) Set(val *TlsCsrResponseAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsCsrResponseAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsCsrResponseAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsCsrResponseAttributes returns a pointer to a new instance of NullableTlsCsrResponseAttributes.
func NewNullableTlsCsrResponseAttributes(val *TlsCsrResponseAttributes) *NullableTlsCsrResponseAttributes {
	return &NullableTlsCsrResponseAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsCsrResponseAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsCsrResponseAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
