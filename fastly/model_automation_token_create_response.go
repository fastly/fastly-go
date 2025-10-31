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
	// A UTC timestamp of when the token expires.
	ExpiresAt  *string             `json:"expires_at,omitempty"`
	Id         *ReadOnlyId         `json:"id,omitempty"`
	UserId     *ReadOnlyUserId     `json:"user_id,omitempty"`
	CustomerId *ReadOnlyCustomerId `json:"customer_id,omitempty"`
	// A UTC timestamp of when the token was created.
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	AccessToken *string    `json:"access_token,omitempty"`
	// Indicates whether TLS access is enabled for the token.
	TlsAccess *bool `json:"tls_access,omitempty"`
	// A UTC timestamp of when the token was last used.
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

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetId() ReadOnlyId {
	if o == nil || o.Id == nil {
		var ret ReadOnlyId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetIdOk() (*ReadOnlyId, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given ReadOnlyId and assigns it to the Id field.
func (o *AutomationTokenCreateResponse) SetId(v ReadOnlyId) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetUserId() ReadOnlyUserId {
	if o == nil || o.UserId == nil {
		var ret ReadOnlyUserId
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetUserIdOk() (*ReadOnlyUserId, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given ReadOnlyUserId and assigns it to the UserId field.
func (o *AutomationTokenCreateResponse) SetUserId(v ReadOnlyUserId) {
	o.UserId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetCustomerId() ReadOnlyCustomerId {
	if o == nil || o.CustomerId == nil {
		var ret ReadOnlyCustomerId
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetCustomerIdOk() (*ReadOnlyCustomerId, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given ReadOnlyCustomerId and assigns it to the CustomerId field.
func (o *AutomationTokenCreateResponse) SetCustomerId(v ReadOnlyCustomerId) {
	o.CustomerId = &v
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

// GetTlsAccess returns the TlsAccess field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponse) GetTlsAccess() bool {
	if o == nil || o.TlsAccess == nil {
		var ret bool
		return ret
	}
	return *o.TlsAccess
}

// GetTlsAccessOk returns a tuple with the TlsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponse) GetTlsAccessOk() (*bool, bool) {
	if o == nil || o.TlsAccess == nil {
		return nil, false
	}
	return o.TlsAccess, true
}

// HasTlsAccess returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponse) HasTlsAccess() bool {
	if o != nil && o.TlsAccess != nil {
		return true
	}

	return false
}

// SetTlsAccess gets a reference to the given bool and assigns it to the TlsAccess field.
func (o *AutomationTokenCreateResponse) SetTlsAccess(v bool) {
	o.TlsAccess = &v
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
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.AccessToken != nil {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.TlsAccess != nil {
		toSerialize["tls_access"] = o.TlsAccess
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
		delete(additionalProperties, "id")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "access_token")
		delete(additionalProperties, "tls_access")
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
