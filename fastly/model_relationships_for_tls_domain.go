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

// RelationshipsForTLSDomain struct for RelationshipsForTLSDomain
type RelationshipsForTLSDomain struct {
	RelationshipTLSActivations *RelationshipTLSActivations
	RelationshipTLSCertificates *RelationshipTLSCertificates
	RelationshipTLSSubscriptions *RelationshipTLSSubscriptions
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipsForTLSDomain) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipTLSActivations
	err = json.Unmarshal(data, &o.RelationshipTLSActivations);
	if err == nil {
		jsonRelationshipTLSActivations, _ := json.Marshal(o.RelationshipTLSActivations)
		if string(jsonRelationshipTLSActivations) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTLSActivations, return on the first match
		}
    o.RelationshipTLSActivations = nil
	} else {
		o.RelationshipTLSActivations = nil
	}

	// try to unmarshal JSON data into RelationshipTLSCertificates
	err = json.Unmarshal(data, &o.RelationshipTLSCertificates);
	if err == nil {
		jsonRelationshipTLSCertificates, _ := json.Marshal(o.RelationshipTLSCertificates)
		if string(jsonRelationshipTLSCertificates) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTLSCertificates, return on the first match
		}
    o.RelationshipTLSCertificates = nil
	} else {
		o.RelationshipTLSCertificates = nil
	}

	// try to unmarshal JSON data into RelationshipTLSSubscriptions
	err = json.Unmarshal(data, &o.RelationshipTLSSubscriptions);
	if err == nil {
		jsonRelationshipTLSSubscriptions, _ := json.Marshal(o.RelationshipTLSSubscriptions)
		if string(jsonRelationshipTLSSubscriptions) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTLSSubscriptions, return on the first match
		}
    o.RelationshipTLSSubscriptions = nil
	} else {
		o.RelationshipTLSSubscriptions = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForTLSDomain)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForTLSDomain) MarshalJSON() ([]byte, error) {
	if o.RelationshipTLSActivations != nil {
		return json.Marshal(&o.RelationshipTLSActivations)
	}

	if o.RelationshipTLSCertificates != nil {
		return json.Marshal(&o.RelationshipTLSCertificates)
	}

	if o.RelationshipTLSSubscriptions != nil {
		return json.Marshal(&o.RelationshipTLSSubscriptions)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForTLSDomain is a helper abstraction for handling nullable relationshipsfortlsdomain types. 
type NullableRelationshipsForTLSDomain struct {
	value *RelationshipsForTLSDomain
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForTLSDomain) Get() *RelationshipsForTLSDomain {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForTLSDomain) Set(val *RelationshipsForTLSDomain) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForTLSDomain) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForTLSDomain) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForTLSDomain returns a pointer to a new instance of NullableRelationshipsForTLSDomain.
func NewNullableRelationshipsForTLSDomain(val *RelationshipsForTLSDomain) *NullableRelationshipsForTLSDomain {
	return &NullableRelationshipsForTLSDomain{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForTLSDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipsForTLSDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
