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
	ID *ReadOnlyID `json:"id,omitempty"`
	CustomerID *ReadOnlyCustomerID `json:"customer_id,omitempty"`
	Role *string `json:"role,omitempty"`
	// The IP address of the client that last used the token.
	IP *string `json:"ip,omitempty"`
	// The User-Agent header of the client that last used the token.
	UserAgent *string `json:"user_agent,omitempty"`
	SudoExpiresAt *string `json:"sudo_expires_at,omitempty"`
	// A UTC time-stamp of when the token was last used.
	LastUsedAt *time.Time `json:"last_used_at,omitempty"`
	// A UTC time-stamp of when the token was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// (optional) A UTC time-stamp of when the token will expire.
	ExpiresAt *string `json:"expires_at,omitempty"`
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

// GetID returns the ID field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetID() ReadOnlyID {
	if o == nil || o.ID == nil {
		var ret ReadOnlyID
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetIDOk() (*ReadOnlyID, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given ReadOnlyID and assigns it to the ID field.
func (o *AutomationTokenResponseAllOf) SetID(v ReadOnlyID) {
	o.ID = &v
}

// GetCustomerID returns the CustomerID field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetCustomerID() ReadOnlyCustomerID {
	if o == nil || o.CustomerID == nil {
		var ret ReadOnlyCustomerID
		return ret
	}
	return *o.CustomerID
}

// GetCustomerIDOk returns a tuple with the CustomerID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetCustomerIDOk() (*ReadOnlyCustomerID, bool) {
	if o == nil || o.CustomerID == nil {
		return nil, false
	}
	return o.CustomerID, true
}

// HasCustomerID returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasCustomerID() bool {
	if o != nil && o.CustomerID != nil {
		return true
	}

	return false
}

// SetCustomerID gets a reference to the given ReadOnlyCustomerID and assigns it to the CustomerID field.
func (o *AutomationTokenResponseAllOf) SetCustomerID(v ReadOnlyCustomerID) {
	o.CustomerID = &v
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

// GetIP returns the IP field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetIP() string {
	if o == nil || o.IP == nil {
		var ret string
		return ret
	}
	return *o.IP
}

// GetIPOk returns a tuple with the IP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetIPOk() (*string, bool) {
	if o == nil || o.IP == nil {
		return nil, false
	}
	return o.IP, true
}

// HasIP returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasIP() bool {
	if o != nil && o.IP != nil {
		return true
	}

	return false
}

// SetIP gets a reference to the given string and assigns it to the IP field.
func (o *AutomationTokenResponseAllOf) SetIP(v string) {
	o.IP = &v
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

// GetSudoExpiresAt returns the SudoExpiresAt field value if set, zero value otherwise.
func (o *AutomationTokenResponseAllOf) GetSudoExpiresAt() string {
	if o == nil || o.SudoExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.SudoExpiresAt
}

// GetSudoExpiresAtOk returns a tuple with the SudoExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenResponseAllOf) GetSudoExpiresAtOk() (*string, bool) {
	if o == nil || o.SudoExpiresAt == nil {
		return nil, false
	}
	return o.SudoExpiresAt, true
}

// HasSudoExpiresAt returns a boolean if a field has been set.
func (o *AutomationTokenResponseAllOf) HasSudoExpiresAt() bool {
	if o != nil && o.SudoExpiresAt != nil {
		return true
	}

	return false
}

// SetSudoExpiresAt gets a reference to the given string and assigns it to the SudoExpiresAt field.
func (o *AutomationTokenResponseAllOf) SetSudoExpiresAt(v string) {
	o.SudoExpiresAt = &v
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
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.CustomerID != nil {
		toSerialize["customer_id"] = o.CustomerID
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
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
		delete(additionalProperties, "sudo_expires_at")
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
