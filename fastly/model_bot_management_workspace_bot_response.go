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
	BotID string `json:"bot_id"`
	// The display name of the bot.
	Name string `json:"name"`
	// The ID of the category the bot belongs to.
	CategoryID           string `json:"category_id"`
	AdditionalProperties map[string]any
}

type _BotManagementWorkspaceBotResponse BotManagementWorkspaceBotResponse

// NewBotManagementWorkspaceBotResponse instantiates a new BotManagementWorkspaceBotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementWorkspaceBotResponse(action BotManagementBotAction, botID string, name string, categoryID string) *BotManagementWorkspaceBotResponse {
	this := BotManagementWorkspaceBotResponse{}
	this.Action = action
	this.BotID = botID
	this.Name = name
	this.CategoryID = categoryID
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

// GetBotID returns the BotID field value
func (o *BotManagementWorkspaceBotResponse) GetBotID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BotID
}

// GetBotIDOk returns a tuple with the BotID field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceBotResponse) GetBotIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BotID, true
}

// SetBotID sets field value
func (o *BotManagementWorkspaceBotResponse) SetBotID(v string) {
	o.BotID = v
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

// GetCategoryID returns the CategoryID field value
func (o *BotManagementWorkspaceBotResponse) GetCategoryID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryID
}

// GetCategoryIDOk returns a tuple with the CategoryID field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceBotResponse) GetCategoryIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryID, true
}

// SetCategoryID sets field value
func (o *BotManagementWorkspaceBotResponse) SetCategoryID(v string) {
	o.CategoryID = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementWorkspaceBotResponse) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["category_id"] = o.CategoryID
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
