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

// RelationshipsForTLSCsr struct for RelationshipsForTLSCsr
type RelationshipsForTLSCsr struct {
	RelationshipTLSPrivateKey *RelationshipTLSPrivateKey
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipsForTLSCsr) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipTLSPrivateKey
	err = json.Unmarshal(data, &o.RelationshipTLSPrivateKey);
	if err == nil {
		jsonRelationshipTLSPrivateKey, _ := json.Marshal(o.RelationshipTLSPrivateKey)
		if string(jsonRelationshipTLSPrivateKey) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTLSPrivateKey, return on the first match
		}
    o.RelationshipTLSPrivateKey = nil
	} else {
		o.RelationshipTLSPrivateKey = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForTLSCsr)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForTLSCsr) MarshalJSON() ([]byte, error) {
	if o.RelationshipTLSPrivateKey != nil {
		return json.Marshal(&o.RelationshipTLSPrivateKey)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForTLSCsr is a helper abstraction for handling nullable relationshipsfortlscsr types. 
type NullableRelationshipsForTLSCsr struct {
	value *RelationshipsForTLSCsr
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForTLSCsr) Get() *RelationshipsForTLSCsr {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForTLSCsr) Set(val *RelationshipsForTLSCsr) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForTLSCsr) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForTLSCsr) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForTLSCsr returns a pointer to a new instance of NullableRelationshipsForTLSCsr.
func NewNullableRelationshipsForTLSCsr(val *RelationshipsForTLSCsr) *NullableRelationshipsForTLSCsr {
	return &NullableRelationshipsForTLSCsr{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForTLSCsr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipsForTLSCsr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
