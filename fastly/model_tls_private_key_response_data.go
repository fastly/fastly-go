// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// TLSPrivateKeyResponseData struct for TLSPrivateKeyResponseData
type TLSPrivateKeyResponseData struct {
	Type *TypeTLSPrivateKey `json:"type,omitempty"`
	ID *string `json:"id,omitempty"`
	Attributes *TLSPrivateKeyResponseAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSPrivateKeyResponseData TLSPrivateKeyResponseData

// NewTLSPrivateKeyResponseData instantiates a new TLSPrivateKeyResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSPrivateKeyResponseData() *TLSPrivateKeyResponseData {
	this := TLSPrivateKeyResponseData{}
	var resourceType TypeTLSPrivateKey = TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY
	this.Type = &resourceType
	return &this
}

// NewTLSPrivateKeyResponseDataWithDefaults instantiates a new TLSPrivateKeyResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSPrivateKeyResponseDataWithDefaults() *TLSPrivateKeyResponseData {
	this := TLSPrivateKeyResponseData{}
	var resourceType TypeTLSPrivateKey = TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseData) GetType() TypeTLSPrivateKey {
	if o == nil || o.Type == nil {
		var ret TypeTLSPrivateKey
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseData) GetTypeOk() (*TypeTLSPrivateKey, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSPrivateKey and assigns it to the Type field.
func (o *TLSPrivateKeyResponseData) SetType(v TypeTLSPrivateKey) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TLSPrivateKeyResponseData) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSPrivateKeyResponseData) GetAttributes() TLSPrivateKeyResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret TLSPrivateKeyResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyResponseData) GetAttributesOk() (*TLSPrivateKeyResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSPrivateKeyResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSPrivateKeyResponseAttributes and assigns it to the Attributes field.
func (o *TLSPrivateKeyResponseData) SetAttributes(v TLSPrivateKeyResponseAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSPrivateKeyResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
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
func (o *TLSPrivateKeyResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varTLSPrivateKeyResponseData := _TLSPrivateKeyResponseData{}

	if err = json.Unmarshal(bytes, &varTLSPrivateKeyResponseData); err == nil {
		*o = TLSPrivateKeyResponseData(varTLSPrivateKeyResponseData)
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

// NullableTLSPrivateKeyResponseData is a helper abstraction for handling nullable tlsprivatekeyresponsedata types. 
type NullableTLSPrivateKeyResponseData struct {
	value *TLSPrivateKeyResponseData
	isSet bool
}

// Get returns the value.
func (v NullableTLSPrivateKeyResponseData) Get() *TLSPrivateKeyResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSPrivateKeyResponseData) Set(val *TLSPrivateKeyResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSPrivateKeyResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSPrivateKeyResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSPrivateKeyResponseData returns a pointer to a new instance of NullableTLSPrivateKeyResponseData.
func NewNullableTLSPrivateKeyResponseData(val *TLSPrivateKeyResponseData) *NullableTLSPrivateKeyResponseData {
	return &NullableTLSPrivateKeyResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSPrivateKeyResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSPrivateKeyResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
