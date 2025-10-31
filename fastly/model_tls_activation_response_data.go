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

// TlsActivationResponseData struct for TlsActivationResponseData
type TlsActivationResponseData struct {
	Type                 *TypeTlsActivation             `json:"type,omitempty"`
	Relationships        *RelationshipsForTlsActivation `json:"relationships,omitempty"`
	Id                   *string                        `json:"id,omitempty"`
	Attributes           *Timestamps                    `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsActivationResponseData TlsActivationResponseData

// NewTlsActivationResponseData instantiates a new TlsActivationResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsActivationResponseData() *TlsActivationResponseData {
	this := TlsActivationResponseData{}
	var type_ TypeTlsActivation = TYPETLSACTIVATION_TLS_ACTIVATION
	this.Type = &type_
	return &this
}

// NewTlsActivationResponseDataWithDefaults instantiates a new TlsActivationResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsActivationResponseDataWithDefaults() *TlsActivationResponseData {
	this := TlsActivationResponseData{}
	var type_ TypeTlsActivation = TYPETLSACTIVATION_TLS_ACTIVATION
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TlsActivationResponseData) GetType() TypeTlsActivation {
	if o == nil || o.Type == nil {
		var ret TypeTlsActivation
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsActivationResponseData) GetTypeOk() (*TypeTlsActivation, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TlsActivationResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsActivation and assigns it to the Type field.
func (o *TlsActivationResponseData) SetType(v TypeTlsActivation) {
	o.Type = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TlsActivationResponseData) GetRelationships() RelationshipsForTlsActivation {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTlsActivation
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsActivationResponseData) GetRelationshipsOk() (*RelationshipsForTlsActivation, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TlsActivationResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTlsActivation and assigns it to the Relationships field.
func (o *TlsActivationResponseData) SetRelationships(v RelationshipsForTlsActivation) {
	o.Relationships = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TlsActivationResponseData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsActivationResponseData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TlsActivationResponseData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TlsActivationResponseData) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TlsActivationResponseData) GetAttributes() Timestamps {
	if o == nil || o.Attributes == nil {
		var ret Timestamps
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsActivationResponseData) GetAttributesOk() (*Timestamps, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TlsActivationResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Timestamps and assigns it to the Attributes field.
func (o *TlsActivationResponseData) SetAttributes(v Timestamps) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsActivationResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsActivationResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varTlsActivationResponseData := _TlsActivationResponseData{}

	if err = json.Unmarshal(bytes, &varTlsActivationResponseData); err == nil {
		*o = TlsActivationResponseData(varTlsActivationResponseData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsActivationResponseData is a helper abstraction for handling nullable tlsactivationresponsedata types.
type NullableTlsActivationResponseData struct {
	value *TlsActivationResponseData
	isSet bool
}

// Get returns the value.
func (v NullableTlsActivationResponseData) Get() *TlsActivationResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsActivationResponseData) Set(val *TlsActivationResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsActivationResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsActivationResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsActivationResponseData returns a pointer to a new instance of NullableTlsActivationResponseData.
func NewNullableTlsActivationResponseData(val *TlsActivationResponseData) *NullableTlsActivationResponseData {
	return &NullableTlsActivationResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsActivationResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsActivationResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
