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

// DdosProtectionResponseLinksGetAllServicesLinks struct for DdosProtectionResponseLinksGetAllServicesLinks
type DdosProtectionResponseLinksGetAllServicesLinks struct {
	// Location of resource
	Self                 *string `json:"self,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionResponseLinksGetAllServicesLinks DdosProtectionResponseLinksGetAllServicesLinks

// NewDdosProtectionResponseLinksGetAllServicesLinks instantiates a new DdosProtectionResponseLinksGetAllServicesLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionResponseLinksGetAllServicesLinks() *DdosProtectionResponseLinksGetAllServicesLinks {
	this := DdosProtectionResponseLinksGetAllServicesLinks{}
	return &this
}

// NewDdosProtectionResponseLinksGetAllServicesLinksWithDefaults instantiates a new DdosProtectionResponseLinksGetAllServicesLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionResponseLinksGetAllServicesLinksWithDefaults() *DdosProtectionResponseLinksGetAllServicesLinks {
	this := DdosProtectionResponseLinksGetAllServicesLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *DdosProtectionResponseLinksGetAllServicesLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionResponseLinksGetAllServicesLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *DdosProtectionResponseLinksGetAllServicesLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *DdosProtectionResponseLinksGetAllServicesLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionResponseLinksGetAllServicesLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionResponseLinksGetAllServicesLinks) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionResponseLinksGetAllServicesLinks := _DdosProtectionResponseLinksGetAllServicesLinks{}

	if err = json.Unmarshal(bytes, &varDdosProtectionResponseLinksGetAllServicesLinks); err == nil {
		*o = DdosProtectionResponseLinksGetAllServicesLinks(varDdosProtectionResponseLinksGetAllServicesLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionResponseLinksGetAllServicesLinks is a helper abstraction for handling nullable ddosprotectionresponselinksgetallserviceslinks types.
type NullableDdosProtectionResponseLinksGetAllServicesLinks struct {
	value *DdosProtectionResponseLinksGetAllServicesLinks
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionResponseLinksGetAllServicesLinks) Get() *DdosProtectionResponseLinksGetAllServicesLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionResponseLinksGetAllServicesLinks) Set(val *DdosProtectionResponseLinksGetAllServicesLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionResponseLinksGetAllServicesLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionResponseLinksGetAllServicesLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionResponseLinksGetAllServicesLinks returns a pointer to a new instance of NullableDdosProtectionResponseLinksGetAllServicesLinks.
func NewNullableDdosProtectionResponseLinksGetAllServicesLinks(val *DdosProtectionResponseLinksGetAllServicesLinks) *NullableDdosProtectionResponseLinksGetAllServicesLinks {
	return &NullableDdosProtectionResponseLinksGetAllServicesLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionResponseLinksGetAllServicesLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionResponseLinksGetAllServicesLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
