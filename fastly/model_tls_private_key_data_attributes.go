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

// TLSPrivateKeyDataAttributes struct for TLSPrivateKeyDataAttributes
type TLSPrivateKeyDataAttributes struct {
	// A customizable name for your private key. Optional.
	Name *string `json:"name,omitempty"`
	// The contents of the private key. Must be a PEM-formatted key. Not returned in response body. Required.
	Key                  *string `json:"key,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSPrivateKeyDataAttributes TLSPrivateKeyDataAttributes

// NewTLSPrivateKeyDataAttributes instantiates a new TLSPrivateKeyDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSPrivateKeyDataAttributes() *TLSPrivateKeyDataAttributes {
	this := TLSPrivateKeyDataAttributes{}
	return &this
}

// NewTLSPrivateKeyDataAttributesWithDefaults instantiates a new TLSPrivateKeyDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSPrivateKeyDataAttributesWithDefaults() *TLSPrivateKeyDataAttributes {
	this := TLSPrivateKeyDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TLSPrivateKeyDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TLSPrivateKeyDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TLSPrivateKeyDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *TLSPrivateKeyDataAttributes) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyDataAttributes) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *TLSPrivateKeyDataAttributes) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *TLSPrivateKeyDataAttributes) SetKey(v string) {
	o.Key = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSPrivateKeyDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TLSPrivateKeyDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTLSPrivateKeyDataAttributes := _TLSPrivateKeyDataAttributes{}

	if err = json.Unmarshal(bytes, &varTLSPrivateKeyDataAttributes); err == nil {
		*o = TLSPrivateKeyDataAttributes(varTLSPrivateKeyDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSPrivateKeyDataAttributes is a helper abstraction for handling nullable tlsprivatekeydataattributes types.
type NullableTLSPrivateKeyDataAttributes struct {
	value *TLSPrivateKeyDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTLSPrivateKeyDataAttributes) Get() *TLSPrivateKeyDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSPrivateKeyDataAttributes) Set(val *TLSPrivateKeyDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSPrivateKeyDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSPrivateKeyDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSPrivateKeyDataAttributes returns a pointer to a new instance of NullableTLSPrivateKeyDataAttributes.
func NewNullableTLSPrivateKeyDataAttributes(val *TLSPrivateKeyDataAttributes) *NullableTLSPrivateKeyDataAttributes {
	return &NullableTLSPrivateKeyDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSPrivateKeyDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTLSPrivateKeyDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
