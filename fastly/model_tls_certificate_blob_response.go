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

// TlsCertificateBlobResponse struct for TlsCertificateBlobResponse
type TlsCertificateBlobResponse struct {
	// A certificate blob
	CertBlob             *string `json:"cert_blob,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsCertificateBlobResponse TlsCertificateBlobResponse

// NewTlsCertificateBlobResponse instantiates a new TlsCertificateBlobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsCertificateBlobResponse() *TlsCertificateBlobResponse {
	this := TlsCertificateBlobResponse{}
	return &this
}

// NewTlsCertificateBlobResponseWithDefaults instantiates a new TlsCertificateBlobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsCertificateBlobResponseWithDefaults() *TlsCertificateBlobResponse {
	this := TlsCertificateBlobResponse{}
	return &this
}

// GetCertBlob returns the CertBlob field value if set, zero value otherwise.
func (o *TlsCertificateBlobResponse) GetCertBlob() string {
	if o == nil || o.CertBlob == nil {
		var ret string
		return ret
	}
	return *o.CertBlob
}

// GetCertBlobOk returns a tuple with the CertBlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateBlobResponse) GetCertBlobOk() (*string, bool) {
	if o == nil || o.CertBlob == nil {
		return nil, false
	}
	return o.CertBlob, true
}

// HasCertBlob returns a boolean if a field has been set.
func (o *TlsCertificateBlobResponse) HasCertBlob() bool {
	if o != nil && o.CertBlob != nil {
		return true
	}

	return false
}

// SetCertBlob gets a reference to the given string and assigns it to the CertBlob field.
func (o *TlsCertificateBlobResponse) SetCertBlob(v string) {
	o.CertBlob = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsCertificateBlobResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CertBlob != nil {
		toSerialize["cert_blob"] = o.CertBlob
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsCertificateBlobResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTlsCertificateBlobResponse := _TlsCertificateBlobResponse{}

	if err = json.Unmarshal(bytes, &varTlsCertificateBlobResponse); err == nil {
		*o = TlsCertificateBlobResponse(varTlsCertificateBlobResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cert_blob")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsCertificateBlobResponse is a helper abstraction for handling nullable tlscertificateblobresponse types.
type NullableTlsCertificateBlobResponse struct {
	value *TlsCertificateBlobResponse
	isSet bool
}

// Get returns the value.
func (v NullableTlsCertificateBlobResponse) Get() *TlsCertificateBlobResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsCertificateBlobResponse) Set(val *TlsCertificateBlobResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsCertificateBlobResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsCertificateBlobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsCertificateBlobResponse returns a pointer to a new instance of NullableTlsCertificateBlobResponse.
func NewNullableTlsCertificateBlobResponse(val *TlsCertificateBlobResponse) *NullableTlsCertificateBlobResponse {
	return &NullableTlsCertificateBlobResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsCertificateBlobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsCertificateBlobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
