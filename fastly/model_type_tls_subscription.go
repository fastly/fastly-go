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

// TypeTlsSubscription Resource type
type TypeTlsSubscription string

// List of type_tls_subscription
const (
	TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION TypeTlsSubscription = "tls_subscription"
)

// AllowedTypeTlsSubscriptionEnumValues All allowed values of TypeTlsSubscription enum
var AllowedTypeTlsSubscriptionEnumValues = []TypeTlsSubscription{
	"tls_subscription",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeTlsSubscription) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTlsSubscription(value)
	for _, existing := range AllowedTypeTlsSubscriptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTlsSubscription", value)
}

// NewTypeTlsSubscriptionFromValue returns a pointer to a valid TypeTlsSubscription
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTlsSubscriptionFromValue(v string) (*TypeTlsSubscription, error) {
	ev := TypeTlsSubscription(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeTlsSubscription: valid values are %v", v, AllowedTypeTlsSubscriptionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTlsSubscription) IsValid() bool {
	for _, existing := range AllowedTypeTlsSubscriptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to type_tls_subscription value
func (v TypeTlsSubscription) Ptr() *TypeTlsSubscription {
	return &v
}

// NullableTypeTlsSubscription is a helper abstraction for handling nullable typetlssubscription types.
type NullableTypeTlsSubscription struct {
	value *TypeTlsSubscription
	isSet bool
}

// Get returns the value.
func (v NullableTypeTlsSubscription) Get() *TypeTlsSubscription {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTlsSubscription) Set(val *TypeTlsSubscription) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTlsSubscription) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTlsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTlsSubscription returns a pointer to a new instance of NullableTypeTlsSubscription.
func NewNullableTypeTlsSubscription(val *TypeTlsSubscription) *NullableTypeTlsSubscription {
	return &NullableTypeTlsSubscription{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTlsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeTlsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
