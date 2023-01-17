// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// HistoricalResponseAllOf struct for HistoricalResponseAllOf
type HistoricalResponseAllOf struct {
	// Contains the results of the query, organized by *service ID*, into arrays where each element describes one service over a *time span*.
	Data *map[string][]Results `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalResponseAllOf HistoricalResponseAllOf

// NewHistoricalResponseAllOf instantiates a new HistoricalResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalResponseAllOf() *HistoricalResponseAllOf {
	this := HistoricalResponseAllOf{}
	return &this
}

// NewHistoricalResponseAllOfWithDefaults instantiates a new HistoricalResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalResponseAllOfWithDefaults() *HistoricalResponseAllOf {
	this := HistoricalResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HistoricalResponseAllOf) GetData() map[string][]Results {
	if o == nil || o.Data == nil {
		var ret map[string][]Results
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalResponseAllOf) GetDataOk() (*map[string][]Results, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HistoricalResponseAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string][]Results and assigns it to the Data field.
func (o *HistoricalResponseAllOf) SetData(v map[string][]Results) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalResponseAllOf) MarshalJSON() ([]byte, error) {
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
func (o *HistoricalResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalResponseAllOf := _HistoricalResponseAllOf{}

	if err = json.Unmarshal(bytes, &varHistoricalResponseAllOf); err == nil {
		*o = HistoricalResponseAllOf(varHistoricalResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHistoricalResponseAllOf is a helper abstraction for handling nullable historicalresponseallof types. 
type NullableHistoricalResponseAllOf struct {
	value *HistoricalResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalResponseAllOf) Get() *HistoricalResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalResponseAllOf) Set(val *HistoricalResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalResponseAllOf returns a pointer to a new instance of NullableHistoricalResponseAllOf.
func NewNullableHistoricalResponseAllOf(val *HistoricalResponseAllOf) *NullableHistoricalResponseAllOf {
	return &NullableHistoricalResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHistoricalResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
