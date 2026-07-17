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

// Policy struct for Policy
type Policy struct {
	Id                   *string     `json:"id,omitempty"`
	PageId               *string     `json:"page_id,omitempty"`
	Name                 *string     `json:"name,omitempty"`
	Description          *string     `json:"description,omitempty"`
	Mode                 *string     `json:"mode,omitempty"`
	Directives           []Directive `json:"directives,omitempty"`
	CreatedAt            *time.Time  `json:"created_at,omitempty"`
	UpdatedAt            *time.Time  `json:"updated_at,omitempty"`
	AdditionalProperties map[string]any
}

type _Policy Policy

// NewPolicy instantiates a new Policy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicy() *Policy {
	this := Policy{}
	return &this
}

// NewPolicyWithDefaults instantiates a new Policy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyWithDefaults() *Policy {
	this := Policy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Policy) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Policy) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Policy) SetId(v string) {
	o.Id = &v
}

// GetPageId returns the PageId field value if set, zero value otherwise.
func (o *Policy) GetPageId() string {
	if o == nil || o.PageId == nil {
		var ret string
		return ret
	}
	return *o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetPageIdOk() (*string, bool) {
	if o == nil || o.PageId == nil {
		return nil, false
	}
	return o.PageId, true
}

// HasPageId returns a boolean if a field has been set.
func (o *Policy) HasPageId() bool {
	if o != nil && o.PageId != nil {
		return true
	}

	return false
}

// SetPageId gets a reference to the given string and assigns it to the PageId field.
func (o *Policy) SetPageId(v string) {
	o.PageId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Policy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Policy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Policy) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Policy) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Policy) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Policy) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *Policy) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *Policy) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *Policy) SetMode(v string) {
	o.Mode = &v
}

// GetDirectives returns the Directives field value if set, zero value otherwise.
func (o *Policy) GetDirectives() []Directive {
	if o == nil || o.Directives == nil {
		var ret []Directive
		return ret
	}
	return o.Directives
}

// GetDirectivesOk returns a tuple with the Directives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetDirectivesOk() ([]Directive, bool) {
	if o == nil || o.Directives == nil {
		return nil, false
	}
	return o.Directives, true
}

// HasDirectives returns a boolean if a field has been set.
func (o *Policy) HasDirectives() bool {
	if o != nil && o.Directives != nil {
		return true
	}

	return false
}

// SetDirectives gets a reference to the given []Directive and assigns it to the Directives field.
func (o *Policy) SetDirectives(v []Directive) {
	o.Directives = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Policy) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Policy) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Policy) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Policy) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Policy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Policy) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Policy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Policy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PageId != nil {
		toSerialize["page_id"] = o.PageId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Directives != nil {
		toSerialize["directives"] = o.Directives
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
func (o *Policy) UnmarshalJSON(bytes []byte) (err error) {
	varPolicy := _Policy{}

	if err = json.Unmarshal(bytes, &varPolicy); err == nil {
		*o = Policy(varPolicy)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "page_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "directives")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePolicy is a helper abstraction for handling nullable policy types.
type NullablePolicy struct {
	value *Policy
	isSet bool
}

// Get returns the value.
func (v NullablePolicy) Get() *Policy {
	return v.value
}

// Set modifies the value.
func (v *NullablePolicy) Set(val *Policy) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePolicy) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePolicy returns a pointer to a new instance of NullablePolicy.
func NewNullablePolicy(val *Policy) *NullablePolicy {
	return &NullablePolicy{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
