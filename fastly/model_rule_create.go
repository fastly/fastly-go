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

// RuleCreate All attributes for creating a rule. A rule with an empty `conditions` array is the default (catch-all) rule for its path.
type RuleCreate struct {
	Action Action `json:"action"`
	// The conditions a request must satisfy for this rule to match. An empty array indicates the default rule for the path.
	Conditions           []RoutingConfigCondition `json:"conditions"`
	Position             *Position                `json:"position,omitempty"`
	AdditionalProperties map[string]any
}

type _RuleCreate RuleCreate

// NewRuleCreate instantiates a new RuleCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleCreate(action Action, conditions []RoutingConfigCondition) *RuleCreate {
	this := RuleCreate{}
	this.Action = action
	this.Conditions = conditions
	return &this
}

// NewRuleCreateWithDefaults instantiates a new RuleCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleCreateWithDefaults() *RuleCreate {
	this := RuleCreate{}
	return &this
}

// GetAction returns the Action field value
func (o *RuleCreate) GetAction() Action {
	if o == nil {
		var ret Action
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *RuleCreate) GetActionOk() (*Action, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *RuleCreate) SetAction(v Action) {
	o.Action = v
}

// GetConditions returns the Conditions field value
func (o *RuleCreate) GetConditions() []RoutingConfigCondition {
	if o == nil {
		var ret []RoutingConfigCondition
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *RuleCreate) GetConditionsOk() ([]RoutingConfigCondition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *RuleCreate) SetConditions(v []RoutingConfigCondition) {
	o.Conditions = v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *RuleCreate) GetPosition() Position {
	if o == nil || o.Position == nil {
		var ret Position
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleCreate) GetPositionOk() (*Position, bool) {
	if o == nil || o.Position == nil {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *RuleCreate) HasPosition() bool {
	if o != nil && o.Position != nil {
		return true
	}

	return false
}

// SetPosition gets a reference to the given Position and assigns it to the Position field.
func (o *RuleCreate) SetPosition(v Position) {
	o.Position = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RuleCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
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
func (o *RuleCreate) UnmarshalJSON(bytes []byte) (err error) {
	varRuleCreate := _RuleCreate{}

	if err = json.Unmarshal(bytes, &varRuleCreate); err == nil {
		*o = RuleCreate(varRuleCreate)
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

// NullableRuleCreate is a helper abstraction for handling nullable rulecreate types.
type NullableRuleCreate struct {
	value *RuleCreate
	isSet bool
}

// Get returns the value.
func (v NullableRuleCreate) Get() *RuleCreate {
	return v.value
}

// Set modifies the value.
func (v *NullableRuleCreate) Set(val *RuleCreate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRuleCreate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRuleCreate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRuleCreate returns a pointer to a new instance of NullableRuleCreate.
func NewNullableRuleCreate(val *RuleCreate) *NullableRuleCreate {
	return &NullableRuleCreate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRuleCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRuleCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
