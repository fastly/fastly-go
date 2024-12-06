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

// DimensionAttributesRate struct for DimensionAttributesRate
type DimensionAttributesRate struct {
	// The rate at which the value in the current dimension occurs.
	Rate                 *float32 `json:"rate,omitempty"`
	AdditionalProperties map[string]any
}

type _DimensionAttributesRate DimensionAttributesRate

// NewDimensionAttributesRate instantiates a new DimensionAttributesRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionAttributesRate() *DimensionAttributesRate {
	this := DimensionAttributesRate{}
	return &this
}

// NewDimensionAttributesRateWithDefaults instantiates a new DimensionAttributesRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionAttributesRateWithDefaults() *DimensionAttributesRate {
	this := DimensionAttributesRate{}
	return &this
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *DimensionAttributesRate) GetRate() float32 {
	if o == nil || o.Rate == nil {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionAttributesRate) GetRateOk() (*float32, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *DimensionAttributesRate) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *DimensionAttributesRate) SetRate(v float32) {
	o.Rate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DimensionAttributesRate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DimensionAttributesRate) UnmarshalJSON(bytes []byte) (err error) {
	varDimensionAttributesRate := _DimensionAttributesRate{}

	if err = json.Unmarshal(bytes, &varDimensionAttributesRate); err == nil {
		*o = DimensionAttributesRate(varDimensionAttributesRate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "rate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDimensionAttributesRate is a helper abstraction for handling nullable dimensionattributesrate types.
type NullableDimensionAttributesRate struct {
	value *DimensionAttributesRate
	isSet bool
}

// Get returns the value.
func (v NullableDimensionAttributesRate) Get() *DimensionAttributesRate {
	return v.value
}

// Set modifies the value.
func (v *NullableDimensionAttributesRate) Set(val *DimensionAttributesRate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDimensionAttributesRate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDimensionAttributesRate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDimensionAttributesRate returns a pointer to a new instance of NullableDimensionAttributesRate.
func NewNullableDimensionAttributesRate(val *DimensionAttributesRate) *NullableDimensionAttributesRate {
	return &NullableDimensionAttributesRate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDimensionAttributesRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDimensionAttributesRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
