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
	"fmt"
)

// NetworkProtectRule One ordered firewall rule. All match fields are optional, but at least one must be provided. The action is required. Rules have no public ID; replacing the list preserves timestamps only when logical fields, comment, and position are unchanged.
type NetworkProtectRule struct {
	NetworkProtectRuleMatchDestinationIp    *NetworkProtectRuleMatchDestinationIp
	NetworkProtectRuleMatchDestinationPorts *NetworkProtectRuleMatchDestinationPorts
	NetworkProtectRuleMatchIpProtocol       *NetworkProtectRuleMatchIpProtocol
	NetworkProtectRuleMatchSourcePorts      *NetworkProtectRuleMatchSourcePorts
	NetworkProtectRuleMatchSourcePrefixList *NetworkProtectRuleMatchSourcePrefixList
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRule) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NetworkProtectRuleMatchDestinationIp
	err = json.Unmarshal(data, &o.NetworkProtectRuleMatchDestinationIp)
	if err == nil {
		jsonNetworkProtectRuleMatchDestinationIp, _ := json.Marshal(o.NetworkProtectRuleMatchDestinationIp)
		if string(jsonNetworkProtectRuleMatchDestinationIp) != "{}" { // empty struct
			return nil // data stored in o.NetworkProtectRuleMatchDestinationIp, return on the first match
		}
		o.NetworkProtectRuleMatchDestinationIp = nil
	} else {
		o.NetworkProtectRuleMatchDestinationIp = nil
	}

	// try to unmarshal JSON data into NetworkProtectRuleMatchDestinationPorts
	err = json.Unmarshal(data, &o.NetworkProtectRuleMatchDestinationPorts)
	if err == nil {
		jsonNetworkProtectRuleMatchDestinationPorts, _ := json.Marshal(o.NetworkProtectRuleMatchDestinationPorts)
		if string(jsonNetworkProtectRuleMatchDestinationPorts) != "{}" { // empty struct
			return nil // data stored in o.NetworkProtectRuleMatchDestinationPorts, return on the first match
		}
		o.NetworkProtectRuleMatchDestinationPorts = nil
	} else {
		o.NetworkProtectRuleMatchDestinationPorts = nil
	}

	// try to unmarshal JSON data into NetworkProtectRuleMatchIpProtocol
	err = json.Unmarshal(data, &o.NetworkProtectRuleMatchIpProtocol)
	if err == nil {
		jsonNetworkProtectRuleMatchIpProtocol, _ := json.Marshal(o.NetworkProtectRuleMatchIpProtocol)
		if string(jsonNetworkProtectRuleMatchIpProtocol) != "{}" { // empty struct
			return nil // data stored in o.NetworkProtectRuleMatchIpProtocol, return on the first match
		}
		o.NetworkProtectRuleMatchIpProtocol = nil
	} else {
		o.NetworkProtectRuleMatchIpProtocol = nil
	}

	// try to unmarshal JSON data into NetworkProtectRuleMatchSourcePorts
	err = json.Unmarshal(data, &o.NetworkProtectRuleMatchSourcePorts)
	if err == nil {
		jsonNetworkProtectRuleMatchSourcePorts, _ := json.Marshal(o.NetworkProtectRuleMatchSourcePorts)
		if string(jsonNetworkProtectRuleMatchSourcePorts) != "{}" { // empty struct
			return nil // data stored in o.NetworkProtectRuleMatchSourcePorts, return on the first match
		}
		o.NetworkProtectRuleMatchSourcePorts = nil
	} else {
		o.NetworkProtectRuleMatchSourcePorts = nil
	}

	// try to unmarshal JSON data into NetworkProtectRuleMatchSourcePrefixList
	err = json.Unmarshal(data, &o.NetworkProtectRuleMatchSourcePrefixList)
	if err == nil {
		jsonNetworkProtectRuleMatchSourcePrefixList, _ := json.Marshal(o.NetworkProtectRuleMatchSourcePrefixList)
		if string(jsonNetworkProtectRuleMatchSourcePrefixList) != "{}" { // empty struct
			return nil // data stored in o.NetworkProtectRuleMatchSourcePrefixList, return on the first match
		}
		o.NetworkProtectRuleMatchSourcePrefixList = nil
	} else {
		o.NetworkProtectRuleMatchSourcePrefixList = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NetworkProtectRule)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *NetworkProtectRule) MarshalJSON() ([]byte, error) {
	if o.NetworkProtectRuleMatchDestinationIp != nil {
		return json.Marshal(&o.NetworkProtectRuleMatchDestinationIp)
	}

	if o.NetworkProtectRuleMatchDestinationPorts != nil {
		return json.Marshal(&o.NetworkProtectRuleMatchDestinationPorts)
	}

	if o.NetworkProtectRuleMatchIpProtocol != nil {
		return json.Marshal(&o.NetworkProtectRuleMatchIpProtocol)
	}

	if o.NetworkProtectRuleMatchSourcePorts != nil {
		return json.Marshal(&o.NetworkProtectRuleMatchSourcePorts)
	}

	if o.NetworkProtectRuleMatchSourcePrefixList != nil {
		return json.Marshal(&o.NetworkProtectRuleMatchSourcePrefixList)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableNetworkProtectRule is a helper abstraction for handling nullable networkprotectrule types.
type NullableNetworkProtectRule struct {
	value *NetworkProtectRule
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRule) Get() *NetworkProtectRule {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRule) Set(val *NetworkProtectRule) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRule) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRule) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRule returns a pointer to a new instance of NullableNetworkProtectRule.
func NewNullableNetworkProtectRule(val *NetworkProtectRule) *NullableNetworkProtectRule {
	return &NullableNetworkProtectRule{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
