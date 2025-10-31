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

// InlineResponse201 struct for InlineResponse201
type InlineResponse201 struct {
	// Alphanumeric string identifying the address.
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse201 InlineResponse201

// NewInlineResponse201 instantiates a new InlineResponse201 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse201() *InlineResponse201 {
	this := InlineResponse201{}
	return &this
}

// NewInlineResponse201WithDefaults instantiates a new InlineResponse201 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse201WithDefaults() *InlineResponse201 {
	this := InlineResponse201{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InlineResponse201) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse201) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InlineResponse201) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InlineResponse201) SetId(v string) {
	o.Id = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse201) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *InlineResponse201) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse201 := _InlineResponse201{}

	if err = json.Unmarshal(bytes, &varInlineResponse201); err == nil {
		*o = InlineResponse201(varInlineResponse201)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse201 is a helper abstraction for handling nullable inlineresponse201 types.
type NullableInlineResponse201 struct {
	value *InlineResponse201
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse201) Get() *InlineResponse201 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse201) Set(val *InlineResponse201) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse201) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse201) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse201 returns a pointer to a new instance of NullableInlineResponse201.
func NewNullableInlineResponse201(val *InlineResponse201) *NullableInlineResponse201 {
	return &NullableInlineResponse201{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse201) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse201) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
