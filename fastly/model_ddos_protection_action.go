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

// DdosProtectionAction Action types for a rule.
type DdosProtectionAction string

// List of ddos_protection_action
const (
	DDOSPROTECTIONACTION_DEFAULT  DdosProtectionAction = "default"
	DDOSPROTECTIONACTION_BLOCK    DdosProtectionAction = "block"
	DDOSPROTECTIONACTION_LOG      DdosProtectionAction = "log"
	DDOSPROTECTIONACTION_DISABLED DdosProtectionAction = "disabled"
)

// AllowedDdosProtectionActionEnumValues All allowed values of DdosProtectionAction enum
var AllowedDdosProtectionActionEnumValues = []DdosProtectionAction{
	"default",
	"block",
	"log",
	"disabled",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *DdosProtectionAction) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DdosProtectionAction(value)
	for _, existing := range AllowedDdosProtectionActionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DdosProtectionAction", value)
}

// NewDdosProtectionActionFromValue returns a pointer to a valid DdosProtectionAction
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDdosProtectionActionFromValue(v string) (*DdosProtectionAction, error) {
	ev := DdosProtectionAction(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DdosProtectionAction: valid values are %v", v, AllowedDdosProtectionActionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DdosProtectionAction) IsValid() bool {
	for _, existing := range AllowedDdosProtectionActionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ddos_protection_action value
func (v DdosProtectionAction) Ptr() *DdosProtectionAction {
	return &v
}

// NullableDdosProtectionAction is a helper abstraction for handling nullable ddosprotectionaction types.
type NullableDdosProtectionAction struct {
	value *DdosProtectionAction
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionAction) Get() *DdosProtectionAction {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionAction) Set(val *DdosProtectionAction) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionAction) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionAction) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionAction returns a pointer to a new instance of NullableDdosProtectionAction.
func NewNullableDdosProtectionAction(val *DdosProtectionAction) *NullableDdosProtectionAction {
	return &NullableDdosProtectionAction{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
