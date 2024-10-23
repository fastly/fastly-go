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

// TLSBulkCertificateResponseAttributesAllOf struct for TLSBulkCertificateResponseAttributesAllOf
type TLSBulkCertificateResponseAttributesAllOf struct {
	// Time-stamp (GMT) when the certificate will expire. Must be in the future to be used to terminate TLS traffic.
	NotAfter *time.Time `json:"not_after,omitempty"`
	// Time-stamp (GMT) when the certificate will become valid. Must be in the past to be used to terminate TLS traffic.
	NotBefore *time.Time `json:"not_before,omitempty"`
	// A recommendation from Fastly indicating the key associated with this certificate is in need of rotation.
	Replace              *bool `json:"replace,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSBulkCertificateResponseAttributesAllOf TLSBulkCertificateResponseAttributesAllOf

// NewTLSBulkCertificateResponseAttributesAllOf instantiates a new TLSBulkCertificateResponseAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSBulkCertificateResponseAttributesAllOf() *TLSBulkCertificateResponseAttributesAllOf {
	this := TLSBulkCertificateResponseAttributesAllOf{}
	return &this
}

// NewTLSBulkCertificateResponseAttributesAllOfWithDefaults instantiates a new TLSBulkCertificateResponseAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSBulkCertificateResponseAttributesAllOfWithDefaults() *TLSBulkCertificateResponseAttributesAllOf {
	this := TLSBulkCertificateResponseAttributesAllOf{}
	return &this
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *TLSBulkCertificateResponseAttributesAllOf) GetNotAfter() time.Time {
	if o == nil || o.NotAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSBulkCertificateResponseAttributesAllOf) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *TLSBulkCertificateResponseAttributesAllOf) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *TLSBulkCertificateResponseAttributesAllOf) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *TLSBulkCertificateResponseAttributesAllOf) GetNotBefore() time.Time {
	if o == nil || o.NotBefore == nil {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSBulkCertificateResponseAttributesAllOf) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || o.NotBefore == nil {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *TLSBulkCertificateResponseAttributesAllOf) HasNotBefore() bool {
	if o != nil && o.NotBefore != nil {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *TLSBulkCertificateResponseAttributesAllOf) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *TLSBulkCertificateResponseAttributesAllOf) GetReplace() bool {
	if o == nil || o.Replace == nil {
		var ret bool
		return ret
	}
	return *o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSBulkCertificateResponseAttributesAllOf) GetReplaceOk() (*bool, bool) {
	if o == nil || o.Replace == nil {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *TLSBulkCertificateResponseAttributesAllOf) HasReplace() bool {
	if o != nil && o.Replace != nil {
		return true
	}

	return false
}

// SetReplace gets a reference to the given bool and assigns it to the Replace field.
func (o *TLSBulkCertificateResponseAttributesAllOf) SetReplace(v bool) {
	o.Replace = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSBulkCertificateResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *TLSBulkCertificateResponseAttributesAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTLSBulkCertificateResponseAttributesAllOf := _TLSBulkCertificateResponseAttributesAllOf{}

	if err = json.Unmarshal(bytes, &varTLSBulkCertificateResponseAttributesAllOf); err == nil {
		*o = TLSBulkCertificateResponseAttributesAllOf(varTLSBulkCertificateResponseAttributesAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "not_after")
		delete(additionalProperties, "not_before")
		delete(additionalProperties, "replace")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSBulkCertificateResponseAttributesAllOf is a helper abstraction for handling nullable tlsbulkcertificateresponseattributesallof types.
type NullableTLSBulkCertificateResponseAttributesAllOf struct {
	value *TLSBulkCertificateResponseAttributesAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTLSBulkCertificateResponseAttributesAllOf) Get() *TLSBulkCertificateResponseAttributesAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSBulkCertificateResponseAttributesAllOf) Set(val *TLSBulkCertificateResponseAttributesAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSBulkCertificateResponseAttributesAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSBulkCertificateResponseAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSBulkCertificateResponseAttributesAllOf returns a pointer to a new instance of NullableTLSBulkCertificateResponseAttributesAllOf.
func NewNullableTLSBulkCertificateResponseAttributesAllOf(val *TLSBulkCertificateResponseAttributesAllOf) *NullableTLSBulkCertificateResponseAttributesAllOf {
	return &NullableTLSBulkCertificateResponseAttributesAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSBulkCertificateResponseAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTLSBulkCertificateResponseAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
