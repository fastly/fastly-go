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

// NetworkProtectPrefixList Shared Network Protect prefix-list attributes.
type NetworkProtectPrefixList struct {
	// A human-readable name, unique within the customer.
	Name   string                         `json:"name"`
	Family NetworkProtectPrefixListFamily `json:"family"`
	// Complete canonical list contents. Bare addresses are canonicalized, CIDRs are masked, equivalent entries are deduplicated, and overlaps are retained.
	Addresses            []string `json:"addresses"`
	AdditionalProperties map[string]any
}

type _NetworkProtectPrefixList NetworkProtectPrefixList

// NewNetworkProtectPrefixList instantiates a new NetworkProtectPrefixList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectPrefixList(name string, family NetworkProtectPrefixListFamily, addresses []string) *NetworkProtectPrefixList {
	this := NetworkProtectPrefixList{}
	this.Name = name
	this.Family = family
	this.Addresses = addresses
	return &this
}

// NewNetworkProtectPrefixListWithDefaults instantiates a new NetworkProtectPrefixList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectPrefixListWithDefaults() *NetworkProtectPrefixList {
	this := NetworkProtectPrefixList{}
	return &this
}

// GetName returns the Name field value
func (o *NetworkProtectPrefixList) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixList) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkProtectPrefixList) SetName(v string) {
	o.Name = v
}

// GetFamily returns the Family field value
func (o *NetworkProtectPrefixList) GetFamily() NetworkProtectPrefixListFamily {
	if o == nil {
		var ret NetworkProtectPrefixListFamily
		return ret
	}

	return o.Family
}

// GetFamilyOk returns a tuple with the Family field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixList) GetFamilyOk() (*NetworkProtectPrefixListFamily, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Family, true
}

// SetFamily sets field value
func (o *NetworkProtectPrefixList) SetFamily(v NetworkProtectPrefixListFamily) {
	o.Family = v
}

// GetAddresses returns the Addresses field value
func (o *NetworkProtectPrefixList) GetAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixList) GetAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Addresses, true
}

// SetAddresses sets field value
func (o *NetworkProtectPrefixList) SetAddresses(v []string) {
	o.Addresses = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectPrefixList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["family"] = o.Family
	}
	if true {
		toSerialize["addresses"] = o.Addresses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectPrefixList) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectPrefixList := _NetworkProtectPrefixList{}

	if err = json.Unmarshal(bytes, &varNetworkProtectPrefixList); err == nil {
		*o = NetworkProtectPrefixList(varNetworkProtectPrefixList)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "family")
		delete(additionalProperties, "addresses")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectPrefixList is a helper abstraction for handling nullable networkprotectprefixlist types.
type NullableNetworkProtectPrefixList struct {
	value *NetworkProtectPrefixList
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectPrefixList) Get() *NetworkProtectPrefixList {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectPrefixList) Set(val *NetworkProtectPrefixList) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectPrefixList) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectPrefixList) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectPrefixList returns a pointer to a new instance of NullableNetworkProtectPrefixList.
func NewNullableNetworkProtectPrefixList(val *NetworkProtectPrefixList) *NullableNetworkProtectPrefixList {
	return &NullableNetworkProtectPrefixList{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectPrefixList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectPrefixList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
