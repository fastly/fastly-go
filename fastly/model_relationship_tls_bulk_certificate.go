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

// RelationshipTLSBulkCertificate struct for RelationshipTLSBulkCertificate
type RelationshipTLSBulkCertificate struct {
	TLSBulkCertificate *RelationshipTLSBulkCertificateTLSBulkCertificate `json:"tls_bulk_certificate,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSBulkCertificate RelationshipTLSBulkCertificate

// NewRelationshipTLSBulkCertificate instantiates a new RelationshipTLSBulkCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSBulkCertificate() *RelationshipTLSBulkCertificate {
	this := RelationshipTLSBulkCertificate{}
	return &this
}

// NewRelationshipTLSBulkCertificateWithDefaults instantiates a new RelationshipTLSBulkCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSBulkCertificateWithDefaults() *RelationshipTLSBulkCertificate {
	this := RelationshipTLSBulkCertificate{}
	return &this
}

// GetTLSBulkCertificate returns the TLSBulkCertificate field value if set, zero value otherwise.
func (o *RelationshipTLSBulkCertificate) GetTLSBulkCertificate() RelationshipTLSBulkCertificateTLSBulkCertificate {
	if o == nil || o.TLSBulkCertificate == nil {
		var ret RelationshipTLSBulkCertificateTLSBulkCertificate
		return ret
	}
	return *o.TLSBulkCertificate
}

// GetTLSBulkCertificateOk returns a tuple with the TLSBulkCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSBulkCertificate) GetTLSBulkCertificateOk() (*RelationshipTLSBulkCertificateTLSBulkCertificate, bool) {
	if o == nil || o.TLSBulkCertificate == nil {
		return nil, false
	}
	return o.TLSBulkCertificate, true
}

// HasTLSBulkCertificate returns a boolean if a field has been set.
func (o *RelationshipTLSBulkCertificate) HasTLSBulkCertificate() bool {
	if o != nil && o.TLSBulkCertificate != nil {
		return true
	}

	return false
}

// SetTLSBulkCertificate gets a reference to the given RelationshipTLSBulkCertificateTLSBulkCertificate and assigns it to the TLSBulkCertificate field.
func (o *RelationshipTLSBulkCertificate) SetTLSBulkCertificate(v RelationshipTLSBulkCertificateTLSBulkCertificate) {
	o.TLSBulkCertificate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSBulkCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSBulkCertificate != nil {
		toSerialize["tls_bulk_certificate"] = o.TLSBulkCertificate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSBulkCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSBulkCertificate := _RelationshipTLSBulkCertificate{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSBulkCertificate); err == nil {
		*o = RelationshipTLSBulkCertificate(varRelationshipTLSBulkCertificate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_bulk_certificate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSBulkCertificate is a helper abstraction for handling nullable relationshiptlsbulkcertificate types. 
type NullableRelationshipTLSBulkCertificate struct {
	value *RelationshipTLSBulkCertificate
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSBulkCertificate) Get() *RelationshipTLSBulkCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSBulkCertificate) Set(val *RelationshipTLSBulkCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSBulkCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSBulkCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSBulkCertificate returns a pointer to a new instance of NullableRelationshipTLSBulkCertificate.
func NewNullableRelationshipTLSBulkCertificate(val *RelationshipTLSBulkCertificate) *NullableRelationshipTLSBulkCertificate {
	return &NullableRelationshipTLSBulkCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSBulkCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSBulkCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
