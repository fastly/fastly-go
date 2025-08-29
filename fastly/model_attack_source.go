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

// AttackSource struct for AttackSource
type AttackSource struct {
	// Country code of the attack source
	CountryCode string `json:"country_code"`
	// Name of the country
	CountryName string `json:"country_name"`
	// Number of requests from this country
	RequestCount int32 `json:"request_count"`
	// Total number of attacks considered
	TotalCount           int32 `json:"total_count"`
	AdditionalProperties map[string]any
}

type _AttackSource AttackSource

// NewAttackSource instantiates a new AttackSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackSource(countryCode string, countryName string, requestCount int32, totalCount int32) *AttackSource {
	this := AttackSource{}
	this.CountryCode = countryCode
	this.CountryName = countryName
	this.RequestCount = requestCount
	this.TotalCount = totalCount
	return &this
}

// NewAttackSourceWithDefaults instantiates a new AttackSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackSourceWithDefaults() *AttackSource {
	this := AttackSource{}
	return &this
}

// GetCountryCode returns the CountryCode field value
func (o *AttackSource) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *AttackSource) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *AttackSource) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetCountryName returns the CountryName field value
func (o *AttackSource) GetCountryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryName
}

// GetCountryNameOk returns a tuple with the CountryName field value
// and a boolean to check if the value has been set.
func (o *AttackSource) GetCountryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryName, true
}

// SetCountryName sets field value
func (o *AttackSource) SetCountryName(v string) {
	o.CountryName = v
}

// GetRequestCount returns the RequestCount field value
func (o *AttackSource) GetRequestCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RequestCount
}

// GetRequestCountOk returns a tuple with the RequestCount field value
// and a boolean to check if the value has been set.
func (o *AttackSource) GetRequestCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestCount, true
}

// SetRequestCount sets field value
func (o *AttackSource) SetRequestCount(v int32) {
	o.RequestCount = v
}

// GetTotalCount returns the TotalCount field value
func (o *AttackSource) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *AttackSource) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *AttackSource) SetTotalCount(v int32) {
	o.TotalCount = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AttackSource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["country_code"] = o.CountryCode
	}
	if true {
		toSerialize["country_name"] = o.CountryName
	}
	if true {
		toSerialize["request_count"] = o.RequestCount
	}
	if true {
		toSerialize["total_count"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AttackSource) UnmarshalJSON(bytes []byte) (err error) {
	varAttackSource := _AttackSource{}

	if err = json.Unmarshal(bytes, &varAttackSource); err == nil {
		*o = AttackSource(varAttackSource)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "country_code")
		delete(additionalProperties, "country_name")
		delete(additionalProperties, "request_count")
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAttackSource is a helper abstraction for handling nullable attacksource types.
type NullableAttackSource struct {
	value *AttackSource
	isSet bool
}

// Get returns the value.
func (v NullableAttackSource) Get() *AttackSource {
	return v.value
}

// Set modifies the value.
func (v *NullableAttackSource) Set(val *AttackSource) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAttackSource) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAttackSource) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAttackSource returns a pointer to a new instance of NullableAttackSource.
func NewNullableAttackSource(val *AttackSource) *NullableAttackSource {
	return &NullableAttackSource{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAttackSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAttackSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
