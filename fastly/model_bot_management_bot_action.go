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

// BotManagementBotAction The action applied to requests from a single bot. Set the action to 'inherit' to apply the category's action or to a specific value to override the category's action.
type BotManagementBotAction struct {
	// The type of action to apply.
	Type                 string `json:"type"`
	AdditionalProperties map[string]any
}

type _BotManagementBotAction BotManagementBotAction

// NewBotManagementBotAction instantiates a new BotManagementBotAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementBotAction(type_ string) *BotManagementBotAction {
	this := BotManagementBotAction{}
	this.Type = type_
	return &this
}

// NewBotManagementBotActionWithDefaults instantiates a new BotManagementBotAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementBotActionWithDefaults() *BotManagementBotAction {
	this := BotManagementBotAction{}
	return &this
}

// GetType returns the Type field value
func (o *BotManagementBotAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BotManagementBotAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BotManagementBotAction) SetType(v string) {
	o.Type = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementBotAction) MarshalJSON() ([]byte, error) {
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
func (o *BotManagementBotAction) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementBotAction := _BotManagementBotAction{}

	if err = json.Unmarshal(bytes, &varBotManagementBotAction); err == nil {
		*o = BotManagementBotAction(varBotManagementBotAction)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementBotAction is a helper abstraction for handling nullable botmanagementbotaction types.
type NullableBotManagementBotAction struct {
	value *BotManagementBotAction
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementBotAction) Get() *BotManagementBotAction {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementBotAction) Set(val *BotManagementBotAction) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementBotAction) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementBotAction) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementBotAction returns a pointer to a new instance of NullableBotManagementBotAction.
func NewNullableBotManagementBotAction(val *BotManagementBotAction) *NullableBotManagementBotAction {
	return &NullableBotManagementBotAction{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementBotAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementBotAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
