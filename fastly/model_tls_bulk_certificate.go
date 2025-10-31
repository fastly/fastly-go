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

// TlsBulkCertificate struct for TlsBulkCertificate
type TlsBulkCertificate struct {
	Data                 *TlsBulkCertificateData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsBulkCertificate TlsBulkCertificate

// NewTlsBulkCertificate instantiates a new TlsBulkCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsBulkCertificate() *TlsBulkCertificate {
	this := TlsBulkCertificate{}
	return &this
}

// NewTlsBulkCertificateWithDefaults instantiates a new TlsBulkCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsBulkCertificateWithDefaults() *TlsBulkCertificate {
	this := TlsBulkCertificate{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TlsBulkCertificate) GetData() TlsBulkCertificateData {
	if o == nil || o.Data == nil {
		var ret TlsBulkCertificateData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsBulkCertificate) GetDataOk() (*TlsBulkCertificateData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TlsBulkCertificate) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given TlsBulkCertificateData and assigns it to the Data field.
func (o *TlsBulkCertificate) SetData(v TlsBulkCertificateData) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsBulkCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsBulkCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varTlsBulkCertificate := _TlsBulkCertificate{}

	if err = json.Unmarshal(bytes, &varTlsBulkCertificate); err == nil {
		*o = TlsBulkCertificate(varTlsBulkCertificate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsBulkCertificate is a helper abstraction for handling nullable tlsbulkcertificate types.
type NullableTlsBulkCertificate struct {
	value *TlsBulkCertificate
	isSet bool
}

// Get returns the value.
func (v NullableTlsBulkCertificate) Get() *TlsBulkCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsBulkCertificate) Set(val *TlsBulkCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsBulkCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsBulkCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsBulkCertificate returns a pointer to a new instance of NullableTlsBulkCertificate.
func NewNullableTlsBulkCertificate(val *TlsBulkCertificate) *NullableTlsBulkCertificate {
	return &NullableTlsBulkCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsBulkCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsBulkCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
