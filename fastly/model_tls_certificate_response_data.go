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

// TLSCertificateResponseData struct for TLSCertificateResponseData
type TLSCertificateResponseData struct {
	Type *TypeTLSCertificate `json:"type,omitempty"`
	Attributes *TLSCertificateResponseAttributes `json:"attributes,omitempty"`
	Relationships *RelationshipTLSDomains `json:"relationships,omitempty"`
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCertificateResponseData TLSCertificateResponseData

// NewTLSCertificateResponseData instantiates a new TLSCertificateResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCertificateResponseData() *TLSCertificateResponseData {
	this := TLSCertificateResponseData{}
	var resourceType TypeTLSCertificate = TYPETLSCERTIFICATE_TLS_CERTIFICATE
	this.Type = &resourceType
	return &this
}

// NewTLSCertificateResponseDataWithDefaults instantiates a new TLSCertificateResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCertificateResponseDataWithDefaults() *TLSCertificateResponseData {
	this := TLSCertificateResponseData{}
	var resourceType TypeTLSCertificate = TYPETLSCERTIFICATE_TLS_CERTIFICATE
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TLSCertificateResponseData) GetType() TypeTLSCertificate {
	if o == nil || o.Type == nil {
		var ret TypeTLSCertificate
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCertificateResponseData) GetTypeOk() (*TypeTLSCertificate, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TLSCertificateResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSCertificate and assigns it to the Type field.
func (o *TLSCertificateResponseData) SetType(v TypeTLSCertificate) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSCertificateResponseData) GetAttributes() TLSCertificateResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret TLSCertificateResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCertificateResponseData) GetAttributesOk() (*TLSCertificateResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSCertificateResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSCertificateResponseAttributes and assigns it to the Attributes field.
func (o *TLSCertificateResponseData) SetAttributes(v TLSCertificateResponseAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TLSCertificateResponseData) GetRelationships() RelationshipTLSDomains {
	if o == nil || o.Relationships == nil {
		var ret RelationshipTLSDomains
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCertificateResponseData) GetRelationshipsOk() (*RelationshipTLSDomains, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TLSCertificateResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipTLSDomains and assigns it to the Relationships field.
func (o *TLSCertificateResponseData) SetRelationships(v RelationshipTLSDomains) {
	o.Relationships = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TLSCertificateResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCertificateResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TLSCertificateResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TLSCertificateResponseData) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCertificateResponseData) MarshalJSON() ([]byte, error) {
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
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *TLSCertificateResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCertificateResponseData := _TLSCertificateResponseData{}

	if err = json.Unmarshal(bytes, &varTLSCertificateResponseData); err == nil {
		*o = TLSCertificateResponseData(varTLSCertificateResponseData)
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

// NullableTLSCertificateResponseData is a helper abstraction for handling nullable tlscertificateresponsedata types. 
type NullableTLSCertificateResponseData struct {
	value *TLSCertificateResponseData
	isSet bool
}

// Get returns the value.
func (v NullableTLSCertificateResponseData) Get() *TLSCertificateResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCertificateResponseData) Set(val *TLSCertificateResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCertificateResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCertificateResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCertificateResponseData returns a pointer to a new instance of NullableTLSCertificateResponseData.
func NewNullableTLSCertificateResponseData(val *TLSCertificateResponseData) *NullableTLSCertificateResponseData {
	return &NullableTLSCertificateResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCertificateResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSCertificateResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
