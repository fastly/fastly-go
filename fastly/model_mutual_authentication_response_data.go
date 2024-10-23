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

// MutualAuthenticationResponseData struct for MutualAuthenticationResponseData
type MutualAuthenticationResponseData struct {
	Type                 *TypeMutualAuthentication               `json:"type,omitempty"`
	Attributes           *MutualAuthenticationResponseAttributes `json:"attributes,omitempty"`
	Relationships        *RelationshipsForMutualAuthentication   `json:"relationships,omitempty"`
	ID                   *string                                 `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _MutualAuthenticationResponseData MutualAuthenticationResponseData

// NewMutualAuthenticationResponseData instantiates a new MutualAuthenticationResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMutualAuthenticationResponseData() *MutualAuthenticationResponseData {
	this := MutualAuthenticationResponseData{}
	var resourceType TypeMutualAuthentication = TYPEMUTUALAUTHENTICATION_MUTUAL_AUTHENTICATION
	this.Type = &resourceType
	return &this
}

// NewMutualAuthenticationResponseDataWithDefaults instantiates a new MutualAuthenticationResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMutualAuthenticationResponseDataWithDefaults() *MutualAuthenticationResponseData {
	this := MutualAuthenticationResponseData{}
	var resourceType TypeMutualAuthentication = TYPEMUTUALAUTHENTICATION_MUTUAL_AUTHENTICATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MutualAuthenticationResponseData) GetType() TypeMutualAuthentication {
	if o == nil || o.Type == nil {
		var ret TypeMutualAuthentication
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationResponseData) GetTypeOk() (*TypeMutualAuthentication, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MutualAuthenticationResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeMutualAuthentication and assigns it to the Type field.
func (o *MutualAuthenticationResponseData) SetType(v TypeMutualAuthentication) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *MutualAuthenticationResponseData) GetAttributes() MutualAuthenticationResponseAttributes {
	if o == nil || o.Attributes == nil {
		var ret MutualAuthenticationResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationResponseData) GetAttributesOk() (*MutualAuthenticationResponseAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *MutualAuthenticationResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given MutualAuthenticationResponseAttributes and assigns it to the Attributes field.
func (o *MutualAuthenticationResponseData) SetAttributes(v MutualAuthenticationResponseAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *MutualAuthenticationResponseData) GetRelationships() RelationshipsForMutualAuthentication {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForMutualAuthentication
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationResponseData) GetRelationshipsOk() (*RelationshipsForMutualAuthentication, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *MutualAuthenticationResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForMutualAuthentication and assigns it to the Relationships field.
func (o *MutualAuthenticationResponseData) SetRelationships(v RelationshipsForMutualAuthentication) {
	o.Relationships = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *MutualAuthenticationResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MutualAuthenticationResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *MutualAuthenticationResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *MutualAuthenticationResponseData) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o MutualAuthenticationResponseData) MarshalJSON() ([]byte, error) {
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
func (o *MutualAuthenticationResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varMutualAuthenticationResponseData := _MutualAuthenticationResponseData{}

	if err = json.Unmarshal(bytes, &varMutualAuthenticationResponseData); err == nil {
		*o = MutualAuthenticationResponseData(varMutualAuthenticationResponseData)
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

// NullableMutualAuthenticationResponseData is a helper abstraction for handling nullable mutualauthenticationresponsedata types.
type NullableMutualAuthenticationResponseData struct {
	value *MutualAuthenticationResponseData
	isSet bool
}

// Get returns the value.
func (v NullableMutualAuthenticationResponseData) Get() *MutualAuthenticationResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableMutualAuthenticationResponseData) Set(val *MutualAuthenticationResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableMutualAuthenticationResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableMutualAuthenticationResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableMutualAuthenticationResponseData returns a pointer to a new instance of NullableMutualAuthenticationResponseData.
func NewNullableMutualAuthenticationResponseData(val *MutualAuthenticationResponseData) *NullableMutualAuthenticationResponseData {
	return &NullableMutualAuthenticationResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableMutualAuthenticationResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableMutualAuthenticationResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
