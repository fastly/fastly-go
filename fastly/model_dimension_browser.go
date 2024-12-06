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

// DimensionBrowser struct for DimensionBrowser
type DimensionBrowser struct {
	// The client's browser for this dimension.
	Browser              *string `json:"browser,omitempty"`
	AdditionalProperties map[string]any
}

type _DimensionBrowser DimensionBrowser

// NewDimensionBrowser instantiates a new DimensionBrowser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionBrowser() *DimensionBrowser {
	this := DimensionBrowser{}
	return &this
}

// NewDimensionBrowserWithDefaults instantiates a new DimensionBrowser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionBrowserWithDefaults() *DimensionBrowser {
	this := DimensionBrowser{}
	return &this
}

// GetBrowser returns the Browser field value if set, zero value otherwise.
func (o *DimensionBrowser) GetBrowser() string {
	if o == nil || o.Browser == nil {
		var ret string
		return ret
	}
	return *o.Browser
}

// GetBrowserOk returns a tuple with the Browser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionBrowser) GetBrowserOk() (*string, bool) {
	if o == nil || o.Browser == nil {
		return nil, false
	}
	return o.Browser, true
}

// HasBrowser returns a boolean if a field has been set.
func (o *DimensionBrowser) HasBrowser() bool {
	if o != nil && o.Browser != nil {
		return true
	}

	return false
}

// SetBrowser gets a reference to the given string and assigns it to the Browser field.
func (o *DimensionBrowser) SetBrowser(v string) {
	o.Browser = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DimensionBrowser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Browser != nil {
		toSerialize["browser"] = o.Browser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DimensionBrowser) UnmarshalJSON(bytes []byte) (err error) {
	varDimensionBrowser := _DimensionBrowser{}

	if err = json.Unmarshal(bytes, &varDimensionBrowser); err == nil {
		*o = DimensionBrowser(varDimensionBrowser)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "browser")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDimensionBrowser is a helper abstraction for handling nullable dimensionbrowser types.
type NullableDimensionBrowser struct {
	value *DimensionBrowser
	isSet bool
}

// Get returns the value.
func (v NullableDimensionBrowser) Get() *DimensionBrowser {
	return v.value
}

// Set modifies the value.
func (v *NullableDimensionBrowser) Set(val *DimensionBrowser) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDimensionBrowser) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDimensionBrowser) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDimensionBrowser returns a pointer to a new instance of NullableDimensionBrowser.
func NewNullableDimensionBrowser(val *DimensionBrowser) *NullableDimensionBrowser {
	return &NullableDimensionBrowser{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDimensionBrowser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDimensionBrowser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}