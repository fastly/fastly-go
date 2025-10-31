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

// RelationshipTlsDomain The TLS domain being enabled for TLS traffic. Required.
type RelationshipTlsDomain struct {
	TlsDomain            *RelationshipTlsDomainTlsDomain `json:"tls_domain,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsDomain RelationshipTlsDomain

// NewRelationshipTlsDomain instantiates a new RelationshipTlsDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsDomain() *RelationshipTlsDomain {
	this := RelationshipTlsDomain{}
	return &this
}

// NewRelationshipTlsDomainWithDefaults instantiates a new RelationshipTlsDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsDomainWithDefaults() *RelationshipTlsDomain {
	this := RelationshipTlsDomain{}
	return &this
}

// GetTlsDomain returns the TlsDomain field value if set, zero value otherwise.
func (o *RelationshipTlsDomain) GetTlsDomain() RelationshipTlsDomainTlsDomain {
	if o == nil || o.TlsDomain == nil {
		var ret RelationshipTlsDomainTlsDomain
		return ret
	}
	return *o.TlsDomain
}

// GetTlsDomainOk returns a tuple with the TlsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsDomain) GetTlsDomainOk() (*RelationshipTlsDomainTlsDomain, bool) {
	if o == nil || o.TlsDomain == nil {
		return nil, false
	}
	return o.TlsDomain, true
}

// HasTlsDomain returns a boolean if a field has been set.
func (o *RelationshipTlsDomain) HasTlsDomain() bool {
	if o != nil && o.TlsDomain != nil {
		return true
	}

	return false
}

// SetTlsDomain gets a reference to the given RelationshipTlsDomainTlsDomain and assigns it to the TlsDomain field.
func (o *RelationshipTlsDomain) SetTlsDomain(v RelationshipTlsDomainTlsDomain) {
	o.TlsDomain = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsDomain != nil {
		toSerialize["tls_domain"] = o.TlsDomain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsDomain) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsDomain := _RelationshipTlsDomain{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsDomain); err == nil {
		*o = RelationshipTlsDomain(varRelationshipTlsDomain)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_domain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsDomain is a helper abstraction for handling nullable relationshiptlsdomain types.
type NullableRelationshipTlsDomain struct {
	value *RelationshipTlsDomain
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsDomain) Get() *RelationshipTlsDomain {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsDomain) Set(val *RelationshipTlsDomain) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsDomain) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsDomain) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsDomain returns a pointer to a new instance of NullableRelationshipTlsDomain.
func NewNullableRelationshipTlsDomain(val *RelationshipTlsDomain) *NullableRelationshipTlsDomain {
	return &NullableRelationshipTlsDomain{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
