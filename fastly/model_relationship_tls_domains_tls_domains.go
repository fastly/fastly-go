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

// RelationshipTlsDomainsTlsDomains struct for RelationshipTlsDomainsTlsDomains
type RelationshipTlsDomainsTlsDomains struct {
	Data                 []RelationshipMemberTlsDomain `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsDomainsTlsDomains RelationshipTlsDomainsTlsDomains

// NewRelationshipTlsDomainsTlsDomains instantiates a new RelationshipTlsDomainsTlsDomains object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsDomainsTlsDomains() *RelationshipTlsDomainsTlsDomains {
	this := RelationshipTlsDomainsTlsDomains{}
	return &this
}

// NewRelationshipTlsDomainsTlsDomainsWithDefaults instantiates a new RelationshipTlsDomainsTlsDomains object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsDomainsTlsDomainsWithDefaults() *RelationshipTlsDomainsTlsDomains {
	this := RelationshipTlsDomainsTlsDomains{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RelationshipTlsDomainsTlsDomains) GetData() []RelationshipMemberTlsDomain {
	if o == nil || o.Data == nil {
		var ret []RelationshipMemberTlsDomain
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsDomainsTlsDomains) GetDataOk() ([]RelationshipMemberTlsDomain, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RelationshipTlsDomainsTlsDomains) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []RelationshipMemberTlsDomain and assigns it to the Data field.
func (o *RelationshipTlsDomainsTlsDomains) SetData(v []RelationshipMemberTlsDomain) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsDomainsTlsDomains) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipTlsDomainsTlsDomains) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsDomainsTlsDomains := _RelationshipTlsDomainsTlsDomains{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsDomainsTlsDomains); err == nil {
		*o = RelationshipTlsDomainsTlsDomains(varRelationshipTlsDomainsTlsDomains)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsDomainsTlsDomains is a helper abstraction for handling nullable relationshiptlsdomainstlsdomains types.
type NullableRelationshipTlsDomainsTlsDomains struct {
	value *RelationshipTlsDomainsTlsDomains
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsDomainsTlsDomains) Get() *RelationshipTlsDomainsTlsDomains {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsDomainsTlsDomains) Set(val *RelationshipTlsDomainsTlsDomains) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsDomainsTlsDomains) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsDomainsTlsDomains) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsDomainsTlsDomains returns a pointer to a new instance of NullableRelationshipTlsDomainsTlsDomains.
func NewNullableRelationshipTlsDomainsTlsDomains(val *RelationshipTlsDomainsTlsDomains) *NullableRelationshipTlsDomainsTlsDomains {
	return &NullableRelationshipTlsDomainsTlsDomains{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsDomainsTlsDomains) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsDomainsTlsDomains) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
