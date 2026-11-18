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

// NetworkProtectRuleMatchIPProtocol Requires an IP protocol match field.
type NetworkProtectRuleMatchIPProtocol struct {
	IPProtocol           NetworkProtectRuleIPProtocol `json:"ip_protocol"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleMatchIPProtocol NetworkProtectRuleMatchIPProtocol

// NewNetworkProtectRuleMatchIPProtocol instantiates a new NetworkProtectRuleMatchIPProtocol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleMatchIPProtocol(ipProtocol NetworkProtectRuleIPProtocol) *NetworkProtectRuleMatchIPProtocol {
	this := NetworkProtectRuleMatchIPProtocol{}
	this.IPProtocol = ipProtocol
	return &this
}

// NewNetworkProtectRuleMatchIPProtocolWithDefaults instantiates a new NetworkProtectRuleMatchIPProtocol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleMatchIPProtocolWithDefaults() *NetworkProtectRuleMatchIPProtocol {
	this := NetworkProtectRuleMatchIPProtocol{}
	return &this
}

// GetIPProtocol returns the IPProtocol field value
func (o *NetworkProtectRuleMatchIPProtocol) GetIPProtocol() NetworkProtectRuleIPProtocol {
	if o == nil {
		var ret NetworkProtectRuleIPProtocol
		return ret
	}

	return o.IPProtocol
}

// GetIPProtocolOk returns a tuple with the IPProtocol field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleMatchIPProtocol) GetIPProtocolOk() (*NetworkProtectRuleIPProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPProtocol, true
}

// SetIPProtocol sets field value
func (o *NetworkProtectRuleMatchIPProtocol) SetIPProtocol(v NetworkProtectRuleIPProtocol) {
	o.IPProtocol = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleMatchIPProtocol) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["ip_protocol"] = o.IPProtocol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRuleMatchIPProtocol) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleMatchIPProtocol := _NetworkProtectRuleMatchIPProtocol{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleMatchIPProtocol); err == nil {
		*o = NetworkProtectRuleMatchIPProtocol(varNetworkProtectRuleMatchIPProtocol)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ip_protocol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleMatchIPProtocol is a helper abstraction for handling nullable networkprotectrulematchipprotocol types.
type NullableNetworkProtectRuleMatchIPProtocol struct {
	value *NetworkProtectRuleMatchIPProtocol
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleMatchIPProtocol) Get() *NetworkProtectRuleMatchIPProtocol {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleMatchIPProtocol) Set(val *NetworkProtectRuleMatchIPProtocol) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleMatchIPProtocol) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleMatchIPProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleMatchIPProtocol returns a pointer to a new instance of NullableNetworkProtectRuleMatchIPProtocol.
func NewNullableNetworkProtectRuleMatchIPProtocol(val *NetworkProtectRuleMatchIPProtocol) *NullableNetworkProtectRuleMatchIPProtocol {
	return &NullableNetworkProtectRuleMatchIPProtocol{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleMatchIPProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleMatchIPProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
