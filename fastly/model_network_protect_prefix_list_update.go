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

// NetworkProtectPrefixListUpdate Attributes accepted when updating a prefix list.
type NetworkProtectPrefixListUpdate struct {
	// A human-readable name, unique within the customer.
	Name *string `json:"name,omitempty"`
	// Complete request contents. At most 1000 entries may be submitted; bare addresses are canonicalized and CIDRs are masked by the backend.
	Addresses            []string `json:"addresses,omitempty"`
	AdditionalProperties map[string]any
}

type _NetworkProtectPrefixListUpdate NetworkProtectPrefixListUpdate

// NewNetworkProtectPrefixListUpdate instantiates a new NetworkProtectPrefixListUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectPrefixListUpdate() *NetworkProtectPrefixListUpdate {
	this := NetworkProtectPrefixListUpdate{}
	return &this
}

// NewNetworkProtectPrefixListUpdateWithDefaults instantiates a new NetworkProtectPrefixListUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectPrefixListUpdateWithDefaults() *NetworkProtectPrefixListUpdate {
	this := NetworkProtectPrefixListUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkProtectPrefixListUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NetworkProtectPrefixListUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkProtectPrefixListUpdate) SetName(v string) {
	o.Name = &v
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *NetworkProtectPrefixListUpdate) GetAddresses() []string {
	if o == nil || o.Addresses == nil {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListUpdate) GetAddressesOk() ([]string, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *NetworkProtectPrefixListUpdate) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *NetworkProtectPrefixListUpdate) SetAddresses(v []string) {
	o.Addresses = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectPrefixListUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectPrefixListUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectPrefixListUpdate := _NetworkProtectPrefixListUpdate{}

	if err = json.Unmarshal(bytes, &varNetworkProtectPrefixListUpdate); err == nil {
		*o = NetworkProtectPrefixListUpdate(varNetworkProtectPrefixListUpdate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "addresses")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectPrefixListUpdate is a helper abstraction for handling nullable networkprotectprefixlistupdate types.
type NullableNetworkProtectPrefixListUpdate struct {
	value *NetworkProtectPrefixListUpdate
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectPrefixListUpdate) Get() *NetworkProtectPrefixListUpdate {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectPrefixListUpdate) Set(val *NetworkProtectPrefixListUpdate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectPrefixListUpdate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectPrefixListUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectPrefixListUpdate returns a pointer to a new instance of NullableNetworkProtectPrefixListUpdate.
func NewNullableNetworkProtectPrefixListUpdate(val *NetworkProtectPrefixListUpdate) *NullableNetworkProtectPrefixListUpdate {
	return &NullableNetworkProtectPrefixListUpdate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectPrefixListUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectPrefixListUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
