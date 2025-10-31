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

// RelationshipTlsPrivateKeys struct for RelationshipTlsPrivateKeys
type RelationshipTlsPrivateKeys struct {
	TlsPrivateKeys       *RelationshipTlsPrivateKeysTlsPrivateKeys `json:"tls_private_keys,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsPrivateKeys RelationshipTlsPrivateKeys

// NewRelationshipTlsPrivateKeys instantiates a new RelationshipTlsPrivateKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsPrivateKeys() *RelationshipTlsPrivateKeys {
	this := RelationshipTlsPrivateKeys{}
	return &this
}

// NewRelationshipTlsPrivateKeysWithDefaults instantiates a new RelationshipTlsPrivateKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsPrivateKeysWithDefaults() *RelationshipTlsPrivateKeys {
	this := RelationshipTlsPrivateKeys{}
	return &this
}

// GetTlsPrivateKeys returns the TlsPrivateKeys field value if set, zero value otherwise.
func (o *RelationshipTlsPrivateKeys) GetTlsPrivateKeys() RelationshipTlsPrivateKeysTlsPrivateKeys {
	if o == nil || o.TlsPrivateKeys == nil {
		var ret RelationshipTlsPrivateKeysTlsPrivateKeys
		return ret
	}
	return *o.TlsPrivateKeys
}

// GetTlsPrivateKeysOk returns a tuple with the TlsPrivateKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsPrivateKeys) GetTlsPrivateKeysOk() (*RelationshipTlsPrivateKeysTlsPrivateKeys, bool) {
	if o == nil || o.TlsPrivateKeys == nil {
		return nil, false
	}
	return o.TlsPrivateKeys, true
}

// HasTlsPrivateKeys returns a boolean if a field has been set.
func (o *RelationshipTlsPrivateKeys) HasTlsPrivateKeys() bool {
	if o != nil && o.TlsPrivateKeys != nil {
		return true
	}

	return false
}

// SetTlsPrivateKeys gets a reference to the given RelationshipTlsPrivateKeysTlsPrivateKeys and assigns it to the TlsPrivateKeys field.
func (o *RelationshipTlsPrivateKeys) SetTlsPrivateKeys(v RelationshipTlsPrivateKeysTlsPrivateKeys) {
	o.TlsPrivateKeys = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsPrivateKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsPrivateKeys != nil {
		toSerialize["tls_private_keys"] = o.TlsPrivateKeys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsPrivateKeys) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsPrivateKeys := _RelationshipTlsPrivateKeys{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsPrivateKeys); err == nil {
		*o = RelationshipTlsPrivateKeys(varRelationshipTlsPrivateKeys)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_private_keys")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsPrivateKeys is a helper abstraction for handling nullable relationshiptlsprivatekeys types.
type NullableRelationshipTlsPrivateKeys struct {
	value *RelationshipTlsPrivateKeys
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsPrivateKeys) Get() *RelationshipTlsPrivateKeys {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsPrivateKeys) Set(val *RelationshipTlsPrivateKeys) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsPrivateKeys) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsPrivateKeys) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsPrivateKeys returns a pointer to a new instance of NullableRelationshipTlsPrivateKeys.
func NewNullableRelationshipTlsPrivateKeys(val *RelationshipTlsPrivateKeys) *NullableRelationshipTlsPrivateKeys {
	return &NullableRelationshipTlsPrivateKeys{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsPrivateKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsPrivateKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
