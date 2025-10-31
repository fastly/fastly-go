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
	"time"
)

// UserResponse struct for UserResponse
type UserResponse struct {
	Login *string `json:"login,omitempty"`
	// The real life name of the user.
	Name *string `json:"name,omitempty"`
	// Indicates that the user has limited access to the customer's services.
	LimitServices *bool `json:"limit_services,omitempty"`
	// Indicates whether the is account is locked for editing or not.
	Locked NullableBool `json:"locked,omitempty"`
	// Indicates if a new password is required at next login.
	RequireNewPassword NullableBool `json:"require_new_password,omitempty"`
	Role               *RoleUser    `json:"role,omitempty"`
	// A list of role IDs assigned to the user.
	Roles []string `json:"roles,omitempty"`
	// Indicates if 2FA is enabled on the user.
	TwoFactorAuthEnabled NullableBool `json:"two_factor_auth_enabled,omitempty"`
	// Indicates if 2FA is required by the user's customer account.
	TwoFactorSetupRequired *bool `json:"two_factor_setup_required,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	Id        *string      `json:"id,omitempty"`
	// The alphanumeric string identifying a email login.
	EmailHash            *string `json:"email_hash,omitempty"`
	CustomerId           *string `json:"customer_id,omitempty"`
	AdditionalProperties map[string]any
}

type _UserResponse UserResponse

// NewUserResponse instantiates a new UserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserResponse() *UserResponse {
	this := UserResponse{}
	return &this
}

// NewUserResponseWithDefaults instantiates a new UserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserResponseWithDefaults() *UserResponse {
	this := UserResponse{}
	return &this
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *UserResponse) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *UserResponse) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *UserResponse) SetLogin(v string) {
	o.Login = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserResponse) SetName(v string) {
	o.Name = &v
}

// GetLimitServices returns the LimitServices field value if set, zero value otherwise.
func (o *UserResponse) GetLimitServices() bool {
	if o == nil || o.LimitServices == nil {
		var ret bool
		return ret
	}
	return *o.LimitServices
}

// GetLimitServicesOk returns a tuple with the LimitServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetLimitServicesOk() (*bool, bool) {
	if o == nil || o.LimitServices == nil {
		return nil, false
	}
	return o.LimitServices, true
}

// HasLimitServices returns a boolean if a field has been set.
func (o *UserResponse) HasLimitServices() bool {
	if o != nil && o.LimitServices != nil {
		return true
	}

	return false
}

// SetLimitServices gets a reference to the given bool and assigns it to the LimitServices field.
func (o *UserResponse) SetLimitServices(v bool) {
	o.LimitServices = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetLocked() bool {
	if o == nil || o.Locked.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Locked.Get()
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetLockedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Locked.Get(), o.Locked.IsSet()
}

// HasLocked returns a boolean if a field has been set.
func (o *UserResponse) HasLocked() bool {
	if o != nil && o.Locked.IsSet() {
		return true
	}

	return false
}

// SetLocked gets a reference to the given NullableBool and assigns it to the Locked field.
func (o *UserResponse) SetLocked(v bool) {
	o.Locked.Set(&v)
}

// SetLockedNil sets the value for Locked to be an explicit nil
func (o *UserResponse) SetLockedNil() {
	o.Locked.Set(nil)
}

// UnsetLocked ensures that no value is present for Locked, not even an explicit nil
func (o *UserResponse) UnsetLocked() {
	o.Locked.Unset()
}

// GetRequireNewPassword returns the RequireNewPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetRequireNewPassword() bool {
	if o == nil || o.RequireNewPassword.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RequireNewPassword.Get()
}

// GetRequireNewPasswordOk returns a tuple with the RequireNewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetRequireNewPasswordOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequireNewPassword.Get(), o.RequireNewPassword.IsSet()
}

// HasRequireNewPassword returns a boolean if a field has been set.
func (o *UserResponse) HasRequireNewPassword() bool {
	if o != nil && o.RequireNewPassword.IsSet() {
		return true
	}

	return false
}

// SetRequireNewPassword gets a reference to the given NullableBool and assigns it to the RequireNewPassword field.
func (o *UserResponse) SetRequireNewPassword(v bool) {
	o.RequireNewPassword.Set(&v)
}

// SetRequireNewPasswordNil sets the value for RequireNewPassword to be an explicit nil
func (o *UserResponse) SetRequireNewPasswordNil() {
	o.RequireNewPassword.Set(nil)
}

// UnsetRequireNewPassword ensures that no value is present for RequireNewPassword, not even an explicit nil
func (o *UserResponse) UnsetRequireNewPassword() {
	o.RequireNewPassword.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *UserResponse) GetRole() RoleUser {
	if o == nil || o.Role == nil {
		var ret RoleUser
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetRoleOk() (*RoleUser, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *UserResponse) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given RoleUser and assigns it to the Role field.
func (o *UserResponse) SetRole(v RoleUser) {
	o.Role = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserResponse) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetRolesOk() ([]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserResponse) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UserResponse) SetRoles(v []string) {
	o.Roles = v
}

// GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetTwoFactorAuthEnabled() bool {
	if o == nil || o.TwoFactorAuthEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.TwoFactorAuthEnabled.Get()
}

// GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetTwoFactorAuthEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TwoFactorAuthEnabled.Get(), o.TwoFactorAuthEnabled.IsSet()
}

// HasTwoFactorAuthEnabled returns a boolean if a field has been set.
func (o *UserResponse) HasTwoFactorAuthEnabled() bool {
	if o != nil && o.TwoFactorAuthEnabled.IsSet() {
		return true
	}

	return false
}

// SetTwoFactorAuthEnabled gets a reference to the given NullableBool and assigns it to the TwoFactorAuthEnabled field.
func (o *UserResponse) SetTwoFactorAuthEnabled(v bool) {
	o.TwoFactorAuthEnabled.Set(&v)
}

// SetTwoFactorAuthEnabledNil sets the value for TwoFactorAuthEnabled to be an explicit nil
func (o *UserResponse) SetTwoFactorAuthEnabledNil() {
	o.TwoFactorAuthEnabled.Set(nil)
}

// UnsetTwoFactorAuthEnabled ensures that no value is present for TwoFactorAuthEnabled, not even an explicit nil
func (o *UserResponse) UnsetTwoFactorAuthEnabled() {
	o.TwoFactorAuthEnabled.Unset()
}

// GetTwoFactorSetupRequired returns the TwoFactorSetupRequired field value if set, zero value otherwise.
func (o *UserResponse) GetTwoFactorSetupRequired() bool {
	if o == nil || o.TwoFactorSetupRequired == nil {
		var ret bool
		return ret
	}
	return *o.TwoFactorSetupRequired
}

// GetTwoFactorSetupRequiredOk returns a tuple with the TwoFactorSetupRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetTwoFactorSetupRequiredOk() (*bool, bool) {
	if o == nil || o.TwoFactorSetupRequired == nil {
		return nil, false
	}
	return o.TwoFactorSetupRequired, true
}

// HasTwoFactorSetupRequired returns a boolean if a field has been set.
func (o *UserResponse) HasTwoFactorSetupRequired() bool {
	if o != nil && o.TwoFactorSetupRequired != nil {
		return true
	}

	return false
}

// SetTwoFactorSetupRequired gets a reference to the given bool and assigns it to the TwoFactorSetupRequired field.
func (o *UserResponse) SetTwoFactorSetupRequired(v bool) {
	o.TwoFactorSetupRequired = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *UserResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *UserResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *UserResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *UserResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *UserResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *UserResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *UserResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UserResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UserResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *UserResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *UserResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *UserResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *UserResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserResponse) SetId(v string) {
	o.Id = &v
}

// GetEmailHash returns the EmailHash field value if set, zero value otherwise.
func (o *UserResponse) GetEmailHash() string {
	if o == nil || o.EmailHash == nil {
		var ret string
		return ret
	}
	return *o.EmailHash
}

// GetEmailHashOk returns a tuple with the EmailHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetEmailHashOk() (*string, bool) {
	if o == nil || o.EmailHash == nil {
		return nil, false
	}
	return o.EmailHash, true
}

// HasEmailHash returns a boolean if a field has been set.
func (o *UserResponse) HasEmailHash() bool {
	if o != nil && o.EmailHash != nil {
		return true
	}

	return false
}

// SetEmailHash gets a reference to the given string and assigns it to the EmailHash field.
func (o *UserResponse) SetEmailHash(v string) {
	o.EmailHash = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *UserResponse) GetCustomerId() string {
	if o == nil || o.CustomerId == nil {
		var ret string
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserResponse) GetCustomerIdOk() (*string, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *UserResponse) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given string and assigns it to the CustomerId field.
func (o *UserResponse) SetCustomerId(v string) {
	o.CustomerId = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o UserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.LimitServices != nil {
		toSerialize["limit_services"] = o.LimitServices
	}
	if o.Locked.IsSet() {
		toSerialize["locked"] = o.Locked.Get()
	}
	if o.RequireNewPassword.IsSet() {
		toSerialize["require_new_password"] = o.RequireNewPassword.Get()
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.TwoFactorAuthEnabled.IsSet() {
		toSerialize["two_factor_auth_enabled"] = o.TwoFactorAuthEnabled.Get()
	}
	if o.TwoFactorSetupRequired != nil {
		toSerialize["two_factor_setup_required"] = o.TwoFactorSetupRequired
	}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.EmailHash != nil {
		toSerialize["email_hash"] = o.EmailHash
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *UserResponse) UnmarshalJSON(bytes []byte) (err error) {
	varUserResponse := _UserResponse{}

	if err = json.Unmarshal(bytes, &varUserResponse); err == nil {
		*o = UserResponse(varUserResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "login")
		delete(additionalProperties, "name")
		delete(additionalProperties, "limit_services")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "require_new_password")
		delete(additionalProperties, "role")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "two_factor_auth_enabled")
		delete(additionalProperties, "two_factor_setup_required")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "email_hash")
		delete(additionalProperties, "customer_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableUserResponse is a helper abstraction for handling nullable userresponse types.
type NullableUserResponse struct {
	value *UserResponse
	isSet bool
}

// Get returns the value.
func (v NullableUserResponse) Get() *UserResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableUserResponse) Set(val *UserResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableUserResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUserResponse returns a pointer to a new instance of NullableUserResponse.
func NewNullableUserResponse(val *UserResponse) *NullableUserResponse {
	return &NullableUserResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
