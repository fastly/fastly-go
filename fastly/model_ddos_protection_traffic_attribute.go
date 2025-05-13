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

// DdosProtectionTrafficAttribute Name of an attribute type used in traffic stats.
type DdosProtectionTrafficAttribute string

// List of ddos_protection_traffic_attribute
const (
	DDOSPROTECTIONTRAFFICATTRIBUTE_SOURCE_IP        DdosProtectionTrafficAttribute = "source_ip"
	DDOSPROTECTIONTRAFFICATTRIBUTE_COUNTRY_CODE     DdosProtectionTrafficAttribute = "country_code"
	DDOSPROTECTIONTRAFFICATTRIBUTE_HOST             DdosProtectionTrafficAttribute = "host"
	DDOSPROTECTIONTRAFFICATTRIBUTE_ASN              DdosProtectionTrafficAttribute = "asn"
	DDOSPROTECTIONTRAFFICATTRIBUTE_SOURCE_IP_PREFIX DdosProtectionTrafficAttribute = "source_ip_prefix"
	DDOSPROTECTIONTRAFFICATTRIBUTE_USER_AGENT       DdosProtectionTrafficAttribute = "user_agent"
	DDOSPROTECTIONTRAFFICATTRIBUTE_METHOD_PATH      DdosProtectionTrafficAttribute = "method_path"
)

// AllowedDdosProtectionTrafficAttributeEnumValues All allowed values of DdosProtectionTrafficAttribute enum
var AllowedDdosProtectionTrafficAttributeEnumValues = []DdosProtectionTrafficAttribute{
	"source_ip",
	"country_code",
	"host",
	"asn",
	"source_ip_prefix",
	"user_agent",
	"method_path",
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *DdosProtectionTrafficAttribute) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DdosProtectionTrafficAttribute(value)
	for _, existing := range AllowedDdosProtectionTrafficAttributeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DdosProtectionTrafficAttribute", value)
}

// NewDdosProtectionTrafficAttributeFromValue returns a pointer to a valid DdosProtectionTrafficAttribute
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDdosProtectionTrafficAttributeFromValue(v string) (*DdosProtectionTrafficAttribute, error) {
	ev := DdosProtectionTrafficAttribute(v)
	if ev.IsValid() {
		return &ev, nil
	}
	return nil, fmt.Errorf("invalid value '%v' for DdosProtectionTrafficAttribute: valid values are %v", v, AllowedDdosProtectionTrafficAttributeEnumValues)
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DdosProtectionTrafficAttribute) IsValid() bool {
	for _, existing := range AllowedDdosProtectionTrafficAttributeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ddos_protection_traffic_attribute value
func (v DdosProtectionTrafficAttribute) Ptr() *DdosProtectionTrafficAttribute {
	return &v
}

// NullableDdosProtectionTrafficAttribute is a helper abstraction for handling nullable ddosprotectiontrafficattribute types.
type NullableDdosProtectionTrafficAttribute struct {
	value *DdosProtectionTrafficAttribute
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionTrafficAttribute) Get() *DdosProtectionTrafficAttribute {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionTrafficAttribute) Set(val *DdosProtectionTrafficAttribute) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionTrafficAttribute) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionTrafficAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionTrafficAttribute returns a pointer to a new instance of NullableDdosProtectionTrafficAttribute.
func NewNullableDdosProtectionTrafficAttribute(val *DdosProtectionTrafficAttribute) *NullableDdosProtectionTrafficAttribute {
	return &NullableDdosProtectionTrafficAttribute{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionTrafficAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionTrafficAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
