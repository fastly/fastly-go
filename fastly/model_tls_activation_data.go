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

// TlsActivationData struct for TlsActivationData
type TlsActivationData struct {
	Type                 *TypeTlsActivation             `json:"type,omitempty"`
	Relationships        *RelationshipsForTlsActivation `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsActivationData TlsActivationData

// NewTlsActivationData instantiates a new TlsActivationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsActivationData() *TlsActivationData {
	this := TlsActivationData{}
	var type_ TypeTlsActivation = TYPETLSACTIVATION_TLS_ACTIVATION
	this.Type = &type_
	return &this
}

// NewTlsActivationDataWithDefaults instantiates a new TlsActivationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsActivationDataWithDefaults() *TlsActivationData {
	this := TlsActivationData{}
	var type_ TypeTlsActivation = TYPETLSACTIVATION_TLS_ACTIVATION
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TlsActivationData) GetType() TypeTlsActivation {
	if o == nil || o.Type == nil {
		var ret TypeTlsActivation
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsActivationData) GetTypeOk() (*TypeTlsActivation, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TlsActivationData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsActivation and assigns it to the Type field.
func (o *TlsActivationData) SetType(v TypeTlsActivation) {
	o.Type = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TlsActivationData) GetRelationships() RelationshipsForTlsActivation {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTlsActivation
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsActivationData) GetRelationshipsOk() (*RelationshipsForTlsActivation, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TlsActivationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTlsActivation and assigns it to the Relationships field.
func (o *TlsActivationData) SetRelationships(v RelationshipsForTlsActivation) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsActivationData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
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
func (o *TlsActivationData) UnmarshalJSON(bytes []byte) (err error) {
	varTlsActivationData := _TlsActivationData{}

	if err = json.Unmarshal(bytes, &varTlsActivationData); err == nil {
		*o = TlsActivationData(varTlsActivationData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsActivationData is a helper abstraction for handling nullable tlsactivationdata types.
type NullableTlsActivationData struct {
	value *TlsActivationData
	isSet bool
}

// Get returns the value.
func (v NullableTlsActivationData) Get() *TlsActivationData {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsActivationData) Set(val *TlsActivationData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsActivationData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsActivationData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsActivationData returns a pointer to a new instance of NullableTlsActivationData.
func NewNullableTlsActivationData(val *TlsActivationData) *NullableTlsActivationData {
	return &NullableTlsActivationData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsActivationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsActivationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
