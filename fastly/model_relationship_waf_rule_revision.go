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

// RelationshipWafRuleRevision struct for RelationshipWafRuleRevision
type RelationshipWafRuleRevision struct {
	WafRuleRevisions *RelationshipWafRuleRevisionWafRuleRevisions `json:"waf_rule_revisions,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipWafRuleRevision RelationshipWafRuleRevision

// NewRelationshipWafRuleRevision instantiates a new RelationshipWafRuleRevision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipWafRuleRevision() *RelationshipWafRuleRevision {
	this := RelationshipWafRuleRevision{}
	return &this
}

// NewRelationshipWafRuleRevisionWithDefaults instantiates a new RelationshipWafRuleRevision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWafRuleRevisionWithDefaults() *RelationshipWafRuleRevision {
	this := RelationshipWafRuleRevision{}
	return &this
}

// GetWafRuleRevisions returns the WafRuleRevisions field value if set, zero value otherwise.
func (o *RelationshipWafRuleRevision) GetWafRuleRevisions() RelationshipWafRuleRevisionWafRuleRevisions {
	if o == nil || o.WafRuleRevisions == nil {
		var ret RelationshipWafRuleRevisionWafRuleRevisions
		return ret
	}
	return *o.WafRuleRevisions
}

// GetWafRuleRevisionsOk returns a tuple with the WafRuleRevisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipWafRuleRevision) GetWafRuleRevisionsOk() (*RelationshipWafRuleRevisionWafRuleRevisions, bool) {
	if o == nil || o.WafRuleRevisions == nil {
		return nil, false
	}
	return o.WafRuleRevisions, true
}

// HasWafRuleRevisions returns a boolean if a field has been set.
func (o *RelationshipWafRuleRevision) HasWafRuleRevisions() bool {
	if o != nil && o.WafRuleRevisions != nil {
		return true
	}

	return false
}

// SetWafRuleRevisions gets a reference to the given RelationshipWafRuleRevisionWafRuleRevisions and assigns it to the WafRuleRevisions field.
func (o *RelationshipWafRuleRevision) SetWafRuleRevisions(v RelationshipWafRuleRevisionWafRuleRevisions) {
	o.WafRuleRevisions = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipWafRuleRevision) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipWafRuleRevision) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipWafRuleRevision := _RelationshipWafRuleRevision{}

	if err = json.Unmarshal(bytes, &varRelationshipWafRuleRevision); err == nil {
		*o = RelationshipWafRuleRevision(varRelationshipWafRuleRevision)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "waf_rule_revisions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipWafRuleRevision is a helper abstraction for handling nullable relationshipwafrulerevision types. 
type NullableRelationshipWafRuleRevision struct {
	value *RelationshipWafRuleRevision
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipWafRuleRevision) Get() *RelationshipWafRuleRevision {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipWafRuleRevision) Set(val *RelationshipWafRuleRevision) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipWafRuleRevision) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipWafRuleRevision) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipWafRuleRevision returns a pointer to a new instance of NullableRelationshipWafRuleRevision.
func NewNullableRelationshipWafRuleRevision(val *RelationshipWafRuleRevision) *NullableRelationshipWafRuleRevision {
	return &NullableRelationshipWafRuleRevision{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipWafRuleRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipWafRuleRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
