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

// TLSCommon struct for TLSCommon
type TLSCommon struct {
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TLSCaCert NullableString `json:"tls_ca_cert,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TLSClientCert NullableString `json:"tls_client_cert,omitempty"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TLSClientKey NullableString `json:"tls_client_key,omitempty"`
	// The hostname used to verify a server's certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN).
	TLSCertHostname NullableString `json:"tls_cert_hostname,omitempty"`
	// Whether to use TLS.
	UseTLS *int32 `json:"use_tls,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCommon TLSCommon

// NewTLSCommon instantiates a new TLSCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCommon() *TLSCommon {
	this := TLSCommon{}
	var tlsCaCert string = "null"
	this.TLSCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TLSClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TLSClientKey = *NewNullableString(&tlsClientKey)
	var tlsCertHostname string = "null"
	this.TLSCertHostname = *NewNullableString(&tlsCertHostname)
	var useTLS int32 = 0
	this.UseTLS = &useTLS
	return &this
}

// NewTLSCommonWithDefaults instantiates a new TLSCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCommonWithDefaults() *TLSCommon {
	this := TLSCommon{}
	var tlsCaCert string = "null"
	this.TLSCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TLSClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TLSClientKey = *NewNullableString(&tlsClientKey)
	var tlsCertHostname string = "null"
	this.TLSCertHostname = *NewNullableString(&tlsCertHostname)
	var useTLS int32 = 0
	this.UseTLS = &useTLS
	return &this
}

// GetTLSCaCert returns the TLSCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSCommon) GetTLSCaCert() string {
	if o == nil || o.TLSCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCaCert.Get()
}

// GetTLSCaCertOk returns a tuple with the TLSCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSCommon) GetTLSCaCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSCaCert.Get(), o.TLSCaCert.IsSet()
}

// HasTLSCaCert returns a boolean if a field has been set.
func (o *TLSCommon) HasTLSCaCert() bool {
	if o != nil && o.TLSCaCert.IsSet() {
		return true
	}

	return false
}

// SetTLSCaCert gets a reference to the given NullableString and assigns it to the TLSCaCert field.
func (o *TLSCommon) SetTLSCaCert(v string) {
	o.TLSCaCert.Set(&v)
}
// SetTLSCaCertNil sets the value for TLSCaCert to be an explicit nil
func (o *TLSCommon) SetTLSCaCertNil() {
	o.TLSCaCert.Set(nil)
}

// UnsetTLSCaCert ensures that no value is present for TLSCaCert, not even an explicit nil
func (o *TLSCommon) UnsetTLSCaCert() {
	o.TLSCaCert.Unset()
}

// GetTLSClientCert returns the TLSClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSCommon) GetTLSClientCert() string {
	if o == nil || o.TLSClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientCert.Get()
}

// GetTLSClientCertOk returns a tuple with the TLSClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSCommon) GetTLSClientCertOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSClientCert.Get(), o.TLSClientCert.IsSet()
}

// HasTLSClientCert returns a boolean if a field has been set.
func (o *TLSCommon) HasTLSClientCert() bool {
	if o != nil && o.TLSClientCert.IsSet() {
		return true
	}

	return false
}

// SetTLSClientCert gets a reference to the given NullableString and assigns it to the TLSClientCert field.
func (o *TLSCommon) SetTLSClientCert(v string) {
	o.TLSClientCert.Set(&v)
}
// SetTLSClientCertNil sets the value for TLSClientCert to be an explicit nil
func (o *TLSCommon) SetTLSClientCertNil() {
	o.TLSClientCert.Set(nil)
}

// UnsetTLSClientCert ensures that no value is present for TLSClientCert, not even an explicit nil
func (o *TLSCommon) UnsetTLSClientCert() {
	o.TLSClientCert.Unset()
}

// GetTLSClientKey returns the TLSClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSCommon) GetTLSClientKey() string {
	if o == nil || o.TLSClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSClientKey.Get()
}

// GetTLSClientKeyOk returns a tuple with the TLSClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSCommon) GetTLSClientKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSClientKey.Get(), o.TLSClientKey.IsSet()
}

// HasTLSClientKey returns a boolean if a field has been set.
func (o *TLSCommon) HasTLSClientKey() bool {
	if o != nil && o.TLSClientKey.IsSet() {
		return true
	}

	return false
}

// SetTLSClientKey gets a reference to the given NullableString and assigns it to the TLSClientKey field.
func (o *TLSCommon) SetTLSClientKey(v string) {
	o.TLSClientKey.Set(&v)
}
// SetTLSClientKeyNil sets the value for TLSClientKey to be an explicit nil
func (o *TLSCommon) SetTLSClientKeyNil() {
	o.TLSClientKey.Set(nil)
}

// UnsetTLSClientKey ensures that no value is present for TLSClientKey, not even an explicit nil
func (o *TLSCommon) UnsetTLSClientKey() {
	o.TLSClientKey.Unset()
}

// GetTLSCertHostname returns the TLSCertHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSCommon) GetTLSCertHostname() string {
	if o == nil || o.TLSCertHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TLSCertHostname.Get()
}

// GetTLSCertHostnameOk returns a tuple with the TLSCertHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSCommon) GetTLSCertHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TLSCertHostname.Get(), o.TLSCertHostname.IsSet()
}

// HasTLSCertHostname returns a boolean if a field has been set.
func (o *TLSCommon) HasTLSCertHostname() bool {
	if o != nil && o.TLSCertHostname.IsSet() {
		return true
	}

	return false
}

// SetTLSCertHostname gets a reference to the given NullableString and assigns it to the TLSCertHostname field.
func (o *TLSCommon) SetTLSCertHostname(v string) {
	o.TLSCertHostname.Set(&v)
}
// SetTLSCertHostnameNil sets the value for TLSCertHostname to be an explicit nil
func (o *TLSCommon) SetTLSCertHostnameNil() {
	o.TLSCertHostname.Set(nil)
}

// UnsetTLSCertHostname ensures that no value is present for TLSCertHostname, not even an explicit nil
func (o *TLSCommon) UnsetTLSCertHostname() {
	o.TLSCertHostname.Unset()
}

// GetUseTLS returns the UseTLS field value if set, zero value otherwise.
func (o *TLSCommon) GetUseTLS() int32 {
	if o == nil || o.UseTLS == nil {
		var ret int32
		return ret
	}
	return *o.UseTLS
}

// GetUseTLSOk returns a tuple with the UseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCommon) GetUseTLSOk() (*int32, bool) {
	if o == nil || o.UseTLS == nil {
		return nil, false
	}
	return o.UseTLS, true
}

// HasUseTLS returns a boolean if a field has been set.
func (o *TLSCommon) HasUseTLS() bool {
	if o != nil && o.UseTLS != nil {
		return true
	}

	return false
}

// SetUseTLS gets a reference to the given int32 and assigns it to the UseTLS field.
func (o *TLSCommon) SetUseTLS(v int32) {
	o.UseTLS = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCommon) MarshalJSON() ([]byte, error) {
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
func (o *TLSCommon) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCommon := _TLSCommon{}

	if err = json.Unmarshal(bytes, &varTLSCommon); err == nil {
		*o = TLSCommon(varTLSCommon)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_ca_cert")
		delete(additionalProperties, "tls_client_cert")
		delete(additionalProperties, "tls_client_key")
		delete(additionalProperties, "tls_cert_hostname")
		delete(additionalProperties, "use_tls")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSCommon is a helper abstraction for handling nullable tlscommon types. 
type NullableTLSCommon struct {
	value *TLSCommon
	isSet bool
}

// Get returns the value.
func (v NullableTLSCommon) Get() *TLSCommon {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCommon) Set(val *TLSCommon) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCommon) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCommon) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCommon returns a pointer to a new instance of NullableTLSCommon.
func NewNullableTLSCommon(val *TLSCommon) *NullableTLSCommon {
	return &NullableTLSCommon{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
