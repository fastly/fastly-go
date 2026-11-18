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

// TypeTLSDNSRecord Resource type
type TypeTLSDNSRecord string

// List of resourceTypetls_dns_record
const (
	TYPETLSDNSRECORD_DNS_RECORD TypeTLSDNSRecord = "dns_record"
)

// AllowedTypeTLSDNSRecordEnumValues All allowed values of TypeTLSDNSRecord enum
var AllowedTypeTLSDNSRecordEnumValues = []TypeTLSDNSRecord{
	"dns_record",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *TypeTLSDNSRecord) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TypeTLSDNSRecord(value)
	for _, existing := range AllowedTypeTLSDNSRecordEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TypeTLSDNSRecord", value)
}

// NewTypeTLSDNSRecordFromValue returns a pointer to a valid TypeTLSDNSRecord
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTypeTLSDNSRecordFromValue(v string) (*TypeTLSDNSRecord, error) {
	ev := TypeTLSDNSRecord(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for TypeTLSDNSRecord: valid values are %v", v, AllowedTypeTLSDNSRecordEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TypeTLSDNSRecord) IsValid() bool {
	for _, existing := range AllowedTypeTLSDNSRecordEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to resourceTypetls_dns_record value
func (v TypeTLSDNSRecord) Ptr() *TypeTLSDNSRecord {
	return &v
}

// NullableTypeTLSDNSRecord is a helper abstraction for handling nullable typetlsdnsrecord types.
type NullableTypeTLSDNSRecord struct {
	value *TypeTLSDNSRecord
	isSet bool
}

// Get returns the value.
func (v NullableTypeTLSDNSRecord) Get() *TypeTLSDNSRecord {
	return v.value
}

// Set modifies the value.
func (v *NullableTypeTLSDNSRecord) Set(val *TypeTLSDNSRecord) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTypeTLSDNSRecord) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTypeTLSDNSRecord) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTypeTLSDNSRecord returns a pointer to a new instance of NullableTypeTLSDNSRecord.
func NewNullableTypeTLSDNSRecord(val *TypeTLSDNSRecord) *NullableTypeTLSDNSRecord {
	return &NullableTypeTLSDNSRecord{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTypeTLSDNSRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTypeTLSDNSRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
