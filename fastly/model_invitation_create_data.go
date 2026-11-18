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

// InvitationCreateData struct for InvitationCreateData
type InvitationCreateData struct {
	Type                 *TypeInvitation                       `json:"type,omitempty"`
	Attributes           *InvitationDataAttributes             `json:"attributes,omitempty"`
	Relationships        *RelationshipServiceInvitationsCreate `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _InvitationCreateData InvitationCreateData

// NewInvitationCreateData instantiates a new InvitationCreateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationCreateData() *InvitationCreateData {
	this := InvitationCreateData{}
	var resourceType TypeInvitation = TYPEINVITATION_INVITATION
	this.Type = &resourceType
	return &this
}

// NewInvitationCreateDataWithDefaults instantiates a new InvitationCreateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationCreateDataWithDefaults() *InvitationCreateData {
	this := InvitationCreateData{}
	var resourceType TypeInvitation = TYPEINVITATION_INVITATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InvitationCreateData) GetType() TypeInvitation {
	if o == nil || o.Type == nil {
		var ret TypeInvitation
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationCreateData) GetTypeOk() (*TypeInvitation, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InvitationCreateData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeInvitation and assigns it to the Type field.
func (o *InvitationCreateData) SetType(v TypeInvitation) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InvitationCreateData) GetAttributes() InvitationDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret InvitationDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationCreateData) GetAttributesOk() (*InvitationDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InvitationCreateData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given InvitationDataAttributes and assigns it to the Attributes field.
func (o *InvitationCreateData) SetAttributes(v InvitationDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InvitationCreateData) GetRelationships() RelationshipServiceInvitationsCreate {
	if o == nil || o.Relationships == nil {
		var ret RelationshipServiceInvitationsCreate
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationCreateData) GetRelationshipsOk() (*RelationshipServiceInvitationsCreate, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InvitationCreateData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipServiceInvitationsCreate and assigns it to the Relationships field.
func (o *InvitationCreateData) SetRelationships(v RelationshipServiceInvitationsCreate) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InvitationCreateData) MarshalJSON() ([]byte, error) {
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
func (o *InvitationCreateData) UnmarshalJSON(bytes []byte) (err error) {
	varInvitationCreateData := _InvitationCreateData{}

	if err = json.Unmarshal(bytes, &varInvitationCreateData); err == nil {
		*o = InvitationCreateData(varInvitationCreateData)
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

// NullableInvitationCreateData is a helper abstraction for handling nullable invitationcreatedata types.
type NullableInvitationCreateData struct {
	value *InvitationCreateData
	isSet bool
}

// Get returns the value.
func (v NullableInvitationCreateData) Get() *InvitationCreateData {
	return v.value
}

// Set modifies the value.
func (v *NullableInvitationCreateData) Set(val *InvitationCreateData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInvitationCreateData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInvitationCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInvitationCreateData returns a pointer to a new instance of NullableInvitationCreateData.
func NewNullableInvitationCreateData(val *InvitationCreateData) *NullableInvitationCreateData {
	return &NullableInvitationCreateData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInvitationCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInvitationCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
