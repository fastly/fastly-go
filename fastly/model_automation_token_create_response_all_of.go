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

// AutomationTokenCreateResponseAllOf struct for AutomationTokenCreateResponseAllOf
type AutomationTokenCreateResponseAllOf struct {
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

type _AutomationTokenCreateResponseAllOf AutomationTokenCreateResponseAllOf

// NewAutomationTokenCreateResponseAllOf instantiates a new AutomationTokenCreateResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenCreateResponseAllOf() *AutomationTokenCreateResponseAllOf {
	this := AutomationTokenCreateResponseAllOf{}
	return &this
}

// NewAutomationTokenCreateResponseAllOfWithDefaults instantiates a new AutomationTokenCreateResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenCreateResponseAllOfWithDefaults() *AutomationTokenCreateResponseAllOf {
	this := AutomationTokenCreateResponseAllOf{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponseAllOf) GetId() ReadOnlyId {
	if o == nil || o.Id == nil {
		var ret ReadOnlyId
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponseAllOf) GetIdOk() (*ReadOnlyId, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponseAllOf) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given ReadOnlyId and assigns it to the Id field.
func (o *AutomationTokenCreateResponseAllOf) SetId(v ReadOnlyId) {
	o.Id = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponseAllOf) GetUserId() ReadOnlyUserId {
	if o == nil || o.UserId == nil {
		var ret ReadOnlyUserId
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponseAllOf) GetUserIdOk() (*ReadOnlyUserId, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponseAllOf) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given ReadOnlyUserId and assigns it to the UserId field.
func (o *AutomationTokenCreateResponseAllOf) SetUserId(v ReadOnlyUserId) {
	o.UserId = &v
}

// GetCustomerId returns the CustomerId field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponseAllOf) GetCustomerId() ReadOnlyCustomerId {
	if o == nil || o.CustomerId == nil {
		var ret ReadOnlyCustomerId
		return ret
	}
	return *o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponseAllOf) GetCustomerIdOk() (*ReadOnlyCustomerId, bool) {
	if o == nil || o.CustomerId == nil {
		return nil, false
	}
	return o.CustomerId, true
}

// HasCustomerId returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponseAllOf) HasCustomerId() bool {
	if o != nil && o.CustomerId != nil {
		return true
	}

	return false
}

// SetCustomerId gets a reference to the given ReadOnlyCustomerId and assigns it to the CustomerId field.
func (o *AutomationTokenCreateResponseAllOf) SetCustomerId(v ReadOnlyCustomerId) {
	o.CustomerId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponseAllOf) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponseAllOf) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponseAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AutomationTokenCreateResponseAllOf) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponseAllOf) GetAccessToken() string {
	if o == nil || o.AccessToken == nil {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponseAllOf) GetAccessTokenOk() (*string, bool) {
	if o == nil || o.AccessToken == nil {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponseAllOf) HasAccessToken() bool {
	if o != nil && o.AccessToken != nil {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *AutomationTokenCreateResponseAllOf) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetTlsAccess returns the TlsAccess field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponseAllOf) GetTlsAccess() bool {
	if o == nil || o.TlsAccess == nil {
		var ret bool
		return ret
	}
	return *o.TlsAccess
}

// GetTlsAccessOk returns a tuple with the TlsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponseAllOf) GetTlsAccessOk() (*bool, bool) {
	if o == nil || o.TlsAccess == nil {
		return nil, false
	}
	return o.TlsAccess, true
}

// HasTlsAccess returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponseAllOf) HasTlsAccess() bool {
	if o != nil && o.TlsAccess != nil {
		return true
	}

	return false
}

// SetTlsAccess gets a reference to the given bool and assigns it to the TlsAccess field.
func (o *AutomationTokenCreateResponseAllOf) SetTlsAccess(v bool) {
	o.TlsAccess = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponseAllOf) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponseAllOf) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil || o.LastUsedAt == nil {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponseAllOf) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt != nil {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given time.Time and assigns it to the LastUsedAt field.
func (o *AutomationTokenCreateResponseAllOf) SetLastUsedAt(v time.Time) {
	o.LastUsedAt = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponseAllOf) GetUserAgent() string {
	if o == nil || o.UserAgent == nil {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponseAllOf) GetUserAgentOk() (*string, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponseAllOf) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *AutomationTokenCreateResponseAllOf) SetUserAgent(v string) {
	o.UserAgent = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenCreateResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *AutomationTokenCreateResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenCreateResponseAllOf := _AutomationTokenCreateResponseAllOf{}

	if err = json.Unmarshal(bytes, &varAutomationTokenCreateResponseAllOf); err == nil {
		*o = AutomationTokenCreateResponseAllOf(varAutomationTokenCreateResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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

// NullableAutomationTokenCreateResponseAllOf is a helper abstraction for handling nullable automationtokencreateresponseallof types.
type NullableAutomationTokenCreateResponseAllOf struct {
	value *AutomationTokenCreateResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenCreateResponseAllOf) Get() *AutomationTokenCreateResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenCreateResponseAllOf) Set(val *AutomationTokenCreateResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenCreateResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenCreateResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenCreateResponseAllOf returns a pointer to a new instance of NullableAutomationTokenCreateResponseAllOf.
func NewNullableAutomationTokenCreateResponseAllOf(val *AutomationTokenCreateResponseAllOf) *NullableAutomationTokenCreateResponseAllOf {
	return &NullableAutomationTokenCreateResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenCreateResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAutomationTokenCreateResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
