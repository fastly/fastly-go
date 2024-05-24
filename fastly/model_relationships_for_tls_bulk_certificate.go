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

// RelationshipsForTLSBulkCertificate struct for RelationshipsForTLSBulkCertificate
type RelationshipsForTLSBulkCertificate struct {
	RelationshipTLSConfigurations *RelationshipTLSConfigurations
	RelationshipTLSDomains *RelationshipTLSDomains
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipsForTLSBulkCertificate) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipTLSConfigurations
	err = json.Unmarshal(data, &o.RelationshipTLSConfigurations);
	if err == nil {
		jsonRelationshipTLSConfigurations, _ := json.Marshal(o.RelationshipTLSConfigurations)
		if string(jsonRelationshipTLSConfigurations) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTLSConfigurations, return on the first match
		}
    o.RelationshipTLSConfigurations = nil
	} else {
		o.RelationshipTLSConfigurations = nil
	}

	// try to unmarshal JSON data into RelationshipTLSDomains
	err = json.Unmarshal(data, &o.RelationshipTLSDomains);
	if err == nil {
		jsonRelationshipTLSDomains, _ := json.Marshal(o.RelationshipTLSDomains)
		if string(jsonRelationshipTLSDomains) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTLSDomains, return on the first match
		}
    o.RelationshipTLSDomains = nil
	} else {
		o.RelationshipTLSDomains = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForTLSBulkCertificate)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForTLSBulkCertificate) MarshalJSON() ([]byte, error) {
	if o.RelationshipTLSConfigurations != nil {
		return json.Marshal(&o.RelationshipTLSConfigurations)
	}

	if o.RelationshipTLSDomains != nil {
		return json.Marshal(&o.RelationshipTLSDomains)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForTLSBulkCertificate is a helper abstraction for handling nullable relationshipsfortlsbulkcertificate types. 
type NullableRelationshipsForTLSBulkCertificate struct {
	value *RelationshipsForTLSBulkCertificate
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForTLSBulkCertificate) Get() *RelationshipsForTLSBulkCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForTLSBulkCertificate) Set(val *RelationshipsForTLSBulkCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForTLSBulkCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForTLSBulkCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForTLSBulkCertificate returns a pointer to a new instance of NullableRelationshipsForTLSBulkCertificate.
func NewNullableRelationshipsForTLSBulkCertificate(val *RelationshipsForTLSBulkCertificate) *NullableRelationshipsForTLSBulkCertificate {
	return &NullableRelationshipsForTLSBulkCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForTLSBulkCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipsForTLSBulkCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
