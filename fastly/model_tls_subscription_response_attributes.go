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

// TlsSubscriptionResponseAttributes struct for TlsSubscriptionResponseAttributes
type TlsSubscriptionResponseAttributes struct {
	// Date and time in ISO 8601 format.
	CreatedAt NullableTime `json:"created_at,omitempty"`
	// Date and time in ISO 8601 format.
	DeletedAt NullableTime `json:"deleted_at,omitempty"`
	// Date and time in ISO 8601 format.
	UpdatedAt NullableTime `json:"updated_at,omitempty"`
	// The current state of your subscription.
	State *string `json:"state,omitempty"`
	// Subscription has an active order
	HasActiveOrder       *bool `json:"has_active_order,omitempty"`
	AdditionalProperties map[string]any
}

type _TlsSubscriptionResponseAttributes TlsSubscriptionResponseAttributes

// NewTlsSubscriptionResponseAttributes instantiates a new TlsSubscriptionResponseAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTlsSubscriptionResponseAttributes() *TlsSubscriptionResponseAttributes {
	this := TlsSubscriptionResponseAttributes{}
	return &this
}

// NewTlsSubscriptionResponseAttributesWithDefaults instantiates a new TlsSubscriptionResponseAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTlsSubscriptionResponseAttributesWithDefaults() *TlsSubscriptionResponseAttributes {
	this := TlsSubscriptionResponseAttributes{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsSubscriptionResponseAttributes) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsSubscriptionResponseAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *TlsSubscriptionResponseAttributes) HasCreatedAt() bool {
	if o != nil && o.CreatedAt.IsSet() {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given NullableTime and assigns it to the CreatedAt field.
func (o *TlsSubscriptionResponseAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// SetCreatedAtNil sets the value for CreatedAt to be an explicit nil
func (o *TlsSubscriptionResponseAttributes) SetCreatedAtNil() {
	o.CreatedAt.Set(nil)
}

// UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
func (o *TlsSubscriptionResponseAttributes) UnsetCreatedAt() {
	o.CreatedAt.Unset()
}

// GetDeletedAt returns the DeletedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsSubscriptionResponseAttributes) GetDeletedAt() time.Time {
	if o == nil || o.DeletedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.DeletedAt.Get()
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsSubscriptionResponseAttributes) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DeletedAt.Get(), o.DeletedAt.IsSet()
}

// HasDeletedAt returns a boolean if a field has been set.
func (o *TlsSubscriptionResponseAttributes) HasDeletedAt() bool {
	if o != nil && o.DeletedAt.IsSet() {
		return true
	}

	return false
}

// SetDeletedAt gets a reference to the given NullableTime and assigns it to the DeletedAt field.
func (o *TlsSubscriptionResponseAttributes) SetDeletedAt(v time.Time) {
	o.DeletedAt.Set(&v)
}

// SetDeletedAtNil sets the value for DeletedAt to be an explicit nil
func (o *TlsSubscriptionResponseAttributes) SetDeletedAtNil() {
	o.DeletedAt.Set(nil)
}

// UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
func (o *TlsSubscriptionResponseAttributes) UnsetDeletedAt() {
	o.DeletedAt.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TlsSubscriptionResponseAttributes) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TlsSubscriptionResponseAttributes) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *TlsSubscriptionResponseAttributes) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt.IsSet() {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given NullableTime and assigns it to the UpdatedAt field.
func (o *TlsSubscriptionResponseAttributes) SetUpdatedAt(v time.Time) {
	o.UpdatedAt.Set(&v)
}

// SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil
func (o *TlsSubscriptionResponseAttributes) SetUpdatedAtNil() {
	o.UpdatedAt.Set(nil)
}

// UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
func (o *TlsSubscriptionResponseAttributes) UnsetUpdatedAt() {
	o.UpdatedAt.Unset()
}

// GetState returns the State field value if set, zero value otherwise.
func (o *TlsSubscriptionResponseAttributes) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsSubscriptionResponseAttributes) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *TlsSubscriptionResponseAttributes) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *TlsSubscriptionResponseAttributes) SetState(v string) {
	o.State = &v
}

// GetHasActiveOrder returns the HasActiveOrder field value if set, zero value otherwise.
func (o *TlsSubscriptionResponseAttributes) GetHasActiveOrder() bool {
	if o == nil || o.HasActiveOrder == nil {
		var ret bool
		return ret
	}
	return *o.HasActiveOrder
}

// GetHasActiveOrderOk returns a tuple with the HasActiveOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TlsSubscriptionResponseAttributes) GetHasActiveOrderOk() (*bool, bool) {
	if o == nil || o.HasActiveOrder == nil {
		return nil, false
	}
	return o.HasActiveOrder, true
}

// HasHasActiveOrder returns a boolean if a field has been set.
func (o *TlsSubscriptionResponseAttributes) HasHasActiveOrder() bool {
	if o != nil && o.HasActiveOrder != nil {
		return true
	}

	return false
}

// SetHasActiveOrder gets a reference to the given bool and assigns it to the HasActiveOrder field.
func (o *TlsSubscriptionResponseAttributes) SetHasActiveOrder(v bool) {
	o.HasActiveOrder = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TlsSubscriptionResponseAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt.IsSet() {
		toSerialize["created_at"] = o.CreatedAt.Get()
	}
	if o.DeletedAt.IsSet() {
		toSerialize["deleted_at"] = o.DeletedAt.Get()
	}
	if o.UpdatedAt.IsSet() {
		toSerialize["updated_at"] = o.UpdatedAt.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.HasActiveOrder != nil {
		toSerialize["has_active_order"] = o.HasActiveOrder
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TlsSubscriptionResponseAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varTlsSubscriptionResponseAttributes := _TlsSubscriptionResponseAttributes{}

	if err = json.Unmarshal(bytes, &varTlsSubscriptionResponseAttributes); err == nil {
		*o = TlsSubscriptionResponseAttributes(varTlsSubscriptionResponseAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "deleted_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "state")
		delete(additionalProperties, "has_active_order")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTlsSubscriptionResponseAttributes is a helper abstraction for handling nullable tlssubscriptionresponseattributes types.
type NullableTlsSubscriptionResponseAttributes struct {
	value *TlsSubscriptionResponseAttributes
	isSet bool
}

// Get returns the value.
func (v NullableTlsSubscriptionResponseAttributes) Get() *TlsSubscriptionResponseAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableTlsSubscriptionResponseAttributes) Set(val *TlsSubscriptionResponseAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTlsSubscriptionResponseAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTlsSubscriptionResponseAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTlsSubscriptionResponseAttributes returns a pointer to a new instance of NullableTlsSubscriptionResponseAttributes.
func NewNullableTlsSubscriptionResponseAttributes(val *TlsSubscriptionResponseAttributes) *NullableTlsSubscriptionResponseAttributes {
	return &NullableTlsSubscriptionResponseAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTlsSubscriptionResponseAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTlsSubscriptionResponseAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
