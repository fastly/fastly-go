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

// TlsConfigurationData struct for TlsConfigurationData
type TlsConfigurationData struct {
	Type                 *TypeTlsConfiguration             `json:"type,omitempty"`
	Attributes           *TlsConfigurationDataAttributes   `json:"attributes,omitempty"`
	Relationships        *RelationshipsForTlsConfiguration `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsConfigurationData TlsConfigurationData

// NewTlsConfigurationData instantiates a new TlsConfigurationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsConfigurationData() *TlsConfigurationData {
	this := TlsConfigurationData{}
	var type_ TypeTlsConfiguration = TYPETLSCONFIGURATION_TLS_CONFIGURATION
	this.Type = &type_
	return &this
}

// NewTlsConfigurationDataWithDefaults instantiates a new TlsConfigurationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsConfigurationDataWithDefaults() *TlsConfigurationData {
	this := TlsConfigurationData{}
	var type_ TypeTlsConfiguration = TYPETLSCONFIGURATION_TLS_CONFIGURATION
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TlsConfigurationData) GetType() TypeTlsConfiguration {
	if o == nil || o.Type == nil {
		var ret TypeTlsConfiguration
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationData) GetTypeOk() (*TypeTlsConfiguration, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TlsConfigurationData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsConfiguration and assigns it to the Type field.
func (o *TlsConfigurationData) SetType(v TypeTlsConfiguration) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TlsConfigurationData) GetAttributes() TlsConfigurationDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret TlsConfigurationDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationData) GetAttributesOk() (*TlsConfigurationDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TlsConfigurationData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TlsConfigurationDataAttributes and assigns it to the Attributes field.
func (o *TlsConfigurationData) SetAttributes(v TlsConfigurationDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TlsConfigurationData) GetRelationships() RelationshipsForTlsConfiguration {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTlsConfiguration
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsConfigurationData) GetRelationshipsOk() (*RelationshipsForTlsConfiguration, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TlsConfigurationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTlsConfiguration and assigns it to the Relationships field.
func (o *TlsConfigurationData) SetRelationships(v RelationshipsForTlsConfiguration) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsConfigurationData) MarshalJSON() ([]byte, error) {
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
func (o *TlsConfigurationData) UnmarshalJSON(bytes []byte) (err error) {
	varTlsConfigurationData := _TlsConfigurationData{}

	if err = json.Unmarshal(bytes, &varTlsConfigurationData); err == nil {
		*o = TlsConfigurationData(varTlsConfigurationData)
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

// NullableTlsConfigurationData is a helper abstraction for handling nullable tlsconfigurationdata types.
type NullableTlsConfigurationData struct {
	value *TlsConfigurationData
	isSet bool
}

// Get returns the value.
func (v NullableTlsConfigurationData) Get() *TlsConfigurationData {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsConfigurationData) Set(val *TlsConfigurationData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsConfigurationData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsConfigurationData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsConfigurationData returns a pointer to a new instance of NullableTlsConfigurationData.
func NewNullableTlsConfigurationData(val *TlsConfigurationData) *NullableTlsConfigurationData {
	return &NullableTlsConfigurationData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsConfigurationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsConfigurationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
