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

// SchemasUserResponse struct for SchemasUserResponse
type SchemasUserResponse struct {
	Login *string `json:"login,omitempty"`
	// The real life name of the user.
	Name *string `json:"name,omitempty"`
	// Indicates that the user has limited access to the customer's services.
	LimitServices *bool `json:"limit_services,omitempty"`
	// Indicates whether the is account is locked for editing or not.
	Locked *bool `json:"locked,omitempty"`
	// Indicates if a new password is required at next login.
	RequireNewPassword *bool `json:"require_new_password,omitempty"`
	Role *RoleUser `json:"role,omitempty"`
	// Indicates if 2FA is enabled on the user.
	TwoFactorAuthEnabled *bool `json:"two_factor_auth_enabled,omitempty"`
	// Indicates if 2FA is required by the user's customer account.
	TwoFactorSetupRequired *bool `json:"two_factor_setup_required,omitempty"`
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	ID *string `json:"id,omitempty"`
	// The alphanumeric string identifying a email login.
	EmailHash *string `json:"email_hash,omitempty"`
	CustomerID *string `json:"customer_id,omitempty"`
	AdditionalProperties map[string]any
}

type _SchemasUserResponse SchemasUserResponse

// NewSchemasUserResponse instantiates a new SchemasUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUserResponse() *SchemasUserResponse {
	this := SchemasUserResponse{}
	return &this
}

// NewSchemasUserResponseWithDefaults instantiates a new SchemasUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUserResponseWithDefaults() *SchemasUserResponse {
	this := SchemasUserResponse{}
	return &this
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *SchemasUserResponse) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *SchemasUserResponse) SetLogin(v string) {
	o.Login = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SchemasUserResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SchemasUserResponse) SetName(v string) {
	o.Name = &v
}

// GetLimitServices returns the LimitServices field value if set, zero value otherwise.
func (o *SchemasUserResponse) GetLimitServices() bool {
	if o == nil || o.LimitServices == nil {
		var ret bool
		return ret
	}
	return *o.LimitServices
}

// GetLimitServicesOk returns a tuple with the LimitServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetLimitServicesOk() (*bool, bool) {
	if o == nil || o.LimitServices == nil {
		return nil, false
	}
	return o.LimitServices, true
}

// HasLimitServices returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasLimitServices() bool {
	if o != nil && o.LimitServices != nil {
		return true
	}

	return false
}

// SetLimitServices gets a reference to the given bool and assigns it to the LimitServices field.
func (o *SchemasUserResponse) SetLimitServices(v bool) {
	o.LimitServices = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *SchemasUserResponse) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *SchemasUserResponse) SetLocked(v bool) {
	o.Locked = &v
}

// GetRequireNewPassword returns the RequireNewPassword field value if set, zero value otherwise.
func (o *SchemasUserResponse) GetRequireNewPassword() bool {
	if o == nil || o.RequireNewPassword == nil {
		var ret bool
		return ret
	}
	return *o.RequireNewPassword
}

// GetRequireNewPasswordOk returns a tuple with the RequireNewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetRequireNewPasswordOk() (*bool, bool) {
	if o == nil || o.RequireNewPassword == nil {
		return nil, false
	}
	return o.RequireNewPassword, true
}

// HasRequireNewPassword returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasRequireNewPassword() bool {
	if o != nil && o.RequireNewPassword != nil {
		return true
	}

	return false
}

// SetRequireNewPassword gets a reference to the given bool and assigns it to the RequireNewPassword field.
func (o *SchemasUserResponse) SetRequireNewPassword(v bool) {
	o.RequireNewPassword = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *SchemasUserResponse) GetRole() RoleUser {
	if o == nil || o.Role == nil {
		var ret RoleUser
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetRoleOk() (*RoleUser, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given RoleUser and assigns it to the Role field.
func (o *SchemasUserResponse) SetRole(v RoleUser) {
	o.Role = &v
}

// GetTwoFactorAuthEnabled returns the TwoFactorAuthEnabled field value if set, zero value otherwise.
func (o *SchemasUserResponse) GetTwoFactorAuthEnabled() bool {
	if o == nil || o.TwoFactorAuthEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TwoFactorAuthEnabled
}

// GetTwoFactorAuthEnabledOk returns a tuple with the TwoFactorAuthEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetTwoFactorAuthEnabledOk() (*bool, bool) {
	if o == nil || o.TwoFactorAuthEnabled == nil {
		return nil, false
	}
	return o.TwoFactorAuthEnabled, true
}

// HasTwoFactorAuthEnabled returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasTwoFactorAuthEnabled() bool {
	if o != nil && o.TwoFactorAuthEnabled != nil {
		return true
	}

	return false
}

// SetTwoFactorAuthEnabled gets a reference to the given bool and assigns it to the TwoFactorAuthEnabled field.
func (o *SchemasUserResponse) SetTwoFactorAuthEnabled(v bool) {
	o.TwoFactorAuthEnabled = &v
}

// GetTwoFactorSetupRequired returns the TwoFactorSetupRequired field value if set, zero value otherwise.
func (o *SchemasUserResponse) GetTwoFactorSetupRequired() bool {
	if o == nil || o.TwoFactorSetupRequired == nil {
		var ret bool
		return ret
	}
	return *o.TwoFactorSetupRequired
}

// GetTwoFactorSetupRequiredOk returns a tuple with the TwoFactorSetupRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetTwoFactorSetupRequiredOk() (*bool, bool) {
	if o == nil || o.TwoFactorSetupRequired == nil {
		return nil, false
	}
	return o.TwoFactorSetupRequired, true
}

// HasTwoFactorSetupRequired returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasTwoFactorSetupRequired() bool {
	if o != nil && o.TwoFactorSetupRequired != nil {
		return true
	}

	return false
}

// SetTwoFactorSetupRequired gets a reference to the given bool and assigns it to the TwoFactorSetupRequired field.
func (o *SchemasUserResponse) SetTwoFactorSetupRequired(v bool) {
	o.TwoFactorSetupRequired = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasUserResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasUserResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *SchemasUserResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}
// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *SchemasUserResponse) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *SchemasUserResponse) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasUserResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasUserResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *SchemasUserResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}
// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *SchemasUserResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *SchemasUserResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SchemasUserResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SchemasUserResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *SchemasUserResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}
// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *SchemasUserResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *SchemasUserResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *SchemasUserResponse) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *SchemasUserResponse) SetID(v string) {
	o.ID = &v
}

// GetEmailHash returns the EmailHash field value if set, zero value otherwise.
func (o *SchemasUserResponse) GetEmailHash() string {
	if o == nil || o.EmailHash == nil {
		var ret string
		return ret
	}
	return *o.EmailHash
}

// GetEmailHashOk returns a tuple with the EmailHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetEmailHashOk() (*string, bool) {
	if o == nil || o.EmailHash == nil {
		return nil, false
	}
	return o.EmailHash, true
}

// HasEmailHash returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasEmailHash() bool {
	if o != nil && o.EmailHash != nil {
		return true
	}

	return false
}

// SetEmailHash gets a reference to the given string and assigns it to the EmailHash field.
func (o *SchemasUserResponse) SetEmailHash(v string) {
	o.EmailHash = &v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *SchemasUserResponse) GetCustomerID() string {
	if o == nil || o.CustomerID == nil {
		var ret string
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetCustomerIDOk() (*string, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *SchemasUserResponse) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given string and assigns it to the CustomerID field.
func (o *SchemasUserResponse) SetCustomerID(v string) {
	o.CustomerID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SchemasUserResponse) MarshalJSON() ([]byte, error) {
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
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}
	if o.RequireNewPassword != nil {
		toSerialize["require_new_password"] = o.RequireNewPassword
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.TwoFactorAuthEnabled != nil {
		toSerialize["two_factor_auth_enabled"] = o.TwoFactorAuthEnabled
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
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.EmailHash != nil {
		toSerialize["email_hash"] = o.EmailHash
	}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *SchemasUserResponse) UnmarshalJSON(bytes []byte) (err error) {
	varSchemasUserResponse := _SchemasUserResponse{}

	if err = json.Unmarshal(bytes, &varSchemasUserResponse); err == nil {
		*o = SchemasUserResponse(varSchemasUserResponse)
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

// NullableSchemasUserResponse is a helper abstraction for handling nullable schemasuserresponse types. 
type NullableSchemasUserResponse struct {
	value *SchemasUserResponse
	isSet bool
}

// Get returns the value.
func (v NullableSchemasUserResponse) Get() *SchemasUserResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableSchemasUserResponse) Set(val *SchemasUserResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSchemasUserResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSchemasUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSchemasUserResponse returns a pointer to a new instance of NullableSchemasUserResponse.
func NewNullableSchemasUserResponse(val *SchemasUserResponse) *NullableSchemasUserResponse {
	return &NullableSchemasUserResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSchemasUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableSchemasUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
