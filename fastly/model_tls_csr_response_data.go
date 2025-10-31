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

// TlsCsrResponseData struct for TlsCsrResponseData
type TlsCsrResponseData struct {
	Id                   *string                    `json:"id,omitempty"`
	Type                 *TypeTlsCsr                `json:"type,omitempty"`
	Attributes           *TlsCsrResponseAttributes  `json:"attributes,omitempty"`
	Relationships        *RelationshipTlsPrivateKey `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsCsrResponseData TlsCsrResponseData

// NewTlsCsrResponseData instantiates a new TlsCsrResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsCsrResponseData() *TlsCsrResponseData {
	this := TlsCsrResponseData{}
	var type_ TypeTlsCsr = TYPETLSCSR_CSR
	this.Type = &type_
	return &this
}

// NewTlsCsrResponseDataWithDefaults instantiates a new TlsCsrResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsCsrResponseDataWithDefaults() *TlsCsrResponseData {
	this := TlsCsrResponseData{}
	var type_ TypeTlsCsr = TYPETLSCSR_CSR
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TlsCsrResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCsrResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TlsCsrResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TlsCsrResponseData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TlsCsrResponseData) GetType() TypeTlsCsr {
	if o == nil || o.Type == nil {
		var ret TypeTlsCsr
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCsrResponseData) GetTypeOk() (*TypeTlsCsr, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TlsCsrResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsCsr and assigns it to the Type field.
func (o *TlsCsrResponseData) SetType(v TypeTlsCsr) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TlsCsrResponseData) GetAttributes() TlsCsrResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret TlsCsrResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCsrResponseData) GetAttributesOk() (*TlsCsrResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TlsCsrResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TlsCsrResponseAttributes and assigns it to the Attributes field.
func (o *TlsCsrResponseData) SetAttributes(v TlsCsrResponseAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TlsCsrResponseData) GetRelationships() RelationshipTlsPrivateKey {
	if o == nil || o.Relationships == nil {
		var ret RelationshipTlsPrivateKey
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCsrResponseData) GetRelationshipsOk() (*RelationshipTlsPrivateKey, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TlsCsrResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipTlsPrivateKey and assigns it to the Relationships field.
func (o *TlsCsrResponseData) SetRelationships(v RelationshipTlsPrivateKey) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsCsrResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsCsrResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varTlsCsrResponseData := _TlsCsrResponseData{}

	if err = json.Unmarshal(bytes, &varTlsCsrResponseData); err == nil {
		*o = TlsCsrResponseData(varTlsCsrResponseData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsCsrResponseData is a helper abstraction for handling nullable tlscsrresponsedata types.
type NullableTlsCsrResponseData struct {
	value *TlsCsrResponseData
	isSet bool
}

// Get returns the value.
func (v NullableTlsCsrResponseData) Get() *TlsCsrResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsCsrResponseData) Set(val *TlsCsrResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsCsrResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsCsrResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsCsrResponseData returns a pointer to a new instance of NullableTlsCsrResponseData.
func NewNullableTlsCsrResponseData(val *TlsCsrResponseData) *NullableTlsCsrResponseData {
	return &NullableTlsCsrResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsCsrResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsCsrResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
