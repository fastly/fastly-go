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

// InvitationResponseDataAllOf struct for InvitationResponseDataAllOf
type InvitationResponseDataAllOf struct {
	Id                   *string                     `json:"id,omitempty"`
	Attributes           *Timestamps                 `json:"attributes,omitempty"`
	Relationships        *RelationshipsForInvitation `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _InvitationResponseDataAllOf InvitationResponseDataAllOf

// NewInvitationResponseDataAllOf instantiates a new InvitationResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvitationResponseDataAllOf() *InvitationResponseDataAllOf {
	this := InvitationResponseDataAllOf{}
	return &this
}

// NewInvitationResponseDataAllOfWithDefaults instantiates a new InvitationResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvitationResponseDataAllOfWithDefaults() *InvitationResponseDataAllOf {
	this := InvitationResponseDataAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InvitationResponseDataAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationResponseDataAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InvitationResponseDataAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *InvitationResponseDataAllOf) SetId(v string) {
	o.Id = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *InvitationResponseDataAllOf) GetAttributes() Timestamps {
	if o == nil || o.Attributes == nil {
		var ret Timestamps
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationResponseDataAllOf) GetAttributesOk() (*Timestamps, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *InvitationResponseDataAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given Timestamps and assigns it to the Attributes field.
func (o *InvitationResponseDataAllOf) SetAttributes(v Timestamps) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *InvitationResponseDataAllOf) GetRelationships() RelationshipsForInvitation {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForInvitation
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvitationResponseDataAllOf) GetRelationshipsOk() (*RelationshipsForInvitation, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *InvitationResponseDataAllOf) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForInvitation and assigns it to the Relationships field.
func (o *InvitationResponseDataAllOf) SetRelationships(v RelationshipsForInvitation) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o InvitationResponseDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
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
func (o *InvitationResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varInvitationResponseDataAllOf := _InvitationResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varInvitationResponseDataAllOf); err == nil {
		*o = InvitationResponseDataAllOf(varInvitationResponseDataAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableInvitationResponseDataAllOf is a helper abstraction for handling nullable invitationresponsedataallof types.
type NullableInvitationResponseDataAllOf struct {
	value *InvitationResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableInvitationResponseDataAllOf) Get() *InvitationResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableInvitationResponseDataAllOf) Set(val *InvitationResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableInvitationResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableInvitationResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableInvitationResponseDataAllOf returns a pointer to a new instance of NullableInvitationResponseDataAllOf.
func NewNullableInvitationResponseDataAllOf(val *InvitationResponseDataAllOf) *NullableInvitationResponseDataAllOf {
	return &NullableInvitationResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableInvitationResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableInvitationResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
