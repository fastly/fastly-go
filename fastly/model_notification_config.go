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

// NotificationConfig Type-specific configuration
type NotificationConfig struct {
	// Email address for mailinglist notifications
	Address              *string `json:"address,omitempty"`
	AdditionalProperties map[string]any
}

type _NotificationConfig NotificationConfig

// NewNotificationConfig instantiates a new NotificationConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationConfig() *NotificationConfig {
	this := NotificationConfig{}
	return &this
}

// NewNotificationConfigWithDefaults instantiates a new NotificationConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationConfigWithDefaults() *NotificationConfig {
	this := NotificationConfig{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NotificationConfig) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationConfig) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NotificationConfig) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NotificationConfig) SetAddress(v string) {
	o.Address = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NotificationConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NotificationConfig) UnmarshalJSON(bytes []byte) (err error) {
	varNotificationConfig := _NotificationConfig{}

	if err = json.Unmarshal(bytes, &varNotificationConfig); err == nil {
		*o = NotificationConfig(varNotificationConfig)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNotificationConfig is a helper abstraction for handling nullable notificationconfig types.
type NullableNotificationConfig struct {
	value *NotificationConfig
	isSet bool
}

// Get returns the value.
func (v NullableNotificationConfig) Get() *NotificationConfig {
	return v.value
}

// Set modifies the value.
func (v *NullableNotificationConfig) Set(val *NotificationConfig) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNotificationConfig) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNotificationConfig) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNotificationConfig returns a pointer to a new instance of NullableNotificationConfig.
func NewNullableNotificationConfig(val *NotificationConfig) *NullableNotificationConfig {
	return &NullableNotificationConfig{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNotificationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNotificationConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
