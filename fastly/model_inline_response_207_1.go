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

// InlineResponse2071 struct for InlineResponse2071
type InlineResponse2071 struct {
	// Results for each operation in the request.
	Data                 []BulkOperationResult `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _InlineResponse2071 InlineResponse2071

// NewInlineResponse2071 instantiates a new InlineResponse2071 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2071() *InlineResponse2071 {
	this := InlineResponse2071{}
	return &this
}

// NewInlineResponse2071WithDefaults instantiates a new InlineResponse2071 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2071WithDefaults() *InlineResponse2071 {
	this := InlineResponse2071{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InlineResponse2071) GetData() []BulkOperationResult {
	if o == nil || o.Data == nil {
		var ret []BulkOperationResult
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse2071) GetDataOk() ([]BulkOperationResult, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InlineResponse2071) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []BulkOperationResult and assigns it to the Data field.
func (o *InlineResponse2071) SetData(v []BulkOperationResult) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InlineResponse2071) MarshalJSON() ([]byte, error) {
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
func (o *InlineResponse2071) UnmarshalJSON(bytes []byte) (err error) {
	varInlineResponse2071 := _InlineResponse2071{}

	if err = json.Unmarshal(bytes, &varInlineResponse2071); err == nil {
		*o = InlineResponse2071(varInlineResponse2071)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInlineResponse2071 is a helper abstraction for handling nullable inlineresponse2071 types.
type NullableInlineResponse2071 struct {
	value *InlineResponse2071
	isSet bool
}

// Get returns the value.
func (v NullableInlineResponse2071) Get() *InlineResponse2071 {
	return v.value
}

// Set modifies the value.
func (v *NullableInlineResponse2071) Set(val *InlineResponse2071) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInlineResponse2071) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInlineResponse2071) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInlineResponse2071 returns a pointer to a new instance of NullableInlineResponse2071.
func NewNullableInlineResponse2071(val *InlineResponse2071) *NullableInlineResponse2071 {
	return &NullableInlineResponse2071{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInlineResponse2071) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInlineResponse2071) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
