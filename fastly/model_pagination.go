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

// Pagination struct for Pagination
type Pagination struct {
	Links                *PaginationLinks `json:"links,omitempty"`
	Meta                 *PaginationMeta  `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _Pagination Pagination

// NewPagination instantiates a new Pagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagination() *Pagination {
	this := Pagination{}
	return &this
}

// NewPaginationWithDefaults instantiates a new Pagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationWithDefaults() *Pagination {
	this := Pagination{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Pagination) GetLinks() PaginationLinks {
	if o == nil || o.Links == nil {
		var ret PaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetLinksOk() (*PaginationLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Pagination) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PaginationLinks and assigns it to the Links field.
func (o *Pagination) SetLinks(v PaginationLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Pagination) GetMeta() PaginationMeta {
	if o == nil || o.Meta == nil {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pagination) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Pagination) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *Pagination) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Pagination) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Pagination) UnmarshalJSON(bytes []byte) (err error) {
	varPagination := _Pagination{}

	if err = json.Unmarshal(bytes, &varPagination); err == nil {
		*o = Pagination(varPagination)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "links")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePagination is a helper abstraction for handling nullable pagination types.
type NullablePagination struct {
	value *Pagination
	isSet bool
}

// Get returns the value.
func (v NullablePagination) Get() *Pagination {
	return v.value
}

// Set modifies the value.
func (v *NullablePagination) Set(val *Pagination) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePagination) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePagination) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePagination returns a pointer to a new instance of NullablePagination.
func NewNullablePagination(val *Pagination) *NullablePagination {
	return &NullablePagination{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePagination) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
