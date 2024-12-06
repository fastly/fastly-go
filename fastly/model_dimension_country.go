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

// DimensionCountry struct for DimensionCountry
type DimensionCountry struct {
	// The client's country for this dimension.
	Country              *string `json:"country,omitempty"`
	AdditionalProperties map[string]any
}

type _DimensionCountry DimensionCountry

// NewDimensionCountry instantiates a new DimensionCountry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionCountry() *DimensionCountry {
	this := DimensionCountry{}
	return &this
}

// NewDimensionCountryWithDefaults instantiates a new DimensionCountry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionCountryWithDefaults() *DimensionCountry {
	this := DimensionCountry{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *DimensionCountry) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionCountry) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *DimensionCountry) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *DimensionCountry) SetCountry(v string) {
	o.Country = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DimensionCountry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DimensionCountry) UnmarshalJSON(bytes []byte) (err error) {
	varDimensionCountry := _DimensionCountry{}

	if err = json.Unmarshal(bytes, &varDimensionCountry); err == nil {
		*o = DimensionCountry(varDimensionCountry)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDimensionCountry is a helper abstraction for handling nullable dimensioncountry types.
type NullableDimensionCountry struct {
	value *DimensionCountry
	isSet bool
}

// Get returns the value.
func (v NullableDimensionCountry) Get() *DimensionCountry {
	return v.value
}

// Set modifies the value.
func (v *NullableDimensionCountry) Set(val *DimensionCountry) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDimensionCountry) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDimensionCountry) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDimensionCountry returns a pointer to a new instance of NullableDimensionCountry.
func NewNullableDimensionCountry(val *DimensionCountry) *NullableDimensionCountry {
	return &NullableDimensionCountry{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDimensionCountry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDimensionCountry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
