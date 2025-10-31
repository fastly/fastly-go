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

// RelationshipsForTlsPrivateKey struct for RelationshipsForTlsPrivateKey
type RelationshipsForTlsPrivateKey struct {
	RelationshipTlsActivations *RelationshipTlsActivations
	RelationshipTlsDomains     *RelationshipTlsDomains
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipsForTlsPrivateKey) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipTlsActivations
	err = json.Unmarshal(data, &o.RelationshipTlsActivations)
	if err == nil {
		jsonRelationshipTlsActivations, _ := json.Marshal(o.RelationshipTlsActivations)
		if string(jsonRelationshipTlsActivations) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTlsActivations, return on the first match
		}
		o.RelationshipTlsActivations = nil
	} else {
		o.RelationshipTlsActivations = nil
	}

	// try to unmarshal JSON data into RelationshipTlsDomains
	err = json.Unmarshal(data, &o.RelationshipTlsDomains)
	if err == nil {
		jsonRelationshipTlsDomains, _ := json.Marshal(o.RelationshipTlsDomains)
		if string(jsonRelationshipTlsDomains) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTlsDomains, return on the first match
		}
		o.RelationshipTlsDomains = nil
	} else {
		o.RelationshipTlsDomains = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForTlsPrivateKey)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForTlsPrivateKey) MarshalJSON() ([]byte, error) {
	if o.RelationshipTlsActivations != nil {
		return json.Marshal(&o.RelationshipTlsActivations)
	}

	if o.RelationshipTlsDomains != nil {
		return json.Marshal(&o.RelationshipTlsDomains)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForTlsPrivateKey is a helper abstraction for handling nullable relationshipsfortlsprivatekey types.
type NullableRelationshipsForTlsPrivateKey struct {
	value *RelationshipsForTlsPrivateKey
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForTlsPrivateKey) Get() *RelationshipsForTlsPrivateKey {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForTlsPrivateKey) Set(val *RelationshipsForTlsPrivateKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForTlsPrivateKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForTlsPrivateKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForTlsPrivateKey returns a pointer to a new instance of NullableRelationshipsForTlsPrivateKey.
func NewNullableRelationshipsForTlsPrivateKey(val *RelationshipsForTlsPrivateKey) *NullableRelationshipsForTlsPrivateKey {
	return &NullableRelationshipsForTlsPrivateKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForTlsPrivateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipsForTlsPrivateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
