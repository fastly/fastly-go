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

// NetworkProtectRuleSet Shared ordered rules and default action.
type NetworkProtectRuleSet struct {
	DefaultAction NetworkProtectRuleSetDefaultAction `json:"default_action"`
	// Rules in evaluation order with read-only timestamps.
	Rules                []NetworkProtectRuleResponse `json:"rules"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleSet NetworkProtectRuleSet

// NewNetworkProtectRuleSet instantiates a new NetworkProtectRuleSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleSet(defaultAction NetworkProtectRuleSetDefaultAction, rules []NetworkProtectRuleResponse) *NetworkProtectRuleSet {
	this := NetworkProtectRuleSet{}
	this.DefaultAction = defaultAction
	this.Rules = rules
	return &this
}

// NewNetworkProtectRuleSetWithDefaults instantiates a new NetworkProtectRuleSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleSetWithDefaults() *NetworkProtectRuleSet {
	this := NetworkProtectRuleSet{}
	return &this
}

// GetDefaultAction returns the DefaultAction field value
func (o *NetworkProtectRuleSet) GetDefaultAction() NetworkProtectRuleSetDefaultAction {
	if o == nil {
		var ret NetworkProtectRuleSetDefaultAction
		return ret
	}

	return o.DefaultAction
}

// GetDefaultActionOk returns a tuple with the DefaultAction field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSet) GetDefaultActionOk() (*NetworkProtectRuleSetDefaultAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultAction, true
}

// SetDefaultAction sets field value
func (o *NetworkProtectRuleSet) SetDefaultAction(v NetworkProtectRuleSetDefaultAction) {
	o.DefaultAction = v
}

// GetRules returns the Rules field value
func (o *NetworkProtectRuleSet) GetRules() []NetworkProtectRuleResponse {
	if o == nil {
		var ret []NetworkProtectRuleResponse
		return ret
	}

	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSet) GetRulesOk() ([]NetworkProtectRuleResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rules, true
}

// SetRules sets field value
func (o *NetworkProtectRuleSet) SetRules(v []NetworkProtectRuleResponse) {
	o.Rules = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleSet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["default_action"] = o.DefaultAction
	}
	if true {
		toSerialize["rules"] = o.Rules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRuleSet) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleSet := _NetworkProtectRuleSet{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleSet); err == nil {
		*o = NetworkProtectRuleSet(varNetworkProtectRuleSet)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "default_action")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleSet is a helper abstraction for handling nullable networkprotectruleset types.
type NullableNetworkProtectRuleSet struct {
	value *NetworkProtectRuleSet
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleSet) Get() *NetworkProtectRuleSet {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleSet) Set(val *NetworkProtectRuleSet) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleSet) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleSet) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleSet returns a pointer to a new instance of NullableNetworkProtectRuleSet.
func NewNullableNetworkProtectRuleSet(val *NetworkProtectRuleSet) *NullableNetworkProtectRuleSet {
	return &NullableNetworkProtectRuleSet{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
