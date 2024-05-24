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

// TypeBillingAddress Resource type
type TypeBillingAddress string

// List of resourceTypebilling_address
const (
	TYPEBILLINGADDRESS_BILLING_ADDRESS TypeBillingAddress = "billing_address"
)

// AllowedTypeBillingAddressEnumValues All allowed values of TypeBillingAddress enum
var AllowedTypeBillingAddressEnumValues = []TypeBillingAddress{
	"billing_address",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeBillingAddress) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeBillingAddress(value)
	for _, existing := range AllowedTypeBillingAddressEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeBillingAddress", value)
}

// NewTypeBillingAddressFromValue returns a pointer to a valid TypeBillingAddress
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeBillingAddressFromValue(v string) (*TypeBillingAddress, error) {
	ev := TypeBillingAddress(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeBillingAddress: valid values are %v", v, AllowedTypeBillingAddressEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeBillingAddress) IsValid() bool {
	for _, existing := range AllowedTypeBillingAddressEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypebilling_address value
func (v TypeBillingAddress) Ptr() *TypeBillingAddress {
	return &v
}

// NullableTypeBillingAddress is a helper abstraction for handling nullable typebillingaddress types. 
type NullableTypeBillingAddress struct {
	value *TypeBillingAddress
	isSet bool
}

// Get returns the value.
func (v NullableTypeBillingAddress) Get() *TypeBillingAddress {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeBillingAddress) Set(val *TypeBillingAddress) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeBillingAddress) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeBillingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeBillingAddress returns a pointer to a new instance of NullableTypeBillingAddress.
func NewNullableTypeBillingAddress(val *TypeBillingAddress) *NullableTypeBillingAddress {
	return &NullableTypeBillingAddress{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeBillingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeBillingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
