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

// NetworkProtectRuleMatchIpProtocol Requires an IP protocol match field.
type NetworkProtectRuleMatchIpProtocol struct {
	IpProtocol           NetworkProtectRuleIpProtocol `json:"ip_protocol"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleMatchIpProtocol NetworkProtectRuleMatchIpProtocol

// NewNetworkProtectRuleMatchIpProtocol instantiates a new NetworkProtectRuleMatchIpProtocol object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleMatchIpProtocol(ipProtocol NetworkProtectRuleIpProtocol) *NetworkProtectRuleMatchIpProtocol {
	this := NetworkProtectRuleMatchIpProtocol{}
	this.IpProtocol = ipProtocol
	return &this
}

// NewNetworkProtectRuleMatchIpProtocolWithDefaults instantiates a new NetworkProtectRuleMatchIpProtocol object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleMatchIpProtocolWithDefaults() *NetworkProtectRuleMatchIpProtocol {
	this := NetworkProtectRuleMatchIpProtocol{}
	return &this
}

// GetIpProtocol returns the IpProtocol field value
func (o *NetworkProtectRuleMatchIpProtocol) GetIpProtocol() NetworkProtectRuleIpProtocol {
	if o == nil {
		var ret NetworkProtectRuleIpProtocol
		return ret
	}

	return o.IpProtocol
}

// GetIpProtocolOk returns a tuple with the IpProtocol field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleMatchIpProtocol) GetIpProtocolOk() (*NetworkProtectRuleIpProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpProtocol, true
}

// SetIpProtocol sets field value
func (o *NetworkProtectRuleMatchIpProtocol) SetIpProtocol(v NetworkProtectRuleIpProtocol) {
	o.IpProtocol = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleMatchIpProtocol) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["ip_protocol"] = o.IpProtocol
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRuleMatchIpProtocol) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleMatchIpProtocol := _NetworkProtectRuleMatchIpProtocol{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleMatchIpProtocol); err == nil {
		*o = NetworkProtectRuleMatchIpProtocol(varNetworkProtectRuleMatchIpProtocol)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ip_protocol")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleMatchIpProtocol is a helper abstraction for handling nullable networkprotectrulematchipprotocol types.
type NullableNetworkProtectRuleMatchIpProtocol struct {
	value *NetworkProtectRuleMatchIpProtocol
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleMatchIpProtocol) Get() *NetworkProtectRuleMatchIpProtocol {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleMatchIpProtocol) Set(val *NetworkProtectRuleMatchIpProtocol) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleMatchIpProtocol) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleMatchIpProtocol) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleMatchIpProtocol returns a pointer to a new instance of NullableNetworkProtectRuleMatchIpProtocol.
func NewNullableNetworkProtectRuleMatchIpProtocol(val *NetworkProtectRuleMatchIpProtocol) *NullableNetworkProtectRuleMatchIpProtocol {
	return &NullableNetworkProtectRuleMatchIpProtocol{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleMatchIpProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleMatchIpProtocol) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
