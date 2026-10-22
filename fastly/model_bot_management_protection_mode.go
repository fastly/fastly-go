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
	"fmt"
)

// BotManagementProtectionMode The protection mode of the workspace. `off` disables all enforcement. `log` detects and logs bots without blocking or challenging requests. `block` enforces the actions set on categories and bots.
type BotManagementProtectionMode string

// List of bot_management_protection_mode
const (
	BOTMANAGEMENTPROTECTIONMODE_FALSE BotManagementProtectionMode = "false"
	BOTMANAGEMENTPROTECTIONMODE_LOG   BotManagementProtectionMode = "log"
	BOTMANAGEMENTPROTECTIONMODE_BLOCK BotManagementProtectionMode = "block"
)

// AllowedBotManagementProtectionModeEnumValues All allowed values of BotManagementProtectionMode enum
var AllowedBotManagementProtectionModeEnumValues = []BotManagementProtectionMode{
	"false",
	"log",
	"block",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *BotManagementProtectionMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BotManagementProtectionMode(value)
	for _, existing := range AllowedBotManagementProtectionModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BotManagementProtectionMode", value)
}

// NewBotManagementProtectionModeFromValue returns a pointer to a valid BotManagementProtectionMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBotManagementProtectionModeFromValue(v string) (*BotManagementProtectionMode, error) {
	ev := BotManagementProtectionMode(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for BotManagementProtectionMode: valid values are %v", v, AllowedBotManagementProtectionModeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BotManagementProtectionMode) IsValid() bool {
	for _, existing := range AllowedBotManagementProtectionModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to bot_management_protection_mode value
func (v BotManagementProtectionMode) Ptr() *BotManagementProtectionMode {
	return &v
}

// NullableBotManagementProtectionMode is a helper abstraction for handling nullable botmanagementprotectionmode types.
type NullableBotManagementProtectionMode struct {
	value *BotManagementProtectionMode
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementProtectionMode) Get() *BotManagementProtectionMode {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementProtectionMode) Set(val *BotManagementProtectionMode) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementProtectionMode) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementProtectionMode) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementProtectionMode returns a pointer to a new instance of NullableBotManagementProtectionMode.
func NewNullableBotManagementProtectionMode(val *BotManagementProtectionMode) *NullableBotManagementProtectionMode {
	return &NullableBotManagementProtectionMode{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementProtectionMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementProtectionMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
