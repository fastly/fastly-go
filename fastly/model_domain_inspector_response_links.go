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

// DomainInspectorResponseLinks struct for DomainInspectorResponseLinks
type DomainInspectorResponseLinks struct {
	Links                *DomainInspectorResponseLinksLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _DomainInspectorResponseLinks DomainInspectorResponseLinks

// NewDomainInspectorResponseLinks instantiates a new DomainInspectorResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainInspectorResponseLinks() *DomainInspectorResponseLinks {
	this := DomainInspectorResponseLinks{}
	return &this
}

// NewDomainInspectorResponseLinksWithDefaults instantiates a new DomainInspectorResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainInspectorResponseLinksWithDefaults() *DomainInspectorResponseLinks {
	this := DomainInspectorResponseLinks{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DomainInspectorResponseLinks) GetLinks() DomainInspectorResponseLinksLinks {
	if o == nil || o.Links == nil {
		var ret DomainInspectorResponseLinksLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainInspectorResponseLinks) GetLinksOk() (*DomainInspectorResponseLinksLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DomainInspectorResponseLinks) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given DomainInspectorResponseLinksLinks and assigns it to the Links field.
func (o *DomainInspectorResponseLinks) SetLinks(v DomainInspectorResponseLinksLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DomainInspectorResponseLinks) MarshalJSON() ([]byte, error) {
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
func (o *DomainInspectorResponseLinks) UnmarshalJSON(bytes []byte) (err error) {
	varDomainInspectorResponseLinks := _DomainInspectorResponseLinks{}

	if err = json.Unmarshal(bytes, &varDomainInspectorResponseLinks); err == nil {
		*o = DomainInspectorResponseLinks(varDomainInspectorResponseLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDomainInspectorResponseLinks is a helper abstraction for handling nullable domaininspectorresponselinks types.
type NullableDomainInspectorResponseLinks struct {
	value *DomainInspectorResponseLinks
	isSet bool
}

// Get returns the value.
func (v NullableDomainInspectorResponseLinks) Get() *DomainInspectorResponseLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableDomainInspectorResponseLinks) Set(val *DomainInspectorResponseLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDomainInspectorResponseLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDomainInspectorResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDomainInspectorResponseLinks returns a pointer to a new instance of NullableDomainInspectorResponseLinks.
func NewNullableDomainInspectorResponseLinks(val *DomainInspectorResponseLinks) *NullableDomainInspectorResponseLinks {
	return &NullableDomainInspectorResponseLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDomainInspectorResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDomainInspectorResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
