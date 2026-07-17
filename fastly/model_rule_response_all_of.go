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

// RuleResponseAllOf struct for RuleResponseAllOf
type RuleResponseAllOf struct {
	// Alphanumeric string identifying the rule. Stable across versions of the routing config.
	Id *string `json:"id,omitempty"`
	// Whether this is the default (catch-all) rule for the path.
	IsDefault *bool   `json:"is_default,omitempty"`
	Action    *Action `json:"action,omitempty"`
	// The conditions a request must satisfy for this rule to match. Empty for the default rule.
	Conditions           []RoutingConfigCondition `json:"conditions,omitempty"`
	AdditionalProperties map[string]any
}

type _RuleResponseAllOf RuleResponseAllOf

// NewRuleResponseAllOf instantiates a new RuleResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRuleResponseAllOf() *RuleResponseAllOf {
	this := RuleResponseAllOf{}
	return &this
}

// NewRuleResponseAllOfWithDefaults instantiates a new RuleResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRuleResponseAllOfWithDefaults() *RuleResponseAllOf {
	this := RuleResponseAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RuleResponseAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleResponseAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RuleResponseAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RuleResponseAllOf) SetId(v string) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *RuleResponseAllOf) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleResponseAllOf) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *RuleResponseAllOf) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *RuleResponseAllOf) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *RuleResponseAllOf) GetAction() Action {
	if o == nil || o.Action == nil {
		var ret Action
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleResponseAllOf) GetActionOk() (*Action, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *RuleResponseAllOf) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given Action and assigns it to the Action field.
func (o *RuleResponseAllOf) SetAction(v Action) {
	o.Action = &v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *RuleResponseAllOf) GetConditions() []RoutingConfigCondition {
	if o == nil || o.Conditions == nil {
		var ret []RoutingConfigCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RuleResponseAllOf) GetConditionsOk() ([]RoutingConfigCondition, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *RuleResponseAllOf) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []RoutingConfigCondition and assigns it to the Conditions field.
func (o *RuleResponseAllOf) SetConditions(v []RoutingConfigCondition) {
	o.Conditions = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RuleResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsDefault != nil {
		toSerialize["is_default"] = o.IsDefault
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RuleResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varRuleResponseAllOf := _RuleResponseAllOf{}

	if err = json.Unmarshal(bytes, &varRuleResponseAllOf); err == nil {
		*o = RuleResponseAllOf(varRuleResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "is_default")
		delete(additionalProperties, "action")
		delete(additionalProperties, "conditions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRuleResponseAllOf is a helper abstraction for handling nullable ruleresponseallof types.
type NullableRuleResponseAllOf struct {
	value *RuleResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableRuleResponseAllOf) Get() *RuleResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableRuleResponseAllOf) Set(val *RuleResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRuleResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRuleResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRuleResponseAllOf returns a pointer to a new instance of NullableRuleResponseAllOf.
func NewNullableRuleResponseAllOf(val *RuleResponseAllOf) *NullableRuleResponseAllOf {
	return &NullableRuleResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRuleResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRuleResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
