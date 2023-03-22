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

// TypeServiceInvitation Resource type
type TypeServiceInvitation string

// List of resourceTypeservice_invitation
const (
	TYPESERVICEINVITATION_SERVICE_INVITATION TypeServiceInvitation = "service_invitation"
)

// AllowedTypeServiceInvitationEnumValues All allowed values of TypeServiceInvitation enum
var AllowedTypeServiceInvitationEnumValues = []TypeServiceInvitation{
	"service_invitation",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeServiceInvitation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeServiceInvitation(value)
	for _, existing := range AllowedTypeServiceInvitationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeServiceInvitation", value)
}

// NewTypeServiceInvitationFromValue returns a pointer to a valid TypeServiceInvitation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeServiceInvitationFromValue(v string) (*TypeServiceInvitation, error) {
	ev := TypeServiceInvitation(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeServiceInvitation: valid values are %v", v, AllowedTypeServiceInvitationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeServiceInvitation) IsValid() bool {
	for _, existing := range AllowedTypeServiceInvitationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypeservice_invitation value
func (v TypeServiceInvitation) Ptr() *TypeServiceInvitation {
	return &v
}

// NullableTypeServiceInvitation is a helper abstraction for handling nullable typeserviceinvitation types. 
type NullableTypeServiceInvitation struct {
	value *TypeServiceInvitation
	isSet bool
}

// Get returns the value.
func (v NullableTypeServiceInvitation) Get() *TypeServiceInvitation {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeServiceInvitation) Set(val *TypeServiceInvitation) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeServiceInvitation) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeServiceInvitation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeServiceInvitation returns a pointer to a new instance of NullableTypeServiceInvitation.
func NewNullableTypeServiceInvitation(val *TypeServiceInvitation) *NullableTypeServiceInvitation {
	return &NullableTypeServiceInvitation{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeServiceInvitation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeServiceInvitation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
