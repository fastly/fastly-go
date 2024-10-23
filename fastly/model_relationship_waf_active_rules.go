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

// RelationshipWafActiveRules struct for RelationshipWafActiveRules
type RelationshipWafActiveRules struct {
	WafActiveRules       *RelationshipWafActiveRulesWafActiveRules `json:"waf_active_rules,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipWafActiveRules RelationshipWafActiveRules

// NewRelationshipWafActiveRules instantiates a new RelationshipWafActiveRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipWafActiveRules() *RelationshipWafActiveRules {
	this := RelationshipWafActiveRules{}
	return &this
}

// NewRelationshipWafActiveRulesWithDefaults instantiates a new RelationshipWafActiveRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWafActiveRulesWithDefaults() *RelationshipWafActiveRules {
	this := RelationshipWafActiveRules{}
	return &this
}

// GetWafActiveRules returns the WafActiveRules field value if set, zero value otherwise.
func (o *RelationshipWafActiveRules) GetWafActiveRules() RelationshipWafActiveRulesWafActiveRules {
	if o == nil || o.WafActiveRules == nil {
		var ret RelationshipWafActiveRulesWafActiveRules
		return ret
	}
	return *o.WafActiveRules
}

// GetWafActiveRulesOk returns a tuple with the WafActiveRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipWafActiveRules) GetWafActiveRulesOk() (*RelationshipWafActiveRulesWafActiveRules, bool) {
	if o == nil || o.WafActiveRules == nil {
		return nil, false
	}
	return o.WafActiveRules, true
}

// HasWafActiveRules returns a boolean if a field has been set.
func (o *RelationshipWafActiveRules) HasWafActiveRules() bool {
	if o != nil && o.WafActiveRules != nil {
		return true
	}

	return false
}

// SetWafActiveRules gets a reference to the given RelationshipWafActiveRulesWafActiveRules and assigns it to the WafActiveRules field.
func (o *RelationshipWafActiveRules) SetWafActiveRules(v RelationshipWafActiveRulesWafActiveRules) {
	o.WafActiveRules = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipWafActiveRules) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.WafActiveRules != nil {
		toSerialize["waf_active_rules"] = o.WafActiveRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipWafActiveRules) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipWafActiveRules := _RelationshipWafActiveRules{}

	if err = json.Unmarshal(bytes, &varRelationshipWafActiveRules); err == nil {
		*o = RelationshipWafActiveRules(varRelationshipWafActiveRules)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "waf_active_rules")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipWafActiveRules is a helper abstraction for handling nullable relationshipwafactiverules types.
type NullableRelationshipWafActiveRules struct {
	value *RelationshipWafActiveRules
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipWafActiveRules) Get() *RelationshipWafActiveRules {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipWafActiveRules) Set(val *RelationshipWafActiveRules) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipWafActiveRules) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipWafActiveRules) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipWafActiveRules returns a pointer to a new instance of NullableRelationshipWafActiveRules.
func NewNullableRelationshipWafActiveRules(val *RelationshipWafActiveRules) *NullableRelationshipWafActiveRules {
	return &NullableRelationshipWafActiveRules{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipWafActiveRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipWafActiveRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
