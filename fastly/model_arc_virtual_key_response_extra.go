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

// ArcVirtualKeyResponseExtra struct for ArcVirtualKeyResponseExtra
type ArcVirtualKeyResponseExtra struct {
	// The unique identifier of the virtual key.
	ID string `json:"id"`
	// The type of the resource.
	Type string `json:"type"`
	// The ID of the user who created the key.
	UserID string `json:"user_id"`
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

type _ArcVirtualKeyResponseExtra ArcVirtualKeyResponseExtra

// NewArcVirtualKeyResponseExtra instantiates a new ArcVirtualKeyResponseExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcVirtualKeyResponseExtra(id string, resourceType string, userID string, createdAt time.Time, updatedAt time.Time) *ArcVirtualKeyResponseExtra {
	this := ArcVirtualKeyResponseExtra{}
	this.ID = id
	this.Type = resourceType
	this.UserID = userID
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewArcVirtualKeyResponseExtraWithDefaults instantiates a new ArcVirtualKeyResponseExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcVirtualKeyResponseExtraWithDefaults() *ArcVirtualKeyResponseExtra {
	this := ArcVirtualKeyResponseExtra{}
	return &this
}

// GetID returns the ID field value
func (o *ArcVirtualKeyResponseExtra) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponseExtra) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *ArcVirtualKeyResponseExtra) SetID(v string) {
	o.ID = v
}

// GetType returns the Type field value
func (o *ArcVirtualKeyResponseExtra) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponseExtra) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArcVirtualKeyResponseExtra) SetType(v string) {
	o.Type = v
}

// GetUserID returns the UserID field value
func (o *ArcVirtualKeyResponseExtra) GetUserID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponseExtra) GetUserIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserID, true
}

// SetUserID sets field value
func (o *ArcVirtualKeyResponseExtra) SetUserID(v string) {
	o.UserID = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ArcVirtualKeyResponseExtra) GetCreatedBy() string {
	if o == nil || o.CreatedBy == nil {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponseExtra) GetCreatedByOk() (*string, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponseExtra) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ArcVirtualKeyResponseExtra) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKeyResponseExtra) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt.Get()
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKeyResponseExtra) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsedAt.Get(), o.LastUsedAt.IsSet()
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponseExtra) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt.IsSet() {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given NullableTime and assigns it to the LastUsedAt field.
func (o *ArcVirtualKeyResponseExtra) SetLastUsedAt(v time.Time) {
	o.LastUsedAt.Set(&v)
}

// SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil
func (o *ArcVirtualKeyResponseExtra) SetLastUsedAtNil() {
	o.LastUsedAt.Set(nil)
}

// UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
func (o *ArcVirtualKeyResponseExtra) UnsetLastUsedAt() {
	o.LastUsedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKeyResponseExtra) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKeyResponseExtra) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponseExtra) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *ArcVirtualKeyResponseExtra) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *ArcVirtualKeyResponseExtra) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *ArcVirtualKeyResponseExtra) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *ArcVirtualKeyResponseExtra) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponseExtra) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ArcVirtualKeyResponseExtra) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ArcVirtualKeyResponseExtra) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponseExtra) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ArcVirtualKeyResponseExtra) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetSpendLimitEnabled returns the SpendLimitEnabled field value if set, zero value otherwise.
func (o *ArcVirtualKeyResponseExtra) GetSpendLimitEnabled() bool {
	if o == nil || o.SpendLimitEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SpendLimitEnabled
}

// GetSpendLimitEnabledOk returns a tuple with the SpendLimitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyResponseExtra) GetSpendLimitEnabledOk() (*bool, bool) {
	if o == nil || o.SpendLimitEnabled == nil {
		return nil, false
	}
	return o.SpendLimitEnabled, true
}

// HasSpendLimitEnabled returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponseExtra) HasSpendLimitEnabled() bool {
	if o != nil && o.SpendLimitEnabled != nil {
		return true
	}

	return false
}

// SetSpendLimitEnabled gets a reference to the given bool and assigns it to the SpendLimitEnabled field.
func (o *ArcVirtualKeyResponseExtra) SetSpendLimitEnabled(v bool) {
	o.SpendLimitEnabled = &v
}

// GetSpendLimitEnabledOn returns the SpendLimitEnabledOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKeyResponseExtra) GetSpendLimitEnabledOn() time.Time {
	if o == nil || o.SpendLimitEnabledOn.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SpendLimitEnabledOn.Get()
}

// GetSpendLimitEnabledOnOk returns a tuple with the SpendLimitEnabledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKeyResponseExtra) GetSpendLimitEnabledOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpendLimitEnabledOn.Get(), o.SpendLimitEnabledOn.IsSet()
}

// HasSpendLimitEnabledOn returns a boolean if a field has been set.
func (o *ArcVirtualKeyResponseExtra) HasSpendLimitEnabledOn() bool {
	if o != nil && o.SpendLimitEnabledOn.IsSet() {
		return true
	}

	return false
}

// SetSpendLimitEnabledOn gets a reference to the given NullableTime and assigns it to the SpendLimitEnabledOn field.
func (o *ArcVirtualKeyResponseExtra) SetSpendLimitEnabledOn(v time.Time) {
	o.SpendLimitEnabledOn.Set(&v)
}

// SetSpendLimitEnabledOnNil sets the value for SpendLimitEnabledOn to be an explicit nil
func (o *ArcVirtualKeyResponseExtra) SetSpendLimitEnabledOnNil() {
	o.SpendLimitEnabledOn.Set(nil)
}

// UnsetSpendLimitEnabledOn ensures that no value is present for SpendLimitEnabledOn, not even an explicit nil
func (o *ArcVirtualKeyResponseExtra) UnsetSpendLimitEnabledOn() {
	o.SpendLimitEnabledOn.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcVirtualKeyResponseExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.ID
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["user_id"] = o.UserID
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
func (o *ArcVirtualKeyResponseExtra) UnmarshalJSON(bytes []byte) (err error) {
	varArcVirtualKeyResponseExtra := _ArcVirtualKeyResponseExtra{}

	if err = json.Unmarshal(bytes, &varArcVirtualKeyResponseExtra); err == nil {
		*o = ArcVirtualKeyResponseExtra(varArcVirtualKeyResponseExtra)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
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

// NullableArcVirtualKeyResponseExtra is a helper abstraction for handling nullable arcvirtualkeyresponseextra types.
type NullableArcVirtualKeyResponseExtra struct {
	value *ArcVirtualKeyResponseExtra
	isSet bool
}

// Get returns the value.
func (v NullableArcVirtualKeyResponseExtra) Get() *ArcVirtualKeyResponseExtra {
	return v.value
}

// Set modifies the value.
func (v *NullableArcVirtualKeyResponseExtra) Set(val *ArcVirtualKeyResponseExtra) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcVirtualKeyResponseExtra) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcVirtualKeyResponseExtra) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcVirtualKeyResponseExtra returns a pointer to a new instance of NullableArcVirtualKeyResponseExtra.
func NewNullableArcVirtualKeyResponseExtra(val *ArcVirtualKeyResponseExtra) *NullableArcVirtualKeyResponseExtra {
	return &NullableArcVirtualKeyResponseExtra{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcVirtualKeyResponseExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcVirtualKeyResponseExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
