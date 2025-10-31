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

// TlsCertificateResponseData struct for TlsCertificateResponseData
type TlsCertificateResponseData struct {
	Type                 *TypeTlsCertificate               `json:"type,omitempty"`
	Attributes           *TlsCertificateResponseAttributes `json:"attributes,omitempty"`
	Relationships        *RelationshipTlsDomains           `json:"relationships,omitempty"`
	Id                   *string                           `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsCertificateResponseData TlsCertificateResponseData

// NewTlsCertificateResponseData instantiates a new TlsCertificateResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsCertificateResponseData() *TlsCertificateResponseData {
	this := TlsCertificateResponseData{}
	var type_ TypeTlsCertificate = TYPETLSCERTIFICATE_TLS_CERTIFICATE
	this.Type = &type_
	return &this
}

// NewTlsCertificateResponseDataWithDefaults instantiates a new TlsCertificateResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsCertificateResponseDataWithDefaults() *TlsCertificateResponseData {
	this := TlsCertificateResponseData{}
	var type_ TypeTlsCertificate = TYPETLSCERTIFICATE_TLS_CERTIFICATE
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TlsCertificateResponseData) GetType() TypeTlsCertificate {
	if o == nil || o.Type == nil {
		var ret TypeTlsCertificate
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseData) GetTypeOk() (*TypeTlsCertificate, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TlsCertificateResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsCertificate and assigns it to the Type field.
func (o *TlsCertificateResponseData) SetType(v TypeTlsCertificate) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TlsCertificateResponseData) GetAttributes() TlsCertificateResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret TlsCertificateResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseData) GetAttributesOk() (*TlsCertificateResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TlsCertificateResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TlsCertificateResponseAttributes and assigns it to the Attributes field.
func (o *TlsCertificateResponseData) SetAttributes(v TlsCertificateResponseAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TlsCertificateResponseData) GetRelationships() RelationshipTlsDomains {
	if o == nil || o.Relationships == nil {
		var ret RelationshipTlsDomains
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseData) GetRelationshipsOk() (*RelationshipTlsDomains, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TlsCertificateResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipTlsDomains and assigns it to the Relationships field.
func (o *TlsCertificateResponseData) SetRelationships(v RelationshipTlsDomains) {
	o.Relationships = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TlsCertificateResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TlsCertificateResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TlsCertificateResponseData) SetId(v string) {
	o.Id = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsCertificateResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsCertificateResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varTlsCertificateResponseData := _TlsCertificateResponseData{}

	if err = json.Unmarshal(bytes, &varTlsCertificateResponseData); err == nil {
		*o = TlsCertificateResponseData(varTlsCertificateResponseData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsCertificateResponseData is a helper abstraction for handling nullable tlscertificateresponsedata types.
type NullableTlsCertificateResponseData struct {
	value *TlsCertificateResponseData
	isSet bool
}

// Get returns the value.
func (v NullableTlsCertificateResponseData) Get() *TlsCertificateResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsCertificateResponseData) Set(val *TlsCertificateResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsCertificateResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsCertificateResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsCertificateResponseData returns a pointer to a new instance of NullableTlsCertificateResponseData.
func NewNullableTlsCertificateResponseData(val *TlsCertificateResponseData) *NullableTlsCertificateResponseData {
	return &NullableTlsCertificateResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsCertificateResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsCertificateResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
