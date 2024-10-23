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

// RelationshipsForStar struct for RelationshipsForStar
type RelationshipsForStar struct {
	User                 *RelationshipUserUser      `json:"user,omitempty"`
	Service              *RelationshipMemberService `json:"service,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipsForStar RelationshipsForStar

// NewRelationshipsForStar instantiates a new RelationshipsForStar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipsForStar() *RelationshipsForStar {
	this := RelationshipsForStar{}
	return &this
}

// NewRelationshipsForStarWithDefaults instantiates a new RelationshipsForStar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipsForStarWithDefaults() *RelationshipsForStar {
	this := RelationshipsForStar{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RelationshipsForStar) GetUser() RelationshipUserUser {
	if o == nil || o.User == nil {
		var ret RelationshipUserUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipsForStar) GetUserOk() (*RelationshipUserUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RelationshipsForStar) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given RelationshipUserUser and assigns it to the User field.
func (o *RelationshipsForStar) SetUser(v RelationshipUserUser) {
	o.User = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *RelationshipsForStar) GetService() RelationshipMemberService {
	if o == nil || o.Service == nil {
		var ret RelationshipMemberService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipsForStar) GetServiceOk() (*RelationshipMemberService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *RelationshipsForStar) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given RelationshipMemberService and assigns it to the Service field.
func (o *RelationshipsForStar) SetService(v RelationshipMemberService) {
	o.Service = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipsForStar) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipsForStar) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipsForStar := _RelationshipsForStar{}

	if err = json.Unmarshal(bytes, &varRelationshipsForStar); err == nil {
		*o = RelationshipsForStar(varRelationshipsForStar)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		delete(additionalProperties, "service")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipsForStar is a helper abstraction for handling nullable relationshipsforstar types.
type NullableRelationshipsForStar struct {
	value *RelationshipsForStar
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipsForStar) Get() *RelationshipsForStar {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipsForStar) Set(val *RelationshipsForStar) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipsForStar) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipsForStar) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipsForStar returns a pointer to a new instance of NullableRelationshipsForStar.
func NewNullableRelationshipsForStar(val *RelationshipsForStar) *NullableRelationshipsForStar {
	return &NullableRelationshipsForStar{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipsForStar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipsForStar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
