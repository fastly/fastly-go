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

// NetworkProtectRuleSetUpdate Attributes accepted when updating a rule set.
type NetworkProtectRuleSetUpdate struct {
	DefaultAction *NetworkProtectRuleSetDefaultAction `json:"default_action,omitempty"`
	// Rules in evaluation order; the first match wins.
	Rules                []NetworkProtectRule `json:"rules,omitempty"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleSetUpdate NetworkProtectRuleSetUpdate

// NewNetworkProtectRuleSetUpdate instantiates a new NetworkProtectRuleSetUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleSetUpdate() *NetworkProtectRuleSetUpdate {
	this := NetworkProtectRuleSetUpdate{}
	return &this
}

// NewNetworkProtectRuleSetUpdateWithDefaults instantiates a new NetworkProtectRuleSetUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleSetUpdateWithDefaults() *NetworkProtectRuleSetUpdate {
	this := NetworkProtectRuleSetUpdate{}
	return &this
}

// GetDefaultAction returns the DefaultAction field value if set, zero value otherwise.
func (o *NetworkProtectRuleSetUpdate) GetDefaultAction() NetworkProtectRuleSetDefaultAction {
	if o == nil || o.DefaultAction == nil {
		var ret NetworkProtectRuleSetDefaultAction
		return ret
	}
	return *o.DefaultAction
}

// GetDefaultActionOk returns a tuple with the DefaultAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSetUpdate) GetDefaultActionOk() (*NetworkProtectRuleSetDefaultAction, bool) {
	if o == nil || o.DefaultAction == nil {
		return nil, false
	}
	return o.DefaultAction, true
}

// HasDefaultAction returns a boolean if a field has been set.
func (o *NetworkProtectRuleSetUpdate) HasDefaultAction() bool {
	if o != nil && o.DefaultAction != nil {
		return true
	}

	return false
}

// SetDefaultAction gets a reference to the given NetworkProtectRuleSetDefaultAction and assigns it to the DefaultAction field.
func (o *NetworkProtectRuleSetUpdate) SetDefaultAction(v NetworkProtectRuleSetDefaultAction) {
	o.DefaultAction = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *NetworkProtectRuleSetUpdate) GetRules() []NetworkProtectRule {
	if o == nil || o.Rules == nil {
		var ret []NetworkProtectRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSetUpdate) GetRulesOk() ([]NetworkProtectRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *NetworkProtectRuleSetUpdate) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworkProtectRule and assigns it to the Rules field.
func (o *NetworkProtectRuleSetUpdate) SetRules(v []NetworkProtectRule) {
	o.Rules = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleSetUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.DefaultAction != nil {
		toSerialize["default_action"] = o.DefaultAction
	}
	if o.Rules != nil {
		toSerialize["rules"] = o.Rules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRuleSetUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleSetUpdate := _NetworkProtectRuleSetUpdate{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleSetUpdate); err == nil {
		*o = NetworkProtectRuleSetUpdate(varNetworkProtectRuleSetUpdate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "default_action")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleSetUpdate is a helper abstraction for handling nullable networkprotectrulesetupdate types.
type NullableNetworkProtectRuleSetUpdate struct {
	value *NetworkProtectRuleSetUpdate
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleSetUpdate) Get() *NetworkProtectRuleSetUpdate {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleSetUpdate) Set(val *NetworkProtectRuleSetUpdate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleSetUpdate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleSetUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleSetUpdate returns a pointer to a new instance of NullableNetworkProtectRuleSetUpdate.
func NewNullableNetworkProtectRuleSetUpdate(val *NetworkProtectRuleSetUpdate) *NullableNetworkProtectRuleSetUpdate {
	return &NullableNetworkProtectRuleSetUpdate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleSetUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleSetUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
