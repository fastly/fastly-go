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

// NetworkProtectRuleMatchDestinationPorts Requires a destination-port match field.
type NetworkProtectRuleMatchDestinationPorts struct {
	DestinationPorts     NetworkProtectRuleDestinationPorts `json:"destination_ports"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleMatchDestinationPorts NetworkProtectRuleMatchDestinationPorts

// NewNetworkProtectRuleMatchDestinationPorts instantiates a new NetworkProtectRuleMatchDestinationPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleMatchDestinationPorts(destinationPorts NetworkProtectRuleDestinationPorts) *NetworkProtectRuleMatchDestinationPorts {
	this := NetworkProtectRuleMatchDestinationPorts{}
	this.DestinationPorts = destinationPorts
	return &this
}

// NewNetworkProtectRuleMatchDestinationPortsWithDefaults instantiates a new NetworkProtectRuleMatchDestinationPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleMatchDestinationPortsWithDefaults() *NetworkProtectRuleMatchDestinationPorts {
	this := NetworkProtectRuleMatchDestinationPorts{}
	return &this
}

// GetDestinationPorts returns the DestinationPorts field value
func (o *NetworkProtectRuleMatchDestinationPorts) GetDestinationPorts() NetworkProtectRuleDestinationPorts {
	if o == nil {
		var ret NetworkProtectRuleDestinationPorts
		return ret
	}

	return o.DestinationPorts
}

// GetDestinationPortsOk returns a tuple with the DestinationPorts field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleMatchDestinationPorts) GetDestinationPortsOk() (*NetworkProtectRuleDestinationPorts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationPorts, true
}

// SetDestinationPorts sets field value
func (o *NetworkProtectRuleMatchDestinationPorts) SetDestinationPorts(v NetworkProtectRuleDestinationPorts) {
	o.DestinationPorts = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleMatchDestinationPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["destination_ports"] = o.DestinationPorts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRuleMatchDestinationPorts) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleMatchDestinationPorts := _NetworkProtectRuleMatchDestinationPorts{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleMatchDestinationPorts); err == nil {
		*o = NetworkProtectRuleMatchDestinationPorts(varNetworkProtectRuleMatchDestinationPorts)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "destination_ports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleMatchDestinationPorts is a helper abstraction for handling nullable networkprotectrulematchdestinationports types.
type NullableNetworkProtectRuleMatchDestinationPorts struct {
	value *NetworkProtectRuleMatchDestinationPorts
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleMatchDestinationPorts) Get() *NetworkProtectRuleMatchDestinationPorts {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleMatchDestinationPorts) Set(val *NetworkProtectRuleMatchDestinationPorts) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleMatchDestinationPorts) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleMatchDestinationPorts) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleMatchDestinationPorts returns a pointer to a new instance of NullableNetworkProtectRuleMatchDestinationPorts.
func NewNullableNetworkProtectRuleMatchDestinationPorts(val *NetworkProtectRuleMatchDestinationPorts) *NullableNetworkProtectRuleMatchDestinationPorts {
	return &NullableNetworkProtectRuleMatchDestinationPorts{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleMatchDestinationPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleMatchDestinationPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
