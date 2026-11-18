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

// TLSCsrResponseAttributes struct for TLSCsrResponseAttributes
type TLSCsrResponseAttributes struct {
	// The PEM encoded CSR.
	Content              *string `json:"content,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCsrResponseAttributes TLSCsrResponseAttributes

// NewTLSCsrResponseAttributes instantiates a new TLSCsrResponseAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCsrResponseAttributes() *TLSCsrResponseAttributes {
	this := TLSCsrResponseAttributes{}
	return &this
}

// NewTLSCsrResponseAttributesWithDefaults instantiates a new TLSCsrResponseAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCsrResponseAttributesWithDefaults() *TLSCsrResponseAttributes {
	this := TLSCsrResponseAttributes{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *TLSCsrResponseAttributes) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrResponseAttributes) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *TLSCsrResponseAttributes) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *TLSCsrResponseAttributes) SetContent(v string) {
	o.Content = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCsrResponseAttributes) MarshalJSON() ([]byte, error) {
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
func (o *TLSCsrResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCsrResponseAttributes := _TLSCsrResponseAttributes{}

	if err = json.Unmarshal(bytes, &varTLSCsrResponseAttributes); err == nil {
		*o = TLSCsrResponseAttributes(varTLSCsrResponseAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "content")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSCsrResponseAttributes is a helper abstraction for handling nullable tlscsrresponseattributes types.
type NullableTLSCsrResponseAttributes struct {
	value *TLSCsrResponseAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTLSCsrResponseAttributes) Get() *TLSCsrResponseAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCsrResponseAttributes) Set(val *TLSCsrResponseAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCsrResponseAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCsrResponseAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCsrResponseAttributes returns a pointer to a new instance of NullableTLSCsrResponseAttributes.
func NewNullableTLSCsrResponseAttributes(val *TLSCsrResponseAttributes) *NullableTLSCsrResponseAttributes {
	return &NullableTLSCsrResponseAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCsrResponseAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTLSCsrResponseAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
