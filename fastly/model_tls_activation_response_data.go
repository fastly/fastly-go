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

// TLSActivationResponseData struct for TLSActivationResponseData
type TLSActivationResponseData struct {
	Type *TypeTLSActivation `json:"type,omitempty"`
	Relationships *RelationshipsForTLSActivation `json:"relationships,omitempty"`
	ID *string `json:"id,omitempty"`
	Attributes *Timestamps `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSActivationResponseData TLSActivationResponseData

// NewTLSActivationResponseData instantiates a new TLSActivationResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSActivationResponseData() *TLSActivationResponseData {
	this := TLSActivationResponseData{}
	var resourceType TypeTLSActivation = TYPETLSACTIVATION_TLS_ACTIVATION
	this.Type = &resourceType
	return &this
}

// NewTLSActivationResponseDataWithDefaults instantiates a new TLSActivationResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSActivationResponseDataWithDefaults() *TLSActivationResponseData {
	this := TLSActivationResponseData{}
	var resourceType TypeTLSActivation = TYPETLSACTIVATION_TLS_ACTIVATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *TLSActivationResponseData) GetType() TypeTLSActivation {
	if o == nil || o.Type == nil {
		var ret TypeTLSActivation
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSActivationResponseData) GetTypeOk() (*TypeTLSActivation, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *TLSActivationResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSActivation and assigns it to the Type field.
func (o *TLSActivationResponseData) SetType(v TypeTLSActivation) {
	o.Type = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *TLSActivationResponseData) GetRelationships() RelationshipsForTLSActivation {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForTLSActivation
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSActivationResponseData) GetRelationshipsOk() (*RelationshipsForTLSActivation, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *TLSActivationResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForTLSActivation and assigns it to the Relationships field.
func (o *TLSActivationResponseData) SetRelationships(v RelationshipsForTLSActivation) {
	o.Relationships = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TLSActivationResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSActivationResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TLSActivationResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TLSActivationResponseData) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *TLSActivationResponseData) GetAttributes() Timestamps {
	if o == nil || o.Attributes == nil {
		var ret Timestamps
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSActivationResponseData) GetAttributesOk() (*Timestamps, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *TLSActivationResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Timestamps and assigns it to the Attributes field.
func (o *TLSActivationResponseData) SetAttributes(v Timestamps) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSActivationResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
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
func (o *TLSActivationResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varTLSActivationResponseData := _TLSActivationResponseData{}

	if err = json.Unmarshal(bytes, &varTLSActivationResponseData); err == nil {
		*o = TLSActivationResponseData(varTLSActivationResponseData)
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

// NullableTLSActivationResponseData is a helper abstraction for handling nullable tlsactivationresponsedata types. 
type NullableTLSActivationResponseData struct {
	value *TLSActivationResponseData
	isSet bool
}

// Get returns the value.
func (v NullableTLSActivationResponseData) Get() *TLSActivationResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSActivationResponseData) Set(val *TLSActivationResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSActivationResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSActivationResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSActivationResponseData returns a pointer to a new instance of NullableTLSActivationResponseData.
func NewNullableTLSActivationResponseData(val *TLSActivationResponseData) *NullableTLSActivationResponseData {
	return &NullableTLSActivationResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSActivationResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSActivationResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
