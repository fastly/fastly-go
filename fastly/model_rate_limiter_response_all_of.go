// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// RateLimiterResponseAllOf struct for RateLimiterResponseAllOf
type RateLimiterResponseAllOf struct {
	// Alphanumeric string identifying the rate limiter.
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RateLimiterResponseAllOf RateLimiterResponseAllOf

// NewRateLimiterResponseAllOf instantiates a new RateLimiterResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimiterResponseAllOf() *RateLimiterResponseAllOf {
	this := RateLimiterResponseAllOf{}
	return &this
}

// NewRateLimiterResponseAllOfWithDefaults instantiates a new RateLimiterResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimiterResponseAllOfWithDefaults() *RateLimiterResponseAllOf {
	this := RateLimiterResponseAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *RateLimiterResponseAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimiterResponseAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *RateLimiterResponseAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *RateLimiterResponseAllOf) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RateLimiterResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RateLimiterResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRateLimiterResponseAllOf := _RateLimiterResponseAllOf{}

	if err = json.Unmarshal(bytes, &varRateLimiterResponseAllOf); err == nil {
		*o = RateLimiterResponseAllOf(varRateLimiterResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRateLimiterResponseAllOf is a helper abstraction for handling nullable ratelimiterresponseallof types. 
type NullableRateLimiterResponseAllOf struct {
	value *RateLimiterResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableRateLimiterResponseAllOf) Get() *RateLimiterResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableRateLimiterResponseAllOf) Set(val *RateLimiterResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRateLimiterResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRateLimiterResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRateLimiterResponseAllOf returns a pointer to a new instance of NullableRateLimiterResponseAllOf.
func NewNullableRateLimiterResponseAllOf(val *RateLimiterResponseAllOf) *NullableRateLimiterResponseAllOf {
	return &NullableRateLimiterResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRateLimiterResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRateLimiterResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
