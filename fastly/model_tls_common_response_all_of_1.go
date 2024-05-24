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

// TLSCommonResponseAllOf1 struct for TLSCommonResponseAllOf1
type TLSCommonResponseAllOf1 struct {
	// Whether to use TLS.
	UseTLS *string `json:"use_tls,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCommonResponseAllOf1 TLSCommonResponseAllOf1

// NewTLSCommonResponseAllOf1 instantiates a new TLSCommonResponseAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCommonResponseAllOf1() *TLSCommonResponseAllOf1 {
	this := TLSCommonResponseAllOf1{}
	var useTLS string = "0"
	this.UseTLS = &useTLS
	return &this
}

// NewTLSCommonResponseAllOf1WithDefaults instantiates a new TLSCommonResponseAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCommonResponseAllOf1WithDefaults() *TLSCommonResponseAllOf1 {
	this := TLSCommonResponseAllOf1{}
	var useTLS string = "0"
	this.UseTLS = &useTLS
	return &this
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *TLSCommonResponseAllOf1) GetUseTLS() string {
	if o == nil || o.UseTLS == nil {
		var ret string
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCommonResponseAllOf1) GetUseTLSOk() (*string, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *TLSCommonResponseAllOf1) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given string and assigns it to the UseTLS field.
func (o *TLSCommonResponseAllOf1) SetUseTLS(v string) {
	o.UseTLS = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCommonResponseAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.UseTLS != nil {
		toSerialize["use_tls"] = o.UseTLS
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *TLSCommonResponseAllOf1) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCommonResponseAllOf1 := _TLSCommonResponseAllOf1{}

	if err = json.Unmarshal(bytes, &varTLSCommonResponseAllOf1); err == nil {
		*o = TLSCommonResponseAllOf1(varTLSCommonResponseAllOf1)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "use_tls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSCommonResponseAllOf1 is a helper abstraction for handling nullable tlscommonresponseallof1 types. 
type NullableTLSCommonResponseAllOf1 struct {
	value *TLSCommonResponseAllOf1
	isSet bool
}

// Get returns the value.
func (v NullableTLSCommonResponseAllOf1) Get() *TLSCommonResponseAllOf1 {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCommonResponseAllOf1) Set(val *TLSCommonResponseAllOf1) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCommonResponseAllOf1) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCommonResponseAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCommonResponseAllOf1 returns a pointer to a new instance of NullableTLSCommonResponseAllOf1.
func NewNullableTLSCommonResponseAllOf1(val *TLSCommonResponseAllOf1) *NullableTLSCommonResponseAllOf1 {
	return &NullableTLSCommonResponseAllOf1{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCommonResponseAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSCommonResponseAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
