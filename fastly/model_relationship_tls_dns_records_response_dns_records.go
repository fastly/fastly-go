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

// RelationshipTlsDnsRecordsResponseDnsRecords struct for RelationshipTlsDnsRecordsResponseDnsRecords
type RelationshipTlsDnsRecordsResponseDnsRecords struct {
	Data                 []RelationshipTlsDnsRecordsResponseDnsRecordsData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsDnsRecordsResponseDnsRecords RelationshipTlsDnsRecordsResponseDnsRecords

// NewRelationshipTlsDnsRecordsResponseDnsRecords instantiates a new RelationshipTlsDnsRecordsResponseDnsRecords object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsDnsRecordsResponseDnsRecords() *RelationshipTlsDnsRecordsResponseDnsRecords {
	this := RelationshipTlsDnsRecordsResponseDnsRecords{}
	return &this
}

// NewRelationshipTlsDnsRecordsResponseDnsRecordsWithDefaults instantiates a new RelationshipTlsDnsRecordsResponseDnsRecords object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsDnsRecordsResponseDnsRecordsWithDefaults() *RelationshipTlsDnsRecordsResponseDnsRecords {
	this := RelationshipTlsDnsRecordsResponseDnsRecords{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RelationshipTlsDnsRecordsResponseDnsRecords) GetData() []RelationshipTlsDnsRecordsResponseDnsRecordsData {
	if o == nil || o.Data == nil {
		var ret []RelationshipTlsDnsRecordsResponseDnsRecordsData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsDnsRecordsResponseDnsRecords) GetDataOk() ([]RelationshipTlsDnsRecordsResponseDnsRecordsData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RelationshipTlsDnsRecordsResponseDnsRecords) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []RelationshipTlsDnsRecordsResponseDnsRecordsData and assigns it to the Data field.
func (o *RelationshipTlsDnsRecordsResponseDnsRecords) SetData(v []RelationshipTlsDnsRecordsResponseDnsRecordsData) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsDnsRecordsResponseDnsRecords) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsDnsRecordsResponseDnsRecords) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsDnsRecordsResponseDnsRecords := _RelationshipTlsDnsRecordsResponseDnsRecords{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsDnsRecordsResponseDnsRecords); err == nil {
		*o = RelationshipTlsDnsRecordsResponseDnsRecords(varRelationshipTlsDnsRecordsResponseDnsRecords)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsDnsRecordsResponseDnsRecords is a helper abstraction for handling nullable relationshiptlsdnsrecordsresponsednsrecords types.
type NullableRelationshipTlsDnsRecordsResponseDnsRecords struct {
	value *RelationshipTlsDnsRecordsResponseDnsRecords
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsDnsRecordsResponseDnsRecords) Get() *RelationshipTlsDnsRecordsResponseDnsRecords {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsDnsRecordsResponseDnsRecords) Set(val *RelationshipTlsDnsRecordsResponseDnsRecords) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsDnsRecordsResponseDnsRecords) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsDnsRecordsResponseDnsRecords) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsDnsRecordsResponseDnsRecords returns a pointer to a new instance of NullableRelationshipTlsDnsRecordsResponseDnsRecords.
func NewNullableRelationshipTlsDnsRecordsResponseDnsRecords(val *RelationshipTlsDnsRecordsResponseDnsRecords) *NullableRelationshipTlsDnsRecordsResponseDnsRecords {
	return &NullableRelationshipTlsDnsRecordsResponseDnsRecords{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsDnsRecordsResponseDnsRecords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsDnsRecordsResponseDnsRecords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
