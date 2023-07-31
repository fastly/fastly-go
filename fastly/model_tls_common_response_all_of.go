// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// TLSCommonResponseAllOf struct for TLSCommonResponseAllOf
type TLSCommonResponseAllOf struct {
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TLSCaCert NullableString `json:"tls_ca_cert,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TLSClientCert NullableString `json:"tls_client_cert,omitempty"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TLSClientKey NullableString `json:"tls_client_key,omitempty"`
	// The hostname used to verify a server's certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN).
	TLSCertHostname NullableString `json:"tls_cert_hostname,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCommonResponseAllOf TLSCommonResponseAllOf

// NewTLSCommonResponseAllOf instantiates a new TLSCommonResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCommonResponseAllOf() *TLSCommonResponseAllOf {
	this := TLSCommonResponseAllOf{}
	var tlsCaCert string = "null"
	this.TLSCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TLSClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TLSClientKey = *NewNullableString(&tlsClientKey)
	var tlsCertHostname string = "null"
	this.TLSCertHostname = *NewNullableString(&tlsCertHostname)
	return &this
}

// NewTLSCommonResponseAllOfWithDefaults instantiates a new TLSCommonResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCommonResponseAllOfWithDefaults() *TLSCommonResponseAllOf {
	this := TLSCommonResponseAllOf{}
	var tlsCaCert string = "null"
	this.TLSCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TLSClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TLSClientKey = *NewNullableString(&tlsClientKey)
	var tlsCertHostname string = "null"
	this.TLSCertHostname = *NewNullableString(&tlsCertHostname)
	return &this
}

// GetTLSCaCert returns the TLSCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSCommonResponseAllOf) GetTLSCaCert() string {
	if o == nil || o.TLSCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCaCert.Get()
}

// GetTLSCaCertOk returns a tuple with the TLSCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSCommonResponseAllOf) GetTLSCaCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSCaCert.Get(), o.TLSCaCert.IsSet()
}

// HasTLSCaCert returns a boolean if a field has been set.
func (o *TLSCommonResponseAllOf) HasTLSCaCert() bool {
	if o != nil && o.TLSCaCert.IsSet() {
		return true
	}

	return false
}

// SetTLSCaCert gets a reference to the given NullableString and assigns it to the TLSCaCert field.
func (o *TLSCommonResponseAllOf) SetTLSCaCert(v string) {
	o.TLSCaCert.Set(&v)
}
// SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil
func (o *TLSCommonResponseAllOf) SetTLSCaCertNil() {
	o.TLSCaCert.Set(nil)
}

// UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
func (o *TLSCommonResponseAllOf) UnsetTLSCaCert() {
	o.TLSCaCert.Unset()
}

// GetTLSClientCert returns the TLSClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSCommonResponseAllOf) GetTLSClientCert() string {
	if o == nil || o.TLSClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientCert.Get()
}

// GetTLSClientCertOk returns a tuple with the TLSClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSCommonResponseAllOf) GetTLSClientCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSClientCert.Get(), o.TLSClientCert.IsSet()
}

// HasTLSClientCert returns a boolean if a field has been set.
func (o *TLSCommonResponseAllOf) HasTLSClientCert() bool {
	if o != nil && o.TLSClientCert.IsSet() {
		return true
	}

	return false
}

// SetTLSClientCert gets a reference to the given NullableString and assigns it to the TLSClientCert field.
func (o *TLSCommonResponseAllOf) SetTLSClientCert(v string) {
	o.TLSClientCert.Set(&v)
}
// SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil
func (o *TLSCommonResponseAllOf) SetTLSClientCertNil() {
	o.TLSClientCert.Set(nil)
}

// UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
func (o *TLSCommonResponseAllOf) UnsetTLSClientCert() {
	o.TLSClientCert.Unset()
}

// GetTLSClientKey returns the TLSClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSCommonResponseAllOf) GetTLSClientKey() string {
	if o == nil || o.TLSClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientKey.Get()
}

// GetTLSClientKeyOk returns a tuple with the TLSClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSCommonResponseAllOf) GetTLSClientKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSClientKey.Get(), o.TLSClientKey.IsSet()
}

// HasTLSClientKey returns a boolean if a field has been set.
func (o *TLSCommonResponseAllOf) HasTLSClientKey() bool {
	if o != nil && o.TLSClientKey.IsSet() {
		return true
	}

	return false
}

// SetTLSClientKey gets a reference to the given NullableString and assigns it to the TLSClientKey field.
func (o *TLSCommonResponseAllOf) SetTLSClientKey(v string) {
	o.TLSClientKey.Set(&v)
}
// SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil
func (o *TLSCommonResponseAllOf) SetTLSClientKeyNil() {
	o.TLSClientKey.Set(nil)
}

// UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
func (o *TLSCommonResponseAllOf) UnsetTLSClientKey() {
	o.TLSClientKey.Unset()
}

// GetTLSCertHostname returns the TLSCertHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSCommonResponseAllOf) GetTLSCertHostname() string {
	if o == nil || o.TLSCertHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCertHostname.Get()
}

// GetTLSCertHostnameOk returns a tuple with the TLSCertHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSCommonResponseAllOf) GetTLSCertHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSCertHostname.Get(), o.TLSCertHostname.IsSet()
}

// HasTLSCertHostname returns a boolean if a field has been set.
func (o *TLSCommonResponseAllOf) HasTLSCertHostname() bool {
	if o != nil && o.TLSCertHostname.IsSet() {
		return true
	}

	return false
}

// SetTLSCertHostname gets a reference to the given NullableString and assigns it to the TLSCertHostname field.
func (o *TLSCommonResponseAllOf) SetTLSCertHostname(v string) {
	o.TLSCertHostname.Set(&v)
}
// SetTLSCertHostnameNil sets the value for TLSCertHostname to be an explicit nil
func (o *TLSCommonResponseAllOf) SetTLSCertHostnameNil() {
	o.TLSCertHostname.Set(nil)
}

// UnsetTLSCertHostname ensures that no value is present for TLSCertHostname, not even an explicit nil
func (o *TLSCommonResponseAllOf) UnsetTLSCertHostname() {
	o.TLSCertHostname.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCommonResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSCaCert.IsSet() {
		toSerialize["tls_ca_cert"] = o.TLSCaCert.Get()
	}
	if o.TLSClientCert.IsSet() {
		toSerialize["tls_client_cert"] = o.TLSClientCert.Get()
	}
	if o.TLSClientKey.IsSet() {
		toSerialize["tls_client_key"] = o.TLSClientKey.Get()
	}
	if o.TLSCertHostname.IsSet() {
		toSerialize["tls_cert_hostname"] = o.TLSCertHostname.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *TLSCommonResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCommonResponseAllOf := _TLSCommonResponseAllOf{}

	if err = json.Unmarshal(bytes, &varTLSCommonResponseAllOf); err == nil {
		*o = TLSCommonResponseAllOf(varTLSCommonResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_ca_cert")
		delete(additionalProperties, "tls_client_cert")
		delete(additionalProperties, "tls_client_key")
		delete(additionalProperties, "tls_cert_hostname")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSCommonResponseAllOf is a helper abstraction for handling nullable tlscommonresponseallof types. 
type NullableTLSCommonResponseAllOf struct {
	value *TLSCommonResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTLSCommonResponseAllOf) Get() *TLSCommonResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCommonResponseAllOf) Set(val *TLSCommonResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCommonResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCommonResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCommonResponseAllOf returns a pointer to a new instance of NullableTLSCommonResponseAllOf.
func NewNullableTLSCommonResponseAllOf(val *TLSCommonResponseAllOf) *NullableTLSCommonResponseAllOf {
	return &NullableTLSCommonResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCommonResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSCommonResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
