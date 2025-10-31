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

// RelationshipTlsDnsRecordsRequest struct for RelationshipTlsDnsRecordsRequest
type RelationshipTlsDnsRecordsRequest struct {
	DnsRecords           *RelationshipTlsDnsRecordsRequestDnsRecords `json:"dns_records,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsDnsRecordsRequest RelationshipTlsDnsRecordsRequest

// NewRelationshipTlsDnsRecordsRequest instantiates a new RelationshipTlsDnsRecordsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsDnsRecordsRequest() *RelationshipTlsDnsRecordsRequest {
	this := RelationshipTlsDnsRecordsRequest{}
	return &this
}

// NewRelationshipTlsDnsRecordsRequestWithDefaults instantiates a new RelationshipTlsDnsRecordsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsDnsRecordsRequestWithDefaults() *RelationshipTlsDnsRecordsRequest {
	this := RelationshipTlsDnsRecordsRequest{}
	return &this
}

// GetDnsRecords returns the DnsRecords field value if set, zero value otherwise.
func (o *RelationshipTlsDnsRecordsRequest) GetDnsRecords() RelationshipTlsDnsRecordsRequestDnsRecords {
	if o == nil || o.DnsRecords == nil {
		var ret RelationshipTlsDnsRecordsRequestDnsRecords
		return ret
	}
	return *o.DnsRecords
}

// GetDnsRecordsOk returns a tuple with the DnsRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsDnsRecordsRequest) GetDnsRecordsOk() (*RelationshipTlsDnsRecordsRequestDnsRecords, bool) {
	if o == nil || o.DnsRecords == nil {
		return nil, false
	}
	return o.DnsRecords, true
}

// HasDnsRecords returns a boolean if a field has been set.
func (o *RelationshipTlsDnsRecordsRequest) HasDnsRecords() bool {
	if o != nil && o.DnsRecords != nil {
		return true
	}

	return false
}

// SetDnsRecords gets a reference to the given RelationshipTlsDnsRecordsRequestDnsRecords and assigns it to the DnsRecords field.
func (o *RelationshipTlsDnsRecordsRequest) SetDnsRecords(v RelationshipTlsDnsRecordsRequestDnsRecords) {
	o.DnsRecords = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsDnsRecordsRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.DnsRecords != nil {
		toSerialize["dns_records"] = o.DnsRecords
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsDnsRecordsRequest) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsDnsRecordsRequest := _RelationshipTlsDnsRecordsRequest{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsDnsRecordsRequest); err == nil {
		*o = RelationshipTlsDnsRecordsRequest(varRelationshipTlsDnsRecordsRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dns_records")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsDnsRecordsRequest is a helper abstraction for handling nullable relationshiptlsdnsrecordsrequest types.
type NullableRelationshipTlsDnsRecordsRequest struct {
	value *RelationshipTlsDnsRecordsRequest
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsDnsRecordsRequest) Get() *RelationshipTlsDnsRecordsRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsDnsRecordsRequest) Set(val *RelationshipTlsDnsRecordsRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsDnsRecordsRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsDnsRecordsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsDnsRecordsRequest returns a pointer to a new instance of NullableRelationshipTlsDnsRecordsRequest.
func NewNullableRelationshipTlsDnsRecordsRequest(val *RelationshipTlsDnsRecordsRequest) *NullableRelationshipTlsDnsRecordsRequest {
	return &NullableRelationshipTlsDnsRecordsRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsDnsRecordsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsDnsRecordsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
