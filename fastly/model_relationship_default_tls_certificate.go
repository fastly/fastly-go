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

// RelationshipDefaultTlsCertificate struct for RelationshipDefaultTlsCertificate
type RelationshipDefaultTlsCertificate struct {
	DefaultCertificate   *RelationshipDefaultTlsCertificateDefaultCertificate `json:"default_certificate,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipDefaultTlsCertificate RelationshipDefaultTlsCertificate

// NewRelationshipDefaultTlsCertificate instantiates a new RelationshipDefaultTlsCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipDefaultTlsCertificate() *RelationshipDefaultTlsCertificate {
	this := RelationshipDefaultTlsCertificate{}
	return &this
}

// NewRelationshipDefaultTlsCertificateWithDefaults instantiates a new RelationshipDefaultTlsCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipDefaultTlsCertificateWithDefaults() *RelationshipDefaultTlsCertificate {
	this := RelationshipDefaultTlsCertificate{}
	return &this
}

// GetDefaultCertificate returns the DefaultCertificate field value if set, zero value otherwise.
func (o *RelationshipDefaultTlsCertificate) GetDefaultCertificate() RelationshipDefaultTlsCertificateDefaultCertificate {
	if o == nil || o.DefaultCertificate == nil {
		var ret RelationshipDefaultTlsCertificateDefaultCertificate
		return ret
	}
	return *o.DefaultCertificate
}

// GetDefaultCertificateOk returns a tuple with the DefaultCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipDefaultTlsCertificate) GetDefaultCertificateOk() (*RelationshipDefaultTlsCertificateDefaultCertificate, bool) {
	if o == nil || o.DefaultCertificate == nil {
		return nil, false
	}
	return o.DefaultCertificate, true
}

// HasDefaultCertificate returns a boolean if a field has been set.
func (o *RelationshipDefaultTlsCertificate) HasDefaultCertificate() bool {
	if o != nil && o.DefaultCertificate != nil {
		return true
	}

	return false
}

// SetDefaultCertificate gets a reference to the given RelationshipDefaultTlsCertificateDefaultCertificate and assigns it to the DefaultCertificate field.
func (o *RelationshipDefaultTlsCertificate) SetDefaultCertificate(v RelationshipDefaultTlsCertificateDefaultCertificate) {
	o.DefaultCertificate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipDefaultTlsCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.DefaultCertificate != nil {
		toSerialize["default_certificate"] = o.DefaultCertificate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipDefaultTlsCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipDefaultTlsCertificate := _RelationshipDefaultTlsCertificate{}

	if err = json.Unmarshal(bytes, &varRelationshipDefaultTlsCertificate); err == nil {
		*o = RelationshipDefaultTlsCertificate(varRelationshipDefaultTlsCertificate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "default_certificate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipDefaultTlsCertificate is a helper abstraction for handling nullable relationshipdefaulttlscertificate types.
type NullableRelationshipDefaultTlsCertificate struct {
	value *RelationshipDefaultTlsCertificate
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipDefaultTlsCertificate) Get() *RelationshipDefaultTlsCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipDefaultTlsCertificate) Set(val *RelationshipDefaultTlsCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipDefaultTlsCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipDefaultTlsCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipDefaultTlsCertificate returns a pointer to a new instance of NullableRelationshipDefaultTlsCertificate.
func NewNullableRelationshipDefaultTlsCertificate(val *RelationshipDefaultTlsCertificate) *NullableRelationshipDefaultTlsCertificate {
	return &NullableRelationshipDefaultTlsCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipDefaultTlsCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipDefaultTlsCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
