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

// NetworkProtectPolicyEndpoint One endpoint attached to a network policy.
type NetworkProtectPolicyEndpoint struct {
	// A unique identifier for a Network Protect resource.
	EndpointID string `json:"endpoint_id"`
	// Delivery preference within a policy. Lower values are preferred first. Two endpoints in the same policy cannot share a value.
	Preference int32 `json:"preference"`
	// An ISO 8601 timestamp in UTC.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// An ISO 8601 timestamp in UTC.
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]any
}

type _NetworkProtectPolicyEndpoint NetworkProtectPolicyEndpoint

// NewNetworkProtectPolicyEndpoint instantiates a new NetworkProtectPolicyEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectPolicyEndpoint(endpointID string, preference int32) *NetworkProtectPolicyEndpoint {
	this := NetworkProtectPolicyEndpoint{}
	this.EndpointID = endpointID
	this.Preference = preference
	return &this
}

// NewNetworkProtectPolicyEndpointWithDefaults instantiates a new NetworkProtectPolicyEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectPolicyEndpointWithDefaults() *NetworkProtectPolicyEndpoint {
	this := NetworkProtectPolicyEndpoint{}
	return &this
}

// GetEndpointID returns the EndpointID field value
func (o *NetworkProtectPolicyEndpoint) GetEndpointID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndpointID
}

// GetEndpointIDOk returns a tuple with the EndpointID field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPolicyEndpoint) GetEndpointIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndpointID, true
}

// SetEndpointID sets field value
func (o *NetworkProtectPolicyEndpoint) SetEndpointID(v string) {
	o.EndpointID = v
}

// GetPreference returns the Preference field value
func (o *NetworkProtectPolicyEndpoint) GetPreference() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Preference
}

// GetPreferenceOk returns a tuple with the Preference field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectPolicyEndpoint) GetPreferenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Preference, true
}

// SetPreference sets field value
func (o *NetworkProtectPolicyEndpoint) SetPreference(v int32) {
	o.Preference = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *NetworkProtectPolicyEndpoint) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectPolicyEndpoint) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *NetworkProtectPolicyEndpoint) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *NetworkProtectPolicyEndpoint) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *NetworkProtectPolicyEndpoint) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectPolicyEndpoint) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *NetworkProtectPolicyEndpoint) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *NetworkProtectPolicyEndpoint) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectPolicyEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["endpoint_id"] = o.EndpointID
	}
	if true {
		toSerialize["preference"] = o.Preference
	}
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
func (o *NetworkProtectPolicyEndpoint) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectPolicyEndpoint := _NetworkProtectPolicyEndpoint{}

	if err = json.Unmarshal(bytes, &varNetworkProtectPolicyEndpoint); err == nil {
		*o = NetworkProtectPolicyEndpoint(varNetworkProtectPolicyEndpoint)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "endpoint_id")
		delete(additionalProperties, "preference")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectPolicyEndpoint is a helper abstraction for handling nullable networkprotectpolicyendpoint types.
type NullableNetworkProtectPolicyEndpoint struct {
	value *NetworkProtectPolicyEndpoint
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectPolicyEndpoint) Get() *NetworkProtectPolicyEndpoint {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectPolicyEndpoint) Set(val *NetworkProtectPolicyEndpoint) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectPolicyEndpoint) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectPolicyEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectPolicyEndpoint returns a pointer to a new instance of NullableNetworkProtectPolicyEndpoint.
func NewNullableNetworkProtectPolicyEndpoint(val *NetworkProtectPolicyEndpoint) *NullableNetworkProtectPolicyEndpoint {
	return &NullableNetworkProtectPolicyEndpoint{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectPolicyEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectPolicyEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
