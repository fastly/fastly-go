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
	ID *ReadOnlyID `json:"id,omitempty"`
	UserID *ReadOnlyUserID `json:"user_id,omitempty"`
	CustomerID *ReadOnlyCustomerID `json:"customer_id,omitempty"`
	SudoExpiresAt *time.Time `json:"sudo_expires_at,omitempty"`
	// A UTC time-stamp of when the token was created. 
	CreatedAt *time.Time `json:"created_at,omitempty"`
	AccessToken *string `json:"access_token,omitempty"`
	// A UTC time-stamp of when the token was last used.
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	// The User-Agent header of the client that last used the token.
	UserAgent *string `json:"user_agent,omitempty"`
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

// GetID returns the ID field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponseAllOf) GetID() ReadOnlyID {
	if o == nil || o.ID == nil {
		var ret ReadOnlyID
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponseAllOf) GetIDOk() (*ReadOnlyID, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponseAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given ReadOnlyID and assigns it to the ID field.
func (o *AutomationTokenCreateResponseAllOf) SetID(v ReadOnlyID) {
	o.ID = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponseAllOf) GetUserID() ReadOnlyUserID {
	if o == nil || o.UserID == nil {
		var ret ReadOnlyUserID
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponseAllOf) GetUserIDOk() (*ReadOnlyUserID, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponseAllOf) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given ReadOnlyUserID and assigns it to the UserID field.
func (o *AutomationTokenCreateResponseAllOf) SetUserID(v ReadOnlyUserID) {
	o.UserID = &v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponseAllOf) GetCustomerID() ReadOnlyCustomerID {
	if o == nil || o.CustomerID == nil {
		var ret ReadOnlyCustomerID
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponseAllOf) GetCustomerIDOk() (*ReadOnlyCustomerID, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponseAllOf) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given ReadOnlyCustomerID and assigns it to the CustomerID field.
func (o *AutomationTokenCreateResponseAllOf) SetCustomerID(v ReadOnlyCustomerID) {
	o.CustomerID = &v
}

// GetSudoExpiresAt returns the SudoExpiresAt field value if set, zero value otherwise.
func (o *AutomationTokenCreateResponseAllOf) GetSudoExpiresAt() time.Time {
	if o == nil || o.SudoExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.SudoExpiresAt
}

// GetSudoExpiresAtOk returns a tuple with the SudoExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenCreateResponseAllOf) GetSudoExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.SudoExpiresAt == nil {
		return nil, false
	}
	return o.SudoExpiresAt, true
}

// HasSudoExpiresAt returns a boolean if a field has been set.
func (o *AutomationTokenCreateResponseAllOf) HasSudoExpiresAt() bool {
	if o != nil && o.SudoExpiresAt != nil {
		return true
	}

	return false
}

// SetSudoExpiresAt gets a reference to the given time.Time and assigns it to the SudoExpiresAt field.
func (o *AutomationTokenCreateResponseAllOf) SetSudoExpiresAt(v time.Time) {
	o.SudoExpiresAt = &v
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
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
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
		delete(additionalProperties, "sudo_expires_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "access_token")
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
