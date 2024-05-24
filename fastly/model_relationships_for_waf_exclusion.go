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
	"fmt"
)

// RelationshipsForWafExclusion struct for RelationshipsForWafExclusion
type RelationshipsForWafExclusion struct {
	RelationshipWafRuleRevisions *RelationshipWafRuleRevisions
	RelationshipWafRules *RelationshipWafRules
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipsForWafExclusion) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into RelationshipWafRules
	err = json.Unmarshal(data, &o.RelationshipWafRules);
	if err == nil {
		jsonRelationshipWafRules, _ := json.Marshal(o.RelationshipWafRules)
		if string(jsonRelationshipWafRules) != "{}" { // empty struct
			return nil // data stored in o.RelationshipWafRules, return on the first match
		}
    o.RelationshipWafRules = nil
	} else {
		o.RelationshipWafRules = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForWafExclusion)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForWafExclusion) MarshalJSON() ([]byte, error) {
	if o.RelationshipWafRuleRevisions != nil {
		return json.Marshal(&o.RelationshipWafRuleRevisions)
	}

	if o.RelationshipWafRules != nil {
		return json.Marshal(&o.RelationshipWafRules)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForWafExclusion is a helper abstraction for handling nullable relationshipsforwafexclusion types. 
type NullableRelationshipsForWafExclusion struct {
	value *RelationshipsForWafExclusion
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForWafExclusion) Get() *RelationshipsForWafExclusion {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForWafExclusion) Set(val *RelationshipsForWafExclusion) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForWafExclusion) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForWafExclusion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForWafExclusion returns a pointer to a new instance of NullableRelationshipsForWafExclusion.
func NewNullableRelationshipsForWafExclusion(val *RelationshipsForWafExclusion) *NullableRelationshipsForWafExclusion {
	return &NullableRelationshipsForWafExclusion{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForWafExclusion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipsForWafExclusion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
