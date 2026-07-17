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

// BotManagementResponseConfigurationConfiguration struct for BotManagementResponseConfigurationConfiguration
type BotManagementResponseConfigurationConfiguration struct {
	// ContentGuard status
	Contentguard         *string `json:"contentguard,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementResponseConfigurationConfiguration BotManagementResponseConfigurationConfiguration

// NewBotManagementResponseConfigurationConfiguration instantiates a new BotManagementResponseConfigurationConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementResponseConfigurationConfiguration() *BotManagementResponseConfigurationConfiguration {
	this := BotManagementResponseConfigurationConfiguration{}
	return &this
}

// NewBotManagementResponseConfigurationConfigurationWithDefaults instantiates a new BotManagementResponseConfigurationConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementResponseConfigurationConfigurationWithDefaults() *BotManagementResponseConfigurationConfiguration {
	this := BotManagementResponseConfigurationConfiguration{}
	return &this
}

// GetContentguard returns the Contentguard field value if set, zero value otherwise.
func (o *BotManagementResponseConfigurationConfiguration) GetContentguard() string {
	if o == nil || o.Contentguard == nil {
		var ret string
		return ret
	}
	return *o.Contentguard
}

// GetContentguardOk returns a tuple with the Contentguard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementResponseConfigurationConfiguration) GetContentguardOk() (*string, bool) {
	if o == nil || o.Contentguard == nil {
		return nil, false
	}
	return o.Contentguard, true
}

// HasContentguard returns a boolean if a field has been set.
func (o *BotManagementResponseConfigurationConfiguration) HasContentguard() bool {
	if o != nil && o.Contentguard != nil {
		return true
	}

	return false
}

// SetContentguard gets a reference to the given string and assigns it to the Contentguard field.
func (o *BotManagementResponseConfigurationConfiguration) SetContentguard(v string) {
	o.Contentguard = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementResponseConfigurationConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Contentguard != nil {
		toSerialize["contentguard"] = o.Contentguard
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementResponseConfigurationConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementResponseConfigurationConfiguration := _BotManagementResponseConfigurationConfiguration{}

	if err = json.Unmarshal(bytes, &varBotManagementResponseConfigurationConfiguration); err == nil {
		*o = BotManagementResponseConfigurationConfiguration(varBotManagementResponseConfigurationConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contentguard")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementResponseConfigurationConfiguration is a helper abstraction for handling nullable botmanagementresponseconfigurationconfiguration types.
type NullableBotManagementResponseConfigurationConfiguration struct {
	value *BotManagementResponseConfigurationConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementResponseConfigurationConfiguration) Get() *BotManagementResponseConfigurationConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementResponseConfigurationConfiguration) Set(val *BotManagementResponseConfigurationConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementResponseConfigurationConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementResponseConfigurationConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementResponseConfigurationConfiguration returns a pointer to a new instance of NullableBotManagementResponseConfigurationConfiguration.
func NewNullableBotManagementResponseConfigurationConfiguration(val *BotManagementResponseConfigurationConfiguration) *NullableBotManagementResponseConfigurationConfiguration {
	return &NullableBotManagementResponseConfigurationConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementResponseConfigurationConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementResponseConfigurationConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
