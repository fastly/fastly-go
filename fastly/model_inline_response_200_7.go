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

// InlineResponse2007 struct for InlineResponse2007
type InlineResponse2007 struct {
	Data []ServiceAuthorizationResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse2007 InlineResponse2007

// NewInlineResponse2007 instantiates a new InlineResponse2007 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2007() *InlineResponse2007 {
	this := InlineResponse2007{}
	return &this
}

// NewInlineResponse2007WithDefaults instantiates a new InlineResponse2007 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2007WithDefaults() *InlineResponse2007 {
	this := InlineResponse2007{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse2007) GetData() []ServiceAuthorizationResponseData {
	if o == nil || o.Data == nil {
		var ret []ServiceAuthorizationResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2007) GetDataOk() ([]ServiceAuthorizationResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse2007) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ServiceAuthorizationResponseData and assigns it to the Data field.
func (o *InlineResponse2007) SetData(v []ServiceAuthorizationResponseData) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse2007) MarshalJSON() ([]byte, error) {
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
func (o *InlineResponse2007) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2007 := _InlineResponse2007{}

	if err = json.Unmarshal(bytes, &varInlineResponse2007); err == nil {
		*o = InlineResponse2007(varInlineResponse2007)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse2007 is a helper abstraction for handling nullable inlineresponse2007 types. 
type NullableInlineResponse2007 struct {
	value *InlineResponse2007
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse2007) Get() *InlineResponse2007 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse2007) Set(val *InlineResponse2007) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse2007) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse2007) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse2007 returns a pointer to a new instance of NullableInlineResponse2007.
func NewNullableInlineResponse2007(val *InlineResponse2007) *NullableInlineResponse2007 {
	return &NullableInlineResponse2007{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse2007) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableInlineResponse2007) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
