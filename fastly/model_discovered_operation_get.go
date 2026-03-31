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

// DiscoveredOperationGet struct for DiscoveredOperationGet
type DiscoveredOperationGet struct {
	// The HTTP method for the operation.
	Method string `json:"method"`
	// The domain for the operation.
	Domain string `json:"domain"`
	// The path for the operation, which may include path parameters.
	Path string `json:"path"`
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

type _DiscoveredOperationGet DiscoveredOperationGet

// NewDiscoveredOperationGet instantiates a new DiscoveredOperationGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveredOperationGet(method string, domain string, path string, id string) *DiscoveredOperationGet {
	this := DiscoveredOperationGet{}
	this.Method = method
	this.Domain = domain
	this.Path = path
	this.Id = id
	return &this
}

// NewDiscoveredOperationGetWithDefaults instantiates a new DiscoveredOperationGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveredOperationGetWithDefaults() *DiscoveredOperationGet {
	this := DiscoveredOperationGet{}
	return &this
}

// GetMethod returns the Method field value
func (o *DiscoveredOperationGet) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationGet) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *DiscoveredOperationGet) SetMethod(v string) {
	o.Method = v
}

// GetDomain returns the Domain field value
func (o *DiscoveredOperationGet) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationGet) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DiscoveredOperationGet) SetDomain(v string) {
	o.Domain = v
}

// GetPath returns the Path field value
func (o *DiscoveredOperationGet) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationGet) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DiscoveredOperationGet) SetPath(v string) {
	o.Path = v
}

// GetId returns the Id field value
func (o *DiscoveredOperationGet) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationGet) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DiscoveredOperationGet) SetId(v string) {
	o.Id = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DiscoveredOperationGet) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationGet) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DiscoveredOperationGet) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DiscoveredOperationGet) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *DiscoveredOperationGet) GetLastSeenAt() time.Time {
	if o == nil || o.LastSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationGet) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || o.LastSeenAt == nil {
		return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *DiscoveredOperationGet) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt != nil {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *DiscoveredOperationGet) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

// GetRps returns the Rps field value if set, zero value otherwise.
func (o *DiscoveredOperationGet) GetRps() float32 {
	if o == nil || o.Rps == nil {
		var ret float32
		return ret
	}
	return *o.Rps
}

// GetRpsOk returns a tuple with the Rps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveredOperationGet) GetRpsOk() (*float32, bool) {
	if o == nil || o.Rps == nil {
		return nil, false
	}
	return o.Rps, true
}

// HasRps returns a boolean if a field has been set.
func (o *DiscoveredOperationGet) HasRps() bool {
	if o != nil && o.Rps != nil {
		return true
	}

	return false
}

// SetRps gets a reference to the given float32 and assigns it to the Rps field.
func (o *DiscoveredOperationGet) SetRps(v float32) {
	o.Rps = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DiscoveredOperationGet) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["method"] = o.Method
	}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["path"] = o.Path
	}
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
func (o *DiscoveredOperationGet) UnmarshalJSON(bytes []byte) (err error) {
	varDiscoveredOperationGet := _DiscoveredOperationGet{}

	if err = json.Unmarshal(bytes, &varDiscoveredOperationGet); err == nil {
		*o = DiscoveredOperationGet(varDiscoveredOperationGet)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "path")
		delete(additionalProperties, "id")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "last_seen_at")
		delete(additionalProperties, "rps")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDiscoveredOperationGet is a helper abstraction for handling nullable discoveredoperationget types.
type NullableDiscoveredOperationGet struct {
	value *DiscoveredOperationGet
	isSet bool
}

// Get returns the value.
func (v NullableDiscoveredOperationGet) Get() *DiscoveredOperationGet {
	return v.value
}

// Set modifies the value.
func (v *NullableDiscoveredOperationGet) Set(val *DiscoveredOperationGet) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDiscoveredOperationGet) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDiscoveredOperationGet) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDiscoveredOperationGet returns a pointer to a new instance of NullableDiscoveredOperationGet.
func NewNullableDiscoveredOperationGet(val *DiscoveredOperationGet) *NullableDiscoveredOperationGet {
	return &NullableDiscoveredOperationGet{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDiscoveredOperationGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDiscoveredOperationGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
