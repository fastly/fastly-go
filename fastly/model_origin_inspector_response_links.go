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

// OriginInspectorResponseLinks struct for OriginInspectorResponseLinks
type OriginInspectorResponseLinks struct {
	Links                *OriginInspectorResponseLinksLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _OriginInspectorResponseLinks OriginInspectorResponseLinks

// NewOriginInspectorResponseLinks instantiates a new OriginInspectorResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspectorResponseLinks() *OriginInspectorResponseLinks {
	this := OriginInspectorResponseLinks{}
	return &this
}

// NewOriginInspectorResponseLinksWithDefaults instantiates a new OriginInspectorResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorResponseLinksWithDefaults() *OriginInspectorResponseLinks {
	this := OriginInspectorResponseLinks{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OriginInspectorResponseLinks) GetLinks() OriginInspectorResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret OriginInspectorResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorResponseLinks) GetLinksOk() (*OriginInspectorResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OriginInspectorResponseLinks) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OriginInspectorResponseLinksLinks and assigns it to the Links field.
func (o *OriginInspectorResponseLinks) SetLinks(v OriginInspectorResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspectorResponseLinks) MarshalJSON() ([]byte, error) {
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
func (o *OriginInspectorResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspectorResponseLinks := _OriginInspectorResponseLinks{}

	if err = json.Unmarshal(bytes, &varOriginInspectorResponseLinks); err == nil {
		*o = OriginInspectorResponseLinks(varOriginInspectorResponseLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOriginInspectorResponseLinks is a helper abstraction for handling nullable origininspectorresponselinks types.
type NullableOriginInspectorResponseLinks struct {
	value *OriginInspectorResponseLinks
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspectorResponseLinks) Get() *OriginInspectorResponseLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspectorResponseLinks) Set(val *OriginInspectorResponseLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspectorResponseLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspectorResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspectorResponseLinks returns a pointer to a new instance of NullableOriginInspectorResponseLinks.
func NewNullableOriginInspectorResponseLinks(val *OriginInspectorResponseLinks) *NullableOriginInspectorResponseLinks {
	return &NullableOriginInspectorResponseLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspectorResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOriginInspectorResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
