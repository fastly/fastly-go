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

// TLSBulkCertificateResponseAttributes struct for TLSBulkCertificateResponseAttributes
type TLSBulkCertificateResponseAttributes struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	// Time-stamp (GMT) when the certificate will expire. Must be in the future to be used to terminate TLS traffic.
	NotAfter *time.Time `json:"not_after,omitempty"`
	// Time-stamp (GMT) when the certificate will become valid. Must be in the past to be used to terminate TLS traffic.
	NotBefore *time.Time `json:"not_before,omitempty"`
	// A recommendation from Fastly indicating the key associated with this certificate is in need of rotation.
	Replace *bool `json:"replace,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSBulkCertificateResponseAttributes TLSBulkCertificateResponseAttributes

// NewTLSBulkCertificateResponseAttributes instantiates a new TLSBulkCertificateResponseAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSBulkCertificateResponseAttributes() *TLSBulkCertificateResponseAttributes {
	this := TLSBulkCertificateResponseAttributes{}
	return &this
}

// NewTLSBulkCertificateResponseAttributesWithDefaults instantiates a new TLSBulkCertificateResponseAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSBulkCertificateResponseAttributesWithDefaults() *TLSBulkCertificateResponseAttributes {
	this := TLSBulkCertificateResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSBulkCertificateResponseAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSBulkCertificateResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TLSBulkCertificateResponseAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *TLSBulkCertificateResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *TLSBulkCertificateResponseAttributes) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *TLSBulkCertificateResponseAttributes) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSBulkCertificateResponseAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSBulkCertificateResponseAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *TLSBulkCertificateResponseAttributes) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *TLSBulkCertificateResponseAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *TLSBulkCertificateResponseAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *TLSBulkCertificateResponseAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TLSBulkCertificateResponseAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TLSBulkCertificateResponseAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TLSBulkCertificateResponseAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *TLSBulkCertificateResponseAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *TLSBulkCertificateResponseAttributes) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *TLSBulkCertificateResponseAttributes) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *TLSBulkCertificateResponseAttributes) GetNotAfter() time.Time {
	if o == nil || o.NotAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSBulkCertificateResponseAttributes) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *TLSBulkCertificateResponseAttributes) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *TLSBulkCertificateResponseAttributes) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *TLSBulkCertificateResponseAttributes) GetNotBefore() time.Time {
	if o == nil || o.NotBefore == nil {
		var ret time.Time
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSBulkCertificateResponseAttributes) GetNotBeforeOk() (*time.Time, bool) {
	if o == nil || o.NotBefore == nil {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *TLSBulkCertificateResponseAttributes) HasNotBefore() bool {
	if o != nil && o.NotBefore != nil {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given time.Time and assigns it to the NotBefore field.
func (o *TLSBulkCertificateResponseAttributes) SetNotBefore(v time.Time) {
	o.NotBefore = &v
}

// GetReplace returns the Replace field value if set, zero value otherwise.
func (o *TLSBulkCertificateResponseAttributes) GetReplace() bool {
	if o == nil || o.Replace == nil {
		var ret bool
		return ret
	}
	return *o.Replace
}

// GetReplaceOk returns a tuple with the Replace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSBulkCertificateResponseAttributes) GetReplaceOk() (*bool, bool) {
	if o == nil || o.Replace == nil {
		return nil, false
	}
	return o.Replace, true
}

// HasReplace returns a boolean if a field has been set.
func (o *TLSBulkCertificateResponseAttributes) HasReplace() bool {
	if o != nil && o.Replace != nil {
		return true
	}

	return false
}

// SetReplace gets a reference to the given bool and assigns it to the Replace field.
func (o *TLSBulkCertificateResponseAttributes) SetReplace(v bool) {
	o.Replace = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSBulkCertificateResponseAttributes) MarshalJSON() ([]byte, error) {
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
func (o *TLSBulkCertificateResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTLSBulkCertificateResponseAttributes := _TLSBulkCertificateResponseAttributes{}

	if err = json.Unmarshal(bytes, &varTLSBulkCertificateResponseAttributes); err == nil {
		*o = TLSBulkCertificateResponseAttributes(varTLSBulkCertificateResponseAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "not_after")
		delete(additionalProperties, "not_before")
		delete(additionalProperties, "replace")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSBulkCertificateResponseAttributes is a helper abstraction for handling nullable tlsbulkcertificateresponseattributes types. 
type NullableTLSBulkCertificateResponseAttributes struct {
	value *TLSBulkCertificateResponseAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTLSBulkCertificateResponseAttributes) Get() *TLSBulkCertificateResponseAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSBulkCertificateResponseAttributes) Set(val *TLSBulkCertificateResponseAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSBulkCertificateResponseAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSBulkCertificateResponseAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSBulkCertificateResponseAttributes returns a pointer to a new instance of NullableTLSBulkCertificateResponseAttributes.
func NewNullableTLSBulkCertificateResponseAttributes(val *TLSBulkCertificateResponseAttributes) *NullableTLSBulkCertificateResponseAttributes {
	return &NullableTLSBulkCertificateResponseAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSBulkCertificateResponseAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSBulkCertificateResponseAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
