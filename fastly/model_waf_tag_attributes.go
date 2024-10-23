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

// WafTagAttributes struct for WafTagAttributes
type WafTagAttributes struct {
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _WafTagAttributes WafTagAttributes

// NewWafTagAttributes instantiates a new WafTagAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafTagAttributes() *WafTagAttributes {
	this := WafTagAttributes{}
	return &this
}

// NewWafTagAttributesWithDefaults instantiates a new WafTagAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafTagAttributesWithDefaults() *WafTagAttributes {
	this := WafTagAttributes{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WafTagAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTagAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WafTagAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WafTagAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafTagAttributes) MarshalJSON() ([]byte, error) {
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
func (o *WafTagAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varWafTagAttributes := _WafTagAttributes{}

	if err = json.Unmarshal(bytes, &varWafTagAttributes); err == nil {
		*o = WafTagAttributes(varWafTagAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafTagAttributes is a helper abstraction for handling nullable waftagattributes types.
type NullableWafTagAttributes struct {
	value *WafTagAttributes
	isSet bool
}

// Get returns the value.
func (v NullableWafTagAttributes) Get() *WafTagAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableWafTagAttributes) Set(val *WafTagAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafTagAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafTagAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafTagAttributes returns a pointer to a new instance of NullableWafTagAttributes.
func NewNullableWafTagAttributes(val *WafTagAttributes) *NullableWafTagAttributes {
	return &NullableWafTagAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafTagAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafTagAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
