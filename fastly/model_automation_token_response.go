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

// AutomationTokenResponse struct for AutomationTokenResponse
type AutomationTokenResponse struct {
	// The name of the token.
	Name *string `json:"name,omitempty"`
	Role *string `json:"role,omitempty"`
	// (Optional) The service IDs of the services the token will have access to. Separate service IDs with a space. If no services are specified, the token will have access to all services on the account.
	Services []string `json:"services,omitempty"`
	// A space-delimited list of authorization scope.
	Scope *string `json:"scope,omitempty"`
	// (optional) A UTC time-stamp of when the token will expire.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// A UTC time-stamp of when the token was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt  NullableTime        `json:"updated_at,omitempty"`
	ID         *ReadOnlyID         `json:"id,omitempty"`
	CustomerID *ReadOnlyCustomerID `json:"customer_id,omitempty"`
	// The IP address of the client that last used the token.
	IP *string `json:"ip,omitempty"`
	// The User-Agent header of the client that last used the token.
	UserAgent     *string `json:"user_agent,omitempty"`
	SudoExpiresAt *string `json:"sudo_expires_at,omitempty"`
	// A UTC time-stamp of when the token was last used.
	LastUsedAt           *time.Time `json:"last_used_at,omitempty"`
	AdditionalProperties map[string]any
}

type _AutomationTokenResponse AutomationTokenResponse

// NewAutomationTokenResponse instantiates a new AutomationTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenResponse() *AutomationTokenResponse {
	this := AutomationTokenResponse{}
	var scope string = "global"
	this.Scope = &scope
	return &this
}

// NewAutomationTokenResponseWithDefaults instantiates a new AutomationTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenResponseWithDefaults() *AutomationTokenResponse {
	this := AutomationTokenResponse{}
	var scope string = "global"
	this.Scope = &scope
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AutomationTokenResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AutomationTokenResponse) SetName(v string) {
	o.Name = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AutomationTokenResponse) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponse) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AutomationTokenResponse) SetRole(v string) {
	o.Role = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *AutomationTokenResponse) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponse) GetServicesOk() ([]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *AutomationTokenResponse) SetServices(v []string) {
	o.Services = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *AutomationTokenResponse) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponse) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *AutomationTokenResponse) SetScope(v string) {
	o.Scope = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *AutomationTokenResponse) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponse) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *AutomationTokenResponse) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AutomationTokenResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AutomationTokenResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutomationTokenResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutomationTokenResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *AutomationTokenResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *AutomationTokenResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *AutomationTokenResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AutomationTokenResponse) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AutomationTokenResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *AutomationTokenResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *AutomationTokenResponse) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *AutomationTokenResponse) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *AutomationTokenResponse) GetID() ReadOnlyID {
	if o == nil || o.ID == nil {
		var ret ReadOnlyID
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponse) GetIDOk() (*ReadOnlyID, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given ReadOnlyID and assigns it to the ID field.
func (o *AutomationTokenResponse) SetID(v ReadOnlyID) {
	o.ID = &v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *AutomationTokenResponse) GetCustomerID() ReadOnlyCustomerID {
	if o == nil || o.CustomerID == nil {
		var ret ReadOnlyCustomerID
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponse) GetCustomerIDOk() (*ReadOnlyCustomerID, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given ReadOnlyCustomerID and assigns it to the CustomerID field.
func (o *AutomationTokenResponse) SetCustomerID(v ReadOnlyCustomerID) {
	o.CustomerID = &v
}

// GetIP returns the IP field value if set, zero value otherwise.
func (o *AutomationTokenResponse) GetIP() string {
	if o == nil || o.IP == nil {
		var ret string
		return ret
	}
	return *o.IP
}

// GetIPOk returns a tuple with the IP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponse) GetIPOk() (*string, bool) {
	if o == nil || o.IP == nil {
		return nil, false
	}
	return o.IP, true
}

// HasIP returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasIP() bool {
	if o != nil && o.IP != nil {
		return true
	}

	return false
}

// SetIP gets a reference to the given string and assigns it to the IP field.
func (o *AutomationTokenResponse) SetIP(v string) {
	o.IP = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *AutomationTokenResponse) GetUserAgent() string {
	if o == nil || o.UserAgent == nil {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponse) GetUserAgentOk() (*string, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *AutomationTokenResponse) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetSudoExpiresAt returns the SudoExpiresAt field value if set, zero value otherwise.
func (o *AutomationTokenResponse) GetSudoExpiresAt() string {
	if o == nil || o.SudoExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.SudoExpiresAt
}

// GetSudoExpiresAtOk returns a tuple with the SudoExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponse) GetSudoExpiresAtOk() (*string, bool) {
	if o == nil || o.SudoExpiresAt == nil {
		return nil, false
	}
	return o.SudoExpiresAt, true
}

// HasSudoExpiresAt returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasSudoExpiresAt() bool {
	if o != nil && o.SudoExpiresAt != nil {
		return true
	}

	return false
}

// SetSudoExpiresAt gets a reference to the given string and assigns it to the SudoExpiresAt field.
func (o *AutomationTokenResponse) SetSudoExpiresAt(v string) {
	o.SudoExpiresAt = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *AutomationTokenResponse) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponse) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || o.LastUsedAt == nil {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *AutomationTokenResponse) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt != nil {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *AutomationTokenResponse) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenResponse) MarshalJSON() ([]byte, error) {
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
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.IP != nil {
		toSerialize["ip"] = o.IP
	}
	if o.UserAgent != nil {
		toSerialize["user_agent"] = o.UserAgent
	}
	if o.SudoExpiresAt != nil {
		toSerialize["sudo_expires_at"] = o.SudoExpiresAt
	}
	if o.LastUsedAt != nil {
		toSerialize["last_used_at"] = o.LastUsedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AutomationTokenResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenResponse := _AutomationTokenResponse{}

	if err = json.Unmarshal(bytes, &varAutomationTokenResponse); err == nil {
		*o = AutomationTokenResponse(varAutomationTokenResponse)
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
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "user_agent")
		delete(additionalProperties, "sudo_expires_at")
		delete(additionalProperties, "last_used_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenResponse is a helper abstraction for handling nullable automationtokenresponse types.
type NullableAutomationTokenResponse struct {
	value *AutomationTokenResponse
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenResponse) Get() *AutomationTokenResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenResponse) Set(val *AutomationTokenResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenResponse returns a pointer to a new instance of NullableAutomationTokenResponse.
func NewNullableAutomationTokenResponse(val *AutomationTokenResponse) *NullableAutomationTokenResponse {
	return &NullableAutomationTokenResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAutomationTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
