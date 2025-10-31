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

// TlsConfigurationResponseDataAllOf struct for TlsConfigurationResponseDataAllOf
type TlsConfigurationResponseDataAllOf struct {
	Id                   *string                             `json:"id,omitempty"`
	Attributes           *TlsConfigurationResponseAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsConfigurationResponseDataAllOf TlsConfigurationResponseDataAllOf

// NewTlsConfigurationResponseDataAllOf instantiates a new TlsConfigurationResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsConfigurationResponseDataAllOf() *TlsConfigurationResponseDataAllOf {
	this := TlsConfigurationResponseDataAllOf{}
	return &this
}

// NewTlsConfigurationResponseDataAllOfWithDefaults instantiates a new TlsConfigurationResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsConfigurationResponseDataAllOfWithDefaults() *TlsConfigurationResponseDataAllOf {
	this := TlsConfigurationResponseDataAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TlsConfigurationResponseDataAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationResponseDataAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TlsConfigurationResponseDataAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TlsConfigurationResponseDataAllOf) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TlsConfigurationResponseDataAllOf) GetAttributes() TlsConfigurationResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret TlsConfigurationResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationResponseDataAllOf) GetAttributesOk() (*TlsConfigurationResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TlsConfigurationResponseDataAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TlsConfigurationResponseAttributes and assigns it to the Attributes field.
func (o *TlsConfigurationResponseDataAllOf) SetAttributes(v TlsConfigurationResponseAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsConfigurationResponseDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsConfigurationResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTlsConfigurationResponseDataAllOf := _TlsConfigurationResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varTlsConfigurationResponseDataAllOf); err == nil {
		*o = TlsConfigurationResponseDataAllOf(varTlsConfigurationResponseDataAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsConfigurationResponseDataAllOf is a helper abstraction for handling nullable tlsconfigurationresponsedataallof types.
type NullableTlsConfigurationResponseDataAllOf struct {
	value *TlsConfigurationResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTlsConfigurationResponseDataAllOf) Get() *TlsConfigurationResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsConfigurationResponseDataAllOf) Set(val *TlsConfigurationResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsConfigurationResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsConfigurationResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsConfigurationResponseDataAllOf returns a pointer to a new instance of NullableTlsConfigurationResponseDataAllOf.
func NewNullableTlsConfigurationResponseDataAllOf(val *TlsConfigurationResponseDataAllOf) *NullableTlsConfigurationResponseDataAllOf {
	return &NullableTlsConfigurationResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsConfigurationResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsConfigurationResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
