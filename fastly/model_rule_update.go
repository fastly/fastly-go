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

// RuleUpdate All attributes for updating a rule. At least one of `action`, `conditions`, or `position` must be provided.
type RuleUpdate struct {
	Action               *Action                  `json:"action,omitempty"`
	Conditions           []RoutingConfigCondition `json:"conditions,omitempty"`
	Position             *Position                `json:"position,omitempty"`
	AdditionalProperties map[string]any
}

type _RuleUpdate RuleUpdate

// NewRuleUpdate instantiates a new RuleUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleUpdate() *RuleUpdate {
	this := RuleUpdate{}
	return &this
}

// NewRuleUpdateWithDefaults instantiates a new RuleUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleUpdateWithDefaults() *RuleUpdate {
	this := RuleUpdate{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RuleUpdate) GetAction() Action {
	if o == nil || o.Action == nil {
		var ret Action
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleUpdate) GetActionOk() (*Action, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RuleUpdate) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given Action and assigns it to the Action field.
func (o *RuleUpdate) SetAction(v Action) {
	o.Action = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *RuleUpdate) GetConditions() []RoutingConfigCondition {
	if o == nil || o.Conditions == nil {
		var ret []RoutingConfigCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleUpdate) GetConditionsOk() ([]RoutingConfigCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *RuleUpdate) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []RoutingConfigCondition and assigns it to the Conditions field.
func (o *RuleUpdate) SetConditions(v []RoutingConfigCondition) {
	o.Conditions = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *RuleUpdate) GetPosition() Position {
	if o == nil || o.Position == nil {
		var ret Position
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleUpdate) GetPositionOk() (*Position, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *RuleUpdate) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given Position and assigns it to the Position field.
func (o *RuleUpdate) SetPosition(v Position) {
	o.Position = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RuleUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}
	if o.Position != nil {
		toSerialize["position"] = o.Position
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RuleUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varRuleUpdate := _RuleUpdate{}

	if err = json.Unmarshal(bytes, &varRuleUpdate); err == nil {
		*o = RuleUpdate(varRuleUpdate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "conditions")
		delete(additionalProperties, "position")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRuleUpdate is a helper abstraction for handling nullable ruleupdate types.
type NullableRuleUpdate struct {
	value *RuleUpdate
	isSet bool
}

// Get returns the value.
func (v NullableRuleUpdate) Get() *RuleUpdate {
	return v.value
}

// Set modifies the value.
func (v *NullableRuleUpdate) Set(val *RuleUpdate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRuleUpdate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRuleUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRuleUpdate returns a pointer to a new instance of NullableRuleUpdate.
func NewNullableRuleUpdate(val *RuleUpdate) *NullableRuleUpdate {
	return &NullableRuleUpdate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRuleUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRuleUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
