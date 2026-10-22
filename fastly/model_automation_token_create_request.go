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

// AutomationTokenCreateRequest struct for AutomationTokenCreateRequest
type AutomationTokenCreateRequest struct {
	// The name of the token.
	Name string `json:"name"`
	// The role for the token. The `roles` field is the preferred way to assign roles. This field is supported for backwards compatibility with single-role tokens.
	Role *string `json:"role,omitempty"`
	// Array of role IDs to assign to the token. Use this field for tokens with multiple roles. When provided, this takes precedence over the `role` field. To get available role IDs, use [GET /roles](https://www.fastly.com/documentation/reference/api/account/roles/).
	Roles []string `json:"roles,omitempty"`
	// List of service IDs the token can access. Only used when `limit_services` is `true`.
	Services []string `json:"services,omitempty"`
	// When `true`, restricts the token to only the services specified in the `services` array. When `false` (default), the token has access to all services.
	LimitServices *bool `json:"limit_services,omitempty"`
	// When `true` (default), restricts the token to only the workspaces specified in the `workspaces` array. When `false`, the token has access to all workspaces. Only applies to roles that support workspace access (such as Next-Gen WAF roles). For roles without workspace support, this field has no effect.
	LimitWorkspaces *bool `json:"limit_workspaces,omitempty"`
	// List of workspace IDs the token can access. Only used when `limit_workspaces` is `true` and the token has roles that support workspace access.
	Workspaces []string `json:"workspaces,omitempty"`
	Scope      *string  `json:"scope,omitempty"`
	// A UTC timestamp of when the token will expire.
	ExpiresAt NullableTime `json:"expires_at"`
	// Indicates whether TLS access is enabled for the token. Only applicable when using the single-role `role` field. For tokens with multiple roles, assign TLS roles (`tls_admin` or `tls_viewer`) via the `roles` field instead.
	TlsAccess            *bool `json:"tls_access,omitempty"`
	AdditionalProperties map[string]any
}

type _AutomationTokenCreateRequest AutomationTokenCreateRequest

// NewAutomationTokenCreateRequest instantiates a new AutomationTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenCreateRequest(name string, expiresAt NullableTime) *AutomationTokenCreateRequest {
	this := AutomationTokenCreateRequest{}
	this.Name = name
	var limitServices bool = false
	this.LimitServices = &limitServices
	var limitWorkspaces bool = true
	this.LimitWorkspaces = &limitWorkspaces
	var scope string = "global"
	this.Scope = &scope
	this.ExpiresAt = expiresAt
	return &this
}

// NewAutomationTokenCreateRequestWithDefaults instantiates a new AutomationTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenCreateRequestWithDefaults() *AutomationTokenCreateRequest {
	this := AutomationTokenCreateRequest{}
	var limitServices bool = false
	this.LimitServices = &limitServices
	var limitWorkspaces bool = true
	this.LimitWorkspaces = &limitWorkspaces
	var scope string = "global"
	this.Scope = &scope
	return &this
}

// GetName returns the Name field value
func (o *AutomationTokenCreateRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutomationTokenCreateRequest) SetName(v string) {
	o.Name = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequest) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequest) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequest) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AutomationTokenCreateRequest) SetRole(v string) {
	o.Role = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequest) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequest) GetRolesOk() ([]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequest) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *AutomationTokenCreateRequest) SetRoles(v []string) {
	o.Roles = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequest) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequest) GetServicesOk() ([]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequest) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *AutomationTokenCreateRequest) SetServices(v []string) {
	o.Services = v
}

// GetLimitServices returns the LimitServices field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequest) GetLimitServices() bool {
	if o == nil || o.LimitServices == nil {
		var ret bool
		return ret
	}
	return *o.LimitServices
}

// GetLimitServicesOk returns a tuple with the LimitServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequest) GetLimitServicesOk() (*bool, bool) {
	if o == nil || o.LimitServices == nil {
		return nil, false
	}
	return o.LimitServices, true
}

// HasLimitServices returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequest) HasLimitServices() bool {
	if o != nil && o.LimitServices != nil {
		return true
	}

	return false
}

// SetLimitServices gets a reference to the given bool and assigns it to the LimitServices field.
func (o *AutomationTokenCreateRequest) SetLimitServices(v bool) {
	o.LimitServices = &v
}

// GetLimitWorkspaces returns the LimitWorkspaces field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequest) GetLimitWorkspaces() bool {
	if o == nil || o.LimitWorkspaces == nil {
		var ret bool
		return ret
	}
	return *o.LimitWorkspaces
}

// GetLimitWorkspacesOk returns a tuple with the LimitWorkspaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequest) GetLimitWorkspacesOk() (*bool, bool) {
	if o == nil || o.LimitWorkspaces == nil {
		return nil, false
	}
	return o.LimitWorkspaces, true
}

// HasLimitWorkspaces returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequest) HasLimitWorkspaces() bool {
	if o != nil && o.LimitWorkspaces != nil {
		return true
	}

	return false
}

// SetLimitWorkspaces gets a reference to the given bool and assigns it to the LimitWorkspaces field.
func (o *AutomationTokenCreateRequest) SetLimitWorkspaces(v bool) {
	o.LimitWorkspaces = &v
}

// GetWorkspaces returns the Workspaces field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequest) GetWorkspaces() []string {
	if o == nil || o.Workspaces == nil {
		var ret []string
		return ret
	}
	return o.Workspaces
}

// GetWorkspacesOk returns a tuple with the Workspaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequest) GetWorkspacesOk() ([]string, bool) {
	if o == nil || o.Workspaces == nil {
		return nil, false
	}
	return o.Workspaces, true
}

// HasWorkspaces returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequest) HasWorkspaces() bool {
	if o != nil && o.Workspaces != nil {
		return true
	}

	return false
}

// SetWorkspaces gets a reference to the given []string and assigns it to the Workspaces field.
func (o *AutomationTokenCreateRequest) SetWorkspaces(v []string) {
	o.Workspaces = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequest) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequest) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequest) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AutomationTokenCreateRequest) SetScope(v string) {
	o.Scope = &v
}

// GetExpiresAt returns the ExpiresAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *AutomationTokenCreateRequest) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutomationTokenCreateRequest) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// SetExpiresAt sets field value
func (o *AutomationTokenCreateRequest) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// GetTlsAccess returns the TlsAccess field value if set, zero value otherwise.
func (o *AutomationTokenCreateRequest) GetTlsAccess() bool {
	if o == nil || o.TlsAccess == nil {
		var ret bool
		return ret
	}
	return *o.TlsAccess
}

// GetTlsAccessOk returns a tuple with the TlsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateRequest) GetTlsAccessOk() (*bool, bool) {
	if o == nil || o.TlsAccess == nil {
		return nil, false
	}
	return o.TlsAccess, true
}

// HasTlsAccess returns a boolean if a field has been set.
func (o *AutomationTokenCreateRequest) HasTlsAccess() bool {
	if o != nil && o.TlsAccess != nil {
		return true
	}

	return false
}

// SetTlsAccess gets a reference to the given bool and assigns it to the TlsAccess field.
func (o *AutomationTokenCreateRequest) SetTlsAccess(v bool) {
	o.TlsAccess = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.LimitServices != nil {
		toSerialize["limit_services"] = o.LimitServices
	}
	if o.LimitWorkspaces != nil {
		toSerialize["limit_workspaces"] = o.LimitWorkspaces
	}
	if o.Workspaces != nil {
		toSerialize["workspaces"] = o.Workspaces
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if o.TlsAccess != nil {
		toSerialize["tls_access"] = o.TlsAccess
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AutomationTokenCreateRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenCreateRequest := _AutomationTokenCreateRequest{}

	if err = json.Unmarshal(bytes, &varAutomationTokenCreateRequest); err == nil {
		*o = AutomationTokenCreateRequest(varAutomationTokenCreateRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "role")
		delete(additionalProperties, "roles")
		delete(additionalProperties, "services")
		delete(additionalProperties, "limit_services")
		delete(additionalProperties, "limit_workspaces")
		delete(additionalProperties, "workspaces")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "tls_access")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenCreateRequest is a helper abstraction for handling nullable automationtokencreaterequest types.
type NullableAutomationTokenCreateRequest struct {
	value *AutomationTokenCreateRequest
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenCreateRequest) Get() *AutomationTokenCreateRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenCreateRequest) Set(val *AutomationTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenCreateRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenCreateRequest returns a pointer to a new instance of NullableAutomationTokenCreateRequest.
func NewNullableAutomationTokenCreateRequest(val *AutomationTokenCreateRequest) *NullableAutomationTokenCreateRequest {
	return &NullableAutomationTokenCreateRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAutomationTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
