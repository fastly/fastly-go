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

// ValuesBrowser struct for ValuesBrowser
type ValuesBrowser struct {
	// The version of the client's browser.
	BrowserVersion *string `json:"browser_version,omitempty"`
	// The percentage of requests by this version of the browser specified in the dimension.
	Rate                 *float32 `json:"rate,omitempty"`
	AdditionalProperties map[string]any
}

type _ValuesBrowser ValuesBrowser

// NewValuesBrowser instantiates a new ValuesBrowser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuesBrowser() *ValuesBrowser {
	this := ValuesBrowser{}
	return &this
}

// NewValuesBrowserWithDefaults instantiates a new ValuesBrowser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuesBrowserWithDefaults() *ValuesBrowser {
	this := ValuesBrowser{}
	return &this
}

// GetBrowserVersion returns the BrowserVersion field value if set, zero value otherwise.
func (o *ValuesBrowser) GetBrowserVersion() string {
	if o == nil || o.BrowserVersion == nil {
		var ret string
		return ret
	}
	return *o.BrowserVersion
}

// GetBrowserVersionOk returns a tuple with the BrowserVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesBrowser) GetBrowserVersionOk() (*string, bool) {
	if o == nil || o.BrowserVersion == nil {
		return nil, false
	}
	return o.BrowserVersion, true
}

// HasBrowserVersion returns a boolean if a field has been set.
func (o *ValuesBrowser) HasBrowserVersion() bool {
	if o != nil && o.BrowserVersion != nil {
		return true
	}

	return false
}

// SetBrowserVersion gets a reference to the given string and assigns it to the BrowserVersion field.
func (o *ValuesBrowser) SetBrowserVersion(v string) {
	o.BrowserVersion = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *ValuesBrowser) GetRate() float32 {
	if o == nil || o.Rate == nil {
		var ret float32
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesBrowser) GetRateOk() (*float32, bool) {
	if o == nil || o.Rate == nil {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *ValuesBrowser) HasRate() bool {
	if o != nil && o.Rate != nil {
		return true
	}

	return false
}

// SetRate gets a reference to the given float32 and assigns it to the Rate field.
func (o *ValuesBrowser) SetRate(v float32) {
	o.Rate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValuesBrowser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.BrowserVersion != nil {
		toSerialize["browser_version"] = o.BrowserVersion
	}
	if o.Rate != nil {
		toSerialize["rate"] = o.Rate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ValuesBrowser) UnmarshalJSON(bytes []byte) (err error) {
	varValuesBrowser := _ValuesBrowser{}

	if err = json.Unmarshal(bytes, &varValuesBrowser); err == nil {
		*o = ValuesBrowser(varValuesBrowser)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "browser_version")
		delete(additionalProperties, "rate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValuesBrowser is a helper abstraction for handling nullable valuesbrowser types.
type NullableValuesBrowser struct {
	value *ValuesBrowser
	isSet bool
}

// Get returns the value.
func (v NullableValuesBrowser) Get() *ValuesBrowser {
	return v.value
}

// Set modifies the value.
func (v *NullableValuesBrowser) Set(val *ValuesBrowser) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValuesBrowser) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValuesBrowser) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValuesBrowser returns a pointer to a new instance of NullableValuesBrowser.
func NewNullableValuesBrowser(val *ValuesBrowser) *NullableValuesBrowser {
	return &NullableValuesBrowser{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValuesBrowser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableValuesBrowser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
