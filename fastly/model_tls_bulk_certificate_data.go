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

// TlsBulkCertificateData struct for TlsBulkCertificateData
type TlsBulkCertificateData struct {
	Type                 *TypeTlsBulkCertificate             `json:"type,omitempty"`
	Attributes           *TlsBulkCertificateDataAttributes   `json:"attributes,omitempty"`
	Relationships        *RelationshipsForTlsBulkCertificate `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsBulkCertificateData TlsBulkCertificateData

// NewTlsBulkCertificateData instantiates a new TlsBulkCertificateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsBulkCertificateData() *TlsBulkCertificateData {
	this := TlsBulkCertificateData{}
	var type_ TypeTlsBulkCertificate = TYPETLSBULKCERTIFICATE_TLS_BULK_CERTIFICATE
	this.Type = &type_
	return &this
}

// NewTlsBulkCertificateDataWithDefaults instantiates a new TlsBulkCertificateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsBulkCertificateDataWithDefaults() *TlsBulkCertificateData {
	this := TlsBulkCertificateData{}
	var type_ TypeTlsBulkCertificate = TYPETLSBULKCERTIFICATE_TLS_BULK_CERTIFICATE
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TlsBulkCertificateData) GetType() TypeTlsBulkCertificate {
	if o == nil || o.Type == nil {
		var ret TypeTlsBulkCertificate
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsBulkCertificateData) GetTypeOk() (*TypeTlsBulkCertificate, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TlsBulkCertificateData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsBulkCertificate and assigns it to the Type field.
func (o *TlsBulkCertificateData) SetType(v TypeTlsBulkCertificate) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TlsBulkCertificateData) GetAttributes() TlsBulkCertificateDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TlsBulkCertificateDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsBulkCertificateData) GetAttributesOk() (*TlsBulkCertificateDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TlsBulkCertificateData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TlsBulkCertificateDataAttributes and assigns it to the Attributes field.
func (o *TlsBulkCertificateData) SetAttributes(v TlsBulkCertificateDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TlsBulkCertificateData) GetRelationships() RelationshipsForTlsBulkCertificate {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTlsBulkCertificate
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsBulkCertificateData) GetRelationshipsOk() (*RelationshipsForTlsBulkCertificate, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TlsBulkCertificateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTlsBulkCertificate and assigns it to the Relationships field.
func (o *TlsBulkCertificateData) SetRelationships(v RelationshipsForTlsBulkCertificate) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsBulkCertificateData) MarshalJSON() ([]byte, error) {
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
func (o *TlsBulkCertificateData) UnmarshalJSON(bytes []byte) (err error) {
	varTlsBulkCertificateData := _TlsBulkCertificateData{}

	if err = json.Unmarshal(bytes, &varTlsBulkCertificateData); err == nil {
		*o = TlsBulkCertificateData(varTlsBulkCertificateData)
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

// NullableTlsBulkCertificateData is a helper abstraction for handling nullable tlsbulkcertificatedata types.
type NullableTlsBulkCertificateData struct {
	value *TlsBulkCertificateData
	isSet bool
}

// Get returns the value.
func (v NullableTlsBulkCertificateData) Get() *TlsBulkCertificateData {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsBulkCertificateData) Set(val *TlsBulkCertificateData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsBulkCertificateData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsBulkCertificateData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsBulkCertificateData returns a pointer to a new instance of NullableTlsBulkCertificateData.
func NewNullableTlsBulkCertificateData(val *TlsBulkCertificateData) *NullableTlsBulkCertificateData {
	return &NullableTlsBulkCertificateData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsBulkCertificateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsBulkCertificateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
