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

// RelationshipsForWafFirewallVersion struct for RelationshipsForWafFirewallVersion
type RelationshipsForWafFirewallVersion struct {
	RelationshipWafActiveRules      *RelationshipWafActiveRules
	RelationshipWafFirewallVersions *RelationshipWafFirewallVersions
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipsForWafFirewallVersion) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipWafActiveRules
	err = json.Unmarshal(data, &o.RelationshipWafActiveRules)
	if err == nil {
		jsonRelationshipWafActiveRules, _ := json.Marshal(o.RelationshipWafActiveRules)
		if string(jsonRelationshipWafActiveRules) != "{}" { // empty struct
			return nil // data stored in o.RelationshipWafActiveRules, return on the first match
		}
		o.RelationshipWafActiveRules = nil
	} else {
		o.RelationshipWafActiveRules = nil
	}

	// try to unmarshal JSON data into RelationshipWafFirewallVersions
	err = json.Unmarshal(data, &o.RelationshipWafFirewallVersions)
	if err == nil {
		jsonRelationshipWafFirewallVersions, _ := json.Marshal(o.RelationshipWafFirewallVersions)
		if string(jsonRelationshipWafFirewallVersions) != "{}" { // empty struct
			return nil // data stored in o.RelationshipWafFirewallVersions, return on the first match
		}
		o.RelationshipWafFirewallVersions = nil
	} else {
		o.RelationshipWafFirewallVersions = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForWafFirewallVersion)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForWafFirewallVersion) MarshalJSON() ([]byte, error) {
	if o.RelationshipWafActiveRules != nil {
		return json.Marshal(&o.RelationshipWafActiveRules)
	}

	if o.RelationshipWafFirewallVersions != nil {
		return json.Marshal(&o.RelationshipWafFirewallVersions)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForWafFirewallVersion is a helper abstraction for handling nullable relationshipsforwaffirewallversion types.
type NullableRelationshipsForWafFirewallVersion struct {
	value *RelationshipsForWafFirewallVersion
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForWafFirewallVersion) Get() *RelationshipsForWafFirewallVersion {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForWafFirewallVersion) Set(val *RelationshipsForWafFirewallVersion) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForWafFirewallVersion) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForWafFirewallVersion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForWafFirewallVersion returns a pointer to a new instance of NullableRelationshipsForWafFirewallVersion.
func NewNullableRelationshipsForWafFirewallVersion(val *RelationshipsForWafFirewallVersion) *NullableRelationshipsForWafFirewallVersion {
	return &NullableRelationshipsForWafFirewallVersion{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForWafFirewallVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipsForWafFirewallVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
