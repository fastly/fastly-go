// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
	"fmt"
)

// RelationshipsForWafRule struct for RelationshipsForWafRule
type RelationshipsForWafRule struct {
	RelationshipWafRuleRevisions *RelationshipWafRuleRevisions
	RelationshipWafTags *RelationshipWafTags
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipsForWafRule) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipWafRuleRevisions
	err = json.Unmarshal(data, &o.RelationshipWafRuleRevisions);
	if err == nil {
		jsonRelationshipWafRuleRevisions, _ := json.Marshal(o.RelationshipWafRuleRevisions)
		if string(jsonRelationshipWafRuleRevisions) != "{}" { // empty struct
			return nil // data stored in o.RelationshipWafRuleRevisions, return on the first match
		}
    o.RelationshipWafRuleRevisions = nil
	} else {
		o.RelationshipWafRuleRevisions = nil
	}

	// try to unmarshal JSON data into RelationshipWafTags
	err = json.Unmarshal(data, &o.RelationshipWafTags);
	if err == nil {
		jsonRelationshipWafTags, _ := json.Marshal(o.RelationshipWafTags)
		if string(jsonRelationshipWafTags) != "{}" { // empty struct
			return nil // data stored in o.RelationshipWafTags, return on the first match
		}
    o.RelationshipWafTags = nil
	} else {
		o.RelationshipWafTags = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForWafRule)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForWafRule) MarshalJSON() ([]byte, error) {
	if o.RelationshipWafRuleRevisions != nil {
		return json.Marshal(&o.RelationshipWafRuleRevisions)
	}

	if o.RelationshipWafTags != nil {
		return json.Marshal(&o.RelationshipWafTags)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForWafRule is a helper abstraction for handling nullable relationshipsforwafrule types. 
type NullableRelationshipsForWafRule struct {
	value *RelationshipsForWafRule
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForWafRule) Get() *RelationshipsForWafRule {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForWafRule) Set(val *RelationshipsForWafRule) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForWafRule) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForWafRule) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForWafRule returns a pointer to a new instance of NullableRelationshipsForWafRule.
func NewNullableRelationshipsForWafRule(val *RelationshipsForWafRule) *NullableRelationshipsForWafRule {
	return &NullableRelationshipsForWafRule{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForWafRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipsForWafRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
