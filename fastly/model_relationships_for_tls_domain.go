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

// RelationshipsForTlsDomain struct for RelationshipsForTlsDomain
type RelationshipsForTlsDomain struct {
	RelationshipTlsActivations   *RelationshipTlsActivations
	RelationshipTlsCertificates  *RelationshipTlsCertificates
	RelationshipTlsSubscriptions *RelationshipTlsSubscriptions
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipsForTlsDomain) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into RelationshipTlsCertificates
	err = json.Unmarshal(data, &o.RelationshipTlsCertificates)
	if err == nil {
		jsonRelationshipTlsCertificates, _ := json.Marshal(o.RelationshipTlsCertificates)
		if string(jsonRelationshipTlsCertificates) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTlsCertificates, return on the first match
		}
		o.RelationshipTlsCertificates = nil
	} else {
		o.RelationshipTlsCertificates = nil
	}

	// try to unmarshal JSON data into RelationshipTlsSubscriptions
	err = json.Unmarshal(data, &o.RelationshipTlsSubscriptions)
	if err == nil {
		jsonRelationshipTlsSubscriptions, _ := json.Marshal(o.RelationshipTlsSubscriptions)
		if string(jsonRelationshipTlsSubscriptions) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTlsSubscriptions, return on the first match
		}
		o.RelationshipTlsSubscriptions = nil
	} else {
		o.RelationshipTlsSubscriptions = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForTlsDomain)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForTlsDomain) MarshalJSON() ([]byte, error) {
	if o.RelationshipTlsActivations != nil {
		return json.Marshal(&o.RelationshipTlsActivations)
	}

	if o.RelationshipTlsCertificates != nil {
		return json.Marshal(&o.RelationshipTlsCertificates)
	}

	if o.RelationshipTlsSubscriptions != nil {
		return json.Marshal(&o.RelationshipTlsSubscriptions)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForTlsDomain is a helper abstraction for handling nullable relationshipsfortlsdomain types.
type NullableRelationshipsForTlsDomain struct {
	value *RelationshipsForTlsDomain
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForTlsDomain) Get() *RelationshipsForTlsDomain {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForTlsDomain) Set(val *RelationshipsForTlsDomain) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForTlsDomain) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForTlsDomain) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForTlsDomain returns a pointer to a new instance of NullableRelationshipsForTlsDomain.
func NewNullableRelationshipsForTlsDomain(val *RelationshipsForTlsDomain) *NullableRelationshipsForTlsDomain {
	return &NullableRelationshipsForTlsDomain{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForTlsDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipsForTlsDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
