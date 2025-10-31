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

// RelationshipTlsPrivateKey struct for RelationshipTlsPrivateKey
type RelationshipTlsPrivateKey struct {
	TlsPrivateKey        *RelationshipTlsPrivateKeyTlsPrivateKey `json:"tls_private_key,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsPrivateKey RelationshipTlsPrivateKey

// NewRelationshipTlsPrivateKey instantiates a new RelationshipTlsPrivateKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsPrivateKey() *RelationshipTlsPrivateKey {
	this := RelationshipTlsPrivateKey{}
	return &this
}

// NewRelationshipTlsPrivateKeyWithDefaults instantiates a new RelationshipTlsPrivateKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsPrivateKeyWithDefaults() *RelationshipTlsPrivateKey {
	this := RelationshipTlsPrivateKey{}
	return &this
}

// GetTlsPrivateKey returns the TlsPrivateKey field value if set, zero value otherwise.
func (o *RelationshipTlsPrivateKey) GetTlsPrivateKey() RelationshipTlsPrivateKeyTlsPrivateKey {
	if o == nil || o.TlsPrivateKey == nil {
		var ret RelationshipTlsPrivateKeyTlsPrivateKey
		return ret
	}
	return *o.TlsPrivateKey
}

// GetTlsPrivateKeyOk returns a tuple with the TlsPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsPrivateKey) GetTlsPrivateKeyOk() (*RelationshipTlsPrivateKeyTlsPrivateKey, bool) {
	if o == nil || o.TlsPrivateKey == nil {
		return nil, false
	}
	return o.TlsPrivateKey, true
}

// HasTlsPrivateKey returns a boolean if a field has been set.
func (o *RelationshipTlsPrivateKey) HasTlsPrivateKey() bool {
	if o != nil && o.TlsPrivateKey != nil {
		return true
	}

	return false
}

// SetTlsPrivateKey gets a reference to the given RelationshipTlsPrivateKeyTlsPrivateKey and assigns it to the TlsPrivateKey field.
func (o *RelationshipTlsPrivateKey) SetTlsPrivateKey(v RelationshipTlsPrivateKeyTlsPrivateKey) {
	o.TlsPrivateKey = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsPrivateKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsPrivateKey != nil {
		toSerialize["tls_private_key"] = o.TlsPrivateKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsPrivateKey) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsPrivateKey := _RelationshipTlsPrivateKey{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsPrivateKey); err == nil {
		*o = RelationshipTlsPrivateKey(varRelationshipTlsPrivateKey)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_private_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsPrivateKey is a helper abstraction for handling nullable relationshiptlsprivatekey types.
type NullableRelationshipTlsPrivateKey struct {
	value *RelationshipTlsPrivateKey
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsPrivateKey) Get() *RelationshipTlsPrivateKey {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsPrivateKey) Set(val *RelationshipTlsPrivateKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsPrivateKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsPrivateKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsPrivateKey returns a pointer to a new instance of NullableRelationshipTlsPrivateKey.
func NewNullableRelationshipTlsPrivateKey(val *RelationshipTlsPrivateKey) *NullableRelationshipTlsPrivateKey {
	return &NullableRelationshipTlsPrivateKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsPrivateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsPrivateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
