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

// PaginationMeta Cursor-based pagination metadata.
type PaginationMeta struct {
	// The number of records returned per page.
	Limit *int32 `json:"limit,omitempty"`
	// Cursor value used to retrieve the next page of results. Empty if there are no more results.
	NextCursor *string `json:"next_cursor,omitempty"`
	// Cursor value used to retrieve the previous page of results. Empty if there is no previous page.
	PreviousCursor *string `json:"previous_cursor,omitempty"`
	// The sort order applied to the results.
	Sort                 *string `json:"sort,omitempty"`
	AdditionalProperties map[string]any
}

type _PaginationMeta PaginationMeta

// NewPaginationMeta instantiates a new PaginationMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationMeta() *PaginationMeta {
	this := PaginationMeta{}
	return &this
}

// NewPaginationMetaWithDefaults instantiates a new PaginationMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationMetaWithDefaults() *PaginationMeta {
	this := PaginationMeta{}
	return &this
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PaginationMeta) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationMeta) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PaginationMeta) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PaginationMeta) SetLimit(v int32) {
	o.Limit = &v
}

// GetNextCursor returns the NextCursor field value if set, zero value otherwise.
func (o *PaginationMeta) GetNextCursor() string {
	if o == nil || o.NextCursor == nil {
		var ret string
		return ret
	}
	return *o.NextCursor
}

// GetNextCursorOk returns a tuple with the NextCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationMeta) GetNextCursorOk() (*string, bool) {
	if o == nil || o.NextCursor == nil {
		return nil, false
	}
	return o.NextCursor, true
}

// HasNextCursor returns a boolean if a field has been set.
func (o *PaginationMeta) HasNextCursor() bool {
	if o != nil && o.NextCursor != nil {
		return true
	}

	return false
}

// SetNextCursor gets a reference to the given string and assigns it to the NextCursor field.
func (o *PaginationMeta) SetNextCursor(v string) {
	o.NextCursor = &v
}

// GetPreviousCursor returns the PreviousCursor field value if set, zero value otherwise.
func (o *PaginationMeta) GetPreviousCursor() string {
	if o == nil || o.PreviousCursor == nil {
		var ret string
		return ret
	}
	return *o.PreviousCursor
}

// GetPreviousCursorOk returns a tuple with the PreviousCursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationMeta) GetPreviousCursorOk() (*string, bool) {
	if o == nil || o.PreviousCursor == nil {
		return nil, false
	}
	return o.PreviousCursor, true
}

// HasPreviousCursor returns a boolean if a field has been set.
func (o *PaginationMeta) HasPreviousCursor() bool {
	if o != nil && o.PreviousCursor != nil {
		return true
	}

	return false
}

// SetPreviousCursor gets a reference to the given string and assigns it to the PreviousCursor field.
func (o *PaginationMeta) SetPreviousCursor(v string) {
	o.PreviousCursor = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *PaginationMeta) GetSort() string {
	if o == nil || o.Sort == nil {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationMeta) GetSortOk() (*string, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *PaginationMeta) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *PaginationMeta) SetSort(v string) {
	o.Sort = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PaginationMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.NextCursor != nil {
		toSerialize["next_cursor"] = o.NextCursor
	}
	if o.PreviousCursor != nil {
		toSerialize["previous_cursor"] = o.PreviousCursor
	}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PaginationMeta) UnmarshalJSON(bytes []byte) (err error) {
	varPaginationMeta := _PaginationMeta{}

	if err = json.Unmarshal(bytes, &varPaginationMeta); err == nil {
		*o = PaginationMeta(varPaginationMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "limit")
		delete(additionalProperties, "next_cursor")
		delete(additionalProperties, "previous_cursor")
		delete(additionalProperties, "sort")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePaginationMeta is a helper abstraction for handling nullable paginationmeta types.
type NullablePaginationMeta struct {
	value *PaginationMeta
	isSet bool
}

// Get returns the value.
func (v NullablePaginationMeta) Get() *PaginationMeta {
	return v.value
}

// Set modifies the value.
func (v *NullablePaginationMeta) Set(val *PaginationMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePaginationMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePaginationMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePaginationMeta returns a pointer to a new instance of NullablePaginationMeta.
func NewNullablePaginationMeta(val *PaginationMeta) *NullablePaginationMeta {
	return &NullablePaginationMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePaginationMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePaginationMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
