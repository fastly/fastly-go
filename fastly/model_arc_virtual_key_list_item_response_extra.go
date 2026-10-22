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

// ArcVirtualKeyListItemResponseExtra struct for ArcVirtualKeyListItemResponseExtra
type ArcVirtualKeyListItemResponseExtra struct {
	// The unique identifier of the virtual key.
	Id string `json:"id"`
	// The ID of the customer account that owns the key.
	CustomerId string `json:"customer_id"`
	// The type of the resource.
	Type string `json:"type"`
	// The ID of the user who created the key.
	UserId string `json:"user_id"`
	// The display name of the user who created the key.
	UserName *string `json:"user_name,omitempty"`
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

type _ArcVirtualKeyListItemResponseExtra ArcVirtualKeyListItemResponseExtra

// NewArcVirtualKeyListItemResponseExtra instantiates a new ArcVirtualKeyListItemResponseExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcVirtualKeyListItemResponseExtra(id string, customerId string, type_ string, userId string, createdAt time.Time, updatedAt time.Time) *ArcVirtualKeyListItemResponseExtra {
	this := ArcVirtualKeyListItemResponseExtra{}
	this.Id = id
	this.CustomerId = customerId
	this.Type = type_
	this.UserId = userId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewArcVirtualKeyListItemResponseExtraWithDefaults instantiates a new ArcVirtualKeyListItemResponseExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcVirtualKeyListItemResponseExtraWithDefaults() *ArcVirtualKeyListItemResponseExtra {
	this := ArcVirtualKeyListItemResponseExtra{}
	return &this
}

// GetId returns the Id field value
func (o *ArcVirtualKeyListItemResponseExtra) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyListItemResponseExtra) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArcVirtualKeyListItemResponseExtra) SetId(v string) {
	o.Id = v
}

// GetCustomerId returns the CustomerId field value
func (o *ArcVirtualKeyListItemResponseExtra) GetCustomerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomerId
}

// GetCustomerIdOk returns a tuple with the CustomerId field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyListItemResponseExtra) GetCustomerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomerId, true
}

// SetCustomerId sets field value
func (o *ArcVirtualKeyListItemResponseExtra) SetCustomerId(v string) {
	o.CustomerId = v
}

// GetType returns the Type field value
func (o *ArcVirtualKeyListItemResponseExtra) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyListItemResponseExtra) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ArcVirtualKeyListItemResponseExtra) SetType(v string) {
	o.Type = v
}

// GetUserId returns the UserId field value
func (o *ArcVirtualKeyListItemResponseExtra) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyListItemResponseExtra) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ArcVirtualKeyListItemResponseExtra) SetUserId(v string) {
	o.UserId = v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *ArcVirtualKeyListItemResponseExtra) GetUserName() string {
	if o == nil || o.UserName == nil {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyListItemResponseExtra) GetUserNameOk() (*string, bool) {
	if o == nil || o.UserName == nil {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *ArcVirtualKeyListItemResponseExtra) HasUserName() bool {
	if o != nil && o.UserName != nil {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *ArcVirtualKeyListItemResponseExtra) SetUserName(v string) {
	o.UserName = &v
}

// GetLastUsedAt returns the LastUsedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKeyListItemResponseExtra) GetLastUsedAt() time.Time {
	if o == nil || o.LastUsedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUsedAt.Get()
}

// GetLastUsedAtOk returns a tuple with the LastUsedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKeyListItemResponseExtra) GetLastUsedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUsedAt.Get(), o.LastUsedAt.IsSet()
}

// HasLastUsedAt returns a boolean if a field has been set.
func (o *ArcVirtualKeyListItemResponseExtra) HasLastUsedAt() bool {
	if o != nil && o.LastUsedAt.IsSet() {
		return true
	}

	return false
}

// SetLastUsedAt gets a reference to the given NullableTime and assigns it to the LastUsedAt field.
func (o *ArcVirtualKeyListItemResponseExtra) SetLastUsedAt(v time.Time) {
	o.LastUsedAt.Set(&v)
}

// SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil
func (o *ArcVirtualKeyListItemResponseExtra) SetLastUsedAtNil() {
	o.LastUsedAt.Set(nil)
}

// UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
func (o *ArcVirtualKeyListItemResponseExtra) UnsetLastUsedAt() {
	o.LastUsedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKeyListItemResponseExtra) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKeyListItemResponseExtra) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *ArcVirtualKeyListItemResponseExtra) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *ArcVirtualKeyListItemResponseExtra) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *ArcVirtualKeyListItemResponseExtra) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *ArcVirtualKeyListItemResponseExtra) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetCreatedAt returns the CreatedAt field value
func (o *ArcVirtualKeyListItemResponseExtra) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyListItemResponseExtra) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ArcVirtualKeyListItemResponseExtra) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ArcVirtualKeyListItemResponseExtra) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyListItemResponseExtra) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ArcVirtualKeyListItemResponseExtra) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetSpendLimitEnabled returns the SpendLimitEnabled field value if set, zero value otherwise.
func (o *ArcVirtualKeyListItemResponseExtra) GetSpendLimitEnabled() bool {
	if o == nil || o.SpendLimitEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SpendLimitEnabled
}

// GetSpendLimitEnabledOk returns a tuple with the SpendLimitEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyListItemResponseExtra) GetSpendLimitEnabledOk() (*bool, bool) {
	if o == nil || o.SpendLimitEnabled == nil {
		return nil, false
	}
	return o.SpendLimitEnabled, true
}

// HasSpendLimitEnabled returns a boolean if a field has been set.
func (o *ArcVirtualKeyListItemResponseExtra) HasSpendLimitEnabled() bool {
	if o != nil && o.SpendLimitEnabled != nil {
		return true
	}

	return false
}

// SetSpendLimitEnabled gets a reference to the given bool and assigns it to the SpendLimitEnabled field.
func (o *ArcVirtualKeyListItemResponseExtra) SetSpendLimitEnabled(v bool) {
	o.SpendLimitEnabled = &v
}

// GetSpendLimitEnabledOn returns the SpendLimitEnabledOn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ArcVirtualKeyListItemResponseExtra) GetSpendLimitEnabledOn() time.Time {
	if o == nil || o.SpendLimitEnabledOn.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SpendLimitEnabledOn.Get()
}

// GetSpendLimitEnabledOnOk returns a tuple with the SpendLimitEnabledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ArcVirtualKeyListItemResponseExtra) GetSpendLimitEnabledOnOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpendLimitEnabledOn.Get(), o.SpendLimitEnabledOn.IsSet()
}

// HasSpendLimitEnabledOn returns a boolean if a field has been set.
func (o *ArcVirtualKeyListItemResponseExtra) HasSpendLimitEnabledOn() bool {
	if o != nil && o.SpendLimitEnabledOn.IsSet() {
		return true
	}

	return false
}

// SetSpendLimitEnabledOn gets a reference to the given NullableTime and assigns it to the SpendLimitEnabledOn field.
func (o *ArcVirtualKeyListItemResponseExtra) SetSpendLimitEnabledOn(v time.Time) {
	o.SpendLimitEnabledOn.Set(&v)
}

// SetSpendLimitEnabledOnNil sets the value for SpendLimitEnabledOn to be an explicit nil
func (o *ArcVirtualKeyListItemResponseExtra) SetSpendLimitEnabledOnNil() {
	o.SpendLimitEnabledOn.Set(nil)
}

// UnsetSpendLimitEnabledOn ensures that no value is present for SpendLimitEnabledOn, not even an explicit nil
func (o *ArcVirtualKeyListItemResponseExtra) UnsetSpendLimitEnabledOn() {
	o.SpendLimitEnabledOn.Unset()
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcVirtualKeyListItemResponseExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["customer_id"] = o.CustomerId
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["user_id"] = o.UserId
	}
	if o.UserName != nil {
		toSerialize["user_name"] = o.UserName
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
func (o *ArcVirtualKeyListItemResponseExtra) UnmarshalJSON(bytes []byte) (err error) {
	varArcVirtualKeyListItemResponseExtra := _ArcVirtualKeyListItemResponseExtra{}

	if err = json.Unmarshal(bytes, &varArcVirtualKeyListItemResponseExtra); err == nil {
		*o = ArcVirtualKeyListItemResponseExtra(varArcVirtualKeyListItemResponseExtra)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "customer_id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "user_id")
		delete(additionalProperties, "user_name")
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

// NullableArcVirtualKeyListItemResponseExtra is a helper abstraction for handling nullable arcvirtualkeylistitemresponseextra types.
type NullableArcVirtualKeyListItemResponseExtra struct {
	value *ArcVirtualKeyListItemResponseExtra
	isSet bool
}

// Get returns the value.
func (v NullableArcVirtualKeyListItemResponseExtra) Get() *ArcVirtualKeyListItemResponseExtra {
	return v.value
}

// Set modifies the value.
func (v *NullableArcVirtualKeyListItemResponseExtra) Set(val *ArcVirtualKeyListItemResponseExtra) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcVirtualKeyListItemResponseExtra) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcVirtualKeyListItemResponseExtra) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcVirtualKeyListItemResponseExtra returns a pointer to a new instance of NullableArcVirtualKeyListItemResponseExtra.
func NewNullableArcVirtualKeyListItemResponseExtra(val *ArcVirtualKeyListItemResponseExtra) *NullableArcVirtualKeyListItemResponseExtra {
	return &NullableArcVirtualKeyListItemResponseExtra{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcVirtualKeyListItemResponseExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcVirtualKeyListItemResponseExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
