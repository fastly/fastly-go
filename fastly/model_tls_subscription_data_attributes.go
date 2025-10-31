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

// TlsSubscriptionDataAttributes struct for TlsSubscriptionDataAttributes
type TlsSubscriptionDataAttributes struct {
	// The entity that issues and certifies the TLS certificates for your subscription, either `certainly`, `lets-encrypt`, or `globalsign`. To migrate the subscription from one certificate authority to another, such as to migrate from 'lets-encrypt' to 'certainly',  pass `certificate_authority` to the PATCH endpoint. To migrate from 'globalsign' to 'certainly', contact Fastly Support.
	CertificateAuthority *string `json:"certificate_authority,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsSubscriptionDataAttributes TlsSubscriptionDataAttributes

// NewTlsSubscriptionDataAttributes instantiates a new TlsSubscriptionDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsSubscriptionDataAttributes() *TlsSubscriptionDataAttributes {
	this := TlsSubscriptionDataAttributes{}
	return &this
}

// NewTlsSubscriptionDataAttributesWithDefaults instantiates a new TlsSubscriptionDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsSubscriptionDataAttributesWithDefaults() *TlsSubscriptionDataAttributes {
	this := TlsSubscriptionDataAttributes{}
	return &this
}

// GetCertificateAuthority returns the CertificateAuthority field value if set, zero value otherwise.
func (o *TlsSubscriptionDataAttributes) GetCertificateAuthority() string {
	if o == nil || o.CertificateAuthority == nil {
		var ret string
		return ret
	}
	return *o.CertificateAuthority
}

// GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsSubscriptionDataAttributes) GetCertificateAuthorityOk() (*string, bool) {
	if o == nil || o.CertificateAuthority == nil {
		return nil, false
	}
	return o.CertificateAuthority, true
}

// HasCertificateAuthority returns a boolean if a field has been set.
func (o *TlsSubscriptionDataAttributes) HasCertificateAuthority() bool {
	if o != nil && o.CertificateAuthority != nil {
		return true
	}

	return false
}

// SetCertificateAuthority gets a reference to the given string and assigns it to the CertificateAuthority field.
func (o *TlsSubscriptionDataAttributes) SetCertificateAuthority(v string) {
	o.CertificateAuthority = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsSubscriptionDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CertificateAuthority != nil {
		toSerialize["certificate_authority"] = o.CertificateAuthority
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsSubscriptionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTlsSubscriptionDataAttributes := _TlsSubscriptionDataAttributes{}

	if err = json.Unmarshal(bytes, &varTlsSubscriptionDataAttributes); err == nil {
		*o = TlsSubscriptionDataAttributes(varTlsSubscriptionDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "certificate_authority")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsSubscriptionDataAttributes is a helper abstraction for handling nullable tlssubscriptiondataattributes types.
type NullableTlsSubscriptionDataAttributes struct {
	value *TlsSubscriptionDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTlsSubscriptionDataAttributes) Get() *TlsSubscriptionDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsSubscriptionDataAttributes) Set(val *TlsSubscriptionDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsSubscriptionDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsSubscriptionDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsSubscriptionDataAttributes returns a pointer to a new instance of NullableTlsSubscriptionDataAttributes.
func NewNullableTlsSubscriptionDataAttributes(val *TlsSubscriptionDataAttributes) *NullableTlsSubscriptionDataAttributes {
	return &NullableTlsSubscriptionDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsSubscriptionDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsSubscriptionDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
