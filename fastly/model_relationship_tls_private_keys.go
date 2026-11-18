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

// RelationshipTLSPrivateKeys struct for RelationshipTLSPrivateKeys
type RelationshipTLSPrivateKeys struct {
	TLSPrivateKeys       *RelationshipTLSPrivateKeysTLSPrivateKeys `json:"tls_private_keys,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSPrivateKeys RelationshipTLSPrivateKeys

// NewRelationshipTLSPrivateKeys instantiates a new RelationshipTLSPrivateKeys object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSPrivateKeys() *RelationshipTLSPrivateKeys {
	this := RelationshipTLSPrivateKeys{}
	return &this
}

// NewRelationshipTLSPrivateKeysWithDefaults instantiates a new RelationshipTLSPrivateKeys object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSPrivateKeysWithDefaults() *RelationshipTLSPrivateKeys {
	this := RelationshipTLSPrivateKeys{}
	return &this
}

// GetTLSPrivateKeys returns the TLSPrivateKeys field value if set, zero value otherwise.
func (o *RelationshipTLSPrivateKeys) GetTLSPrivateKeys() RelationshipTLSPrivateKeysTLSPrivateKeys {
	if o == nil || o.TLSPrivateKeys == nil {
		var ret RelationshipTLSPrivateKeysTLSPrivateKeys
		return ret
	}
	return *o.TLSPrivateKeys
}

// GetTLSPrivateKeysOk returns a tuple with the TLSPrivateKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSPrivateKeys) GetTLSPrivateKeysOk() (*RelationshipTLSPrivateKeysTLSPrivateKeys, bool) {
	if o == nil || o.TLSPrivateKeys == nil {
		return nil, false
	}
	return o.TLSPrivateKeys, true
}

// HasTLSPrivateKeys returns a boolean if a field has been set.
func (o *RelationshipTLSPrivateKeys) HasTLSPrivateKeys() bool {
	if o != nil && o.TLSPrivateKeys != nil {
		return true
	}

	return false
}

// SetTLSPrivateKeys gets a reference to the given RelationshipTLSPrivateKeysTLSPrivateKeys and assigns it to the TLSPrivateKeys field.
func (o *RelationshipTLSPrivateKeys) SetTLSPrivateKeys(v RelationshipTLSPrivateKeysTLSPrivateKeys) {
	o.TLSPrivateKeys = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSPrivateKeys) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSPrivateKeys != nil {
		toSerialize["tls_private_keys"] = o.TLSPrivateKeys
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTLSPrivateKeys) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSPrivateKeys := _RelationshipTLSPrivateKeys{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSPrivateKeys); err == nil {
		*o = RelationshipTLSPrivateKeys(varRelationshipTLSPrivateKeys)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_private_keys")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSPrivateKeys is a helper abstraction for handling nullable relationshiptlsprivatekeys types.
type NullableRelationshipTLSPrivateKeys struct {
	value *RelationshipTLSPrivateKeys
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSPrivateKeys) Get() *RelationshipTLSPrivateKeys {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSPrivateKeys) Set(val *RelationshipTLSPrivateKeys) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSPrivateKeys) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSPrivateKeys) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSPrivateKeys returns a pointer to a new instance of NullableRelationshipTLSPrivateKeys.
func NewNullableRelationshipTLSPrivateKeys(val *RelationshipTLSPrivateKeys) *NullableRelationshipTLSPrivateKeys {
	return &NullableRelationshipTLSPrivateKeys{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSPrivateKeys) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTLSPrivateKeys) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
