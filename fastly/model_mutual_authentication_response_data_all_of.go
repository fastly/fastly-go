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

// MutualAuthenticationResponseDataAllOf struct for MutualAuthenticationResponseDataAllOf
type MutualAuthenticationResponseDataAllOf struct {
	ID *string `json:"id,omitempty"`
	Attributes *MutualAuthenticationResponseAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _MutualAuthenticationResponseDataAllOf MutualAuthenticationResponseDataAllOf

// NewMutualAuthenticationResponseDataAllOf instantiates a new MutualAuthenticationResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutualAuthenticationResponseDataAllOf() *MutualAuthenticationResponseDataAllOf {
	this := MutualAuthenticationResponseDataAllOf{}
	return &this
}

// NewMutualAuthenticationResponseDataAllOfWithDefaults instantiates a new MutualAuthenticationResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutualAuthenticationResponseDataAllOfWithDefaults() *MutualAuthenticationResponseDataAllOf {
	this := MutualAuthenticationResponseDataAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *MutualAuthenticationResponseDataAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationResponseDataAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *MutualAuthenticationResponseDataAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *MutualAuthenticationResponseDataAllOf) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MutualAuthenticationResponseDataAllOf) GetAttributes() MutualAuthenticationResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret MutualAuthenticationResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationResponseDataAllOf) GetAttributesOk() (*MutualAuthenticationResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MutualAuthenticationResponseDataAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MutualAuthenticationResponseAttributes and assigns it to the Attributes field.
func (o *MutualAuthenticationResponseDataAllOf) SetAttributes(v MutualAuthenticationResponseAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o MutualAuthenticationResponseDataAllOf) MarshalJSON() ([]byte, error) {
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
func (o *MutualAuthenticationResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varMutualAuthenticationResponseDataAllOf := _MutualAuthenticationResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varMutualAuthenticationResponseDataAllOf); err == nil {
		*o = MutualAuthenticationResponseDataAllOf(varMutualAuthenticationResponseDataAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableMutualAuthenticationResponseDataAllOf is a helper abstraction for handling nullable mutualauthenticationresponsedataallof types. 
type NullableMutualAuthenticationResponseDataAllOf struct {
	value *MutualAuthenticationResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableMutualAuthenticationResponseDataAllOf) Get() *MutualAuthenticationResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableMutualAuthenticationResponseDataAllOf) Set(val *MutualAuthenticationResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableMutualAuthenticationResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableMutualAuthenticationResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableMutualAuthenticationResponseDataAllOf returns a pointer to a new instance of NullableMutualAuthenticationResponseDataAllOf.
func NewNullableMutualAuthenticationResponseDataAllOf(val *MutualAuthenticationResponseDataAllOf) *NullableMutualAuthenticationResponseDataAllOf {
	return &NullableMutualAuthenticationResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableMutualAuthenticationResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableMutualAuthenticationResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
