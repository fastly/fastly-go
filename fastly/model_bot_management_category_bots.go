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

// BotManagementCategoryBots An action to apply to every bot in a category.
type BotManagementCategoryBots struct {
	Action               BotManagementCategoryBotsAction `json:"action"`
	AdditionalProperties map[string]any
}

type _BotManagementCategoryBots BotManagementCategoryBots

// NewBotManagementCategoryBots instantiates a new BotManagementCategoryBots object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementCategoryBots(action BotManagementCategoryBotsAction) *BotManagementCategoryBots {
	this := BotManagementCategoryBots{}
	this.Action = action
	return &this
}

// NewBotManagementCategoryBotsWithDefaults instantiates a new BotManagementCategoryBots object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementCategoryBotsWithDefaults() *BotManagementCategoryBots {
	this := BotManagementCategoryBots{}
	return &this
}

// GetAction returns the Action field value
func (o *BotManagementCategoryBots) GetAction() BotManagementCategoryBotsAction {
	if o == nil {
		var ret BotManagementCategoryBotsAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *BotManagementCategoryBots) GetActionOk() (*BotManagementCategoryBotsAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *BotManagementCategoryBots) SetAction(v BotManagementCategoryBotsAction) {
	o.Action = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementCategoryBots) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["action"] = o.Action
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementCategoryBots) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementCategoryBots := _BotManagementCategoryBots{}

	if err = json.Unmarshal(bytes, &varBotManagementCategoryBots); err == nil {
		*o = BotManagementCategoryBots(varBotManagementCategoryBots)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementCategoryBots is a helper abstraction for handling nullable botmanagementcategorybots types.
type NullableBotManagementCategoryBots struct {
	value *BotManagementCategoryBots
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementCategoryBots) Get() *BotManagementCategoryBots {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementCategoryBots) Set(val *BotManagementCategoryBots) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementCategoryBots) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementCategoryBots) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementCategoryBots returns a pointer to a new instance of NullableBotManagementCategoryBots.
func NewNullableBotManagementCategoryBots(val *BotManagementCategoryBots) *NullableBotManagementCategoryBots {
	return &NullableBotManagementCategoryBots{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementCategoryBots) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementCategoryBots) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
