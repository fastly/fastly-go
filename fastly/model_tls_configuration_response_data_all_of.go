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

// TLSConfigurationResponseDataAllOf struct for TLSConfigurationResponseDataAllOf
type TLSConfigurationResponseDataAllOf struct {
	ID *string `json:"id,omitempty"`
	Attributes *TLSConfigurationResponseAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSConfigurationResponseDataAllOf TLSConfigurationResponseDataAllOf

// NewTLSConfigurationResponseDataAllOf instantiates a new TLSConfigurationResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSConfigurationResponseDataAllOf() *TLSConfigurationResponseDataAllOf {
	this := TLSConfigurationResponseDataAllOf{}
	return &this
}

// NewTLSConfigurationResponseDataAllOfWithDefaults instantiates a new TLSConfigurationResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSConfigurationResponseDataAllOfWithDefaults() *TLSConfigurationResponseDataAllOf {
	this := TLSConfigurationResponseDataAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TLSConfigurationResponseDataAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationResponseDataAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TLSConfigurationResponseDataAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TLSConfigurationResponseDataAllOf) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSConfigurationResponseDataAllOf) GetAttributes() TLSConfigurationResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret TLSConfigurationResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationResponseDataAllOf) GetAttributesOk() (*TLSConfigurationResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSConfigurationResponseDataAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSConfigurationResponseAttributes and assigns it to the Attributes field.
func (o *TLSConfigurationResponseDataAllOf) SetAttributes(v TLSConfigurationResponseAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSConfigurationResponseDataAllOf) MarshalJSON() ([]byte, error) {
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
func (o *TLSConfigurationResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTLSConfigurationResponseDataAllOf := _TLSConfigurationResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varTLSConfigurationResponseDataAllOf); err == nil {
		*o = TLSConfigurationResponseDataAllOf(varTLSConfigurationResponseDataAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSConfigurationResponseDataAllOf is a helper abstraction for handling nullable tlsconfigurationresponsedataallof types. 
type NullableTLSConfigurationResponseDataAllOf struct {
	value *TLSConfigurationResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTLSConfigurationResponseDataAllOf) Get() *TLSConfigurationResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSConfigurationResponseDataAllOf) Set(val *TLSConfigurationResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSConfigurationResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSConfigurationResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSConfigurationResponseDataAllOf returns a pointer to a new instance of NullableTLSConfigurationResponseDataAllOf.
func NewNullableTLSConfigurationResponseDataAllOf(val *TLSConfigurationResponseDataAllOf) *NullableTLSConfigurationResponseDataAllOf {
	return &NullableTLSConfigurationResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSConfigurationResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSConfigurationResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
