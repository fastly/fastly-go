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

// InlineResponse2004 struct for InlineResponse2004
type InlineResponse2004 struct {
	Data                 []string                `json:"data,omitempty"`
	Meta                 *InlineResponse2004Meta `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse2004 InlineResponse2004

// NewInlineResponse2004 instantiates a new InlineResponse2004 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2004() *InlineResponse2004 {
	this := InlineResponse2004{}
	return &this
}

// NewInlineResponse2004WithDefaults instantiates a new InlineResponse2004 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2004WithDefaults() *InlineResponse2004 {
	this := InlineResponse2004{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse2004) GetData() []string {
	if o == nil || o.Data == nil {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetDataOk() ([]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse2004) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *InlineResponse2004) SetData(v []string) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse2004) GetMeta() InlineResponse2004Meta {
	if o == nil || o.Meta == nil {
		var ret InlineResponse2004Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2004) GetMetaOk() (*InlineResponse2004Meta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse2004) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse2004Meta and assigns it to the Meta field.
func (o *InlineResponse2004) SetMeta(v InlineResponse2004Meta) {
	o.Meta = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse2004) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *InlineResponse2004) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2004 := _InlineResponse2004{}

	if err = json.Unmarshal(bytes, &varInlineResponse2004); err == nil {
		*o = InlineResponse2004(varInlineResponse2004)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse2004 is a helper abstraction for handling nullable inlineresponse2004 types.
type NullableInlineResponse2004 struct {
	value *InlineResponse2004
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse2004) Get() *InlineResponse2004 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse2004) Set(val *InlineResponse2004) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse2004) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse2004) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse2004 returns a pointer to a new instance of NullableInlineResponse2004.
func NewNullableInlineResponse2004(val *InlineResponse2004) *NullableInlineResponse2004 {
	return &NullableInlineResponse2004{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse2004) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse2004) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
