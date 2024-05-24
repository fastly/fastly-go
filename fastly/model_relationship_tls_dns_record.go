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
)

// RelationshipTLSDNSRecord struct for RelationshipTLSDNSRecord
type RelationshipTLSDNSRecord struct {
	DNSRecord *RelationshipTLSDNSRecordDNSRecord `json:"dns_record,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSDNSRecord RelationshipTLSDNSRecord

// NewRelationshipTLSDNSRecord instantiates a new RelationshipTLSDNSRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSDNSRecord() *RelationshipTLSDNSRecord {
	this := RelationshipTLSDNSRecord{}
	return &this
}

// NewRelationshipTLSDNSRecordWithDefaults instantiates a new RelationshipTLSDNSRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSDNSRecordWithDefaults() *RelationshipTLSDNSRecord {
	this := RelationshipTLSDNSRecord{}
	return &this
}

// GetDNSRecord returns the DNSRecord field value if set, zero value otherwise.
func (o *RelationshipTLSDNSRecord) GetDNSRecord() RelationshipTLSDNSRecordDNSRecord {
	if o == nil || o.DNSRecord == nil {
		var ret RelationshipTLSDNSRecordDNSRecord
		return ret
	}
	return *o.DNSRecord
}

// GetDNSRecordOk returns a tuple with the DNSRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSDNSRecord) GetDNSRecordOk() (*RelationshipTLSDNSRecordDNSRecord, bool) {
	if o == nil || o.DNSRecord == nil {
		return nil, false
	}
	return o.DNSRecord, true
}

// HasDNSRecord returns a boolean if a field has been set.
func (o *RelationshipTLSDNSRecord) HasDNSRecord() bool {
	if o != nil && o.DNSRecord != nil {
		return true
	}

	return false
}

// SetDNSRecord gets a reference to the given RelationshipTLSDNSRecordDNSRecord and assigns it to the DNSRecord field.
func (o *RelationshipTLSDNSRecord) SetDNSRecord(v RelationshipTLSDNSRecordDNSRecord) {
	o.DNSRecord = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSDNSRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.DNSRecord != nil {
		toSerialize["dns_record"] = o.DNSRecord
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSDNSRecord) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSDNSRecord := _RelationshipTLSDNSRecord{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSDNSRecord); err == nil {
		*o = RelationshipTLSDNSRecord(varRelationshipTLSDNSRecord)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dns_record")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSDNSRecord is a helper abstraction for handling nullable relationshiptlsdnsrecord types. 
type NullableRelationshipTLSDNSRecord struct {
	value *RelationshipTLSDNSRecord
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSDNSRecord) Get() *RelationshipTLSDNSRecord {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSDNSRecord) Set(val *RelationshipTLSDNSRecord) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSDNSRecord) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSDNSRecord) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSDNSRecord returns a pointer to a new instance of NullableRelationshipTLSDNSRecord.
func NewNullableRelationshipTLSDNSRecord(val *RelationshipTLSDNSRecord) *NullableRelationshipTLSDNSRecord {
	return &NullableRelationshipTLSDNSRecord{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSDNSRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSDNSRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
