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

// RelationshipTlsCertificate The [TLS certificate](https://www.fastly.com/documentation/reference/api/tls/custom-certs/certificates/) being used to terminate TLS traffic for a domain. Required.
type RelationshipTlsCertificate struct {
	TlsCertificate       *RelationshipTlsCertificateTlsCertificate `json:"tls_certificate,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsCertificate RelationshipTlsCertificate

// NewRelationshipTlsCertificate instantiates a new RelationshipTlsCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsCertificate() *RelationshipTlsCertificate {
	this := RelationshipTlsCertificate{}
	return &this
}

// NewRelationshipTlsCertificateWithDefaults instantiates a new RelationshipTlsCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsCertificateWithDefaults() *RelationshipTlsCertificate {
	this := RelationshipTlsCertificate{}
	return &this
}

// GetTlsCertificate returns the TlsCertificate field value if set, zero value otherwise.
func (o *RelationshipTlsCertificate) GetTlsCertificate() RelationshipTlsCertificateTlsCertificate {
	if o == nil || o.TlsCertificate == nil {
		var ret RelationshipTlsCertificateTlsCertificate
		return ret
	}
	return *o.TlsCertificate
}

// GetTlsCertificateOk returns a tuple with the TlsCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsCertificate) GetTlsCertificateOk() (*RelationshipTlsCertificateTlsCertificate, bool) {
	if o == nil || o.TlsCertificate == nil {
		return nil, false
	}
	return o.TlsCertificate, true
}

// HasTlsCertificate returns a boolean if a field has been set.
func (o *RelationshipTlsCertificate) HasTlsCertificate() bool {
	if o != nil && o.TlsCertificate != nil {
		return true
	}

	return false
}

// SetTlsCertificate gets a reference to the given RelationshipTlsCertificateTlsCertificate and assigns it to the TlsCertificate field.
func (o *RelationshipTlsCertificate) SetTlsCertificate(v RelationshipTlsCertificateTlsCertificate) {
	o.TlsCertificate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsCertificate != nil {
		toSerialize["tls_certificate"] = o.TlsCertificate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsCertificate := _RelationshipTlsCertificate{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsCertificate); err == nil {
		*o = RelationshipTlsCertificate(varRelationshipTlsCertificate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_certificate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsCertificate is a helper abstraction for handling nullable relationshiptlscertificate types.
type NullableRelationshipTlsCertificate struct {
	value *RelationshipTlsCertificate
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsCertificate) Get() *RelationshipTlsCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsCertificate) Set(val *RelationshipTlsCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsCertificate returns a pointer to a new instance of NullableRelationshipTlsCertificate.
func NewNullableRelationshipTlsCertificate(val *RelationshipTlsCertificate) *NullableRelationshipTlsCertificate {
	return &NullableRelationshipTlsCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
