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

// WafExclusionResponseDataRelationships struct for WafExclusionResponseDataRelationships
type WafExclusionResponseDataRelationships struct {
	RelationshipWafRuleRevisions *RelationshipWafRuleRevisions
	RelationshipWafRules *RelationshipWafRules
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *WafExclusionResponseDataRelationships) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(WafExclusionResponseDataRelationships)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *WafExclusionResponseDataRelationships) MarshalJSON() ([]byte, error) {
	if o.RelationshipWafRuleRevisions != nil {
		return json.Marshal(&o.RelationshipWafRuleRevisions)
	}

	if o.RelationshipWafRules != nil {
		return json.Marshal(&o.RelationshipWafRules)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableWafExclusionResponseDataRelationships is a helper abstraction for handling nullable wafexclusionresponsedatarelationships types. 
type NullableWafExclusionResponseDataRelationships struct {
	value *WafExclusionResponseDataRelationships
	isSet bool
}

// Get returns the value.
func (v NullableWafExclusionResponseDataRelationships) Get() *WafExclusionResponseDataRelationships {
	return v.value
}

// Set modifies the value.
func (v *NullableWafExclusionResponseDataRelationships) Set(val *WafExclusionResponseDataRelationships) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafExclusionResponseDataRelationships) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafExclusionResponseDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafExclusionResponseDataRelationships returns a pointer to a new instance of NullableWafExclusionResponseDataRelationships.
func NewNullableWafExclusionResponseDataRelationships(val *WafExclusionResponseDataRelationships) *NullableWafExclusionResponseDataRelationships {
	return &NullableWafExclusionResponseDataRelationships{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafExclusionResponseDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafExclusionResponseDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
