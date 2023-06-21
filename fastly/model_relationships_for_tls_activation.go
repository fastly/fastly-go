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

// RelationshipsForTLSActivation struct for RelationshipsForTLSActivation
type RelationshipsForTLSActivation struct {
	TLSCertificate *RelationshipTLSCertificateTLSCertificate `json:"tls_certificate,omitempty"`
	TLSConfiguration *RelationshipTLSConfigurationTLSConfiguration `json:"tls_configuration,omitempty"`
	TLSDomain *RelationshipTLSDomainTLSDomain `json:"tls_domain,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipsForTLSActivation RelationshipsForTLSActivation

// NewRelationshipsForTLSActivation instantiates a new RelationshipsForTLSActivation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipsForTLSActivation() *RelationshipsForTLSActivation {
	this := RelationshipsForTLSActivation{}
	return &this
}

// NewRelationshipsForTLSActivationWithDefaults instantiates a new RelationshipsForTLSActivation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipsForTLSActivationWithDefaults() *RelationshipsForTLSActivation {
	this := RelationshipsForTLSActivation{}
	return &this
}

// GetTLSCertificate returns the TLSCertificate field value if set, zero value otherwise.
func (o *RelationshipsForTLSActivation) GetTLSCertificate() RelationshipTLSCertificateTLSCertificate {
	if o == nil || o.TLSCertificate == nil {
		var ret RelationshipTLSCertificateTLSCertificate
		return ret
	}
	return *o.TLSCertificate
}

// GetTLSCertificateOk returns a tuple with the TLSCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipsForTLSActivation) GetTLSCertificateOk() (*RelationshipTLSCertificateTLSCertificate, bool) {
	if o == nil || o.TLSCertificate == nil {
		return nil, false
	}
	return o.TLSCertificate, true
}

// HasTLSCertificate returns a boolean if a field has been set.
func (o *RelationshipsForTLSActivation) HasTLSCertificate() bool {
	if o != nil && o.TLSCertificate != nil {
		return true
	}

	return false
}

// SetTLSCertificate gets a reference to the given RelationshipTLSCertificateTLSCertificate and assigns it to the TLSCertificate field.
func (o *RelationshipsForTLSActivation) SetTLSCertificate(v RelationshipTLSCertificateTLSCertificate) {
	o.TLSCertificate = &v
}

// GetTLSConfiguration returns the TLSConfiguration field value if set, zero value otherwise.
func (o *RelationshipsForTLSActivation) GetTLSConfiguration() RelationshipTLSConfigurationTLSConfiguration {
	if o == nil || o.TLSConfiguration == nil {
		var ret RelationshipTLSConfigurationTLSConfiguration
		return ret
	}
	return *o.TLSConfiguration
}

// GetTLSConfigurationOk returns a tuple with the TLSConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipsForTLSActivation) GetTLSConfigurationOk() (*RelationshipTLSConfigurationTLSConfiguration, bool) {
	if o == nil || o.TLSConfiguration == nil {
		return nil, false
	}
	return o.TLSConfiguration, true
}

// HasTLSConfiguration returns a boolean if a field has been set.
func (o *RelationshipsForTLSActivation) HasTLSConfiguration() bool {
	if o != nil && o.TLSConfiguration != nil {
		return true
	}

	return false
}

// SetTLSConfiguration gets a reference to the given RelationshipTLSConfigurationTLSConfiguration and assigns it to the TLSConfiguration field.
func (o *RelationshipsForTLSActivation) SetTLSConfiguration(v RelationshipTLSConfigurationTLSConfiguration) {
	o.TLSConfiguration = &v
}

// GetTLSDomain returns the TLSDomain field value if set, zero value otherwise.
func (o *RelationshipsForTLSActivation) GetTLSDomain() RelationshipTLSDomainTLSDomain {
	if o == nil || o.TLSDomain == nil {
		var ret RelationshipTLSDomainTLSDomain
		return ret
	}
	return *o.TLSDomain
}

// GetTLSDomainOk returns a tuple with the TLSDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipsForTLSActivation) GetTLSDomainOk() (*RelationshipTLSDomainTLSDomain, bool) {
	if o == nil || o.TLSDomain == nil {
		return nil, false
	}
	return o.TLSDomain, true
}

// HasTLSDomain returns a boolean if a field has been set.
func (o *RelationshipsForTLSActivation) HasTLSDomain() bool {
	if o != nil && o.TLSDomain != nil {
		return true
	}

	return false
}

// SetTLSDomain gets a reference to the given RelationshipTLSDomainTLSDomain and assigns it to the TLSDomain field.
func (o *RelationshipsForTLSActivation) SetTLSDomain(v RelationshipTLSDomainTLSDomain) {
	o.TLSDomain = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipsForTLSActivation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSCertificate != nil {
		toSerialize["tls_certificate"] = o.TLSCertificate
	}
	if o.TLSConfiguration != nil {
		toSerialize["tls_configuration"] = o.TLSConfiguration
	}
	if o.TLSDomain != nil {
		toSerialize["tls_domain"] = o.TLSDomain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipsForTLSActivation) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipsForTLSActivation := _RelationshipsForTLSActivation{}

	if err = json.Unmarshal(bytes, &varRelationshipsForTLSActivation); err == nil {
		*o = RelationshipsForTLSActivation(varRelationshipsForTLSActivation)
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

// NullableRelationshipsForTLSActivation is a helper abstraction for handling nullable relationshipsfortlsactivation types. 
type NullableRelationshipsForTLSActivation struct {
	value *RelationshipsForTLSActivation
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForTLSActivation) Get() *RelationshipsForTLSActivation {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForTLSActivation) Set(val *RelationshipsForTLSActivation) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForTLSActivation) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForTLSActivation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForTLSActivation returns a pointer to a new instance of NullableRelationshipsForTLSActivation.
func NewNullableRelationshipsForTLSActivation(val *RelationshipsForTLSActivation) *NullableRelationshipsForTLSActivation {
	return &NullableRelationshipsForTLSActivation{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForTLSActivation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipsForTLSActivation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
