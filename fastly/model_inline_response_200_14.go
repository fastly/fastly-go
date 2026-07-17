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

// InlineResponse20014 struct for InlineResponse20014
type InlineResponse20014 struct {
	Data []SuccessfulResponseAsObject `json:"data,omitempty"`
	// Meta for the pagination.
	Meta                 interface{} `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse20014 InlineResponse20014

// NewInlineResponse20014 instantiates a new InlineResponse20014 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse20014() *InlineResponse20014 {
	this := InlineResponse20014{}
	return &this
}

// NewInlineResponse20014WithDefaults instantiates a new InlineResponse20014 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse20014WithDefaults() *InlineResponse20014 {
	this := InlineResponse20014{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse20014) GetData() []SuccessfulResponseAsObject {
	if o == nil || o.Data == nil {
		var ret []SuccessfulResponseAsObject
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse20014) GetDataOk() ([]SuccessfulResponseAsObject, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse20014) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []SuccessfulResponseAsObject and assigns it to the Data field.
func (o *InlineResponse20014) SetData(v []SuccessfulResponseAsObject) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineResponse20014) GetMeta() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineResponse20014) GetMetaOk() (*interface{}, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return &o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse20014) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given interface{} and assigns it to the Meta field.
func (o *InlineResponse20014) SetMeta(v interface{}) {
	o.Meta = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse20014) MarshalJSON() ([]byte, error) {
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
func (o *InlineResponse20014) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse20014 := _InlineResponse20014{}

	if err = json.Unmarshal(bytes, &varInlineResponse20014); err == nil {
		*o = InlineResponse20014(varInlineResponse20014)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse20014 is a helper abstraction for handling nullable inlineresponse20014 types.
type NullableInlineResponse20014 struct {
	value *InlineResponse20014
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse20014) Get() *InlineResponse20014 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse20014) Set(val *InlineResponse20014) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse20014) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse20014) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse20014 returns a pointer to a new instance of NullableInlineResponse20014.
func NewNullableInlineResponse20014(val *InlineResponse20014) *NullableInlineResponse20014 {
	return &NullableInlineResponse20014{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse20014) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse20014) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
