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

// PaginationCursorMeta struct for PaginationCursorMeta
type PaginationCursorMeta struct {
	// Cursor for the next page.
	NextCursor *string `json:"next_cursor,omitempty"`
	// Maximum number of results returned.
	Limit *int32 `json:"limit,omitempty"`
	AdditionalProperties map[string]any
}

type _PaginationCursorMeta PaginationCursorMeta

// NewPaginationCursorMeta instantiates a new PaginationCursorMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationCursorMeta() *PaginationCursorMeta {
	this := PaginationCursorMeta{}
	return &this
}

// NewPaginationCursorMetaWithDefaults instantiates a new PaginationCursorMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationCursorMetaWithDefaults() *PaginationCursorMeta {
	this := PaginationCursorMeta{}
	return &this
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *PaginationCursorMeta) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationCursorMeta) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *PaginationCursorMeta) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *PaginationCursorMeta) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PaginationCursorMeta) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationCursorMeta) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PaginationCursorMeta) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PaginationCursorMeta) SetLimit(v int32) {
	o.Limit = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PaginationCursorMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *PaginationCursorMeta) UnmarshalJSON(bytes []byte) (err error) {
	varPaginationCursorMeta := _PaginationCursorMeta{}

	if err = json.Unmarshal(bytes, &varPaginationCursorMeta); err == nil {
		*o = PaginationCursorMeta(varPaginationCursorMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePaginationCursorMeta is a helper abstraction for handling nullable paginationcursormeta types. 
type NullablePaginationCursorMeta struct {
	value *PaginationCursorMeta
	isSet bool
}

// Get returns the value.
func (v NullablePaginationCursorMeta) Get() *PaginationCursorMeta {
	return v.value
}

// Set modifies the value.
func (v *NullablePaginationCursorMeta) Set(val *PaginationCursorMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePaginationCursorMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePaginationCursorMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePaginationCursorMeta returns a pointer to a new instance of NullablePaginationCursorMeta.
func NewNullablePaginationCursorMeta(val *PaginationCursorMeta) *NullablePaginationCursorMeta {
	return &NullablePaginationCursorMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePaginationCursorMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullablePaginationCursorMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
