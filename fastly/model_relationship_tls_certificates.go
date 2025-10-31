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

// RelationshipTlsCertificates struct for RelationshipTlsCertificates
type RelationshipTlsCertificates struct {
	TlsCertificates      *RelationshipTlsCertificatesTlsCertificates `json:"tls_certificates,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsCertificates RelationshipTlsCertificates

// NewRelationshipTlsCertificates instantiates a new RelationshipTlsCertificates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsCertificates() *RelationshipTlsCertificates {
	this := RelationshipTlsCertificates{}
	return &this
}

// NewRelationshipTlsCertificatesWithDefaults instantiates a new RelationshipTlsCertificates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsCertificatesWithDefaults() *RelationshipTlsCertificates {
	this := RelationshipTlsCertificates{}
	return &this
}

// GetTlsCertificates returns the TlsCertificates field value if set, zero value otherwise.
func (o *RelationshipTlsCertificates) GetTlsCertificates() RelationshipTlsCertificatesTlsCertificates {
	if o == nil || o.TlsCertificates == nil {
		var ret RelationshipTlsCertificatesTlsCertificates
		return ret
	}
	return *o.TlsCertificates
}

// GetTlsCertificatesOk returns a tuple with the TlsCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsCertificates) GetTlsCertificatesOk() (*RelationshipTlsCertificatesTlsCertificates, bool) {
	if o == nil || o.TlsCertificates == nil {
		return nil, false
	}
	return o.TlsCertificates, true
}

// HasTlsCertificates returns a boolean if a field has been set.
func (o *RelationshipTlsCertificates) HasTlsCertificates() bool {
	if o != nil && o.TlsCertificates != nil {
		return true
	}

	return false
}

// SetTlsCertificates gets a reference to the given RelationshipTlsCertificatesTlsCertificates and assigns it to the TlsCertificates field.
func (o *RelationshipTlsCertificates) SetTlsCertificates(v RelationshipTlsCertificatesTlsCertificates) {
	o.TlsCertificates = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsCertificates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsCertificates != nil {
		toSerialize["tls_certificates"] = o.TlsCertificates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsCertificates) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsCertificates := _RelationshipTlsCertificates{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsCertificates); err == nil {
		*o = RelationshipTlsCertificates(varRelationshipTlsCertificates)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_certificates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsCertificates is a helper abstraction for handling nullable relationshiptlscertificates types.
type NullableRelationshipTlsCertificates struct {
	value *RelationshipTlsCertificates
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsCertificates) Get() *RelationshipTlsCertificates {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsCertificates) Set(val *RelationshipTlsCertificates) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsCertificates) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsCertificates) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsCertificates returns a pointer to a new instance of NullableRelationshipTlsCertificates.
func NewNullableRelationshipTlsCertificates(val *RelationshipTlsCertificates) *NullableRelationshipTlsCertificates {
	return &NullableRelationshipTlsCertificates{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsCertificates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsCertificates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
