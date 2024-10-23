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

// TLSCertificateData struct for TLSCertificateData
type TLSCertificateData struct {
	Type                 *TypeTLSCertificate           `json:"type,omitempty"`
	Attributes           *TLSCertificateDataAttributes `json:"attributes,omitempty"`
	Relationships        *RelationshipTLSDomains       `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCertificateData TLSCertificateData

// NewTLSCertificateData instantiates a new TLSCertificateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCertificateData() *TLSCertificateData {
	this := TLSCertificateData{}
	var resourceType TypeTLSCertificate = TYPETLSCERTIFICATE_TLS_CERTIFICATE
	this.Type = &resourceType
	return &this
}

// NewTLSCertificateDataWithDefaults instantiates a new TLSCertificateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCertificateDataWithDefaults() *TLSCertificateData {
	this := TLSCertificateData{}
	var resourceType TypeTLSCertificate = TYPETLSCERTIFICATE_TLS_CERTIFICATE
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TLSCertificateData) GetType() TypeTLSCertificate {
	if o == nil || o.Type == nil {
		var ret TypeTLSCertificate
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCertificateData) GetTypeOk() (*TypeTLSCertificate, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TLSCertificateData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSCertificate and assigns it to the Type field.
func (o *TLSCertificateData) SetType(v TypeTLSCertificate) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSCertificateData) GetAttributes() TLSCertificateDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TLSCertificateDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCertificateData) GetAttributesOk() (*TLSCertificateDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSCertificateData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSCertificateDataAttributes and assigns it to the Attributes field.
func (o *TLSCertificateData) SetAttributes(v TLSCertificateDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TLSCertificateData) GetRelationships() RelationshipTLSDomains {
	if o == nil || o.Relationships == nil {
		var ret RelationshipTLSDomains
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCertificateData) GetRelationshipsOk() (*RelationshipTLSDomains, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TLSCertificateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipTLSDomains and assigns it to the Relationships field.
func (o *TLSCertificateData) SetRelationships(v RelationshipTLSDomains) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCertificateData) MarshalJSON() ([]byte, error) {
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
func (o *TLSCertificateData) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCertificateData := _TLSCertificateData{}

	if err = json.Unmarshal(bytes, &varTLSCertificateData); err == nil {
		*o = TLSCertificateData(varTLSCertificateData)
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

// NullableTLSCertificateData is a helper abstraction for handling nullable tlscertificatedata types.
type NullableTLSCertificateData struct {
	value *TLSCertificateData
	isSet bool
}

// Get returns the value.
func (v NullableTLSCertificateData) Get() *TLSCertificateData {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCertificateData) Set(val *TLSCertificateData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCertificateData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCertificateData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCertificateData returns a pointer to a new instance of NullableTLSCertificateData.
func NewNullableTLSCertificateData(val *TLSCertificateData) *NullableTLSCertificateData {
	return &NullableTLSCertificateData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCertificateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTLSCertificateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
