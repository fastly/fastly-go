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

// RelationshipServiceInvitationsCreate struct for RelationshipServiceInvitationsCreate
type RelationshipServiceInvitationsCreate struct {
	ServiceInvitations *RelationshipServiceInvitationsCreateServiceInvitations `json:"service_invitations,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipServiceInvitationsCreate RelationshipServiceInvitationsCreate

// NewRelationshipServiceInvitationsCreate instantiates a new RelationshipServiceInvitationsCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipServiceInvitationsCreate() *RelationshipServiceInvitationsCreate {
	this := RelationshipServiceInvitationsCreate{}
	return &this
}

// NewRelationshipServiceInvitationsCreateWithDefaults instantiates a new RelationshipServiceInvitationsCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipServiceInvitationsCreateWithDefaults() *RelationshipServiceInvitationsCreate {
	this := RelationshipServiceInvitationsCreate{}
	return &this
}

// GetServiceInvitations returns the ServiceInvitations field value if set, zero value otherwise.
func (o *RelationshipServiceInvitationsCreate) GetServiceInvitations() RelationshipServiceInvitationsCreateServiceInvitations {
	if o == nil || o.ServiceInvitations == nil {
		var ret RelationshipServiceInvitationsCreateServiceInvitations
		return ret
	}
	return *o.ServiceInvitations
}

// GetServiceInvitationsOk returns a tuple with the ServiceInvitations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipServiceInvitationsCreate) GetServiceInvitationsOk() (*RelationshipServiceInvitationsCreateServiceInvitations, bool) {
	if o == nil || o.ServiceInvitations == nil {
		return nil, false
	}
	return o.ServiceInvitations, true
}

// HasServiceInvitations returns a boolean if a field has been set.
func (o *RelationshipServiceInvitationsCreate) HasServiceInvitations() bool {
	if o != nil && o.ServiceInvitations != nil {
		return true
	}

	return false
}

// SetServiceInvitations gets a reference to the given RelationshipServiceInvitationsCreateServiceInvitations and assigns it to the ServiceInvitations field.
func (o *RelationshipServiceInvitationsCreate) SetServiceInvitations(v RelationshipServiceInvitationsCreateServiceInvitations) {
	o.ServiceInvitations = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipServiceInvitationsCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceInvitations != nil {
		toSerialize["service_invitations"] = o.ServiceInvitations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipServiceInvitationsCreate) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipServiceInvitationsCreate := _RelationshipServiceInvitationsCreate{}

	if err = json.Unmarshal(bytes, &varRelationshipServiceInvitationsCreate); err == nil {
		*o = RelationshipServiceInvitationsCreate(varRelationshipServiceInvitationsCreate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service_invitations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipServiceInvitationsCreate is a helper abstraction for handling nullable relationshipserviceinvitationscreate types. 
type NullableRelationshipServiceInvitationsCreate struct {
	value *RelationshipServiceInvitationsCreate
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipServiceInvitationsCreate) Get() *RelationshipServiceInvitationsCreate {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipServiceInvitationsCreate) Set(val *RelationshipServiceInvitationsCreate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipServiceInvitationsCreate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipServiceInvitationsCreate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipServiceInvitationsCreate returns a pointer to a new instance of NullableRelationshipServiceInvitationsCreate.
func NewNullableRelationshipServiceInvitationsCreate(val *RelationshipServiceInvitationsCreate) *NullableRelationshipServiceInvitationsCreate {
	return &NullableRelationshipServiceInvitationsCreate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipServiceInvitationsCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipServiceInvitationsCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
