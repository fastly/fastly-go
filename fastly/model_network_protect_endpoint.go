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

// NetworkProtectEndpoint Shared Network Protect endpoint attributes.
type NetworkProtectEndpoint struct {
	// A human-readable name, unique within the customer.
	Name string `json:"name"`
	// The GRE destination address, unique within the customer. It cannot be changed after endpoint creation.
	IpAddress            string `json:"ip_address"`
	AdditionalProperties map[string]any
}

type _NetworkProtectEndpoint NetworkProtectEndpoint

// NewNetworkProtectEndpoint instantiates a new NetworkProtectEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectEndpoint(name string, ipAddress string) *NetworkProtectEndpoint {
	this := NetworkProtectEndpoint{}
	this.Name = name
	this.IpAddress = ipAddress
	return &this
}

// NewNetworkProtectEndpointWithDefaults instantiates a new NetworkProtectEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectEndpointWithDefaults() *NetworkProtectEndpoint {
	this := NetworkProtectEndpoint{}
	return &this
}

// GetName returns the Name field value
func (o *NetworkProtectEndpoint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpoint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkProtectEndpoint) SetName(v string) {
	o.Name = v
}

// GetIpAddress returns the IpAddress field value
func (o *NetworkProtectEndpoint) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpoint) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *NetworkProtectEndpoint) SetIpAddress(v string) {
	o.IpAddress = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["ip_address"] = o.IpAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectEndpoint := _NetworkProtectEndpoint{}

	if err = json.Unmarshal(bytes, &varNetworkProtectEndpoint); err == nil {
		*o = NetworkProtectEndpoint(varNetworkProtectEndpoint)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "ip_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectEndpoint is a helper abstraction for handling nullable networkprotectendpoint types.
type NullableNetworkProtectEndpoint struct {
	value *NetworkProtectEndpoint
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectEndpoint) Get() *NetworkProtectEndpoint {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectEndpoint) Set(val *NetworkProtectEndpoint) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectEndpoint) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectEndpoint returns a pointer to a new instance of NullableNetworkProtectEndpoint.
func NewNullableNetworkProtectEndpoint(val *NetworkProtectEndpoint) *NullableNetworkProtectEndpoint {
	return &NullableNetworkProtectEndpoint{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
