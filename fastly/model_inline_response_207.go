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

// InlineResponse207 struct for InlineResponse207
type InlineResponse207 struct {
	// Results for each operation in the request.
	Data                 []BulkOperationResult `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse207 InlineResponse207

// NewInlineResponse207 instantiates a new InlineResponse207 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse207() *InlineResponse207 {
	this := InlineResponse207{}
	return &this
}

// NewInlineResponse207WithDefaults instantiates a new InlineResponse207 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse207WithDefaults() *InlineResponse207 {
	this := InlineResponse207{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse207) GetData() []BulkOperationResult {
	if o == nil || o.Data == nil {
		var ret []BulkOperationResult
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse207) GetDataOk() ([]BulkOperationResult, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse207) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []BulkOperationResult and assigns it to the Data field.
func (o *InlineResponse207) SetData(v []BulkOperationResult) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse207) MarshalJSON() ([]byte, error) {
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
func (o *InlineResponse207) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse207 := _InlineResponse207{}

	if err = json.Unmarshal(bytes, &varInlineResponse207); err == nil {
		*o = InlineResponse207(varInlineResponse207)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse207 is a helper abstraction for handling nullable inlineresponse207 types.
type NullableInlineResponse207 struct {
	value *InlineResponse207
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse207) Get() *InlineResponse207 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse207) Set(val *InlineResponse207) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse207) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse207) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse207 returns a pointer to a new instance of NullableInlineResponse207.
func NewNullableInlineResponse207(val *InlineResponse207) *NullableInlineResponse207 {
	return &NullableInlineResponse207{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse207) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse207) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
