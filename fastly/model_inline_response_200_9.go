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

// InlineResponse2009 struct for InlineResponse2009
type InlineResponse2009 struct {
	Data                 []ServiceAuthorizationResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse2009 InlineResponse2009

// NewInlineResponse2009 instantiates a new InlineResponse2009 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2009() *InlineResponse2009 {
	this := InlineResponse2009{}
	return &this
}

// NewInlineResponse2009WithDefaults instantiates a new InlineResponse2009 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2009WithDefaults() *InlineResponse2009 {
	this := InlineResponse2009{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse2009) GetData() []ServiceAuthorizationResponseData {
	if o == nil || o.Data == nil {
		var ret []ServiceAuthorizationResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2009) GetDataOk() ([]ServiceAuthorizationResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse2009) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ServiceAuthorizationResponseData and assigns it to the Data field.
func (o *InlineResponse2009) SetData(v []ServiceAuthorizationResponseData) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse2009) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *InlineResponse2009) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2009 := _InlineResponse2009{}

	if err = json.Unmarshal(bytes, &varInlineResponse2009); err == nil {
		*o = InlineResponse2009(varInlineResponse2009)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse2009 is a helper abstraction for handling nullable inlineresponse2009 types.
type NullableInlineResponse2009 struct {
	value *InlineResponse2009
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse2009) Get() *InlineResponse2009 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse2009) Set(val *InlineResponse2009) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse2009) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse2009) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse2009 returns a pointer to a new instance of NullableInlineResponse2009.
func NewNullableInlineResponse2009(val *InlineResponse2009) *NullableInlineResponse2009 {
	return &NullableInlineResponse2009{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse2009) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse2009) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
