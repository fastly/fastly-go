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
	"time"
)

// IamUserGroup struct for IamUserGroup
type IamUserGroup struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	// Alphanumeric string identifying the user group.
	ID *string `json:"id,omitempty"`
	// Name of the user group.
	Name *string `json:"name,omitempty"`
	// Description of the user group.
	Description *string `json:"description,omitempty"`
	// Number of invitations added to the user group.
	InvitationsCount *int32 `json:"invitations_count,omitempty"`
	// Number of users added to the user group.
	UsersCount *int32 `json:"users_count,omitempty"`
	// Number of roles added to the user group.
	RolesCount *int32 `json:"roles_count,omitempty"`
	AdditionalProperties map[string]any
}

type _IamUserGroup IamUserGroup

// NewIamUserGroup instantiates a new IamUserGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamUserGroup() *IamUserGroup {
	this := IamUserGroup{}
	return &this
}

// NewIamUserGroupWithDefaults instantiates a new IamUserGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamUserGroupWithDefaults() *IamUserGroup {
	this := IamUserGroup{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserGroup) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserGroup) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IamUserGroup) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *IamUserGroup) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *IamUserGroup) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *IamUserGroup) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IamUserGroup) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IamUserGroup) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IamUserGroup) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *IamUserGroup) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *IamUserGroup) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *IamUserGroup) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *IamUserGroup) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroup) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *IamUserGroup) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *IamUserGroup) SetID(v string) {
	o.ID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamUserGroup) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroup) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamUserGroup) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamUserGroup) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamUserGroup) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamUserGroup) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamUserGroup) SetDescription(v string) {
	o.Description = &v
}

// GetInvitationsCount returns the InvitationsCount field value if set, zero value otherwise.
func (o *IamUserGroup) GetInvitationsCount() int32 {
	if o == nil || o.InvitationsCount == nil {
		var ret int32
		return ret
	}
	return *o.InvitationsCount
}

// GetInvitationsCountOk returns a tuple with the InvitationsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroup) GetInvitationsCountOk() (*int32, bool) {
	if o == nil || o.InvitationsCount == nil {
		return nil, false
	}
	return o.InvitationsCount, true
}

// HasInvitationsCount returns a boolean if a field has been set.
func (o *IamUserGroup) HasInvitationsCount() bool {
	if o != nil && o.InvitationsCount != nil {
		return true
	}

	return false
}

// SetInvitationsCount gets a reference to the given int32 and assigns it to the InvitationsCount field.
func (o *IamUserGroup) SetInvitationsCount(v int32) {
	o.InvitationsCount = &v
}

// GetUsersCount returns the UsersCount field value if set, zero value otherwise.
func (o *IamUserGroup) GetUsersCount() int32 {
	if o == nil || o.UsersCount == nil {
		var ret int32
		return ret
	}
	return *o.UsersCount
}

// GetUsersCountOk returns a tuple with the UsersCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroup) GetUsersCountOk() (*int32, bool) {
	if o == nil || o.UsersCount == nil {
		return nil, false
	}
	return o.UsersCount, true
}

// HasUsersCount returns a boolean if a field has been set.
func (o *IamUserGroup) HasUsersCount() bool {
	if o != nil && o.UsersCount != nil {
		return true
	}

	return false
}

// SetUsersCount gets a reference to the given int32 and assigns it to the UsersCount field.
func (o *IamUserGroup) SetUsersCount(v int32) {
	o.UsersCount = &v
}

// GetRolesCount returns the RolesCount field value if set, zero value otherwise.
func (o *IamUserGroup) GetRolesCount() int32 {
	if o == nil || o.RolesCount == nil {
		var ret int32
		return ret
	}
	return *o.RolesCount
}

// GetRolesCountOk returns a tuple with the RolesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamUserGroup) GetRolesCountOk() (*int32, bool) {
	if o == nil || o.RolesCount == nil {
		return nil, false
	}
	return o.RolesCount, true
}

// HasRolesCount returns a boolean if a field has been set.
func (o *IamUserGroup) HasRolesCount() bool {
	if o != nil && o.RolesCount != nil {
		return true
	}

	return false
}

// SetRolesCount gets a reference to the given int32 and assigns it to the RolesCount field.
func (o *IamUserGroup) SetRolesCount(v int32) {
	o.RolesCount = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o IamUserGroup) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
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
func (o *IamUserGroup) UnmarshalJSON(bytes []byte) (err error) {
	varIamUserGroup := _IamUserGroup{}

	if err = json.Unmarshal(bytes, &varIamUserGroup); err == nil {
		*o = IamUserGroup(varIamUserGroup)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
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

// NullableIamUserGroup is a helper abstraction for handling nullable iamusergroup types. 
type NullableIamUserGroup struct {
	value *IamUserGroup
	isSet bool
}

// Get returns the value.
func (v NullableIamUserGroup) Get() *IamUserGroup {
	return v.value
}

// Set modifies the value.
func (v *NullableIamUserGroup) Set(val *IamUserGroup) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIamUserGroup) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIamUserGroup) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIamUserGroup returns a pointer to a new instance of NullableIamUserGroup.
func NewNullableIamUserGroup(val *IamUserGroup) *NullableIamUserGroup {
	return &NullableIamUserGroup{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIamUserGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableIamUserGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
