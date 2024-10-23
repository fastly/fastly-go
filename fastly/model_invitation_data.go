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

// InvitationData struct for InvitationData
type InvitationData struct {
	Type                 *TypeInvitation                       `json:"type,omitempty"`
	Attributes           *InvitationDataAttributes             `json:"attributes,omitempty"`
	Relationships        *RelationshipServiceInvitationsCreate `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _InvitationData InvitationData

// NewInvitationData instantiates a new InvitationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationData() *InvitationData {
	this := InvitationData{}
	var resourceType TypeInvitation = TYPEINVITATION_INVITATION
	this.Type = &resourceType
	return &this
}

// NewInvitationDataWithDefaults instantiates a new InvitationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationDataWithDefaults() *InvitationData {
	this := InvitationData{}
	var resourceType TypeInvitation = TYPEINVITATION_INVITATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *InvitationData) GetType() TypeInvitation {
	if o == nil || o.Type == nil {
		var ret TypeInvitation
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationData) GetTypeOk() (*TypeInvitation, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *InvitationData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeInvitation and assigns it to the Type field.
func (o *InvitationData) SetType(v TypeInvitation) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InvitationData) GetAttributes() InvitationDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret InvitationDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationData) GetAttributesOk() (*InvitationDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InvitationData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given InvitationDataAttributes and assigns it to the Attributes field.
func (o *InvitationData) SetAttributes(v InvitationDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InvitationData) GetRelationships() RelationshipServiceInvitationsCreate {
	if o == nil || o.Relationships == nil {
		var ret RelationshipServiceInvitationsCreate
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationData) GetRelationshipsOk() (*RelationshipServiceInvitationsCreate, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InvitationData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipServiceInvitationsCreate and assigns it to the Relationships field.
func (o *InvitationData) SetRelationships(v RelationshipServiceInvitationsCreate) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InvitationData) MarshalJSON() ([]byte, error) {
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
func (o *InvitationData) UnmarshalJSON(bytes []byte) (err error) {
	varInvitationData := _InvitationData{}

	if err = json.Unmarshal(bytes, &varInvitationData); err == nil {
		*o = InvitationData(varInvitationData)
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

// NullableInvitationData is a helper abstraction for handling nullable invitationdata types.
type NullableInvitationData struct {
	value *InvitationData
	isSet bool
}

// Get returns the value.
func (v NullableInvitationData) Get() *InvitationData {
	return v.value
}

// Set modifies the value.
func (v *NullableInvitationData) Set(val *InvitationData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInvitationData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInvitationData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInvitationData returns a pointer to a new instance of NullableInvitationData.
func NewNullableInvitationData(val *InvitationData) *NullableInvitationData {
	return &NullableInvitationData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInvitationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInvitationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
