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

// RelationshipWafRuleRevisions struct for RelationshipWafRuleRevisions
type RelationshipWafRuleRevisions struct {
	WafRuleRevisions *RelationshipWafRuleRevisionWafRuleRevisions `json:"waf_rule_revisions,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipWafRuleRevisions RelationshipWafRuleRevisions

// NewRelationshipWafRuleRevisions instantiates a new RelationshipWafRuleRevisions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipWafRuleRevisions() *RelationshipWafRuleRevisions {
	this := RelationshipWafRuleRevisions{}
	return &this
}

// NewRelationshipWafRuleRevisionsWithDefaults instantiates a new RelationshipWafRuleRevisions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWafRuleRevisionsWithDefaults() *RelationshipWafRuleRevisions {
	this := RelationshipWafRuleRevisions{}
	return &this
}

// GetWafRuleRevisions returns the WafRuleRevisions field value if set, zero value otherwise.
func (o *RelationshipWafRuleRevisions) GetWafRuleRevisions() RelationshipWafRuleRevisionWafRuleRevisions {
	if o == nil || o.WafRuleRevisions == nil {
		var ret RelationshipWafRuleRevisionWafRuleRevisions
		return ret
	}
	return *o.WafRuleRevisions
}

// GetWafRuleRevisionsOk returns a tuple with the WafRuleRevisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipWafRuleRevisions) GetWafRuleRevisionsOk() (*RelationshipWafRuleRevisionWafRuleRevisions, bool) {
	if o == nil || o.WafRuleRevisions == nil {
		return nil, false
	}
	return o.WafRuleRevisions, true
}

// HasWafRuleRevisions returns a boolean if a field has been set.
func (o *RelationshipWafRuleRevisions) HasWafRuleRevisions() bool {
	if o != nil && o.WafRuleRevisions != nil {
		return true
	}

	return false
}

// SetWafRuleRevisions gets a reference to the given RelationshipWafRuleRevisionWafRuleRevisions and assigns it to the WafRuleRevisions field.
func (o *RelationshipWafRuleRevisions) SetWafRuleRevisions(v RelationshipWafRuleRevisionWafRuleRevisions) {
	o.WafRuleRevisions = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipWafRuleRevisions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.WafRuleRevisions != nil {
		toSerialize["waf_rule_revisions"] = o.WafRuleRevisions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipWafRuleRevisions) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipWafRuleRevisions := _RelationshipWafRuleRevisions{}

	if err = json.Unmarshal(bytes, &varRelationshipWafRuleRevisions); err == nil {
		*o = RelationshipWafRuleRevisions(varRelationshipWafRuleRevisions)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "waf_rule_revisions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipWafRuleRevisions is a helper abstraction for handling nullable relationshipwafrulerevisions types. 
type NullableRelationshipWafRuleRevisions struct {
	value *RelationshipWafRuleRevisions
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipWafRuleRevisions) Get() *RelationshipWafRuleRevisions {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipWafRuleRevisions) Set(val *RelationshipWafRuleRevisions) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipWafRuleRevisions) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipWafRuleRevisions) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipWafRuleRevisions returns a pointer to a new instance of NullableRelationshipWafRuleRevisions.
func NewNullableRelationshipWafRuleRevisions(val *RelationshipWafRuleRevisions) *NullableRelationshipWafRuleRevisions {
	return &NullableRelationshipWafRuleRevisions{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipWafRuleRevisions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipWafRuleRevisions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
