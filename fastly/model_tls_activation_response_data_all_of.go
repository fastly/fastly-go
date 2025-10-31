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

// TlsActivationResponseDataAllOf struct for TlsActivationResponseDataAllOf
type TlsActivationResponseDataAllOf struct {
	Id                   *string     `json:"id,omitempty"`
	Attributes           *Timestamps `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsActivationResponseDataAllOf TlsActivationResponseDataAllOf

// NewTlsActivationResponseDataAllOf instantiates a new TlsActivationResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsActivationResponseDataAllOf() *TlsActivationResponseDataAllOf {
	this := TlsActivationResponseDataAllOf{}
	return &this
}

// NewTlsActivationResponseDataAllOfWithDefaults instantiates a new TlsActivationResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsActivationResponseDataAllOfWithDefaults() *TlsActivationResponseDataAllOf {
	this := TlsActivationResponseDataAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TlsActivationResponseDataAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsActivationResponseDataAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TlsActivationResponseDataAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TlsActivationResponseDataAllOf) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TlsActivationResponseDataAllOf) GetAttributes() Timestamps {
	if o == nil || o.Attributes == nil {
		var ret Timestamps
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsActivationResponseDataAllOf) GetAttributesOk() (*Timestamps, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TlsActivationResponseDataAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Timestamps and assigns it to the Attributes field.
func (o *TlsActivationResponseDataAllOf) SetAttributes(v Timestamps) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsActivationResponseDataAllOf) MarshalJSON() ([]byte, error) {
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
func (o *TlsActivationResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTlsActivationResponseDataAllOf := _TlsActivationResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varTlsActivationResponseDataAllOf); err == nil {
		*o = TlsActivationResponseDataAllOf(varTlsActivationResponseDataAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsActivationResponseDataAllOf is a helper abstraction for handling nullable tlsactivationresponsedataallof types.
type NullableTlsActivationResponseDataAllOf struct {
	value *TlsActivationResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTlsActivationResponseDataAllOf) Get() *TlsActivationResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsActivationResponseDataAllOf) Set(val *TlsActivationResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsActivationResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsActivationResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsActivationResponseDataAllOf returns a pointer to a new instance of NullableTlsActivationResponseDataAllOf.
func NewNullableTlsActivationResponseDataAllOf(val *TlsActivationResponseDataAllOf) *NullableTlsActivationResponseDataAllOf {
	return &NullableTlsActivationResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsActivationResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsActivationResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
