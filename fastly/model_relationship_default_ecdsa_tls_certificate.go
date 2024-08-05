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

// RelationshipDefaultEcdsaTLSCertificate struct for RelationshipDefaultEcdsaTLSCertificate
type RelationshipDefaultEcdsaTLSCertificate struct {
	DefaultEcdsaCertificate *RelationshipDefaultEcdsaTLSCertificateDefaultEcdsaCertificate `json:"default_ecdsa_certificate,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipDefaultEcdsaTLSCertificate RelationshipDefaultEcdsaTLSCertificate

// NewRelationshipDefaultEcdsaTLSCertificate instantiates a new RelationshipDefaultEcdsaTLSCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipDefaultEcdsaTLSCertificate() *RelationshipDefaultEcdsaTLSCertificate {
	this := RelationshipDefaultEcdsaTLSCertificate{}
	return &this
}

// NewRelationshipDefaultEcdsaTLSCertificateWithDefaults instantiates a new RelationshipDefaultEcdsaTLSCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipDefaultEcdsaTLSCertificateWithDefaults() *RelationshipDefaultEcdsaTLSCertificate {
	this := RelationshipDefaultEcdsaTLSCertificate{}
	return &this
}

// GetDefaultEcdsaCertificate returns the DefaultEcdsaCertificate field value if set, zero value otherwise.
func (o *RelationshipDefaultEcdsaTLSCertificate) GetDefaultEcdsaCertificate() RelationshipDefaultEcdsaTLSCertificateDefaultEcdsaCertificate {
	if o == nil || o.DefaultEcdsaCertificate == nil {
		var ret RelationshipDefaultEcdsaTLSCertificateDefaultEcdsaCertificate
		return ret
	}
	return *o.DefaultEcdsaCertificate
}

// GetDefaultEcdsaCertificateOk returns a tuple with the DefaultEcdsaCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipDefaultEcdsaTLSCertificate) GetDefaultEcdsaCertificateOk() (*RelationshipDefaultEcdsaTLSCertificateDefaultEcdsaCertificate, bool) {
	if o == nil || o.DefaultEcdsaCertificate == nil {
		return nil, false
	}
	return o.DefaultEcdsaCertificate, true
}

// HasDefaultEcdsaCertificate returns a boolean if a field has been set.
func (o *RelationshipDefaultEcdsaTLSCertificate) HasDefaultEcdsaCertificate() bool {
	if o != nil && o.DefaultEcdsaCertificate != nil {
		return true
	}

	return false
}

// SetDefaultEcdsaCertificate gets a reference to the given RelationshipDefaultEcdsaTLSCertificateDefaultEcdsaCertificate and assigns it to the DefaultEcdsaCertificate field.
func (o *RelationshipDefaultEcdsaTLSCertificate) SetDefaultEcdsaCertificate(v RelationshipDefaultEcdsaTLSCertificateDefaultEcdsaCertificate) {
	o.DefaultEcdsaCertificate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipDefaultEcdsaTLSCertificate) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipDefaultEcdsaTLSCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipDefaultEcdsaTLSCertificate := _RelationshipDefaultEcdsaTLSCertificate{}

	if err = json.Unmarshal(bytes, &varRelationshipDefaultEcdsaTLSCertificate); err == nil {
		*o = RelationshipDefaultEcdsaTLSCertificate(varRelationshipDefaultEcdsaTLSCertificate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "default_ecdsa_certificate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipDefaultEcdsaTLSCertificate is a helper abstraction for handling nullable relationshipdefaultecdsatlscertificate types. 
type NullableRelationshipDefaultEcdsaTLSCertificate struct {
	value *RelationshipDefaultEcdsaTLSCertificate
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipDefaultEcdsaTLSCertificate) Get() *RelationshipDefaultEcdsaTLSCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipDefaultEcdsaTLSCertificate) Set(val *RelationshipDefaultEcdsaTLSCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipDefaultEcdsaTLSCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipDefaultEcdsaTLSCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipDefaultEcdsaTLSCertificate returns a pointer to a new instance of NullableRelationshipDefaultEcdsaTLSCertificate.
func NewNullableRelationshipDefaultEcdsaTLSCertificate(val *RelationshipDefaultEcdsaTLSCertificate) *NullableRelationshipDefaultEcdsaTLSCertificate {
	return &NullableRelationshipDefaultEcdsaTLSCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipDefaultEcdsaTLSCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipDefaultEcdsaTLSCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
