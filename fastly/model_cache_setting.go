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

// CacheSetting struct for CacheSetting
type CacheSetting struct {
	// If set, will cause vcl_fetch to terminate after processing this rule with the return state specified. If not set, other configuration logic in vcl_fetch with a lower priority will run after this rule.
	Action NullableString `json:"action,omitempty"`
	// Name of the cache condition controlling when this configuration applies.
	CacheCondition NullableString `json:"cache_condition,omitempty"`
	// Name for the cache settings object.
	Name *string `json:"name,omitempty"`
	// Maximum time in seconds to continue to use a stale version of the object if future requests to your backend server fail (also known as 'stale if error').
	StaleTTL *string `json:"stale_ttl,omitempty"`
	// Maximum time to consider the object fresh in the cache (the cache 'time to live').
	TTL                  *string `json:"ttl,omitempty"`
	AdditionalProperties map[string]any
}

type _CacheSetting CacheSetting

// NewCacheSetting instantiates a new CacheSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCacheSetting() *CacheSetting {
	this := CacheSetting{}
	return &this
}

// NewCacheSettingWithDefaults instantiates a new CacheSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCacheSettingWithDefaults() *CacheSetting {
	this := CacheSetting{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CacheSetting) GetAction() string {
	if o == nil || o.Action.Get() == nil {
		var ret string
		return ret
	}
	return *o.Action.Get()
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CacheSetting) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Action.Get(), o.Action.IsSet()
}

// HasAction returns a boolean if a field has been set.
func (o *CacheSetting) HasAction() bool {
	if o != nil && o.Action.IsSet() {
		return true
	}

	return false
}

// SetAction gets a reference to the given NullableString and assigns it to the Action field.
func (o *CacheSetting) SetAction(v string) {
	o.Action.Set(&v)
}

// SetActionNil sets the value for Action to be an explicit nil
func (o *CacheSetting) SetActionNil() {
	o.Action.Set(nil)
}

// UnsetAction ensures that no value is present for Action, not even an explicit nil
func (o *CacheSetting) UnsetAction() {
	o.Action.Unset()
}

// GetCacheCondition returns the CacheCondition field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CacheSetting) GetCacheCondition() string {
	if o == nil || o.CacheCondition.Get() == nil {
		var ret string
		return ret
	}
	return *o.CacheCondition.Get()
}

// GetCacheConditionOk returns a tuple with the CacheCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CacheSetting) GetCacheConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CacheCondition.Get(), o.CacheCondition.IsSet()
}

// HasCacheCondition returns a boolean if a field has been set.
func (o *CacheSetting) HasCacheCondition() bool {
	if o != nil && o.CacheCondition.IsSet() {
		return true
	}

	return false
}

// SetCacheCondition gets a reference to the given NullableString and assigns it to the CacheCondition field.
func (o *CacheSetting) SetCacheCondition(v string) {
	o.CacheCondition.Set(&v)
}

// SetCacheConditionNil sets the value for CacheCondition to be an explicit nil
func (o *CacheSetting) SetCacheConditionNil() {
	o.CacheCondition.Set(nil)
}

// UnsetCacheCondition ensures that no value is present for CacheCondition, not even an explicit nil
func (o *CacheSetting) UnsetCacheCondition() {
	o.CacheCondition.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CacheSetting) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheSetting) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CacheSetting) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CacheSetting) SetName(v string) {
	o.Name = &v
}

// GetStaleTTL returns the StaleTTL field value if set, zero value otherwise.
func (o *CacheSetting) GetStaleTTL() string {
	if o == nil || o.StaleTTL == nil {
		var ret string
		return ret
	}
	return *o.StaleTTL
}

// GetStaleTTLOk returns a tuple with the StaleTTL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheSetting) GetStaleTTLOk() (*string, bool) {
	if o == nil || o.StaleTTL == nil {
		return nil, false
	}
	return o.StaleTTL, true
}

// HasStaleTTL returns a boolean if a field has been set.
func (o *CacheSetting) HasStaleTTL() bool {
	if o != nil && o.StaleTTL != nil {
		return true
	}

	return false
}

// SetStaleTTL gets a reference to the given string and assigns it to the StaleTTL field.
func (o *CacheSetting) SetStaleTTL(v string) {
	o.StaleTTL = &v
}

// GetTTL returns the TTL field value if set, zero value otherwise.
func (o *CacheSetting) GetTTL() string {
	if o == nil || o.TTL == nil {
		var ret string
		return ret
	}
	return *o.TTL
}

// GetTTLOk returns a tuple with the TTL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CacheSetting) GetTTLOk() (*string, bool) {
	if o == nil || o.TTL == nil {
		return nil, false
	}
	return o.TTL, true
}

// HasTTL returns a boolean if a field has been set.
func (o *CacheSetting) HasTTL() bool {
	if o != nil && o.TTL != nil {
		return true
	}

	return false
}

// SetTTL gets a reference to the given string and assigns it to the TTL field.
func (o *CacheSetting) SetTTL(v string) {
	o.TTL = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o CacheSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Action.IsSet() {
		toSerialize["action"] = o.Action.Get()
	}
	if o.CacheCondition.IsSet() {
		toSerialize["cache_condition"] = o.CacheCondition.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.StaleTTL != nil {
		toSerialize["stale_ttl"] = o.StaleTTL
	}
	if o.TTL != nil {
		toSerialize["ttl"] = o.TTL
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *CacheSetting) UnmarshalJSON(bytes []byte) (err error) {
	varCacheSetting := _CacheSetting{}

	if err = json.Unmarshal(bytes, &varCacheSetting); err == nil {
		*o = CacheSetting(varCacheSetting)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "action")
		delete(additionalProperties, "cache_condition")
		delete(additionalProperties, "name")
		delete(additionalProperties, "stale_ttl")
		delete(additionalProperties, "ttl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableCacheSetting is a helper abstraction for handling nullable cachesetting types.
type NullableCacheSetting struct {
	value *CacheSetting
	isSet bool
}

// Get returns the value.
func (v NullableCacheSetting) Get() *CacheSetting {
	return v.value
}

// Set modifies the value.
func (v *NullableCacheSetting) Set(val *CacheSetting) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableCacheSetting) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableCacheSetting) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableCacheSetting returns a pointer to a new instance of NullableCacheSetting.
func NewNullableCacheSetting(val *CacheSetting) *NullableCacheSetting {
	return &NullableCacheSetting{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableCacheSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableCacheSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
