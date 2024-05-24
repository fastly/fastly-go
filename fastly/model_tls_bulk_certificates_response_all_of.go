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

// TLSBulkCertificatesResponseAllOf struct for TLSBulkCertificatesResponseAllOf
type TLSBulkCertificatesResponseAllOf struct {
	Data []TLSBulkCertificateResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSBulkCertificatesResponseAllOf TLSBulkCertificatesResponseAllOf

// NewTLSBulkCertificatesResponseAllOf instantiates a new TLSBulkCertificatesResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSBulkCertificatesResponseAllOf() *TLSBulkCertificatesResponseAllOf {
	this := TLSBulkCertificatesResponseAllOf{}
	return &this
}

// NewTLSBulkCertificatesResponseAllOfWithDefaults instantiates a new TLSBulkCertificatesResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSBulkCertificatesResponseAllOfWithDefaults() *TLSBulkCertificatesResponseAllOf {
	this := TLSBulkCertificatesResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TLSBulkCertificatesResponseAllOf) GetData() []TLSBulkCertificateResponseData {
	if o == nil || o.Data == nil {
		var ret []TLSBulkCertificateResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSBulkCertificatesResponseAllOf) GetDataOk() ([]TLSBulkCertificateResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TLSBulkCertificatesResponseAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []TLSBulkCertificateResponseData and assigns it to the Data field.
func (o *TLSBulkCertificatesResponseAllOf) SetData(v []TLSBulkCertificateResponseData) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSBulkCertificatesResponseAllOf) MarshalJSON() ([]byte, error) {
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
func (o *TLSBulkCertificatesResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTLSBulkCertificatesResponseAllOf := _TLSBulkCertificatesResponseAllOf{}

	if err = json.Unmarshal(bytes, &varTLSBulkCertificatesResponseAllOf); err == nil {
		*o = TLSBulkCertificatesResponseAllOf(varTLSBulkCertificatesResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSBulkCertificatesResponseAllOf is a helper abstraction for handling nullable tlsbulkcertificatesresponseallof types. 
type NullableTLSBulkCertificatesResponseAllOf struct {
	value *TLSBulkCertificatesResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTLSBulkCertificatesResponseAllOf) Get() *TLSBulkCertificatesResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSBulkCertificatesResponseAllOf) Set(val *TLSBulkCertificatesResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSBulkCertificatesResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSBulkCertificatesResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSBulkCertificatesResponseAllOf returns a pointer to a new instance of NullableTLSBulkCertificatesResponseAllOf.
func NewNullableTLSBulkCertificatesResponseAllOf(val *TLSBulkCertificatesResponseAllOf) *NullableTLSBulkCertificatesResponseAllOf {
	return &NullableTLSBulkCertificatesResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSBulkCertificatesResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSBulkCertificatesResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
