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

// BotManagementAction The action applied to matching requests.
type BotManagementAction struct {
	// The type of action to apply.
	Type                 string `json:"type"`
	AdditionalProperties map[string]any
}

type _BotManagementAction BotManagementAction

// NewBotManagementAction instantiates a new BotManagementAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementAction(type_ string) *BotManagementAction {
	this := BotManagementAction{}
	this.Type = type_
	return &this
}

// NewBotManagementActionWithDefaults instantiates a new BotManagementAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementActionWithDefaults() *BotManagementAction {
	this := BotManagementAction{}
	return &this
}

// GetType returns the Type field value
func (o *BotManagementAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BotManagementAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BotManagementAction) SetType(v string) {
	o.Type = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementAction) MarshalJSON() ([]byte, error) {
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
func (o *BotManagementAction) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementAction := _BotManagementAction{}

	if err = json.Unmarshal(bytes, &varBotManagementAction); err == nil {
		*o = BotManagementAction(varBotManagementAction)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementAction is a helper abstraction for handling nullable botmanagementaction types.
type NullableBotManagementAction struct {
	value *BotManagementAction
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementAction) Get() *BotManagementAction {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementAction) Set(val *BotManagementAction) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementAction) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementAction) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementAction returns a pointer to a new instance of NullableBotManagementAction.
func NewNullableBotManagementAction(val *BotManagementAction) *NullableBotManagementAction {
	return &NullableBotManagementAction{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
