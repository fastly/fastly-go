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

// AutomationTokenResponseAllOf struct for AutomationTokenResponseAllOf
type AutomationTokenResponseAllOf struct {
	Id         *ReadOnlyId         `json:"id,omitempty"`
	CustomerId *ReadOnlyCustomerId `json:"customer_id,omitempty"`
	Role       *string             `json:"role,omitempty"`
	// The IP address of the client that last used the token.
	Ip *string `json:"ip,omitempty"`
	// The User-Agent header of the client that last used the token.
	UserAgent *string `json:"user_agent,omitempty"`
	// Indicates whether TLS access is enabled for the token.
	TlsAccess *bool `json:"tls_access,omitempty"`
	// A UTC timestamp of when the token was last used.
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	// A UTC timestamp of when the token was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// (optional) A UTC timestamp of when the token will expire.
	ExpiresAt            *string `json:"expires_at,omitempty"`
	AdditionalProperties map[string]any
}

type _AutomationTokenResponseAllOf AutomationTokenResponseAllOf

// NewAutomationTokenResponseAllOf instantiates a new AutomationTokenResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenResponseAllOf() *AutomationTokenResponseAllOf {
	this := AutomationTokenResponseAllOf{}
	return &this
}

// NewAutomationTokenResponseAllOfWithDefaults instantiates a new AutomationTokenResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenResponseAllOfWithDefaults() *AutomationTokenResponseAllOf {
	this := AutomationTokenResponseAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetId() ReadOnlyId {
	if o == nil || o.Id == nil {
		var ret ReadOnlyId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetIdOk() (*ReadOnlyId, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given ReadOnlyId and assigns it to the Id field.
func (o *AutomationTokenResponseAllOf) SetId(v ReadOnlyId) {
	o.Id = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetCustomerId() ReadOnlyCustomerId {
	if o == nil || o.CustomerId == nil {
		var ret ReadOnlyCustomerId
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetCustomerIdOk() (*ReadOnlyCustomerId, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given ReadOnlyCustomerId and assigns it to the CustomerId field.
func (o *AutomationTokenResponseAllOf) SetCustomerId(v ReadOnlyCustomerId) {
	o.CustomerId = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetRole() string {
	if o == nil || o.Role == nil {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetRoleOk() (*string, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *AutomationTokenResponseAllOf) SetRole(v string) {
	o.Role = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *AutomationTokenResponseAllOf) SetIp(v string) {
	o.Ip = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetUserAgent() string {
	if o == nil || o.UserAgent == nil {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetUserAgentOk() (*string, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *AutomationTokenResponseAllOf) SetUserAgent(v string) {
	o.UserAgent = &v
}

// GetTlsAccess returns the TlsAccess field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetTlsAccess() bool {
	if o == nil || o.TlsAccess == nil {
		var ret bool
		return ret
	}
	return *o.TlsAccess
}

// GetTlsAccessOk returns a tuple with the TlsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetTlsAccessOk() (*bool, bool) {
	if o == nil || o.TlsAccess == nil {
		return nil, false
	}
	return o.TlsAccess, true
}

// HasTlsAccess returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasTlsAccess() bool {
	if o != nil && o.TlsAccess != nil {
		return true
	}

	return false
}

// SetTlsAccess gets a reference to the given bool and assigns it to the TlsAccess field.
func (o *AutomationTokenResponseAllOf) SetTlsAccess(v bool) {
	o.TlsAccess = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || o.LastUsedAt == nil {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt != nil {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *AutomationTokenResponseAllOf) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AutomationTokenResponseAllOf) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *AutomationTokenResponseAllOf) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CustomerId != nil {
		toSerialize["customer_id"] = o.CustomerId
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.UserAgent != nil {
		toSerialize["user_agent"] = o.UserAgent
	}
	if o.TlsAccess != nil {
		toSerialize["tls_access"] = o.TlsAccess
	}
	if o.LastUsedAt != nil {
		toSerialize["last_used_at"] = o.LastUsedAt
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AutomationTokenResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenResponseAllOf := _AutomationTokenResponseAllOf{}

	if err = json.Unmarshal(bytes, &varAutomationTokenResponseAllOf); err == nil {
		*o = AutomationTokenResponseAllOf(varAutomationTokenResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "role")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "user_agent")
		delete(additionalProperties, "tls_access")
		delete(additionalProperties, "last_used_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "expires_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenResponseAllOf is a helper abstraction for handling nullable automationtokenresponseallof types.
type NullableAutomationTokenResponseAllOf struct {
	value *AutomationTokenResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenResponseAllOf) Get() *AutomationTokenResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenResponseAllOf) Set(val *AutomationTokenResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenResponseAllOf returns a pointer to a new instance of NullableAutomationTokenResponseAllOf.
func NewNullableAutomationTokenResponseAllOf(val *AutomationTokenResponseAllOf) *NullableAutomationTokenResponseAllOf {
	return &NullableAutomationTokenResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAutomationTokenResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
