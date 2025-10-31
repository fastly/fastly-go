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

// RelationshipTlsBulkCertificate struct for RelationshipTlsBulkCertificate
type RelationshipTlsBulkCertificate struct {
	TlsBulkCertificate   *RelationshipTlsBulkCertificateTlsBulkCertificate `json:"tls_bulk_certificate,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsBulkCertificate RelationshipTlsBulkCertificate

// NewRelationshipTlsBulkCertificate instantiates a new RelationshipTlsBulkCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsBulkCertificate() *RelationshipTlsBulkCertificate {
	this := RelationshipTlsBulkCertificate{}
	return &this
}

// NewRelationshipTlsBulkCertificateWithDefaults instantiates a new RelationshipTlsBulkCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsBulkCertificateWithDefaults() *RelationshipTlsBulkCertificate {
	this := RelationshipTlsBulkCertificate{}
	return &this
}

// GetTlsBulkCertificate returns the TlsBulkCertificate field value if set, zero value otherwise.
func (o *RelationshipTlsBulkCertificate) GetTlsBulkCertificate() RelationshipTlsBulkCertificateTlsBulkCertificate {
	if o == nil || o.TlsBulkCertificate == nil {
		var ret RelationshipTlsBulkCertificateTlsBulkCertificate
		return ret
	}
	return *o.TlsBulkCertificate
}

// GetTlsBulkCertificateOk returns a tuple with the TlsBulkCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsBulkCertificate) GetTlsBulkCertificateOk() (*RelationshipTlsBulkCertificateTlsBulkCertificate, bool) {
	if o == nil || o.TlsBulkCertificate == nil {
		return nil, false
	}
	return o.TlsBulkCertificate, true
}

// HasTlsBulkCertificate returns a boolean if a field has been set.
func (o *RelationshipTlsBulkCertificate) HasTlsBulkCertificate() bool {
	if o != nil && o.TlsBulkCertificate != nil {
		return true
	}

	return false
}

// SetTlsBulkCertificate gets a reference to the given RelationshipTlsBulkCertificateTlsBulkCertificate and assigns it to the TlsBulkCertificate field.
func (o *RelationshipTlsBulkCertificate) SetTlsBulkCertificate(v RelationshipTlsBulkCertificateTlsBulkCertificate) {
	o.TlsBulkCertificate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsBulkCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsBulkCertificate != nil {
		toSerialize["tls_bulk_certificate"] = o.TlsBulkCertificate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsBulkCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsBulkCertificate := _RelationshipTlsBulkCertificate{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsBulkCertificate); err == nil {
		*o = RelationshipTlsBulkCertificate(varRelationshipTlsBulkCertificate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_bulk_certificate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsBulkCertificate is a helper abstraction for handling nullable relationshiptlsbulkcertificate types.
type NullableRelationshipTlsBulkCertificate struct {
	value *RelationshipTlsBulkCertificate
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsBulkCertificate) Get() *RelationshipTlsBulkCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsBulkCertificate) Set(val *RelationshipTlsBulkCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsBulkCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsBulkCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsBulkCertificate returns a pointer to a new instance of NullableRelationshipTlsBulkCertificate.
func NewNullableRelationshipTlsBulkCertificate(val *RelationshipTlsBulkCertificate) *NullableRelationshipTlsBulkCertificate {
	return &NullableRelationshipTlsBulkCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsBulkCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsBulkCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
