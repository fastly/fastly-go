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

// TlsCsrData struct for TlsCsrData
type TlsCsrData struct {
	Type                 *TypeTlsCsr           `json:"type,omitempty"`
	Attributes           *TlsCsrDataAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsCsrData TlsCsrData

// NewTlsCsrData instantiates a new TlsCsrData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsCsrData() *TlsCsrData {
	this := TlsCsrData{}
	var type_ TypeTlsCsr = TYPETLSCSR_CSR
	this.Type = &type_
	return &this
}

// NewTlsCsrDataWithDefaults instantiates a new TlsCsrData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsCsrDataWithDefaults() *TlsCsrData {
	this := TlsCsrData{}
	var type_ TypeTlsCsr = TYPETLSCSR_CSR
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TlsCsrData) GetType() TypeTlsCsr {
	if o == nil || o.Type == nil {
		var ret TypeTlsCsr
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCsrData) GetTypeOk() (*TypeTlsCsr, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TlsCsrData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsCsr and assigns it to the Type field.
func (o *TlsCsrData) SetType(v TypeTlsCsr) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TlsCsrData) GetAttributes() TlsCsrDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TlsCsrDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCsrData) GetAttributesOk() (*TlsCsrDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TlsCsrData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TlsCsrDataAttributes and assigns it to the Attributes field.
func (o *TlsCsrData) SetAttributes(v TlsCsrDataAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsCsrData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsCsrData) UnmarshalJSON(bytes []byte) (err error) {
	varTlsCsrData := _TlsCsrData{}

	if err = json.Unmarshal(bytes, &varTlsCsrData); err == nil {
		*o = TlsCsrData(varTlsCsrData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsCsrData is a helper abstraction for handling nullable tlscsrdata types.
type NullableTlsCsrData struct {
	value *TlsCsrData
	isSet bool
}

// Get returns the value.
func (v NullableTlsCsrData) Get() *TlsCsrData {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsCsrData) Set(val *TlsCsrData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsCsrData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsCsrData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsCsrData returns a pointer to a new instance of NullableTlsCsrData.
func NewNullableTlsCsrData(val *TlsCsrData) *NullableTlsCsrData {
	return &NullableTlsCsrData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsCsrData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsCsrData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
