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

// RelationshipWafRules struct for RelationshipWafRules
type RelationshipWafRules struct {
	WafRules             *RelationshipWafRuleWafRule `json:"waf_rules,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipWafRules RelationshipWafRules

// NewRelationshipWafRules instantiates a new RelationshipWafRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipWafRules() *RelationshipWafRules {
	this := RelationshipWafRules{}
	return &this
}

// NewRelationshipWafRulesWithDefaults instantiates a new RelationshipWafRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWafRulesWithDefaults() *RelationshipWafRules {
	this := RelationshipWafRules{}
	return &this
}

// GetWafRules returns the WafRules field value if set, zero value otherwise.
func (o *RelationshipWafRules) GetWafRules() RelationshipWafRuleWafRule {
	if o == nil || o.WafRules == nil {
		var ret RelationshipWafRuleWafRule
		return ret
	}
	return *o.WafRules
}

// GetWafRulesOk returns a tuple with the WafRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipWafRules) GetWafRulesOk() (*RelationshipWafRuleWafRule, bool) {
	if o == nil || o.WafRules == nil {
		return nil, false
	}
	return o.WafRules, true
}

// HasWafRules returns a boolean if a field has been set.
func (o *RelationshipWafRules) HasWafRules() bool {
	if o != nil && o.WafRules != nil {
		return true
	}

	return false
}

// SetWafRules gets a reference to the given RelationshipWafRuleWafRule and assigns it to the WafRules field.
func (o *RelationshipWafRules) SetWafRules(v RelationshipWafRuleWafRule) {
	o.WafRules = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipWafRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.WafRules != nil {
		toSerialize["waf_rules"] = o.WafRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipWafRules) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipWafRules := _RelationshipWafRules{}

	if err = json.Unmarshal(bytes, &varRelationshipWafRules); err == nil {
		*o = RelationshipWafRules(varRelationshipWafRules)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "waf_rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipWafRules is a helper abstraction for handling nullable relationshipwafrules types.
type NullableRelationshipWafRules struct {
	value *RelationshipWafRules
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipWafRules) Get() *RelationshipWafRules {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipWafRules) Set(val *RelationshipWafRules) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipWafRules) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipWafRules) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipWafRules returns a pointer to a new instance of NullableRelationshipWafRules.
func NewNullableRelationshipWafRules(val *RelationshipWafRules) *NullableRelationshipWafRules {
	return &NullableRelationshipWafRules{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipWafRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipWafRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
