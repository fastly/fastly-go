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

// ArcVirtualKey A virtual key used to authenticate requests through the AI Runtime Control gateway.
type ArcVirtualKey struct {
	// A display name for the virtual key.
	Name string `json:"name"`
	// The date and time the key expires, in RFC 3339 format. Omit for a key that doesn't expire.
	ExpiresAt NullableTime `json:"expires_at,omitempty"`
	// Whether AI Firewall security is enabled for requests made with this key.
	SecurityEnabled *bool `json:"security_enabled,omitempty"`
	// The action AI Firewall takes when it detects a violation on a request made with this key. Only meaningful when `security_enabled` is `true`.
	SecurityAction *string `json:"security_action,omitempty"`
	// The requests-per-minute limit for this key. Omit for no limit.
	RpmLimit NullableInt32 `json:"rpm_limit,omitempty"`
	// The tokens-per-minute limit for this key. Omit for no limit.
	TpmLimit             NullableInt32 `json:"tpm_limit,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcVirtualKey ArcVirtualKey

// NewArcVirtualKey instantiates a new ArcVirtualKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcVirtualKey(name string) *ArcVirtualKey {
	this := ArcVirtualKey{}
	this.Name = name
	var securityEnabled bool = false
	this.SecurityEnabled = &securityEnabled
	return &this
}

// NewArcVirtualKeyWithDefaults instantiates a new ArcVirtualKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcVirtualKeyWithDefaults() *ArcVirtualKey {
	this := ArcVirtualKey{}
	var securityEnabled bool = false
	this.SecurityEnabled = &securityEnabled
	return &this
}

// GetName returns the Name field value
func (o *ArcVirtualKey) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKey) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArcVirtualKey) SetName(v string) {
	o.Name = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKey) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKey) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ArcVirtualKey) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *ArcVirtualKey) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *ArcVirtualKey) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *ArcVirtualKey) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetSecurityEnabled returns the SecurityEnabled field value if set, zero value otherwise.
func (o *ArcVirtualKey) GetSecurityEnabled() bool {
	if o == nil || o.SecurityEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SecurityEnabled
}

// GetSecurityEnabledOk returns a tuple with the SecurityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKey) GetSecurityEnabledOk() (*bool, bool) {
	if o == nil || o.SecurityEnabled == nil {
		return nil, false
	}
	return o.SecurityEnabled, true
}

// HasSecurityEnabled returns a boolean if a field has been set.
func (o *ArcVirtualKey) HasSecurityEnabled() bool {
	if o != nil && o.SecurityEnabled != nil {
		return true
	}

	return false
}

// SetSecurityEnabled gets a reference to the given bool and assigns it to the SecurityEnabled field.
func (o *ArcVirtualKey) SetSecurityEnabled(v bool) {
	o.SecurityEnabled = &v
}

// GetSecurityAction returns the SecurityAction field value if set, zero value otherwise.
func (o *ArcVirtualKey) GetSecurityAction() string {
	if o == nil || o.SecurityAction == nil {
		var ret string
		return ret
	}
	return *o.SecurityAction
}

// GetSecurityActionOk returns a tuple with the SecurityAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKey) GetSecurityActionOk() (*string, bool) {
	if o == nil || o.SecurityAction == nil {
		return nil, false
	}
	return o.SecurityAction, true
}

// HasSecurityAction returns a boolean if a field has been set.
func (o *ArcVirtualKey) HasSecurityAction() bool {
	if o != nil && o.SecurityAction != nil {
		return true
	}

	return false
}

// SetSecurityAction gets a reference to the given string and assigns it to the SecurityAction field.
func (o *ArcVirtualKey) SetSecurityAction(v string) {
	o.SecurityAction = &v
}

// GetRpmLimit returns the RpmLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKey) GetRpmLimit() int32 {
	if o == nil || o.RpmLimit.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RpmLimit.Get()
}

// GetRpmLimitOk returns a tuple with the RpmLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKey) GetRpmLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RpmLimit.Get(), o.RpmLimit.IsSet()
}

// HasRpmLimit returns a boolean if a field has been set.
func (o *ArcVirtualKey) HasRpmLimit() bool {
	if o != nil && o.RpmLimit.IsSet() {
		return true
	}

	return false
}

// SetRpmLimit gets a reference to the given NullableInt32 and assigns it to the RpmLimit field.
func (o *ArcVirtualKey) SetRpmLimit(v int32) {
	o.RpmLimit.Set(&v)
}

// SetRpmLimitNil sets the value for RpmLimit to be an explicit nil
func (o *ArcVirtualKey) SetRpmLimitNil() {
	o.RpmLimit.Set(nil)
}

// UnsetRpmLimit ensures that no value is present for RpmLimit, not even an explicit nil
func (o *ArcVirtualKey) UnsetRpmLimit() {
	o.RpmLimit.Unset()
}

// GetTpmLimit returns the TpmLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKey) GetTpmLimit() int32 {
	if o == nil || o.TpmLimit.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TpmLimit.Get()
}

// GetTpmLimitOk returns a tuple with the TpmLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKey) GetTpmLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TpmLimit.Get(), o.TpmLimit.IsSet()
}

// HasTpmLimit returns a boolean if a field has been set.
func (o *ArcVirtualKey) HasTpmLimit() bool {
	if o != nil && o.TpmLimit.IsSet() {
		return true
	}

	return false
}

// SetTpmLimit gets a reference to the given NullableInt32 and assigns it to the TpmLimit field.
func (o *ArcVirtualKey) SetTpmLimit(v int32) {
	o.TpmLimit.Set(&v)
}

// SetTpmLimitNil sets the value for TpmLimit to be an explicit nil
func (o *ArcVirtualKey) SetTpmLimitNil() {
	o.TpmLimit.Set(nil)
}

// UnsetTpmLimit ensures that no value is present for TpmLimit, not even an explicit nil
func (o *ArcVirtualKey) UnsetTpmLimit() {
	o.TpmLimit.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcVirtualKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ExpiresAt.IsSet() {
		toSerialize["expires_at"] = o.ExpiresAt.Get()
	}
	if o.SecurityEnabled != nil {
		toSerialize["security_enabled"] = o.SecurityEnabled
	}
	if o.SecurityAction != nil {
		toSerialize["security_action"] = o.SecurityAction
	}
	if o.RpmLimit.IsSet() {
		toSerialize["rpm_limit"] = o.RpmLimit.Get()
	}
	if o.TpmLimit.IsSet() {
		toSerialize["tpm_limit"] = o.TpmLimit.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcVirtualKey) UnmarshalJSON(bytes []byte) (err error) {
	varArcVirtualKey := _ArcVirtualKey{}

	if err = json.Unmarshal(bytes, &varArcVirtualKey); err == nil {
		*o = ArcVirtualKey(varArcVirtualKey)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "security_enabled")
		delete(additionalProperties, "security_action")
		delete(additionalProperties, "rpm_limit")
		delete(additionalProperties, "tpm_limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcVirtualKey is a helper abstraction for handling nullable arcvirtualkey types.
type NullableArcVirtualKey struct {
	value *ArcVirtualKey
	isSet bool
}

// Get returns the value.
func (v NullableArcVirtualKey) Get() *ArcVirtualKey {
	return v.value
}

// Set modifies the value.
func (v *NullableArcVirtualKey) Set(val *ArcVirtualKey) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcVirtualKey) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcVirtualKey) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcVirtualKey returns a pointer to a new instance of NullableArcVirtualKey.
func NewNullableArcVirtualKey(val *ArcVirtualKey) *NullableArcVirtualKey {
	return &NullableArcVirtualKey{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcVirtualKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcVirtualKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
