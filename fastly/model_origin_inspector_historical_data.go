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

// OriginInspectorHistoricalData struct for OriginInspectorHistoricalData
type OriginInspectorHistoricalData struct {
	Dimensions *OriginInspectorDimensions `json:"dimensions,omitempty"`
	// An array of values representing the metric values at each point in time. Note that this dataset is sparse: only the keys with non-zero values will be included in the record. 
	Values []OriginInspectorValues `json:"values,omitempty"`
	AdditionalProperties map[string]any
}

type _OriginInspectorHistoricalData OriginInspectorHistoricalData

// NewOriginInspectorHistoricalData instantiates a new OriginInspectorHistoricalData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspectorHistoricalData() *OriginInspectorHistoricalData {
	this := OriginInspectorHistoricalData{}
	return &this
}

// NewOriginInspectorHistoricalDataWithDefaults instantiates a new OriginInspectorHistoricalData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorHistoricalDataWithDefaults() *OriginInspectorHistoricalData {
	this := OriginInspectorHistoricalData{}
	return &this
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalData) GetDimensions() OriginInspectorDimensions {
	if o == nil || o.Dimensions == nil {
		var ret OriginInspectorDimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalData) GetDimensionsOk() (*OriginInspectorDimensions, bool) {
	if o == nil || o.Dimensions == nil {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalData) HasDimensions() bool {
	if o != nil && o.Dimensions != nil {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given OriginInspectorDimensions and assigns it to the Dimensions field.
func (o *OriginInspectorHistoricalData) SetDimensions(v OriginInspectorDimensions) {
	o.Dimensions = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *OriginInspectorHistoricalData) GetValues() []OriginInspectorValues {
	if o == nil || o.Values == nil {
		var ret []OriginInspectorValues
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorHistoricalData) GetValuesOk() ([]OriginInspectorValues, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *OriginInspectorHistoricalData) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []OriginInspectorValues and assigns it to the Values field.
func (o *OriginInspectorHistoricalData) SetValues(v []OriginInspectorValues) {
	o.Values = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspectorHistoricalData) MarshalJSON() ([]byte, error) {
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
func (o *OriginInspectorHistoricalData) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspectorHistoricalData := _OriginInspectorHistoricalData{}

	if err = json.Unmarshal(bytes, &varOriginInspectorHistoricalData); err == nil {
		*o = OriginInspectorHistoricalData(varOriginInspectorHistoricalData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dimensions")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOriginInspectorHistoricalData is a helper abstraction for handling nullable origininspectorhistoricaldata types. 
type NullableOriginInspectorHistoricalData struct {
	value *OriginInspectorHistoricalData
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspectorHistoricalData) Get() *OriginInspectorHistoricalData {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspectorHistoricalData) Set(val *OriginInspectorHistoricalData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspectorHistoricalData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspectorHistoricalData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspectorHistoricalData returns a pointer to a new instance of NullableOriginInspectorHistoricalData.
func NewNullableOriginInspectorHistoricalData(val *OriginInspectorHistoricalData) *NullableOriginInspectorHistoricalData {
	return &NullableOriginInspectorHistoricalData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspectorHistoricalData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableOriginInspectorHistoricalData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
