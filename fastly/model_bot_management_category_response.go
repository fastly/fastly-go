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

// BotManagementCategoryResponse struct for BotManagementCategoryResponse
type BotManagementCategoryResponse struct {
	Action BotManagementAction `json:"action"`
	// The unique identifier of the category.
	CategoryID string `json:"category_id"`
	// The display name of the category.
	Name string `json:"name"`
	// A description of the bots the category contains.
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementCategoryResponse BotManagementCategoryResponse

// NewBotManagementCategoryResponse instantiates a new BotManagementCategoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementCategoryResponse(action BotManagementAction, categoryID string, name string) *BotManagementCategoryResponse {
	this := BotManagementCategoryResponse{}
	this.Action = action
	this.CategoryID = categoryID
	this.Name = name
	return &this
}

// NewBotManagementCategoryResponseWithDefaults instantiates a new BotManagementCategoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementCategoryResponseWithDefaults() *BotManagementCategoryResponse {
	this := BotManagementCategoryResponse{}
	return &this
}

// GetAction returns the Action field value
func (o *BotManagementCategoryResponse) GetAction() BotManagementAction {
	if o == nil {
		var ret BotManagementAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *BotManagementCategoryResponse) GetActionOk() (*BotManagementAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *BotManagementCategoryResponse) SetAction(v BotManagementAction) {
	o.Action = v
}

// GetCategoryID returns the CategoryID field value
func (o *BotManagementCategoryResponse) GetCategoryID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryID
}

// GetCategoryIDOk returns a tuple with the CategoryID field value
// and a boolean to check if the value has been set.
func (o *BotManagementCategoryResponse) GetCategoryIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryID, true
}

// SetCategoryID sets field value
func (o *BotManagementCategoryResponse) SetCategoryID(v string) {
	o.CategoryID = v
}

// GetName returns the Name field value
func (o *BotManagementCategoryResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BotManagementCategoryResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BotManagementCategoryResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BotManagementCategoryResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementCategoryResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BotManagementCategoryResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BotManagementCategoryResponse) SetDescription(v string) {
	o.Description = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementCategoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["category_id"] = o.CategoryID
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementCategoryResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementCategoryResponse := _BotManagementCategoryResponse{}

	if err = json.Unmarshal(bytes, &varBotManagementCategoryResponse); err == nil {
		*o = BotManagementCategoryResponse(varBotManagementCategoryResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "category_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementCategoryResponse is a helper abstraction for handling nullable botmanagementcategoryresponse types.
type NullableBotManagementCategoryResponse struct {
	value *BotManagementCategoryResponse
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementCategoryResponse) Get() *BotManagementCategoryResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementCategoryResponse) Set(val *BotManagementCategoryResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementCategoryResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementCategoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementCategoryResponse returns a pointer to a new instance of NullableBotManagementCategoryResponse.
func NewNullableBotManagementCategoryResponse(val *BotManagementCategoryResponse) *NullableBotManagementCategoryResponse {
	return &NullableBotManagementCategoryResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementCategoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementCategoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
