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

// TLSPrivateKey struct for TLSPrivateKey
type TLSPrivateKey struct {
	Data *TLSPrivateKeyData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSPrivateKey TLSPrivateKey

// NewTLSPrivateKey instantiates a new TLSPrivateKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSPrivateKey() *TLSPrivateKey {
	this := TLSPrivateKey{}
	return &this
}

// NewTLSPrivateKeyWithDefaults instantiates a new TLSPrivateKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSPrivateKeyWithDefaults() *TLSPrivateKey {
	this := TLSPrivateKey{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TLSPrivateKey) GetData() TLSPrivateKeyData {
	if o == nil || o.Data == nil {
		var ret TLSPrivateKeyData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKey) GetDataOk() (*TLSPrivateKeyData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TLSPrivateKey) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given TLSPrivateKeyData and assigns it to the Data field.
func (o *TLSPrivateKey) SetData(v TLSPrivateKeyData) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSPrivateKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *TLSPrivateKey) UnmarshalJSON(bytes []byte) (err error) {
	varTLSPrivateKey := _TLSPrivateKey{}

	if err = json.Unmarshal(bytes, &varTLSPrivateKey); err == nil {
		*o = TLSPrivateKey(varTLSPrivateKey)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSPrivateKey is a helper abstraction for handling nullable tlsprivatekey types. 
type NullableTLSPrivateKey struct {
	value *TLSPrivateKey
	isSet bool
}

// Get returns the value.
func (v NullableTLSPrivateKey) Get() *TLSPrivateKey {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSPrivateKey) Set(val *TLSPrivateKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSPrivateKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSPrivateKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSPrivateKey returns a pointer to a new instance of NullableTLSPrivateKey.
func NewNullableTLSPrivateKey(val *TLSPrivateKey) *NullableTLSPrivateKey {
	return &NullableTLSPrivateKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSPrivateKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSPrivateKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
