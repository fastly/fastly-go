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

// TLSConfigurationResponseData struct for TLSConfigurationResponseData
type TLSConfigurationResponseData struct {
	Type                 *TypeTLSConfiguration               `json:"type,omitempty"`
	Attributes           *TLSConfigurationResponseAttributes `json:"attributes,omitempty"`
	Relationships        *RelationshipsForTLSConfiguration   `json:"relationships,omitempty"`
	ID                   *string                             `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSConfigurationResponseData TLSConfigurationResponseData

// NewTLSConfigurationResponseData instantiates a new TLSConfigurationResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSConfigurationResponseData() *TLSConfigurationResponseData {
	this := TLSConfigurationResponseData{}
	var resourceType TypeTLSConfiguration = TYPETLSCONFIGURATION_TLS_CONFIGURATION
	this.Type = &resourceType
	return &this
}

// NewTLSConfigurationResponseDataWithDefaults instantiates a new TLSConfigurationResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSConfigurationResponseDataWithDefaults() *TLSConfigurationResponseData {
	this := TLSConfigurationResponseData{}
	var resourceType TypeTLSConfiguration = TYPETLSCONFIGURATION_TLS_CONFIGURATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TLSConfigurationResponseData) GetType() TypeTLSConfiguration {
	if o == nil || o.Type == nil {
		var ret TypeTLSConfiguration
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationResponseData) GetTypeOk() (*TypeTLSConfiguration, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TLSConfigurationResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSConfiguration and assigns it to the Type field.
func (o *TLSConfigurationResponseData) SetType(v TypeTLSConfiguration) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSConfigurationResponseData) GetAttributes() TLSConfigurationResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret TLSConfigurationResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationResponseData) GetAttributesOk() (*TLSConfigurationResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSConfigurationResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given TLSConfigurationResponseAttributes and assigns it to the Attributes field.
func (o *TLSConfigurationResponseData) SetAttributes(v TLSConfigurationResponseAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TLSConfigurationResponseData) GetRelationships() RelationshipsForTLSConfiguration {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTLSConfiguration
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationResponseData) GetRelationshipsOk() (*RelationshipsForTLSConfiguration, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TLSConfigurationResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTLSConfiguration and assigns it to the Relationships field.
func (o *TLSConfigurationResponseData) SetRelationships(v RelationshipsForTLSConfiguration) {
	o.Relationships = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TLSConfigurationResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfigurationResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TLSConfigurationResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TLSConfigurationResponseData) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSConfigurationResponseData) MarshalJSON() ([]byte, error) {
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
func (o *TLSConfigurationResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varTLSConfigurationResponseData := _TLSConfigurationResponseData{}

	if err = json.Unmarshal(bytes, &varTLSConfigurationResponseData); err == nil {
		*o = TLSConfigurationResponseData(varTLSConfigurationResponseData)
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

// NullableTLSConfigurationResponseData is a helper abstraction for handling nullable tlsconfigurationresponsedata types.
type NullableTLSConfigurationResponseData struct {
	value *TLSConfigurationResponseData
	isSet bool
}

// Get returns the value.
func (v NullableTLSConfigurationResponseData) Get() *TLSConfigurationResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSConfigurationResponseData) Set(val *TLSConfigurationResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSConfigurationResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSConfigurationResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSConfigurationResponseData returns a pointer to a new instance of NullableTLSConfigurationResponseData.
func NewNullableTLSConfigurationResponseData(val *TLSConfigurationResponseData) *NullableTLSConfigurationResponseData {
	return &NullableTLSConfigurationResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSConfigurationResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTLSConfigurationResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
