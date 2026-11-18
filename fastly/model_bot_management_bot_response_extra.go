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

// BotManagementBotResponseExtra struct for BotManagementBotResponseExtra
type BotManagementBotResponseExtra struct {
	// The unique identifier of the bot.
	BotID string `json:"bot_id"`
	// The display name of the bot.
	Name                 string `json:"name"`
	AdditionalProperties map[string]any
}

type _BotManagementBotResponseExtra BotManagementBotResponseExtra

// NewBotManagementBotResponseExtra instantiates a new BotManagementBotResponseExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementBotResponseExtra(botID string, name string) *BotManagementBotResponseExtra {
	this := BotManagementBotResponseExtra{}
	this.BotID = botID
	this.Name = name
	return &this
}

// NewBotManagementBotResponseExtraWithDefaults instantiates a new BotManagementBotResponseExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementBotResponseExtraWithDefaults() *BotManagementBotResponseExtra {
	this := BotManagementBotResponseExtra{}
	return &this
}

// GetBotID returns the BotID field value
func (o *BotManagementBotResponseExtra) GetBotID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BotID
}

// GetBotIDOk returns a tuple with the BotID field value
// and a boolean to check if the value has been set.
func (o *BotManagementBotResponseExtra) GetBotIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BotID, true
}

// SetBotID sets field value
func (o *BotManagementBotResponseExtra) SetBotID(v string) {
	o.BotID = v
}

// GetName returns the Name field value
func (o *BotManagementBotResponseExtra) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BotManagementBotResponseExtra) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BotManagementBotResponseExtra) SetName(v string) {
	o.Name = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementBotResponseExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["bot_id"] = o.BotID
	}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementBotResponseExtra) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementBotResponseExtra := _BotManagementBotResponseExtra{}

	if err = json.Unmarshal(bytes, &varBotManagementBotResponseExtra); err == nil {
		*o = BotManagementBotResponseExtra(varBotManagementBotResponseExtra)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bot_id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementBotResponseExtra is a helper abstraction for handling nullable botmanagementbotresponseextra types.
type NullableBotManagementBotResponseExtra struct {
	value *BotManagementBotResponseExtra
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementBotResponseExtra) Get() *BotManagementBotResponseExtra {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementBotResponseExtra) Set(val *BotManagementBotResponseExtra) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementBotResponseExtra) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementBotResponseExtra) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementBotResponseExtra returns a pointer to a new instance of NullableBotManagementBotResponseExtra.
func NewNullableBotManagementBotResponseExtra(val *BotManagementBotResponseExtra) *NullableBotManagementBotResponseExtra {
	return &NullableBotManagementBotResponseExtra{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementBotResponseExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementBotResponseExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
