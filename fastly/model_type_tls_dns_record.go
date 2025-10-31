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

// TypeTlsDnsRecord Resource type
type TypeTlsDnsRecord string

// List of type_tls_dns_record
const (
	TYPETLSDNSRECORD_DNS_RECORD TypeTlsDnsRecord = "dns_record"
)

// AllowedTypeTlsDnsRecordEnumValues All allowed values of TypeTlsDnsRecord enum
var AllowedTypeTlsDnsRecordEnumValues = []TypeTlsDnsRecord{
	"dns_record",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeTlsDnsRecord) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTlsDnsRecord(value)
	for _, existing := range AllowedTypeTlsDnsRecordEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTlsDnsRecord", value)
}

// NewTypeTlsDnsRecordFromValue returns a pointer to a valid TypeTlsDnsRecord
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTlsDnsRecordFromValue(v string) (*TypeTlsDnsRecord, error) {
	ev := TypeTlsDnsRecord(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeTlsDnsRecord: valid values are %v", v, AllowedTypeTlsDnsRecordEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTlsDnsRecord) IsValid() bool {
	for _, existing := range AllowedTypeTlsDnsRecordEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to type_tls_dns_record value
func (v TypeTlsDnsRecord) Ptr() *TypeTlsDnsRecord {
	return &v
}

// NullableTypeTlsDnsRecord is a helper abstraction for handling nullable typetlsdnsrecord types.
type NullableTypeTlsDnsRecord struct {
	value *TypeTlsDnsRecord
	isSet bool
}

// Get returns the value.
func (v NullableTypeTlsDnsRecord) Get() *TypeTlsDnsRecord {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTlsDnsRecord) Set(val *TypeTlsDnsRecord) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTlsDnsRecord) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTlsDnsRecord) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTlsDnsRecord returns a pointer to a new instance of NullableTypeTlsDnsRecord.
func NewNullableTypeTlsDnsRecord(val *TypeTlsDnsRecord) *NullableTypeTlsDnsRecord {
	return &NullableTypeTlsDnsRecord{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTlsDnsRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeTlsDnsRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
