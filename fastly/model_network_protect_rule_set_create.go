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

// NetworkProtectRuleSetCreate Attributes accepted when creating a rule set.
type NetworkProtectRuleSetCreate struct {
	DefaultAction NetworkProtectRuleSetDefaultAction `json:"default_action"`
	// Rules in evaluation order; the first match wins.
	Rules                []NetworkProtectRule `json:"rules,omitempty"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleSetCreate NetworkProtectRuleSetCreate

// NewNetworkProtectRuleSetCreate instantiates a new NetworkProtectRuleSetCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleSetCreate(defaultAction NetworkProtectRuleSetDefaultAction) *NetworkProtectRuleSetCreate {
	this := NetworkProtectRuleSetCreate{}
	this.DefaultAction = defaultAction
	return &this
}

// NewNetworkProtectRuleSetCreateWithDefaults instantiates a new NetworkProtectRuleSetCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleSetCreateWithDefaults() *NetworkProtectRuleSetCreate {
	this := NetworkProtectRuleSetCreate{}
	return &this
}

// GetDefaultAction returns the DefaultAction field value
func (o *NetworkProtectRuleSetCreate) GetDefaultAction() NetworkProtectRuleSetDefaultAction {
	if o == nil {
		var ret NetworkProtectRuleSetDefaultAction
		return ret
	}

	return o.DefaultAction
}

// GetDefaultActionOk returns a tuple with the DefaultAction field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSetCreate) GetDefaultActionOk() (*NetworkProtectRuleSetDefaultAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultAction, true
}

// SetDefaultAction sets field value
func (o *NetworkProtectRuleSetCreate) SetDefaultAction(v NetworkProtectRuleSetDefaultAction) {
	o.DefaultAction = v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *NetworkProtectRuleSetCreate) GetRules() []NetworkProtectRule {
	if o == nil || o.Rules == nil {
		var ret []NetworkProtectRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSetCreate) GetRulesOk() ([]NetworkProtectRule, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *NetworkProtectRuleSetCreate) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given []NetworkProtectRule and assigns it to the Rules field.
func (o *NetworkProtectRuleSetCreate) SetRules(v []NetworkProtectRule) {
	o.Rules = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleSetCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
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
func (o *NetworkProtectRuleSetCreate) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleSetCreate := _NetworkProtectRuleSetCreate{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleSetCreate); err == nil {
		*o = NetworkProtectRuleSetCreate(varNetworkProtectRuleSetCreate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "default_action")
		delete(additionalProperties, "rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleSetCreate is a helper abstraction for handling nullable networkprotectrulesetcreate types.
type NullableNetworkProtectRuleSetCreate struct {
	value *NetworkProtectRuleSetCreate
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleSetCreate) Get() *NetworkProtectRuleSetCreate {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleSetCreate) Set(val *NetworkProtectRuleSetCreate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleSetCreate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleSetCreate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleSetCreate returns a pointer to a new instance of NullableNetworkProtectRuleSetCreate.
func NewNullableNetworkProtectRuleSetCreate(val *NetworkProtectRuleSetCreate) *NullableNetworkProtectRuleSetCreate {
	return &NullableNetworkProtectRuleSetCreate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleSetCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleSetCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
