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

// DimensionAttributesCountryStats struct for DimensionAttributesCountryStats
type DimensionAttributesCountryStats struct {
	// The cache hit ratio for the country.
	CountryChr *float32 `json:"country_chr,omitempty"`
	// The error rate for the country.
	CountryErrorRate *float32 `json:"country_error_rate,omitempty"`
	// This country's percentage of the total requests.
	CountryRequestRate   *float32 `json:"country_request_rate,omitempty"`
	AdditionalProperties map[string]any
}

type _DimensionAttributesCountryStats DimensionAttributesCountryStats

// NewDimensionAttributesCountryStats instantiates a new DimensionAttributesCountryStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionAttributesCountryStats() *DimensionAttributesCountryStats {
	this := DimensionAttributesCountryStats{}
	return &this
}

// NewDimensionAttributesCountryStatsWithDefaults instantiates a new DimensionAttributesCountryStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionAttributesCountryStatsWithDefaults() *DimensionAttributesCountryStats {
	this := DimensionAttributesCountryStats{}
	return &this
}

// GetCountryChr returns the CountryChr field value if set, zero value otherwise.
func (o *DimensionAttributesCountryStats) GetCountryChr() float32 {
	if o == nil || o.CountryChr == nil {
		var ret float32
		return ret
	}
	return *o.CountryChr
}

// GetCountryChrOk returns a tuple with the CountryChr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionAttributesCountryStats) GetCountryChrOk() (*float32, bool) {
	if o == nil || o.CountryChr == nil {
		return nil, false
	}
	return o.CountryChr, true
}

// HasCountryChr returns a boolean if a field has been set.
func (o *DimensionAttributesCountryStats) HasCountryChr() bool {
	if o != nil && o.CountryChr != nil {
		return true
	}

	return false
}

// SetCountryChr gets a reference to the given float32 and assigns it to the CountryChr field.
func (o *DimensionAttributesCountryStats) SetCountryChr(v float32) {
	o.CountryChr = &v
}

// GetCountryErrorRate returns the CountryErrorRate field value if set, zero value otherwise.
func (o *DimensionAttributesCountryStats) GetCountryErrorRate() float32 {
	if o == nil || o.CountryErrorRate == nil {
		var ret float32
		return ret
	}
	return *o.CountryErrorRate
}

// GetCountryErrorRateOk returns a tuple with the CountryErrorRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionAttributesCountryStats) GetCountryErrorRateOk() (*float32, bool) {
	if o == nil || o.CountryErrorRate == nil {
		return nil, false
	}
	return o.CountryErrorRate, true
}

// HasCountryErrorRate returns a boolean if a field has been set.
func (o *DimensionAttributesCountryStats) HasCountryErrorRate() bool {
	if o != nil && o.CountryErrorRate != nil {
		return true
	}

	return false
}

// SetCountryErrorRate gets a reference to the given float32 and assigns it to the CountryErrorRate field.
func (o *DimensionAttributesCountryStats) SetCountryErrorRate(v float32) {
	o.CountryErrorRate = &v
}

// GetCountryRequestRate returns the CountryRequestRate field value if set, zero value otherwise.
func (o *DimensionAttributesCountryStats) GetCountryRequestRate() float32 {
	if o == nil || o.CountryRequestRate == nil {
		var ret float32
		return ret
	}
	return *o.CountryRequestRate
}

// GetCountryRequestRateOk returns a tuple with the CountryRequestRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionAttributesCountryStats) GetCountryRequestRateOk() (*float32, bool) {
	if o == nil || o.CountryRequestRate == nil {
		return nil, false
	}
	return o.CountryRequestRate, true
}

// HasCountryRequestRate returns a boolean if a field has been set.
func (o *DimensionAttributesCountryStats) HasCountryRequestRate() bool {
	if o != nil && o.CountryRequestRate != nil {
		return true
	}

	return false
}

// SetCountryRequestRate gets a reference to the given float32 and assigns it to the CountryRequestRate field.
func (o *DimensionAttributesCountryStats) SetCountryRequestRate(v float32) {
	o.CountryRequestRate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DimensionAttributesCountryStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CountryChr != nil {
		toSerialize["country_chr"] = o.CountryChr
	}
	if o.CountryErrorRate != nil {
		toSerialize["country_error_rate"] = o.CountryErrorRate
	}
	if o.CountryRequestRate != nil {
		toSerialize["country_request_rate"] = o.CountryRequestRate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DimensionAttributesCountryStats) UnmarshalJSON(bytes []byte) (err error) {
	varDimensionAttributesCountryStats := _DimensionAttributesCountryStats{}

	if err = json.Unmarshal(bytes, &varDimensionAttributesCountryStats); err == nil {
		*o = DimensionAttributesCountryStats(varDimensionAttributesCountryStats)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "country_chr")
		delete(additionalProperties, "country_error_rate")
		delete(additionalProperties, "country_request_rate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDimensionAttributesCountryStats is a helper abstraction for handling nullable dimensionattributescountrystats types.
type NullableDimensionAttributesCountryStats struct {
	value *DimensionAttributesCountryStats
	isSet bool
}

// Get returns the value.
func (v NullableDimensionAttributesCountryStats) Get() *DimensionAttributesCountryStats {
	return v.value
}

// Set modifies the value.
func (v *NullableDimensionAttributesCountryStats) Set(val *DimensionAttributesCountryStats) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDimensionAttributesCountryStats) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDimensionAttributesCountryStats) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDimensionAttributesCountryStats returns a pointer to a new instance of NullableDimensionAttributesCountryStats.
func NewNullableDimensionAttributesCountryStats(val *DimensionAttributesCountryStats) *NullableDimensionAttributesCountryStats {
	return &NullableDimensionAttributesCountryStats{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDimensionAttributesCountryStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDimensionAttributesCountryStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
