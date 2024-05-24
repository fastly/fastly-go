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

// TypeServiceAuthorization Resource type
type TypeServiceAuthorization string

// List of resourceTypeservice_authorization
const (
	TYPESERVICEAUTHORIZATION_SERVICE_AUTHORIZATION TypeServiceAuthorization = "service_authorization"
)

// AllowedTypeServiceAuthorizationEnumValues All allowed values of TypeServiceAuthorization enum
var AllowedTypeServiceAuthorizationEnumValues = []TypeServiceAuthorization{
	"service_authorization",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeServiceAuthorization) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeServiceAuthorization(value)
	for _, existing := range AllowedTypeServiceAuthorizationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeServiceAuthorization", value)
}

// NewTypeServiceAuthorizationFromValue returns a pointer to a valid TypeServiceAuthorization
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeServiceAuthorizationFromValue(v string) (*TypeServiceAuthorization, error) {
	ev := TypeServiceAuthorization(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeServiceAuthorization: valid values are %v", v, AllowedTypeServiceAuthorizationEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeServiceAuthorization) IsValid() bool {
	for _, existing := range AllowedTypeServiceAuthorizationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypeservice_authorization value
func (v TypeServiceAuthorization) Ptr() *TypeServiceAuthorization {
	return &v
}

// NullableTypeServiceAuthorization is a helper abstraction for handling nullable typeserviceauthorization types. 
type NullableTypeServiceAuthorization struct {
	value *TypeServiceAuthorization
	isSet bool
}

// Get returns the value.
func (v NullableTypeServiceAuthorization) Get() *TypeServiceAuthorization {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeServiceAuthorization) Set(val *TypeServiceAuthorization) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeServiceAuthorization) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeServiceAuthorization) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeServiceAuthorization returns a pointer to a new instance of NullableTypeServiceAuthorization.
func NewNullableTypeServiceAuthorization(val *TypeServiceAuthorization) *NullableTypeServiceAuthorization {
	return &NullableTypeServiceAuthorization{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeServiceAuthorization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeServiceAuthorization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
