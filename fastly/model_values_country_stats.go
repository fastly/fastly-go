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

// ValuesCountryStats struct for ValuesCountryStats
type ValuesCountryStats struct {
	// The client's country subdivision code as defined by ISO 3166-2.
	Region *string `json:"region,omitempty"`
	// The cache hit ratio for the region.
	RegionChr *float32 `json:"region_chr,omitempty"`
	// The error rate for the region.
	RegionErrorRate      *float32 `json:"region_error_rate,omitempty"`
	AdditionalProperties map[string]any
}

type _ValuesCountryStats ValuesCountryStats

// NewValuesCountryStats instantiates a new ValuesCountryStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuesCountryStats() *ValuesCountryStats {
	this := ValuesCountryStats{}
	return &this
}

// NewValuesCountryStatsWithDefaults instantiates a new ValuesCountryStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuesCountryStatsWithDefaults() *ValuesCountryStats {
	this := ValuesCountryStats{}
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *ValuesCountryStats) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesCountryStats) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *ValuesCountryStats) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *ValuesCountryStats) SetRegion(v string) {
	o.Region = &v
}

// GetRegionChr returns the RegionChr field value if set, zero value otherwise.
func (o *ValuesCountryStats) GetRegionChr() float32 {
	if o == nil || o.RegionChr == nil {
		var ret float32
		return ret
	}
	return *o.RegionChr
}

// GetRegionChrOk returns a tuple with the RegionChr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesCountryStats) GetRegionChrOk() (*float32, bool) {
	if o == nil || o.RegionChr == nil {
		return nil, false
	}
	return o.RegionChr, true
}

// HasRegionChr returns a boolean if a field has been set.
func (o *ValuesCountryStats) HasRegionChr() bool {
	if o != nil && o.RegionChr != nil {
		return true
	}

	return false
}

// SetRegionChr gets a reference to the given float32 and assigns it to the RegionChr field.
func (o *ValuesCountryStats) SetRegionChr(v float32) {
	o.RegionChr = &v
}

// GetRegionErrorRate returns the RegionErrorRate field value if set, zero value otherwise.
func (o *ValuesCountryStats) GetRegionErrorRate() float32 {
	if o == nil || o.RegionErrorRate == nil {
		var ret float32
		return ret
	}
	return *o.RegionErrorRate
}

// GetRegionErrorRateOk returns a tuple with the RegionErrorRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesCountryStats) GetRegionErrorRateOk() (*float32, bool) {
	if o == nil || o.RegionErrorRate == nil {
		return nil, false
	}
	return o.RegionErrorRate, true
}

// HasRegionErrorRate returns a boolean if a field has been set.
func (o *ValuesCountryStats) HasRegionErrorRate() bool {
	if o != nil && o.RegionErrorRate != nil {
		return true
	}

	return false
}

// SetRegionErrorRate gets a reference to the given float32 and assigns it to the RegionErrorRate field.
func (o *ValuesCountryStats) SetRegionErrorRate(v float32) {
	o.RegionErrorRate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValuesCountryStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.RegionChr != nil {
		toSerialize["region_chr"] = o.RegionChr
	}
	if o.RegionErrorRate != nil {
		toSerialize["region_error_rate"] = o.RegionErrorRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ValuesCountryStats) UnmarshalJSON(bytes []byte) (err error) {
	varValuesCountryStats := _ValuesCountryStats{}

	if err = json.Unmarshal(bytes, &varValuesCountryStats); err == nil {
		*o = ValuesCountryStats(varValuesCountryStats)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "region")
		delete(additionalProperties, "region_chr")
		delete(additionalProperties, "region_error_rate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValuesCountryStats is a helper abstraction for handling nullable valuescountrystats types.
type NullableValuesCountryStats struct {
	value *ValuesCountryStats
	isSet bool
}

// Get returns the value.
func (v NullableValuesCountryStats) Get() *ValuesCountryStats {
	return v.value
}

// Set modifies the value.
func (v *NullableValuesCountryStats) Set(val *ValuesCountryStats) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValuesCountryStats) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValuesCountryStats) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValuesCountryStats returns a pointer to a new instance of NullableValuesCountryStats.
func NewNullableValuesCountryStats(val *ValuesCountryStats) *NullableValuesCountryStats {
	return &NullableValuesCountryStats{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValuesCountryStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableValuesCountryStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
