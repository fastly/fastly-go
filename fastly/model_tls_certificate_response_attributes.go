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

// TlsCertificateResponseAttributes struct for TlsCertificateResponseAttributes
type TlsCertificateResponseAttributes struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
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

type _TlsCertificateResponseAttributes TlsCertificateResponseAttributes

// NewTlsCertificateResponseAttributes instantiates a new TlsCertificateResponseAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsCertificateResponseAttributes() *TlsCertificateResponseAttributes {
	this := TlsCertificateResponseAttributes{}
	return &this
}

// NewTlsCertificateResponseAttributesWithDefaults instantiates a new TlsCertificateResponseAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsCertificateResponseAttributesWithDefaults() *TlsCertificateResponseAttributes {
	this := TlsCertificateResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsCertificateResponseAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsCertificateResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *TlsCertificateResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *TlsCertificateResponseAttributes) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *TlsCertificateResponseAttributes) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsCertificateResponseAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsCertificateResponseAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributes) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *TlsCertificateResponseAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *TlsCertificateResponseAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *TlsCertificateResponseAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsCertificateResponseAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsCertificateResponseAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *TlsCertificateResponseAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *TlsCertificateResponseAttributes) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *TlsCertificateResponseAttributes) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetIssuedTo returns the IssuedTo field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributes) GetIssuedTo() string {
	if o == nil || o.IssuedTo == nil {
		var ret string
		return ret
	}
	return *o.IssuedTo
}

// GetIssuedToOk returns a tuple with the IssuedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributes) GetIssuedToOk() (*string, bool) {
	if o == nil || o.IssuedTo == nil {
		return nil, false
	}
	return o.IssuedTo, true
}

// HasIssuedTo returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributes) HasIssuedTo() bool {
	if o != nil && o.IssuedTo != nil {
		return true
	}

	return false
}

// SetIssuedTo gets a reference to the given string and assigns it to the IssuedTo field.
func (o *TlsCertificateResponseAttributes) SetIssuedTo(v string) {
	o.IssuedTo = &v
}

// GetIssuer returns the Issuer field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributes) GetIssuer() string {
	if o == nil || o.Issuer == nil {
		var ret string
		return ret
	}
	return *o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributes) GetIssuerOk() (*string, bool) {
	if o == nil || o.Issuer == nil {
		return nil, false
	}
	return o.Issuer, true
}

// HasIssuer returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributes) HasIssuer() bool {
	if o != nil && o.Issuer != nil {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given string and assigns it to the Issuer field.
func (o *TlsCertificateResponseAttributes) SetIssuer(v string) {
	o.Issuer = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributes) GetSerialNumber() string {
	if o == nil || o.SerialNumber == nil {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributes) GetSerialNumberOk() (*string, bool) {
	if o == nil || o.SerialNumber == nil {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributes) HasSerialNumber() bool {
	if o != nil && o.SerialNumber != nil {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *TlsCertificateResponseAttributes) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetSignatureAlgorithm returns the SignatureAlgorithm field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributes) GetSignatureAlgorithm() string {
	if o == nil || o.SignatureAlgorithm == nil {
		var ret string
		return ret
	}
	return *o.SignatureAlgorithm
}

// GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributes) GetSignatureAlgorithmOk() (*string, bool) {
	if o == nil || o.SignatureAlgorithm == nil {
		return nil, false
	}
	return o.SignatureAlgorithm, true
}

// HasSignatureAlgorithm returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributes) HasSignatureAlgorithm() bool {
	if o != nil && o.SignatureAlgorithm != nil {
		return true
	}

	return false
}

// SetSignatureAlgorithm gets a reference to the given string and assigns it to the SignatureAlgorithm field.
func (o *TlsCertificateResponseAttributes) SetSignatureAlgorithm(v string) {
	o.SignatureAlgorithm = &v
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributes) GetNotAfter() time.Time {
	if o == nil || o.NotAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributes) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributes) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *TlsCertificateResponseAttributes) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributes) GetNotBefore() time.Time {
	if o == nil || o.NotBefore == nil {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributes) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || o.NotBefore == nil {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributes) HasNotBefore() bool {
	if o != nil && o.NotBefore != nil {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *TlsCertificateResponseAttributes) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *TlsCertificateResponseAttributes) GetReplace() bool {
	if o == nil || o.Replace == nil {
		var ret bool
		return ret
	}
	return *o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseAttributes) GetReplaceOk() (*bool, bool) {
	if o == nil || o.Replace == nil {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *TlsCertificateResponseAttributes) HasReplace() bool {
	if o != nil && o.Replace != nil {
		return true
	}

	return false
}

// SetReplace gets a reference to the given bool and assigns it to the Replace field.
func (o *TlsCertificateResponseAttributes) SetReplace(v bool) {
	o.Replace = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsCertificateResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
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
func (o *TlsCertificateResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTlsCertificateResponseAttributes := _TlsCertificateResponseAttributes{}

	if err = json.Unmarshal(bytes, &varTlsCertificateResponseAttributes); err == nil {
		*o = TlsCertificateResponseAttributes(varTlsCertificateResponseAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
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

// NullableTlsCertificateResponseAttributes is a helper abstraction for handling nullable tlscertificateresponseattributes types.
type NullableTlsCertificateResponseAttributes struct {
	value *TlsCertificateResponseAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTlsCertificateResponseAttributes) Get() *TlsCertificateResponseAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsCertificateResponseAttributes) Set(val *TlsCertificateResponseAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsCertificateResponseAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsCertificateResponseAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsCertificateResponseAttributes returns a pointer to a new instance of NullableTlsCertificateResponseAttributes.
func NewNullableTlsCertificateResponseAttributes(val *TlsCertificateResponseAttributes) *NullableTlsCertificateResponseAttributes {
	return &NullableTlsCertificateResponseAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsCertificateResponseAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsCertificateResponseAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
