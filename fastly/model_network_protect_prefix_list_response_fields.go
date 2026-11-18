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

// NetworkProtectPrefixListResponseFields Read-only fields and required fields in a prefix-list response.
type NetworkProtectPrefixListResponseFields struct {
	// A read-only unique identifier for a Network Protect resource.
	ID string `json:"id"`
	// An ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at"`
	// An ISO 8601 timestamp in UTC.
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]any
}

type _NetworkProtectPrefixListResponseFields NetworkProtectPrefixListResponseFields

// NewNetworkProtectPrefixListResponseFields instantiates a new NetworkProtectPrefixListResponseFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectPrefixListResponseFields(id string, createdAt time.Time, updatedAt time.Time) *NetworkProtectPrefixListResponseFields {
	this := NetworkProtectPrefixListResponseFields{}
	this.ID = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewNetworkProtectPrefixListResponseFieldsWithDefaults instantiates a new NetworkProtectPrefixListResponseFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectPrefixListResponseFieldsWithDefaults() *NetworkProtectPrefixListResponseFields {
	this := NetworkProtectPrefixListResponseFields{}
	return &this
}

// GetID returns the ID field value
func (o *NetworkProtectPrefixListResponseFields) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListResponseFields) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *NetworkProtectPrefixListResponseFields) SetID(v string) {
	o.ID = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NetworkProtectPrefixListResponseFields) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListResponseFields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NetworkProtectPrefixListResponseFields) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *NetworkProtectPrefixListResponseFields) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPrefixListResponseFields) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *NetworkProtectPrefixListResponseFields) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectPrefixListResponseFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.ID
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
func (o *NetworkProtectPrefixListResponseFields) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectPrefixListResponseFields := _NetworkProtectPrefixListResponseFields{}

	if err = json.Unmarshal(bytes, &varNetworkProtectPrefixListResponseFields); err == nil {
		*o = NetworkProtectPrefixListResponseFields(varNetworkProtectPrefixListResponseFields)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectPrefixListResponseFields is a helper abstraction for handling nullable networkprotectprefixlistresponsefields types.
type NullableNetworkProtectPrefixListResponseFields struct {
	value *NetworkProtectPrefixListResponseFields
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectPrefixListResponseFields) Get() *NetworkProtectPrefixListResponseFields {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectPrefixListResponseFields) Set(val *NetworkProtectPrefixListResponseFields) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectPrefixListResponseFields) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectPrefixListResponseFields) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectPrefixListResponseFields returns a pointer to a new instance of NullableNetworkProtectPrefixListResponseFields.
func NewNullableNetworkProtectPrefixListResponseFields(val *NetworkProtectPrefixListResponseFields) *NullableNetworkProtectPrefixListResponseFields {
	return &NullableNetworkProtectPrefixListResponseFields{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectPrefixListResponseFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectPrefixListResponseFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
