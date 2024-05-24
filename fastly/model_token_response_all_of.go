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

// TokenResponseAllOf struct for TokenResponseAllOf
type TokenResponseAllOf struct {
	ID *string `json:"id,omitempty"`
	UserID *string `json:"user_id,omitempty"`
	// Time-stamp (UTC) of when the token was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// Time-stamp (UTC) of when the token was last used.
	LastUsedAt *string `json:"last_used_at,omitempty"`
	// Time-stamp (UTC) of when the token will expire (optional).
	ExpiresAt *string `json:"expires_at,omitempty"`
	// IP Address of the client that last used the token.
	IP *string `json:"ip,omitempty"`
	// User-Agent header of the client that last used the token.
	UserAgent *string `json:"user_agent,omitempty"`
	AdditionalProperties map[string]any
}

type _TokenResponseAllOf TokenResponseAllOf

// NewTokenResponseAllOf instantiates a new TokenResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenResponseAllOf() *TokenResponseAllOf {
	this := TokenResponseAllOf{}
	return &this
}

// NewTokenResponseAllOfWithDefaults instantiates a new TokenResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenResponseAllOfWithDefaults() *TokenResponseAllOf {
	this := TokenResponseAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TokenResponseAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TokenResponseAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TokenResponseAllOf) SetID(v string) {
	o.ID = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *TokenResponseAllOf) GetUserID() string {
	if o == nil || o.UserID == nil {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseAllOf) GetUserIDOk() (*string, bool) {
	if o == nil || o.UserID == nil {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *TokenResponseAllOf) HasUserID() bool {
	if o != nil && o.UserID != nil {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *TokenResponseAllOf) SetUserID(v string) {
	o.UserID = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *TokenResponseAllOf) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseAllOf) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TokenResponseAllOf) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *TokenResponseAllOf) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise.
func (o *TokenResponseAllOf) GetLastUsedAt() string {
	if o == nil || o.LastUsedAt == nil {
		var ret string
		return ret
	}
	return *o.LastUsedAt
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseAllOf) GetLastUsedAtOk() (*string, bool) {
	if o == nil || o.LastUsedAt == nil {
		return nil, false
	}
	return o.LastUsedAt, true
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *TokenResponseAllOf) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt != nil {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given string and assigns it to the LastUsedAt field.
func (o *TokenResponseAllOf) SetLastUsedAt(v string) {
	o.LastUsedAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *TokenResponseAllOf) GetExpiresAt() string {
	if o == nil || o.ExpiresAt == nil {
		var ret string
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseAllOf) GetExpiresAtOk() (*string, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *TokenResponseAllOf) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given string and assigns it to the ExpiresAt field.
func (o *TokenResponseAllOf) SetExpiresAt(v string) {
	o.ExpiresAt = &v
}

// GetIP returns the IP field value if set, zero value otherwise.
func (o *TokenResponseAllOf) GetIP() string {
	if o == nil || o.IP == nil {
		var ret string
		return ret
	}
	return *o.IP
}

// GetIPOk returns a tuple with the IP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseAllOf) GetIPOk() (*string, bool) {
	if o == nil || o.IP == nil {
		return nil, false
	}
	return o.IP, true
}

// HasIP returns a boolean if a field has been set.
func (o *TokenResponseAllOf) HasIP() bool {
	if o != nil && o.IP != nil {
		return true
	}

	return false
}

// SetIP gets a reference to the given string and assigns it to the IP field.
func (o *TokenResponseAllOf) SetIP(v string) {
	o.IP = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *TokenResponseAllOf) GetUserAgent() string {
	if o == nil || o.UserAgent == nil {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenResponseAllOf) GetUserAgentOk() (*string, bool) {
	if o == nil || o.UserAgent == nil {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *TokenResponseAllOf) HasUserAgent() bool {
	if o != nil && o.UserAgent != nil {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *TokenResponseAllOf) SetUserAgent(v string) {
	o.UserAgent = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TokenResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.UserID != nil {
		toSerialize["user_id"] = o.UserID
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.LastUsedAt != nil {
		toSerialize["last_used_at"] = o.LastUsedAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.IP != nil {
		toSerialize["ip"] = o.IP
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
func (o *TokenResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTokenResponseAllOf := _TokenResponseAllOf{}

	if err = json.Unmarshal(bytes, &varTokenResponseAllOf); err == nil {
		*o = TokenResponseAllOf(varTokenResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "last_used_at")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "ip")
		delete(additionalProperties, "user_agent")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTokenResponseAllOf is a helper abstraction for handling nullable tokenresponseallof types. 
type NullableTokenResponseAllOf struct {
	value *TokenResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableTokenResponseAllOf) Get() *TokenResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableTokenResponseAllOf) Set(val *TokenResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTokenResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTokenResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTokenResponseAllOf returns a pointer to a new instance of NullableTokenResponseAllOf.
func NewNullableTokenResponseAllOf(val *TokenResponseAllOf) *NullableTokenResponseAllOf {
	return &NullableTokenResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTokenResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTokenResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
