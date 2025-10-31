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

// RelationshipTlsDomains struct for RelationshipTlsDomains
type RelationshipTlsDomains struct {
	TlsDomains           *RelationshipTlsDomainsTlsDomains `json:"tls_domains,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsDomains RelationshipTlsDomains

// NewRelationshipTlsDomains instantiates a new RelationshipTlsDomains object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsDomains() *RelationshipTlsDomains {
	this := RelationshipTlsDomains{}
	return &this
}

// NewRelationshipTlsDomainsWithDefaults instantiates a new RelationshipTlsDomains object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsDomainsWithDefaults() *RelationshipTlsDomains {
	this := RelationshipTlsDomains{}
	return &this
}

// GetTlsDomains returns the TlsDomains field value if set, zero value otherwise.
func (o *RelationshipTlsDomains) GetTlsDomains() RelationshipTlsDomainsTlsDomains {
	if o == nil || o.TlsDomains == nil {
		var ret RelationshipTlsDomainsTlsDomains
		return ret
	}
	return *o.TlsDomains
}

// GetTlsDomainsOk returns a tuple with the TlsDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsDomains) GetTlsDomainsOk() (*RelationshipTlsDomainsTlsDomains, bool) {
	if o == nil || o.TlsDomains == nil {
		return nil, false
	}
	return o.TlsDomains, true
}

// HasTlsDomains returns a boolean if a field has been set.
func (o *RelationshipTlsDomains) HasTlsDomains() bool {
	if o != nil && o.TlsDomains != nil {
		return true
	}

	return false
}

// SetTlsDomains gets a reference to the given RelationshipTlsDomainsTlsDomains and assigns it to the TlsDomains field.
func (o *RelationshipTlsDomains) SetTlsDomains(v RelationshipTlsDomainsTlsDomains) {
	o.TlsDomains = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsDomains) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsDomains != nil {
		toSerialize["tls_domains"] = o.TlsDomains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsDomains) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsDomains := _RelationshipTlsDomains{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsDomains); err == nil {
		*o = RelationshipTlsDomains(varRelationshipTlsDomains)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_domains")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsDomains is a helper abstraction for handling nullable relationshiptlsdomains types.
type NullableRelationshipTlsDomains struct {
	value *RelationshipTlsDomains
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsDomains) Get() *RelationshipTlsDomains {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsDomains) Set(val *RelationshipTlsDomains) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsDomains) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsDomains) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsDomains returns a pointer to a new instance of NullableRelationshipTlsDomains.
func NewNullableRelationshipTlsDomains(val *RelationshipTlsDomains) *NullableRelationshipTlsDomains {
	return &NullableRelationshipTlsDomains{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsDomains) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsDomains) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
