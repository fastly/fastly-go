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

// RelationshipUser struct for RelationshipUser
type RelationshipUser struct {
	User *RelationshipUserUser `json:"user,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipUser RelationshipUser

// NewRelationshipUser instantiates a new RelationshipUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipUser() *RelationshipUser {
	this := RelationshipUser{}
	return &this
}

// NewRelationshipUserWithDefaults instantiates a new RelationshipUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipUserWithDefaults() *RelationshipUser {
	this := RelationshipUser{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RelationshipUser) GetUser() RelationshipUserUser {
	if o == nil || o.User == nil {
		var ret RelationshipUserUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipUser) GetUserOk() (*RelationshipUserUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RelationshipUser) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given RelationshipUserUser and assigns it to the User field.
func (o *RelationshipUser) SetUser(v RelationshipUserUser) {
	o.User = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipUser) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipUser := _RelationshipUser{}

	if err = json.Unmarshal(bytes, &varRelationshipUser); err == nil {
		*o = RelationshipUser(varRelationshipUser)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipUser is a helper abstraction for handling nullable relationshipuser types. 
type NullableRelationshipUser struct {
	value *RelationshipUser
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipUser) Get() *RelationshipUser {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipUser) Set(val *RelationshipUser) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipUser) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipUser) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipUser returns a pointer to a new instance of NullableRelationshipUser.
func NewNullableRelationshipUser(val *RelationshipUser) *NullableRelationshipUser {
	return &NullableRelationshipUser{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
