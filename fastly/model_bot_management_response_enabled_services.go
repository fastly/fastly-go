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

// BotManagementResponseEnabledServices struct for BotManagementResponseEnabledServices
type BotManagementResponseEnabledServices struct {
	// A list of services with Bot Management enabled.
	Services             []string `json:"services,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementResponseEnabledServices BotManagementResponseEnabledServices

// NewBotManagementResponseEnabledServices instantiates a new BotManagementResponseEnabledServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementResponseEnabledServices() *BotManagementResponseEnabledServices {
	this := BotManagementResponseEnabledServices{}
	return &this
}

// NewBotManagementResponseEnabledServicesWithDefaults instantiates a new BotManagementResponseEnabledServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementResponseEnabledServicesWithDefaults() *BotManagementResponseEnabledServices {
	this := BotManagementResponseEnabledServices{}
	return &this
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *BotManagementResponseEnabledServices) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementResponseEnabledServices) GetServicesOk() ([]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *BotManagementResponseEnabledServices) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *BotManagementResponseEnabledServices) SetServices(v []string) {
	o.Services = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementResponseEnabledServices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementResponseEnabledServices) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementResponseEnabledServices := _BotManagementResponseEnabledServices{}

	if err = json.Unmarshal(bytes, &varBotManagementResponseEnabledServices); err == nil {
		*o = BotManagementResponseEnabledServices(varBotManagementResponseEnabledServices)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "services")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementResponseEnabledServices is a helper abstraction for handling nullable botmanagementresponseenabledservices types.
type NullableBotManagementResponseEnabledServices struct {
	value *BotManagementResponseEnabledServices
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementResponseEnabledServices) Get() *BotManagementResponseEnabledServices {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementResponseEnabledServices) Set(val *BotManagementResponseEnabledServices) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementResponseEnabledServices) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementResponseEnabledServices) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementResponseEnabledServices returns a pointer to a new instance of NullableBotManagementResponseEnabledServices.
func NewNullableBotManagementResponseEnabledServices(val *BotManagementResponseEnabledServices) *NullableBotManagementResponseEnabledServices {
	return &NullableBotManagementResponseEnabledServices{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementResponseEnabledServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementResponseEnabledServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
