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

// RelationshipsForTlsSubscription struct for RelationshipsForTlsSubscription
type RelationshipsForTlsSubscription struct {
	RelationshipCommonName                         *RelationshipCommonName
	RelationshipTlsCertificates                    *RelationshipTlsCertificates
	RelationshipTlsConfigurationForTlsSubscription *RelationshipTlsConfigurationForTlsSubscription
	RelationshipTlsDomains                         *RelationshipTlsDomains
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipsForTlsSubscription) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipCommonName
	err = json.Unmarshal(data, &o.RelationshipCommonName)
	if err == nil {
		jsonRelationshipCommonName, _ := json.Marshal(o.RelationshipCommonName)
		if string(jsonRelationshipCommonName) != "{}" { // empty struct
			return nil // data stored in o.RelationshipCommonName, return on the first match
		}
		o.RelationshipCommonName = nil
	} else {
		o.RelationshipCommonName = nil
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

	// try to unmarshal JSON data into RelationshipTlsConfigurationForTlsSubscription
	err = json.Unmarshal(data, &o.RelationshipTlsConfigurationForTlsSubscription)
	if err == nil {
		jsonRelationshipTlsConfigurationForTlsSubscription, _ := json.Marshal(o.RelationshipTlsConfigurationForTlsSubscription)
		if string(jsonRelationshipTlsConfigurationForTlsSubscription) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTlsConfigurationForTlsSubscription, return on the first match
		}
		o.RelationshipTlsConfigurationForTlsSubscription = nil
	} else {
		o.RelationshipTlsConfigurationForTlsSubscription = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForTlsSubscription)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForTlsSubscription) MarshalJSON() ([]byte, error) {
	if o.RelationshipCommonName != nil {
		return json.Marshal(&o.RelationshipCommonName)
	}

	if o.RelationshipTlsCertificates != nil {
		return json.Marshal(&o.RelationshipTlsCertificates)
	}

	if o.RelationshipTlsConfigurationForTlsSubscription != nil {
		return json.Marshal(&o.RelationshipTlsConfigurationForTlsSubscription)
	}

	if o.RelationshipTlsDomains != nil {
		return json.Marshal(&o.RelationshipTlsDomains)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForTlsSubscription is a helper abstraction for handling nullable relationshipsfortlssubscription types.
type NullableRelationshipsForTlsSubscription struct {
	value *RelationshipsForTlsSubscription
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForTlsSubscription) Get() *RelationshipsForTlsSubscription {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForTlsSubscription) Set(val *RelationshipsForTlsSubscription) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForTlsSubscription) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForTlsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForTlsSubscription returns a pointer to a new instance of NullableRelationshipsForTlsSubscription.
func NewNullableRelationshipsForTlsSubscription(val *RelationshipsForTlsSubscription) *NullableRelationshipsForTlsSubscription {
	return &NullableRelationshipsForTlsSubscription{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForTlsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipsForTlsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
