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

// BotManagementBot A single bot within a category. Bots are defined by Fastly and cannot be created or deleted.
type BotManagementBot struct {
	Action               BotManagementBotAction `json:"action"`
	AdditionalProperties map[string]any
}

type _BotManagementBot BotManagementBot

// NewBotManagementBot instantiates a new BotManagementBot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementBot(action BotManagementBotAction) *BotManagementBot {
	this := BotManagementBot{}
	this.Action = action
	return &this
}

// NewBotManagementBotWithDefaults instantiates a new BotManagementBot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementBotWithDefaults() *BotManagementBot {
	this := BotManagementBot{}
	return &this
}

// GetAction returns the Action field value
func (o *BotManagementBot) GetAction() BotManagementBotAction {
	if o == nil {
		var ret BotManagementBotAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *BotManagementBot) GetActionOk() (*BotManagementBotAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *BotManagementBot) SetAction(v BotManagementBotAction) {
	o.Action = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementBot) MarshalJSON() ([]byte, error) {
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
func (o *BotManagementBot) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementBot := _BotManagementBot{}

	if err = json.Unmarshal(bytes, &varBotManagementBot); err == nil {
		*o = BotManagementBot(varBotManagementBot)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementBot is a helper abstraction for handling nullable botmanagementbot types.
type NullableBotManagementBot struct {
	value *BotManagementBot
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementBot) Get() *BotManagementBot {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementBot) Set(val *BotManagementBot) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementBot) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementBot) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementBot returns a pointer to a new instance of NullableBotManagementBot.
func NewNullableBotManagementBot(val *BotManagementBot) *NullableBotManagementBot {
	return &NullableBotManagementBot{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementBot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementBot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
