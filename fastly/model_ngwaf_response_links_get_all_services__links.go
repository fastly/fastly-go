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

// NgwafResponseLinksGetAllServicesLinks struct for NgwafResponseLinksGetAllServicesLinks
type NgwafResponseLinksGetAllServicesLinks struct {
	// Location of resource
	Self                 *string `json:"self,omitempty"`
	AdditionalProperties map[string]any
}

type _NgwafResponseLinksGetAllServicesLinks NgwafResponseLinksGetAllServicesLinks

// NewNgwafResponseLinksGetAllServicesLinks instantiates a new NgwafResponseLinksGetAllServicesLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgwafResponseLinksGetAllServicesLinks() *NgwafResponseLinksGetAllServicesLinks {
	this := NgwafResponseLinksGetAllServicesLinks{}
	return &this
}

// NewNgwafResponseLinksGetAllServicesLinksWithDefaults instantiates a new NgwafResponseLinksGetAllServicesLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgwafResponseLinksGetAllServicesLinksWithDefaults() *NgwafResponseLinksGetAllServicesLinks {
	this := NgwafResponseLinksGetAllServicesLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *NgwafResponseLinksGetAllServicesLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgwafResponseLinksGetAllServicesLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *NgwafResponseLinksGetAllServicesLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *NgwafResponseLinksGetAllServicesLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NgwafResponseLinksGetAllServicesLinks) MarshalJSON() ([]byte, error) {
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
func (o *NgwafResponseLinksGetAllServicesLinks) UnmarshalJSON(bytes []byte) (err error) {
	varNgwafResponseLinksGetAllServicesLinks := _NgwafResponseLinksGetAllServicesLinks{}

	if err = json.Unmarshal(bytes, &varNgwafResponseLinksGetAllServicesLinks); err == nil {
		*o = NgwafResponseLinksGetAllServicesLinks(varNgwafResponseLinksGetAllServicesLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNgwafResponseLinksGetAllServicesLinks is a helper abstraction for handling nullable ngwafresponselinksgetallserviceslinks types.
type NullableNgwafResponseLinksGetAllServicesLinks struct {
	value *NgwafResponseLinksGetAllServicesLinks
	isSet bool
}

// Get returns the value.
func (v NullableNgwafResponseLinksGetAllServicesLinks) Get() *NgwafResponseLinksGetAllServicesLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableNgwafResponseLinksGetAllServicesLinks) Set(val *NgwafResponseLinksGetAllServicesLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNgwafResponseLinksGetAllServicesLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNgwafResponseLinksGetAllServicesLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNgwafResponseLinksGetAllServicesLinks returns a pointer to a new instance of NullableNgwafResponseLinksGetAllServicesLinks.
func NewNullableNgwafResponseLinksGetAllServicesLinks(val *NgwafResponseLinksGetAllServicesLinks) *NullableNgwafResponseLinksGetAllServicesLinks {
	return &NullableNgwafResponseLinksGetAllServicesLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNgwafResponseLinksGetAllServicesLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNgwafResponseLinksGetAllServicesLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
