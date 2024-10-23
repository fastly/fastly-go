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

// InlineResponse2005 struct for InlineResponse2005
type InlineResponse2005 struct {
	Data                 []SecretStoreResponse `json:"data,omitempty"`
	Meta                 *PaginationCursorMeta `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse2005 InlineResponse2005

// NewInlineResponse2005 instantiates a new InlineResponse2005 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2005() *InlineResponse2005 {
	this := InlineResponse2005{}
	return &this
}

// NewInlineResponse2005WithDefaults instantiates a new InlineResponse2005 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2005WithDefaults() *InlineResponse2005 {
	this := InlineResponse2005{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse2005) GetData() []SecretStoreResponse {
	if o == nil || o.Data == nil {
		var ret []SecretStoreResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005) GetDataOk() ([]SecretStoreResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse2005) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []SecretStoreResponse and assigns it to the Data field.
func (o *InlineResponse2005) SetData(v []SecretStoreResponse) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse2005) GetMeta() PaginationCursorMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationCursorMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2005) GetMetaOk() (*PaginationCursorMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse2005) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationCursorMeta and assigns it to the Meta field.
func (o *InlineResponse2005) SetMeta(v PaginationCursorMeta) {
	o.Meta = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse2005) MarshalJSON() ([]byte, error) {
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
func (o *InlineResponse2005) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2005 := _InlineResponse2005{}

	if err = json.Unmarshal(bytes, &varInlineResponse2005); err == nil {
		*o = InlineResponse2005(varInlineResponse2005)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse2005 is a helper abstraction for handling nullable inlineresponse2005 types.
type NullableInlineResponse2005 struct {
	value *InlineResponse2005
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse2005) Get() *InlineResponse2005 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse2005) Set(val *InlineResponse2005) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse2005) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse2005) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse2005 returns a pointer to a new instance of NullableInlineResponse2005.
func NewNullableInlineResponse2005(val *InlineResponse2005) *NullableInlineResponse2005 {
	return &NullableInlineResponse2005{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse2005) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse2005) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
