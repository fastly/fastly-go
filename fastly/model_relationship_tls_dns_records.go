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

// RelationshipTLSDNSRecords struct for RelationshipTLSDNSRecords
type RelationshipTLSDNSRecords struct {
	DNSRecords *RelationshipTLSDNSRecordDNSRecord `json:"dns_records,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSDNSRecords RelationshipTLSDNSRecords

// NewRelationshipTLSDNSRecords instantiates a new RelationshipTLSDNSRecords object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSDNSRecords() *RelationshipTLSDNSRecords {
	this := RelationshipTLSDNSRecords{}
	return &this
}

// NewRelationshipTLSDNSRecordsWithDefaults instantiates a new RelationshipTLSDNSRecords object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSDNSRecordsWithDefaults() *RelationshipTLSDNSRecords {
	this := RelationshipTLSDNSRecords{}
	return &this
}

// GetDNSRecords returns the DNSRecords field value if set, zero value otherwise.
func (o *RelationshipTLSDNSRecords) GetDNSRecords() RelationshipTLSDNSRecordDNSRecord {
	if o == nil || o.DNSRecords == nil {
		var ret RelationshipTLSDNSRecordDNSRecord
		return ret
	}
	return *o.DNSRecords
}

// GetDNSRecordsOk returns a tuple with the DNSRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSDNSRecords) GetDNSRecordsOk() (*RelationshipTLSDNSRecordDNSRecord, bool) {
	if o == nil || o.DNSRecords == nil {
		return nil, false
	}
	return o.DNSRecords, true
}

// HasDNSRecords returns a boolean if a field has been set.
func (o *RelationshipTLSDNSRecords) HasDNSRecords() bool {
	if o != nil && o.DNSRecords != nil {
		return true
	}

	return false
}

// SetDNSRecords gets a reference to the given RelationshipTLSDNSRecordDNSRecord and assigns it to the DNSRecords field.
func (o *RelationshipTLSDNSRecords) SetDNSRecords(v RelationshipTLSDNSRecordDNSRecord) {
	o.DNSRecords = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSDNSRecords) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.DNSRecords != nil {
		toSerialize["dns_records"] = o.DNSRecords
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSDNSRecords) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSDNSRecords := _RelationshipTLSDNSRecords{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSDNSRecords); err == nil {
		*o = RelationshipTLSDNSRecords(varRelationshipTLSDNSRecords)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dns_records")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSDNSRecords is a helper abstraction for handling nullable relationshiptlsdnsrecords types. 
type NullableRelationshipTLSDNSRecords struct {
	value *RelationshipTLSDNSRecords
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSDNSRecords) Get() *RelationshipTLSDNSRecords {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSDNSRecords) Set(val *RelationshipTLSDNSRecords) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSDNSRecords) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSDNSRecords) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSDNSRecords returns a pointer to a new instance of NullableRelationshipTLSDNSRecords.
func NewNullableRelationshipTLSDNSRecords(val *RelationshipTLSDNSRecords) *NullableRelationshipTLSDNSRecords {
	return &NullableRelationshipTLSDNSRecords{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSDNSRecords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSDNSRecords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
