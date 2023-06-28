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

// InlineResponse2001 struct for InlineResponse2001
type InlineResponse2001 struct {
	// The service IDs of the services the token will have access to. Separate service IDs with a space.
	Data []string `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse2001 InlineResponse2001

// NewInlineResponse2001 instantiates a new InlineResponse2001 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// NewInlineResponse2001WithDefaults instantiates a new InlineResponse2001 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001WithDefaults() *InlineResponse2001 {
	this := InlineResponse2001{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse2001) GetData() []string {
	if o == nil || o.Data == nil {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2001) GetDataOk() ([]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse2001) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *InlineResponse2001) SetData(v []string) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse2001) MarshalJSON() ([]byte, error) {
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
func (o *InlineResponse2001) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2001 := _InlineResponse2001{}

	if err = json.Unmarshal(bytes, &varInlineResponse2001); err == nil {
		*o = InlineResponse2001(varInlineResponse2001)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse2001 is a helper abstraction for handling nullable inlineresponse2001 types. 
type NullableInlineResponse2001 struct {
	value *InlineResponse2001
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse2001) Get() *InlineResponse2001 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse2001) Set(val *InlineResponse2001) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse2001) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse2001) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse2001 returns a pointer to a new instance of NullableInlineResponse2001.
func NewNullableInlineResponse2001(val *InlineResponse2001) *NullableInlineResponse2001 {
	return &NullableInlineResponse2001{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse2001) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableInlineResponse2001) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
