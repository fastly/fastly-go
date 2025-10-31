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

// RelationshipTlsDnsRecordsRequestDnsRecordsData struct for RelationshipTlsDnsRecordsRequestDnsRecordsData
type RelationshipTlsDnsRecordsRequestDnsRecordsData struct {
	Type                 *TypeTlsDnsRecord `json:"type,omitempty"`
	Id                   *string           `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsDnsRecordsRequestDnsRecordsData RelationshipTlsDnsRecordsRequestDnsRecordsData

// NewRelationshipTlsDnsRecordsRequestDnsRecordsData instantiates a new RelationshipTlsDnsRecordsRequestDnsRecordsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsDnsRecordsRequestDnsRecordsData() *RelationshipTlsDnsRecordsRequestDnsRecordsData {
	this := RelationshipTlsDnsRecordsRequestDnsRecordsData{}
	var type_ TypeTlsDnsRecord = TYPETLSDNSRECORD_DNS_RECORD
	this.Type = &type_
	return &this
}

// NewRelationshipTlsDnsRecordsRequestDnsRecordsDataWithDefaults instantiates a new RelationshipTlsDnsRecordsRequestDnsRecordsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsDnsRecordsRequestDnsRecordsDataWithDefaults() *RelationshipTlsDnsRecordsRequestDnsRecordsData {
	this := RelationshipTlsDnsRecordsRequestDnsRecordsData{}
	var type_ TypeTlsDnsRecord = TYPETLSDNSRECORD_DNS_RECORD
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipTlsDnsRecordsRequestDnsRecordsData) GetType() TypeTlsDnsRecord {
	if o == nil || o.Type == nil {
		var ret TypeTlsDnsRecord
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsDnsRecordsRequestDnsRecordsData) GetTypeOk() (*TypeTlsDnsRecord, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipTlsDnsRecordsRequestDnsRecordsData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsDnsRecord and assigns it to the Type field.
func (o *RelationshipTlsDnsRecordsRequestDnsRecordsData) SetType(v TypeTlsDnsRecord) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RelationshipTlsDnsRecordsRequestDnsRecordsData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsDnsRecordsRequestDnsRecordsData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RelationshipTlsDnsRecordsRequestDnsRecordsData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RelationshipTlsDnsRecordsRequestDnsRecordsData) SetId(v string) {
	o.Id = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsDnsRecordsRequestDnsRecordsData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsDnsRecordsRequestDnsRecordsData) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsDnsRecordsRequestDnsRecordsData := _RelationshipTlsDnsRecordsRequestDnsRecordsData{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsDnsRecordsRequestDnsRecordsData); err == nil {
		*o = RelationshipTlsDnsRecordsRequestDnsRecordsData(varRelationshipTlsDnsRecordsRequestDnsRecordsData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsDnsRecordsRequestDnsRecordsData is a helper abstraction for handling nullable relationshiptlsdnsrecordsrequestdnsrecordsdata types.
type NullableRelationshipTlsDnsRecordsRequestDnsRecordsData struct {
	value *RelationshipTlsDnsRecordsRequestDnsRecordsData
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsDnsRecordsRequestDnsRecordsData) Get() *RelationshipTlsDnsRecordsRequestDnsRecordsData {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsDnsRecordsRequestDnsRecordsData) Set(val *RelationshipTlsDnsRecordsRequestDnsRecordsData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsDnsRecordsRequestDnsRecordsData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsDnsRecordsRequestDnsRecordsData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsDnsRecordsRequestDnsRecordsData returns a pointer to a new instance of NullableRelationshipTlsDnsRecordsRequestDnsRecordsData.
func NewNullableRelationshipTlsDnsRecordsRequestDnsRecordsData(val *RelationshipTlsDnsRecordsRequestDnsRecordsData) *NullableRelationshipTlsDnsRecordsRequestDnsRecordsData {
	return &NullableRelationshipTlsDnsRecordsRequestDnsRecordsData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsDnsRecordsRequestDnsRecordsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsDnsRecordsRequestDnsRecordsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
