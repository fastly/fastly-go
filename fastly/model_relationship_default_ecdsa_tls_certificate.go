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

// RelationshipDefaultEcdsaTlsCertificate struct for RelationshipDefaultEcdsaTlsCertificate
type RelationshipDefaultEcdsaTlsCertificate struct {
	DefaultEcdsaCertificate *RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate `json:"default_ecdsa_certificate,omitempty"`
	AdditionalProperties    map[string]any
}

type _RelationshipDefaultEcdsaTlsCertificate RelationshipDefaultEcdsaTlsCertificate

// NewRelationshipDefaultEcdsaTlsCertificate instantiates a new RelationshipDefaultEcdsaTlsCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipDefaultEcdsaTlsCertificate() *RelationshipDefaultEcdsaTlsCertificate {
	this := RelationshipDefaultEcdsaTlsCertificate{}
	return &this
}

// NewRelationshipDefaultEcdsaTlsCertificateWithDefaults instantiates a new RelationshipDefaultEcdsaTlsCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipDefaultEcdsaTlsCertificateWithDefaults() *RelationshipDefaultEcdsaTlsCertificate {
	this := RelationshipDefaultEcdsaTlsCertificate{}
	return &this
}

// GetDefaultEcdsaCertificate returns the DefaultEcdsaCertificate field value if set, zero value otherwise.
func (o *RelationshipDefaultEcdsaTlsCertificate) GetDefaultEcdsaCertificate() RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate {
	if o == nil || o.DefaultEcdsaCertificate == nil {
		var ret RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate
		return ret
	}
	return *o.DefaultEcdsaCertificate
}

// GetDefaultEcdsaCertificateOk returns a tuple with the DefaultEcdsaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipDefaultEcdsaTlsCertificate) GetDefaultEcdsaCertificateOk() (*RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate, bool) {
	if o == nil || o.DefaultEcdsaCertificate == nil {
		return nil, false
	}
	return o.DefaultEcdsaCertificate, true
}

// HasDefaultEcdsaCertificate returns a boolean if a field has been set.
func (o *RelationshipDefaultEcdsaTlsCertificate) HasDefaultEcdsaCertificate() bool {
	if o != nil && o.DefaultEcdsaCertificate != nil {
		return true
	}

	return false
}

// SetDefaultEcdsaCertificate gets a reference to the given RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate and assigns it to the DefaultEcdsaCertificate field.
func (o *RelationshipDefaultEcdsaTlsCertificate) SetDefaultEcdsaCertificate(v RelationshipDefaultEcdsaTlsCertificateDefaultEcdsaCertificate) {
	o.DefaultEcdsaCertificate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipDefaultEcdsaTlsCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.DefaultEcdsaCertificate != nil {
		toSerialize["default_ecdsa_certificate"] = o.DefaultEcdsaCertificate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipDefaultEcdsaTlsCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipDefaultEcdsaTlsCertificate := _RelationshipDefaultEcdsaTlsCertificate{}

	if err = json.Unmarshal(bytes, &varRelationshipDefaultEcdsaTlsCertificate); err == nil {
		*o = RelationshipDefaultEcdsaTlsCertificate(varRelationshipDefaultEcdsaTlsCertificate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "default_ecdsa_certificate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipDefaultEcdsaTlsCertificate is a helper abstraction for handling nullable relationshipdefaultecdsatlscertificate types.
type NullableRelationshipDefaultEcdsaTlsCertificate struct {
	value *RelationshipDefaultEcdsaTlsCertificate
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipDefaultEcdsaTlsCertificate) Get() *RelationshipDefaultEcdsaTlsCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipDefaultEcdsaTlsCertificate) Set(val *RelationshipDefaultEcdsaTlsCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipDefaultEcdsaTlsCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipDefaultEcdsaTlsCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipDefaultEcdsaTlsCertificate returns a pointer to a new instance of NullableRelationshipDefaultEcdsaTlsCertificate.
func NewNullableRelationshipDefaultEcdsaTlsCertificate(val *RelationshipDefaultEcdsaTlsCertificate) *NullableRelationshipDefaultEcdsaTlsCertificate {
	return &NullableRelationshipDefaultEcdsaTlsCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipDefaultEcdsaTlsCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipDefaultEcdsaTlsCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
