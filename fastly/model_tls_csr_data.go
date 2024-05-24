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

// TLSCsrData struct for TLSCsrData
type TLSCsrData struct {
	Type *TypeTLSCsr `json:"type,omitempty"`
	Attributes *TLSCsrDataAttributes `json:"attributes,omitempty"`
	Relationships *RelationshipTLSPrivateKey `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCsrData TLSCsrData

// NewTLSCsrData instantiates a new TLSCsrData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCsrData() *TLSCsrData {
	this := TLSCsrData{}
	var resourceType TypeTLSCsr = TYPETLSCSR_CSR
	this.Type = &resourceType
	return &this
}

// NewTLSCsrDataWithDefaults instantiates a new TLSCsrData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCsrDataWithDefaults() *TLSCsrData {
	this := TLSCsrData{}
	var resourceType TypeTLSCsr = TYPETLSCSR_CSR
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TLSCsrData) GetType() TypeTLSCsr {
	if o == nil || o.Type == nil {
		var ret TypeTLSCsr
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrData) GetTypeOk() (*TypeTLSCsr, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TLSCsrData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSCsr and assigns it to the Type field.
func (o *TLSCsrData) SetType(v TypeTLSCsr) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSCsrData) GetAttributes() TLSCsrDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TLSCsrDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrData) GetAttributesOk() (*TLSCsrDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSCsrData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSCsrDataAttributes and assigns it to the Attributes field.
func (o *TLSCsrData) SetAttributes(v TLSCsrDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TLSCsrData) GetRelationships() RelationshipTLSPrivateKey {
	if o == nil || o.Relationships == nil {
		var ret RelationshipTLSPrivateKey
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrData) GetRelationshipsOk() (*RelationshipTLSPrivateKey, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TLSCsrData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipTLSPrivateKey and assigns it to the Relationships field.
func (o *TLSCsrData) SetRelationships(v RelationshipTLSPrivateKey) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCsrData) MarshalJSON() ([]byte, error) {
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
func (o *TLSCsrData) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCsrData := _TLSCsrData{}

	if err = json.Unmarshal(bytes, &varTLSCsrData); err == nil {
		*o = TLSCsrData(varTLSCsrData)
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

// NullableTLSCsrData is a helper abstraction for handling nullable tlscsrdata types. 
type NullableTLSCsrData struct {
	value *TLSCsrData
	isSet bool
}

// Get returns the value.
func (v NullableTLSCsrData) Get() *TLSCsrData {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCsrData) Set(val *TLSCsrData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCsrData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCsrData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCsrData returns a pointer to a new instance of NullableTLSCsrData.
func NewNullableTLSCsrData(val *TLSCsrData) *NullableTLSCsrData {
	return &NullableTLSCsrData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCsrData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSCsrData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
