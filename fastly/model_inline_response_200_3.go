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

// InlineResponse2003 struct for InlineResponse2003
type InlineResponse2003 struct {
	Data                 []DdosProtectionRuleWithStats `json:"data"`
	Meta                 PaginationCursorMeta          `json:"meta"`
	AdditionalProperties map[string]any
}

type _InlineResponse2003 InlineResponse2003

// NewInlineResponse2003 instantiates a new InlineResponse2003 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003(data []DdosProtectionRuleWithStats, meta PaginationCursorMeta) *InlineResponse2003 {
	this := InlineResponse2003{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003WithDefaults() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// GetData returns the Data field value
func (o *InlineResponse2003) GetData() []DdosProtectionRuleWithStats {
	if o == nil {
		var ret []DdosProtectionRuleWithStats
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetDataOk() ([]DdosProtectionRuleWithStats, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *InlineResponse2003) SetData(v []DdosProtectionRuleWithStats) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *InlineResponse2003) GetMeta() PaginationCursorMeta {
	if o == nil {
		var ret PaginationCursorMeta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetMetaOk() (*PaginationCursorMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *InlineResponse2003) SetMeta(v PaginationCursorMeta) {
	o.Meta = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse2003) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *InlineResponse2003) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2003 := _InlineResponse2003{}

	if err = json.Unmarshal(bytes, &varInlineResponse2003); err == nil {
		*o = InlineResponse2003(varInlineResponse2003)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse2003 is a helper abstraction for handling nullable inlineresponse2003 types.
type NullableInlineResponse2003 struct {
	value *InlineResponse2003
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse2003) Get() *InlineResponse2003 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse2003) Set(val *InlineResponse2003) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse2003) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse2003) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse2003 returns a pointer to a new instance of NullableInlineResponse2003.
func NewNullableInlineResponse2003(val *InlineResponse2003) *NullableInlineResponse2003 {
	return &NullableInlineResponse2003{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse2003) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse2003) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
