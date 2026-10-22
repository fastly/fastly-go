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

// BotManagementNotFoundError struct for BotManagementNotFoundError
type BotManagementNotFoundError struct {
	Title  string `json:"title"`
	Status int32  `json:"status"`
	// An explanation of the problem, specific to this occurrence.
	Detail               *string `json:"detail,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementNotFoundError BotManagementNotFoundError

// NewBotManagementNotFoundError instantiates a new BotManagementNotFoundError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementNotFoundError(title string, status int32) *BotManagementNotFoundError {
	this := BotManagementNotFoundError{}
	this.Title = title
	this.Status = status
	return &this
}

// NewBotManagementNotFoundErrorWithDefaults instantiates a new BotManagementNotFoundError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementNotFoundErrorWithDefaults() *BotManagementNotFoundError {
	this := BotManagementNotFoundError{}
	return &this
}

// GetTitle returns the Title field value
func (o *BotManagementNotFoundError) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *BotManagementNotFoundError) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *BotManagementNotFoundError) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *BotManagementNotFoundError) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BotManagementNotFoundError) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BotManagementNotFoundError) SetStatus(v int32) {
	o.Status = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *BotManagementNotFoundError) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementNotFoundError) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *BotManagementNotFoundError) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *BotManagementNotFoundError) SetDetail(v string) {
	o.Detail = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementNotFoundError) MarshalJSON() ([]byte, error) {
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
func (o *BotManagementNotFoundError) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementNotFoundError := _BotManagementNotFoundError{}

	if err = json.Unmarshal(bytes, &varBotManagementNotFoundError); err == nil {
		*o = BotManagementNotFoundError(varBotManagementNotFoundError)
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

// NullableBotManagementNotFoundError is a helper abstraction for handling nullable botmanagementnotfounderror types.
type NullableBotManagementNotFoundError struct {
	value *BotManagementNotFoundError
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementNotFoundError) Get() *BotManagementNotFoundError {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementNotFoundError) Set(val *BotManagementNotFoundError) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementNotFoundError) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementNotFoundError) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementNotFoundError returns a pointer to a new instance of NullableBotManagementNotFoundError.
func NewNullableBotManagementNotFoundError(val *BotManagementNotFoundError) *NullableBotManagementNotFoundError {
	return &NullableBotManagementNotFoundError{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementNotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementNotFoundError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
