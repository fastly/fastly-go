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

// InlineResponse2002 struct for InlineResponse2002
type InlineResponse2002 struct {
	Data []StoreResponse `json:"data,omitempty"`
	Meta *InlineResponse2002Meta `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse2002 InlineResponse2002

// NewInlineResponse2002 instantiates a new InlineResponse2002 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2002() *InlineResponse2002 {
	this := InlineResponse2002{}
	return &this
}

// NewInlineResponse2002WithDefaults instantiates a new InlineResponse2002 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2002WithDefaults() *InlineResponse2002 {
	this := InlineResponse2002{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse2002) GetData() []StoreResponse {
	if o == nil || o.Data == nil {
		var ret []StoreResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002) GetDataOk() ([]StoreResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse2002) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []StoreResponse and assigns it to the Data field.
func (o *InlineResponse2002) SetData(v []StoreResponse) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *InlineResponse2002) GetMeta() InlineResponse2002Meta {
	if o == nil || o.Meta == nil {
		var ret InlineResponse2002Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2002) GetMetaOk() (*InlineResponse2002Meta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *InlineResponse2002) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given InlineResponse2002Meta and assigns it to the Meta field.
func (o *InlineResponse2002) SetMeta(v InlineResponse2002Meta) {
	o.Meta = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse2002) MarshalJSON() ([]byte, error) {
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
func (o *InlineResponse2002) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2002 := _InlineResponse2002{}

	if err = json.Unmarshal(bytes, &varInlineResponse2002); err == nil {
		*o = InlineResponse2002(varInlineResponse2002)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse2002 is a helper abstraction for handling nullable inlineresponse2002 types. 
type NullableInlineResponse2002 struct {
	value *InlineResponse2002
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse2002) Get() *InlineResponse2002 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse2002) Set(val *InlineResponse2002) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse2002) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse2002) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse2002 returns a pointer to a new instance of NullableInlineResponse2002.
func NewNullableInlineResponse2002(val *InlineResponse2002) *NullableInlineResponse2002 {
	return &NullableInlineResponse2002{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse2002) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableInlineResponse2002) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
