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

// RelationshipsForTlsActivation struct for RelationshipsForTlsActivation
type RelationshipsForTlsActivation struct {
	TlsCertificate       *RelationshipTlsCertificateTlsCertificate     `json:"tls_certificate,omitempty"`
	TlsConfiguration     *RelationshipTlsConfigurationTlsConfiguration `json:"tls_configuration,omitempty"`
	TlsDomain            *RelationshipTlsDomainTlsDomain               `json:"tls_domain,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipsForTlsActivation RelationshipsForTlsActivation

// NewRelationshipsForTlsActivation instantiates a new RelationshipsForTlsActivation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipsForTlsActivation() *RelationshipsForTlsActivation {
	this := RelationshipsForTlsActivation{}
	return &this
}

// NewRelationshipsForTlsActivationWithDefaults instantiates a new RelationshipsForTlsActivation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipsForTlsActivationWithDefaults() *RelationshipsForTlsActivation {
	this := RelationshipsForTlsActivation{}
	return &this
}

// GetTlsCertificate returns the TlsCertificate field value if set, zero value otherwise.
func (o *RelationshipsForTlsActivation) GetTlsCertificate() RelationshipTlsCertificateTlsCertificate {
	if o == nil || o.TlsCertificate == nil {
		var ret RelationshipTlsCertificateTlsCertificate
		return ret
	}
	return *o.TlsCertificate
}

// GetTlsCertificateOk returns a tuple with the TlsCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipsForTlsActivation) GetTlsCertificateOk() (*RelationshipTlsCertificateTlsCertificate, bool) {
	if o == nil || o.TlsCertificate == nil {
		return nil, false
	}
	return o.TlsCertificate, true
}

// HasTlsCertificate returns a boolean if a field has been set.
func (o *RelationshipsForTlsActivation) HasTlsCertificate() bool {
	if o != nil && o.TlsCertificate != nil {
		return true
	}

	return false
}

// SetTlsCertificate gets a reference to the given RelationshipTlsCertificateTlsCertificate and assigns it to the TlsCertificate field.
func (o *RelationshipsForTlsActivation) SetTlsCertificate(v RelationshipTlsCertificateTlsCertificate) {
	o.TlsCertificate = &v
}

// GetTlsConfiguration returns the TlsConfiguration field value if set, zero value otherwise.
func (o *RelationshipsForTlsActivation) GetTlsConfiguration() RelationshipTlsConfigurationTlsConfiguration {
	if o == nil || o.TlsConfiguration == nil {
		var ret RelationshipTlsConfigurationTlsConfiguration
		return ret
	}
	return *o.TlsConfiguration
}

// GetTlsConfigurationOk returns a tuple with the TlsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipsForTlsActivation) GetTlsConfigurationOk() (*RelationshipTlsConfigurationTlsConfiguration, bool) {
	if o == nil || o.TlsConfiguration == nil {
		return nil, false
	}
	return o.TlsConfiguration, true
}

// HasTlsConfiguration returns a boolean if a field has been set.
func (o *RelationshipsForTlsActivation) HasTlsConfiguration() bool {
	if o != nil && o.TlsConfiguration != nil {
		return true
	}

	return false
}

// SetTlsConfiguration gets a reference to the given RelationshipTlsConfigurationTlsConfiguration and assigns it to the TlsConfiguration field.
func (o *RelationshipsForTlsActivation) SetTlsConfiguration(v RelationshipTlsConfigurationTlsConfiguration) {
	o.TlsConfiguration = &v
}

// GetTlsDomain returns the TlsDomain field value if set, zero value otherwise.
func (o *RelationshipsForTlsActivation) GetTlsDomain() RelationshipTlsDomainTlsDomain {
	if o == nil || o.TlsDomain == nil {
		var ret RelationshipTlsDomainTlsDomain
		return ret
	}
	return *o.TlsDomain
}

// GetTlsDomainOk returns a tuple with the TlsDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipsForTlsActivation) GetTlsDomainOk() (*RelationshipTlsDomainTlsDomain, bool) {
	if o == nil || o.TlsDomain == nil {
		return nil, false
	}
	return o.TlsDomain, true
}

// HasTlsDomain returns a boolean if a field has been set.
func (o *RelationshipsForTlsActivation) HasTlsDomain() bool {
	if o != nil && o.TlsDomain != nil {
		return true
	}

	return false
}

// SetTlsDomain gets a reference to the given RelationshipTlsDomainTlsDomain and assigns it to the TlsDomain field.
func (o *RelationshipsForTlsActivation) SetTlsDomain(v RelationshipTlsDomainTlsDomain) {
	o.TlsDomain = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipsForTlsActivation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsCertificate != nil {
		toSerialize["tls_certificate"] = o.TlsCertificate
	}
	if o.TlsConfiguration != nil {
		toSerialize["tls_configuration"] = o.TlsConfiguration
	}
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
func (o *RelationshipsForTlsActivation) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipsForTlsActivation := _RelationshipsForTlsActivation{}

	if err = json.Unmarshal(bytes, &varRelationshipsForTlsActivation); err == nil {
		*o = RelationshipsForTlsActivation(varRelationshipsForTlsActivation)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_certificate")
		delete(additionalProperties, "tls_configuration")
		delete(additionalProperties, "tls_domain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipsForTlsActivation is a helper abstraction for handling nullable relationshipsfortlsactivation types.
type NullableRelationshipsForTlsActivation struct {
	value *RelationshipsForTlsActivation
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForTlsActivation) Get() *RelationshipsForTlsActivation {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForTlsActivation) Set(val *RelationshipsForTlsActivation) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForTlsActivation) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForTlsActivation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForTlsActivation returns a pointer to a new instance of NullableRelationshipsForTlsActivation.
func NewNullableRelationshipsForTlsActivation(val *RelationshipsForTlsActivation) *NullableRelationshipsForTlsActivation {
	return &NullableRelationshipsForTlsActivation{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForTlsActivation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipsForTlsActivation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
