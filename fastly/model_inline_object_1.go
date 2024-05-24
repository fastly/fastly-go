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

// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	Vcl string `json:"vcl"`
	AdditionalProperties map[string]any
}

type _InlineObject1 InlineObject1

// NewInlineObject1 instantiates a new InlineObject1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1(vcl string) *InlineObject1 {
	this := InlineObject1{}
	this.Vcl = vcl
	return &this
}

// NewInlineObject1WithDefaults instantiates a new InlineObject1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1WithDefaults() *InlineObject1 {
	this := InlineObject1{}
	return &this
}

// GetVcl returns the Vcl field value
func (o *InlineObject1) GetVcl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vcl
}

// GetVclOk returns a tuple with the Vcl field value
// and a boolean to check if the value has been set.
func (o *InlineObject1) GetVclOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vcl, true
}

// SetVcl sets field value
func (o *InlineObject1) SetVcl(v string) {
	o.Vcl = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineObject1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["vcl"] = o.Vcl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *InlineObject1) UnmarshalJSON(bytes []byte) (err error) {
	varInlineObject1 := _InlineObject1{}

	if err = json.Unmarshal(bytes, &varInlineObject1); err == nil {
		*o = InlineObject1(varInlineObject1)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "vcl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineObject1 is a helper abstraction for handling nullable inlineobject1 types. 
type NullableInlineObject1 struct {
	value *InlineObject1
	isSet bool
}

// Get returns the value.
func (v NullableInlineObject1) Get() *InlineObject1 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineObject1) Set(val *InlineObject1) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineObject1) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineObject1) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineObject1 returns a pointer to a new instance of NullableInlineObject1.
func NewNullableInlineObject1(val *InlineObject1) *NullableInlineObject1 {
	return &NullableInlineObject1{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineObject1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableInlineObject1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
