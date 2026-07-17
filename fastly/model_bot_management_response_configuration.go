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

// BotManagementResponseConfiguration struct for BotManagementResponseConfiguration
type BotManagementResponseConfiguration struct {
	Configuration        *BotManagementResponseConfigurationConfiguration `json:"configuration,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementResponseConfiguration BotManagementResponseConfiguration

// NewBotManagementResponseConfiguration instantiates a new BotManagementResponseConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementResponseConfiguration() *BotManagementResponseConfiguration {
	this := BotManagementResponseConfiguration{}
	return &this
}

// NewBotManagementResponseConfigurationWithDefaults instantiates a new BotManagementResponseConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementResponseConfigurationWithDefaults() *BotManagementResponseConfiguration {
	this := BotManagementResponseConfiguration{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BotManagementResponseConfiguration) GetConfiguration() BotManagementResponseConfigurationConfiguration {
	if o == nil || o.Configuration == nil {
		var ret BotManagementResponseConfigurationConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementResponseConfiguration) GetConfigurationOk() (*BotManagementResponseConfigurationConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BotManagementResponseConfiguration) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given BotManagementResponseConfigurationConfiguration and assigns it to the Configuration field.
func (o *BotManagementResponseConfiguration) SetConfiguration(v BotManagementResponseConfigurationConfiguration) {
	o.Configuration = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementResponseConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementResponseConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementResponseConfiguration := _BotManagementResponseConfiguration{}

	if err = json.Unmarshal(bytes, &varBotManagementResponseConfiguration); err == nil {
		*o = BotManagementResponseConfiguration(varBotManagementResponseConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "configuration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementResponseConfiguration is a helper abstraction for handling nullable botmanagementresponseconfiguration types.
type NullableBotManagementResponseConfiguration struct {
	value *BotManagementResponseConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementResponseConfiguration) Get() *BotManagementResponseConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementResponseConfiguration) Set(val *BotManagementResponseConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementResponseConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementResponseConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementResponseConfiguration returns a pointer to a new instance of NullableBotManagementResponseConfiguration.
func NewNullableBotManagementResponseConfiguration(val *BotManagementResponseConfiguration) *NullableBotManagementResponseConfiguration {
	return &NullableBotManagementResponseConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementResponseConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementResponseConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
