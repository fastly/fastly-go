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

// TLSConfigurationData struct for TLSConfigurationData
type TLSConfigurationData struct {
	Type *TypeTLSConfiguration `json:"type,omitempty"`
	Attributes *TLSConfigurationDataAttributes `json:"attributes,omitempty"`
	Relationships *RelationshipsForTLSConfiguration `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSConfigurationData TLSConfigurationData

// NewTLSConfigurationData instantiates a new TLSConfigurationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSConfigurationData() *TLSConfigurationData {
	this := TLSConfigurationData{}
	var resourceType TypeTLSConfiguration = TYPETLSCONFIGURATION_TLS_CONFIGURATION
	this.Type = &resourceType
	return &this
}

// NewTLSConfigurationDataWithDefaults instantiates a new TLSConfigurationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSConfigurationDataWithDefaults() *TLSConfigurationData {
	this := TLSConfigurationData{}
	var resourceType TypeTLSConfiguration = TYPETLSCONFIGURATION_TLS_CONFIGURATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TLSConfigurationData) GetType() TypeTLSConfiguration {
	if o == nil || o.Type == nil {
		var ret TypeTLSConfiguration
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationData) GetTypeOk() (*TypeTLSConfiguration, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TLSConfigurationData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSConfiguration and assigns it to the Type field.
func (o *TLSConfigurationData) SetType(v TypeTLSConfiguration) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSConfigurationData) GetAttributes() TLSConfigurationDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TLSConfigurationDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationData) GetAttributesOk() (*TLSConfigurationDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSConfigurationData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSConfigurationDataAttributes and assigns it to the Attributes field.
func (o *TLSConfigurationData) SetAttributes(v TLSConfigurationDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TLSConfigurationData) GetRelationships() RelationshipsForTLSConfiguration {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTLSConfiguration
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationData) GetRelationshipsOk() (*RelationshipsForTLSConfiguration, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TLSConfigurationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTLSConfiguration and assigns it to the Relationships field.
func (o *TLSConfigurationData) SetRelationships(v RelationshipsForTLSConfiguration) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSConfigurationData) MarshalJSON() ([]byte, error) {
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
func (o *TLSConfigurationData) UnmarshalJSON(bytes []byte) (err error) {
	varTLSConfigurationData := _TLSConfigurationData{}

	if err = json.Unmarshal(bytes, &varTLSConfigurationData); err == nil {
		*o = TLSConfigurationData(varTLSConfigurationData)
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

// NullableTLSConfigurationData is a helper abstraction for handling nullable tlsconfigurationdata types. 
type NullableTLSConfigurationData struct {
	value *TLSConfigurationData
	isSet bool
}

// Get returns the value.
func (v NullableTLSConfigurationData) Get() *TLSConfigurationData {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSConfigurationData) Set(val *TLSConfigurationData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSConfigurationData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSConfigurationData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSConfigurationData returns a pointer to a new instance of NullableTLSConfigurationData.
func NewNullableTLSConfigurationData(val *TLSConfigurationData) *NullableTLSConfigurationData {
	return &NullableTLSConfigurationData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSConfigurationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSConfigurationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
