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

// LoggingTlsCommon struct for LoggingTlsCommon
type LoggingTlsCommon struct {
	// A secure certificate to authenticate a server with. Must be in PEM format.
	TlsCaCert NullableString `json:"tls_ca_cert,omitempty"`
	// The client certificate used to make authenticated requests. Must be in PEM format.
	TlsClientCert NullableString `json:"tls_client_cert,omitempty"`
	// The client private key used to make authenticated requests. Must be in PEM format.
	TlsClientKey NullableString `json:"tls_client_key,omitempty"`
	// The hostname to verify the server's certificate. This should be one of the Subject Alternative Name (SAN) fields for the certificate. Common Names (CN) are not supported.
	TlsHostname          NullableString `json:"tls_hostname,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingTlsCommon LoggingTlsCommon

// NewLoggingTlsCommon instantiates a new LoggingTlsCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingTlsCommon() *LoggingTlsCommon {
	this := LoggingTlsCommon{}
	var tlsCaCert string = "null"
	this.TlsCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TlsClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TlsClientKey = *NewNullableString(&tlsClientKey)
	var tlsHostname string = "null"
	this.TlsHostname = *NewNullableString(&tlsHostname)
	return &this
}

// NewLoggingTlsCommonWithDefaults instantiates a new LoggingTlsCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingTlsCommonWithDefaults() *LoggingTlsCommon {
	this := LoggingTlsCommon{}
	var tlsCaCert string = "null"
	this.TlsCaCert = *NewNullableString(&tlsCaCert)
	var tlsClientCert string = "null"
	this.TlsClientCert = *NewNullableString(&tlsClientCert)
	var tlsClientKey string = "null"
	this.TlsClientKey = *NewNullableString(&tlsClientKey)
	var tlsHostname string = "null"
	this.TlsHostname = *NewNullableString(&tlsHostname)
	return &this
}

// GetTlsCaCert returns the TlsCaCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingTlsCommon) GetTlsCaCert() string {
	if o == nil || o.TlsCaCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsCaCert.Get()
}

// GetTlsCaCertOk returns a tuple with the TlsCaCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingTlsCommon) GetTlsCaCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsCaCert.Get(), o.TlsCaCert.IsSet()
}

// HasTlsCaCert returns a boolean if a field has been set.
func (o *LoggingTlsCommon) HasTlsCaCert() bool {
	if o != nil && o.TlsCaCert.IsSet() {
		return true
	}

	return false
}

// SetTlsCaCert gets a reference to the given NullableString and assigns it to the TlsCaCert field.
func (o *LoggingTlsCommon) SetTlsCaCert(v string) {
	o.TlsCaCert.Set(&v)
}

// SetTlsCaCertNil sets the value for TlsCaCert to be an explicit nil
func (o *LoggingTlsCommon) SetTlsCaCertNil() {
	o.TlsCaCert.Set(nil)
}

// UnsetTlsCaCert ensures that no value is present for TlsCaCert, not even an explicit nil
func (o *LoggingTlsCommon) UnsetTlsCaCert() {
	o.TlsCaCert.Unset()
}

// GetTlsClientCert returns the TlsClientCert field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingTlsCommon) GetTlsClientCert() string {
	if o == nil || o.TlsClientCert.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsClientCert.Get()
}

// GetTlsClientCertOk returns a tuple with the TlsClientCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingTlsCommon) GetTlsClientCertOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsClientCert.Get(), o.TlsClientCert.IsSet()
}

// HasTlsClientCert returns a boolean if a field has been set.
func (o *LoggingTlsCommon) HasTlsClientCert() bool {
	if o != nil && o.TlsClientCert.IsSet() {
		return true
	}

	return false
}

// SetTlsClientCert gets a reference to the given NullableString and assigns it to the TlsClientCert field.
func (o *LoggingTlsCommon) SetTlsClientCert(v string) {
	o.TlsClientCert.Set(&v)
}

// SetTlsClientCertNil sets the value for TlsClientCert to be an explicit nil
func (o *LoggingTlsCommon) SetTlsClientCertNil() {
	o.TlsClientCert.Set(nil)
}

// UnsetTlsClientCert ensures that no value is present for TlsClientCert, not even an explicit nil
func (o *LoggingTlsCommon) UnsetTlsClientCert() {
	o.TlsClientCert.Unset()
}

// GetTlsClientKey returns the TlsClientKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingTlsCommon) GetTlsClientKey() string {
	if o == nil || o.TlsClientKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsClientKey.Get()
}

// GetTlsClientKeyOk returns a tuple with the TlsClientKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingTlsCommon) GetTlsClientKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsClientKey.Get(), o.TlsClientKey.IsSet()
}

// HasTlsClientKey returns a boolean if a field has been set.
func (o *LoggingTlsCommon) HasTlsClientKey() bool {
	if o != nil && o.TlsClientKey.IsSet() {
		return true
	}

	return false
}

// SetTlsClientKey gets a reference to the given NullableString and assigns it to the TlsClientKey field.
func (o *LoggingTlsCommon) SetTlsClientKey(v string) {
	o.TlsClientKey.Set(&v)
}

// SetTlsClientKeyNil sets the value for TlsClientKey to be an explicit nil
func (o *LoggingTlsCommon) SetTlsClientKeyNil() {
	o.TlsClientKey.Set(nil)
}

// UnsetTlsClientKey ensures that no value is present for TlsClientKey, not even an explicit nil
func (o *LoggingTlsCommon) UnsetTlsClientKey() {
	o.TlsClientKey.Unset()
}

// GetTlsHostname returns the TlsHostname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingTlsCommon) GetTlsHostname() string {
	if o == nil || o.TlsHostname.Get() == nil {
		var ret string
		return ret
	}
	return *o.TlsHostname.Get()
}

// GetTlsHostnameOk returns a tuple with the TlsHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingTlsCommon) GetTlsHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TlsHostname.Get(), o.TlsHostname.IsSet()
}

// HasTlsHostname returns a boolean if a field has been set.
func (o *LoggingTlsCommon) HasTlsHostname() bool {
	if o != nil && o.TlsHostname.IsSet() {
		return true
	}

	return false
}

// SetTlsHostname gets a reference to the given NullableString and assigns it to the TlsHostname field.
func (o *LoggingTlsCommon) SetTlsHostname(v string) {
	o.TlsHostname.Set(&v)
}

// SetTlsHostnameNil sets the value for TlsHostname to be an explicit nil
func (o *LoggingTlsCommon) SetTlsHostnameNil() {
	o.TlsHostname.Set(nil)
}

// UnsetTlsHostname ensures that no value is present for TlsHostname, not even an explicit nil
func (o *LoggingTlsCommon) UnsetTlsHostname() {
	o.TlsHostname.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingTlsCommon) MarshalJSON() ([]byte, error) {
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
	if o.TlsHostname.IsSet() {
		toSerialize["tls_hostname"] = o.TlsHostname.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LoggingTlsCommon) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingTlsCommon := _LoggingTlsCommon{}

	if err = json.Unmarshal(bytes, &varLoggingTlsCommon); err == nil {
		*o = LoggingTlsCommon(varLoggingTlsCommon)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_ca_cert")
		delete(additionalProperties, "tls_client_cert")
		delete(additionalProperties, "tls_client_key")
		delete(additionalProperties, "tls_hostname")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingTlsCommon is a helper abstraction for handling nullable loggingtlscommon types.
type NullableLoggingTlsCommon struct {
	value *LoggingTlsCommon
	isSet bool
}

// Get returns the value.
func (v NullableLoggingTlsCommon) Get() *LoggingTlsCommon {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingTlsCommon) Set(val *LoggingTlsCommon) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingTlsCommon) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingTlsCommon) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingTlsCommon returns a pointer to a new instance of NullableLoggingTlsCommon.
func NewNullableLoggingTlsCommon(val *LoggingTlsCommon) *NullableLoggingTlsCommon {
	return &NullableLoggingTlsCommon{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingTlsCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLoggingTlsCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
