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
	"time"
)

// ClientKey A Base64-encoded X25519 public key.
type ClientKey struct {
	// A Base64-encoded X25519 public key that can be used with a [libsodium-compatible sealed box](https://libsodium.gitbook.io/doc/public-key_cryptography/sealed_boxes) to encrypt secrets before upload.
	ClientKey *string `json:"client_key,omitempty"`
	// A Base64-encoded signature of the client key. The signature is generated using the signing key and must be verified before using the client key.
	Signature *string `json:"signature,omitempty"`
	// Date and time in ISO 8601 format.
	ExpiresAt            NullableTime `json:"expires_at,omitempty"`
	AdditionalProperties map[string]any
}

type _ClientKey ClientKey

// NewClientKey instantiates a new ClientKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientKey() *ClientKey {
	this := ClientKey{}
	return &this
}

// NewClientKeyWithDefaults instantiates a new ClientKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientKeyWithDefaults() *ClientKey {
	this := ClientKey{}
	return &this
}

// GetClientKey returns the ClientKey field value if set, zero value otherwise.
func (o *ClientKey) GetClientKey() string {
	if o == nil || o.ClientKey == nil {
		var ret string
		return ret
	}
	return *o.ClientKey
}

// GetClientKeyOk returns a tuple with the ClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientKey) GetClientKeyOk() (*string, bool) {
	if o == nil || o.ClientKey == nil {
		return nil, false
	}
	return o.ClientKey, true
}

// HasClientKey returns a boolean if a field has been set.
func (o *ClientKey) HasClientKey() bool {
	if o != nil && o.ClientKey != nil {
		return true
	}

	return false
}

// SetClientKey gets a reference to the given string and assigns it to the ClientKey field.
func (o *ClientKey) SetClientKey(v string) {
	o.ClientKey = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *ClientKey) GetSignature() string {
	if o == nil || o.Signature == nil {
		var ret string
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientKey) GetSignatureOk() (*string, bool) {
	if o == nil || o.Signature == nil {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *ClientKey) HasSignature() bool {
	if o != nil && o.Signature != nil {
		return true
	}

	return false
}

// SetSignature gets a reference to the given string and assigns it to the Signature field.
func (o *ClientKey) SetSignature(v string) {
	o.Signature = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClientKey) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientKey) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ClientKey) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *ClientKey) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *ClientKey) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *ClientKey) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ClientKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ClientKey != nil {
		toSerialize["client_key"] = o.ClientKey
	}
	if o.Signature != nil {
		toSerialize["signature"] = o.Signature
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ClientKey) UnmarshalJSON(bytes []byte) (err error) {
	varClientKey := _ClientKey{}

	if err = json.Unmarshal(bytes, &varClientKey); err == nil {
		*o = ClientKey(varClientKey)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "client_key")
		delete(additionalProperties, "signature")
		delete(additionalProperties, "expires_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableClientKey is a helper abstraction for handling nullable clientkey types.
type NullableClientKey struct {
	value *ClientKey
	isSet bool
}

// Get returns the value.
func (v NullableClientKey) Get() *ClientKey {
	return v.value
}

// Set modifies the value.
func (v *NullableClientKey) Set(val *ClientKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableClientKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableClientKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableClientKey returns a pointer to a new instance of NullableClientKey.
func NewNullableClientKey(val *ClientKey) *NullableClientKey {
	return &NullableClientKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableClientKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableClientKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
