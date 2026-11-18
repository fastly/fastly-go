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

// RelationshipTLSDNSRecordsResponse struct for RelationshipTLSDNSRecordsResponse
type RelationshipTLSDNSRecordsResponse struct {
	DNSRecords           *RelationshipTLSDNSRecordsResponseDNSRecords `json:"dns_records,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSDNSRecordsResponse RelationshipTLSDNSRecordsResponse

// NewRelationshipTLSDNSRecordsResponse instantiates a new RelationshipTLSDNSRecordsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSDNSRecordsResponse() *RelationshipTLSDNSRecordsResponse {
	this := RelationshipTLSDNSRecordsResponse{}
	return &this
}

// NewRelationshipTLSDNSRecordsResponseWithDefaults instantiates a new RelationshipTLSDNSRecordsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSDNSRecordsResponseWithDefaults() *RelationshipTLSDNSRecordsResponse {
	this := RelationshipTLSDNSRecordsResponse{}
	return &this
}

// GetDNSRecords returns the DNSRecords field value if set, zero value otherwise.
func (o *RelationshipTLSDNSRecordsResponse) GetDNSRecords() RelationshipTLSDNSRecordsResponseDNSRecords {
	if o == nil || o.DNSRecords == nil {
		var ret RelationshipTLSDNSRecordsResponseDNSRecords
		return ret
	}
	return *o.DNSRecords
}

// GetDNSRecordsOk returns a tuple with the DNSRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSDNSRecordsResponse) GetDNSRecordsOk() (*RelationshipTLSDNSRecordsResponseDNSRecords, bool) {
	if o == nil || o.DNSRecords == nil {
		return nil, false
	}
	return o.DNSRecords, true
}

// HasDNSRecords returns a boolean if a field has been set.
func (o *RelationshipTLSDNSRecordsResponse) HasDNSRecords() bool {
	if o != nil && o.DNSRecords != nil {
		return true
	}

	return false
}

// SetDNSRecords gets a reference to the given RelationshipTLSDNSRecordsResponseDNSRecords and assigns it to the DNSRecords field.
func (o *RelationshipTLSDNSRecordsResponse) SetDNSRecords(v RelationshipTLSDNSRecordsResponseDNSRecords) {
	o.DNSRecords = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSDNSRecordsResponse) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipTLSDNSRecordsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSDNSRecordsResponse := _RelationshipTLSDNSRecordsResponse{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSDNSRecordsResponse); err == nil {
		*o = RelationshipTLSDNSRecordsResponse(varRelationshipTLSDNSRecordsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dns_records")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSDNSRecordsResponse is a helper abstraction for handling nullable relationshiptlsdnsrecordsresponse types.
type NullableRelationshipTLSDNSRecordsResponse struct {
	value *RelationshipTLSDNSRecordsResponse
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSDNSRecordsResponse) Get() *RelationshipTLSDNSRecordsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSDNSRecordsResponse) Set(val *RelationshipTLSDNSRecordsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSDNSRecordsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSDNSRecordsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSDNSRecordsResponse returns a pointer to a new instance of NullableRelationshipTLSDNSRecordsResponse.
func NewNullableRelationshipTLSDNSRecordsResponse(val *RelationshipTLSDNSRecordsResponse) *NullableRelationshipTLSDNSRecordsResponse {
	return &NullableRelationshipTLSDNSRecordsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSDNSRecordsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTLSDNSRecordsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
