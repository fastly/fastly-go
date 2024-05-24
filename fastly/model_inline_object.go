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

// InlineObject struct for InlineObject
type InlineObject struct {
	Vcl string `json:"vcl"`
	AdditionalProperties map[string]any
}

type _InlineObject InlineObject

// NewInlineObject instantiates a new InlineObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject(vcl string) *InlineObject {
	this := InlineObject{}
	this.Vcl = vcl
	return &this
}

// NewInlineObjectWithDefaults instantiates a new InlineObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObjectWithDefaults() *InlineObject {
	this := InlineObject{}
	return &this
}

// GetVcl returns the Vcl field value
func (o *InlineObject) GetVcl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vcl
}

// GetVclOk returns a tuple with the Vcl field value
// and a boolean to check if the value has been set.
func (o *InlineObject) GetVclOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vcl, true
}

// SetVcl sets field value
func (o *InlineObject) SetVcl(v string) {
	o.Vcl = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineObject) MarshalJSON() ([]byte, error) {
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
func (o *InlineObject) UnmarshalJSON(bytes []byte) (err error) {
	varInlineObject := _InlineObject{}

	if err = json.Unmarshal(bytes, &varInlineObject); err == nil {
		*o = InlineObject(varInlineObject)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "vcl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineObject is a helper abstraction for handling nullable inlineobject types. 
type NullableInlineObject struct {
	value *InlineObject
	isSet bool
}

// Get returns the value.
func (v NullableInlineObject) Get() *InlineObject {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineObject) Set(val *InlineObject) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineObject) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineObject) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineObject returns a pointer to a new instance of NullableInlineObject.
func NewNullableInlineObject(val *InlineObject) *NullableInlineObject {
	return &NullableInlineObject{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableInlineObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
