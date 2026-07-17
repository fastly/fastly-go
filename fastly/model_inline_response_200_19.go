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

// InlineResponse20019 struct for InlineResponse20019
type InlineResponse20019 struct {
	Meta *AgentKeyMeta `json:"meta,omitempty"`
	// The agent keys returned by the request.
	Data                 []AgentKey `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse20019 InlineResponse20019

// NewInlineResponse20019 instantiates a new InlineResponse20019 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20019() *InlineResponse20019 {
	this := InlineResponse20019{}
	return &this
}

// NewInlineResponse20019WithDefaults instantiates a new InlineResponse20019 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20019WithDefaults() *InlineResponse20019 {
	this := InlineResponse20019{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse20019) GetMeta() AgentKeyMeta {
	if o == nil || o.Meta == nil {
		var ret AgentKeyMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20019) GetMetaOk() (*AgentKeyMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse20019) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given AgentKeyMeta and assigns it to the Meta field.
func (o *InlineResponse20019) SetMeta(v AgentKeyMeta) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse20019) GetData() []AgentKey {
	if o == nil || o.Data == nil {
		var ret []AgentKey
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20019) GetDataOk() ([]AgentKey, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse20019) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []AgentKey and assigns it to the Data field.
func (o *InlineResponse20019) SetData(v []AgentKey) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse20019) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *InlineResponse20019) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse20019 := _InlineResponse20019{}

	if err = json.Unmarshal(bytes, &varInlineResponse20019); err == nil {
		*o = InlineResponse20019(varInlineResponse20019)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse20019 is a helper abstraction for handling nullable inlineresponse20019 types.
type NullableInlineResponse20019 struct {
	value *InlineResponse20019
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse20019) Get() *InlineResponse20019 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse20019) Set(val *InlineResponse20019) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse20019) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse20019) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse20019 returns a pointer to a new instance of NullableInlineResponse20019.
func NewNullableInlineResponse20019(val *InlineResponse20019) *NullableInlineResponse20019 {
	return &NullableInlineResponse20019{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse20019) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse20019) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
