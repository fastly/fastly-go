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

// RelationshipTLSCertificates struct for RelationshipTLSCertificates
type RelationshipTLSCertificates struct {
	TLSCertificates      *RelationshipTLSCertificatesTLSCertificates `json:"tls_certificates,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSCertificates RelationshipTLSCertificates

// NewRelationshipTLSCertificates instantiates a new RelationshipTLSCertificates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSCertificates() *RelationshipTLSCertificates {
	this := RelationshipTLSCertificates{}
	return &this
}

// NewRelationshipTLSCertificatesWithDefaults instantiates a new RelationshipTLSCertificates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSCertificatesWithDefaults() *RelationshipTLSCertificates {
	this := RelationshipTLSCertificates{}
	return &this
}

// GetTLSCertificates returns the TLSCertificates field value if set, zero value otherwise.
func (o *RelationshipTLSCertificates) GetTLSCertificates() RelationshipTLSCertificatesTLSCertificates {
	if o == nil || o.TLSCertificates == nil {
		var ret RelationshipTLSCertificatesTLSCertificates
		return ret
	}
	return *o.TLSCertificates
}

// GetTLSCertificatesOk returns a tuple with the TLSCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSCertificates) GetTLSCertificatesOk() (*RelationshipTLSCertificatesTLSCertificates, bool) {
	if o == nil || o.TLSCertificates == nil {
		return nil, false
	}
	return o.TLSCertificates, true
}

// HasTLSCertificates returns a boolean if a field has been set.
func (o *RelationshipTLSCertificates) HasTLSCertificates() bool {
	if o != nil && o.TLSCertificates != nil {
		return true
	}

	return false
}

// SetTLSCertificates gets a reference to the given RelationshipTLSCertificatesTLSCertificates and assigns it to the TLSCertificates field.
func (o *RelationshipTLSCertificates) SetTLSCertificates(v RelationshipTLSCertificatesTLSCertificates) {
	o.TLSCertificates = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSCertificates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSCertificates != nil {
		toSerialize["tls_certificates"] = o.TLSCertificates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTLSCertificates) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSCertificates := _RelationshipTLSCertificates{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSCertificates); err == nil {
		*o = RelationshipTLSCertificates(varRelationshipTLSCertificates)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_certificates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSCertificates is a helper abstraction for handling nullable relationshiptlscertificates types.
type NullableRelationshipTLSCertificates struct {
	value *RelationshipTLSCertificates
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSCertificates) Get() *RelationshipTLSCertificates {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSCertificates) Set(val *RelationshipTLSCertificates) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSCertificates) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSCertificates) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSCertificates returns a pointer to a new instance of NullableRelationshipTLSCertificates.
func NewNullableRelationshipTLSCertificates(val *RelationshipTLSCertificates) *NullableRelationshipTLSCertificates {
	return &NullableRelationshipTLSCertificates{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSCertificates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTLSCertificates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
