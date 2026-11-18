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

// BotManagementCategoryBotsAction The action to apply to every bot in a category.
type BotManagementCategoryBotsAction struct {
	// The type of action to apply. Only `inherit` is supported.
	Type                 string `json:"type"`
	AdditionalProperties map[string]any
}

type _BotManagementCategoryBotsAction BotManagementCategoryBotsAction

// NewBotManagementCategoryBotsAction instantiates a new BotManagementCategoryBotsAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementCategoryBotsAction(resourceType string) *BotManagementCategoryBotsAction {
	this := BotManagementCategoryBotsAction{}
	this.Type = resourceType
	return &this
}

// NewBotManagementCategoryBotsActionWithDefaults instantiates a new BotManagementCategoryBotsAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementCategoryBotsActionWithDefaults() *BotManagementCategoryBotsAction {
	this := BotManagementCategoryBotsAction{}
	return &this
}

// GetType returns the Type field value
func (o *BotManagementCategoryBotsAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BotManagementCategoryBotsAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BotManagementCategoryBotsAction) SetType(v string) {
	o.Type = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementCategoryBotsAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementCategoryBotsAction) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementCategoryBotsAction := _BotManagementCategoryBotsAction{}

	if err = json.Unmarshal(bytes, &varBotManagementCategoryBotsAction); err == nil {
		*o = BotManagementCategoryBotsAction(varBotManagementCategoryBotsAction)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementCategoryBotsAction is a helper abstraction for handling nullable botmanagementcategorybotsaction types.
type NullableBotManagementCategoryBotsAction struct {
	value *BotManagementCategoryBotsAction
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementCategoryBotsAction) Get() *BotManagementCategoryBotsAction {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementCategoryBotsAction) Set(val *BotManagementCategoryBotsAction) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementCategoryBotsAction) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementCategoryBotsAction) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementCategoryBotsAction returns a pointer to a new instance of NullableBotManagementCategoryBotsAction.
func NewNullableBotManagementCategoryBotsAction(val *BotManagementCategoryBotsAction) *NullableBotManagementCategoryBotsAction {
	return &NullableBotManagementCategoryBotsAction{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementCategoryBotsAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementCategoryBotsAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
