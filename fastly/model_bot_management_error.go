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

// BotManagementError An error response, formatted as a problem details object.
type BotManagementError struct {
	// A short summary of the problem.
	Title string `json:"title"`
	// The HTTP status code of the response.
	Status int32 `json:"status"`
	// An explanation of the problem, specific to this occurrence.
	Detail               *string `json:"detail,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementError BotManagementError

// NewBotManagementError instantiates a new BotManagementError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementError(title string, status int32) *BotManagementError {
	this := BotManagementError{}
	this.Title = title
	this.Status = status
	return &this
}

// NewBotManagementErrorWithDefaults instantiates a new BotManagementError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementErrorWithDefaults() *BotManagementError {
	this := BotManagementError{}
	return &this
}

// GetTitle returns the Title field value
func (o *BotManagementError) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *BotManagementError) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *BotManagementError) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *BotManagementError) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BotManagementError) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BotManagementError) SetStatus(v int32) {
	o.Status = v
}

// GetDetail returns the Detail field value if set, zero value otherwise.
func (o *BotManagementError) GetDetail() string {
	if o == nil || o.Detail == nil {
		var ret string
		return ret
	}
	return *o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementError) GetDetailOk() (*string, bool) {
	if o == nil || o.Detail == nil {
		return nil, false
	}
	return o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *BotManagementError) HasDetail() bool {
	if o != nil && o.Detail != nil {
		return true
	}

	return false
}

// SetDetail gets a reference to the given string and assigns it to the Detail field.
func (o *BotManagementError) SetDetail(v string) {
	o.Detail = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementError) MarshalJSON() ([]byte, error) {
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
func (o *BotManagementError) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementError := _BotManagementError{}

	if err = json.Unmarshal(bytes, &varBotManagementError); err == nil {
		*o = BotManagementError(varBotManagementError)
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

// NullableBotManagementError is a helper abstraction for handling nullable botmanagementerror types.
type NullableBotManagementError struct {
	value *BotManagementError
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementError) Get() *BotManagementError {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementError) Set(val *BotManagementError) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementError) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementError) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementError returns a pointer to a new instance of NullableBotManagementError.
func NewNullableBotManagementError(val *BotManagementError) *NullableBotManagementError {
	return &NullableBotManagementError{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
