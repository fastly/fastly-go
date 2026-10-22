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

// NetworkProtectRuleMatchSourcePorts Requires a source-port match field.
type NetworkProtectRuleMatchSourcePorts struct {
	SourcePorts          NetworkProtectRuleSourcePorts `json:"source_ports"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleMatchSourcePorts NetworkProtectRuleMatchSourcePorts

// NewNetworkProtectRuleMatchSourcePorts instantiates a new NetworkProtectRuleMatchSourcePorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleMatchSourcePorts(sourcePorts NetworkProtectRuleSourcePorts) *NetworkProtectRuleMatchSourcePorts {
	this := NetworkProtectRuleMatchSourcePorts{}
	this.SourcePorts = sourcePorts
	return &this
}

// NewNetworkProtectRuleMatchSourcePortsWithDefaults instantiates a new NetworkProtectRuleMatchSourcePorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleMatchSourcePortsWithDefaults() *NetworkProtectRuleMatchSourcePorts {
	this := NetworkProtectRuleMatchSourcePorts{}
	return &this
}

// GetSourcePorts returns the SourcePorts field value
func (o *NetworkProtectRuleMatchSourcePorts) GetSourcePorts() NetworkProtectRuleSourcePorts {
	if o == nil {
		var ret NetworkProtectRuleSourcePorts
		return ret
	}

	return o.SourcePorts
}

// GetSourcePortsOk returns a tuple with the SourcePorts field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleMatchSourcePorts) GetSourcePortsOk() (*NetworkProtectRuleSourcePorts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourcePorts, true
}

// SetSourcePorts sets field value
func (o *NetworkProtectRuleMatchSourcePorts) SetSourcePorts(v NetworkProtectRuleSourcePorts) {
	o.SourcePorts = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleMatchSourcePorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["source_ports"] = o.SourcePorts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRuleMatchSourcePorts) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleMatchSourcePorts := _NetworkProtectRuleMatchSourcePorts{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleMatchSourcePorts); err == nil {
		*o = NetworkProtectRuleMatchSourcePorts(varNetworkProtectRuleMatchSourcePorts)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "source_ports")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleMatchSourcePorts is a helper abstraction for handling nullable networkprotectrulematchsourceports types.
type NullableNetworkProtectRuleMatchSourcePorts struct {
	value *NetworkProtectRuleMatchSourcePorts
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleMatchSourcePorts) Get() *NetworkProtectRuleMatchSourcePorts {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleMatchSourcePorts) Set(val *NetworkProtectRuleMatchSourcePorts) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleMatchSourcePorts) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleMatchSourcePorts) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleMatchSourcePorts returns a pointer to a new instance of NullableNetworkProtectRuleMatchSourcePorts.
func NewNullableNetworkProtectRuleMatchSourcePorts(val *NetworkProtectRuleMatchSourcePorts) *NullableNetworkProtectRuleMatchSourcePorts {
	return &NullableNetworkProtectRuleMatchSourcePorts{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleMatchSourcePorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleMatchSourcePorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
