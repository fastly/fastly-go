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

// ValuesStatusCodes struct for ValuesStatusCodes
type ValuesStatusCodes struct {
	// The HTTP request path.
	URL *string `json:"url,omitempty"`
	// The URL accounts for this percentage of the status code in this dimension.
	RatePerStatus *float32 `json:"rate_per_status,omitempty"`
	// The rate at which the status code in this dimension occurs for this URL.
	RatePerURL           *float32 `json:"rate_per_url,omitempty"`
	AdditionalProperties map[string]any
}

type _ValuesStatusCodes ValuesStatusCodes

// NewValuesStatusCodes instantiates a new ValuesStatusCodes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuesStatusCodes() *ValuesStatusCodes {
	this := ValuesStatusCodes{}
	return &this
}

// NewValuesStatusCodesWithDefaults instantiates a new ValuesStatusCodes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuesStatusCodesWithDefaults() *ValuesStatusCodes {
	this := ValuesStatusCodes{}
	return &this
}

// GetURL returns the URL field value if set, zero value otherwise.
func (o *ValuesStatusCodes) GetURL() string {
	if o == nil || o.URL == nil {
		var ret string
		return ret
	}
	return *o.URL
}

// GetURLOk returns a tuple with the URL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesStatusCodes) GetURLOk() (*string, bool) {
	if o == nil || o.URL == nil {
		return nil, false
	}
	return o.URL, true
}

// HasURL returns a boolean if a field has been set.
func (o *ValuesStatusCodes) HasURL() bool {
	if o != nil && o.URL != nil {
		return true
	}

	return false
}

// SetURL gets a reference to the given string and assigns it to the URL field.
func (o *ValuesStatusCodes) SetURL(v string) {
	o.URL = &v
}

// GetRatePerStatus returns the RatePerStatus field value if set, zero value otherwise.
func (o *ValuesStatusCodes) GetRatePerStatus() float32 {
	if o == nil || o.RatePerStatus == nil {
		var ret float32
		return ret
	}
	return *o.RatePerStatus
}

// GetRatePerStatusOk returns a tuple with the RatePerStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesStatusCodes) GetRatePerStatusOk() (*float32, bool) {
	if o == nil || o.RatePerStatus == nil {
		return nil, false
	}
	return o.RatePerStatus, true
}

// HasRatePerStatus returns a boolean if a field has been set.
func (o *ValuesStatusCodes) HasRatePerStatus() bool {
	if o != nil && o.RatePerStatus != nil {
		return true
	}

	return false
}

// SetRatePerStatus gets a reference to the given float32 and assigns it to the RatePerStatus field.
func (o *ValuesStatusCodes) SetRatePerStatus(v float32) {
	o.RatePerStatus = &v
}

// GetRatePerURL returns the RatePerURL field value if set, zero value otherwise.
func (o *ValuesStatusCodes) GetRatePerURL() float32 {
	if o == nil || o.RatePerURL == nil {
		var ret float32
		return ret
	}
	return *o.RatePerURL
}

// GetRatePerURLOk returns a tuple with the RatePerURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesStatusCodes) GetRatePerURLOk() (*float32, bool) {
	if o == nil || o.RatePerURL == nil {
		return nil, false
	}
	return o.RatePerURL, true
}

// HasRatePerURL returns a boolean if a field has been set.
func (o *ValuesStatusCodes) HasRatePerURL() bool {
	if o != nil && o.RatePerURL != nil {
		return true
	}

	return false
}

// SetRatePerURL gets a reference to the given float32 and assigns it to the RatePerURL field.
func (o *ValuesStatusCodes) SetRatePerURL(v float32) {
	o.RatePerURL = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValuesStatusCodes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.URL != nil {
		toSerialize["url"] = o.URL
	}
	if o.RatePerStatus != nil {
		toSerialize["rate_per_status"] = o.RatePerStatus
	}
	if o.RatePerURL != nil {
		toSerialize["rate_per_url"] = o.RatePerURL
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ValuesStatusCodes) UnmarshalJSON(bytes []byte) (err error) {
	varValuesStatusCodes := _ValuesStatusCodes{}

	if err = json.Unmarshal(bytes, &varValuesStatusCodes); err == nil {
		*o = ValuesStatusCodes(varValuesStatusCodes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "url")
		delete(additionalProperties, "rate_per_status")
		delete(additionalProperties, "rate_per_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValuesStatusCodes is a helper abstraction for handling nullable valuesstatuscodes types.
type NullableValuesStatusCodes struct {
	value *ValuesStatusCodes
	isSet bool
}

// Get returns the value.
func (v NullableValuesStatusCodes) Get() *ValuesStatusCodes {
	return v.value
}

// Set modifies the value.
func (v *NullableValuesStatusCodes) Set(val *ValuesStatusCodes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValuesStatusCodes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValuesStatusCodes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValuesStatusCodes returns a pointer to a new instance of NullableValuesStatusCodes.
func NewNullableValuesStatusCodes(val *ValuesStatusCodes) *NullableValuesStatusCodes {
	return &NullableValuesStatusCodes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValuesStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableValuesStatusCodes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
