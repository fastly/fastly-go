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

// OriginInspectorResponseLinksGetAllServices struct for OriginInspectorResponseLinksGetAllServices
type OriginInspectorResponseLinksGetAllServices struct {
	Links                *OriginInspectorResponseLinksGetAllServicesLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _OriginInspectorResponseLinksGetAllServices OriginInspectorResponseLinksGetAllServices

// NewOriginInspectorResponseLinksGetAllServices instantiates a new OriginInspectorResponseLinksGetAllServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspectorResponseLinksGetAllServices() *OriginInspectorResponseLinksGetAllServices {
	this := OriginInspectorResponseLinksGetAllServices{}
	return &this
}

// NewOriginInspectorResponseLinksGetAllServicesWithDefaults instantiates a new OriginInspectorResponseLinksGetAllServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorResponseLinksGetAllServicesWithDefaults() *OriginInspectorResponseLinksGetAllServices {
	this := OriginInspectorResponseLinksGetAllServices{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *OriginInspectorResponseLinksGetAllServices) GetLinks() OriginInspectorResponseLinksGetAllServicesLinks {
	if o == nil || o.Links == nil {
		var ret OriginInspectorResponseLinksGetAllServicesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorResponseLinksGetAllServices) GetLinksOk() (*OriginInspectorResponseLinksGetAllServicesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *OriginInspectorResponseLinksGetAllServices) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given OriginInspectorResponseLinksGetAllServicesLinks and assigns it to the Links field.
func (o *OriginInspectorResponseLinksGetAllServices) SetLinks(v OriginInspectorResponseLinksGetAllServicesLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspectorResponseLinksGetAllServices) MarshalJSON() ([]byte, error) {
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
func (o *OriginInspectorResponseLinksGetAllServices) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspectorResponseLinksGetAllServices := _OriginInspectorResponseLinksGetAllServices{}

	if err = json.Unmarshal(bytes, &varOriginInspectorResponseLinksGetAllServices); err == nil {
		*o = OriginInspectorResponseLinksGetAllServices(varOriginInspectorResponseLinksGetAllServices)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOriginInspectorResponseLinksGetAllServices is a helper abstraction for handling nullable origininspectorresponselinksgetallservices types.
type NullableOriginInspectorResponseLinksGetAllServices struct {
	value *OriginInspectorResponseLinksGetAllServices
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspectorResponseLinksGetAllServices) Get() *OriginInspectorResponseLinksGetAllServices {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspectorResponseLinksGetAllServices) Set(val *OriginInspectorResponseLinksGetAllServices) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspectorResponseLinksGetAllServices) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspectorResponseLinksGetAllServices) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspectorResponseLinksGetAllServices returns a pointer to a new instance of NullableOriginInspectorResponseLinksGetAllServices.
func NewNullableOriginInspectorResponseLinksGetAllServices(val *OriginInspectorResponseLinksGetAllServices) *NullableOriginInspectorResponseLinksGetAllServices {
	return &NullableOriginInspectorResponseLinksGetAllServices{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspectorResponseLinksGetAllServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOriginInspectorResponseLinksGetAllServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
