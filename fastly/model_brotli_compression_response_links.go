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

// BrotliCompressionResponseLinks struct for BrotliCompressionResponseLinks
type BrotliCompressionResponseLinks struct {
	Links                *BrotliCompressionResponseLinksLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _BrotliCompressionResponseLinks BrotliCompressionResponseLinks

// NewBrotliCompressionResponseLinks instantiates a new BrotliCompressionResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrotliCompressionResponseLinks() *BrotliCompressionResponseLinks {
	this := BrotliCompressionResponseLinks{}
	return &this
}

// NewBrotliCompressionResponseLinksWithDefaults instantiates a new BrotliCompressionResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrotliCompressionResponseLinksWithDefaults() *BrotliCompressionResponseLinks {
	this := BrotliCompressionResponseLinks{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BrotliCompressionResponseLinks) GetLinks() BrotliCompressionResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret BrotliCompressionResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrotliCompressionResponseLinks) GetLinksOk() (*BrotliCompressionResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BrotliCompressionResponseLinks) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BrotliCompressionResponseLinksLinks and assigns it to the Links field.
func (o *BrotliCompressionResponseLinks) SetLinks(v BrotliCompressionResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BrotliCompressionResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BrotliCompressionResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	varBrotliCompressionResponseLinks := _BrotliCompressionResponseLinks{}

	if err = json.Unmarshal(bytes, &varBrotliCompressionResponseLinks); err == nil {
		*o = BrotliCompressionResponseLinks(varBrotliCompressionResponseLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBrotliCompressionResponseLinks is a helper abstraction for handling nullable brotlicompressionresponselinks types.
type NullableBrotliCompressionResponseLinks struct {
	value *BrotliCompressionResponseLinks
	isSet bool
}

// Get returns the value.
func (v NullableBrotliCompressionResponseLinks) Get() *BrotliCompressionResponseLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableBrotliCompressionResponseLinks) Set(val *BrotliCompressionResponseLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBrotliCompressionResponseLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBrotliCompressionResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBrotliCompressionResponseLinks returns a pointer to a new instance of NullableBrotliCompressionResponseLinks.
func NewNullableBrotliCompressionResponseLinks(val *BrotliCompressionResponseLinks) *NullableBrotliCompressionResponseLinks {
	return &NullableBrotliCompressionResponseLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBrotliCompressionResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBrotliCompressionResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
