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

// TlsCommonResponseAllOf1 struct for TlsCommonResponseAllOf1
type TlsCommonResponseAllOf1 struct {
	// Whether to use TLS.
	UseTls               *string `json:"use_tls,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsCommonResponseAllOf1 TlsCommonResponseAllOf1

// NewTlsCommonResponseAllOf1 instantiates a new TlsCommonResponseAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsCommonResponseAllOf1() *TlsCommonResponseAllOf1 {
	this := TlsCommonResponseAllOf1{}
	var useTls string = "0"
	this.UseTls = &useTls
	return &this
}

// NewTlsCommonResponseAllOf1WithDefaults instantiates a new TlsCommonResponseAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsCommonResponseAllOf1WithDefaults() *TlsCommonResponseAllOf1 {
	this := TlsCommonResponseAllOf1{}
	var useTls string = "0"
	this.UseTls = &useTls
	return &this
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *TlsCommonResponseAllOf1) GetUseTls() string {
	if o == nil || o.UseTls == nil {
		var ret string
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCommonResponseAllOf1) GetUseTlsOk() (*string, bool) {
	if o == nil || o.UseTls == nil {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *TlsCommonResponseAllOf1) HasUseTls() bool {
	if o != nil && o.UseTls != nil {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given string and assigns it to the UseTls field.
func (o *TlsCommonResponseAllOf1) SetUseTls(v string) {
	o.UseTls = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsCommonResponseAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.UseTls != nil {
		toSerialize["use_tls"] = o.UseTls
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsCommonResponseAllOf1) UnmarshalJSON(bytes []byte) (err error) {
	varTlsCommonResponseAllOf1 := _TlsCommonResponseAllOf1{}

	if err = json.Unmarshal(bytes, &varTlsCommonResponseAllOf1); err == nil {
		*o = TlsCommonResponseAllOf1(varTlsCommonResponseAllOf1)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "use_tls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsCommonResponseAllOf1 is a helper abstraction for handling nullable tlscommonresponseallof1 types.
type NullableTlsCommonResponseAllOf1 struct {
	value *TlsCommonResponseAllOf1
	isSet bool
}

// Get returns the value.
func (v NullableTlsCommonResponseAllOf1) Get() *TlsCommonResponseAllOf1 {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsCommonResponseAllOf1) Set(val *TlsCommonResponseAllOf1) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsCommonResponseAllOf1) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsCommonResponseAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsCommonResponseAllOf1 returns a pointer to a new instance of NullableTlsCommonResponseAllOf1.
func NewNullableTlsCommonResponseAllOf1(val *TlsCommonResponseAllOf1) *NullableTlsCommonResponseAllOf1 {
	return &NullableTlsCommonResponseAllOf1{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsCommonResponseAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsCommonResponseAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
