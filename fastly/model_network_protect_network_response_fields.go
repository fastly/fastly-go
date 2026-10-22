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

// NetworkProtectNetworkResponseFields Read-only fields and required fields in a network response.
type NetworkProtectNetworkResponseFields struct {
	// A read-only unique identifier for a Network Protect resource.
	Id        string                   `json:"id"`
	RoaStatus *NetworkProtectRoaStatus `json:"roa_status,omitempty"`
	// When Fastly last completed a ROA check. Null while the status is pending. A failed check leaves this unchanged; an old value can indicate delayed or failing checks. Null also means the controller did not report a check time.
	RoaCheckedAt NullableTime `json:"roa_checked_at,omitempty"`
	// When the RPKI dataset used by the last check was built. This is not the same as `roa_checked_at`; a recent check can use an older dataset. Null while the status is pending or when no dataset build time was reported.
	RoaDataAsOf NullableTime             `json:"roa_data_as_of,omitempty"`
	LoaStatus   *NetworkProtectLoaStatus `json:"loa_status,omitempty"`
	// An ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at"`
	// An ISO 8601 timestamp in UTC.
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]any
}

type _NetworkProtectNetworkResponseFields NetworkProtectNetworkResponseFields

// NewNetworkProtectNetworkResponseFields instantiates a new NetworkProtectNetworkResponseFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectNetworkResponseFields(id string, createdAt time.Time, updatedAt time.Time) *NetworkProtectNetworkResponseFields {
	this := NetworkProtectNetworkResponseFields{}
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewNetworkProtectNetworkResponseFieldsWithDefaults instantiates a new NetworkProtectNetworkResponseFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectNetworkResponseFieldsWithDefaults() *NetworkProtectNetworkResponseFields {
	this := NetworkProtectNetworkResponseFields{}
	return &this
}

// GetId returns the Id field value
func (o *NetworkProtectNetworkResponseFields) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponseFields) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *NetworkProtectNetworkResponseFields) SetId(v string) {
	o.Id = v
}

// GetRoaStatus returns the RoaStatus field value if set, zero value otherwise.
func (o *NetworkProtectNetworkResponseFields) GetRoaStatus() NetworkProtectRoaStatus {
	if o == nil || o.RoaStatus == nil {
		var ret NetworkProtectRoaStatus
		return ret
	}
	return *o.RoaStatus
}

// GetRoaStatusOk returns a tuple with the RoaStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponseFields) GetRoaStatusOk() (*NetworkProtectRoaStatus, bool) {
	if o == nil || o.RoaStatus == nil {
		return nil, false
	}
	return o.RoaStatus, true
}

// HasRoaStatus returns a boolean if a field has been set.
func (o *NetworkProtectNetworkResponseFields) HasRoaStatus() bool {
	if o != nil && o.RoaStatus != nil {
		return true
	}

	return false
}

// SetRoaStatus gets a reference to the given NetworkProtectRoaStatus and assigns it to the RoaStatus field.
func (o *NetworkProtectNetworkResponseFields) SetRoaStatus(v NetworkProtectRoaStatus) {
	o.RoaStatus = &v
}

// GetRoaCheckedAt returns the RoaCheckedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkProtectNetworkResponseFields) GetRoaCheckedAt() time.Time {
	if o == nil || o.RoaCheckedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RoaCheckedAt.Get()
}

// GetRoaCheckedAtOk returns a tuple with the RoaCheckedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkProtectNetworkResponseFields) GetRoaCheckedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoaCheckedAt.Get(), o.RoaCheckedAt.IsSet()
}

// HasRoaCheckedAt returns a boolean if a field has been set.
func (o *NetworkProtectNetworkResponseFields) HasRoaCheckedAt() bool {
	if o != nil && o.RoaCheckedAt.IsSet() {
		return true
	}

	return false
}

// SetRoaCheckedAt gets a reference to the given NullableTime and assigns it to the RoaCheckedAt field.
func (o *NetworkProtectNetworkResponseFields) SetRoaCheckedAt(v time.Time) {
	o.RoaCheckedAt.Set(&v)
}

// SetRoaCheckedAtNil sets the value for RoaCheckedAt to be an explicit nil
func (o *NetworkProtectNetworkResponseFields) SetRoaCheckedAtNil() {
	o.RoaCheckedAt.Set(nil)
}

// UnsetRoaCheckedAt ensures that no value is present for RoaCheckedAt, not even an explicit nil
func (o *NetworkProtectNetworkResponseFields) UnsetRoaCheckedAt() {
	o.RoaCheckedAt.Unset()
}

// GetRoaDataAsOf returns the RoaDataAsOf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NetworkProtectNetworkResponseFields) GetRoaDataAsOf() time.Time {
	if o == nil || o.RoaDataAsOf.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.RoaDataAsOf.Get()
}

// GetRoaDataAsOfOk returns a tuple with the RoaDataAsOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NetworkProtectNetworkResponseFields) GetRoaDataAsOfOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.RoaDataAsOf.Get(), o.RoaDataAsOf.IsSet()
}

// HasRoaDataAsOf returns a boolean if a field has been set.
func (o *NetworkProtectNetworkResponseFields) HasRoaDataAsOf() bool {
	if o != nil && o.RoaDataAsOf.IsSet() {
		return true
	}

	return false
}

// SetRoaDataAsOf gets a reference to the given NullableTime and assigns it to the RoaDataAsOf field.
func (o *NetworkProtectNetworkResponseFields) SetRoaDataAsOf(v time.Time) {
	o.RoaDataAsOf.Set(&v)
}

// SetRoaDataAsOfNil sets the value for RoaDataAsOf to be an explicit nil
func (o *NetworkProtectNetworkResponseFields) SetRoaDataAsOfNil() {
	o.RoaDataAsOf.Set(nil)
}

// UnsetRoaDataAsOf ensures that no value is present for RoaDataAsOf, not even an explicit nil
func (o *NetworkProtectNetworkResponseFields) UnsetRoaDataAsOf() {
	o.RoaDataAsOf.Unset()
}

// GetLoaStatus returns the LoaStatus field value if set, zero value otherwise.
func (o *NetworkProtectNetworkResponseFields) GetLoaStatus() NetworkProtectLoaStatus {
	if o == nil || o.LoaStatus == nil {
		var ret NetworkProtectLoaStatus
		return ret
	}
	return *o.LoaStatus
}

// GetLoaStatusOk returns a tuple with the LoaStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponseFields) GetLoaStatusOk() (*NetworkProtectLoaStatus, bool) {
	if o == nil || o.LoaStatus == nil {
		return nil, false
	}
	return o.LoaStatus, true
}

// HasLoaStatus returns a boolean if a field has been set.
func (o *NetworkProtectNetworkResponseFields) HasLoaStatus() bool {
	if o != nil && o.LoaStatus != nil {
		return true
	}

	return false
}

// SetLoaStatus gets a reference to the given NetworkProtectLoaStatus and assigns it to the LoaStatus field.
func (o *NetworkProtectNetworkResponseFields) SetLoaStatus(v NetworkProtectLoaStatus) {
	o.LoaStatus = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NetworkProtectNetworkResponseFields) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponseFields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NetworkProtectNetworkResponseFields) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *NetworkProtectNetworkResponseFields) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectNetworkResponseFields) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *NetworkProtectNetworkResponseFields) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectNetworkResponseFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.RoaStatus != nil {
		toSerialize["roa_status"] = o.RoaStatus
	}
	if o.RoaCheckedAt.IsSet() {
		toSerialize["roa_checked_at"] = o.RoaCheckedAt.Get()
	}
	if o.RoaDataAsOf.IsSet() {
		toSerialize["roa_data_as_of"] = o.RoaDataAsOf.Get()
	}
	if o.LoaStatus != nil {
		toSerialize["loa_status"] = o.LoaStatus
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectNetworkResponseFields) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectNetworkResponseFields := _NetworkProtectNetworkResponseFields{}

	if err = json.Unmarshal(bytes, &varNetworkProtectNetworkResponseFields); err == nil {
		*o = NetworkProtectNetworkResponseFields(varNetworkProtectNetworkResponseFields)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "roa_status")
		delete(additionalProperties, "roa_checked_at")
		delete(additionalProperties, "roa_data_as_of")
		delete(additionalProperties, "loa_status")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectNetworkResponseFields is a helper abstraction for handling nullable networkprotectnetworkresponsefields types.
type NullableNetworkProtectNetworkResponseFields struct {
	value *NetworkProtectNetworkResponseFields
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectNetworkResponseFields) Get() *NetworkProtectNetworkResponseFields {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectNetworkResponseFields) Set(val *NetworkProtectNetworkResponseFields) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectNetworkResponseFields) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectNetworkResponseFields) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectNetworkResponseFields returns a pointer to a new instance of NullableNetworkProtectNetworkResponseFields.
func NewNullableNetworkProtectNetworkResponseFields(val *NetworkProtectNetworkResponseFields) *NullableNetworkProtectNetworkResponseFields {
	return &NullableNetworkProtectNetworkResponseFields{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectNetworkResponseFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectNetworkResponseFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
