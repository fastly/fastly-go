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

// ValuesCacheHitRatio struct for ValuesCacheHitRatio
type ValuesCacheHitRatio struct {
	// The cache hit ratio for the URL specified in the dimension.
	CacheHitRatio        *float32 `json:"cache_hit_ratio,omitempty"`
	AdditionalProperties map[string]any
}

type _ValuesCacheHitRatio ValuesCacheHitRatio

// NewValuesCacheHitRatio instantiates a new ValuesCacheHitRatio object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValuesCacheHitRatio() *ValuesCacheHitRatio {
	this := ValuesCacheHitRatio{}
	return &this
}

// NewValuesCacheHitRatioWithDefaults instantiates a new ValuesCacheHitRatio object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValuesCacheHitRatioWithDefaults() *ValuesCacheHitRatio {
	this := ValuesCacheHitRatio{}
	return &this
}

// GetCacheHitRatio returns the CacheHitRatio field value if set, zero value otherwise.
func (o *ValuesCacheHitRatio) GetCacheHitRatio() float32 {
	if o == nil || o.CacheHitRatio == nil {
		var ret float32
		return ret
	}
	return *o.CacheHitRatio
}

// GetCacheHitRatioOk returns a tuple with the CacheHitRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValuesCacheHitRatio) GetCacheHitRatioOk() (*float32, bool) {
	if o == nil || o.CacheHitRatio == nil {
		return nil, false
	}
	return o.CacheHitRatio, true
}

// HasCacheHitRatio returns a boolean if a field has been set.
func (o *ValuesCacheHitRatio) HasCacheHitRatio() bool {
	if o != nil && o.CacheHitRatio != nil {
		return true
	}

	return false
}

// SetCacheHitRatio gets a reference to the given float32 and assigns it to the CacheHitRatio field.
func (o *ValuesCacheHitRatio) SetCacheHitRatio(v float32) {
	o.CacheHitRatio = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ValuesCacheHitRatio) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CacheHitRatio != nil {
		toSerialize["cache_hit_ratio"] = o.CacheHitRatio
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ValuesCacheHitRatio) UnmarshalJSON(bytes []byte) (err error) {
	varValuesCacheHitRatio := _ValuesCacheHitRatio{}

	if err = json.Unmarshal(bytes, &varValuesCacheHitRatio); err == nil {
		*o = ValuesCacheHitRatio(varValuesCacheHitRatio)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cache_hit_ratio")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableValuesCacheHitRatio is a helper abstraction for handling nullable valuescachehitratio types.
type NullableValuesCacheHitRatio struct {
	value *ValuesCacheHitRatio
	isSet bool
}

// Get returns the value.
func (v NullableValuesCacheHitRatio) Get() *ValuesCacheHitRatio {
	return v.value
}

// Set modifies the value.
func (v *NullableValuesCacheHitRatio) Set(val *ValuesCacheHitRatio) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableValuesCacheHitRatio) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableValuesCacheHitRatio) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableValuesCacheHitRatio returns a pointer to a new instance of NullableValuesCacheHitRatio.
func NewNullableValuesCacheHitRatio(val *ValuesCacheHitRatio) *NullableValuesCacheHitRatio {
	return &NullableValuesCacheHitRatio{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableValuesCacheHitRatio) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableValuesCacheHitRatio) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
