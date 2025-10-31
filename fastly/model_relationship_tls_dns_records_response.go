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

// RelationshipTlsDnsRecordsResponse struct for RelationshipTlsDnsRecordsResponse
type RelationshipTlsDnsRecordsResponse struct {
	DnsRecords           *RelationshipTlsDnsRecordsResponseDnsRecords `json:"dns_records,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsDnsRecordsResponse RelationshipTlsDnsRecordsResponse

// NewRelationshipTlsDnsRecordsResponse instantiates a new RelationshipTlsDnsRecordsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsDnsRecordsResponse() *RelationshipTlsDnsRecordsResponse {
	this := RelationshipTlsDnsRecordsResponse{}
	return &this
}

// NewRelationshipTlsDnsRecordsResponseWithDefaults instantiates a new RelationshipTlsDnsRecordsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsDnsRecordsResponseWithDefaults() *RelationshipTlsDnsRecordsResponse {
	this := RelationshipTlsDnsRecordsResponse{}
	return &this
}

// GetDnsRecords returns the DnsRecords field value if set, zero value otherwise.
func (o *RelationshipTlsDnsRecordsResponse) GetDnsRecords() RelationshipTlsDnsRecordsResponseDnsRecords {
	if o == nil || o.DnsRecords == nil {
		var ret RelationshipTlsDnsRecordsResponseDnsRecords
		return ret
	}
	return *o.DnsRecords
}

// GetDnsRecordsOk returns a tuple with the DnsRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsDnsRecordsResponse) GetDnsRecordsOk() (*RelationshipTlsDnsRecordsResponseDnsRecords, bool) {
	if o == nil || o.DnsRecords == nil {
		return nil, false
	}
	return o.DnsRecords, true
}

// HasDnsRecords returns a boolean if a field has been set.
func (o *RelationshipTlsDnsRecordsResponse) HasDnsRecords() bool {
	if o != nil && o.DnsRecords != nil {
		return true
	}

	return false
}

// SetDnsRecords gets a reference to the given RelationshipTlsDnsRecordsResponseDnsRecords and assigns it to the DnsRecords field.
func (o *RelationshipTlsDnsRecordsResponse) SetDnsRecords(v RelationshipTlsDnsRecordsResponseDnsRecords) {
	o.DnsRecords = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsDnsRecordsResponse) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipTlsDnsRecordsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsDnsRecordsResponse := _RelationshipTlsDnsRecordsResponse{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsDnsRecordsResponse); err == nil {
		*o = RelationshipTlsDnsRecordsResponse(varRelationshipTlsDnsRecordsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dns_records")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsDnsRecordsResponse is a helper abstraction for handling nullable relationshiptlsdnsrecordsresponse types.
type NullableRelationshipTlsDnsRecordsResponse struct {
	value *RelationshipTlsDnsRecordsResponse
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsDnsRecordsResponse) Get() *RelationshipTlsDnsRecordsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsDnsRecordsResponse) Set(val *RelationshipTlsDnsRecordsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsDnsRecordsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsDnsRecordsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsDnsRecordsResponse returns a pointer to a new instance of NullableRelationshipTlsDnsRecordsResponse.
func NewNullableRelationshipTlsDnsRecordsResponse(val *RelationshipTlsDnsRecordsResponse) *NullableRelationshipTlsDnsRecordsResponse {
	return &NullableRelationshipTlsDnsRecordsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsDnsRecordsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsDnsRecordsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
