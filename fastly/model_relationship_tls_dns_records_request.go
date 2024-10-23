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

// RelationshipTLSDNSRecordsRequest struct for RelationshipTLSDNSRecordsRequest
type RelationshipTLSDNSRecordsRequest struct {
	DNSRecords           *RelationshipTLSDNSRecordsRequestDNSRecords `json:"dns_records,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSDNSRecordsRequest RelationshipTLSDNSRecordsRequest

// NewRelationshipTLSDNSRecordsRequest instantiates a new RelationshipTLSDNSRecordsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSDNSRecordsRequest() *RelationshipTLSDNSRecordsRequest {
	this := RelationshipTLSDNSRecordsRequest{}
	return &this
}

// NewRelationshipTLSDNSRecordsRequestWithDefaults instantiates a new RelationshipTLSDNSRecordsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSDNSRecordsRequestWithDefaults() *RelationshipTLSDNSRecordsRequest {
	this := RelationshipTLSDNSRecordsRequest{}
	return &this
}

// GetDNSRecords returns the DNSRecords field value if set, zero value otherwise.
func (o *RelationshipTLSDNSRecordsRequest) GetDNSRecords() RelationshipTLSDNSRecordsRequestDNSRecords {
	if o == nil || o.DNSRecords == nil {
		var ret RelationshipTLSDNSRecordsRequestDNSRecords
		return ret
	}
	return *o.DNSRecords
}

// GetDNSRecordsOk returns a tuple with the DNSRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSDNSRecordsRequest) GetDNSRecordsOk() (*RelationshipTLSDNSRecordsRequestDNSRecords, bool) {
	if o == nil || o.DNSRecords == nil {
		return nil, false
	}
	return o.DNSRecords, true
}

// HasDNSRecords returns a boolean if a field has been set.
func (o *RelationshipTLSDNSRecordsRequest) HasDNSRecords() bool {
	if o != nil && o.DNSRecords != nil {
		return true
	}

	return false
}

// SetDNSRecords gets a reference to the given RelationshipTLSDNSRecordsRequestDNSRecords and assigns it to the DNSRecords field.
func (o *RelationshipTLSDNSRecordsRequest) SetDNSRecords(v RelationshipTLSDNSRecordsRequestDNSRecords) {
	o.DNSRecords = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSDNSRecordsRequest) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipTLSDNSRecordsRequest) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSDNSRecordsRequest := _RelationshipTLSDNSRecordsRequest{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSDNSRecordsRequest); err == nil {
		*o = RelationshipTLSDNSRecordsRequest(varRelationshipTLSDNSRecordsRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dns_records")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSDNSRecordsRequest is a helper abstraction for handling nullable relationshiptlsdnsrecordsrequest types.
type NullableRelationshipTLSDNSRecordsRequest struct {
	value *RelationshipTLSDNSRecordsRequest
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSDNSRecordsRequest) Get() *RelationshipTLSDNSRecordsRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSDNSRecordsRequest) Set(val *RelationshipTLSDNSRecordsRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSDNSRecordsRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSDNSRecordsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSDNSRecordsRequest returns a pointer to a new instance of NullableRelationshipTLSDNSRecordsRequest.
func NewNullableRelationshipTLSDNSRecordsRequest(val *RelationshipTLSDNSRecordsRequest) *NullableRelationshipTLSDNSRecordsRequest {
	return &NullableRelationshipTLSDNSRecordsRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSDNSRecordsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTLSDNSRecordsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
