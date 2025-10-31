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

// TlsPrivateKeyResponseData struct for TlsPrivateKeyResponseData
type TlsPrivateKeyResponseData struct {
	Type                 *TypeTlsPrivateKey               `json:"type,omitempty"`
	Id                   *string                          `json:"id,omitempty"`
	Attributes           *TlsPrivateKeyResponseAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsPrivateKeyResponseData TlsPrivateKeyResponseData

// NewTlsPrivateKeyResponseData instantiates a new TlsPrivateKeyResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsPrivateKeyResponseData() *TlsPrivateKeyResponseData {
	this := TlsPrivateKeyResponseData{}
	var type_ TypeTlsPrivateKey = TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY
	this.Type = &type_
	return &this
}

// NewTlsPrivateKeyResponseDataWithDefaults instantiates a new TlsPrivateKeyResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsPrivateKeyResponseDataWithDefaults() *TlsPrivateKeyResponseData {
	this := TlsPrivateKeyResponseData{}
	var type_ TypeTlsPrivateKey = TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TlsPrivateKeyResponseData) GetType() TypeTlsPrivateKey {
	if o == nil || o.Type == nil {
		var ret TypeTlsPrivateKey
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsPrivateKeyResponseData) GetTypeOk() (*TypeTlsPrivateKey, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TlsPrivateKeyResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsPrivateKey and assigns it to the Type field.
func (o *TlsPrivateKeyResponseData) SetType(v TypeTlsPrivateKey) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TlsPrivateKeyResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsPrivateKeyResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TlsPrivateKeyResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TlsPrivateKeyResponseData) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TlsPrivateKeyResponseData) GetAttributes() TlsPrivateKeyResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret TlsPrivateKeyResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsPrivateKeyResponseData) GetAttributesOk() (*TlsPrivateKeyResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TlsPrivateKeyResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TlsPrivateKeyResponseAttributes and assigns it to the Attributes field.
func (o *TlsPrivateKeyResponseData) SetAttributes(v TlsPrivateKeyResponseAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsPrivateKeyResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
func (o *TlsPrivateKeyResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varTlsPrivateKeyResponseData := _TlsPrivateKeyResponseData{}

	if err = json.Unmarshal(bytes, &varTlsPrivateKeyResponseData); err == nil {
		*o = TlsPrivateKeyResponseData(varTlsPrivateKeyResponseData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsPrivateKeyResponseData is a helper abstraction for handling nullable tlsprivatekeyresponsedata types.
type NullableTlsPrivateKeyResponseData struct {
	value *TlsPrivateKeyResponseData
	isSet bool
}

// Get returns the value.
func (v NullableTlsPrivateKeyResponseData) Get() *TlsPrivateKeyResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsPrivateKeyResponseData) Set(val *TlsPrivateKeyResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsPrivateKeyResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsPrivateKeyResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsPrivateKeyResponseData returns a pointer to a new instance of NullableTlsPrivateKeyResponseData.
func NewNullableTlsPrivateKeyResponseData(val *TlsPrivateKeyResponseData) *NullableTlsPrivateKeyResponseData {
	return &NullableTlsPrivateKeyResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsPrivateKeyResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsPrivateKeyResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
