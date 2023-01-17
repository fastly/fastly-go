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

// TLSCsrResponseData struct for TLSCsrResponseData
type TLSCsrResponseData struct {
	ID *string `json:"id,omitempty"`
	Type *TypeTLSCsr `json:"type,omitempty"`
	Attributes *TLSCsrResponseAttributes `json:"attributes,omitempty"`
	Relationships *RelationshipsForTLSCsr `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSCsrResponseData TLSCsrResponseData

// NewTLSCsrResponseData instantiates a new TLSCsrResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSCsrResponseData() *TLSCsrResponseData {
	this := TLSCsrResponseData{}
	var resourceType TypeTLSCsr = TYPETLSCSR_CSR
	this.Type = &resourceType
	return &this
}

// NewTLSCsrResponseDataWithDefaults instantiates a new TLSCsrResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSCsrResponseDataWithDefaults() *TLSCsrResponseData {
	this := TLSCsrResponseData{}
	var resourceType TypeTLSCsr = TYPETLSCSR_CSR
	this.Type = &resourceType
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TLSCsrResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TLSCsrResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TLSCsrResponseData) SetID(v string) {
	o.ID = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TLSCsrResponseData) GetType() TypeTLSCsr {
	if o == nil || o.Type == nil {
		var ret TypeTLSCsr
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrResponseData) GetTypeOk() (*TypeTLSCsr, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TLSCsrResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSCsr and assigns it to the Type field.
func (o *TLSCsrResponseData) SetType(v TypeTLSCsr) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSCsrResponseData) GetAttributes() TLSCsrResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret TLSCsrResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrResponseData) GetAttributesOk() (*TLSCsrResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSCsrResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSCsrResponseAttributes and assigns it to the Attributes field.
func (o *TLSCsrResponseData) SetAttributes(v TLSCsrResponseAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TLSCsrResponseData) GetRelationships() RelationshipsForTLSCsr {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTLSCsr
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSCsrResponseData) GetRelationshipsOk() (*RelationshipsForTLSCsr, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TLSCsrResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTLSCsr and assigns it to the Relationships field.
func (o *TLSCsrResponseData) SetRelationships(v RelationshipsForTLSCsr) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSCsrResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
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
func (o *TLSCsrResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varTLSCsrResponseData := _TLSCsrResponseData{}

	if err = json.Unmarshal(bytes, &varTLSCsrResponseData); err == nil {
		*o = TLSCsrResponseData(varTLSCsrResponseData)
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

// NullableTLSCsrResponseData is a helper abstraction for handling nullable tlscsrresponsedata types. 
type NullableTLSCsrResponseData struct {
	value *TLSCsrResponseData
	isSet bool
}

// Get returns the value.
func (v NullableTLSCsrResponseData) Get() *TLSCsrResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSCsrResponseData) Set(val *TLSCsrResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSCsrResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSCsrResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSCsrResponseData returns a pointer to a new instance of NullableTLSCsrResponseData.
func NewNullableTLSCsrResponseData(val *TLSCsrResponseData) *NullableTLSCsrResponseData {
	return &NullableTLSCsrResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSCsrResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSCsrResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
