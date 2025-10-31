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

// ApiDiscoveryResponseLinks struct for ApiDiscoveryResponseLinks
type ApiDiscoveryResponseLinks struct {
	Links                *ApiDiscoveryResponseLinksLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _ApiDiscoveryResponseLinks ApiDiscoveryResponseLinks

// NewApiDiscoveryResponseLinks instantiates a new ApiDiscoveryResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDiscoveryResponseLinks() *ApiDiscoveryResponseLinks {
	this := ApiDiscoveryResponseLinks{}
	return &this
}

// NewApiDiscoveryResponseLinksWithDefaults instantiates a new ApiDiscoveryResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDiscoveryResponseLinksWithDefaults() *ApiDiscoveryResponseLinks {
	this := ApiDiscoveryResponseLinks{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ApiDiscoveryResponseLinks) GetLinks() ApiDiscoveryResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret ApiDiscoveryResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDiscoveryResponseLinks) GetLinksOk() (*ApiDiscoveryResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ApiDiscoveryResponseLinks) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ApiDiscoveryResponseLinksLinks and assigns it to the Links field.
func (o *ApiDiscoveryResponseLinks) SetLinks(v ApiDiscoveryResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ApiDiscoveryResponseLinks) MarshalJSON() ([]byte, error) {
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
func (o *ApiDiscoveryResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	varApiDiscoveryResponseLinks := _ApiDiscoveryResponseLinks{}

	if err = json.Unmarshal(bytes, &varApiDiscoveryResponseLinks); err == nil {
		*o = ApiDiscoveryResponseLinks(varApiDiscoveryResponseLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableApiDiscoveryResponseLinks is a helper abstraction for handling nullable apidiscoveryresponselinks types.
type NullableApiDiscoveryResponseLinks struct {
	value *ApiDiscoveryResponseLinks
	isSet bool
}

// Get returns the value.
func (v NullableApiDiscoveryResponseLinks) Get() *ApiDiscoveryResponseLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableApiDiscoveryResponseLinks) Set(val *ApiDiscoveryResponseLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableApiDiscoveryResponseLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableApiDiscoveryResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableApiDiscoveryResponseLinks returns a pointer to a new instance of NullableApiDiscoveryResponseLinks.
func NewNullableApiDiscoveryResponseLinks(val *ApiDiscoveryResponseLinks) *NullableApiDiscoveryResponseLinks {
	return &NullableApiDiscoveryResponseLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableApiDiscoveryResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableApiDiscoveryResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
