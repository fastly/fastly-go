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

// RelationshipTLSDNSRecordsResponseDNSRecordsData struct for RelationshipTLSDNSRecordsResponseDNSRecordsData
type RelationshipTLSDNSRecordsResponseDNSRecordsData struct {
	Type                 *TypeTLSDNSRecord `json:"type,omitempty"`
	ID                   *string           `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSDNSRecordsResponseDNSRecordsData RelationshipTLSDNSRecordsResponseDNSRecordsData

// NewRelationshipTLSDNSRecordsResponseDNSRecordsData instantiates a new RelationshipTLSDNSRecordsResponseDNSRecordsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSDNSRecordsResponseDNSRecordsData() *RelationshipTLSDNSRecordsResponseDNSRecordsData {
	this := RelationshipTLSDNSRecordsResponseDNSRecordsData{}
	var resourceType TypeTLSDNSRecord = TYPETLSDNSRECORD_DNS_RECORD
	this.Type = &resourceType
	return &this
}

// NewRelationshipTLSDNSRecordsResponseDNSRecordsDataWithDefaults instantiates a new RelationshipTLSDNSRecordsResponseDNSRecordsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSDNSRecordsResponseDNSRecordsDataWithDefaults() *RelationshipTLSDNSRecordsResponseDNSRecordsData {
	this := RelationshipTLSDNSRecordsResponseDNSRecordsData{}
	var resourceType TypeTLSDNSRecord = TYPETLSDNSRECORD_DNS_RECORD
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) GetType() TypeTLSDNSRecord {
	if o == nil || o.Type == nil {
		var ret TypeTLSDNSRecord
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) GetTypeOk() (*TypeTLSDNSRecord, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSDNSRecord and assigns it to the Type field.
func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) SetType(v TypeTLSDNSRecord) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSDNSRecordsResponseDNSRecordsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTLSDNSRecordsResponseDNSRecordsData) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSDNSRecordsResponseDNSRecordsData := _RelationshipTLSDNSRecordsResponseDNSRecordsData{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSDNSRecordsResponseDNSRecordsData); err == nil {
		*o = RelationshipTLSDNSRecordsResponseDNSRecordsData(varRelationshipTLSDNSRecordsResponseDNSRecordsData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSDNSRecordsResponseDNSRecordsData is a helper abstraction for handling nullable relationshiptlsdnsrecordsresponsednsrecordsdata types.
type NullableRelationshipTLSDNSRecordsResponseDNSRecordsData struct {
	value *RelationshipTLSDNSRecordsResponseDNSRecordsData
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSDNSRecordsResponseDNSRecordsData) Get() *RelationshipTLSDNSRecordsResponseDNSRecordsData {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSDNSRecordsResponseDNSRecordsData) Set(val *RelationshipTLSDNSRecordsResponseDNSRecordsData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSDNSRecordsResponseDNSRecordsData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSDNSRecordsResponseDNSRecordsData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSDNSRecordsResponseDNSRecordsData returns a pointer to a new instance of NullableRelationshipTLSDNSRecordsResponseDNSRecordsData.
func NewNullableRelationshipTLSDNSRecordsResponseDNSRecordsData(val *RelationshipTLSDNSRecordsResponseDNSRecordsData) *NullableRelationshipTLSDNSRecordsResponseDNSRecordsData {
	return &NullableRelationshipTLSDNSRecordsResponseDNSRecordsData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSDNSRecordsResponseDNSRecordsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTLSDNSRecordsResponseDNSRecordsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
