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

// TLSPrivateKeyData struct for TLSPrivateKeyData
type TLSPrivateKeyData struct {
	Type *TypeTLSPrivateKey `json:"type,omitempty"`
	Attributes *TLSPrivateKeyDataAttributes `json:"attributes,omitempty"`
	Relationships *RelationshipsForTLSPrivateKey `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSPrivateKeyData TLSPrivateKeyData

// NewTLSPrivateKeyData instantiates a new TLSPrivateKeyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSPrivateKeyData() *TLSPrivateKeyData {
	this := TLSPrivateKeyData{}
	var resourceType TypeTLSPrivateKey = TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY
	this.Type = &resourceType
	return &this
}

// NewTLSPrivateKeyDataWithDefaults instantiates a new TLSPrivateKeyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSPrivateKeyDataWithDefaults() *TLSPrivateKeyData {
	this := TLSPrivateKeyData{}
	var resourceType TypeTLSPrivateKey = TYPETLSPRIVATEKEY_TLS_PRIVATE_KEY
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TLSPrivateKeyData) GetType() TypeTLSPrivateKey {
	if o == nil || o.Type == nil {
		var ret TypeTLSPrivateKey
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyData) GetTypeOk() (*TypeTLSPrivateKey, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TLSPrivateKeyData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSPrivateKey and assigns it to the Type field.
func (o *TLSPrivateKeyData) SetType(v TypeTLSPrivateKey) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSPrivateKeyData) GetAttributes() TLSPrivateKeyDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TLSPrivateKeyDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyData) GetAttributesOk() (*TLSPrivateKeyDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSPrivateKeyData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSPrivateKeyDataAttributes and assigns it to the Attributes field.
func (o *TLSPrivateKeyData) SetAttributes(v TLSPrivateKeyDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TLSPrivateKeyData) GetRelationships() RelationshipsForTLSPrivateKey {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTLSPrivateKey
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSPrivateKeyData) GetRelationshipsOk() (*RelationshipsForTLSPrivateKey, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TLSPrivateKeyData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTLSPrivateKey and assigns it to the Relationships field.
func (o *TLSPrivateKeyData) SetRelationships(v RelationshipsForTLSPrivateKey) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSPrivateKeyData) MarshalJSON() ([]byte, error) {
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
func (o *TLSPrivateKeyData) UnmarshalJSON(bytes []byte) (err error) {
	varTLSPrivateKeyData := _TLSPrivateKeyData{}

	if err = json.Unmarshal(bytes, &varTLSPrivateKeyData); err == nil {
		*o = TLSPrivateKeyData(varTLSPrivateKeyData)
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

// NullableTLSPrivateKeyData is a helper abstraction for handling nullable tlsprivatekeydata types. 
type NullableTLSPrivateKeyData struct {
	value *TLSPrivateKeyData
	isSet bool
}

// Get returns the value.
func (v NullableTLSPrivateKeyData) Get() *TLSPrivateKeyData {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSPrivateKeyData) Set(val *TLSPrivateKeyData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSPrivateKeyData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSPrivateKeyData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSPrivateKeyData returns a pointer to a new instance of NullableTLSPrivateKeyData.
func NewNullableTLSPrivateKeyData(val *TLSPrivateKeyData) *NullableTLSPrivateKeyData {
	return &NullableTLSPrivateKeyData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSPrivateKeyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSPrivateKeyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
