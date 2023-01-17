// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// RelationshipTLSDomains struct for RelationshipTLSDomains
type RelationshipTLSDomains struct {
	TLSDomains *RelationshipTLSDomainsTLSDomains `json:"tls_domains,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSDomains RelationshipTLSDomains

// NewRelationshipTLSDomains instantiates a new RelationshipTLSDomains object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSDomains() *RelationshipTLSDomains {
	this := RelationshipTLSDomains{}
	return &this
}

// NewRelationshipTLSDomainsWithDefaults instantiates a new RelationshipTLSDomains object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSDomainsWithDefaults() *RelationshipTLSDomains {
	this := RelationshipTLSDomains{}
	return &this
}

// GetTLSDomains returns the TLSDomains field value if set, zero value otherwise.
func (o *RelationshipTLSDomains) GetTLSDomains() RelationshipTLSDomainsTLSDomains {
	if o == nil || o.TLSDomains == nil {
		var ret RelationshipTLSDomainsTLSDomains
		return ret
	}
	return *o.TLSDomains
}

// GetTLSDomainsOk returns a tuple with the TLSDomains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSDomains) GetTLSDomainsOk() (*RelationshipTLSDomainsTLSDomains, bool) {
	if o == nil || o.TLSDomains == nil {
		return nil, false
	}
	return o.TLSDomains, true
}

// HasTLSDomains returns a boolean if a field has been set.
func (o *RelationshipTLSDomains) HasTLSDomains() bool {
	if o != nil && o.TLSDomains != nil {
		return true
	}

	return false
}

// SetTLSDomains gets a reference to the given RelationshipTLSDomainsTLSDomains and assigns it to the TLSDomains field.
func (o *RelationshipTLSDomains) SetTLSDomains(v RelationshipTLSDomainsTLSDomains) {
	o.TLSDomains = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSDomains) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSDomains != nil {
		toSerialize["tls_domains"] = o.TLSDomains
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSDomains) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSDomains := _RelationshipTLSDomains{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSDomains); err == nil {
		*o = RelationshipTLSDomains(varRelationshipTLSDomains)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_domains")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSDomains is a helper abstraction for handling nullable relationshiptlsdomains types. 
type NullableRelationshipTLSDomains struct {
	value *RelationshipTLSDomains
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSDomains) Get() *RelationshipTLSDomains {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSDomains) Set(val *RelationshipTLSDomains) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSDomains) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSDomains) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSDomains returns a pointer to a new instance of NullableRelationshipTLSDomains.
func NewNullableRelationshipTLSDomains(val *RelationshipTLSDomains) *NullableRelationshipTLSDomains {
	return &NullableRelationshipTLSDomains{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSDomains) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSDomains) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
