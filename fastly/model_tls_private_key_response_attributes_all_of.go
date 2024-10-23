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

// TLSPrivateKeyResponseAttributesAllOf struct for TLSPrivateKeyResponseAttributesAllOf
type TLSPrivateKeyResponseAttributesAllOf struct {
	// A customizable name for your private key.
	Name *string `json:"name,omitempty"`
	// The key length used to generate the private key.
	KeyLength *int32 `json:"key_length,omitempty"`
	// The algorithm used to generate the private key. Must be `RSA`.
	KeyType *string `json:"key_type,omitempty"`
	// A recommendation from Fastly to replace this private key and all associated certificates.
	Replace *bool `json:"replace,omitempty"`
	// Useful for safely identifying the key.
	PublicKeySha1        *string `json:"public_key_sha1,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSPrivateKeyResponseAttributesAllOf TLSPrivateKeyResponseAttributesAllOf

// NewTLSPrivateKeyResponseAttributesAllOf instantiates a new TLSPrivateKeyResponseAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSPrivateKeyResponseAttributesAllOf() *TLSPrivateKeyResponseAttributesAllOf {
	this := TLSPrivateKeyResponseAttributesAllOf{}
	return &this
}

// NewTLSPrivateKeyResponseAttributesAllOfWithDefaults instantiates a new TLSPrivateKeyResponseAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSPrivateKeyResponseAttributesAllOfWithDefaults() *TLSPrivateKeyResponseAttributesAllOf {
	this := TLSPrivateKeyResponseAttributesAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseAttributesAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseAttributesAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributesAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TLSPrivateKeyResponseAttributesAllOf) SetName(v string) {
	o.Name = &v
}

// GetKeyLength returns the KeyLength field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseAttributesAllOf) GetKeyLength() int32 {
	if o == nil || o.KeyLength == nil {
		var ret int32
		return ret
	}
	return *o.KeyLength
}

// GetKeyLengthOk returns a tuple with the KeyLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseAttributesAllOf) GetKeyLengthOk() (*int32, bool) {
	if o == nil || o.KeyLength == nil {
		return nil, false
	}
	return o.KeyLength, true
}

// HasKeyLength returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributesAllOf) HasKeyLength() bool {
	if o != nil && o.KeyLength != nil {
		return true
	}

	return false
}

// SetKeyLength gets a reference to the given int32 and assigns it to the KeyLength field.
func (o *TLSPrivateKeyResponseAttributesAllOf) SetKeyLength(v int32) {
	o.KeyLength = &v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseAttributesAllOf) GetKeyType() string {
	if o == nil || o.KeyType == nil {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseAttributesAllOf) GetKeyTypeOk() (*string, bool) {
	if o == nil || o.KeyType == nil {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributesAllOf) HasKeyType() bool {
	if o != nil && o.KeyType != nil {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *TLSPrivateKeyResponseAttributesAllOf) SetKeyType(v string) {
	o.KeyType = &v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseAttributesAllOf) GetReplace() bool {
	if o == nil || o.Replace == nil {
		var ret bool
		return ret
	}
	return *o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseAttributesAllOf) GetReplaceOk() (*bool, bool) {
	if o == nil || o.Replace == nil {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributesAllOf) HasReplace() bool {
	if o != nil && o.Replace != nil {
		return true
	}

	return false
}

// SetReplace gets a reference to the given bool and assigns it to the Replace field.
func (o *TLSPrivateKeyResponseAttributesAllOf) SetReplace(v bool) {
	o.Replace = &v
}

// GetPublicKeySha1 returns the PublicKeySha1 field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseAttributesAllOf) GetPublicKeySha1() string {
	if o == nil || o.PublicKeySha1 == nil {
		var ret string
		return ret
	}
	return *o.PublicKeySha1
}

// GetPublicKeySha1Ok returns a tuple with the PublicKeySha1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseAttributesAllOf) GetPublicKeySha1Ok() (*string, bool) {
	if o == nil || o.PublicKeySha1 == nil {
		return nil, false
	}
	return o.PublicKeySha1, true
}

// HasPublicKeySha1 returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseAttributesAllOf) HasPublicKeySha1() bool {
	if o != nil && o.PublicKeySha1 != nil {
		return true
	}

	return false
}

// SetPublicKeySha1 gets a reference to the given string and assigns it to the PublicKeySha1 field.
func (o *TLSPrivateKeyResponseAttributesAllOf) SetPublicKeySha1(v string) {
	o.PublicKeySha1 = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSPrivateKeyResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.KeyLength != nil {
		toSerialize["key_length"] = o.KeyLength
	}
	if o.KeyType != nil {
		toSerialize["key_type"] = o.KeyType
	}
	if o.Replace != nil {
		toSerialize["replace"] = o.Replace
	}
	if o.PublicKeySha1 != nil {
		toSerialize["public_key_sha1"] = o.PublicKeySha1
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TLSPrivateKeyResponseAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTLSPrivateKeyResponseAttributesAllOf := _TLSPrivateKeyResponseAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varTLSPrivateKeyResponseAttributesAllOf); err == nil {
		*o = TLSPrivateKeyResponseAttributesAllOf(varTLSPrivateKeyResponseAttributesAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "key_length")
		delete(additionalProperties, "key_type")
		delete(additionalProperties, "replace")
		delete(additionalProperties, "public_key_sha1")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSPrivateKeyResponseAttributesAllOf is a helper abstraction for handling nullable tlsprivatekeyresponseattributesallof types.
type NullableTLSPrivateKeyResponseAttributesAllOf struct {
	value *TLSPrivateKeyResponseAttributesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTLSPrivateKeyResponseAttributesAllOf) Get() *TLSPrivateKeyResponseAttributesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSPrivateKeyResponseAttributesAllOf) Set(val *TLSPrivateKeyResponseAttributesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSPrivateKeyResponseAttributesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSPrivateKeyResponseAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSPrivateKeyResponseAttributesAllOf returns a pointer to a new instance of NullableTLSPrivateKeyResponseAttributesAllOf.
func NewNullableTLSPrivateKeyResponseAttributesAllOf(val *TLSPrivateKeyResponseAttributesAllOf) *NullableTLSPrivateKeyResponseAttributesAllOf {
	return &NullableTLSPrivateKeyResponseAttributesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSPrivateKeyResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTLSPrivateKeyResponseAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
