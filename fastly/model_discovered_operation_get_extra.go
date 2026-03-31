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

// DiscoveredOperationGetExtra struct for DiscoveredOperationGetExtra
type DiscoveredOperationGetExtra struct {
	// The unique identifier of the discovered operation.
	Id string `json:"id"`
	// The timestamp when the operation was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The timestamp when the operation was last seen in traffic.
	LastSeenAt *time.Time `json:"last_seen_at,omitempty"`
	// Requests per second observed for this operation.
	Rps                  *float32 `json:"rps,omitempty"`
	AdditionalProperties map[string]any
}

type _DiscoveredOperationGetExtra DiscoveredOperationGetExtra

// NewDiscoveredOperationGetExtra instantiates a new DiscoveredOperationGetExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveredOperationGetExtra(id string) *DiscoveredOperationGetExtra {
	this := DiscoveredOperationGetExtra{}
	this.Id = id
	return &this
}

// NewDiscoveredOperationGetExtraWithDefaults instantiates a new DiscoveredOperationGetExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveredOperationGetExtraWithDefaults() *DiscoveredOperationGetExtra {
	this := DiscoveredOperationGetExtra{}
	return &this
}

// GetId returns the Id field value
func (o *DiscoveredOperationGetExtra) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationGetExtra) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiscoveredOperationGetExtra) SetId(v string) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DiscoveredOperationGetExtra) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationGetExtra) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DiscoveredOperationGetExtra) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DiscoveredOperationGetExtra) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *DiscoveredOperationGetExtra) GetLastSeenAt() time.Time {
	if o == nil || o.LastSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationGetExtra) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || o.LastSeenAt == nil {
		return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *DiscoveredOperationGetExtra) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt != nil {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *DiscoveredOperationGetExtra) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

// GetRps returns the Rps field value if set, zero value otherwise.
func (o *DiscoveredOperationGetExtra) GetRps() float32 {
	if o == nil || o.Rps == nil {
		var ret float32
		return ret
	}
	return *o.Rps
}

// GetRpsOk returns a tuple with the Rps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationGetExtra) GetRpsOk() (*float32, bool) {
	if o == nil || o.Rps == nil {
		return nil, false
	}
	return o.Rps, true
}

// HasRps returns a boolean if a field has been set.
func (o *DiscoveredOperationGetExtra) HasRps() bool {
	if o != nil && o.Rps != nil {
		return true
	}

	return false
}

// SetRps gets a reference to the given float32 and assigns it to the Rps field.
func (o *DiscoveredOperationGetExtra) SetRps(v float32) {
	o.Rps = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DiscoveredOperationGetExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.LastSeenAt != nil {
		toSerialize["last_seen_at"] = o.LastSeenAt
	}
	if o.Rps != nil {
		toSerialize["rps"] = o.Rps
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DiscoveredOperationGetExtra) UnmarshalJSON(bytes []byte) (err error) {
	varDiscoveredOperationGetExtra := _DiscoveredOperationGetExtra{}

	if err = json.Unmarshal(bytes, &varDiscoveredOperationGetExtra); err == nil {
		*o = DiscoveredOperationGetExtra(varDiscoveredOperationGetExtra)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "last_seen_at")
		delete(additionalProperties, "rps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDiscoveredOperationGetExtra is a helper abstraction for handling nullable discoveredoperationgetextra types.
type NullableDiscoveredOperationGetExtra struct {
	value *DiscoveredOperationGetExtra
	isSet bool
}

// Get returns the value.
func (v NullableDiscoveredOperationGetExtra) Get() *DiscoveredOperationGetExtra {
	return v.value
}

// Set modifies the value.
func (v *NullableDiscoveredOperationGetExtra) Set(val *DiscoveredOperationGetExtra) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDiscoveredOperationGetExtra) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDiscoveredOperationGetExtra) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDiscoveredOperationGetExtra returns a pointer to a new instance of NullableDiscoveredOperationGetExtra.
func NewNullableDiscoveredOperationGetExtra(val *DiscoveredOperationGetExtra) *NullableDiscoveredOperationGetExtra {
	return &NullableDiscoveredOperationGetExtra{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDiscoveredOperationGetExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDiscoveredOperationGetExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
