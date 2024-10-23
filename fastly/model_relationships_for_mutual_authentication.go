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

// RelationshipsForMutualAuthentication struct for RelationshipsForMutualAuthentication
type RelationshipsForMutualAuthentication struct {
	RelationshipTLSActivations *RelationshipTLSActivations
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipsForMutualAuthentication) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForMutualAuthentication)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForMutualAuthentication) MarshalJSON() ([]byte, error) {
	if o.RelationshipTLSActivations != nil {
		return json.Marshal(&o.RelationshipTLSActivations)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForMutualAuthentication is a helper abstraction for handling nullable relationshipsformutualauthentication types.
type NullableRelationshipsForMutualAuthentication struct {
	value *RelationshipsForMutualAuthentication
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForMutualAuthentication) Get() *RelationshipsForMutualAuthentication {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForMutualAuthentication) Set(val *RelationshipsForMutualAuthentication) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForMutualAuthentication) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForMutualAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForMutualAuthentication returns a pointer to a new instance of NullableRelationshipsForMutualAuthentication.
func NewNullableRelationshipsForMutualAuthentication(val *RelationshipsForMutualAuthentication) *NullableRelationshipsForMutualAuthentication {
	return &NullableRelationshipsForMutualAuthentication{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForMutualAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipsForMutualAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
