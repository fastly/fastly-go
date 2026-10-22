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

// ArcVirtualKeyResponse struct for ArcVirtualKeyResponse
type ArcVirtualKeyResponse struct {
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
	TpmLimit NullableInt32 `json:"tpm_limit,omitempty"`
	// The unique identifier of the virtual key.
	Id string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
	// The ID of the user who created the key.
	UserId string `json:"user_id"`
	// The display name of the user who created the key.
	CreatedBy *string `json:"created_by,omitempty"`
	// The date and time the key was last used to make a request, in RFC 3339 format. `null` if the key has never been used.
	LastUsedAt NullableTime `json:"last_used_at,omitempty"`
	// The date and time the key was deleted, in RFC 3339 format. `null` unless the key was deleted and `include_deleted` was used to retrieve it.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// The date and time the key was created, in RFC 3339 format.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the key was last updated, in RFC 3339 format.
	UpdatedAt time.Time `json:"updated_at"`
	// Whether a spend limit is enabled for this key.
	SpendLimitEnabled *bool `json:"spend_limit_enabled,omitempty"`
	// The date and time `spend_limit_enabled` was last set to `true`, in RFC 3339 format. `null` if the spend limit has never been enabled. This timestamp is not cleared when `spend_limit_enabled` is later set back to `false`, so it can be non-null even when the spend limit is currently disabled.
	SpendLimitEnabledOn  NullableTime `json:"spend_limit_enabled_on,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcVirtualKeyResponse ArcVirtualKeyResponse

// NewArcVirtualKeyResponse instantiates a new ArcVirtualKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcVirtualKeyResponse(name string, id string, type_ string, userId string, createdAt time.Time, updatedAt time.Time) *ArcVirtualKeyResponse {
	this := ArcVirtualKeyResponse{}
	this.Name = name
	var securityEnabled bool = false
	this.SecurityEnabled = &securityEnabled
	this.Id = id
	this.Type = type_
	this.UserId = userId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewArcVirtualKeyResponseWithDefaults instantiates a new ArcVirtualKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcVirtualKeyResponseWithDefaults() *ArcVirtualKeyResponse {
	this := ArcVirtualKeyResponse{}
	var securityEnabled bool = false
	this.SecurityEnabled = &securityEnabled
	return &this
}

// GetName returns the Name field value
func (o *ArcVirtualKeyResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ArcVirtualKeyResponse) SetName(v string) {
	o.Name = v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKeyResponse) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt.Get()
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKeyResponse) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt.Get(), o.ExpiresAt.IsSet()
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponse) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given NullableTime and assigns it to the ExpiresAt field.
func (o *ArcVirtualKeyResponse) SetExpiresAt(v time.Time) {
	o.ExpiresAt.Set(&v)
}

// SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil
func (o *ArcVirtualKeyResponse) SetExpiresAtNil() {
	o.ExpiresAt.Set(nil)
}

// UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil
func (o *ArcVirtualKeyResponse) UnsetExpiresAt() {
	o.ExpiresAt.Unset()
}

// GetSecurityEnabled returns the SecurityEnabled field value if set, zero value otherwise.
func (o *ArcVirtualKeyResponse) GetSecurityEnabled() bool {
	if o == nil || o.SecurityEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SecurityEnabled
}

// GetSecurityEnabledOk returns a tuple with the SecurityEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponse) GetSecurityEnabledOk() (*bool, bool) {
	if o == nil || o.SecurityEnabled == nil {
		return nil, false
	}
	return o.SecurityEnabled, true
}

// HasSecurityEnabled returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponse) HasSecurityEnabled() bool {
	if o != nil && o.SecurityEnabled != nil {
		return true
	}

	return false
}

// SetSecurityEnabled gets a reference to the given bool and assigns it to the SecurityEnabled field.
func (o *ArcVirtualKeyResponse) SetSecurityEnabled(v bool) {
	o.SecurityEnabled = &v
}

// GetSecurityAction returns the SecurityAction field value if set, zero value otherwise.
func (o *ArcVirtualKeyResponse) GetSecurityAction() string {
	if o == nil || o.SecurityAction == nil {
		var ret string
		return ret
	}
	return *o.SecurityAction
}

// GetSecurityActionOk returns a tuple with the SecurityAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponse) GetSecurityActionOk() (*string, bool) {
	if o == nil || o.SecurityAction == nil {
		return nil, false
	}
	return o.SecurityAction, true
}

// HasSecurityAction returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponse) HasSecurityAction() bool {
	if o != nil && o.SecurityAction != nil {
		return true
	}

	return false
}

// SetSecurityAction gets a reference to the given string and assigns it to the SecurityAction field.
func (o *ArcVirtualKeyResponse) SetSecurityAction(v string) {
	o.SecurityAction = &v
}

// GetRpmLimit returns the RpmLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKeyResponse) GetRpmLimit() int32 {
	if o == nil || o.RpmLimit.Get() == nil {
		var ret int32
		return ret
	}
	return *o.RpmLimit.Get()
}

// GetRpmLimitOk returns a tuple with the RpmLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKeyResponse) GetRpmLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.RpmLimit.Get(), o.RpmLimit.IsSet()
}

// HasRpmLimit returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponse) HasRpmLimit() bool {
	if o != nil && o.RpmLimit.IsSet() {
		return true
	}

	return false
}

// SetRpmLimit gets a reference to the given NullableInt32 and assigns it to the RpmLimit field.
func (o *ArcVirtualKeyResponse) SetRpmLimit(v int32) {
	o.RpmLimit.Set(&v)
}

// SetRpmLimitNil sets the value for RpmLimit to be an explicit nil
func (o *ArcVirtualKeyResponse) SetRpmLimitNil() {
	o.RpmLimit.Set(nil)
}

// UnsetRpmLimit ensures that no value is present for RpmLimit, not even an explicit nil
func (o *ArcVirtualKeyResponse) UnsetRpmLimit() {
	o.RpmLimit.Unset()
}

// GetTpmLimit returns the TpmLimit field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKeyResponse) GetTpmLimit() int32 {
	if o == nil || o.TpmLimit.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TpmLimit.Get()
}

// GetTpmLimitOk returns a tuple with the TpmLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKeyResponse) GetTpmLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TpmLimit.Get(), o.TpmLimit.IsSet()
}

// HasTpmLimit returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponse) HasTpmLimit() bool {
	if o != nil && o.TpmLimit.IsSet() {
		return true
	}

	return false
}

// SetTpmLimit gets a reference to the given NullableInt32 and assigns it to the TpmLimit field.
func (o *ArcVirtualKeyResponse) SetTpmLimit(v int32) {
	o.TpmLimit.Set(&v)
}

// SetTpmLimitNil sets the value for TpmLimit to be an explicit nil
func (o *ArcVirtualKeyResponse) SetTpmLimitNil() {
	o.TpmLimit.Set(nil)
}

// UnsetTpmLimit ensures that no value is present for TpmLimit, not even an explicit nil
func (o *ArcVirtualKeyResponse) UnsetTpmLimit() {
	o.TpmLimit.Unset()
}

// GetId returns the Id field value
func (o *ArcVirtualKeyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArcVirtualKeyResponse) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ArcVirtualKeyResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArcVirtualKeyResponse) SetType(v string) {
	o.Type = v
}

// GetUserId returns the UserId field value
func (o *ArcVirtualKeyResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ArcVirtualKeyResponse) SetUserId(v string) {
	o.UserId = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ArcVirtualKeyResponse) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponse) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ArcVirtualKeyResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKeyResponse) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt.Get()
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKeyResponse) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsedAt.Get(), o.LastUsedAt.IsSet()
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponse) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt.IsSet() {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given NullableTime and assigns it to the LastUsedAt field.
func (o *ArcVirtualKeyResponse) SetLastUsedAt(v time.Time) {
	o.LastUsedAt.Set(&v)
}

// SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil
func (o *ArcVirtualKeyResponse) SetLastUsedAtNil() {
	o.LastUsedAt.Set(nil)
}

// UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
func (o *ArcVirtualKeyResponse) UnsetLastUsedAt() {
	o.LastUsedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKeyResponse) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKeyResponse) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponse) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *ArcVirtualKeyResponse) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *ArcVirtualKeyResponse) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *ArcVirtualKeyResponse) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *ArcVirtualKeyResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ArcVirtualKeyResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ArcVirtualKeyResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ArcVirtualKeyResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetSpendLimitEnabled returns the SpendLimitEnabled field value if set, zero value otherwise.
func (o *ArcVirtualKeyResponse) GetSpendLimitEnabled() bool {
	if o == nil || o.SpendLimitEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SpendLimitEnabled
}

// GetSpendLimitEnabledOk returns a tuple with the SpendLimitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponse) GetSpendLimitEnabledOk() (*bool, bool) {
	if o == nil || o.SpendLimitEnabled == nil {
		return nil, false
	}
	return o.SpendLimitEnabled, true
}

// HasSpendLimitEnabled returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponse) HasSpendLimitEnabled() bool {
	if o != nil && o.SpendLimitEnabled != nil {
		return true
	}

	return false
}

// SetSpendLimitEnabled gets a reference to the given bool and assigns it to the SpendLimitEnabled field.
func (o *ArcVirtualKeyResponse) SetSpendLimitEnabled(v bool) {
	o.SpendLimitEnabled = &v
}

// GetSpendLimitEnabledOn returns the SpendLimitEnabledOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKeyResponse) GetSpendLimitEnabledOn() time.Time {
	if o == nil || o.SpendLimitEnabledOn.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SpendLimitEnabledOn.Get()
}

// GetSpendLimitEnabledOnOk returns a tuple with the SpendLimitEnabledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKeyResponse) GetSpendLimitEnabledOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpendLimitEnabledOn.Get(), o.SpendLimitEnabledOn.IsSet()
}

// HasSpendLimitEnabledOn returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponse) HasSpendLimitEnabledOn() bool {
	if o != nil && o.SpendLimitEnabledOn.IsSet() {
		return true
	}

	return false
}

// SetSpendLimitEnabledOn gets a reference to the given NullableTime and assigns it to the SpendLimitEnabledOn field.
func (o *ArcVirtualKeyResponse) SetSpendLimitEnabledOn(v time.Time) {
	o.SpendLimitEnabledOn.Set(&v)
}

// SetSpendLimitEnabledOnNil sets the value for SpendLimitEnabledOn to be an explicit nil
func (o *ArcVirtualKeyResponse) SetSpendLimitEnabledOnNil() {
	o.SpendLimitEnabledOn.Set(nil)
}

// UnsetSpendLimitEnabledOn ensures that no value is present for SpendLimitEnabledOn, not even an explicit nil
func (o *ArcVirtualKeyResponse) UnsetSpendLimitEnabledOn() {
	o.SpendLimitEnabledOn.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcVirtualKeyResponse) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if o.CreatedBy != nil {
		toSerialize["created_by"] = o.CreatedBy
	}
	if o.LastUsedAt.IsSet() {
		toSerialize["last_used_at"] = o.LastUsedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.SpendLimitEnabled != nil {
		toSerialize["spend_limit_enabled"] = o.SpendLimitEnabled
	}
	if o.SpendLimitEnabledOn.IsSet() {
		toSerialize["spend_limit_enabled_on"] = o.SpendLimitEnabledOn.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcVirtualKeyResponse) UnmarshalJSON(bytes []byte) (err error) {
	varArcVirtualKeyResponse := _ArcVirtualKeyResponse{}

	if err = json.Unmarshal(bytes, &varArcVirtualKeyResponse); err == nil {
		*o = ArcVirtualKeyResponse(varArcVirtualKeyResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "expires_at")
		delete(additionalProperties, "security_enabled")
		delete(additionalProperties, "security_action")
		delete(additionalProperties, "rpm_limit")
		delete(additionalProperties, "tpm_limit")
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "created_by")
		delete(additionalProperties, "last_used_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "spend_limit_enabled")
		delete(additionalProperties, "spend_limit_enabled_on")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcVirtualKeyResponse is a helper abstraction for handling nullable arcvirtualkeyresponse types.
type NullableArcVirtualKeyResponse struct {
	value *ArcVirtualKeyResponse
	isSet bool
}

// Get returns the value.
func (v NullableArcVirtualKeyResponse) Get() *ArcVirtualKeyResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableArcVirtualKeyResponse) Set(val *ArcVirtualKeyResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcVirtualKeyResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcVirtualKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcVirtualKeyResponse returns a pointer to a new instance of NullableArcVirtualKeyResponse.
func NewNullableArcVirtualKeyResponse(val *ArcVirtualKeyResponse) *NullableArcVirtualKeyResponse {
	return &NullableArcVirtualKeyResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcVirtualKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcVirtualKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
