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

// RelationshipsForTLSSubscription struct for RelationshipsForTLSSubscription
type RelationshipsForTLSSubscription struct {
	RelationshipCommonName *RelationshipCommonName
	RelationshipTLSCertificates *RelationshipTLSCertificates
	RelationshipTLSConfigurationForTLSSubscription *RelationshipTLSConfigurationForTLSSubscription
	RelationshipTLSDomains *RelationshipTLSDomains
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipsForTLSSubscription) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipCommonName
	err = json.Unmarshal(data, &o.RelationshipCommonName);
	if err == nil {
		jsonRelationshipCommonName, _ := json.Marshal(o.RelationshipCommonName)
		if string(jsonRelationshipCommonName) != "{}" { // empty struct
			return nil // data stored in o.RelationshipCommonName, return on the first match
		}
    o.RelationshipCommonName = nil
	} else {
		o.RelationshipCommonName = nil
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

	// try to unmarshal JSON data into RelationshipTLSConfigurationForTLSSubscription
	err = json.Unmarshal(data, &o.RelationshipTLSConfigurationForTLSSubscription);
	if err == nil {
		jsonRelationshipTLSConfigurationForTLSSubscription, _ := json.Marshal(o.RelationshipTLSConfigurationForTLSSubscription)
		if string(jsonRelationshipTLSConfigurationForTLSSubscription) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTLSConfigurationForTLSSubscription, return on the first match
		}
    o.RelationshipTLSConfigurationForTLSSubscription = nil
	} else {
		o.RelationshipTLSConfigurationForTLSSubscription = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForTLSSubscription)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForTLSSubscription) MarshalJSON() ([]byte, error) {
	if o.RelationshipCommonName != nil {
		return json.Marshal(&o.RelationshipCommonName)
	}

	if o.RelationshipTLSCertificates != nil {
		return json.Marshal(&o.RelationshipTLSCertificates)
	}

	if o.RelationshipTLSConfigurationForTLSSubscription != nil {
		return json.Marshal(&o.RelationshipTLSConfigurationForTLSSubscription)
	}

	if o.RelationshipTLSDomains != nil {
		return json.Marshal(&o.RelationshipTLSDomains)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForTLSSubscription is a helper abstraction for handling nullable relationshipsfortlssubscription types. 
type NullableRelationshipsForTLSSubscription struct {
	value *RelationshipsForTLSSubscription
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForTLSSubscription) Get() *RelationshipsForTLSSubscription {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForTLSSubscription) Set(val *RelationshipsForTLSSubscription) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForTLSSubscription) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForTLSSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForTLSSubscription returns a pointer to a new instance of NullableRelationshipsForTLSSubscription.
func NewNullableRelationshipsForTLSSubscription(val *RelationshipsForTLSSubscription) *NullableRelationshipsForTLSSubscription {
	return &NullableRelationshipsForTLSSubscription{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForTLSSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipsForTLSSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
