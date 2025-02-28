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

// BotManagementResponseLinksLinks struct for BotManagementResponseLinksLinks
type BotManagementResponseLinksLinks struct {
	// Location of resource
	Self *string `json:"self,omitempty"`
	// Location of the service resource
	Service              *string `json:"service,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementResponseLinksLinks BotManagementResponseLinksLinks

// NewBotManagementResponseLinksLinks instantiates a new BotManagementResponseLinksLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementResponseLinksLinks() *BotManagementResponseLinksLinks {
	this := BotManagementResponseLinksLinks{}
	return &this
}

// NewBotManagementResponseLinksLinksWithDefaults instantiates a new BotManagementResponseLinksLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementResponseLinksLinksWithDefaults() *BotManagementResponseLinksLinks {
	this := BotManagementResponseLinksLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *BotManagementResponseLinksLinks) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementResponseLinksLinks) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *BotManagementResponseLinksLinks) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *BotManagementResponseLinksLinks) SetSelf(v string) {
	o.Self = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *BotManagementResponseLinksLinks) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementResponseLinksLinks) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *BotManagementResponseLinksLinks) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *BotManagementResponseLinksLinks) SetService(v string) {
	o.Service = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementResponseLinksLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementResponseLinksLinks) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementResponseLinksLinks := _BotManagementResponseLinksLinks{}

	if err = json.Unmarshal(bytes, &varBotManagementResponseLinksLinks); err == nil {
		*o = BotManagementResponseLinksLinks(varBotManagementResponseLinksLinks)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "service")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementResponseLinksLinks is a helper abstraction for handling nullable botmanagementresponselinkslinks types.
type NullableBotManagementResponseLinksLinks struct {
	value *BotManagementResponseLinksLinks
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementResponseLinksLinks) Get() *BotManagementResponseLinksLinks {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementResponseLinksLinks) Set(val *BotManagementResponseLinksLinks) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementResponseLinksLinks) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementResponseLinksLinks) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementResponseLinksLinks returns a pointer to a new instance of NullableBotManagementResponseLinksLinks.
func NewNullableBotManagementResponseLinksLinks(val *BotManagementResponseLinksLinks) *NullableBotManagementResponseLinksLinks {
	return &NullableBotManagementResponseLinksLinks{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementResponseLinksLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementResponseLinksLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
