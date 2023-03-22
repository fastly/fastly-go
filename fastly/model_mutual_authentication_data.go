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

// MutualAuthenticationData struct for MutualAuthenticationData
type MutualAuthenticationData struct {
	Type *TypeMutualAuthentication `json:"type,omitempty"`
	Attributes *MutualAuthenticationDataAttributes `json:"attributes,omitempty"`
	Relationships *RelationshipsForMutualAuthentication `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _MutualAuthenticationData MutualAuthenticationData

// NewMutualAuthenticationData instantiates a new MutualAuthenticationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutualAuthenticationData() *MutualAuthenticationData {
	this := MutualAuthenticationData{}
	var resourceType TypeMutualAuthentication = TYPEMUTUALAUTHENTICATION_MUTUAL_AUTHENTICATION
	this.Type = &resourceType
	return &this
}

// NewMutualAuthenticationDataWithDefaults instantiates a new MutualAuthenticationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutualAuthenticationDataWithDefaults() *MutualAuthenticationData {
	this := MutualAuthenticationData{}
	var resourceType TypeMutualAuthentication = TYPEMUTUALAUTHENTICATION_MUTUAL_AUTHENTICATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MutualAuthenticationData) GetType() TypeMutualAuthentication {
	if o == nil || o.Type == nil {
		var ret TypeMutualAuthentication
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationData) GetTypeOk() (*TypeMutualAuthentication, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MutualAuthenticationData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeMutualAuthentication and assigns it to the Type field.
func (o *MutualAuthenticationData) SetType(v TypeMutualAuthentication) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MutualAuthenticationData) GetAttributes() MutualAuthenticationDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret MutualAuthenticationDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationData) GetAttributesOk() (*MutualAuthenticationDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MutualAuthenticationData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MutualAuthenticationDataAttributes and assigns it to the Attributes field.
func (o *MutualAuthenticationData) SetAttributes(v MutualAuthenticationDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *MutualAuthenticationData) GetRelationships() RelationshipsForMutualAuthentication {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForMutualAuthentication
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationData) GetRelationshipsOk() (*RelationshipsForMutualAuthentication, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *MutualAuthenticationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForMutualAuthentication and assigns it to the Relationships field.
func (o *MutualAuthenticationData) SetRelationships(v RelationshipsForMutualAuthentication) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o MutualAuthenticationData) MarshalJSON() ([]byte, error) {
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
func (o *MutualAuthenticationData) UnmarshalJSON(bytes []byte) (err error) {
	varMutualAuthenticationData := _MutualAuthenticationData{}

	if err = json.Unmarshal(bytes, &varMutualAuthenticationData); err == nil {
		*o = MutualAuthenticationData(varMutualAuthenticationData)
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

// NullableMutualAuthenticationData is a helper abstraction for handling nullable mutualauthenticationdata types. 
type NullableMutualAuthenticationData struct {
	value *MutualAuthenticationData
	isSet bool
}

// Get returns the value.
func (v NullableMutualAuthenticationData) Get() *MutualAuthenticationData {
	return v.value
}

// Set modifies the value.
func (v *NullableMutualAuthenticationData) Set(val *MutualAuthenticationData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableMutualAuthenticationData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableMutualAuthenticationData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableMutualAuthenticationData returns a pointer to a new instance of NullableMutualAuthenticationData.
func NewNullableMutualAuthenticationData(val *MutualAuthenticationData) *NullableMutualAuthenticationData {
	return &NullableMutualAuthenticationData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableMutualAuthenticationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableMutualAuthenticationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
