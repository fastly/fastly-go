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

// RuleChange A modification to an existing rule's action between versions.
type RuleChange struct {
	// Alphanumeric string identifying the rule. Stable across versions of the routing config.
	RuleId               *string `json:"rule_id,omitempty"`
	OldAction            *Action `json:"old_action,omitempty"`
	NewAction            *Action `json:"new_action,omitempty"`
	AdditionalProperties map[string]any
}

type _RuleChange RuleChange

// NewRuleChange instantiates a new RuleChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleChange() *RuleChange {
	this := RuleChange{}
	return &this
}

// NewRuleChangeWithDefaults instantiates a new RuleChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleChangeWithDefaults() *RuleChange {
	this := RuleChange{}
	return &this
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *RuleChange) GetRuleId() string {
	if o == nil || o.RuleId == nil {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleChange) GetRuleIdOk() (*string, bool) {
	if o == nil || o.RuleId == nil {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *RuleChange) HasRuleId() bool {
	if o != nil && o.RuleId != nil {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *RuleChange) SetRuleId(v string) {
	o.RuleId = &v
}

// GetOldAction returns the OldAction field value if set, zero value otherwise.
func (o *RuleChange) GetOldAction() Action {
	if o == nil || o.OldAction == nil {
		var ret Action
		return ret
	}
	return *o.OldAction
}

// GetOldActionOk returns a tuple with the OldAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleChange) GetOldActionOk() (*Action, bool) {
	if o == nil || o.OldAction == nil {
		return nil, false
	}
	return o.OldAction, true
}

// HasOldAction returns a boolean if a field has been set.
func (o *RuleChange) HasOldAction() bool {
	if o != nil && o.OldAction != nil {
		return true
	}

	return false
}

// SetOldAction gets a reference to the given Action and assigns it to the OldAction field.
func (o *RuleChange) SetOldAction(v Action) {
	o.OldAction = &v
}

// GetNewAction returns the NewAction field value if set, zero value otherwise.
func (o *RuleChange) GetNewAction() Action {
	if o == nil || o.NewAction == nil {
		var ret Action
		return ret
	}
	return *o.NewAction
}

// GetNewActionOk returns a tuple with the NewAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleChange) GetNewActionOk() (*Action, bool) {
	if o == nil || o.NewAction == nil {
		return nil, false
	}
	return o.NewAction, true
}

// HasNewAction returns a boolean if a field has been set.
func (o *RuleChange) HasNewAction() bool {
	if o != nil && o.NewAction != nil {
		return true
	}

	return false
}

// SetNewAction gets a reference to the given Action and assigns it to the NewAction field.
func (o *RuleChange) SetNewAction(v Action) {
	o.NewAction = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RuleChange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.RuleId != nil {
		toSerialize["rule_id"] = o.RuleId
	}
	if o.OldAction != nil {
		toSerialize["old_action"] = o.OldAction
	}
	if o.NewAction != nil {
		toSerialize["new_action"] = o.NewAction
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RuleChange) UnmarshalJSON(bytes []byte) (err error) {
	varRuleChange := _RuleChange{}

	if err = json.Unmarshal(bytes, &varRuleChange); err == nil {
		*o = RuleChange(varRuleChange)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "rule_id")
		delete(additionalProperties, "old_action")
		delete(additionalProperties, "new_action")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRuleChange is a helper abstraction for handling nullable rulechange types.
type NullableRuleChange struct {
	value *RuleChange
	isSet bool
}

// Get returns the value.
func (v NullableRuleChange) Get() *RuleChange {
	return v.value
}

// Set modifies the value.
func (v *NullableRuleChange) Set(val *RuleChange) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRuleChange) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRuleChange) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRuleChange returns a pointer to a new instance of NullableRuleChange.
func NewNullableRuleChange(val *RuleChange) *NullableRuleChange {
	return &NullableRuleChange{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRuleChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRuleChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
