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

// RelationshipsForTlsConfiguration struct for RelationshipsForTlsConfiguration
type RelationshipsForTlsConfiguration struct {
	RelationshipService               *RelationshipService
	RelationshipTlsDnsRecordsResponse *RelationshipTlsDnsRecordsResponse
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipsForTlsConfiguration) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipService
	err = json.Unmarshal(data, &o.RelationshipService)
	if err == nil {
		jsonRelationshipService, _ := json.Marshal(o.RelationshipService)
		if string(jsonRelationshipService) != "{}" { // empty struct
			return nil // data stored in o.RelationshipService, return on the first match
		}
		o.RelationshipService = nil
	} else {
		o.RelationshipService = nil
	}

	// try to unmarshal JSON data into RelationshipTlsDnsRecordsResponse
	err = json.Unmarshal(data, &o.RelationshipTlsDnsRecordsResponse)
	if err == nil {
		jsonRelationshipTlsDnsRecordsResponse, _ := json.Marshal(o.RelationshipTlsDnsRecordsResponse)
		if string(jsonRelationshipTlsDnsRecordsResponse) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTlsDnsRecordsResponse, return on the first match
		}
		o.RelationshipTlsDnsRecordsResponse = nil
	} else {
		o.RelationshipTlsDnsRecordsResponse = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForTlsConfiguration)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForTlsConfiguration) MarshalJSON() ([]byte, error) {
	if o.RelationshipService != nil {
		return json.Marshal(&o.RelationshipService)
	}

	if o.RelationshipTlsDnsRecordsResponse != nil {
		return json.Marshal(&o.RelationshipTlsDnsRecordsResponse)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForTlsConfiguration is a helper abstraction for handling nullable relationshipsfortlsconfiguration types.
type NullableRelationshipsForTlsConfiguration struct {
	value *RelationshipsForTlsConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForTlsConfiguration) Get() *RelationshipsForTlsConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForTlsConfiguration) Set(val *RelationshipsForTlsConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForTlsConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForTlsConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForTlsConfiguration returns a pointer to a new instance of NullableRelationshipsForTlsConfiguration.
func NewNullableRelationshipsForTlsConfiguration(val *RelationshipsForTlsConfiguration) *NullableRelationshipsForTlsConfiguration {
	return &NullableRelationshipsForTlsConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForTlsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipsForTlsConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
