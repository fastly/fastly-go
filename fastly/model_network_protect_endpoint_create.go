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

// NetworkProtectEndpointCreate Attributes accepted when creating an endpoint.
type NetworkProtectEndpointCreate struct {
	// A human-readable name, unique within the customer.
	Name string `json:"name"`
	// The GRE destination address, unique within the customer. It cannot be changed after endpoint creation.
	IPAddress            string `json:"ip_address"`
	AdditionalProperties map[string]any
}

type _NetworkProtectEndpointCreate NetworkProtectEndpointCreate

// NewNetworkProtectEndpointCreate instantiates a new NetworkProtectEndpointCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectEndpointCreate(name string, ipAddress string) *NetworkProtectEndpointCreate {
	this := NetworkProtectEndpointCreate{}
	this.Name = name
	this.IPAddress = ipAddress
	return &this
}

// NewNetworkProtectEndpointCreateWithDefaults instantiates a new NetworkProtectEndpointCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectEndpointCreateWithDefaults() *NetworkProtectEndpointCreate {
	this := NetworkProtectEndpointCreate{}
	return &this
}

// GetName returns the Name field value
func (o *NetworkProtectEndpointCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpointCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkProtectEndpointCreate) SetName(v string) {
	o.Name = v
}

// GetIPAddress returns the IPAddress field value
func (o *NetworkProtectEndpointCreate) GetIPAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IPAddress
}

// GetIPAddressOk returns a tuple with the IPAddress field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpointCreate) GetIPAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IPAddress, true
}

// SetIPAddress sets field value
func (o *NetworkProtectEndpointCreate) SetIPAddress(v string) {
	o.IPAddress = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectEndpointCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["ip_address"] = o.IPAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectEndpointCreate) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectEndpointCreate := _NetworkProtectEndpointCreate{}

	if err = json.Unmarshal(bytes, &varNetworkProtectEndpointCreate); err == nil {
		*o = NetworkProtectEndpointCreate(varNetworkProtectEndpointCreate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "ip_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectEndpointCreate is a helper abstraction for handling nullable networkprotectendpointcreate types.
type NullableNetworkProtectEndpointCreate struct {
	value *NetworkProtectEndpointCreate
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectEndpointCreate) Get() *NetworkProtectEndpointCreate {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectEndpointCreate) Set(val *NetworkProtectEndpointCreate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectEndpointCreate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectEndpointCreate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectEndpointCreate returns a pointer to a new instance of NullableNetworkProtectEndpointCreate.
func NewNullableNetworkProtectEndpointCreate(val *NetworkProtectEndpointCreate) *NullableNetworkProtectEndpointCreate {
	return &NullableNetworkProtectEndpointCreate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectEndpointCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectEndpointCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
