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

// AutomationTokenCreateResponse struct for AutomationTokenCreateResponse
type AutomationTokenCreateResponse struct {
	// The name of the token.
	Name *string `json:"name,omitempty"`
	// The role on the token.
	Role *string `json:"role,omitempty"`
	// (Optional) The service IDs of the services the token will have access to. Separate service IDs with a space. If no services are specified, the token will have access to all services on the account.
	Services []string `json:"services,omitempty"`
	// A space-delimited list of authorization scope.
	Scope *string `json:"scope,omitempty"`
	// A UTC time-stamp of when the token expires.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// A UTC time-stamp of when the token was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt     NullableTime        `json:"updated_at,omitempty"`
	ID            *ReadOnlyID         `json:"id,omitempty"`
	UserID        *ReadOnlyUserID     `json:"user_id,omitempty"`
	CustomerID    *ReadOnlyCustomerID `json:"customer_id,omitempty"`
	SudoExpiresAt *time.Time          `json:"sudo_expires_at,omitempty"`
	AccessToken   *string             `json:"access_token,omitempty"`
	// A UTC time-stamp of when the token was last used.
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	// The User-Agent header of the client that last used the token.
	UserAgent            *string `json:"user_agent,omitempty"`
	AdditionalProperties map[string]any
}

type _AutomationTokenCreateResponse AutomationTokenCreateResponse

// NewAutomationTokenCreateResponse instantiates a new AutomationTokenCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenCreateResponse() *AutomationTokenCreateResponse {
	this := AutomationTokenCreateResponse{}
	var scope string = "global"
	this.Scope = &scope
	return &this
}

// NewAutomationTokenCreateResponseWithDefaults instantiates a new AutomationTokenCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenCreateResponseWithDefaults() *AutomationTokenCreateResponse {
	this := AutomationTokenCreateResponse{}
	var scope string = "global"
	this.Scope = &scope
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AutomationTokenCreateResponse) SetName(v string) {
	o.Name = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AutomationTokenCreateResponse) SetRole(v string) {
	o.Role = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetServicesOk() ([]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *AutomationTokenCreateResponse) SetServices(v []string) {
	o.Services = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AutomationTokenCreateResponse) SetScope(v string) {
	o.Scope = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *AutomationTokenCreateResponse) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AutomationTokenCreateResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutomationTokenCreateResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutomationTokenCreateResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *AutomationTokenCreateResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *AutomationTokenCreateResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *AutomationTokenCreateResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutomationTokenCreateResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutomationTokenCreateResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *AutomationTokenCreateResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *AutomationTokenCreateResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *AutomationTokenCreateResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetID() ReadOnlyID {
	if o == nil || o.ID == nil {
		var ret ReadOnlyID
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetIDOk() (*ReadOnlyID, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given ReadOnlyID and assigns it to the ID field.
func (o *AutomationTokenCreateResponse) SetID(v ReadOnlyID) {
	o.ID = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetUserID() ReadOnlyUserID {
	if o == nil || o.UserID == nil {
		var ret ReadOnlyUserID
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetUserIDOk() (*ReadOnlyUserID, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given ReadOnlyUserID and assigns it to the UserID field.
func (o *AutomationTokenCreateResponse) SetUserID(v ReadOnlyUserID) {
	o.UserID = &v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetCustomerID() ReadOnlyCustomerID {
	if o == nil || o.CustomerID == nil {
		var ret ReadOnlyCustomerID
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetCustomerIDOk() (*ReadOnlyCustomerID, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given ReadOnlyCustomerID and assigns it to the CustomerID field.
func (o *AutomationTokenCreateResponse) SetCustomerID(v ReadOnlyCustomerID) {
	o.CustomerID = &v
}

// GetSudoExpiresAt returns the SudoExpiresAt field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetSudoExpiresAt() time.Time {
	if o == nil || o.SudoExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.SudoExpiresAt
}

// GetSudoExpiresAtOk returns a tuple with the SudoExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetSudoExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.SudoExpiresAt == nil {
		return nil, false
	}
	return o.SudoExpiresAt, true
}

// HasSudoExpiresAt returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasSudoExpiresAt() bool {
	if o != nil && o.SudoExpiresAt != nil {
		return true
	}

	return false
}

// SetSudoExpiresAt gets a reference to the given time.Time and assigns it to the SudoExpiresAt field.
func (o *AutomationTokenCreateResponse) SetSudoExpiresAt(v time.Time) {
	o.SudoExpiresAt = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *AutomationTokenCreateResponse) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || o.LastUsedAt == nil {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt != nil {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *AutomationTokenCreateResponse) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetUserAgent() string {
	if o == nil || o.UserAgent == nil {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetUserAgentOk() (*string, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *AutomationTokenCreateResponse) SetUserAgent(v string) {
	o.UserAgent = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
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
	if o.UserID != nil {
		toSerialize["user_id"] = o.UserID
	}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.SudoExpiresAt != nil {
		toSerialize["sudo_expires_at"] = o.SudoExpiresAt
	}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.LastUsedAt != nil {
		toSerialize["last_used_at"] = o.LastUsedAt
	}
	if o.UserAgent != nil {
		toSerialize["user_agent"] = o.UserAgent
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AutomationTokenCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenCreateResponse := _AutomationTokenCreateResponse{}

	if err = json.Unmarshal(bytes, &varAutomationTokenCreateResponse); err == nil {
		*o = AutomationTokenCreateResponse(varAutomationTokenCreateResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "role")
		delete(additionalProperties, "services")
		delete(additionalProperties, "scope")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "id")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "sudo_expires_at")
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "last_used_at")
		delete(additionalProperties, "user_agent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenCreateResponse is a helper abstraction for handling nullable automationtokencreateresponse types.
type NullableAutomationTokenCreateResponse struct {
	value *AutomationTokenCreateResponse
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenCreateResponse) Get() *AutomationTokenCreateResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenCreateResponse) Set(val *AutomationTokenCreateResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenCreateResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenCreateResponse returns a pointer to a new instance of NullableAutomationTokenCreateResponse.
func NewNullableAutomationTokenCreateResponse(val *AutomationTokenCreateResponse) *NullableAutomationTokenCreateResponse {
	return &NullableAutomationTokenCreateResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAutomationTokenCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
