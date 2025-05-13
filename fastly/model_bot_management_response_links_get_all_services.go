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

// BotManagementResponseLinksGetAllServices struct for BotManagementResponseLinksGetAllServices
type BotManagementResponseLinksGetAllServices struct {
	Links                *BotManagementResponseLinksGetAllServicesLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementResponseLinksGetAllServices BotManagementResponseLinksGetAllServices

// NewBotManagementResponseLinksGetAllServices instantiates a new BotManagementResponseLinksGetAllServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementResponseLinksGetAllServices() *BotManagementResponseLinksGetAllServices {
	this := BotManagementResponseLinksGetAllServices{}
	return &this
}

// NewBotManagementResponseLinksGetAllServicesWithDefaults instantiates a new BotManagementResponseLinksGetAllServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementResponseLinksGetAllServicesWithDefaults() *BotManagementResponseLinksGetAllServices {
	this := BotManagementResponseLinksGetAllServices{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BotManagementResponseLinksGetAllServices) GetLinks() BotManagementResponseLinksGetAllServicesLinks {
	if o == nil || o.Links == nil {
		var ret BotManagementResponseLinksGetAllServicesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementResponseLinksGetAllServices) GetLinksOk() (*BotManagementResponseLinksGetAllServicesLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BotManagementResponseLinksGetAllServices) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given BotManagementResponseLinksGetAllServicesLinks and assigns it to the Links field.
func (o *BotManagementResponseLinksGetAllServices) SetLinks(v BotManagementResponseLinksGetAllServicesLinks) {
	o.Links = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementResponseLinksGetAllServices) MarshalJSON() ([]byte, error) {
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
func (o *BotManagementResponseLinksGetAllServices) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementResponseLinksGetAllServices := _BotManagementResponseLinksGetAllServices{}

	if err = json.Unmarshal(bytes, &varBotManagementResponseLinksGetAllServices); err == nil {
		*o = BotManagementResponseLinksGetAllServices(varBotManagementResponseLinksGetAllServices)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementResponseLinksGetAllServices is a helper abstraction for handling nullable botmanagementresponselinksgetallservices types.
type NullableBotManagementResponseLinksGetAllServices struct {
	value *BotManagementResponseLinksGetAllServices
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementResponseLinksGetAllServices) Get() *BotManagementResponseLinksGetAllServices {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementResponseLinksGetAllServices) Set(val *BotManagementResponseLinksGetAllServices) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementResponseLinksGetAllServices) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementResponseLinksGetAllServices) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementResponseLinksGetAllServices returns a pointer to a new instance of NullableBotManagementResponseLinksGetAllServices.
func NewNullableBotManagementResponseLinksGetAllServices(val *BotManagementResponseLinksGetAllServices) *NullableBotManagementResponseLinksGetAllServices {
	return &NullableBotManagementResponseLinksGetAllServices{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementResponseLinksGetAllServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementResponseLinksGetAllServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
