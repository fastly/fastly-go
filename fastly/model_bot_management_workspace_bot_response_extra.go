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

// BotManagementWorkspaceBotResponseExtra struct for BotManagementWorkspaceBotResponseExtra
type BotManagementWorkspaceBotResponseExtra struct {
	// The ID of the category the bot belongs to.
	CategoryID           string `json:"category_id"`
	AdditionalProperties map[string]any
}

type _BotManagementWorkspaceBotResponseExtra BotManagementWorkspaceBotResponseExtra

// NewBotManagementWorkspaceBotResponseExtra instantiates a new BotManagementWorkspaceBotResponseExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementWorkspaceBotResponseExtra(categoryID string) *BotManagementWorkspaceBotResponseExtra {
	this := BotManagementWorkspaceBotResponseExtra{}
	this.CategoryID = categoryID
	return &this
}

// NewBotManagementWorkspaceBotResponseExtraWithDefaults instantiates a new BotManagementWorkspaceBotResponseExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementWorkspaceBotResponseExtraWithDefaults() *BotManagementWorkspaceBotResponseExtra {
	this := BotManagementWorkspaceBotResponseExtra{}
	return &this
}

// GetCategoryID returns the CategoryID field value
func (o *BotManagementWorkspaceBotResponseExtra) GetCategoryID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryID
}

// GetCategoryIDOk returns a tuple with the CategoryID field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceBotResponseExtra) GetCategoryIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryID, true
}

// SetCategoryID sets field value
func (o *BotManagementWorkspaceBotResponseExtra) SetCategoryID(v string) {
	o.CategoryID = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementWorkspaceBotResponseExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *BotManagementWorkspaceBotResponseExtra) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementWorkspaceBotResponseExtra := _BotManagementWorkspaceBotResponseExtra{}

	if err = json.Unmarshal(bytes, &varBotManagementWorkspaceBotResponseExtra); err == nil {
		*o = BotManagementWorkspaceBotResponseExtra(varBotManagementWorkspaceBotResponseExtra)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "category_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementWorkspaceBotResponseExtra is a helper abstraction for handling nullable botmanagementworkspacebotresponseextra types.
type NullableBotManagementWorkspaceBotResponseExtra struct {
	value *BotManagementWorkspaceBotResponseExtra
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementWorkspaceBotResponseExtra) Get() *BotManagementWorkspaceBotResponseExtra {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementWorkspaceBotResponseExtra) Set(val *BotManagementWorkspaceBotResponseExtra) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementWorkspaceBotResponseExtra) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementWorkspaceBotResponseExtra) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementWorkspaceBotResponseExtra returns a pointer to a new instance of NullableBotManagementWorkspaceBotResponseExtra.
func NewNullableBotManagementWorkspaceBotResponseExtra(val *BotManagementWorkspaceBotResponseExtra) *NullableBotManagementWorkspaceBotResponseExtra {
	return &NullableBotManagementWorkspaceBotResponseExtra{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementWorkspaceBotResponseExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementWorkspaceBotResponseExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
