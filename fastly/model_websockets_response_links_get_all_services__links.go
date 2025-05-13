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

// WebsocketsResponseLinksGetAllServicesLinks struct for WebsocketsResponseLinksGetAllServicesLinks
type WebsocketsResponseLinksGetAllServicesLinks struct {
	// Location of resource
	Self                 *string `json:"self,omitempty"`
	AdditionalProperties map[string]any
}

type _WebsocketsResponseLinksGetAllServicesLinks WebsocketsResponseLinksGetAllServicesLinks

// NewWebsocketsResponseLinksGetAllServicesLinks instantiates a new WebsocketsResponseLinksGetAllServicesLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebsocketsResponseLinksGetAllServicesLinks() *WebsocketsResponseLinksGetAllServicesLinks {
	this := WebsocketsResponseLinksGetAllServicesLinks{}
	return &this
}

// NewWebsocketsResponseLinksGetAllServicesLinksWithDefaults instantiates a new WebsocketsResponseLinksGetAllServicesLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebsocketsResponseLinksGetAllServicesLinksWithDefaults() *WebsocketsResponseLinksGetAllServicesLinks {
	this := WebsocketsResponseLinksGetAllServicesLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *WebsocketsResponseLinksGetAllServicesLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebsocketsResponseLinksGetAllServicesLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *WebsocketsResponseLinksGetAllServicesLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *WebsocketsResponseLinksGetAllServicesLinks) SetSelf(v string) {
	o.Self = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WebsocketsResponseLinksGetAllServicesLinks) MarshalJSON() ([]byte, error) {
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
func (o *WebsocketsResponseLinksGetAllServicesLinks) UnmarshalJSON(bytes []byte) (err error) {
	varWebsocketsResponseLinksGetAllServicesLinks := _WebsocketsResponseLinksGetAllServicesLinks{}

	if err = json.Unmarshal(bytes, &varWebsocketsResponseLinksGetAllServicesLinks); err == nil {
		*o = WebsocketsResponseLinksGetAllServicesLinks(varWebsocketsResponseLinksGetAllServicesLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWebsocketsResponseLinksGetAllServicesLinks is a helper abstraction for handling nullable websocketsresponselinksgetallserviceslinks types.
type NullableWebsocketsResponseLinksGetAllServicesLinks struct {
	value *WebsocketsResponseLinksGetAllServicesLinks
	isSet bool
}

// Get returns the value.
func (v NullableWebsocketsResponseLinksGetAllServicesLinks) Get() *WebsocketsResponseLinksGetAllServicesLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableWebsocketsResponseLinksGetAllServicesLinks) Set(val *WebsocketsResponseLinksGetAllServicesLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWebsocketsResponseLinksGetAllServicesLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWebsocketsResponseLinksGetAllServicesLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWebsocketsResponseLinksGetAllServicesLinks returns a pointer to a new instance of NullableWebsocketsResponseLinksGetAllServicesLinks.
func NewNullableWebsocketsResponseLinksGetAllServicesLinks(val *WebsocketsResponseLinksGetAllServicesLinks) *NullableWebsocketsResponseLinksGetAllServicesLinks {
	return &NullableWebsocketsResponseLinksGetAllServicesLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWebsocketsResponseLinksGetAllServicesLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWebsocketsResponseLinksGetAllServicesLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
