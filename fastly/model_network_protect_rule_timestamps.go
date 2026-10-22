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

// NetworkProtectRuleTimestamps Read-only timestamps for an ordered rule. Moving a rule or changing a logical field or comment creates a new timestamp; appending a rule does not change timestamps for rules that remain in place.
type NetworkProtectRuleTimestamps struct {
	// An ISO 8601 timestamp in UTC.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// An ISO 8601 timestamp in UTC.
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleTimestamps NetworkProtectRuleTimestamps

// NewNetworkProtectRuleTimestamps instantiates a new NetworkProtectRuleTimestamps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleTimestamps() *NetworkProtectRuleTimestamps {
	this := NetworkProtectRuleTimestamps{}
	return &this
}

// NewNetworkProtectRuleTimestampsWithDefaults instantiates a new NetworkProtectRuleTimestamps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleTimestampsWithDefaults() *NetworkProtectRuleTimestamps {
	this := NetworkProtectRuleTimestamps{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NetworkProtectRuleTimestamps) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleTimestamps) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NetworkProtectRuleTimestamps) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NetworkProtectRuleTimestamps) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NetworkProtectRuleTimestamps) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleTimestamps) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NetworkProtectRuleTimestamps) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NetworkProtectRuleTimestamps) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleTimestamps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRuleTimestamps) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleTimestamps := _NetworkProtectRuleTimestamps{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleTimestamps); err == nil {
		*o = NetworkProtectRuleTimestamps(varNetworkProtectRuleTimestamps)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleTimestamps is a helper abstraction for handling nullable networkprotectruletimestamps types.
type NullableNetworkProtectRuleTimestamps struct {
	value *NetworkProtectRuleTimestamps
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleTimestamps) Get() *NetworkProtectRuleTimestamps {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleTimestamps) Set(val *NetworkProtectRuleTimestamps) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleTimestamps) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleTimestamps) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleTimestamps returns a pointer to a new instance of NullableNetworkProtectRuleTimestamps.
func NewNullableNetworkProtectRuleTimestamps(val *NetworkProtectRuleTimestamps) *NullableNetworkProtectRuleTimestamps {
	return &NullableNetworkProtectRuleTimestamps{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleTimestamps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleTimestamps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
