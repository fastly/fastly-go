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

// NetworkProtectRuleDestinationPorts Match traffic sent to this port or port range.
type NetworkProtectRuleDestinationPorts struct {
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleDestinationPorts NetworkProtectRuleDestinationPorts

// NewNetworkProtectRuleDestinationPorts instantiates a new NetworkProtectRuleDestinationPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleDestinationPorts() *NetworkProtectRuleDestinationPorts {
	this := NetworkProtectRuleDestinationPorts{}
	return &this
}

// NewNetworkProtectRuleDestinationPortsWithDefaults instantiates a new NetworkProtectRuleDestinationPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleDestinationPortsWithDefaults() *NetworkProtectRuleDestinationPorts {
	this := NetworkProtectRuleDestinationPorts{}
	return &this
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleDestinationPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRuleDestinationPorts) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleDestinationPorts := _NetworkProtectRuleDestinationPorts{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleDestinationPorts); err == nil {
		*o = NetworkProtectRuleDestinationPorts(varNetworkProtectRuleDestinationPorts)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleDestinationPorts is a helper abstraction for handling nullable networkprotectruledestinationports types.
type NullableNetworkProtectRuleDestinationPorts struct {
	value *NetworkProtectRuleDestinationPorts
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleDestinationPorts) Get() *NetworkProtectRuleDestinationPorts {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleDestinationPorts) Set(val *NetworkProtectRuleDestinationPorts) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleDestinationPorts) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleDestinationPorts) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleDestinationPorts returns a pointer to a new instance of NullableNetworkProtectRuleDestinationPorts.
func NewNullableNetworkProtectRuleDestinationPorts(val *NetworkProtectRuleDestinationPorts) *NullableNetworkProtectRuleDestinationPorts {
	return &NullableNetworkProtectRuleDestinationPorts{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleDestinationPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleDestinationPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
