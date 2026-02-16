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

// OperationGet struct for OperationGet
type OperationGet struct {
	// The HTTP method for the operation.
	Method string `json:"method"`
	// The domain for the operation.
	Domain string `json:"domain"`
	// The path for the operation, which may include path parameters.
	Path string `json:"path"`
	// A description of what the operation does.
	Description *string `json:"description,omitempty"`
	// An array of operation tag IDs associated with this operation.
	TagIds []string `json:"tag_ids,omitempty"`
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

type _OperationGet OperationGet

// NewOperationGet instantiates a new OperationGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationGet(method string, domain string, path string, id string, updatedAt time.Time) *OperationGet {
	this := OperationGet{}
	this.Method = method
	this.Domain = domain
	this.Path = path
	this.Id = id
	this.UpdatedAt = updatedAt
	return &this
}

// NewOperationGetWithDefaults instantiates a new OperationGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationGetWithDefaults() *OperationGet {
	this := OperationGet{}
	return &this
}

// GetMethod returns the Method field value
func (o *OperationGet) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *OperationGet) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *OperationGet) SetMethod(v string) {
	o.Method = v
}

// GetDomain returns the Domain field value
func (o *OperationGet) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *OperationGet) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *OperationGet) SetDomain(v string) {
	o.Domain = v
}

// GetPath returns the Path field value
func (o *OperationGet) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *OperationGet) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *OperationGet) SetPath(v string) {
	o.Path = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OperationGet) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationGet) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OperationGet) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OperationGet) SetDescription(v string) {
	o.Description = &v
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *OperationGet) GetTagIds() []string {
	if o == nil || o.TagIds == nil {
		var ret []string
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationGet) GetTagIdsOk() ([]string, bool) {
	if o == nil || o.TagIds == nil {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *OperationGet) HasTagIds() bool {
	if o != nil && o.TagIds != nil {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []string and assigns it to the TagIds field.
func (o *OperationGet) SetTagIds(v []string) {
	o.TagIds = v
}

// GetId returns the Id field value
func (o *OperationGet) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OperationGet) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OperationGet) SetId(v string) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *OperationGet) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationGet) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *OperationGet) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *OperationGet) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *OperationGet) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *OperationGet) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *OperationGet) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetLastSeenAt returns the LastSeenAt field value if set, zero value otherwise.
func (o *OperationGet) GetLastSeenAt() time.Time {
	if o == nil || o.LastSeenAt == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSeenAt
}

// GetLastSeenAtOk returns a tuple with the LastSeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OperationGet) GetLastSeenAtOk() (*time.Time, bool) {
	if o == nil || o.LastSeenAt == nil {
		return nil, false
	}
	return o.LastSeenAt, true
}

// HasLastSeenAt returns a boolean if a field has been set.
func (o *OperationGet) HasLastSeenAt() bool {
	if o != nil && o.LastSeenAt != nil {
		return true
	}

	return false
}

// SetLastSeenAt gets a reference to the given time.Time and assigns it to the LastSeenAt field.
func (o *OperationGet) SetLastSeenAt(v time.Time) {
	o.LastSeenAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OperationGet) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.TagIds != nil {
		toSerialize["tag_ids"] = o.TagIds
	}
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
func (o *OperationGet) UnmarshalJSON(bytes []byte) (err error) {
	varOperationGet := _OperationGet{}

	if err = json.Unmarshal(bytes, &varOperationGet); err == nil {
		*o = OperationGet(varOperationGet)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "method")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "path")
		delete(additionalProperties, "description")
		delete(additionalProperties, "tag_ids")
		delete(additionalProperties, "id")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "last_seen_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOperationGet is a helper abstraction for handling nullable operationget types.
type NullableOperationGet struct {
	value *OperationGet
	isSet bool
}

// Get returns the value.
func (v NullableOperationGet) Get() *OperationGet {
	return v.value
}

// Set modifies the value.
func (v *NullableOperationGet) Set(val *OperationGet) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOperationGet) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOperationGet) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOperationGet returns a pointer to a new instance of NullableOperationGet.
func NewNullableOperationGet(val *OperationGet) *NullableOperationGet {
	return &NullableOperationGet{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOperationGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOperationGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
