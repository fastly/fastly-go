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

// PublicIPList struct for PublicIPList
type PublicIPList struct {
	// Fastly's IPv4 ranges.
	Addresses []string `json:"addresses,omitempty"`
	// Fastly's IPv6 ranges.
	Ipv6Addresses []string `json:"ipv6_addresses,omitempty"`
	AdditionalProperties map[string]any
}

type _PublicIPList PublicIPList

// NewPublicIPList instantiates a new PublicIPList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIPList() *PublicIPList {
	this := PublicIPList{}
	return &this
}

// NewPublicIPListWithDefaults instantiates a new PublicIPList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIPListWithDefaults() *PublicIPList {
	this := PublicIPList{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *PublicIPList) GetAddresses() []string {
	if o == nil || o.Addresses == nil {
		var ret []string
		return ret
	}
	return o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIPList) GetAddressesOk() ([]string, bool) {
	if o == nil || o.Addresses == nil {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *PublicIPList) HasAddresses() bool {
	if o != nil && o.Addresses != nil {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given []string and assigns it to the Addresses field.
func (o *PublicIPList) SetAddresses(v []string) {
	o.Addresses = v
}

// GetIpv6Addresses returns the Ipv6Addresses field value if set, zero value otherwise.
func (o *PublicIPList) GetIpv6Addresses() []string {
	if o == nil || o.Ipv6Addresses == nil {
		var ret []string
		return ret
	}
	return o.Ipv6Addresses
}

// GetIpv6AddressesOk returns a tuple with the Ipv6Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIPList) GetIpv6AddressesOk() ([]string, bool) {
	if o == nil || o.Ipv6Addresses == nil {
		return nil, false
	}
	return o.Ipv6Addresses, true
}

// HasIpv6Addresses returns a boolean if a field has been set.
func (o *PublicIPList) HasIpv6Addresses() bool {
	if o != nil && o.Ipv6Addresses != nil {
		return true
	}

	return false
}

// SetIpv6Addresses gets a reference to the given []string and assigns it to the Ipv6Addresses field.
func (o *PublicIPList) SetIpv6Addresses(v []string) {
	o.Ipv6Addresses = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PublicIPList) MarshalJSON() ([]byte, error) {
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
func (o *PublicIPList) UnmarshalJSON(bytes []byte) (err error) {
	varPublicIPList := _PublicIPList{}

	if err = json.Unmarshal(bytes, &varPublicIPList); err == nil {
		*o = PublicIPList(varPublicIPList)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "addresses")
		delete(additionalProperties, "ipv6_addresses")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePublicIPList is a helper abstraction for handling nullable publiciplist types. 
type NullablePublicIPList struct {
	value *PublicIPList
	isSet bool
}

// Get returns the value.
func (v NullablePublicIPList) Get() *PublicIPList {
	return v.value
}

// Set modifies the value.
func (v *NullablePublicIPList) Set(val *PublicIPList) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePublicIPList) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePublicIPList) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePublicIPList returns a pointer to a new instance of NullablePublicIPList.
func NewNullablePublicIPList(val *PublicIPList) *NullablePublicIPList {
	return &NullablePublicIPList{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePublicIPList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullablePublicIPList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
