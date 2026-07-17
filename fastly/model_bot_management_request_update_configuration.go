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

// BotManagementRequestUpdateConfiguration struct for BotManagementRequestUpdateConfiguration
type BotManagementRequestUpdateConfiguration struct {
	// ContentGuard status
	Contentguard         string `json:"contentguard"`
	AdditionalProperties map[string]any
}

type _BotManagementRequestUpdateConfiguration BotManagementRequestUpdateConfiguration

// NewBotManagementRequestUpdateConfiguration instantiates a new BotManagementRequestUpdateConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementRequestUpdateConfiguration(contentguard string) *BotManagementRequestUpdateConfiguration {
	this := BotManagementRequestUpdateConfiguration{}
	this.Contentguard = contentguard
	return &this
}

// NewBotManagementRequestUpdateConfigurationWithDefaults instantiates a new BotManagementRequestUpdateConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementRequestUpdateConfigurationWithDefaults() *BotManagementRequestUpdateConfiguration {
	this := BotManagementRequestUpdateConfiguration{}
	return &this
}

// GetContentguard returns the Contentguard field value
func (o *BotManagementRequestUpdateConfiguration) GetContentguard() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Contentguard
}

// GetContentguardOk returns a tuple with the Contentguard field value
// and a boolean to check if the value has been set.
func (o *BotManagementRequestUpdateConfiguration) GetContentguardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contentguard, true
}

// SetContentguard sets field value
func (o *BotManagementRequestUpdateConfiguration) SetContentguard(v string) {
	o.Contentguard = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementRequestUpdateConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["contentguard"] = o.Contentguard
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementRequestUpdateConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementRequestUpdateConfiguration := _BotManagementRequestUpdateConfiguration{}

	if err = json.Unmarshal(bytes, &varBotManagementRequestUpdateConfiguration); err == nil {
		*o = BotManagementRequestUpdateConfiguration(varBotManagementRequestUpdateConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "contentguard")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementRequestUpdateConfiguration is a helper abstraction for handling nullable botmanagementrequestupdateconfiguration types.
type NullableBotManagementRequestUpdateConfiguration struct {
	value *BotManagementRequestUpdateConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementRequestUpdateConfiguration) Get() *BotManagementRequestUpdateConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementRequestUpdateConfiguration) Set(val *BotManagementRequestUpdateConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementRequestUpdateConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementRequestUpdateConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementRequestUpdateConfiguration returns a pointer to a new instance of NullableBotManagementRequestUpdateConfiguration.
func NewNullableBotManagementRequestUpdateConfiguration(val *BotManagementRequestUpdateConfiguration) *NullableBotManagementRequestUpdateConfiguration {
	return &NullableBotManagementRequestUpdateConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementRequestUpdateConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementRequestUpdateConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
