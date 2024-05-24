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

// RelationshipTLSPrivateKey struct for RelationshipTLSPrivateKey
type RelationshipTLSPrivateKey struct {
	TLSPrivateKey *RelationshipTLSPrivateKeyTLSPrivateKey `json:"tls_private_key,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSPrivateKey RelationshipTLSPrivateKey

// NewRelationshipTLSPrivateKey instantiates a new RelationshipTLSPrivateKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSPrivateKey() *RelationshipTLSPrivateKey {
	this := RelationshipTLSPrivateKey{}
	return &this
}

// NewRelationshipTLSPrivateKeyWithDefaults instantiates a new RelationshipTLSPrivateKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSPrivateKeyWithDefaults() *RelationshipTLSPrivateKey {
	this := RelationshipTLSPrivateKey{}
	return &this
}

// GetTLSPrivateKey returns the TLSPrivateKey field value if set, zero value otherwise.
func (o *RelationshipTLSPrivateKey) GetTLSPrivateKey() RelationshipTLSPrivateKeyTLSPrivateKey {
	if o == nil || o.TLSPrivateKey == nil {
		var ret RelationshipTLSPrivateKeyTLSPrivateKey
		return ret
	}
	return *o.TLSPrivateKey
}

// GetTLSPrivateKeyOk returns a tuple with the TLSPrivateKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSPrivateKey) GetTLSPrivateKeyOk() (*RelationshipTLSPrivateKeyTLSPrivateKey, bool) {
	if o == nil || o.TLSPrivateKey == nil {
		return nil, false
	}
	return o.TLSPrivateKey, true
}

// HasTLSPrivateKey returns a boolean if a field has been set.
func (o *RelationshipTLSPrivateKey) HasTLSPrivateKey() bool {
	if o != nil && o.TLSPrivateKey != nil {
		return true
	}

	return false
}

// SetTLSPrivateKey gets a reference to the given RelationshipTLSPrivateKeyTLSPrivateKey and assigns it to the TLSPrivateKey field.
func (o *RelationshipTLSPrivateKey) SetTLSPrivateKey(v RelationshipTLSPrivateKeyTLSPrivateKey) {
	o.TLSPrivateKey = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSPrivateKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSPrivateKey != nil {
		toSerialize["tls_private_key"] = o.TLSPrivateKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSPrivateKey) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSPrivateKey := _RelationshipTLSPrivateKey{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSPrivateKey); err == nil {
		*o = RelationshipTLSPrivateKey(varRelationshipTLSPrivateKey)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_private_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSPrivateKey is a helper abstraction for handling nullable relationshiptlsprivatekey types. 
type NullableRelationshipTLSPrivateKey struct {
	value *RelationshipTLSPrivateKey
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSPrivateKey) Get() *RelationshipTLSPrivateKey {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSPrivateKey) Set(val *RelationshipTLSPrivateKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSPrivateKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSPrivateKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSPrivateKey returns a pointer to a new instance of NullableRelationshipTLSPrivateKey.
func NewNullableRelationshipTLSPrivateKey(val *RelationshipTLSPrivateKey) *NullableRelationshipTLSPrivateKey {
	return &NullableRelationshipTLSPrivateKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSPrivateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSPrivateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
