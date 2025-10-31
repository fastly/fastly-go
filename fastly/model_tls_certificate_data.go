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

// TlsCertificateData struct for TlsCertificateData
type TlsCertificateData struct {
	Type                 *TypeTlsCertificate           `json:"type,omitempty"`
	Attributes           *TlsCertificateDataAttributes `json:"attributes,omitempty"`
	Relationships        *RelationshipTlsDomains       `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsCertificateData TlsCertificateData

// NewTlsCertificateData instantiates a new TlsCertificateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsCertificateData() *TlsCertificateData {
	this := TlsCertificateData{}
	var type_ TypeTlsCertificate = TYPETLSCERTIFICATE_TLS_CERTIFICATE
	this.Type = &type_
	return &this
}

// NewTlsCertificateDataWithDefaults instantiates a new TlsCertificateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsCertificateDataWithDefaults() *TlsCertificateData {
	this := TlsCertificateData{}
	var type_ TypeTlsCertificate = TYPETLSCERTIFICATE_TLS_CERTIFICATE
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TlsCertificateData) GetType() TypeTlsCertificate {
	if o == nil || o.Type == nil {
		var ret TypeTlsCertificate
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateData) GetTypeOk() (*TypeTlsCertificate, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TlsCertificateData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsCertificate and assigns it to the Type field.
func (o *TlsCertificateData) SetType(v TypeTlsCertificate) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TlsCertificateData) GetAttributes() TlsCertificateDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TlsCertificateDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateData) GetAttributesOk() (*TlsCertificateDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TlsCertificateData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TlsCertificateDataAttributes and assigns it to the Attributes field.
func (o *TlsCertificateData) SetAttributes(v TlsCertificateDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TlsCertificateData) GetRelationships() RelationshipTlsDomains {
	if o == nil || o.Relationships == nil {
		var ret RelationshipTlsDomains
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsCertificateData) GetRelationshipsOk() (*RelationshipTlsDomains, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TlsCertificateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipTlsDomains and assigns it to the Relationships field.
func (o *TlsCertificateData) SetRelationships(v RelationshipTlsDomains) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsCertificateData) MarshalJSON() ([]byte, error) {
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

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsCertificateData) UnmarshalJSON(bytes []byte) (err error) {
	varTlsCertificateData := _TlsCertificateData{}

	if err = json.Unmarshal(bytes, &varTlsCertificateData); err == nil {
		*o = TlsCertificateData(varTlsCertificateData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsCertificateData is a helper abstraction for handling nullable tlscertificatedata types.
type NullableTlsCertificateData struct {
	value *TlsCertificateData
	isSet bool
}

// Get returns the value.
func (v NullableTlsCertificateData) Get() *TlsCertificateData {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsCertificateData) Set(val *TlsCertificateData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsCertificateData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsCertificateData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsCertificateData returns a pointer to a new instance of NullableTlsCertificateData.
func NewNullableTlsCertificateData(val *TlsCertificateData) *NullableTlsCertificateData {
	return &NullableTlsCertificateData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsCertificateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsCertificateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
