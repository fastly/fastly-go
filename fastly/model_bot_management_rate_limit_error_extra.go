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

// BotManagementRateLimitErrorExtra struct for BotManagementRateLimitErrorExtra
type BotManagementRateLimitErrorExtra struct {
	Title                *string `json:"title,omitempty"`
	Status               *int32  `json:"status,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementRateLimitErrorExtra BotManagementRateLimitErrorExtra

// NewBotManagementRateLimitErrorExtra instantiates a new BotManagementRateLimitErrorExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementRateLimitErrorExtra() *BotManagementRateLimitErrorExtra {
	this := BotManagementRateLimitErrorExtra{}
	return &this
}

// NewBotManagementRateLimitErrorExtraWithDefaults instantiates a new BotManagementRateLimitErrorExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementRateLimitErrorExtraWithDefaults() *BotManagementRateLimitErrorExtra {
	this := BotManagementRateLimitErrorExtra{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BotManagementRateLimitErrorExtra) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementRateLimitErrorExtra) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BotManagementRateLimitErrorExtra) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BotManagementRateLimitErrorExtra) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BotManagementRateLimitErrorExtra) GetStatus() int32 {
	if o == nil || o.Status == nil {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementRateLimitErrorExtra) GetStatusOk() (*int32, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BotManagementRateLimitErrorExtra) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *BotManagementRateLimitErrorExtra) SetStatus(v int32) {
	o.Status = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementRateLimitErrorExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementRateLimitErrorExtra) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementRateLimitErrorExtra := _BotManagementRateLimitErrorExtra{}

	if err = json.Unmarshal(bytes, &varBotManagementRateLimitErrorExtra); err == nil {
		*o = BotManagementRateLimitErrorExtra(varBotManagementRateLimitErrorExtra)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementRateLimitErrorExtra is a helper abstraction for handling nullable botmanagementratelimiterrorextra types.
type NullableBotManagementRateLimitErrorExtra struct {
	value *BotManagementRateLimitErrorExtra
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementRateLimitErrorExtra) Get() *BotManagementRateLimitErrorExtra {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementRateLimitErrorExtra) Set(val *BotManagementRateLimitErrorExtra) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementRateLimitErrorExtra) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementRateLimitErrorExtra) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementRateLimitErrorExtra returns a pointer to a new instance of NullableBotManagementRateLimitErrorExtra.
func NewNullableBotManagementRateLimitErrorExtra(val *BotManagementRateLimitErrorExtra) *NullableBotManagementRateLimitErrorExtra {
	return &NullableBotManagementRateLimitErrorExtra{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementRateLimitErrorExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementRateLimitErrorExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
