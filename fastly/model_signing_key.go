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

// SigningKey Used to verify signatures of client keys.
type SigningKey struct {
	// A Base64-encoded Ed25519 public key that can be used to verify signatures of client keys.
	SigningKey *string `json:"signing_key,omitempty"`
	AdditionalProperties map[string]any
}

type _SigningKey SigningKey

// NewSigningKey instantiates a new SigningKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSigningKey() *SigningKey {
	this := SigningKey{}
	return &this
}

// NewSigningKeyWithDefaults instantiates a new SigningKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSigningKeyWithDefaults() *SigningKey {
	this := SigningKey{}
	return &this
}

// GetSigningKey returns the SigningKey field value if set, zero value otherwise.
func (o *SigningKey) GetSigningKey() string {
	if o == nil || o.SigningKey == nil {
		var ret string
		return ret
	}
	return *o.SigningKey
}

// GetSigningKeyOk returns a tuple with the SigningKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SigningKey) GetSigningKeyOk() (*string, bool) {
	if o == nil || o.SigningKey == nil {
		return nil, false
	}
	return o.SigningKey, true
}

// HasSigningKey returns a boolean if a field has been set.
func (o *SigningKey) HasSigningKey() bool {
	if o != nil && o.SigningKey != nil {
		return true
	}

	return false
}

// SetSigningKey gets a reference to the given string and assigns it to the SigningKey field.
func (o *SigningKey) SetSigningKey(v string) {
	o.SigningKey = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SigningKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.SigningKey != nil {
		toSerialize["signing_key"] = o.SigningKey
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *SigningKey) UnmarshalJSON(bytes []byte) (err error) {
	varSigningKey := _SigningKey{}

	if err = json.Unmarshal(bytes, &varSigningKey); err == nil {
		*o = SigningKey(varSigningKey)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "signing_key")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSigningKey is a helper abstraction for handling nullable signingkey types. 
type NullableSigningKey struct {
	value *SigningKey
	isSet bool
}

// Get returns the value.
func (v NullableSigningKey) Get() *SigningKey {
	return v.value
}

// Set modifies the value.
func (v *NullableSigningKey) Set(val *SigningKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSigningKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSigningKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSigningKey returns a pointer to a new instance of NullableSigningKey.
func NewNullableSigningKey(val *SigningKey) *NullableSigningKey {
	return &NullableSigningKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSigningKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableSigningKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
