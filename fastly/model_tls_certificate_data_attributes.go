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

// TlsCertificateDataAttributes struct for TlsCertificateDataAttributes
type TlsCertificateDataAttributes struct {
	// The PEM-formatted certificate blob. Required.
	CertBlob *string `json:"cert_blob,omitempty"`
	// A customizable name for your certificate. Defaults to the certificate's Common Name or first Subject Alternative Names (SAN) entry. Optional.
	Name *string `json:"name,omitempty"`
	// Indicates that the supplied certificate was not signed by a trusted CA.
	AllowUntrustedRoot   *bool `json:"allow_untrusted_root,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsCertificateDataAttributes TlsCertificateDataAttributes

// NewTlsCertificateDataAttributes instantiates a new TlsCertificateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsCertificateDataAttributes() *TlsCertificateDataAttributes {
	this := TlsCertificateDataAttributes{}
	return &this
}

// NewTlsCertificateDataAttributesWithDefaults instantiates a new TlsCertificateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsCertificateDataAttributesWithDefaults() *TlsCertificateDataAttributes {
	this := TlsCertificateDataAttributes{}
	return &this
}

// GetCertBlob returns the CertBlob field value if set, zero value otherwise.
func (o *TlsCertificateDataAttributes) GetCertBlob() string {
	if o == nil || o.CertBlob == nil {
		var ret string
		return ret
	}
	return *o.CertBlob
}

// GetCertBlobOk returns a tuple with the CertBlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateDataAttributes) GetCertBlobOk() (*string, bool) {
	if o == nil || o.CertBlob == nil {
		return nil, false
	}
	return o.CertBlob, true
}

// HasCertBlob returns a boolean if a field has been set.
func (o *TlsCertificateDataAttributes) HasCertBlob() bool {
	if o != nil && o.CertBlob != nil {
		return true
	}

	return false
}

// SetCertBlob gets a reference to the given string and assigns it to the CertBlob field.
func (o *TlsCertificateDataAttributes) SetCertBlob(v string) {
	o.CertBlob = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TlsCertificateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TlsCertificateDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TlsCertificateDataAttributes) SetName(v string) {
	o.Name = &v
}

// GetAllowUntrustedRoot returns the AllowUntrustedRoot field value if set, zero value otherwise.
func (o *TlsCertificateDataAttributes) GetAllowUntrustedRoot() bool {
	if o == nil || o.AllowUntrustedRoot == nil {
		var ret bool
		return ret
	}
	return *o.AllowUntrustedRoot
}

// GetAllowUntrustedRootOk returns a tuple with the AllowUntrustedRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateDataAttributes) GetAllowUntrustedRootOk() (*bool, bool) {
	if o == nil || o.AllowUntrustedRoot == nil {
		return nil, false
	}
	return o.AllowUntrustedRoot, true
}

// HasAllowUntrustedRoot returns a boolean if a field has been set.
func (o *TlsCertificateDataAttributes) HasAllowUntrustedRoot() bool {
	if o != nil && o.AllowUntrustedRoot != nil {
		return true
	}

	return false
}

// SetAllowUntrustedRoot gets a reference to the given bool and assigns it to the AllowUntrustedRoot field.
func (o *TlsCertificateDataAttributes) SetAllowUntrustedRoot(v bool) {
	o.AllowUntrustedRoot = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsCertificateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CertBlob != nil {
		toSerialize["cert_blob"] = o.CertBlob
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.AllowUntrustedRoot != nil {
		toSerialize["allow_untrusted_root"] = o.AllowUntrustedRoot
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsCertificateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTlsCertificateDataAttributes := _TlsCertificateDataAttributes{}

	if err = json.Unmarshal(bytes, &varTlsCertificateDataAttributes); err == nil {
		*o = TlsCertificateDataAttributes(varTlsCertificateDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cert_blob")
		delete(additionalProperties, "name")
		delete(additionalProperties, "allow_untrusted_root")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsCertificateDataAttributes is a helper abstraction for handling nullable tlscertificatedataattributes types.
type NullableTlsCertificateDataAttributes struct {
	value *TlsCertificateDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTlsCertificateDataAttributes) Get() *TlsCertificateDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsCertificateDataAttributes) Set(val *TlsCertificateDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsCertificateDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsCertificateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsCertificateDataAttributes returns a pointer to a new instance of NullableTlsCertificateDataAttributes.
func NewNullableTlsCertificateDataAttributes(val *TlsCertificateDataAttributes) *NullableTlsCertificateDataAttributes {
	return &NullableTlsCertificateDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsCertificateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsCertificateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
