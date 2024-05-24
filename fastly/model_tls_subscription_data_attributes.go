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

// TLSSubscriptionDataAttributes struct for TLSSubscriptionDataAttributes
type TLSSubscriptionDataAttributes struct {
	// The entity that issues and certifies the TLS certificates for your subscription, either `certainly`, `lets-encrypt`, or `globalsign`. To migrate the subscription from one certificate authority to another, such as to migrate from 'lets-encrypt' to 'certainly',  pass `certificate_authority` to the PATCH endpoint. To migrate from 'globalsign' to 'certainly', contact Fastly Support.
	CertificateAuthority *string `json:"certificate_authority,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSSubscriptionDataAttributes TLSSubscriptionDataAttributes

// NewTLSSubscriptionDataAttributes instantiates a new TLSSubscriptionDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSSubscriptionDataAttributes() *TLSSubscriptionDataAttributes {
	this := TLSSubscriptionDataAttributes{}
	return &this
}

// NewTLSSubscriptionDataAttributesWithDefaults instantiates a new TLSSubscriptionDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSSubscriptionDataAttributesWithDefaults() *TLSSubscriptionDataAttributes {
	this := TLSSubscriptionDataAttributes{}
	return &this
}

// GetCertificateAuthority returns the CertificateAuthority field value if set, zero value otherwise.
func (o *TLSSubscriptionDataAttributes) GetCertificateAuthority() string {
	if o == nil || o.CertificateAuthority == nil {
		var ret string
		return ret
	}
	return *o.CertificateAuthority
}

// GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSSubscriptionDataAttributes) GetCertificateAuthorityOk() (*string, bool) {
	if o == nil || o.CertificateAuthority == nil {
		return nil, false
	}
	return o.CertificateAuthority, true
}

// HasCertificateAuthority returns a boolean if a field has been set.
func (o *TLSSubscriptionDataAttributes) HasCertificateAuthority() bool {
	if o != nil && o.CertificateAuthority != nil {
		return true
	}

	return false
}

// SetCertificateAuthority gets a reference to the given string and assigns it to the CertificateAuthority field.
func (o *TLSSubscriptionDataAttributes) SetCertificateAuthority(v string) {
	o.CertificateAuthority = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSSubscriptionDataAttributes) MarshalJSON() ([]byte, error) {
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
func (o *TLSSubscriptionDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTLSSubscriptionDataAttributes := _TLSSubscriptionDataAttributes{}

	if err = json.Unmarshal(bytes, &varTLSSubscriptionDataAttributes); err == nil {
		*o = TLSSubscriptionDataAttributes(varTLSSubscriptionDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "certificate_authority")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSSubscriptionDataAttributes is a helper abstraction for handling nullable tlssubscriptiondataattributes types. 
type NullableTLSSubscriptionDataAttributes struct {
	value *TLSSubscriptionDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTLSSubscriptionDataAttributes) Get() *TLSSubscriptionDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSSubscriptionDataAttributes) Set(val *TLSSubscriptionDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSSubscriptionDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSSubscriptionDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSSubscriptionDataAttributes returns a pointer to a new instance of NullableTLSSubscriptionDataAttributes.
func NewNullableTLSSubscriptionDataAttributes(val *TLSSubscriptionDataAttributes) *NullableTLSSubscriptionDataAttributes {
	return &NullableTLSSubscriptionDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSSubscriptionDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSSubscriptionDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
