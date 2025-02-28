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

// FanoutResponseLinks struct for FanoutResponseLinks
type FanoutResponseLinks struct {
	Links                *FanoutResponseLinksLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _FanoutResponseLinks FanoutResponseLinks

// NewFanoutResponseLinks instantiates a new FanoutResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFanoutResponseLinks() *FanoutResponseLinks {
	this := FanoutResponseLinks{}
	return &this
}

// NewFanoutResponseLinksWithDefaults instantiates a new FanoutResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFanoutResponseLinksWithDefaults() *FanoutResponseLinks {
	this := FanoutResponseLinks{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FanoutResponseLinks) GetLinks() FanoutResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret FanoutResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanoutResponseLinks) GetLinksOk() (*FanoutResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FanoutResponseLinks) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given FanoutResponseLinksLinks and assigns it to the Links field.
func (o *FanoutResponseLinks) SetLinks(v FanoutResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o FanoutResponseLinks) MarshalJSON() ([]byte, error) {
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
func (o *FanoutResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	varFanoutResponseLinks := _FanoutResponseLinks{}

	if err = json.Unmarshal(bytes, &varFanoutResponseLinks); err == nil {
		*o = FanoutResponseLinks(varFanoutResponseLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableFanoutResponseLinks is a helper abstraction for handling nullable fanoutresponselinks types.
type NullableFanoutResponseLinks struct {
	value *FanoutResponseLinks
	isSet bool
}

// Get returns the value.
func (v NullableFanoutResponseLinks) Get() *FanoutResponseLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableFanoutResponseLinks) Set(val *FanoutResponseLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableFanoutResponseLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableFanoutResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFanoutResponseLinks returns a pointer to a new instance of NullableFanoutResponseLinks.
func NewNullableFanoutResponseLinks(val *FanoutResponseLinks) *NullableFanoutResponseLinks {
	return &NullableFanoutResponseLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableFanoutResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableFanoutResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
