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

// FanoutResponseLinksGetAllServicesLinks struct for FanoutResponseLinksGetAllServicesLinks
type FanoutResponseLinksGetAllServicesLinks struct {
	// Location of resource
	Self                 *string `json:"self,omitempty"`
	AdditionalProperties map[string]any
}

type _FanoutResponseLinksGetAllServicesLinks FanoutResponseLinksGetAllServicesLinks

// NewFanoutResponseLinksGetAllServicesLinks instantiates a new FanoutResponseLinksGetAllServicesLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFanoutResponseLinksGetAllServicesLinks() *FanoutResponseLinksGetAllServicesLinks {
	this := FanoutResponseLinksGetAllServicesLinks{}
	return &this
}

// NewFanoutResponseLinksGetAllServicesLinksWithDefaults instantiates a new FanoutResponseLinksGetAllServicesLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFanoutResponseLinksGetAllServicesLinksWithDefaults() *FanoutResponseLinksGetAllServicesLinks {
	this := FanoutResponseLinksGetAllServicesLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *FanoutResponseLinksGetAllServicesLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FanoutResponseLinksGetAllServicesLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *FanoutResponseLinksGetAllServicesLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *FanoutResponseLinksGetAllServicesLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o FanoutResponseLinksGetAllServicesLinks) MarshalJSON() ([]byte, error) {
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
func (o *FanoutResponseLinksGetAllServicesLinks) UnmarshalJSON(bytes []byte) (err error) {
	varFanoutResponseLinksGetAllServicesLinks := _FanoutResponseLinksGetAllServicesLinks{}

	if err = json.Unmarshal(bytes, &varFanoutResponseLinksGetAllServicesLinks); err == nil {
		*o = FanoutResponseLinksGetAllServicesLinks(varFanoutResponseLinksGetAllServicesLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableFanoutResponseLinksGetAllServicesLinks is a helper abstraction for handling nullable fanoutresponselinksgetallserviceslinks types.
type NullableFanoutResponseLinksGetAllServicesLinks struct {
	value *FanoutResponseLinksGetAllServicesLinks
	isSet bool
}

// Get returns the value.
func (v NullableFanoutResponseLinksGetAllServicesLinks) Get() *FanoutResponseLinksGetAllServicesLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableFanoutResponseLinksGetAllServicesLinks) Set(val *FanoutResponseLinksGetAllServicesLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableFanoutResponseLinksGetAllServicesLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableFanoutResponseLinksGetAllServicesLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableFanoutResponseLinksGetAllServicesLinks returns a pointer to a new instance of NullableFanoutResponseLinksGetAllServicesLinks.
func NewNullableFanoutResponseLinksGetAllServicesLinks(val *FanoutResponseLinksGetAllServicesLinks) *NullableFanoutResponseLinksGetAllServicesLinks {
	return &NullableFanoutResponseLinksGetAllServicesLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableFanoutResponseLinksGetAllServicesLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableFanoutResponseLinksGetAllServicesLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
