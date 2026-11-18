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

// TLSCertificateBlobResponse struct for TLSCertificateBlobResponse
type TLSCertificateBlobResponse struct {
	// A certificate blob
	CertBlob             *string `json:"cert_blob,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCertificateBlobResponse TLSCertificateBlobResponse

// NewTLSCertificateBlobResponse instantiates a new TLSCertificateBlobResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCertificateBlobResponse() *TLSCertificateBlobResponse {
	this := TLSCertificateBlobResponse{}
	return &this
}

// NewTLSCertificateBlobResponseWithDefaults instantiates a new TLSCertificateBlobResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCertificateBlobResponseWithDefaults() *TLSCertificateBlobResponse {
	this := TLSCertificateBlobResponse{}
	return &this
}

// GetCertBlob returns the CertBlob field value if set, zero value otherwise.
func (o *TLSCertificateBlobResponse) GetCertBlob() string {
	if o == nil || o.CertBlob == nil {
		var ret string
		return ret
	}
	return *o.CertBlob
}

// GetCertBlobOk returns a tuple with the CertBlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCertificateBlobResponse) GetCertBlobOk() (*string, bool) {
	if o == nil || o.CertBlob == nil {
		return nil, false
	}
	return o.CertBlob, true
}

// HasCertBlob returns a boolean if a field has been set.
func (o *TLSCertificateBlobResponse) HasCertBlob() bool {
	if o != nil && o.CertBlob != nil {
		return true
	}

	return false
}

// SetCertBlob gets a reference to the given string and assigns it to the CertBlob field.
func (o *TLSCertificateBlobResponse) SetCertBlob(v string) {
	o.CertBlob = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCertificateBlobResponse) MarshalJSON() ([]byte, error) {
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
func (o *TLSCertificateBlobResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCertificateBlobResponse := _TLSCertificateBlobResponse{}

	if err = json.Unmarshal(bytes, &varTLSCertificateBlobResponse); err == nil {
		*o = TLSCertificateBlobResponse(varTLSCertificateBlobResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cert_blob")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSCertificateBlobResponse is a helper abstraction for handling nullable tlscertificateblobresponse types.
type NullableTLSCertificateBlobResponse struct {
	value *TLSCertificateBlobResponse
	isSet bool
}

// Get returns the value.
func (v NullableTLSCertificateBlobResponse) Get() *TLSCertificateBlobResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCertificateBlobResponse) Set(val *TLSCertificateBlobResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCertificateBlobResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCertificateBlobResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCertificateBlobResponse returns a pointer to a new instance of NullableTLSCertificateBlobResponse.
func NewNullableTLSCertificateBlobResponse(val *TLSCertificateBlobResponse) *NullableTLSCertificateBlobResponse {
	return &NullableTLSCertificateBlobResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCertificateBlobResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTLSCertificateBlobResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
