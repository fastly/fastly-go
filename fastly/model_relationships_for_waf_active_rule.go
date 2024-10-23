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

// RelationshipsForWafActiveRule struct for RelationshipsForWafActiveRule
type RelationshipsForWafActiveRule struct {
	RelationshipWafFirewallVersion *RelationshipWafFirewallVersion
	RelationshipWafRuleRevision    *RelationshipWafRuleRevision
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipsForWafActiveRule) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipWafFirewallVersion
	err = json.Unmarshal(data, &o.RelationshipWafFirewallVersion)
	if err == nil {
		jsonRelationshipWafFirewallVersion, _ := json.Marshal(o.RelationshipWafFirewallVersion)
		if string(jsonRelationshipWafFirewallVersion) != "{}" { // empty struct
			return nil // data stored in o.RelationshipWafFirewallVersion, return on the first match
		}
		o.RelationshipWafFirewallVersion = nil
	} else {
		o.RelationshipWafFirewallVersion = nil
	}

	// try to unmarshal JSON data into RelationshipWafRuleRevision
	err = json.Unmarshal(data, &o.RelationshipWafRuleRevision)
	if err == nil {
		jsonRelationshipWafRuleRevision, _ := json.Marshal(o.RelationshipWafRuleRevision)
		if string(jsonRelationshipWafRuleRevision) != "{}" { // empty struct
			return nil // data stored in o.RelationshipWafRuleRevision, return on the first match
		}
		o.RelationshipWafRuleRevision = nil
	} else {
		o.RelationshipWafRuleRevision = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForWafActiveRule)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForWafActiveRule) MarshalJSON() ([]byte, error) {
	if o.RelationshipWafFirewallVersion != nil {
		return json.Marshal(&o.RelationshipWafFirewallVersion)
	}

	if o.RelationshipWafRuleRevision != nil {
		return json.Marshal(&o.RelationshipWafRuleRevision)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForWafActiveRule is a helper abstraction for handling nullable relationshipsforwafactiverule types.
type NullableRelationshipsForWafActiveRule struct {
	value *RelationshipsForWafActiveRule
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForWafActiveRule) Get() *RelationshipsForWafActiveRule {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForWafActiveRule) Set(val *RelationshipsForWafActiveRule) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForWafActiveRule) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForWafActiveRule) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForWafActiveRule returns a pointer to a new instance of NullableRelationshipsForWafActiveRule.
func NewNullableRelationshipsForWafActiveRule(val *RelationshipsForWafActiveRule) *NullableRelationshipsForWafActiveRule {
	return &NullableRelationshipsForWafActiveRule{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForWafActiveRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipsForWafActiveRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
