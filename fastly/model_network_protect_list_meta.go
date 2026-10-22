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

// NetworkProtectListMeta Pagination metadata for a collection response.
type NetworkProtectListMeta struct {
	// Effective page limit after defaulting and clamping.
	Limit int32 `json:"limit"`
	// Continuation value to send as cursor when more resources remain. Omitted on the final page.
	NextCursor           *string `json:"next_cursor,omitempty"`
	AdditionalProperties map[string]any
}

type _NetworkProtectListMeta NetworkProtectListMeta

// NewNetworkProtectListMeta instantiates a new NetworkProtectListMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectListMeta(limit int32) *NetworkProtectListMeta {
	this := NetworkProtectListMeta{}
	this.Limit = limit
	return &this
}

// NewNetworkProtectListMetaWithDefaults instantiates a new NetworkProtectListMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectListMetaWithDefaults() *NetworkProtectListMeta {
	this := NetworkProtectListMeta{}
	var limit int32 = 100
	this.Limit = limit
	return &this
}

// GetLimit returns the Limit field value
func (o *NetworkProtectListMeta) GetLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectListMeta) GetLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *NetworkProtectListMeta) SetLimit(v int32) {
	o.Limit = v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *NetworkProtectListMeta) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectListMeta) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *NetworkProtectListMeta) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *NetworkProtectListMeta) SetNextCursor(v string) {
	o.NextCursor = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectListMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectListMeta) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectListMeta := _NetworkProtectListMeta{}

	if err = json.Unmarshal(bytes, &varNetworkProtectListMeta); err == nil {
		*o = NetworkProtectListMeta(varNetworkProtectListMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "limit")
		delete(additionalProperties, "next_cursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectListMeta is a helper abstraction for handling nullable networkprotectlistmeta types.
type NullableNetworkProtectListMeta struct {
	value *NetworkProtectListMeta
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectListMeta) Get() *NetworkProtectListMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectListMeta) Set(val *NetworkProtectListMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectListMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectListMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectListMeta returns a pointer to a new instance of NullableNetworkProtectListMeta.
func NewNullableNetworkProtectListMeta(val *NetworkProtectListMeta) *NullableNetworkProtectListMeta {
	return &NullableNetworkProtectListMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectListMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectListMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
