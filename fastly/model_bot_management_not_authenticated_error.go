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

// BotManagementNotAuthenticatedError struct for BotManagementNotAuthenticatedError
type BotManagementNotAuthenticatedError struct {
	Title  string `json:"title"`
	Status int32  `json:"status"`
	// An explanation of the problem, specific to this occurrence.
	Detail               *string `json:"detail,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementNotAuthenticatedError BotManagementNotAuthenticatedError

// NewBotManagementNotAuthenticatedError instantiates a new BotManagementNotAuthenticatedError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementNotAuthenticatedError(title string, status int32) *BotManagementNotAuthenticatedError {
	this := BotManagementNotAuthenticatedError{}
	this.Title = title
	this.Status = status
	return &this
}

// NewBotManagementNotAuthenticatedErrorWithDefaults instantiates a new BotManagementNotAuthenticatedError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementNotAuthenticatedErrorWithDefaults() *BotManagementNotAuthenticatedError {
	this := BotManagementNotAuthenticatedError{}
	return &this
}

// GetTitle returns the Title field value
func (o *BotManagementNotAuthenticatedError) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *BotManagementNotAuthenticatedError) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *BotManagementNotAuthenticatedError) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *BotManagementNotAuthenticatedError) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BotManagementNotAuthenticatedError) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BotManagementNotAuthenticatedError) SetStatus(v int32) {
	o.Status = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *BotManagementNotAuthenticatedError) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementNotAuthenticatedError) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *BotManagementNotAuthenticatedError) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *BotManagementNotAuthenticatedError) SetDetail(v string) {
	o.Detail = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementNotAuthenticatedError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementNotAuthenticatedError) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementNotAuthenticatedError := _BotManagementNotAuthenticatedError{}

	if err = json.Unmarshal(bytes, &varBotManagementNotAuthenticatedError); err == nil {
		*o = BotManagementNotAuthenticatedError(varBotManagementNotAuthenticatedError)
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

// NullableBotManagementNotAuthenticatedError is a helper abstraction for handling nullable botmanagementnotauthenticatederror types.
type NullableBotManagementNotAuthenticatedError struct {
	value *BotManagementNotAuthenticatedError
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementNotAuthenticatedError) Get() *BotManagementNotAuthenticatedError {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementNotAuthenticatedError) Set(val *BotManagementNotAuthenticatedError) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementNotAuthenticatedError) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementNotAuthenticatedError) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementNotAuthenticatedError returns a pointer to a new instance of NullableBotManagementNotAuthenticatedError.
func NewNullableBotManagementNotAuthenticatedError(val *BotManagementNotAuthenticatedError) *NullableBotManagementNotAuthenticatedError {
	return &NullableBotManagementNotAuthenticatedError{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementNotAuthenticatedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementNotAuthenticatedError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
