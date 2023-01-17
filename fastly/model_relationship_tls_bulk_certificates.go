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

// RelationshipTLSBulkCertificates struct for RelationshipTLSBulkCertificates
type RelationshipTLSBulkCertificates struct {
	TLSBulkCertificates *RelationshipTLSBulkCertificateTLSBulkCertificate `json:"tls_bulk_certificates,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSBulkCertificates RelationshipTLSBulkCertificates

// NewRelationshipTLSBulkCertificates instantiates a new RelationshipTLSBulkCertificates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSBulkCertificates() *RelationshipTLSBulkCertificates {
	this := RelationshipTLSBulkCertificates{}
	return &this
}

// NewRelationshipTLSBulkCertificatesWithDefaults instantiates a new RelationshipTLSBulkCertificates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSBulkCertificatesWithDefaults() *RelationshipTLSBulkCertificates {
	this := RelationshipTLSBulkCertificates{}
	return &this
}

// GetTLSBulkCertificates returns the TLSBulkCertificates field value if set, zero value otherwise.
func (o *RelationshipTLSBulkCertificates) GetTLSBulkCertificates() RelationshipTLSBulkCertificateTLSBulkCertificate {
	if o == nil || o.TLSBulkCertificates == nil {
		var ret RelationshipTLSBulkCertificateTLSBulkCertificate
		return ret
	}
	return *o.TLSBulkCertificates
}

// GetTLSBulkCertificatesOk returns a tuple with the TLSBulkCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSBulkCertificates) GetTLSBulkCertificatesOk() (*RelationshipTLSBulkCertificateTLSBulkCertificate, bool) {
	if o == nil || o.TLSBulkCertificates == nil {
		return nil, false
	}
	return o.TLSBulkCertificates, true
}

// HasTLSBulkCertificates returns a boolean if a field has been set.
func (o *RelationshipTLSBulkCertificates) HasTLSBulkCertificates() bool {
	if o != nil && o.TLSBulkCertificates != nil {
		return true
	}

	return false
}

// SetTLSBulkCertificates gets a reference to the given RelationshipTLSBulkCertificateTLSBulkCertificate and assigns it to the TLSBulkCertificates field.
func (o *RelationshipTLSBulkCertificates) SetTLSBulkCertificates(v RelationshipTLSBulkCertificateTLSBulkCertificate) {
	o.TLSBulkCertificates = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSBulkCertificates) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSBulkCertificates != nil {
		toSerialize["tls_bulk_certificates"] = o.TLSBulkCertificates
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSBulkCertificates) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSBulkCertificates := _RelationshipTLSBulkCertificates{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSBulkCertificates); err == nil {
		*o = RelationshipTLSBulkCertificates(varRelationshipTLSBulkCertificates)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_bulk_certificates")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSBulkCertificates is a helper abstraction for handling nullable relationshiptlsbulkcertificates types. 
type NullableRelationshipTLSBulkCertificates struct {
	value *RelationshipTLSBulkCertificates
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSBulkCertificates) Get() *RelationshipTLSBulkCertificates {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSBulkCertificates) Set(val *RelationshipTLSBulkCertificates) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSBulkCertificates) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSBulkCertificates) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSBulkCertificates returns a pointer to a new instance of NullableRelationshipTLSBulkCertificates.
func NewNullableRelationshipTLSBulkCertificates(val *RelationshipTLSBulkCertificates) *NullableRelationshipTLSBulkCertificates {
	return &NullableRelationshipTLSBulkCertificates{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSBulkCertificates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSBulkCertificates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
