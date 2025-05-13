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

// FanoutResponseLinksGetAllServices struct for FanoutResponseLinksGetAllServices
type FanoutResponseLinksGetAllServices struct {
	Links                *FanoutResponseLinksGetAllServicesLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _FanoutResponseLinksGetAllServices FanoutResponseLinksGetAllServices

// NewFanoutResponseLinksGetAllServices instantiates a new FanoutResponseLinksGetAllServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFanoutResponseLinksGetAllServices() *FanoutResponseLinksGetAllServices {
	this := FanoutResponseLinksGetAllServices{}
	return &this
}

// NewFanoutResponseLinksGetAllServicesWithDefaults instantiates a new FanoutResponseLinksGetAllServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFanoutResponseLinksGetAllServicesWithDefaults() *FanoutResponseLinksGetAllServices {
	this := FanoutResponseLinksGetAllServices{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *FanoutResponseLinksGetAllServices) GetLinks() FanoutResponseLinksGetAllServicesLinks {
	if o == nil || o.Links == nil {
		var ret FanoutResponseLinksGetAllServicesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanoutResponseLinksGetAllServices) GetLinksOk() (*FanoutResponseLinksGetAllServicesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *FanoutResponseLinksGetAllServices) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given FanoutResponseLinksGetAllServicesLinks and assigns it to the Links field.
func (o *FanoutResponseLinksGetAllServices) SetLinks(v FanoutResponseLinksGetAllServicesLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o FanoutResponseLinksGetAllServices) MarshalJSON() ([]byte, error) {
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
func (o *FanoutResponseLinksGetAllServices) UnmarshalJSON(bytes []byte) (err error) {
	varFanoutResponseLinksGetAllServices := _FanoutResponseLinksGetAllServices{}

	if err = json.Unmarshal(bytes, &varFanoutResponseLinksGetAllServices); err == nil {
		*o = FanoutResponseLinksGetAllServices(varFanoutResponseLinksGetAllServices)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableFanoutResponseLinksGetAllServices is a helper abstraction for handling nullable fanoutresponselinksgetallservices types.
type NullableFanoutResponseLinksGetAllServices struct {
	value *FanoutResponseLinksGetAllServices
	isSet bool
}

// Get returns the value.
func (v NullableFanoutResponseLinksGetAllServices) Get() *FanoutResponseLinksGetAllServices {
	return v.value
}

// Set modifies the value.
func (v *NullableFanoutResponseLinksGetAllServices) Set(val *FanoutResponseLinksGetAllServices) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableFanoutResponseLinksGetAllServices) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableFanoutResponseLinksGetAllServices) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFanoutResponseLinksGetAllServices returns a pointer to a new instance of NullableFanoutResponseLinksGetAllServices.
func NewNullableFanoutResponseLinksGetAllServices(val *FanoutResponseLinksGetAllServices) *NullableFanoutResponseLinksGetAllServices {
	return &NullableFanoutResponseLinksGetAllServices{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableFanoutResponseLinksGetAllServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableFanoutResponseLinksGetAllServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
