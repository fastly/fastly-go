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

// NetworkProtectRuleMatchDestinationIP Requires a destination address match field.
type NetworkProtectRuleMatchDestinationIP struct {
	// Match traffic sent to this address; omit for any destination. An address outside a Network using this Rule Set never matches and still consumes one of the 64 rules. The API does not enforce containment in a network prefix; attached Rule Sets require a compatible IP family.
	DestinationIPAddress string `json:"destination_ip_address"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleMatchDestinationIP NetworkProtectRuleMatchDestinationIP

// NewNetworkProtectRuleMatchDestinationIP instantiates a new NetworkProtectRuleMatchDestinationIP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleMatchDestinationIP(destinationIPAddress string) *NetworkProtectRuleMatchDestinationIP {
	this := NetworkProtectRuleMatchDestinationIP{}
	this.DestinationIPAddress = destinationIPAddress
	return &this
}

// NewNetworkProtectRuleMatchDestinationIPWithDefaults instantiates a new NetworkProtectRuleMatchDestinationIP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleMatchDestinationIPWithDefaults() *NetworkProtectRuleMatchDestinationIP {
	this := NetworkProtectRuleMatchDestinationIP{}
	return &this
}

// GetDestinationIPAddress returns the DestinationIPAddress field value
func (o *NetworkProtectRuleMatchDestinationIP) GetDestinationIPAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DestinationIPAddress
}

// GetDestinationIPAddressOk returns a tuple with the DestinationIPAddress field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleMatchDestinationIP) GetDestinationIPAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestinationIPAddress, true
}

// SetDestinationIPAddress sets field value
func (o *NetworkProtectRuleMatchDestinationIP) SetDestinationIPAddress(v string) {
	o.DestinationIPAddress = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleMatchDestinationIP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["destination_ip_address"] = o.DestinationIPAddress
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRuleMatchDestinationIP) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleMatchDestinationIP := _NetworkProtectRuleMatchDestinationIP{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleMatchDestinationIP); err == nil {
		*o = NetworkProtectRuleMatchDestinationIP(varNetworkProtectRuleMatchDestinationIP)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "destination_ip_address")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleMatchDestinationIP is a helper abstraction for handling nullable networkprotectrulematchdestinationip types.
type NullableNetworkProtectRuleMatchDestinationIP struct {
	value *NetworkProtectRuleMatchDestinationIP
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleMatchDestinationIP) Get() *NetworkProtectRuleMatchDestinationIP {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleMatchDestinationIP) Set(val *NetworkProtectRuleMatchDestinationIP) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleMatchDestinationIP) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleMatchDestinationIP) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleMatchDestinationIP returns a pointer to a new instance of NullableNetworkProtectRuleMatchDestinationIP.
func NewNullableNetworkProtectRuleMatchDestinationIP(val *NetworkProtectRuleMatchDestinationIP) *NullableNetworkProtectRuleMatchDestinationIP {
	return &NullableNetworkProtectRuleMatchDestinationIP{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleMatchDestinationIP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleMatchDestinationIP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
