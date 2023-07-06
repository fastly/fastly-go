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

// InlineResponse400 struct for InlineResponse400
type InlineResponse400 struct {
	Code *string `json:"code,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse400 InlineResponse400

// NewInlineResponse400 instantiates a new InlineResponse400 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse400() *InlineResponse400 {
	this := InlineResponse400{}
	return &this
}

// NewInlineResponse400WithDefaults instantiates a new InlineResponse400 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse400WithDefaults() *InlineResponse400 {
	this := InlineResponse400{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *InlineResponse400) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *InlineResponse400) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *InlineResponse400) SetCode(v string) {
	o.Code = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse400) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *InlineResponse400) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse400 := _InlineResponse400{}

	if err = json.Unmarshal(bytes, &varInlineResponse400); err == nil {
		*o = InlineResponse400(varInlineResponse400)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse400 is a helper abstraction for handling nullable inlineresponse400 types. 
type NullableInlineResponse400 struct {
	value *InlineResponse400
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse400) Get() *InlineResponse400 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse400) Set(val *InlineResponse400) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse400) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse400) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse400 returns a pointer to a new instance of NullableInlineResponse400.
func NewNullableInlineResponse400(val *InlineResponse400) *NullableInlineResponse400 {
	return &NullableInlineResponse400{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse400) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableInlineResponse400) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
