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

// HistoricalDomainsData struct for HistoricalDomainsData
type HistoricalDomainsData struct {
	Dimensions *DomainInspectorEntryDimensions `json:"dimensions,omitempty"`
	// An array of values representing the metric values at each point in time. Note that this dataset is sparse: only the keys with non-zero values will be included in the record. 
	Values []Values `json:"values,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalDomainsData HistoricalDomainsData

// NewHistoricalDomainsData instantiates a new HistoricalDomainsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalDomainsData() *HistoricalDomainsData {
	this := HistoricalDomainsData{}
	return &this
}

// NewHistoricalDomainsDataWithDefaults instantiates a new HistoricalDomainsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalDomainsDataWithDefaults() *HistoricalDomainsData {
	this := HistoricalDomainsData{}
	return &this
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *HistoricalDomainsData) GetDimensions() DomainInspectorEntryDimensions {
	if o == nil || o.Dimensions == nil {
		var ret DomainInspectorEntryDimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDomainsData) GetDimensionsOk() (*DomainInspectorEntryDimensions, bool) {
	if o == nil || o.Dimensions == nil {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *HistoricalDomainsData) HasDimensions() bool {
	if o != nil && o.Dimensions != nil {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given DomainInspectorEntryDimensions and assigns it to the Dimensions field.
func (o *HistoricalDomainsData) SetDimensions(v DomainInspectorEntryDimensions) {
	o.Dimensions = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *HistoricalDomainsData) GetValues() []Values {
	if o == nil || o.Values == nil {
		var ret []Values
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDomainsData) GetValuesOk() ([]Values, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *HistoricalDomainsData) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []Values and assigns it to the Values field.
func (o *HistoricalDomainsData) SetValues(v []Values) {
	o.Values = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalDomainsData) MarshalJSON() ([]byte, error) {
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
func (o *HistoricalDomainsData) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalDomainsData := _HistoricalDomainsData{}

	if err = json.Unmarshal(bytes, &varHistoricalDomainsData); err == nil {
		*o = HistoricalDomainsData(varHistoricalDomainsData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dimensions")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHistoricalDomainsData is a helper abstraction for handling nullable historicaldomainsdata types. 
type NullableHistoricalDomainsData struct {
	value *HistoricalDomainsData
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalDomainsData) Get() *HistoricalDomainsData {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalDomainsData) Set(val *HistoricalDomainsData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalDomainsData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalDomainsData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalDomainsData returns a pointer to a new instance of NullableHistoricalDomainsData.
func NewNullableHistoricalDomainsData(val *HistoricalDomainsData) *NullableHistoricalDomainsData {
	return &NullableHistoricalDomainsData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalDomainsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHistoricalDomainsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
