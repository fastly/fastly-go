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

// ValuesBandwidth struct for ValuesBandwidth
type ValuesBandwidth struct {
	// The average bandwidth in bytes for responses to requests to the URL in the current dimension.
	AverageBandwidthBytes *float32 `json:"average_bandwidth_bytes,omitempty"`
	// The total bandwidth percentage for all responses to requests to the URL in the current dimension.
	BandwidthPercentage  *float32 `json:"bandwidth_percentage,omitempty"`
	AdditionalProperties map[string]any
}

type _ValuesBandwidth ValuesBandwidth

// NewValuesBandwidth instantiates a new ValuesBandwidth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuesBandwidth() *ValuesBandwidth {
	this := ValuesBandwidth{}
	return &this
}

// NewValuesBandwidthWithDefaults instantiates a new ValuesBandwidth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuesBandwidthWithDefaults() *ValuesBandwidth {
	this := ValuesBandwidth{}
	return &this
}

// GetAverageBandwidthBytes returns the AverageBandwidthBytes field value if set, zero value otherwise.
func (o *ValuesBandwidth) GetAverageBandwidthBytes() float32 {
	if o == nil || o.AverageBandwidthBytes == nil {
		var ret float32
		return ret
	}
	return *o.AverageBandwidthBytes
}

// GetAverageBandwidthBytesOk returns a tuple with the AverageBandwidthBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesBandwidth) GetAverageBandwidthBytesOk() (*float32, bool) {
	if o == nil || o.AverageBandwidthBytes == nil {
		return nil, false
	}
	return o.AverageBandwidthBytes, true
}

// HasAverageBandwidthBytes returns a boolean if a field has been set.
func (o *ValuesBandwidth) HasAverageBandwidthBytes() bool {
	if o != nil && o.AverageBandwidthBytes != nil {
		return true
	}

	return false
}

// SetAverageBandwidthBytes gets a reference to the given float32 and assigns it to the AverageBandwidthBytes field.
func (o *ValuesBandwidth) SetAverageBandwidthBytes(v float32) {
	o.AverageBandwidthBytes = &v
}

// GetBandwidthPercentage returns the BandwidthPercentage field value if set, zero value otherwise.
func (o *ValuesBandwidth) GetBandwidthPercentage() float32 {
	if o == nil || o.BandwidthPercentage == nil {
		var ret float32
		return ret
	}
	return *o.BandwidthPercentage
}

// GetBandwidthPercentageOk returns a tuple with the BandwidthPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesBandwidth) GetBandwidthPercentageOk() (*float32, bool) {
	if o == nil || o.BandwidthPercentage == nil {
		return nil, false
	}
	return o.BandwidthPercentage, true
}

// HasBandwidthPercentage returns a boolean if a field has been set.
func (o *ValuesBandwidth) HasBandwidthPercentage() bool {
	if o != nil && o.BandwidthPercentage != nil {
		return true
	}

	return false
}

// SetBandwidthPercentage gets a reference to the given float32 and assigns it to the BandwidthPercentage field.
func (o *ValuesBandwidth) SetBandwidthPercentage(v float32) {
	o.BandwidthPercentage = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValuesBandwidth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AverageBandwidthBytes != nil {
		toSerialize["average_bandwidth_bytes"] = o.AverageBandwidthBytes
	}
	if o.BandwidthPercentage != nil {
		toSerialize["bandwidth_percentage"] = o.BandwidthPercentage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ValuesBandwidth) UnmarshalJSON(bytes []byte) (err error) {
	varValuesBandwidth := _ValuesBandwidth{}

	if err = json.Unmarshal(bytes, &varValuesBandwidth); err == nil {
		*o = ValuesBandwidth(varValuesBandwidth)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "average_bandwidth_bytes")
		delete(additionalProperties, "bandwidth_percentage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValuesBandwidth is a helper abstraction for handling nullable valuesbandwidth types.
type NullableValuesBandwidth struct {
	value *ValuesBandwidth
	isSet bool
}

// Get returns the value.
func (v NullableValuesBandwidth) Get() *ValuesBandwidth {
	return v.value
}

// Set modifies the value.
func (v *NullableValuesBandwidth) Set(val *ValuesBandwidth) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValuesBandwidth) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValuesBandwidth) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValuesBandwidth returns a pointer to a new instance of NullableValuesBandwidth.
func NewNullableValuesBandwidth(val *ValuesBandwidth) *NullableValuesBandwidth {
	return &NullableValuesBandwidth{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValuesBandwidth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableValuesBandwidth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
