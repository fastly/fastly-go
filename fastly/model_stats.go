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

// Stats struct for Stats
type Stats struct {
	Stats *map[string]Results `json:"stats,omitempty"`
	AdditionalProperties map[string]any
}

type _Stats Stats

// NewStats instantiates a new Stats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStats() *Stats {
	this := Stats{}
	return &this
}

// NewStatsWithDefaults instantiates a new Stats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsWithDefaults() *Stats {
	this := Stats{}
	return &this
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *Stats) GetStats() map[string]Results {
	if o == nil || o.Stats == nil {
		var ret map[string]Results
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Stats) GetStatsOk() (*map[string]Results, bool) {
	if o == nil || o.Stats == nil {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *Stats) HasStats() bool {
	if o != nil && o.Stats != nil {
		return true
	}

	return false
}

// SetStats gets a reference to the given map[string]Results and assigns it to the Stats field.
func (o *Stats) SetStats(v map[string]Results) {
	o.Stats = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Stats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Stats != nil {
		toSerialize["stats"] = o.Stats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *Stats) UnmarshalJSON(bytes []byte) (err error) {
	varStats := _Stats{}

	if err = json.Unmarshal(bytes, &varStats); err == nil {
		*o = Stats(varStats)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "stats")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableStats is a helper abstraction for handling nullable stats types. 
type NullableStats struct {
	value *Stats
	isSet bool
}

// Get returns the value.
func (v NullableStats) Get() *Stats {
	return v.value
}

// Set modifies the value.
func (v *NullableStats) Set(val *Stats) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableStats) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableStats) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableStats returns a pointer to a new instance of NullableStats.
func NewNullableStats(val *Stats) *NullableStats {
	return &NullableStats{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
