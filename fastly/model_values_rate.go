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

// ValuesRate struct for ValuesRate
type ValuesRate struct {
	// The percentage of requests matching the value in the current dimension.
	Rate                 *float32 `json:"rate,omitempty"`
	AdditionalProperties map[string]any
}

type _ValuesRate ValuesRate

// NewValuesRate instantiates a new ValuesRate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuesRate() *ValuesRate {
	this := ValuesRate{}
	return &this
}

// NewValuesRateWithDefaults instantiates a new ValuesRate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuesRateWithDefaults() *ValuesRate {
	this := ValuesRate{}
	return &this
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *ValuesRate) GetRate() float32 {
	if o == nil || o.Rate == nil {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesRate) GetRateOk() (*float32, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *ValuesRate) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *ValuesRate) SetRate(v float32) {
	o.Rate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValuesRate) MarshalJSON() ([]byte, error) {
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
func (o *ValuesRate) UnmarshalJSON(bytes []byte) (err error) {
	varValuesRate := _ValuesRate{}

	if err = json.Unmarshal(bytes, &varValuesRate); err == nil {
		*o = ValuesRate(varValuesRate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "rate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValuesRate is a helper abstraction for handling nullable valuesrate types.
type NullableValuesRate struct {
	value *ValuesRate
	isSet bool
}

// Get returns the value.
func (v NullableValuesRate) Get() *ValuesRate {
	return v.value
}

// Set modifies the value.
func (v *NullableValuesRate) Set(val *ValuesRate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValuesRate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValuesRate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValuesRate returns a pointer to a new instance of NullableValuesRate.
func NewNullableValuesRate(val *ValuesRate) *NullableValuesRate {
	return &NullableValuesRate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValuesRate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableValuesRate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
