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

// SnippetWithDynamicNumberAllOf struct for SnippetWithDynamicNumberAllOf
type SnippetWithDynamicNumberAllOf struct {
	// Sets the snippet version.
	Dynamic              *float32 `json:"dynamic,omitempty"`
	AdditionalProperties map[string]any
}

type _SnippetWithDynamicNumberAllOf SnippetWithDynamicNumberAllOf

// NewSnippetWithDynamicNumberAllOf instantiates a new SnippetWithDynamicNumberAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnippetWithDynamicNumberAllOf() *SnippetWithDynamicNumberAllOf {
	this := SnippetWithDynamicNumberAllOf{}
	return &this
}

// NewSnippetWithDynamicNumberAllOfWithDefaults instantiates a new SnippetWithDynamicNumberAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnippetWithDynamicNumberAllOfWithDefaults() *SnippetWithDynamicNumberAllOf {
	this := SnippetWithDynamicNumberAllOf{}
	return &this
}

// GetDynamic returns the Dynamic field value if set, zero value otherwise.
func (o *SnippetWithDynamicNumberAllOf) GetDynamic() float32 {
	if o == nil || o.Dynamic == nil {
		var ret float32
		return ret
	}
	return *o.Dynamic
}

// GetDynamicOk returns a tuple with the Dynamic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnippetWithDynamicNumberAllOf) GetDynamicOk() (*float32, bool) {
	if o == nil || o.Dynamic == nil {
		return nil, false
	}
	return o.Dynamic, true
}

// HasDynamic returns a boolean if a field has been set.
func (o *SnippetWithDynamicNumberAllOf) HasDynamic() bool {
	if o != nil && o.Dynamic != nil {
		return true
	}

	return false
}

// SetDynamic gets a reference to the given float32 and assigns it to the Dynamic field.
func (o *SnippetWithDynamicNumberAllOf) SetDynamic(v float32) {
	o.Dynamic = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SnippetWithDynamicNumberAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Dynamic != nil {
		toSerialize["dynamic"] = o.Dynamic
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *SnippetWithDynamicNumberAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSnippetWithDynamicNumberAllOf := _SnippetWithDynamicNumberAllOf{}

	if err = json.Unmarshal(bytes, &varSnippetWithDynamicNumberAllOf); err == nil {
		*o = SnippetWithDynamicNumberAllOf(varSnippetWithDynamicNumberAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dynamic")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSnippetWithDynamicNumberAllOf is a helper abstraction for handling nullable snippetwithdynamicnumberallof types.
type NullableSnippetWithDynamicNumberAllOf struct {
	value *SnippetWithDynamicNumberAllOf
	isSet bool
}

// Get returns the value.
func (v NullableSnippetWithDynamicNumberAllOf) Get() *SnippetWithDynamicNumberAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableSnippetWithDynamicNumberAllOf) Set(val *SnippetWithDynamicNumberAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSnippetWithDynamicNumberAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSnippetWithDynamicNumberAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSnippetWithDynamicNumberAllOf returns a pointer to a new instance of NullableSnippetWithDynamicNumberAllOf.
func NewNullableSnippetWithDynamicNumberAllOf(val *SnippetWithDynamicNumberAllOf) *NullableSnippetWithDynamicNumberAllOf {
	return &NullableSnippetWithDynamicNumberAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSnippetWithDynamicNumberAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableSnippetWithDynamicNumberAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
