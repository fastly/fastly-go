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

// InvitationCreateDataAllOf struct for InvitationCreateDataAllOf
type InvitationCreateDataAllOf struct {
	Relationships        *RelationshipServiceInvitationsCreate `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _InvitationCreateDataAllOf InvitationCreateDataAllOf

// NewInvitationCreateDataAllOf instantiates a new InvitationCreateDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationCreateDataAllOf() *InvitationCreateDataAllOf {
	this := InvitationCreateDataAllOf{}
	return &this
}

// NewInvitationCreateDataAllOfWithDefaults instantiates a new InvitationCreateDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationCreateDataAllOfWithDefaults() *InvitationCreateDataAllOf {
	this := InvitationCreateDataAllOf{}
	return &this
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InvitationCreateDataAllOf) GetRelationships() RelationshipServiceInvitationsCreate {
	if o == nil || o.Relationships == nil {
		var ret RelationshipServiceInvitationsCreate
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationCreateDataAllOf) GetRelationshipsOk() (*RelationshipServiceInvitationsCreate, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InvitationCreateDataAllOf) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipServiceInvitationsCreate and assigns it to the Relationships field.
func (o *InvitationCreateDataAllOf) SetRelationships(v RelationshipServiceInvitationsCreate) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InvitationCreateDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *InvitationCreateDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInvitationCreateDataAllOf := _InvitationCreateDataAllOf{}

	if err = json.Unmarshal(bytes, &varInvitationCreateDataAllOf); err == nil {
		*o = InvitationCreateDataAllOf(varInvitationCreateDataAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInvitationCreateDataAllOf is a helper abstraction for handling nullable invitationcreatedataallof types.
type NullableInvitationCreateDataAllOf struct {
	value *InvitationCreateDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableInvitationCreateDataAllOf) Get() *InvitationCreateDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableInvitationCreateDataAllOf) Set(val *InvitationCreateDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInvitationCreateDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInvitationCreateDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInvitationCreateDataAllOf returns a pointer to a new instance of NullableInvitationCreateDataAllOf.
func NewNullableInvitationCreateDataAllOf(val *InvitationCreateDataAllOf) *NullableInvitationCreateDataAllOf {
	return &NullableInvitationCreateDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInvitationCreateDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInvitationCreateDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
