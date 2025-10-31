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
	"time"
)

// TlsCertificateResponseAttributesAllOf struct for TlsCertificateResponseAttributesAllOf
type TlsCertificateResponseAttributesAllOf struct {
	// The hostname for which a certificate was issued.
	IssuedTo *string `json:"issued_to,omitempty"`
	// The certificate authority that issued the certificate.
	Issuer *string `json:"issuer,omitempty"`
	// A value assigned by the issuer that is unique to a certificate.
	SerialNumber *string `json:"serial_number,omitempty"`
	// The algorithm used to sign the certificate.
	SignatureAlgorithm *string `json:"signature_algorithm,omitempty"`
	// Time-stamp (GMT) when the certificate will expire. Must be in the future to be used to terminate TLS traffic.
	NotAfter *time.Time `json:"not_after,omitempty"`
	// Time-stamp (GMT) when the certificate will become valid. Must be in the past to be used to terminate TLS traffic.
	NotBefore *time.Time `json:"not_before,omitempty"`
	// A recommendation from Fastly indicating the key associated with this certificate is in need of rotation.
	Replace              *bool `json:"replace,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsCertificateResponseAttributesAllOf TlsCertificateResponseAttributesAllOf

// NewTlsCertificateResponseAttributesAllOf instantiates a new TlsCertificateResponseAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsCertificateResponseAttributesAllOf() *TlsCertificateResponseAttributesAllOf {
	this := TlsCertificateResponseAttributesAllOf{}
	return &this
}

// NewTlsCertificateResponseAttributesAllOfWithDefaults instantiates a new TlsCertificateResponseAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsCertificateResponseAttributesAllOfWithDefaults() *TlsCertificateResponseAttributesAllOf {
	this := TlsCertificateResponseAttributesAllOf{}
	return &this
}

// GetIssuedTo returns the IssuedTo field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributesAllOf) GetIssuedTo() string {
	if o == nil || o.IssuedTo == nil {
		var ret string
		return ret
	}
	return *o.IssuedTo
}

// GetIssuedToOk returns a tuple with the IssuedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributesAllOf) GetIssuedToOk() (*string, bool) {
	if o == nil || o.IssuedTo == nil {
		return nil, false
	}
	return o.IssuedTo, true
}

// HasIssuedTo returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributesAllOf) HasIssuedTo() bool {
	if o != nil && o.IssuedTo != nil {
		return true
	}

	return false
}

// SetIssuedTo gets a reference to the given string and assigns it to the IssuedTo field.
func (o *TlsCertificateResponseAttributesAllOf) SetIssuedTo(v string) {
	o.IssuedTo = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributesAllOf) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributesAllOf) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributesAllOf) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *TlsCertificateResponseAttributesAllOf) SetIssuer(v string) {
	o.Issuer = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributesAllOf) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributesAllOf) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributesAllOf) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *TlsCertificateResponseAttributesAllOf) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributesAllOf) GetSignatureAlgorithm() string {
	if o == nil || o.SignatureAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributesAllOf) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || o.SignatureAlgorithm == nil {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributesAllOf) HasSignatureAlgorithm() bool {
	if o != nil && o.SignatureAlgorithm != nil {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *TlsCertificateResponseAttributesAllOf) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributesAllOf) GetNotAfter() time.Time {
	if o == nil || o.NotAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributesAllOf) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributesAllOf) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *TlsCertificateResponseAttributesAllOf) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributesAllOf) GetNotBefore() time.Time {
	if o == nil || o.NotBefore == nil {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributesAllOf) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || o.NotBefore == nil {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributesAllOf) HasNotBefore() bool {
	if o != nil && o.NotBefore != nil {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *TlsCertificateResponseAttributesAllOf) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributesAllOf) GetReplace() bool {
	if o == nil || o.Replace == nil {
		var ret bool
		return ret
	}
	return *o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributesAllOf) GetReplaceOk() (*bool, bool) {
	if o == nil || o.Replace == nil {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributesAllOf) HasReplace() bool {
	if o != nil && o.Replace != nil {
		return true
	}

	return false
}

// SetReplace gets a reference to the given bool and assigns it to the Replace field.
func (o *TlsCertificateResponseAttributesAllOf) SetReplace(v bool) {
	o.Replace = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsCertificateResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.IssuedTo != nil {
		toSerialize["issued_to"] = o.IssuedTo
	}
	if o.Issuer != nil {
		toSerialize["issuer"] = o.Issuer
	}
	if o.SerialNumber != nil {
		toSerialize["serial_number"] = o.SerialNumber
	}
	if o.SignatureAlgorithm != nil {
		toSerialize["signature_algorithm"] = o.SignatureAlgorithm
	}
	if o.NotAfter != nil {
		toSerialize["not_after"] = o.NotAfter
	}
	if o.NotBefore != nil {
		toSerialize["not_before"] = o.NotBefore
	}
	if o.Replace != nil {
		toSerialize["replace"] = o.Replace
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsCertificateResponseAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTlsCertificateResponseAttributesAllOf := _TlsCertificateResponseAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varTlsCertificateResponseAttributesAllOf); err == nil {
		*o = TlsCertificateResponseAttributesAllOf(varTlsCertificateResponseAttributesAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "issued_to")
		delete(additionalProperties, "issuer")
		delete(additionalProperties, "serial_number")
		delete(additionalProperties, "signature_algorithm")
		delete(additionalProperties, "not_after")
		delete(additionalProperties, "not_before")
		delete(additionalProperties, "replace")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsCertificateResponseAttributesAllOf is a helper abstraction for handling nullable tlscertificateresponseattributesallof types.
type NullableTlsCertificateResponseAttributesAllOf struct {
	value *TlsCertificateResponseAttributesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTlsCertificateResponseAttributesAllOf) Get() *TlsCertificateResponseAttributesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsCertificateResponseAttributesAllOf) Set(val *TlsCertificateResponseAttributesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsCertificateResponseAttributesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsCertificateResponseAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsCertificateResponseAttributesAllOf returns a pointer to a new instance of NullableTlsCertificateResponseAttributesAllOf.
func NewNullableTlsCertificateResponseAttributesAllOf(val *TlsCertificateResponseAttributesAllOf) *NullableTlsCertificateResponseAttributesAllOf {
	return &NullableTlsCertificateResponseAttributesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsCertificateResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsCertificateResponseAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
