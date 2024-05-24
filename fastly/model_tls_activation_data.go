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

// TLSActivationData struct for TLSActivationData
type TLSActivationData struct {
	Type *TypeTLSActivation `json:"type,omitempty"`
	Relationships *RelationshipsForTLSActivation `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSActivationData TLSActivationData

// NewTLSActivationData instantiates a new TLSActivationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSActivationData() *TLSActivationData {
	this := TLSActivationData{}
	var resourceType TypeTLSActivation = TYPETLSACTIVATION_TLS_ACTIVATION
	this.Type = &resourceType
	return &this
}

// NewTLSActivationDataWithDefaults instantiates a new TLSActivationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSActivationDataWithDefaults() *TLSActivationData {
	this := TLSActivationData{}
	var resourceType TypeTLSActivation = TYPETLSACTIVATION_TLS_ACTIVATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TLSActivationData) GetType() TypeTLSActivation {
	if o == nil || o.Type == nil {
		var ret TypeTLSActivation
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSActivationData) GetTypeOk() (*TypeTLSActivation, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TLSActivationData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSActivation and assigns it to the Type field.
func (o *TLSActivationData) SetType(v TypeTLSActivation) {
	o.Type = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TLSActivationData) GetRelationships() RelationshipsForTLSActivation {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTLSActivation
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSActivationData) GetRelationshipsOk() (*RelationshipsForTLSActivation, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TLSActivationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTLSActivation and assigns it to the Relationships field.
func (o *TLSActivationData) SetRelationships(v RelationshipsForTLSActivation) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSActivationData) MarshalJSON() ([]byte, error) {
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
func (o *TLSActivationData) UnmarshalJSON(bytes []byte) (err error) {
	varTLSActivationData := _TLSActivationData{}

	if err = json.Unmarshal(bytes, &varTLSActivationData); err == nil {
		*o = TLSActivationData(varTLSActivationData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSActivationData is a helper abstraction for handling nullable tlsactivationdata types. 
type NullableTLSActivationData struct {
	value *TLSActivationData
	isSet bool
}

// Get returns the value.
func (v NullableTLSActivationData) Get() *TLSActivationData {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSActivationData) Set(val *TLSActivationData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSActivationData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSActivationData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSActivationData returns a pointer to a new instance of NullableTLSActivationData.
func NewNullableTLSActivationData(val *TLSActivationData) *NullableTLSActivationData {
	return &NullableTLSActivationData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSActivationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSActivationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
