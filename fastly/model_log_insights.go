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

// LogInsights struct for LogInsights
type LogInsights struct {
	Dimensions           *LogInsightsDimensions          `json:"dimensions,omitempty"`
	DimensionAttributes  *LogInsightsDimensionAttributes `json:"dimension_attributes,omitempty"`
	Values               *LogInsightsValues              `json:"values,omitempty"`
	AdditionalProperties map[string]any
}

type _LogInsights LogInsights

// NewLogInsights instantiates a new LogInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogInsights() *LogInsights {
	this := LogInsights{}
	return &this
}

// NewLogInsightsWithDefaults instantiates a new LogInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogInsightsWithDefaults() *LogInsights {
	this := LogInsights{}
	return &this
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *LogInsights) GetDimensions() LogInsightsDimensions {
	if o == nil || o.Dimensions == nil {
		var ret LogInsightsDimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogInsights) GetDimensionsOk() (*LogInsightsDimensions, bool) {
	if o == nil || o.Dimensions == nil {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *LogInsights) HasDimensions() bool {
	if o != nil && o.Dimensions != nil {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given LogInsightsDimensions and assigns it to the Dimensions field.
func (o *LogInsights) SetDimensions(v LogInsightsDimensions) {
	o.Dimensions = &v
}

// GetDimensionAttributes returns the DimensionAttributes field value if set, zero value otherwise.
func (o *LogInsights) GetDimensionAttributes() LogInsightsDimensionAttributes {
	if o == nil || o.DimensionAttributes == nil {
		var ret LogInsightsDimensionAttributes
		return ret
	}
	return *o.DimensionAttributes
}

// GetDimensionAttributesOk returns a tuple with the DimensionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogInsights) GetDimensionAttributesOk() (*LogInsightsDimensionAttributes, bool) {
	if o == nil || o.DimensionAttributes == nil {
		return nil, false
	}
	return o.DimensionAttributes, true
}

// HasDimensionAttributes returns a boolean if a field has been set.
func (o *LogInsights) HasDimensionAttributes() bool {
	if o != nil && o.DimensionAttributes != nil {
		return true
	}

	return false
}

// SetDimensionAttributes gets a reference to the given LogInsightsDimensionAttributes and assigns it to the DimensionAttributes field.
func (o *LogInsights) SetDimensionAttributes(v LogInsightsDimensionAttributes) {
	o.DimensionAttributes = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *LogInsights) GetValues() LogInsightsValues {
	if o == nil || o.Values == nil {
		var ret LogInsightsValues
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogInsights) GetValuesOk() (*LogInsightsValues, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *LogInsights) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given LogInsightsValues and assigns it to the Values field.
func (o *LogInsights) SetValues(v LogInsightsValues) {
	o.Values = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogInsights) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Dimensions != nil {
		toSerialize["dimensions"] = o.Dimensions
	}
	if o.DimensionAttributes != nil {
		toSerialize["dimension_attributes"] = o.DimensionAttributes
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
func (o *LogInsights) UnmarshalJSON(bytes []byte) (err error) {
	varLogInsights := _LogInsights{}

	if err = json.Unmarshal(bytes, &varLogInsights); err == nil {
		*o = LogInsights(varLogInsights)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dimensions")
		delete(additionalProperties, "dimension_attributes")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLogInsights is a helper abstraction for handling nullable loginsights types.
type NullableLogInsights struct {
	value *LogInsights
	isSet bool
}

// Get returns the value.
func (v NullableLogInsights) Get() *LogInsights {
	return v.value
}

// Set modifies the value.
func (v *NullableLogInsights) Set(val *LogInsights) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogInsights) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogInsights) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogInsights returns a pointer to a new instance of NullableLogInsights.
func NewNullableLogInsights(val *LogInsights) *NullableLogInsights {
	return &NullableLogInsights{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
