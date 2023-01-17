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

// Dictionary struct for Dictionary
type Dictionary struct {
	// Name for the Dictionary (must start with an alphabetic character and can contain only alphanumeric characters, underscores, and whitespace).
	Name *string `json:"name,omitempty"`
	// Determines if items in the dictionary are readable or not.
	WriteOnly *bool `json:"write_only,omitempty"`
	AdditionalProperties map[string]any
}

type _Dictionary Dictionary

// NewDictionary instantiates a new Dictionary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDictionary() *Dictionary {
	this := Dictionary{}
	var writeOnly bool = false
	this.WriteOnly = &writeOnly
	return &this
}

// NewDictionaryWithDefaults instantiates a new Dictionary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDictionaryWithDefaults() *Dictionary {
	this := Dictionary{}
	var writeOnly bool = false
	this.WriteOnly = &writeOnly
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Dictionary) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dictionary) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Dictionary) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Dictionary) SetName(v string) {
	o.Name = &v
}

// GetWriteOnly returns the WriteOnly field value if set, zero value otherwise.
func (o *Dictionary) GetWriteOnly() bool {
	if o == nil || o.WriteOnly == nil {
		var ret bool
		return ret
	}
	return *o.WriteOnly
}

// GetWriteOnlyOk returns a tuple with the WriteOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dictionary) GetWriteOnlyOk() (*bool, bool) {
	if o == nil || o.WriteOnly == nil {
		return nil, false
	}
	return o.WriteOnly, true
}

// HasWriteOnly returns a boolean if a field has been set.
func (o *Dictionary) HasWriteOnly() bool {
	if o != nil && o.WriteOnly != nil {
		return true
	}

	return false
}

// SetWriteOnly gets a reference to the given bool and assigns it to the WriteOnly field.
func (o *Dictionary) SetWriteOnly(v bool) {
	o.WriteOnly = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Dictionary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.WriteOnly != nil {
		toSerialize["write_only"] = o.WriteOnly
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Dictionary) UnmarshalJSON(bytes []byte) (err error) {
	varDictionary := _Dictionary{}

	if err = json.Unmarshal(bytes, &varDictionary); err == nil {
		*o = Dictionary(varDictionary)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "write_only")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDictionary is a helper abstraction for handling nullable dictionary types. 
type NullableDictionary struct {
	value *Dictionary
	isSet bool
}

// Get returns the value.
func (v NullableDictionary) Get() *Dictionary {
	return v.value
}

// Set modifies the value.
func (v *NullableDictionary) Set(val *Dictionary) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDictionary) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDictionary) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDictionary returns a pointer to a new instance of NullableDictionary.
func NewNullableDictionary(val *Dictionary) *NullableDictionary {
	return &NullableDictionary{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDictionary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableDictionary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
