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

// NetworkProtectRuleSetResponseFields Read-only fields and required fields in a rule-set response.
type NetworkProtectRuleSetResponseFields struct {
	// A read-only unique identifier for a Network Protect resource.
	ID string `json:"id"`
	// An ISO 8601 timestamp in UTC.
	CreatedAt time.Time `json:"created_at"`
	// An ISO 8601 timestamp in UTC.
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleSetResponseFields NetworkProtectRuleSetResponseFields

// NewNetworkProtectRuleSetResponseFields instantiates a new NetworkProtectRuleSetResponseFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleSetResponseFields(id string, createdAt time.Time, updatedAt time.Time) *NetworkProtectRuleSetResponseFields {
	this := NetworkProtectRuleSetResponseFields{}
	this.ID = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewNetworkProtectRuleSetResponseFieldsWithDefaults instantiates a new NetworkProtectRuleSetResponseFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleSetResponseFieldsWithDefaults() *NetworkProtectRuleSetResponseFields {
	this := NetworkProtectRuleSetResponseFields{}
	return &this
}

// GetID returns the ID field value
func (o *NetworkProtectRuleSetResponseFields) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSetResponseFields) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *NetworkProtectRuleSetResponseFields) SetID(v string) {
	o.ID = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *NetworkProtectRuleSetResponseFields) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSetResponseFields) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *NetworkProtectRuleSetResponseFields) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *NetworkProtectRuleSetResponseFields) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleSetResponseFields) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *NetworkProtectRuleSetResponseFields) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleSetResponseFields) MarshalJSON() ([]byte, error) {
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
func (o *NetworkProtectRuleSetResponseFields) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleSetResponseFields := _NetworkProtectRuleSetResponseFields{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleSetResponseFields); err == nil {
		*o = NetworkProtectRuleSetResponseFields(varNetworkProtectRuleSetResponseFields)
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

// NullableNetworkProtectRuleSetResponseFields is a helper abstraction for handling nullable networkprotectrulesetresponsefields types.
type NullableNetworkProtectRuleSetResponseFields struct {
	value *NetworkProtectRuleSetResponseFields
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleSetResponseFields) Get() *NetworkProtectRuleSetResponseFields {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleSetResponseFields) Set(val *NetworkProtectRuleSetResponseFields) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleSetResponseFields) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleSetResponseFields) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleSetResponseFields returns a pointer to a new instance of NullableNetworkProtectRuleSetResponseFields.
func NewNullableNetworkProtectRuleSetResponseFields(val *NetworkProtectRuleSetResponseFields) *NullableNetworkProtectRuleSetResponseFields {
	return &NullableNetworkProtectRuleSetResponseFields{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleSetResponseFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleSetResponseFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
