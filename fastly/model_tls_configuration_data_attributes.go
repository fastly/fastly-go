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

// TLSConfigurationDataAttributes struct for TLSConfigurationDataAttributes
type TLSConfigurationDataAttributes struct {
	// A custom name for your TLS configuration.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSConfigurationDataAttributes TLSConfigurationDataAttributes

// NewTLSConfigurationDataAttributes instantiates a new TLSConfigurationDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSConfigurationDataAttributes() *TLSConfigurationDataAttributes {
	this := TLSConfigurationDataAttributes{}
	return &this
}

// NewTLSConfigurationDataAttributesWithDefaults instantiates a new TLSConfigurationDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSConfigurationDataAttributesWithDefaults() *TLSConfigurationDataAttributes {
	this := TLSConfigurationDataAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TLSConfigurationDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TLSConfigurationDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TLSConfigurationDataAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSConfigurationDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *TLSConfigurationDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTLSConfigurationDataAttributes := _TLSConfigurationDataAttributes{}

	if err = json.Unmarshal(bytes, &varTLSConfigurationDataAttributes); err == nil {
		*o = TLSConfigurationDataAttributes(varTLSConfigurationDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSConfigurationDataAttributes is a helper abstraction for handling nullable tlsconfigurationdataattributes types. 
type NullableTLSConfigurationDataAttributes struct {
	value *TLSConfigurationDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTLSConfigurationDataAttributes) Get() *TLSConfigurationDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSConfigurationDataAttributes) Set(val *TLSConfigurationDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSConfigurationDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSConfigurationDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSConfigurationDataAttributes returns a pointer to a new instance of NullableTLSConfigurationDataAttributes.
func NewNullableTLSConfigurationDataAttributes(val *TLSConfigurationDataAttributes) *NullableTLSConfigurationDataAttributes {
	return &NullableTLSConfigurationDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSConfigurationDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSConfigurationDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
