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

// TLSBulkCertificateResponseDataAllOf struct for TLSBulkCertificateResponseDataAllOf
type TLSBulkCertificateResponseDataAllOf struct {
	ID                   *string                               `json:"id,omitempty"`
	Attributes           *TLSBulkCertificateResponseAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSBulkCertificateResponseDataAllOf TLSBulkCertificateResponseDataAllOf

// NewTLSBulkCertificateResponseDataAllOf instantiates a new TLSBulkCertificateResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSBulkCertificateResponseDataAllOf() *TLSBulkCertificateResponseDataAllOf {
	this := TLSBulkCertificateResponseDataAllOf{}
	return &this
}

// NewTLSBulkCertificateResponseDataAllOfWithDefaults instantiates a new TLSBulkCertificateResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSBulkCertificateResponseDataAllOfWithDefaults() *TLSBulkCertificateResponseDataAllOf {
	this := TLSBulkCertificateResponseDataAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TLSBulkCertificateResponseDataAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSBulkCertificateResponseDataAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TLSBulkCertificateResponseDataAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TLSBulkCertificateResponseDataAllOf) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSBulkCertificateResponseDataAllOf) GetAttributes() TLSBulkCertificateResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret TLSBulkCertificateResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSBulkCertificateResponseDataAllOf) GetAttributesOk() (*TLSBulkCertificateResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSBulkCertificateResponseDataAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSBulkCertificateResponseAttributes and assigns it to the Attributes field.
func (o *TLSBulkCertificateResponseDataAllOf) SetAttributes(v TLSBulkCertificateResponseAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSBulkCertificateResponseDataAllOf) MarshalJSON() ([]byte, error) {
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
func (o *TLSBulkCertificateResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTLSBulkCertificateResponseDataAllOf := _TLSBulkCertificateResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varTLSBulkCertificateResponseDataAllOf); err == nil {
		*o = TLSBulkCertificateResponseDataAllOf(varTLSBulkCertificateResponseDataAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSBulkCertificateResponseDataAllOf is a helper abstraction for handling nullable tlsbulkcertificateresponsedataallof types.
type NullableTLSBulkCertificateResponseDataAllOf struct {
	value *TLSBulkCertificateResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTLSBulkCertificateResponseDataAllOf) Get() *TLSBulkCertificateResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSBulkCertificateResponseDataAllOf) Set(val *TLSBulkCertificateResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSBulkCertificateResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSBulkCertificateResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSBulkCertificateResponseDataAllOf returns a pointer to a new instance of NullableTLSBulkCertificateResponseDataAllOf.
func NewNullableTLSBulkCertificateResponseDataAllOf(val *TLSBulkCertificateResponseDataAllOf) *NullableTLSBulkCertificateResponseDataAllOf {
	return &NullableTLSBulkCertificateResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSBulkCertificateResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTLSBulkCertificateResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
