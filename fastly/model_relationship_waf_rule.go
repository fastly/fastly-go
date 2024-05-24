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

// RelationshipWafRule struct for RelationshipWafRule
type RelationshipWafRule struct {
	WafRule *RelationshipWafRuleWafRule `json:"waf_rule,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipWafRule RelationshipWafRule

// NewRelationshipWafRule instantiates a new RelationshipWafRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipWafRule() *RelationshipWafRule {
	this := RelationshipWafRule{}
	return &this
}

// NewRelationshipWafRuleWithDefaults instantiates a new RelationshipWafRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWafRuleWithDefaults() *RelationshipWafRule {
	this := RelationshipWafRule{}
	return &this
}

// GetWafRule returns the WafRule field value if set, zero value otherwise.
func (o *RelationshipWafRule) GetWafRule() RelationshipWafRuleWafRule {
	if o == nil || o.WafRule == nil {
		var ret RelationshipWafRuleWafRule
		return ret
	}
	return *o.WafRule
}

// GetWafRuleOk returns a tuple with the WafRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipWafRule) GetWafRuleOk() (*RelationshipWafRuleWafRule, bool) {
	if o == nil || o.WafRule == nil {
		return nil, false
	}
	return o.WafRule, true
}

// HasWafRule returns a boolean if a field has been set.
func (o *RelationshipWafRule) HasWafRule() bool {
	if o != nil && o.WafRule != nil {
		return true
	}

	return false
}

// SetWafRule gets a reference to the given RelationshipWafRuleWafRule and assigns it to the WafRule field.
func (o *RelationshipWafRule) SetWafRule(v RelationshipWafRuleWafRule) {
	o.WafRule = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipWafRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.WafRule != nil {
		toSerialize["waf_rule"] = o.WafRule
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipWafRule) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipWafRule := _RelationshipWafRule{}

	if err = json.Unmarshal(bytes, &varRelationshipWafRule); err == nil {
		*o = RelationshipWafRule(varRelationshipWafRule)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "waf_rule")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipWafRule is a helper abstraction for handling nullable relationshipwafrule types. 
type NullableRelationshipWafRule struct {
	value *RelationshipWafRule
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipWafRule) Get() *RelationshipWafRule {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipWafRule) Set(val *RelationshipWafRule) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipWafRule) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipWafRule) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipWafRule returns a pointer to a new instance of NullableRelationshipWafRule.
func NewNullableRelationshipWafRule(val *RelationshipWafRule) *NullableRelationshipWafRule {
	return &NullableRelationshipWafRule{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipWafRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipWafRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
