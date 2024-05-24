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

// SnippetAllOf struct for SnippetAllOf
type SnippetAllOf struct {
	// Sets the snippet version.
	Dynamic *string `json:"dynamic,omitempty"`
	AdditionalProperties map[string]any
}

type _SnippetAllOf SnippetAllOf

// NewSnippetAllOf instantiates a new SnippetAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnippetAllOf() *SnippetAllOf {
	this := SnippetAllOf{}
	return &this
}

// NewSnippetAllOfWithDefaults instantiates a new SnippetAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnippetAllOfWithDefaults() *SnippetAllOf {
	this := SnippetAllOf{}
	return &this
}

// GetDynamic returns the Dynamic field value if set, zero value otherwise.
func (o *SnippetAllOf) GetDynamic() string {
	if o == nil || o.Dynamic == nil {
		var ret string
		return ret
	}
	return *o.Dynamic
}

// GetDynamicOk returns a tuple with the Dynamic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnippetAllOf) GetDynamicOk() (*string, bool) {
	if o == nil || o.Dynamic == nil {
		return nil, false
	}
	return o.Dynamic, true
}

// HasDynamic returns a boolean if a field has been set.
func (o *SnippetAllOf) HasDynamic() bool {
	if o != nil && o.Dynamic != nil {
		return true
	}

	return false
}

// SetDynamic gets a reference to the given string and assigns it to the Dynamic field.
func (o *SnippetAllOf) SetDynamic(v string) {
	o.Dynamic = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SnippetAllOf) MarshalJSON() ([]byte, error) {
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
func (o *SnippetAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varSnippetAllOf := _SnippetAllOf{}

	if err = json.Unmarshal(bytes, &varSnippetAllOf); err == nil {
		*o = SnippetAllOf(varSnippetAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dynamic")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSnippetAllOf is a helper abstraction for handling nullable snippetallof types. 
type NullableSnippetAllOf struct {
	value *SnippetAllOf
	isSet bool
}

// Get returns the value.
func (v NullableSnippetAllOf) Get() *SnippetAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableSnippetAllOf) Set(val *SnippetAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSnippetAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSnippetAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSnippetAllOf returns a pointer to a new instance of NullableSnippetAllOf.
func NewNullableSnippetAllOf(val *SnippetAllOf) *NullableSnippetAllOf {
	return &NullableSnippetAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSnippetAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableSnippetAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
