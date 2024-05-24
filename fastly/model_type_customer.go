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

// TypeCustomer Resource type
type TypeCustomer string

// List of resourceTypecustomer
const (
	TYPECUSTOMER_CUSTOMER TypeCustomer = "customer"
)

// AllowedTypeCustomerEnumValues All allowed values of TypeCustomer enum
var AllowedTypeCustomerEnumValues = []TypeCustomer{
	"customer",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeCustomer) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeCustomer(value)
	for _, existing := range AllowedTypeCustomerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeCustomer", value)
}

// NewTypeCustomerFromValue returns a pointer to a valid TypeCustomer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeCustomerFromValue(v string) (*TypeCustomer, error) {
	ev := TypeCustomer(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeCustomer: valid values are %v", v, AllowedTypeCustomerEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeCustomer) IsValid() bool {
	for _, existing := range AllowedTypeCustomerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypecustomer value
func (v TypeCustomer) Ptr() *TypeCustomer {
	return &v
}

// NullableTypeCustomer is a helper abstraction for handling nullable typecustomer types. 
type NullableTypeCustomer struct {
	value *TypeCustomer
	isSet bool
}

// Get returns the value.
func (v NullableTypeCustomer) Get() *TypeCustomer {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeCustomer) Set(val *TypeCustomer) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeCustomer) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeCustomer returns a pointer to a new instance of NullableTypeCustomer.
func NewNullableTypeCustomer(val *TypeCustomer) *NullableTypeCustomer {
	return &NullableTypeCustomer{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
