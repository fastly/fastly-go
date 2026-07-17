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

// RoutingConfigState The current state of the routing config's versions.
type RoutingConfigState string

// List of routing_config_state
const (
	ROUTINGCONFIGSTATE_state_draft_only        RoutingConfigState = "draft-only"
	ROUTINGCONFIGSTATE_state_active            RoutingConfigState = "active"
	ROUTINGCONFIGSTATE_state_active_with_draft RoutingConfigState = "active-with-draft"
)

// AllowedRoutingConfigStateEnumValues All allowed values of RoutingConfigState enum
var AllowedRoutingConfigStateEnumValues = []RoutingConfigState{
	"draft-only",
	"active",
	"active-with-draft",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *RoutingConfigState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoutingConfigState(value)
	for _, existing := range AllowedRoutingConfigStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoutingConfigState", value)
}

// NewRoutingConfigStateFromValue returns a pointer to a valid RoutingConfigState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoutingConfigStateFromValue(v string) (*RoutingConfigState, error) {
	ev := RoutingConfigState(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for RoutingConfigState: valid values are %v", v, AllowedRoutingConfigStateEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoutingConfigState) IsValid() bool {
	for _, existing := range AllowedRoutingConfigStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to routing_config_state value
func (v RoutingConfigState) Ptr() *RoutingConfigState {
	return &v
}

// NullableRoutingConfigState is a helper abstraction for handling nullable routingconfigstate types.
type NullableRoutingConfigState struct {
	value *RoutingConfigState
	isSet bool
}

// Get returns the value.
func (v NullableRoutingConfigState) Get() *RoutingConfigState {
	return v.value
}

// Set modifies the value.
func (v *NullableRoutingConfigState) Set(val *RoutingConfigState) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRoutingConfigState) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRoutingConfigState) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRoutingConfigState returns a pointer to a new instance of NullableRoutingConfigState.
func NewNullableRoutingConfigState(val *RoutingConfigState) *NullableRoutingConfigState {
	return &NullableRoutingConfigState{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRoutingConfigState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRoutingConfigState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
