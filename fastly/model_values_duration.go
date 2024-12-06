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

// ValuesDuration struct for ValuesDuration
type ValuesDuration struct {
	// The average time in seconds to respond to requests to the URL in the current dimension.
	AverageResponseTime *float32 `json:"average_response_time,omitempty"`
	// The P95 time in seconds to respond to requests to the URL in the current dimension.
	P95ResponseTime *float32 `json:"p95_response_time,omitempty"`
	// The total percentage of time to respond to all requests to the URL in the current dimension.
	ResponseTimePercentage *float32 `json:"response_time_percentage,omitempty"`
	AdditionalProperties   map[string]any
}

type _ValuesDuration ValuesDuration

// NewValuesDuration instantiates a new ValuesDuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuesDuration() *ValuesDuration {
	this := ValuesDuration{}
	return &this
}

// NewValuesDurationWithDefaults instantiates a new ValuesDuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuesDurationWithDefaults() *ValuesDuration {
	this := ValuesDuration{}
	return &this
}

// GetAverageResponseTime returns the AverageResponseTime field value if set, zero value otherwise.
func (o *ValuesDuration) GetAverageResponseTime() float32 {
	if o == nil || o.AverageResponseTime == nil {
		var ret float32
		return ret
	}
	return *o.AverageResponseTime
}

// GetAverageResponseTimeOk returns a tuple with the AverageResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesDuration) GetAverageResponseTimeOk() (*float32, bool) {
	if o == nil || o.AverageResponseTime == nil {
		return nil, false
	}
	return o.AverageResponseTime, true
}

// HasAverageResponseTime returns a boolean if a field has been set.
func (o *ValuesDuration) HasAverageResponseTime() bool {
	if o != nil && o.AverageResponseTime != nil {
		return true
	}

	return false
}

// SetAverageResponseTime gets a reference to the given float32 and assigns it to the AverageResponseTime field.
func (o *ValuesDuration) SetAverageResponseTime(v float32) {
	o.AverageResponseTime = &v
}

// GetP95ResponseTime returns the P95ResponseTime field value if set, zero value otherwise.
func (o *ValuesDuration) GetP95ResponseTime() float32 {
	if o == nil || o.P95ResponseTime == nil {
		var ret float32
		return ret
	}
	return *o.P95ResponseTime
}

// GetP95ResponseTimeOk returns a tuple with the P95ResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesDuration) GetP95ResponseTimeOk() (*float32, bool) {
	if o == nil || o.P95ResponseTime == nil {
		return nil, false
	}
	return o.P95ResponseTime, true
}

// HasP95ResponseTime returns a boolean if a field has been set.
func (o *ValuesDuration) HasP95ResponseTime() bool {
	if o != nil && o.P95ResponseTime != nil {
		return true
	}

	return false
}

// SetP95ResponseTime gets a reference to the given float32 and assigns it to the P95ResponseTime field.
func (o *ValuesDuration) SetP95ResponseTime(v float32) {
	o.P95ResponseTime = &v
}

// GetResponseTimePercentage returns the ResponseTimePercentage field value if set, zero value otherwise.
func (o *ValuesDuration) GetResponseTimePercentage() float32 {
	if o == nil || o.ResponseTimePercentage == nil {
		var ret float32
		return ret
	}
	return *o.ResponseTimePercentage
}

// GetResponseTimePercentageOk returns a tuple with the ResponseTimePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesDuration) GetResponseTimePercentageOk() (*float32, bool) {
	if o == nil || o.ResponseTimePercentage == nil {
		return nil, false
	}
	return o.ResponseTimePercentage, true
}

// HasResponseTimePercentage returns a boolean if a field has been set.
func (o *ValuesDuration) HasResponseTimePercentage() bool {
	if o != nil && o.ResponseTimePercentage != nil {
		return true
	}

	return false
}

// SetResponseTimePercentage gets a reference to the given float32 and assigns it to the ResponseTimePercentage field.
func (o *ValuesDuration) SetResponseTimePercentage(v float32) {
	o.ResponseTimePercentage = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValuesDuration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AverageResponseTime != nil {
		toSerialize["average_response_time"] = o.AverageResponseTime
	}
	if o.P95ResponseTime != nil {
		toSerialize["p95_response_time"] = o.P95ResponseTime
	}
	if o.ResponseTimePercentage != nil {
		toSerialize["response_time_percentage"] = o.ResponseTimePercentage
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ValuesDuration) UnmarshalJSON(bytes []byte) (err error) {
	varValuesDuration := _ValuesDuration{}

	if err = json.Unmarshal(bytes, &varValuesDuration); err == nil {
		*o = ValuesDuration(varValuesDuration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "average_response_time")
		delete(additionalProperties, "p95_response_time")
		delete(additionalProperties, "response_time_percentage")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValuesDuration is a helper abstraction for handling nullable valuesduration types.
type NullableValuesDuration struct {
	value *ValuesDuration
	isSet bool
}

// Get returns the value.
func (v NullableValuesDuration) Get() *ValuesDuration {
	return v.value
}

// Set modifies the value.
func (v *NullableValuesDuration) Set(val *ValuesDuration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValuesDuration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValuesDuration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValuesDuration returns a pointer to a new instance of NullableValuesDuration.
func NewNullableValuesDuration(val *ValuesDuration) *NullableValuesDuration {
	return &NullableValuesDuration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValuesDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableValuesDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
