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

// RelationshipServiceInvitations struct for RelationshipServiceInvitations
type RelationshipServiceInvitations struct {
	ServiceInvitations   NullableRelationshipServiceInvitationsServiceInvitations `json:"service_invitations,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipServiceInvitations RelationshipServiceInvitations

// NewRelationshipServiceInvitations instantiates a new RelationshipServiceInvitations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipServiceInvitations() *RelationshipServiceInvitations {
	this := RelationshipServiceInvitations{}
	return &this
}

// NewRelationshipServiceInvitationsWithDefaults instantiates a new RelationshipServiceInvitations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipServiceInvitationsWithDefaults() *RelationshipServiceInvitations {
	this := RelationshipServiceInvitations{}
	return &this
}

// GetServiceInvitations returns the ServiceInvitations field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RelationshipServiceInvitations) GetServiceInvitations() RelationshipServiceInvitationsServiceInvitations {
	if o == nil || o.ServiceInvitations.Get() == nil {
		var ret RelationshipServiceInvitationsServiceInvitations
		return ret
	}
	return *o.ServiceInvitations.Get()
}

// GetServiceInvitationsOk returns a tuple with the ServiceInvitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RelationshipServiceInvitations) GetServiceInvitationsOk() (*RelationshipServiceInvitationsServiceInvitations, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceInvitations.Get(), o.ServiceInvitations.IsSet()
}

// HasServiceInvitations returns a boolean if a field has been set.
func (o *RelationshipServiceInvitations) HasServiceInvitations() bool {
	if o != nil && o.ServiceInvitations.IsSet() {
		return true
	}

	return false
}

// SetServiceInvitations gets a reference to the given NullableRelationshipServiceInvitationsServiceInvitations and assigns it to the ServiceInvitations field.
func (o *RelationshipServiceInvitations) SetServiceInvitations(v RelationshipServiceInvitationsServiceInvitations) {
	o.ServiceInvitations.Set(&v)
}

// SetServiceInvitationsNil sets the value for ServiceInvitations to be an explicit nil
func (o *RelationshipServiceInvitations) SetServiceInvitationsNil() {
	o.ServiceInvitations.Set(nil)
}

// UnsetServiceInvitations ensures that no value is present for ServiceInvitations, not even an explicit nil
func (o *RelationshipServiceInvitations) UnsetServiceInvitations() {
	o.ServiceInvitations.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipServiceInvitations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceInvitations.IsSet() {
		toSerialize["service_invitations"] = o.ServiceInvitations.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipServiceInvitations) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipServiceInvitations := _RelationshipServiceInvitations{}

	if err = json.Unmarshal(bytes, &varRelationshipServiceInvitations); err == nil {
		*o = RelationshipServiceInvitations(varRelationshipServiceInvitations)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service_invitations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipServiceInvitations is a helper abstraction for handling nullable relationshipserviceinvitations types.
type NullableRelationshipServiceInvitations struct {
	value *RelationshipServiceInvitations
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipServiceInvitations) Get() *RelationshipServiceInvitations {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipServiceInvitations) Set(val *RelationshipServiceInvitations) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipServiceInvitations) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipServiceInvitations) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipServiceInvitations returns a pointer to a new instance of NullableRelationshipServiceInvitations.
func NewNullableRelationshipServiceInvitations(val *RelationshipServiceInvitations) *NullableRelationshipServiceInvitations {
	return &NullableRelationshipServiceInvitations{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipServiceInvitations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipServiceInvitations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
