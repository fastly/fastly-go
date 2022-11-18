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

// TypeService Resource type
type TypeService string

// List of resourceTypeservice
const (
	TYPESERVICE_SERVICE TypeService = "service"
)

// AllowedTypeServiceEnumValues All allowed values of TypeService enum
var AllowedTypeServiceEnumValues = []TypeService{
	"service",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeService) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeService(value)
	for _, existing := range AllowedTypeServiceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeService", value)
}

// NewTypeServiceFromValue returns a pointer to a valid TypeService
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeServiceFromValue(v string) (*TypeService, error) {
	ev := TypeService(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeService: valid values are %v", v, AllowedTypeServiceEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeService) IsValid() bool {
	for _, existing := range AllowedTypeServiceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypeservice value
func (v TypeService) Ptr() *TypeService {
	return &v
}

// NullableTypeService is a helper abstraction for handling nullable typeservice types. 
type NullableTypeService struct {
	value *TypeService
	isSet bool
}

// Get returns the value.
func (v NullableTypeService) Get() *TypeService {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeService) Set(val *TypeService) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeService) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeService) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeService returns a pointer to a new instance of NullableTypeService.
func NewNullableTypeService(val *TypeService) *NullableTypeService {
	return &NullableTypeService{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
