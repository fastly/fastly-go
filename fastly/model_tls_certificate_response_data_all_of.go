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

// TLSCertificateResponseDataAllOf struct for TLSCertificateResponseDataAllOf
type TLSCertificateResponseDataAllOf struct {
	ID *string `json:"id,omitempty"`
	Attributes *TLSCertificateResponseAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCertificateResponseDataAllOf TLSCertificateResponseDataAllOf

// NewTLSCertificateResponseDataAllOf instantiates a new TLSCertificateResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCertificateResponseDataAllOf() *TLSCertificateResponseDataAllOf {
	this := TLSCertificateResponseDataAllOf{}
	return &this
}

// NewTLSCertificateResponseDataAllOfWithDefaults instantiates a new TLSCertificateResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCertificateResponseDataAllOfWithDefaults() *TLSCertificateResponseDataAllOf {
	this := TLSCertificateResponseDataAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TLSCertificateResponseDataAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCertificateResponseDataAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TLSCertificateResponseDataAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TLSCertificateResponseDataAllOf) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSCertificateResponseDataAllOf) GetAttributes() TLSCertificateResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret TLSCertificateResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCertificateResponseDataAllOf) GetAttributesOk() (*TLSCertificateResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSCertificateResponseDataAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSCertificateResponseAttributes and assigns it to the Attributes field.
func (o *TLSCertificateResponseDataAllOf) SetAttributes(v TLSCertificateResponseAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCertificateResponseDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
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
func (o *TLSCertificateResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCertificateResponseDataAllOf := _TLSCertificateResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varTLSCertificateResponseDataAllOf); err == nil {
		*o = TLSCertificateResponseDataAllOf(varTLSCertificateResponseDataAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSCertificateResponseDataAllOf is a helper abstraction for handling nullable tlscertificateresponsedataallof types. 
type NullableTLSCertificateResponseDataAllOf struct {
	value *TLSCertificateResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTLSCertificateResponseDataAllOf) Get() *TLSCertificateResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCertificateResponseDataAllOf) Set(val *TLSCertificateResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCertificateResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCertificateResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCertificateResponseDataAllOf returns a pointer to a new instance of NullableTLSCertificateResponseDataAllOf.
func NewNullableTLSCertificateResponseDataAllOf(val *TLSCertificateResponseDataAllOf) *NullableTLSCertificateResponseDataAllOf {
	return &NullableTLSCertificateResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCertificateResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSCertificateResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
