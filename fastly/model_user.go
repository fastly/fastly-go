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

// User struct for User
type User struct {
	Login *string `json:"login,omitempty"`
	// The real life name of the user.
	Name *string `json:"name,omitempty"`
	// Indicates that the user has limited access to the customer's services.
	LimitServices *bool `json:"limit_services,omitempty"`
	// Indicates whether the is account is locked for editing or not.
	Locked NullableBool `json:"locked,omitempty"`
	// Indicates if a new password is required at next login.
	RequireNewPassword NullableBool `json:"require_new_password,omitempty"`
	Role *RoleUser `json:"role,omitempty"`
	// Indicates if 2FA is enabled on the user.
	TwoFactorAuthEnabled NullableBool `json:"two_factor_auth_enabled,omitempty"`
	// Indicates if 2FA is required by the user's customer account.
	TwoFactorSetupRequired *bool `json:"two_factor_setup_required,omitempty"`
	AdditionalProperties map[string]any
}

type _User User

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *User) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *User) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *User) SetLogin(v string) {
	o.Login = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

// GetLimitServices returns the LimitServices field value if set, zero value otherwise.
func (o *User) GetLimitServices() bool {
	if o == nil || o.LimitServices == nil {
		var ret bool
		return ret
	}
	return *o.LimitServices
}

// GetLimitServicesOk returns a tuple with the LimitServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLimitServicesOk() (*bool, bool) {
	if o == nil || o.LimitServices == nil {
		return nil, false
	}
	return o.LimitServices, true
}

// HasLimitServices returns a boolean if a field has been set.
func (o *User) HasLimitServices() bool {
	if o != nil && o.LimitServices != nil {
		return true
	}

	return false
}

// SetLimitServices gets a reference to the given bool and assigns it to the LimitServices field.
func (o *User) SetLimitServices(v bool) {
	o.LimitServices = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetLocked() bool {
	if o == nil || o.Locked.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Locked.Get()
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetLockedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Locked.Get(), o.Locked.IsSet()
}

// HasLocked returns a boolean if a field has been set.
func (o *User) HasLocked() bool {
	if o != nil && o.Locked.IsSet() {
		return true
	}

	return false
}

// SetLocked gets a reference to the given NullableBool and assigns it to the Locked field.
func (o *User) SetLocked(v bool) {
	o.Locked.Set(&v)
}
// SetLockedNil sets the value for Locked to be an explicit nil
func (o *User) SetLockedNil() {
	o.Locked.Set(nil)
}

// UnsetLocked ensures that no value is present for Locked, not even an explicit nil
func (o *User) UnsetLocked() {
	o.Locked.Unset()
}

// GetRequireNewPassword returns the RequireNewPassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetRequireNewPassword() bool {
	if o == nil || o.RequireNewPassword.Get() == nil {
		var ret bool
		return ret
	}
	return *o.RequireNewPassword.Get()
}

// GetRequireNewPasswordOk returns a tuple with the RequireNewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetRequireNewPasswordOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RequireNewPassword.Get(), o.RequireNewPassword.IsSet()
}

// HasRequireNewPassword returns a boolean if a field has been set.
func (o *User) HasRequireNewPassword() bool {
	if o != nil && o.RequireNewPassword.IsSet() {
		return true
	}

	return false
}

// SetRequireNewPassword gets a reference to the given NullableBool and assigns it to the RequireNewPassword field.
func (o *User) SetRequireNewPassword(v bool) {
	o.RequireNewPassword.Set(&v)
}
// SetRequireNewPasswordNil sets the value for RequireNewPassword to be an explicit nil
func (o *User) SetRequireNewPasswordNil() {
	o.RequireNewPassword.Set(nil)
}

// UnsetRequireNewPassword ensures that no value is present for RequireNewPassword, not even an explicit nil
func (o *User) UnsetRequireNewPassword() {
	o.RequireNewPassword.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *User) GetRole() RoleUser {
	if o == nil || o.Role == nil {
		var ret RoleUser
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetRoleOk() (*RoleUser, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *User) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given RoleUser and assigns it to the Role field.
func (o *User) SetRole(v RoleUser) {
	o.Role = &v
}

// GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetTwoFactorAuthEnabled() bool {
	if o == nil || o.TwoFactorAuthEnabled.Get() == nil {
		var ret bool
		return ret
	}
	return *o.TwoFactorAuthEnabled.Get()
}

// GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetTwoFactorAuthEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TwoFactorAuthEnabled.Get(), o.TwoFactorAuthEnabled.IsSet()
}

// HasTwoFactorAuthEnabled returns a boolean if a field has been set.
func (o *User) HasTwoFactorAuthEnabled() bool {
	if o != nil && o.TwoFactorAuthEnabled.IsSet() {
		return true
	}

	return false
}

// SetTwoFactorAuthEnabled gets a reference to the given NullableBool and assigns it to the TwoFactorAuthEnabled field.
func (o *User) SetTwoFactorAuthEnabled(v bool) {
	o.TwoFactorAuthEnabled.Set(&v)
}
// SetTwoFactorAuthEnabledNil sets the value for TwoFactorAuthEnabled to be an explicit nil
func (o *User) SetTwoFactorAuthEnabledNil() {
	o.TwoFactorAuthEnabled.Set(nil)
}

// UnsetTwoFactorAuthEnabled ensures that no value is present for TwoFactorAuthEnabled, not even an explicit nil
func (o *User) UnsetTwoFactorAuthEnabled() {
	o.TwoFactorAuthEnabled.Unset()
}

// GetTwoFactorSetupRequired returns the TwoFactorSetupRequired field value if set, zero value otherwise.
func (o *User) GetTwoFactorSetupRequired() bool {
	if o == nil || o.TwoFactorSetupRequired == nil {
		var ret bool
		return ret
	}
	return *o.TwoFactorSetupRequired
}

// GetTwoFactorSetupRequiredOk returns a tuple with the TwoFactorSetupRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTwoFactorSetupRequiredOk() (*bool, bool) {
	if o == nil || o.TwoFactorSetupRequired == nil {
		return nil, false
	}
	return o.TwoFactorSetupRequired, true
}

// HasTwoFactorSetupRequired returns a boolean if a field has been set.
func (o *User) HasTwoFactorSetupRequired() bool {
	if o != nil && o.TwoFactorSetupRequired != nil {
		return true
	}

	return false
}

// SetTwoFactorSetupRequired gets a reference to the given bool and assigns it to the TwoFactorSetupRequired field.
func (o *User) SetTwoFactorSetupRequired(v bool) {
	o.TwoFactorSetupRequired = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o User) MarshalJSON() ([]byte, error) {
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
	if o.TwoFactorAuthEnabled.IsSet() {
		toSerialize["two_factor_auth_enabled"] = o.TwoFactorAuthEnabled.Get()
	}
	if o.TwoFactorSetupRequired != nil {
		toSerialize["two_factor_setup_required"] = o.TwoFactorSetupRequired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *User) UnmarshalJSON(bytes []byte) (err error) {
	varUser := _User{}

	if err = json.Unmarshal(bytes, &varUser); err == nil {
		*o = User(varUser)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "login")
		delete(additionalProperties, "name")
		delete(additionalProperties, "limit_services")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "require_new_password")
		delete(additionalProperties, "role")
		delete(additionalProperties, "two_factor_auth_enabled")
		delete(additionalProperties, "two_factor_setup_required")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableUser is a helper abstraction for handling nullable user types. 
type NullableUser struct {
	value *User
	isSet bool
}

// Get returns the value.
func (v NullableUser) Get() *User {
	return v.value
}

// Set modifies the value.
func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableUser) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUser returns a pointer to a new instance of NullableUser.
func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
