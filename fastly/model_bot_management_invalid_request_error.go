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

// BotManagementInvalidRequestError struct for BotManagementInvalidRequestError
type BotManagementInvalidRequestError struct {
	Title                string `json:"title"`
	Status               int32  `json:"status"`
	Detail               string `json:"detail"`
	AdditionalProperties map[string]any
}

type _BotManagementInvalidRequestError BotManagementInvalidRequestError

// NewBotManagementInvalidRequestError instantiates a new BotManagementInvalidRequestError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementInvalidRequestError(title string, status int32, detail string) *BotManagementInvalidRequestError {
	this := BotManagementInvalidRequestError{}
	this.Title = title
	this.Status = status
	this.Detail = detail
	return &this
}

// NewBotManagementInvalidRequestErrorWithDefaults instantiates a new BotManagementInvalidRequestError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementInvalidRequestErrorWithDefaults() *BotManagementInvalidRequestError {
	this := BotManagementInvalidRequestError{}
	return &this
}

// GetTitle returns the Title field value
func (o *BotManagementInvalidRequestError) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *BotManagementInvalidRequestError) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *BotManagementInvalidRequestError) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *BotManagementInvalidRequestError) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BotManagementInvalidRequestError) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BotManagementInvalidRequestError) SetStatus(v int32) {
	o.Status = v
}

// GetDetail returns the Detail field value
func (o *BotManagementInvalidRequestError) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *BotManagementInvalidRequestError) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *BotManagementInvalidRequestError) SetDetail(v string) {
	o.Detail = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementInvalidRequestError) MarshalJSON() ([]byte, error) {
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
func (o *BotManagementInvalidRequestError) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementInvalidRequestError := _BotManagementInvalidRequestError{}

	if err = json.Unmarshal(bytes, &varBotManagementInvalidRequestError); err == nil {
		*o = BotManagementInvalidRequestError(varBotManagementInvalidRequestError)
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

// NullableBotManagementInvalidRequestError is a helper abstraction for handling nullable botmanagementinvalidrequesterror types.
type NullableBotManagementInvalidRequestError struct {
	value *BotManagementInvalidRequestError
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementInvalidRequestError) Get() *BotManagementInvalidRequestError {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementInvalidRequestError) Set(val *BotManagementInvalidRequestError) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementInvalidRequestError) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementInvalidRequestError) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementInvalidRequestError returns a pointer to a new instance of NullableBotManagementInvalidRequestError.
func NewNullableBotManagementInvalidRequestError(val *BotManagementInvalidRequestError) *NullableBotManagementInvalidRequestError {
	return &NullableBotManagementInvalidRequestError{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementInvalidRequestError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementInvalidRequestError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
