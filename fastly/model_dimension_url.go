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

// DimensionURL struct for DimensionURL
type DimensionURL struct {
	// The URL path for this dimension.
	URL                  *string `json:"url,omitempty"`
	AdditionalProperties map[string]any
}

type _DimensionURL DimensionURL

// NewDimensionURL instantiates a new DimensionURL object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionURL() *DimensionURL {
	this := DimensionURL{}
	return &this
}

// NewDimensionURLWithDefaults instantiates a new DimensionURL object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionURLWithDefaults() *DimensionURL {
	this := DimensionURL{}
	return &this
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *DimensionURL) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionURL) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *DimensionURL) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *DimensionURL) SetURL(v string) {
	o.URL = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DimensionURL) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.URL != nil {
		toSerialize["url"] = o.URL
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DimensionURL) UnmarshalJSON(bytes []byte) (err error) {
	varDimensionURL := _DimensionURL{}

	if err = json.Unmarshal(bytes, &varDimensionURL); err == nil {
		*o = DimensionURL(varDimensionURL)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDimensionURL is a helper abstraction for handling nullable dimensionurl types.
type NullableDimensionURL struct {
	value *DimensionURL
	isSet bool
}

// Get returns the value.
func (v NullableDimensionURL) Get() *DimensionURL {
	return v.value
}

// Set modifies the value.
func (v *NullableDimensionURL) Set(val *DimensionURL) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDimensionURL) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDimensionURL) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDimensionURL returns a pointer to a new instance of NullableDimensionURL.
func NewNullableDimensionURL(val *DimensionURL) *NullableDimensionURL {
	return &NullableDimensionURL{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDimensionURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDimensionURL) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
