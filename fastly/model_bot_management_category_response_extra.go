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

// BotManagementCategoryResponseExtra struct for BotManagementCategoryResponseExtra
type BotManagementCategoryResponseExtra struct {
	// The unique identifier of the category.
	CategoryId string `json:"category_id"`
	// The display name of the category.
	Name string `json:"name"`
	// A description of the bots the category contains.
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementCategoryResponseExtra BotManagementCategoryResponseExtra

// NewBotManagementCategoryResponseExtra instantiates a new BotManagementCategoryResponseExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementCategoryResponseExtra(categoryId string, name string) *BotManagementCategoryResponseExtra {
	this := BotManagementCategoryResponseExtra{}
	this.CategoryId = categoryId
	this.Name = name
	return &this
}

// NewBotManagementCategoryResponseExtraWithDefaults instantiates a new BotManagementCategoryResponseExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementCategoryResponseExtraWithDefaults() *BotManagementCategoryResponseExtra {
	this := BotManagementCategoryResponseExtra{}
	return &this
}

// GetCategoryId returns the CategoryId field value
func (o *BotManagementCategoryResponseExtra) GetCategoryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value
// and a boolean to check if the value has been set.
func (o *BotManagementCategoryResponseExtra) GetCategoryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CategoryId, true
}

// SetCategoryId sets field value
func (o *BotManagementCategoryResponseExtra) SetCategoryId(v string) {
	o.CategoryId = v
}

// GetName returns the Name field value
func (o *BotManagementCategoryResponseExtra) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BotManagementCategoryResponseExtra) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BotManagementCategoryResponseExtra) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BotManagementCategoryResponseExtra) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementCategoryResponseExtra) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BotManagementCategoryResponseExtra) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BotManagementCategoryResponseExtra) SetDescription(v string) {
	o.Description = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementCategoryResponseExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["category_id"] = o.CategoryId
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
func (o *BotManagementCategoryResponseExtra) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementCategoryResponseExtra := _BotManagementCategoryResponseExtra{}

	if err = json.Unmarshal(bytes, &varBotManagementCategoryResponseExtra); err == nil {
		*o = BotManagementCategoryResponseExtra(varBotManagementCategoryResponseExtra)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "category_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementCategoryResponseExtra is a helper abstraction for handling nullable botmanagementcategoryresponseextra types.
type NullableBotManagementCategoryResponseExtra struct {
	value *BotManagementCategoryResponseExtra
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementCategoryResponseExtra) Get() *BotManagementCategoryResponseExtra {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementCategoryResponseExtra) Set(val *BotManagementCategoryResponseExtra) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementCategoryResponseExtra) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementCategoryResponseExtra) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementCategoryResponseExtra returns a pointer to a new instance of NullableBotManagementCategoryResponseExtra.
func NewNullableBotManagementCategoryResponseExtra(val *BotManagementCategoryResponseExtra) *NullableBotManagementCategoryResponseExtra {
	return &NullableBotManagementCategoryResponseExtra{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementCategoryResponseExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementCategoryResponseExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
