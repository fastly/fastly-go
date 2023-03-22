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

// PaginationLinks struct for PaginationLinks
type PaginationLinks struct {
	// The first page of data.
	First NullableString `json:"first,omitempty"`
	// The last page of data.
	Last NullableString `json:"last,omitempty"`
	// The previous page of data.
	Prev NullableString `json:"prev,omitempty"`
	// The next page of data.
	Next NullableString `json:"next,omitempty"`
	AdditionalProperties map[string]any
}

type _PaginationLinks PaginationLinks

// NewPaginationLinks instantiates a new PaginationLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationLinks() *PaginationLinks {
	this := PaginationLinks{}
	return &this
}

// NewPaginationLinksWithDefaults instantiates a new PaginationLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationLinksWithDefaults() *PaginationLinks {
	this := PaginationLinks{}
	return &this
}

// GetFirst returns the First field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginationLinks) GetFirst() string {
	if o == nil || o.First.Get() == nil {
		var ret string
		return ret
	}
	return *o.First.Get()
}

// GetFirstOk returns a tuple with the First field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationLinks) GetFirstOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.First.Get(), o.First.IsSet()
}

// HasFirst returns a boolean if a field has been set.
func (o *PaginationLinks) HasFirst() bool {
	if o != nil && o.First.IsSet() {
		return true
	}

	return false
}

// SetFirst gets a reference to the given NullableString and assigns it to the First field.
func (o *PaginationLinks) SetFirst(v string) {
	o.First.Set(&v)
}
// SetFirstNil sets the value for First to be an explicit nil
func (o *PaginationLinks) SetFirstNil() {
	o.First.Set(nil)
}

// UnsetFirst ensures that no value is present for First, not even an explicit nil
func (o *PaginationLinks) UnsetFirst() {
	o.First.Unset()
}

// GetLast returns the Last field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginationLinks) GetLast() string {
	if o == nil || o.Last.Get() == nil {
		var ret string
		return ret
	}
	return *o.Last.Get()
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationLinks) GetLastOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Last.Get(), o.Last.IsSet()
}

// HasLast returns a boolean if a field has been set.
func (o *PaginationLinks) HasLast() bool {
	if o != nil && o.Last.IsSet() {
		return true
	}

	return false
}

// SetLast gets a reference to the given NullableString and assigns it to the Last field.
func (o *PaginationLinks) SetLast(v string) {
	o.Last.Set(&v)
}
// SetLastNil sets the value for Last to be an explicit nil
func (o *PaginationLinks) SetLastNil() {
	o.Last.Set(nil)
}

// UnsetLast ensures that no value is present for Last, not even an explicit nil
func (o *PaginationLinks) UnsetLast() {
	o.Last.Unset()
}

// GetPrev returns the Prev field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginationLinks) GetPrev() string {
	if o == nil || o.Prev.Get() == nil {
		var ret string
		return ret
	}
	return *o.Prev.Get()
}

// GetPrevOk returns a tuple with the Prev field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationLinks) GetPrevOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Prev.Get(), o.Prev.IsSet()
}

// HasPrev returns a boolean if a field has been set.
func (o *PaginationLinks) HasPrev() bool {
	if o != nil && o.Prev.IsSet() {
		return true
	}

	return false
}

// SetPrev gets a reference to the given NullableString and assigns it to the Prev field.
func (o *PaginationLinks) SetPrev(v string) {
	o.Prev.Set(&v)
}
// SetPrevNil sets the value for Prev to be an explicit nil
func (o *PaginationLinks) SetPrevNil() {
	o.Prev.Set(nil)
}

// UnsetPrev ensures that no value is present for Prev, not even an explicit nil
func (o *PaginationLinks) UnsetPrev() {
	o.Prev.Unset()
}

// GetNext returns the Next field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginationLinks) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}
	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginationLinks) GetNextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// HasNext returns a boolean if a field has been set.
func (o *PaginationLinks) HasNext() bool {
	if o != nil && o.Next.IsSet() {
		return true
	}

	return false
}

// SetNext gets a reference to the given NullableString and assigns it to the Next field.
func (o *PaginationLinks) SetNext(v string) {
	o.Next.Set(&v)
}
// SetNextNil sets the value for Next to be an explicit nil
func (o *PaginationLinks) SetNextNil() {
	o.Next.Set(nil)
}

// UnsetNext ensures that no value is present for Next, not even an explicit nil
func (o *PaginationLinks) UnsetNext() {
	o.Next.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PaginationLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.First.IsSet() {
		toSerialize["first"] = o.First.Get()
	}
	if o.Last.IsSet() {
		toSerialize["last"] = o.Last.Get()
	}
	if o.Prev.IsSet() {
		toSerialize["prev"] = o.Prev.Get()
	}
	if o.Next.IsSet() {
		toSerialize["next"] = o.Next.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *PaginationLinks) UnmarshalJSON(bytes []byte) (err error) {
	varPaginationLinks := _PaginationLinks{}

	if err = json.Unmarshal(bytes, &varPaginationLinks); err == nil {
		*o = PaginationLinks(varPaginationLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "first")
		delete(additionalProperties, "last")
		delete(additionalProperties, "prev")
		delete(additionalProperties, "next")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePaginationLinks is a helper abstraction for handling nullable paginationlinks types. 
type NullablePaginationLinks struct {
	value *PaginationLinks
	isSet bool
}

// Get returns the value.
func (v NullablePaginationLinks) Get() *PaginationLinks {
	return v.value
}

// Set modifies the value.
func (v *NullablePaginationLinks) Set(val *PaginationLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePaginationLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePaginationLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePaginationLinks returns a pointer to a new instance of NullablePaginationLinks.
func NewNullablePaginationLinks(val *PaginationLinks) *NullablePaginationLinks {
	return &NullablePaginationLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePaginationLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullablePaginationLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
