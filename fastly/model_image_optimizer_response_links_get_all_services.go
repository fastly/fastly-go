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

// ImageOptimizerResponseLinksGetAllServices struct for ImageOptimizerResponseLinksGetAllServices
type ImageOptimizerResponseLinksGetAllServices struct {
	Links                *ImageOptimizerResponseLinksGetAllServicesLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _ImageOptimizerResponseLinksGetAllServices ImageOptimizerResponseLinksGetAllServices

// NewImageOptimizerResponseLinksGetAllServices instantiates a new ImageOptimizerResponseLinksGetAllServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageOptimizerResponseLinksGetAllServices() *ImageOptimizerResponseLinksGetAllServices {
	this := ImageOptimizerResponseLinksGetAllServices{}
	return &this
}

// NewImageOptimizerResponseLinksGetAllServicesWithDefaults instantiates a new ImageOptimizerResponseLinksGetAllServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageOptimizerResponseLinksGetAllServicesWithDefaults() *ImageOptimizerResponseLinksGetAllServices {
	this := ImageOptimizerResponseLinksGetAllServices{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ImageOptimizerResponseLinksGetAllServices) GetLinks() ImageOptimizerResponseLinksGetAllServicesLinks {
	if o == nil || o.Links == nil {
		var ret ImageOptimizerResponseLinksGetAllServicesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageOptimizerResponseLinksGetAllServices) GetLinksOk() (*ImageOptimizerResponseLinksGetAllServicesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ImageOptimizerResponseLinksGetAllServices) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ImageOptimizerResponseLinksGetAllServicesLinks and assigns it to the Links field.
func (o *ImageOptimizerResponseLinksGetAllServices) SetLinks(v ImageOptimizerResponseLinksGetAllServicesLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ImageOptimizerResponseLinksGetAllServices) MarshalJSON() ([]byte, error) {
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
func (o *ImageOptimizerResponseLinksGetAllServices) UnmarshalJSON(bytes []byte) (err error) {
	varImageOptimizerResponseLinksGetAllServices := _ImageOptimizerResponseLinksGetAllServices{}

	if err = json.Unmarshal(bytes, &varImageOptimizerResponseLinksGetAllServices); err == nil {
		*o = ImageOptimizerResponseLinksGetAllServices(varImageOptimizerResponseLinksGetAllServices)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableImageOptimizerResponseLinksGetAllServices is a helper abstraction for handling nullable imageoptimizerresponselinksgetallservices types.
type NullableImageOptimizerResponseLinksGetAllServices struct {
	value *ImageOptimizerResponseLinksGetAllServices
	isSet bool
}

// Get returns the value.
func (v NullableImageOptimizerResponseLinksGetAllServices) Get() *ImageOptimizerResponseLinksGetAllServices {
	return v.value
}

// Set modifies the value.
func (v *NullableImageOptimizerResponseLinksGetAllServices) Set(val *ImageOptimizerResponseLinksGetAllServices) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableImageOptimizerResponseLinksGetAllServices) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableImageOptimizerResponseLinksGetAllServices) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableImageOptimizerResponseLinksGetAllServices returns a pointer to a new instance of NullableImageOptimizerResponseLinksGetAllServices.
func NewNullableImageOptimizerResponseLinksGetAllServices(val *ImageOptimizerResponseLinksGetAllServices) *NullableImageOptimizerResponseLinksGetAllServices {
	return &NullableImageOptimizerResponseLinksGetAllServices{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableImageOptimizerResponseLinksGetAllServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableImageOptimizerResponseLinksGetAllServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
