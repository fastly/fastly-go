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

// TimeseriesResult struct for TimeseriesResult
type TimeseriesResult struct {
	// An object containing each requested dimension and time as properties.
	Dimensions *map[string]map[string]interface{} `json:"dimensions,omitempty"`
	// An object containing each requested series as a property.
	Values               map[string]map[string]interface{} `json:"values,omitempty"`
	AdditionalProperties map[string]any
}

type _TimeseriesResult TimeseriesResult

// NewTimeseriesResult instantiates a new TimeseriesResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeseriesResult() *TimeseriesResult {
	this := TimeseriesResult{}
	return &this
}

// NewTimeseriesResultWithDefaults instantiates a new TimeseriesResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeseriesResultWithDefaults() *TimeseriesResult {
	this := TimeseriesResult{}
	return &this
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *TimeseriesResult) GetDimensions() map[string]map[string]interface{} {
	if o == nil || o.Dimensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesResult) GetDimensionsOk() (*map[string]map[string]interface{}, bool) {
	if o == nil || o.Dimensions == nil {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *TimeseriesResult) HasDimensions() bool {
	if o != nil && o.Dimensions != nil {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Dimensions field.
func (o *TimeseriesResult) SetDimensions(v map[string]map[string]interface{}) {
	o.Dimensions = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *TimeseriesResult) GetValues() map[string]map[string]interface{} {
	if o == nil || o.Values == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesResult) GetValuesOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *TimeseriesResult) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given map[string]map[string]interface{} and assigns it to the Values field.
func (o *TimeseriesResult) SetValues(v map[string]map[string]interface{}) {
	o.Values = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TimeseriesResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Dimensions != nil {
		toSerialize["dimensions"] = o.Dimensions
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TimeseriesResult) UnmarshalJSON(bytes []byte) (err error) {
	varTimeseriesResult := _TimeseriesResult{}

	if err = json.Unmarshal(bytes, &varTimeseriesResult); err == nil {
		*o = TimeseriesResult(varTimeseriesResult)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dimensions")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTimeseriesResult is a helper abstraction for handling nullable timeseriesresult types.
type NullableTimeseriesResult struct {
	value *TimeseriesResult
	isSet bool
}

// Get returns the value.
func (v NullableTimeseriesResult) Get() *TimeseriesResult {
	return v.value
}

// Set modifies the value.
func (v *NullableTimeseriesResult) Set(val *TimeseriesResult) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTimeseriesResult) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTimeseriesResult) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTimeseriesResult returns a pointer to a new instance of NullableTimeseriesResult.
func NewNullableTimeseriesResult(val *TimeseriesResult) *NullableTimeseriesResult {
	return &NullableTimeseriesResult{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTimeseriesResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTimeseriesResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
