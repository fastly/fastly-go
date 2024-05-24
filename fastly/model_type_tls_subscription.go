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

// TypeTLSSubscription Resource type
type TypeTLSSubscription string

// List of resourceTypetls_subscription
const (
	TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION TypeTLSSubscription = "tls_subscription"
)

// AllowedTypeTLSSubscriptionEnumValues All allowed values of TypeTLSSubscription enum
var AllowedTypeTLSSubscriptionEnumValues = []TypeTLSSubscription{
	"tls_subscription",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *TypeTLSSubscription) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTLSSubscription(value)
	for _, existing := range AllowedTypeTLSSubscriptionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTLSSubscription", value)
}

// NewTypeTLSSubscriptionFromValue returns a pointer to a valid TypeTLSSubscription
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTLSSubscriptionFromValue(v string) (*TypeTLSSubscription, error) {
	ev := TypeTLSSubscription(v)
	if ev.IsValid() {
		return &ev, nil
	}
  return nil, fmt.Errorf("invalid value '%v' for TypeTLSSubscription: valid values are %v", v, AllowedTypeTLSSubscriptionEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTLSSubscription) IsValid() bool {
	for _, existing := range AllowedTypeTLSSubscriptionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypetls_subscription value
func (v TypeTLSSubscription) Ptr() *TypeTLSSubscription {
	return &v
}

// NullableTypeTLSSubscription is a helper abstraction for handling nullable typetlssubscription types. 
type NullableTypeTLSSubscription struct {
	value *TypeTLSSubscription
	isSet bool
}

// Get returns the value.
func (v NullableTypeTLSSubscription) Get() *TypeTLSSubscription {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTLSSubscription) Set(val *TypeTLSSubscription) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTLSSubscription) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTLSSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTLSSubscription returns a pointer to a new instance of NullableTypeTLSSubscription.
func NewNullableTypeTLSSubscription(val *TypeTLSSubscription) *NullableTypeTLSSubscription {
	return &NullableTypeTLSSubscription{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTLSSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTypeTLSSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
