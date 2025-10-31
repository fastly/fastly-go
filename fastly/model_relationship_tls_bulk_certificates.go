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

// RelationshipTlsBulkCertificates struct for RelationshipTlsBulkCertificates
type RelationshipTlsBulkCertificates struct {
	TlsBulkCertificates  *RelationshipTlsBulkCertificateTlsBulkCertificate `json:"tls_bulk_certificates,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsBulkCertificates RelationshipTlsBulkCertificates

// NewRelationshipTlsBulkCertificates instantiates a new RelationshipTlsBulkCertificates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsBulkCertificates() *RelationshipTlsBulkCertificates {
	this := RelationshipTlsBulkCertificates{}
	return &this
}

// NewRelationshipTlsBulkCertificatesWithDefaults instantiates a new RelationshipTlsBulkCertificates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsBulkCertificatesWithDefaults() *RelationshipTlsBulkCertificates {
	this := RelationshipTlsBulkCertificates{}
	return &this
}

// GetTlsBulkCertificates returns the TlsBulkCertificates field value if set, zero value otherwise.
func (o *RelationshipTlsBulkCertificates) GetTlsBulkCertificates() RelationshipTlsBulkCertificateTlsBulkCertificate {
	if o == nil || o.TlsBulkCertificates == nil {
		var ret RelationshipTlsBulkCertificateTlsBulkCertificate
		return ret
	}
	return *o.TlsBulkCertificates
}

// GetTlsBulkCertificatesOk returns a tuple with the TlsBulkCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsBulkCertificates) GetTlsBulkCertificatesOk() (*RelationshipTlsBulkCertificateTlsBulkCertificate, bool) {
	if o == nil || o.TlsBulkCertificates == nil {
		return nil, false
	}
	return o.TlsBulkCertificates, true
}

// HasTlsBulkCertificates returns a boolean if a field has been set.
func (o *RelationshipTlsBulkCertificates) HasTlsBulkCertificates() bool {
	if o != nil && o.TlsBulkCertificates != nil {
		return true
	}

	return false
}

// SetTlsBulkCertificates gets a reference to the given RelationshipTlsBulkCertificateTlsBulkCertificate and assigns it to the TlsBulkCertificates field.
func (o *RelationshipTlsBulkCertificates) SetTlsBulkCertificates(v RelationshipTlsBulkCertificateTlsBulkCertificate) {
	o.TlsBulkCertificates = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsBulkCertificates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsBulkCertificates != nil {
		toSerialize["tls_bulk_certificates"] = o.TlsBulkCertificates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsBulkCertificates) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsBulkCertificates := _RelationshipTlsBulkCertificates{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsBulkCertificates); err == nil {
		*o = RelationshipTlsBulkCertificates(varRelationshipTlsBulkCertificates)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_bulk_certificates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsBulkCertificates is a helper abstraction for handling nullable relationshiptlsbulkcertificates types.
type NullableRelationshipTlsBulkCertificates struct {
	value *RelationshipTlsBulkCertificates
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsBulkCertificates) Get() *RelationshipTlsBulkCertificates {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsBulkCertificates) Set(val *RelationshipTlsBulkCertificates) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsBulkCertificates) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsBulkCertificates) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsBulkCertificates returns a pointer to a new instance of NullableRelationshipTlsBulkCertificates.
func NewNullableRelationshipTlsBulkCertificates(val *RelationshipTlsBulkCertificates) *NullableRelationshipTlsBulkCertificates {
	return &NullableRelationshipTlsBulkCertificates{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsBulkCertificates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsBulkCertificates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
