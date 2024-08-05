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

// RelationshipsForTLSConfiguration struct for RelationshipsForTLSConfiguration
type RelationshipsForTLSConfiguration struct {
	RelationshipService *RelationshipService
	RelationshipTLSDNSRecordsResponse *RelationshipTLSDNSRecordsResponse
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipsForTLSConfiguration) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipService
	err = json.Unmarshal(data, &o.RelationshipService);
	if err == nil {
		jsonRelationshipService, _ := json.Marshal(o.RelationshipService)
		if string(jsonRelationshipService) != "{}" { // empty struct
			return nil // data stored in o.RelationshipService, return on the first match
		}
    o.RelationshipService = nil
	} else {
		o.RelationshipService = nil
	}

	// try to unmarshal JSON data into RelationshipTLSDNSRecordsResponse
	err = json.Unmarshal(data, &o.RelationshipTLSDNSRecordsResponse);
	if err == nil {
		jsonRelationshipTLSDNSRecordsResponse, _ := json.Marshal(o.RelationshipTLSDNSRecordsResponse)
		if string(jsonRelationshipTLSDNSRecordsResponse) != "{}" { // empty struct
			return nil // data stored in o.RelationshipTLSDNSRecordsResponse, return on the first match
		}
    o.RelationshipTLSDNSRecordsResponse = nil
	} else {
		o.RelationshipTLSDNSRecordsResponse = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForTLSConfiguration)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForTLSConfiguration) MarshalJSON() ([]byte, error) {
	if o.RelationshipService != nil {
		return json.Marshal(&o.RelationshipService)
	}

	if o.RelationshipTLSDNSRecordsResponse != nil {
		return json.Marshal(&o.RelationshipTLSDNSRecordsResponse)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForTLSConfiguration is a helper abstraction for handling nullable relationshipsfortlsconfiguration types. 
type NullableRelationshipsForTLSConfiguration struct {
	value *RelationshipsForTLSConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForTLSConfiguration) Get() *RelationshipsForTLSConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForTLSConfiguration) Set(val *RelationshipsForTLSConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForTLSConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForTLSConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForTLSConfiguration returns a pointer to a new instance of NullableRelationshipsForTLSConfiguration.
func NewNullableRelationshipsForTLSConfiguration(val *RelationshipsForTLSConfiguration) *NullableRelationshipsForTLSConfiguration {
	return &NullableRelationshipsForTLSConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForTLSConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipsForTLSConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
