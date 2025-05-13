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

// DdosProtectionAttributeValue struct for DdosProtectionAttributeValue
type DdosProtectionAttributeValue struct {
	Value *string `json:"value,omitempty"`
	// Percentage of traffic containing a value.
	Percentage           *int32 `json:"percentage,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionAttributeValue DdosProtectionAttributeValue

// NewDdosProtectionAttributeValue instantiates a new DdosProtectionAttributeValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionAttributeValue() *DdosProtectionAttributeValue {
	this := DdosProtectionAttributeValue{}
	return &this
}

// NewDdosProtectionAttributeValueWithDefaults instantiates a new DdosProtectionAttributeValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionAttributeValueWithDefaults() *DdosProtectionAttributeValue {
	this := DdosProtectionAttributeValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DdosProtectionAttributeValue) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionAttributeValue) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DdosProtectionAttributeValue) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *DdosProtectionAttributeValue) SetValue(v string) {
	o.Value = &v
}

// GetPercentage returns the Percentage field value if set, zero value otherwise.
func (o *DdosProtectionAttributeValue) GetPercentage() int32 {
	if o == nil || o.Percentage == nil {
		var ret int32
		return ret
	}
	return *o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionAttributeValue) GetPercentageOk() (*int32, bool) {
	if o == nil || o.Percentage == nil {
		return nil, false
	}
	return o.Percentage, true
}

// HasPercentage returns a boolean if a field has been set.
func (o *DdosProtectionAttributeValue) HasPercentage() bool {
	if o != nil && o.Percentage != nil {
		return true
	}

	return false
}

// SetPercentage gets a reference to the given int32 and assigns it to the Percentage field.
func (o *DdosProtectionAttributeValue) SetPercentage(v int32) {
	o.Percentage = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionAttributeValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Percentage != nil {
		toSerialize["percentage"] = o.Percentage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionAttributeValue) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionAttributeValue := _DdosProtectionAttributeValue{}

	if err = json.Unmarshal(bytes, &varDdosProtectionAttributeValue); err == nil {
		*o = DdosProtectionAttributeValue(varDdosProtectionAttributeValue)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "percentage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionAttributeValue is a helper abstraction for handling nullable ddosprotectionattributevalue types.
type NullableDdosProtectionAttributeValue struct {
	value *DdosProtectionAttributeValue
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionAttributeValue) Get() *DdosProtectionAttributeValue {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionAttributeValue) Set(val *DdosProtectionAttributeValue) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionAttributeValue) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionAttributeValue) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionAttributeValue returns a pointer to a new instance of NullableDdosProtectionAttributeValue.
func NewNullableDdosProtectionAttributeValue(val *DdosProtectionAttributeValue) *NullableDdosProtectionAttributeValue {
	return &NullableDdosProtectionAttributeValue{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionAttributeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionAttributeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
