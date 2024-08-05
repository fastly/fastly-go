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

// RelationshipDefaultTLSCertificate struct for RelationshipDefaultTLSCertificate
type RelationshipDefaultTLSCertificate struct {
	DefaultCertificate *RelationshipDefaultTLSCertificateDefaultCertificate `json:"default_certificate,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipDefaultTLSCertificate RelationshipDefaultTLSCertificate

// NewRelationshipDefaultTLSCertificate instantiates a new RelationshipDefaultTLSCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipDefaultTLSCertificate() *RelationshipDefaultTLSCertificate {
	this := RelationshipDefaultTLSCertificate{}
	return &this
}

// NewRelationshipDefaultTLSCertificateWithDefaults instantiates a new RelationshipDefaultTLSCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipDefaultTLSCertificateWithDefaults() *RelationshipDefaultTLSCertificate {
	this := RelationshipDefaultTLSCertificate{}
	return &this
}

// GetDefaultCertificate returns the DefaultCertificate field value if set, zero value otherwise.
func (o *RelationshipDefaultTLSCertificate) GetDefaultCertificate() RelationshipDefaultTLSCertificateDefaultCertificate {
	if o == nil || o.DefaultCertificate == nil {
		var ret RelationshipDefaultTLSCertificateDefaultCertificate
		return ret
	}
	return *o.DefaultCertificate
}

// GetDefaultCertificateOk returns a tuple with the DefaultCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipDefaultTLSCertificate) GetDefaultCertificateOk() (*RelationshipDefaultTLSCertificateDefaultCertificate, bool) {
	if o == nil || o.DefaultCertificate == nil {
		return nil, false
	}
	return o.DefaultCertificate, true
}

// HasDefaultCertificate returns a boolean if a field has been set.
func (o *RelationshipDefaultTLSCertificate) HasDefaultCertificate() bool {
	if o != nil && o.DefaultCertificate != nil {
		return true
	}

	return false
}

// SetDefaultCertificate gets a reference to the given RelationshipDefaultTLSCertificateDefaultCertificate and assigns it to the DefaultCertificate field.
func (o *RelationshipDefaultTLSCertificate) SetDefaultCertificate(v RelationshipDefaultTLSCertificateDefaultCertificate) {
	o.DefaultCertificate = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipDefaultTLSCertificate) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipDefaultTLSCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipDefaultTLSCertificate := _RelationshipDefaultTLSCertificate{}

	if err = json.Unmarshal(bytes, &varRelationshipDefaultTLSCertificate); err == nil {
		*o = RelationshipDefaultTLSCertificate(varRelationshipDefaultTLSCertificate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "default_certificate")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipDefaultTLSCertificate is a helper abstraction for handling nullable relationshipdefaulttlscertificate types. 
type NullableRelationshipDefaultTLSCertificate struct {
	value *RelationshipDefaultTLSCertificate
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipDefaultTLSCertificate) Get() *RelationshipDefaultTLSCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipDefaultTLSCertificate) Set(val *RelationshipDefaultTLSCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipDefaultTLSCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipDefaultTLSCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipDefaultTLSCertificate returns a pointer to a new instance of NullableRelationshipDefaultTLSCertificate.
func NewNullableRelationshipDefaultTLSCertificate(val *RelationshipDefaultTLSCertificate) *NullableRelationshipDefaultTLSCertificate {
	return &NullableRelationshipDefaultTLSCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipDefaultTLSCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipDefaultTLSCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
