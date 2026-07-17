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

// PaginationMeta1 struct for PaginationMeta1
type PaginationMeta1 struct {
	// Current page.
	CurrentPage *int32 `json:"current_page,omitempty"`
	// Number of records per page.
	PerPage *int32 `json:"per_page,omitempty"`
	// Total records in result set.
	RecordCount *int32 `json:"record_count,omitempty"`
	// Total pages in result set.
	TotalPages           *int32 `json:"total_pages,omitempty"`
	AdditionalProperties map[string]any
}

type _PaginationMeta1 PaginationMeta1

// NewPaginationMeta1 instantiates a new PaginationMeta1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationMeta1() *PaginationMeta1 {
	this := PaginationMeta1{}
	var perPage int32 = 20
	this.PerPage = &perPage
	return &this
}

// NewPaginationMeta1WithDefaults instantiates a new PaginationMeta1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationMeta1WithDefaults() *PaginationMeta1 {
	this := PaginationMeta1{}
	var perPage int32 = 20
	this.PerPage = &perPage
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *PaginationMeta1) GetCurrentPage() int32 {
	if o == nil || o.CurrentPage == nil {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationMeta1) GetCurrentPageOk() (*int32, bool) {
	if o == nil || o.CurrentPage == nil {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *PaginationMeta1) HasCurrentPage() bool {
	if o != nil && o.CurrentPage != nil {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *PaginationMeta1) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *PaginationMeta1) GetPerPage() int32 {
	if o == nil || o.PerPage == nil {
		var ret int32
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationMeta1) GetPerPageOk() (*int32, bool) {
	if o == nil || o.PerPage == nil {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *PaginationMeta1) HasPerPage() bool {
	if o != nil && o.PerPage != nil {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int32 and assigns it to the PerPage field.
func (o *PaginationMeta1) SetPerPage(v int32) {
	o.PerPage = &v
}

// GetRecordCount returns the RecordCount field value if set, zero value otherwise.
func (o *PaginationMeta1) GetRecordCount() int32 {
	if o == nil || o.RecordCount == nil {
		var ret int32
		return ret
	}
	return *o.RecordCount
}

// GetRecordCountOk returns a tuple with the RecordCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationMeta1) GetRecordCountOk() (*int32, bool) {
	if o == nil || o.RecordCount == nil {
		return nil, false
	}
	return o.RecordCount, true
}

// HasRecordCount returns a boolean if a field has been set.
func (o *PaginationMeta1) HasRecordCount() bool {
	if o != nil && o.RecordCount != nil {
		return true
	}

	return false
}

// SetRecordCount gets a reference to the given int32 and assigns it to the RecordCount field.
func (o *PaginationMeta1) SetRecordCount(v int32) {
	o.RecordCount = &v
}

// GetTotalPages returns the TotalPages field value if set, zero value otherwise.
func (o *PaginationMeta1) GetTotalPages() int32 {
	if o == nil || o.TotalPages == nil {
		var ret int32
		return ret
	}
	return *o.TotalPages
}

// GetTotalPagesOk returns a tuple with the TotalPages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationMeta1) GetTotalPagesOk() (*int32, bool) {
	if o == nil || o.TotalPages == nil {
		return nil, false
	}
	return o.TotalPages, true
}

// HasTotalPages returns a boolean if a field has been set.
func (o *PaginationMeta1) HasTotalPages() bool {
	if o != nil && o.TotalPages != nil {
		return true
	}

	return false
}

// SetTotalPages gets a reference to the given int32 and assigns it to the TotalPages field.
func (o *PaginationMeta1) SetTotalPages(v int32) {
	o.TotalPages = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PaginationMeta1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CurrentPage != nil {
		toSerialize["current_page"] = o.CurrentPage
	}
	if o.PerPage != nil {
		toSerialize["per_page"] = o.PerPage
	}
	if o.RecordCount != nil {
		toSerialize["record_count"] = o.RecordCount
	}
	if o.TotalPages != nil {
		toSerialize["total_pages"] = o.TotalPages
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PaginationMeta1) UnmarshalJSON(bytes []byte) (err error) {
	varPaginationMeta1 := _PaginationMeta1{}

	if err = json.Unmarshal(bytes, &varPaginationMeta1); err == nil {
		*o = PaginationMeta1(varPaginationMeta1)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "current_page")
		delete(additionalProperties, "per_page")
		delete(additionalProperties, "record_count")
		delete(additionalProperties, "total_pages")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePaginationMeta1 is a helper abstraction for handling nullable paginationmeta1 types.
type NullablePaginationMeta1 struct {
	value *PaginationMeta1
	isSet bool
}

// Get returns the value.
func (v NullablePaginationMeta1) Get() *PaginationMeta1 {
	return v.value
}

// Set modifies the value.
func (v *NullablePaginationMeta1) Set(val *PaginationMeta1) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePaginationMeta1) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePaginationMeta1) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePaginationMeta1 returns a pointer to a new instance of NullablePaginationMeta1.
func NewNullablePaginationMeta1(val *PaginationMeta1) *NullablePaginationMeta1 {
	return &NullablePaginationMeta1{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePaginationMeta1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePaginationMeta1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
