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

// DimensionUrl struct for DimensionUrl
type DimensionUrl struct {
	// The URL path for this dimension.
	Url                  *string `json:"url,omitempty"`
	AdditionalProperties map[string]any
}

type _DimensionUrl DimensionUrl

// NewDimensionUrl instantiates a new DimensionUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionUrl() *DimensionUrl {
	this := DimensionUrl{}
	return &this
}

// NewDimensionUrlWithDefaults instantiates a new DimensionUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionUrlWithDefaults() *DimensionUrl {
	this := DimensionUrl{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DimensionUrl) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionUrl) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DimensionUrl) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DimensionUrl) SetUrl(v string) {
	o.Url = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DimensionUrl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DimensionUrl) UnmarshalJSON(bytes []byte) (err error) {
	varDimensionUrl := _DimensionUrl{}

	if err = json.Unmarshal(bytes, &varDimensionUrl); err == nil {
		*o = DimensionUrl(varDimensionUrl)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDimensionUrl is a helper abstraction for handling nullable dimensionurl types.
type NullableDimensionUrl struct {
	value *DimensionUrl
	isSet bool
}

// Get returns the value.
func (v NullableDimensionUrl) Get() *DimensionUrl {
	return v.value
}

// Set modifies the value.
func (v *NullableDimensionUrl) Set(val *DimensionUrl) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDimensionUrl) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDimensionUrl) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDimensionUrl returns a pointer to a new instance of NullableDimensionUrl.
func NewNullableDimensionUrl(val *DimensionUrl) *NullableDimensionUrl {
	return &NullableDimensionUrl{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDimensionUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDimensionUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
