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

// BotManagementBotResponse struct for BotManagementBotResponse
type BotManagementBotResponse struct {
	Action BotManagementBotAction `json:"action"`
	// The unique identifier of the bot.
	BotID string `json:"bot_id"`
	// The display name of the bot.
	Name                 string `json:"name"`
	AdditionalProperties map[string]any
}

type _BotManagementBotResponse BotManagementBotResponse

// NewBotManagementBotResponse instantiates a new BotManagementBotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementBotResponse(action BotManagementBotAction, botID string, name string) *BotManagementBotResponse {
	this := BotManagementBotResponse{}
	this.Action = action
	this.BotID = botID
	this.Name = name
	return &this
}

// NewBotManagementBotResponseWithDefaults instantiates a new BotManagementBotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementBotResponseWithDefaults() *BotManagementBotResponse {
	this := BotManagementBotResponse{}
	return &this
}

// GetAction returns the Action field value
func (o *BotManagementBotResponse) GetAction() BotManagementBotAction {
	if o == nil {
		var ret BotManagementBotAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *BotManagementBotResponse) GetActionOk() (*BotManagementBotAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *BotManagementBotResponse) SetAction(v BotManagementBotAction) {
	o.Action = v
}

// GetBotID returns the BotID field value
func (o *BotManagementBotResponse) GetBotID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BotID
}

// GetBotIDOk returns a tuple with the BotID field value
// and a boolean to check if the value has been set.
func (o *BotManagementBotResponse) GetBotIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BotID, true
}

// SetBotID sets field value
func (o *BotManagementBotResponse) SetBotID(v string) {
	o.BotID = v
}

// GetName returns the Name field value
func (o *BotManagementBotResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BotManagementBotResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BotManagementBotResponse) SetName(v string) {
	o.Name = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementBotResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["action"] = o.Action
	}
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
func (o *BotManagementBotResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementBotResponse := _BotManagementBotResponse{}

	if err = json.Unmarshal(bytes, &varBotManagementBotResponse); err == nil {
		*o = BotManagementBotResponse(varBotManagementBotResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "bot_id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementBotResponse is a helper abstraction for handling nullable botmanagementbotresponse types.
type NullableBotManagementBotResponse struct {
	value *BotManagementBotResponse
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementBotResponse) Get() *BotManagementBotResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementBotResponse) Set(val *BotManagementBotResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementBotResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementBotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementBotResponse returns a pointer to a new instance of NullableBotManagementBotResponse.
func NewNullableBotManagementBotResponse(val *BotManagementBotResponse) *NullableBotManagementBotResponse {
	return &NullableBotManagementBotResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementBotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementBotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
