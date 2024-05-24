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

// TLSCertificateDataAttributes struct for TLSCertificateDataAttributes
type TLSCertificateDataAttributes struct {
	// The PEM-formatted certificate blob. Required.
	CertBlob *string `json:"cert_blob,omitempty"`
	// A customizable name for your certificate. Defaults to the certificate's Common Name or first Subject Alternative Names (SAN) entry. Optional.
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCertificateDataAttributes TLSCertificateDataAttributes

// NewTLSCertificateDataAttributes instantiates a new TLSCertificateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCertificateDataAttributes() *TLSCertificateDataAttributes {
	this := TLSCertificateDataAttributes{}
	return &this
}

// NewTLSCertificateDataAttributesWithDefaults instantiates a new TLSCertificateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCertificateDataAttributesWithDefaults() *TLSCertificateDataAttributes {
	this := TLSCertificateDataAttributes{}
	return &this
}

// GetCertBlob returns the CertBlob field value if set, zero value otherwise.
func (o *TLSCertificateDataAttributes) GetCertBlob() string {
	if o == nil || o.CertBlob == nil {
		var ret string
		return ret
	}
	return *o.CertBlob
}

// GetCertBlobOk returns a tuple with the CertBlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCertificateDataAttributes) GetCertBlobOk() (*string, bool) {
	if o == nil || o.CertBlob == nil {
		return nil, false
	}
	return o.CertBlob, true
}

// HasCertBlob returns a boolean if a field has been set.
func (o *TLSCertificateDataAttributes) HasCertBlob() bool {
	if o != nil && o.CertBlob != nil {
		return true
	}

	return false
}

// SetCertBlob gets a reference to the given string and assigns it to the CertBlob field.
func (o *TLSCertificateDataAttributes) SetCertBlob(v string) {
	o.CertBlob = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TLSCertificateDataAttributes) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCertificateDataAttributes) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TLSCertificateDataAttributes) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TLSCertificateDataAttributes) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCertificateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CertBlob != nil {
		toSerialize["cert_blob"] = o.CertBlob
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *TLSCertificateDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCertificateDataAttributes := _TLSCertificateDataAttributes{}

	if err = json.Unmarshal(bytes, &varTLSCertificateDataAttributes); err == nil {
		*o = TLSCertificateDataAttributes(varTLSCertificateDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cert_blob")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSCertificateDataAttributes is a helper abstraction for handling nullable tlscertificatedataattributes types. 
type NullableTLSCertificateDataAttributes struct {
	value *TLSCertificateDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTLSCertificateDataAttributes) Get() *TLSCertificateDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCertificateDataAttributes) Set(val *TLSCertificateDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCertificateDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCertificateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCertificateDataAttributes returns a pointer to a new instance of NullableTLSCertificateDataAttributes.
func NewNullableTLSCertificateDataAttributes(val *TLSCertificateDataAttributes) *NullableTLSCertificateDataAttributes {
	return &NullableTLSCertificateDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCertificateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSCertificateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
