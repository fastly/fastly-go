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

// BotManagementInvalidRequestErrorExtra struct for BotManagementInvalidRequestErrorExtra
type BotManagementInvalidRequestErrorExtra struct {
	Title                string `json:"title"`
	Status               int32  `json:"status"`
	Detail               string `json:"detail"`
	AdditionalProperties map[string]any
}

type _BotManagementInvalidRequestErrorExtra BotManagementInvalidRequestErrorExtra

// NewBotManagementInvalidRequestErrorExtra instantiates a new BotManagementInvalidRequestErrorExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementInvalidRequestErrorExtra(title string, status int32, detail string) *BotManagementInvalidRequestErrorExtra {
	this := BotManagementInvalidRequestErrorExtra{}
	this.Title = title
	this.Status = status
	this.Detail = detail
	return &this
}

// NewBotManagementInvalidRequestErrorExtraWithDefaults instantiates a new BotManagementInvalidRequestErrorExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementInvalidRequestErrorExtraWithDefaults() *BotManagementInvalidRequestErrorExtra {
	this := BotManagementInvalidRequestErrorExtra{}
	return &this
}

// GetTitle returns the Title field value
func (o *BotManagementInvalidRequestErrorExtra) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *BotManagementInvalidRequestErrorExtra) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *BotManagementInvalidRequestErrorExtra) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *BotManagementInvalidRequestErrorExtra) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BotManagementInvalidRequestErrorExtra) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BotManagementInvalidRequestErrorExtra) SetStatus(v int32) {
	o.Status = v
}

// GetDetail returns the Detail field value
func (o *BotManagementInvalidRequestErrorExtra) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *BotManagementInvalidRequestErrorExtra) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *BotManagementInvalidRequestErrorExtra) SetDetail(v string) {
	o.Detail = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementInvalidRequestErrorExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["detail"] = o.Detail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementInvalidRequestErrorExtra) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementInvalidRequestErrorExtra := _BotManagementInvalidRequestErrorExtra{}

	if err = json.Unmarshal(bytes, &varBotManagementInvalidRequestErrorExtra); err == nil {
		*o = BotManagementInvalidRequestErrorExtra(varBotManagementInvalidRequestErrorExtra)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "status")
		delete(additionalProperties, "detail")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementInvalidRequestErrorExtra is a helper abstraction for handling nullable botmanagementinvalidrequesterrorextra types.
type NullableBotManagementInvalidRequestErrorExtra struct {
	value *BotManagementInvalidRequestErrorExtra
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementInvalidRequestErrorExtra) Get() *BotManagementInvalidRequestErrorExtra {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementInvalidRequestErrorExtra) Set(val *BotManagementInvalidRequestErrorExtra) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementInvalidRequestErrorExtra) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementInvalidRequestErrorExtra) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementInvalidRequestErrorExtra returns a pointer to a new instance of NullableBotManagementInvalidRequestErrorExtra.
func NewNullableBotManagementInvalidRequestErrorExtra(val *BotManagementInvalidRequestErrorExtra) *NullableBotManagementInvalidRequestErrorExtra {
	return &NullableBotManagementInvalidRequestErrorExtra{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementInvalidRequestErrorExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementInvalidRequestErrorExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
