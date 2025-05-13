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

// DdosProtectionResponseLinksGetAllServices struct for DdosProtectionResponseLinksGetAllServices
type DdosProtectionResponseLinksGetAllServices struct {
	Links                *DdosProtectionResponseLinksGetAllServicesLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionResponseLinksGetAllServices DdosProtectionResponseLinksGetAllServices

// NewDdosProtectionResponseLinksGetAllServices instantiates a new DdosProtectionResponseLinksGetAllServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionResponseLinksGetAllServices() *DdosProtectionResponseLinksGetAllServices {
	this := DdosProtectionResponseLinksGetAllServices{}
	return &this
}

// NewDdosProtectionResponseLinksGetAllServicesWithDefaults instantiates a new DdosProtectionResponseLinksGetAllServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionResponseLinksGetAllServicesWithDefaults() *DdosProtectionResponseLinksGetAllServices {
	this := DdosProtectionResponseLinksGetAllServices{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DdosProtectionResponseLinksGetAllServices) GetLinks() DdosProtectionResponseLinksGetAllServicesLinks {
	if o == nil || o.Links == nil {
		var ret DdosProtectionResponseLinksGetAllServicesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionResponseLinksGetAllServices) GetLinksOk() (*DdosProtectionResponseLinksGetAllServicesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DdosProtectionResponseLinksGetAllServices) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given DdosProtectionResponseLinksGetAllServicesLinks and assigns it to the Links field.
func (o *DdosProtectionResponseLinksGetAllServices) SetLinks(v DdosProtectionResponseLinksGetAllServicesLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionResponseLinksGetAllServices) MarshalJSON() ([]byte, error) {
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
func (o *DdosProtectionResponseLinksGetAllServices) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionResponseLinksGetAllServices := _DdosProtectionResponseLinksGetAllServices{}

	if err = json.Unmarshal(bytes, &varDdosProtectionResponseLinksGetAllServices); err == nil {
		*o = DdosProtectionResponseLinksGetAllServices(varDdosProtectionResponseLinksGetAllServices)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionResponseLinksGetAllServices is a helper abstraction for handling nullable ddosprotectionresponselinksgetallservices types.
type NullableDdosProtectionResponseLinksGetAllServices struct {
	value *DdosProtectionResponseLinksGetAllServices
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionResponseLinksGetAllServices) Get() *DdosProtectionResponseLinksGetAllServices {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionResponseLinksGetAllServices) Set(val *DdosProtectionResponseLinksGetAllServices) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionResponseLinksGetAllServices) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionResponseLinksGetAllServices) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionResponseLinksGetAllServices returns a pointer to a new instance of NullableDdosProtectionResponseLinksGetAllServices.
func NewNullableDdosProtectionResponseLinksGetAllServices(val *DdosProtectionResponseLinksGetAllServices) *NullableDdosProtectionResponseLinksGetAllServices {
	return &NullableDdosProtectionResponseLinksGetAllServices{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionResponseLinksGetAllServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionResponseLinksGetAllServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
