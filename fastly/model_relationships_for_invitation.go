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

// RelationshipsForInvitation struct for RelationshipsForInvitation
type RelationshipsForInvitation struct {
	RelationshipCustomer *RelationshipCustomer
	RelationshipServiceInvitations *RelationshipServiceInvitations
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipsForInvitation) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RelationshipCustomer
	err = json.Unmarshal(data, &o.RelationshipCustomer);
	if err == nil {
		jsonRelationshipCustomer, _ := json.Marshal(o.RelationshipCustomer)
		if string(jsonRelationshipCustomer) != "{}" { // empty struct
			return nil // data stored in o.RelationshipCustomer, return on the first match
		}
    o.RelationshipCustomer = nil
	} else {
		o.RelationshipCustomer = nil
	}

	// try to unmarshal JSON data into RelationshipServiceInvitations
	err = json.Unmarshal(data, &o.RelationshipServiceInvitations);
	if err == nil {
		jsonRelationshipServiceInvitations, _ := json.Marshal(o.RelationshipServiceInvitations)
		if string(jsonRelationshipServiceInvitations) != "{}" { // empty struct
			return nil // data stored in o.RelationshipServiceInvitations, return on the first match
		}
    o.RelationshipServiceInvitations = nil
	} else {
		o.RelationshipServiceInvitations = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(RelationshipsForInvitation)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *RelationshipsForInvitation) MarshalJSON() ([]byte, error) {
	if o.RelationshipCustomer != nil {
		return json.Marshal(&o.RelationshipCustomer)
	}

	if o.RelationshipServiceInvitations != nil {
		return json.Marshal(&o.RelationshipServiceInvitations)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableRelationshipsForInvitation is a helper abstraction for handling nullable relationshipsforinvitation types. 
type NullableRelationshipsForInvitation struct {
	value *RelationshipsForInvitation
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForInvitation) Get() *RelationshipsForInvitation {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForInvitation) Set(val *RelationshipsForInvitation) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForInvitation) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForInvitation returns a pointer to a new instance of NullableRelationshipsForInvitation.
func NewNullableRelationshipsForInvitation(val *RelationshipsForInvitation) *NullableRelationshipsForInvitation {
	return &NullableRelationshipsForInvitation{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipsForInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
