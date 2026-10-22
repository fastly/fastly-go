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

// BotManagementWorkspaceBotResponse struct for BotManagementWorkspaceBotResponse
type BotManagementWorkspaceBotResponse struct {
	Action BotManagementBotAction `json:"action"`
	// The unique identifier of the bot.
	BotId string `json:"bot_id"`
	// The display name of the bot.
	Name string `json:"name"`
	// The ID of the category the bot belongs to.
	CategoryId           string `json:"category_id"`
	AdditionalProperties map[string]any
}

type _BotManagementWorkspaceBotResponse BotManagementWorkspaceBotResponse

// NewBotManagementWorkspaceBotResponse instantiates a new BotManagementWorkspaceBotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementWorkspaceBotResponse(action BotManagementBotAction, botId string, name string, categoryId string) *BotManagementWorkspaceBotResponse {
	this := BotManagementWorkspaceBotResponse{}
	this.Action = action
	this.BotId = botId
	this.Name = name
	this.CategoryId = categoryId
	return &this
}

// NewBotManagementWorkspaceBotResponseWithDefaults instantiates a new BotManagementWorkspaceBotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementWorkspaceBotResponseWithDefaults() *BotManagementWorkspaceBotResponse {
	this := BotManagementWorkspaceBotResponse{}
	return &this
}

// GetAction returns the Action field value
func (o *BotManagementWorkspaceBotResponse) GetAction() BotManagementBotAction {
	if o == nil {
		var ret BotManagementBotAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceBotResponse) GetActionOk() (*BotManagementBotAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *BotManagementWorkspaceBotResponse) SetAction(v BotManagementBotAction) {
	o.Action = v
}

// GetBotId returns the BotId field value
func (o *BotManagementWorkspaceBotResponse) GetBotId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BotId
}

// GetBotIdOk returns a tuple with the BotId field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceBotResponse) GetBotIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BotId, true
}

// SetBotId sets field value
func (o *BotManagementWorkspaceBotResponse) SetBotId(v string) {
	o.BotId = v
}

// GetName returns the Name field value
func (o *BotManagementWorkspaceBotResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceBotResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BotManagementWorkspaceBotResponse) SetName(v string) {
	o.Name = v
}

// GetCategoryId returns the CategoryId field value
func (o *BotManagementWorkspaceBotResponse) GetCategoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceBotResponse) GetCategoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *BotManagementWorkspaceBotResponse) SetCategoryId(v string) {
	o.CategoryId = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementWorkspaceBotResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["bot_id"] = o.BotId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["category_id"] = o.CategoryId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementWorkspaceBotResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementWorkspaceBotResponse := _BotManagementWorkspaceBotResponse{}

	if err = json.Unmarshal(bytes, &varBotManagementWorkspaceBotResponse); err == nil {
		*o = BotManagementWorkspaceBotResponse(varBotManagementWorkspaceBotResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "bot_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "category_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementWorkspaceBotResponse is a helper abstraction for handling nullable botmanagementworkspacebotresponse types.
type NullableBotManagementWorkspaceBotResponse struct {
	value *BotManagementWorkspaceBotResponse
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementWorkspaceBotResponse) Get() *BotManagementWorkspaceBotResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementWorkspaceBotResponse) Set(val *BotManagementWorkspaceBotResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementWorkspaceBotResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementWorkspaceBotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementWorkspaceBotResponse returns a pointer to a new instance of NullableBotManagementWorkspaceBotResponse.
func NewNullableBotManagementWorkspaceBotResponse(val *BotManagementWorkspaceBotResponse) *NullableBotManagementWorkspaceBotResponse {
	return &NullableBotManagementWorkspaceBotResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementWorkspaceBotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementWorkspaceBotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
