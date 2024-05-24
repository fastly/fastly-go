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

// TLSBulkCertificateDataAttributes struct for TLSBulkCertificateDataAttributes
type TLSBulkCertificateDataAttributes struct {
	// Allow certificates that chain to untrusted roots.
	AllowUntrustedRoot *bool `json:"allow_untrusted_root,omitempty"`
	// The PEM-formatted certificate blob. Required.
	CertBlob *string `json:"cert_blob,omitempty"`
	// The PEM-formatted chain of intermediate blobs. Required.
	IntermediatesBlob *string `json:"intermediates_blob,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSBulkCertificateDataAttributes TLSBulkCertificateDataAttributes

// NewTLSBulkCertificateDataAttributes instantiates a new TLSBulkCertificateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSBulkCertificateDataAttributes() *TLSBulkCertificateDataAttributes {
	this := TLSBulkCertificateDataAttributes{}
	var allowUntrustedRoot bool = false
	this.AllowUntrustedRoot = &allowUntrustedRoot
	return &this
}

// NewTLSBulkCertificateDataAttributesWithDefaults instantiates a new TLSBulkCertificateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSBulkCertificateDataAttributesWithDefaults() *TLSBulkCertificateDataAttributes {
	this := TLSBulkCertificateDataAttributes{}
	var allowUntrustedRoot bool = false
	this.AllowUntrustedRoot = &allowUntrustedRoot
	return &this
}

// GetAllowUntrustedRoot returns the AllowUntrustedRoot field value if set, zero value otherwise.
func (o *TLSBulkCertificateDataAttributes) GetAllowUntrustedRoot() bool {
	if o == nil || o.AllowUntrustedRoot == nil {
		var ret bool
		return ret
	}
	return *o.AllowUntrustedRoot
}

// GetAllowUntrustedRootOk returns a tuple with the AllowUntrustedRoot field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSBulkCertificateDataAttributes) GetAllowUntrustedRootOk() (*bool, bool) {
	if o == nil || o.AllowUntrustedRoot == nil {
		return nil, false
	}
	return o.AllowUntrustedRoot, true
}

// HasAllowUntrustedRoot returns a boolean if a field has been set.
func (o *TLSBulkCertificateDataAttributes) HasAllowUntrustedRoot() bool {
	if o != nil && o.AllowUntrustedRoot != nil {
		return true
	}

	return false
}

// SetAllowUntrustedRoot gets a reference to the given bool and assigns it to the AllowUntrustedRoot field.
func (o *TLSBulkCertificateDataAttributes) SetAllowUntrustedRoot(v bool) {
	o.AllowUntrustedRoot = &v
}

// GetCertBlob returns the CertBlob field value if set, zero value otherwise.
func (o *TLSBulkCertificateDataAttributes) GetCertBlob() string {
	if o == nil || o.CertBlob == nil {
		var ret string
		return ret
	}
	return *o.CertBlob
}

// GetCertBlobOk returns a tuple with the CertBlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSBulkCertificateDataAttributes) GetCertBlobOk() (*string, bool) {
	if o == nil || o.CertBlob == nil {
		return nil, false
	}
	return o.CertBlob, true
}

// HasCertBlob returns a boolean if a field has been set.
func (o *TLSBulkCertificateDataAttributes) HasCertBlob() bool {
	if o != nil && o.CertBlob != nil {
		return true
	}

	return false
}

// SetCertBlob gets a reference to the given string and assigns it to the CertBlob field.
func (o *TLSBulkCertificateDataAttributes) SetCertBlob(v string) {
	o.CertBlob = &v
}

// GetIntermediatesBlob returns the IntermediatesBlob field value if set, zero value otherwise.
func (o *TLSBulkCertificateDataAttributes) GetIntermediatesBlob() string {
	if o == nil || o.IntermediatesBlob == nil {
		var ret string
		return ret
	}
	return *o.IntermediatesBlob
}

// GetIntermediatesBlobOk returns a tuple with the IntermediatesBlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSBulkCertificateDataAttributes) GetIntermediatesBlobOk() (*string, bool) {
	if o == nil || o.IntermediatesBlob == nil {
		return nil, false
	}
	return o.IntermediatesBlob, true
}

// HasIntermediatesBlob returns a boolean if a field has been set.
func (o *TLSBulkCertificateDataAttributes) HasIntermediatesBlob() bool {
	if o != nil && o.IntermediatesBlob != nil {
		return true
	}

	return false
}

// SetIntermediatesBlob gets a reference to the given string and assigns it to the IntermediatesBlob field.
func (o *TLSBulkCertificateDataAttributes) SetIntermediatesBlob(v string) {
	o.IntermediatesBlob = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSBulkCertificateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AllowUntrustedRoot != nil {
		toSerialize["allow_untrusted_root"] = o.AllowUntrustedRoot
	}
	if o.CertBlob != nil {
		toSerialize["cert_blob"] = o.CertBlob
	}
	if o.IntermediatesBlob != nil {
		toSerialize["intermediates_blob"] = o.IntermediatesBlob
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *TLSBulkCertificateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTLSBulkCertificateDataAttributes := _TLSBulkCertificateDataAttributes{}

	if err = json.Unmarshal(bytes, &varTLSBulkCertificateDataAttributes); err == nil {
		*o = TLSBulkCertificateDataAttributes(varTLSBulkCertificateDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "allow_untrusted_root")
		delete(additionalProperties, "cert_blob")
		delete(additionalProperties, "intermediates_blob")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSBulkCertificateDataAttributes is a helper abstraction for handling nullable tlsbulkcertificatedataattributes types. 
type NullableTLSBulkCertificateDataAttributes struct {
	value *TLSBulkCertificateDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTLSBulkCertificateDataAttributes) Get() *TLSBulkCertificateDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSBulkCertificateDataAttributes) Set(val *TLSBulkCertificateDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSBulkCertificateDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSBulkCertificateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSBulkCertificateDataAttributes returns a pointer to a new instance of NullableTLSBulkCertificateDataAttributes.
func NewNullableTLSBulkCertificateDataAttributes(val *TLSBulkCertificateDataAttributes) *NullableTLSBulkCertificateDataAttributes {
	return &NullableTLSBulkCertificateDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSBulkCertificateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSBulkCertificateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
