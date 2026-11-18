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

// ComputeACLListEntriesMeta struct for ComputeACLListEntriesMeta
type ComputeACLListEntriesMeta struct {
	// The maximum number of results shown in this response.
	Limit *int32 `json:"limit,omitempty"`
	// Used for pagination, supply to the next request to get the next block of results.
	NextCursor           *string `json:"next_cursor,omitempty"`
	AdditionalProperties map[string]any
}

type _ComputeACLListEntriesMeta ComputeACLListEntriesMeta

// NewComputeACLListEntriesMeta instantiates a new ComputeACLListEntriesMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeACLListEntriesMeta() *ComputeACLListEntriesMeta {
	this := ComputeACLListEntriesMeta{}
	return &this
}

// NewComputeACLListEntriesMetaWithDefaults instantiates a new ComputeACLListEntriesMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeACLListEntriesMetaWithDefaults() *ComputeACLListEntriesMeta {
	this := ComputeACLListEntriesMeta{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ComputeACLListEntriesMeta) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeACLListEntriesMeta) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ComputeACLListEntriesMeta) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *ComputeACLListEntriesMeta) SetLimit(v int32) {
	o.Limit = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *ComputeACLListEntriesMeta) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeACLListEntriesMeta) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *ComputeACLListEntriesMeta) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *ComputeACLListEntriesMeta) SetNextCursor(v string) {
	o.NextCursor = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ComputeACLListEntriesMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Limit != nil {
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
func (o *ComputeACLListEntriesMeta) UnmarshalJSON(bytes []byte) (err error) {
	varComputeACLListEntriesMeta := _ComputeACLListEntriesMeta{}

	if err = json.Unmarshal(bytes, &varComputeACLListEntriesMeta); err == nil {
		*o = ComputeACLListEntriesMeta(varComputeACLListEntriesMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "limit")
		delete(additionalProperties, "next_cursor")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableComputeACLListEntriesMeta is a helper abstraction for handling nullable computeacllistentriesmeta types.
type NullableComputeACLListEntriesMeta struct {
	value *ComputeACLListEntriesMeta
	isSet bool
}

// Get returns the value.
func (v NullableComputeACLListEntriesMeta) Get() *ComputeACLListEntriesMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableComputeACLListEntriesMeta) Set(val *ComputeACLListEntriesMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableComputeACLListEntriesMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableComputeACLListEntriesMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableComputeACLListEntriesMeta returns a pointer to a new instance of NullableComputeACLListEntriesMeta.
func NewNullableComputeACLListEntriesMeta(val *ComputeACLListEntriesMeta) *NullableComputeACLListEntriesMeta {
	return &NullableComputeACLListEntriesMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableComputeACLListEntriesMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableComputeACLListEntriesMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
