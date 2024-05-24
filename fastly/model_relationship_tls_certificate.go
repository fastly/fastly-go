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

// RelationshipTLSCertificate The [TLS certificate](https://www.fastly.com/documentation/reference/api/tls/custom-certs/certificates/) being used to terminate TLS traffic for a domain. Required.
type RelationshipTLSCertificate struct {
	TLSCertificate *RelationshipTLSCertificateTLSCertificate `json:"tls_certificate,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSCertificate RelationshipTLSCertificate

// NewRelationshipTLSCertificate instantiates a new RelationshipTLSCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSCertificate() *RelationshipTLSCertificate {
	this := RelationshipTLSCertificate{}
	return &this
}

// NewRelationshipTLSCertificateWithDefaults instantiates a new RelationshipTLSCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSCertificateWithDefaults() *RelationshipTLSCertificate {
	this := RelationshipTLSCertificate{}
	return &this
}

// GetTLSCertificate returns the TLSCertificate field value if set, zero value otherwise.
func (o *RelationshipTLSCertificate) GetTLSCertificate() RelationshipTLSCertificateTLSCertificate {
	if o == nil || o.TLSCertificate == nil {
		var ret RelationshipTLSCertificateTLSCertificate
		return ret
	}
	return *o.TLSCertificate
}

// GetTLSCertificateOk returns a tuple with the TLSCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSCertificate) GetTLSCertificateOk() (*RelationshipTLSCertificateTLSCertificate, bool) {
	if o == nil || o.TLSCertificate == nil {
		return nil, false
	}
	return o.TLSCertificate, true
}

// HasTLSCertificate returns a boolean if a field has been set.
func (o *RelationshipTLSCertificate) HasTLSCertificate() bool {
	if o != nil && o.TLSCertificate != nil {
		return true
	}

	return false
}

// SetTLSCertificate gets a reference to the given RelationshipTLSCertificateTLSCertificate and assigns it to the TLSCertificate field.
func (o *RelationshipTLSCertificate) SetTLSCertificate(v RelationshipTLSCertificateTLSCertificate) {
	o.TLSCertificate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSCertificate != nil {
		toSerialize["tls_certificate"] = o.TLSCertificate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSCertificate := _RelationshipTLSCertificate{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSCertificate); err == nil {
		*o = RelationshipTLSCertificate(varRelationshipTLSCertificate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_certificate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSCertificate is a helper abstraction for handling nullable relationshiptlscertificate types. 
type NullableRelationshipTLSCertificate struct {
	value *RelationshipTLSCertificate
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSCertificate) Get() *RelationshipTLSCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSCertificate) Set(val *RelationshipTLSCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSCertificate returns a pointer to a new instance of NullableRelationshipTLSCertificate.
func NewNullableRelationshipTLSCertificate(val *RelationshipTLSCertificate) *NullableRelationshipTLSCertificate {
	return &NullableRelationshipTLSCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
