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

// RelationshipsForTLSPrivateKey struct for RelationshipsForTLSPrivateKey
type RelationshipsForTLSPrivateKey struct {
	RelationshipTLSActivations *RelationshipTLSActivations
	RelationshipTLSDomains     *RelationshipTLSDomains
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipsForTLSPrivateKey) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipTLSActivations
	err = json.Unmarshal(data, &o.RelationshipTLSActivations)
	if err == nil {
		jsonRelationshipTLSActivations, _ := json.Marshal(o.RelationshipTLSActivations)
		if string(jsonRelationshipTLSActivations) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTLSActivations, return on the first match
		}
		o.RelationshipTLSActivations = nil
	} else {
		o.RelationshipTLSActivations = nil
	}

	// try to unmarshal JSON data into RelationshipTLSDomains
	err = json.Unmarshal(data, &o.RelationshipTLSDomains)
	if err == nil {
		jsonRelationshipTLSDomains, _ := json.Marshal(o.RelationshipTLSDomains)
		if string(jsonRelationshipTLSDomains) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTLSDomains, return on the first match
		}
		o.RelationshipTLSDomains = nil
	} else {
		o.RelationshipTLSDomains = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForTLSPrivateKey)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForTLSPrivateKey) MarshalJSON() ([]byte, error) {
	if o.RelationshipTLSActivations != nil {
		return json.Marshal(&o.RelationshipTLSActivations)
	}

	if o.RelationshipTLSDomains != nil {
		return json.Marshal(&o.RelationshipTLSDomains)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForTLSPrivateKey is a helper abstraction for handling nullable relationshipsfortlsprivatekey types.
type NullableRelationshipsForTLSPrivateKey struct {
	value *RelationshipsForTLSPrivateKey
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForTLSPrivateKey) Get() *RelationshipsForTLSPrivateKey {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForTLSPrivateKey) Set(val *RelationshipsForTLSPrivateKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForTLSPrivateKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForTLSPrivateKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForTLSPrivateKey returns a pointer to a new instance of NullableRelationshipsForTLSPrivateKey.
func NewNullableRelationshipsForTLSPrivateKey(val *RelationshipsForTLSPrivateKey) *NullableRelationshipsForTLSPrivateKey {
	return &NullableRelationshipsForTLSPrivateKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForTLSPrivateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipsForTLSPrivateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
