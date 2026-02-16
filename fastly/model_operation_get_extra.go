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

// OperationGetExtra struct for OperationGetExtra
type OperationGetExtra struct {
	// The unique identifier of the operation.
	Id string `json:"id"`
	// The timestamp when the operation was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// The timestamp when the operation was last updated.
	UpdatedAt time.Time `json:"updated_at"`
	// The timestamp when the operation was last seen in traffic.
	LastSeenAt           *time.Time `json:"last_seen_at,omitempty"`
	AdditionalProperties map[string]any
}

type _OperationGetExtra OperationGetExtra

// NewOperationGetExtra instantiates a new OperationGetExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationGetExtra(id string, updatedAt time.Time) *OperationGetExtra {
	this := OperationGetExtra{}
	this.Id = id
	this.UpdatedAt = updatedAt
	return &this
}

// NewOperationGetExtraWithDefaults instantiates a new OperationGetExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationGetExtraWithDefaults() *OperationGetExtra {
	this := OperationGetExtra{}
	return &this
}

// GetId returns the Id field value
func (o *OperationGetExtra) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OperationGetExtra) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OperationGetExtra) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OperationGetExtra) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationGetExtra) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OperationGetExtra) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OperationGetExtra) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *OperationGetExtra) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OperationGetExtra) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *OperationGetExtra) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *OperationGetExtra) GetLastSeenAt() time.Time {
	if o == nil || o.LastSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationGetExtra) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || o.LastSeenAt == nil {
		return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *OperationGetExtra) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt != nil {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *OperationGetExtra) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OperationGetExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.LastSeenAt != nil {
		toSerialize["last_seen_at"] = o.LastSeenAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *OperationGetExtra) UnmarshalJSON(bytes []byte) (err error) {
	varOperationGetExtra := _OperationGetExtra{}

	if err = json.Unmarshal(bytes, &varOperationGetExtra); err == nil {
		*o = OperationGetExtra(varOperationGetExtra)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "last_seen_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOperationGetExtra is a helper abstraction for handling nullable operationgetextra types.
type NullableOperationGetExtra struct {
	value *OperationGetExtra
	isSet bool
}

// Get returns the value.
func (v NullableOperationGetExtra) Get() *OperationGetExtra {
	return v.value
}

// Set modifies the value.
func (v *NullableOperationGetExtra) Set(val *OperationGetExtra) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOperationGetExtra) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOperationGetExtra) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOperationGetExtra returns a pointer to a new instance of NullableOperationGetExtra.
func NewNullableOperationGetExtra(val *OperationGetExtra) *NullableOperationGetExtra {
	return &NullableOperationGetExtra{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOperationGetExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOperationGetExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
