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

// PublicIpList struct for PublicIpList
type PublicIpList struct {
	// Fastly's IPv4 ranges.
	Addresses []string `json:"addresses,omitempty"`
	// Fastly's IPv6 ranges.
	Ipv6Addresses        []string `json:"ipv6_addresses,omitempty"`
	AdditionalProperties map[string]any
}

type _PublicIpList PublicIpList

// NewPublicIpList instantiates a new PublicIpList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIpList() *PublicIpList {
	this := PublicIpList{}
	return &this
}

// NewPublicIpListWithDefaults instantiates a new PublicIpList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIpListWithDefaults() *PublicIpList {
	this := PublicIpList{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *PublicIpList) GetAddresses() []string {
	if o == nil || o.Addresses == nil {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIpList) GetAddressesOk() ([]string, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *PublicIpList) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *PublicIpList) SetAddresses(v []string) {
	o.Addresses = v
}

// GetIpv6Addresses returns the Ipv6Addresses field value if set, zero value otherwise.
func (o *PublicIpList) GetIpv6Addresses() []string {
	if o == nil || o.Ipv6Addresses == nil {
		var ret []string
		return ret
	}
	return o.Ipv6Addresses
}

// GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIpList) GetIpv6AddressesOk() ([]string, bool) {
	if o == nil || o.Ipv6Addresses == nil {
		return nil, false
	}
	return o.Ipv6Addresses, true
}

// HasIpv6Addresses returns a boolean if a field has been set.
func (o *PublicIpList) HasIpv6Addresses() bool {
	if o != nil && o.Ipv6Addresses != nil {
		return true
	}

	return false
}

// SetIpv6Addresses gets a reference to the given []string and assigns it to the Ipv6Addresses field.
func (o *PublicIpList) SetIpv6Addresses(v []string) {
	o.Ipv6Addresses = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PublicIpList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Addresses != nil {
		toSerialize["addresses"] = o.Addresses
	}
	if o.Ipv6Addresses != nil {
		toSerialize["ipv6_addresses"] = o.Ipv6Addresses
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PublicIpList) UnmarshalJSON(bytes []byte) (err error) {
	varPublicIpList := _PublicIpList{}

	if err = json.Unmarshal(bytes, &varPublicIpList); err == nil {
		*o = PublicIpList(varPublicIpList)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "addresses")
		delete(additionalProperties, "ipv6_addresses")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePublicIpList is a helper abstraction for handling nullable publiciplist types.
type NullablePublicIpList struct {
	value *PublicIpList
	isSet bool
}

// Get returns the value.
func (v NullablePublicIpList) Get() *PublicIpList {
	return v.value
}

// Set modifies the value.
func (v *NullablePublicIpList) Set(val *PublicIpList) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePublicIpList) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePublicIpList) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePublicIpList returns a pointer to a new instance of NullablePublicIpList.
func NewNullablePublicIpList(val *PublicIpList) *NullablePublicIpList {
	return &NullablePublicIpList{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePublicIpList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePublicIpList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
