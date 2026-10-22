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

// BotManagementCategory A group of bots with similar characteristics or behaviors. Categories are defined by Fastly and cannot be created or deleted.
type BotManagementCategory struct {
	Action               BotManagementAction `json:"action"`
	AdditionalProperties map[string]any
}

type _BotManagementCategory BotManagementCategory

// NewBotManagementCategory instantiates a new BotManagementCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementCategory(action BotManagementAction) *BotManagementCategory {
	this := BotManagementCategory{}
	this.Action = action
	return &this
}

// NewBotManagementCategoryWithDefaults instantiates a new BotManagementCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementCategoryWithDefaults() *BotManagementCategory {
	this := BotManagementCategory{}
	return &this
}

// GetAction returns the Action field value
func (o *BotManagementCategory) GetAction() BotManagementAction {
	if o == nil {
		var ret BotManagementAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *BotManagementCategory) GetActionOk() (*BotManagementAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *BotManagementCategory) SetAction(v BotManagementAction) {
	o.Action = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementCategory) MarshalJSON() ([]byte, error) {
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
func (o *BotManagementCategory) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementCategory := _BotManagementCategory{}

	if err = json.Unmarshal(bytes, &varBotManagementCategory); err == nil {
		*o = BotManagementCategory(varBotManagementCategory)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementCategory is a helper abstraction for handling nullable botmanagementcategory types.
type NullableBotManagementCategory struct {
	value *BotManagementCategory
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementCategory) Get() *BotManagementCategory {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementCategory) Set(val *BotManagementCategory) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementCategory) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementCategory) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementCategory returns a pointer to a new instance of NullableBotManagementCategory.
func NewNullableBotManagementCategory(val *BotManagementCategory) *NullableBotManagementCategory {
	return &NullableBotManagementCategory{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
