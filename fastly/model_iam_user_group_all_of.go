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

// IamUserGroupAllOf struct for IamUserGroupAllOf
type IamUserGroupAllOf struct {
	// Alphanumeric string identifying the user group.
	Id *string `json:"id,omitempty"`
	// Name of the user group.
	Name *string `json:"name,omitempty"`
	// Description of the user group.
	Description *string `json:"description,omitempty"`
	// Number of invitations added to the user group.
	InvitationsCount *int32 `json:"invitations_count,omitempty"`
	// Number of users added to the user group.
	UsersCount *int32 `json:"users_count,omitempty"`
	// Number of roles added to the user group.
	RolesCount           *int32 `json:"roles_count,omitempty"`
	AdditionalProperties map[string]any
}

type _IamUserGroupAllOf IamUserGroupAllOf

// NewIamUserGroupAllOf instantiates a new IamUserGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamUserGroupAllOf() *IamUserGroupAllOf {
	this := IamUserGroupAllOf{}
	return &this
}

// NewIamUserGroupAllOfWithDefaults instantiates a new IamUserGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamUserGroupAllOfWithDefaults() *IamUserGroupAllOf {
	this := IamUserGroupAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IamUserGroupAllOf) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroupAllOf) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IamUserGroupAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IamUserGroupAllOf) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamUserGroupAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroupAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamUserGroupAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamUserGroupAllOf) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamUserGroupAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroupAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamUserGroupAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamUserGroupAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetInvitationsCount returns the InvitationsCount field value if set, zero value otherwise.
func (o *IamUserGroupAllOf) GetInvitationsCount() int32 {
	if o == nil || o.InvitationsCount == nil {
		var ret int32
		return ret
	}
	return *o.InvitationsCount
}

// GetInvitationsCountOk returns a tuple with the InvitationsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroupAllOf) GetInvitationsCountOk() (*int32, bool) {
	if o == nil || o.InvitationsCount == nil {
		return nil, false
	}
	return o.InvitationsCount, true
}

// HasInvitationsCount returns a boolean if a field has been set.
func (o *IamUserGroupAllOf) HasInvitationsCount() bool {
	if o != nil && o.InvitationsCount != nil {
		return true
	}

	return false
}

// SetInvitationsCount gets a reference to the given int32 and assigns it to the InvitationsCount field.
func (o *IamUserGroupAllOf) SetInvitationsCount(v int32) {
	o.InvitationsCount = &v
}

// GetUsersCount returns the UsersCount field value if set, zero value otherwise.
func (o *IamUserGroupAllOf) GetUsersCount() int32 {
	if o == nil || o.UsersCount == nil {
		var ret int32
		return ret
	}
	return *o.UsersCount
}

// GetUsersCountOk returns a tuple with the UsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroupAllOf) GetUsersCountOk() (*int32, bool) {
	if o == nil || o.UsersCount == nil {
		return nil, false
	}
	return o.UsersCount, true
}

// HasUsersCount returns a boolean if a field has been set.
func (o *IamUserGroupAllOf) HasUsersCount() bool {
	if o != nil && o.UsersCount != nil {
		return true
	}

	return false
}

// SetUsersCount gets a reference to the given int32 and assigns it to the UsersCount field.
func (o *IamUserGroupAllOf) SetUsersCount(v int32) {
	o.UsersCount = &v
}

// GetRolesCount returns the RolesCount field value if set, zero value otherwise.
func (o *IamUserGroupAllOf) GetRolesCount() int32 {
	if o == nil || o.RolesCount == nil {
		var ret int32
		return ret
	}
	return *o.RolesCount
}

// GetRolesCountOk returns a tuple with the RolesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroupAllOf) GetRolesCountOk() (*int32, bool) {
	if o == nil || o.RolesCount == nil {
		return nil, false
	}
	return o.RolesCount, true
}

// HasRolesCount returns a boolean if a field has been set.
func (o *IamUserGroupAllOf) HasRolesCount() bool {
	if o != nil && o.RolesCount != nil {
		return true
	}

	return false
}

// SetRolesCount gets a reference to the given int32 and assigns it to the RolesCount field.
func (o *IamUserGroupAllOf) SetRolesCount(v int32) {
	o.RolesCount = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o IamUserGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.InvitationsCount != nil {
		toSerialize["invitations_count"] = o.InvitationsCount
	}
	if o.UsersCount != nil {
		toSerialize["users_count"] = o.UsersCount
	}
	if o.RolesCount != nil {
		toSerialize["roles_count"] = o.RolesCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *IamUserGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamUserGroupAllOf := _IamUserGroupAllOf{}

	if err = json.Unmarshal(bytes, &varIamUserGroupAllOf); err == nil {
		*o = IamUserGroupAllOf(varIamUserGroupAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "invitations_count")
		delete(additionalProperties, "users_count")
		delete(additionalProperties, "roles_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableIamUserGroupAllOf is a helper abstraction for handling nullable iamusergroupallof types.
type NullableIamUserGroupAllOf struct {
	value *IamUserGroupAllOf
	isSet bool
}

// Get returns the value.
func (v NullableIamUserGroupAllOf) Get() *IamUserGroupAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableIamUserGroupAllOf) Set(val *IamUserGroupAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIamUserGroupAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIamUserGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIamUserGroupAllOf returns a pointer to a new instance of NullableIamUserGroupAllOf.
func NewNullableIamUserGroupAllOf(val *IamUserGroupAllOf) *NullableIamUserGroupAllOf {
	return &NullableIamUserGroupAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIamUserGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableIamUserGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
