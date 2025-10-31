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

// TlsCommonResponse struct for TlsCommonResponse
type TlsCommonResponse struct {
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TlsCaCert NullableString `json:"tls_ca_cert,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TlsClientCert NullableString `json:"tls_client_cert,omitempty"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TlsClientKey NullableString `json:"tls_client_key,omitempty"`
	// The hostname used to verify a server's certificate. It can either be the Common Name (CN) or a Subject Alternative Name (SAN).
	TlsCertHostname NullableString `json:"tls_cert_hostname,omitempty"`
	// Whether to use TLS.
	UseTls               *string `json:"use_tls,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsCommonResponse TlsCommonResponse

// NewTlsCommonResponse instantiates a new TlsCommonResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsCommonResponse() *TlsCommonResponse {
	this := TlsCommonResponse{}
	var tlsCaCert string = "null"
	this.TlsCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TlsClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TlsClientKey = *NewNullableString(&tlsClientKey)
	var tlsCertHostname string = "null"
	this.TlsCertHostname = *NewNullableString(&tlsCertHostname)
	var useTls string = "0"
	this.UseTls = &useTls
	return &this
}

// NewTlsCommonResponseWithDefaults instantiates a new TlsCommonResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsCommonResponseWithDefaults() *TlsCommonResponse {
	this := TlsCommonResponse{}
	var tlsCaCert string = "null"
	this.TlsCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TlsClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TlsClientKey = *NewNullableString(&tlsClientKey)
	var tlsCertHostname string = "null"
	this.TlsCertHostname = *NewNullableString(&tlsCertHostname)
	var useTls string = "0"
	this.UseTls = &useTls
	return &this
}

// GetTlsCaCert returns the TlsCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsCommonResponse) GetTlsCaCert() string {
	if o == nil || o.TlsCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsCaCert.Get()
}

// GetTlsCaCertOk returns a tuple with the TlsCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsCommonResponse) GetTlsCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCaCert.Get(), o.TlsCaCert.IsSet()
}

// HasTlsCaCert returns a boolean if a field has been set.
func (o *TlsCommonResponse) HasTlsCaCert() bool {
	if o != nil && o.TlsCaCert.IsSet() {
		return true
	}

	return false
}

// SetTlsCaCert gets a reference to the given NullableString and assigns it to the TlsCaCert field.
func (o *TlsCommonResponse) SetTlsCaCert(v string) {
	o.TlsCaCert.Set(&v)
}

// SetTlsCaCertNil sets the value for TlsCaCert to be an explicit nil
func (o *TlsCommonResponse) SetTlsCaCertNil() {
	o.TlsCaCert.Set(nil)
}

// UnsetTlsCaCert ensures that no value is present for TlsCaCert, not even an explicit nil
func (o *TlsCommonResponse) UnsetTlsCaCert() {
	o.TlsCaCert.Unset()
}

// GetTlsClientCert returns the TlsClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsCommonResponse) GetTlsClientCert() string {
	if o == nil || o.TlsClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsClientCert.Get()
}

// GetTlsClientCertOk returns a tuple with the TlsClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsCommonResponse) GetTlsClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsClientCert.Get(), o.TlsClientCert.IsSet()
}

// HasTlsClientCert returns a boolean if a field has been set.
func (o *TlsCommonResponse) HasTlsClientCert() bool {
	if o != nil && o.TlsClientCert.IsSet() {
		return true
	}

	return false
}

// SetTlsClientCert gets a reference to the given NullableString and assigns it to the TlsClientCert field.
func (o *TlsCommonResponse) SetTlsClientCert(v string) {
	o.TlsClientCert.Set(&v)
}

// SetTlsClientCertNil sets the value for TlsClientCert to be an explicit nil
func (o *TlsCommonResponse) SetTlsClientCertNil() {
	o.TlsClientCert.Set(nil)
}

// UnsetTlsClientCert ensures that no value is present for TlsClientCert, not even an explicit nil
func (o *TlsCommonResponse) UnsetTlsClientCert() {
	o.TlsClientCert.Unset()
}

// GetTlsClientKey returns the TlsClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsCommonResponse) GetTlsClientKey() string {
	if o == nil || o.TlsClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsClientKey.Get()
}

// GetTlsClientKeyOk returns a tuple with the TlsClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsCommonResponse) GetTlsClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsClientKey.Get(), o.TlsClientKey.IsSet()
}

// HasTlsClientKey returns a boolean if a field has been set.
func (o *TlsCommonResponse) HasTlsClientKey() bool {
	if o != nil && o.TlsClientKey.IsSet() {
		return true
	}

	return false
}

// SetTlsClientKey gets a reference to the given NullableString and assigns it to the TlsClientKey field.
func (o *TlsCommonResponse) SetTlsClientKey(v string) {
	o.TlsClientKey.Set(&v)
}

// SetTlsClientKeyNil sets the value for TlsClientKey to be an explicit nil
func (o *TlsCommonResponse) SetTlsClientKeyNil() {
	o.TlsClientKey.Set(nil)
}

// UnsetTlsClientKey ensures that no value is present for TlsClientKey, not even an explicit nil
func (o *TlsCommonResponse) UnsetTlsClientKey() {
	o.TlsClientKey.Unset()
}

// GetTlsCertHostname returns the TlsCertHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsCommonResponse) GetTlsCertHostname() string {
	if o == nil || o.TlsCertHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsCertHostname.Get()
}

// GetTlsCertHostnameOk returns a tuple with the TlsCertHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsCommonResponse) GetTlsCertHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCertHostname.Get(), o.TlsCertHostname.IsSet()
}

// HasTlsCertHostname returns a boolean if a field has been set.
func (o *TlsCommonResponse) HasTlsCertHostname() bool {
	if o != nil && o.TlsCertHostname.IsSet() {
		return true
	}

	return false
}

// SetTlsCertHostname gets a reference to the given NullableString and assigns it to the TlsCertHostname field.
func (o *TlsCommonResponse) SetTlsCertHostname(v string) {
	o.TlsCertHostname.Set(&v)
}

// SetTlsCertHostnameNil sets the value for TlsCertHostname to be an explicit nil
func (o *TlsCommonResponse) SetTlsCertHostnameNil() {
	o.TlsCertHostname.Set(nil)
}

// UnsetTlsCertHostname ensures that no value is present for TlsCertHostname, not even an explicit nil
func (o *TlsCommonResponse) UnsetTlsCertHostname() {
	o.TlsCertHostname.Unset()
}

// GetUseTls returns the UseTls field value if set, zero value otherwise.
func (o *TlsCommonResponse) GetUseTls() string {
	if o == nil || o.UseTls == nil {
		var ret string
		return ret
	}
	return *o.UseTls
}

// GetUseTlsOk returns a tuple with the UseTls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCommonResponse) GetUseTlsOk() (*string, bool) {
	if o == nil || o.UseTls == nil {
		return nil, false
	}
	return o.UseTls, true
}

// HasUseTls returns a boolean if a field has been set.
func (o *TlsCommonResponse) HasUseTls() bool {
	if o != nil && o.UseTls != nil {
		return true
	}

	return false
}

// SetUseTls gets a reference to the given string and assigns it to the UseTls field.
func (o *TlsCommonResponse) SetUseTls(v string) {
	o.UseTls = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsCommonResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsCaCert.IsSet() {
		toSerialize["tls_ca_cert"] = o.TlsCaCert.Get()
	}
	if o.TlsClientCert.IsSet() {
		toSerialize["tls_client_cert"] = o.TlsClientCert.Get()
	}
	if o.TlsClientKey.IsSet() {
		toSerialize["tls_client_key"] = o.TlsClientKey.Get()
	}
	if o.TlsCertHostname.IsSet() {
		toSerialize["tls_cert_hostname"] = o.TlsCertHostname.Get()
	}
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
func (o *TlsCommonResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTlsCommonResponse := _TlsCommonResponse{}

	if err = json.Unmarshal(bytes, &varTlsCommonResponse); err == nil {
		*o = TlsCommonResponse(varTlsCommonResponse)
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

// NullableTlsCommonResponse is a helper abstraction for handling nullable tlscommonresponse types.
type NullableTlsCommonResponse struct {
	value *TlsCommonResponse
	isSet bool
}

// Get returns the value.
func (v NullableTlsCommonResponse) Get() *TlsCommonResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsCommonResponse) Set(val *TlsCommonResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsCommonResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsCommonResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsCommonResponse returns a pointer to a new instance of NullableTlsCommonResponse.
func NewNullableTlsCommonResponse(val *TlsCommonResponse) *NullableTlsCommonResponse {
	return &NullableTlsCommonResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsCommonResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsCommonResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
