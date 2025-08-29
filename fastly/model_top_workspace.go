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
)

// TopWorkspace This object, found within the `top_workspaces` array, contains the workspace information and count for the requested signal.
type TopWorkspace struct {
	// ID of the workspace.
	ID *string `json:"id,omitempty"`
	// Name of the workspace.
	Name *string `json:"name,omitempty"`
	// Count of attacks on this workspace for the specific attack type.
	Count                *int32 `json:"count,omitempty"`
	AdditionalProperties map[string]any
}

type _TopWorkspace TopWorkspace

// NewTopWorkspace instantiates a new TopWorkspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopWorkspace() *TopWorkspace {
	this := TopWorkspace{}
	return &this
}

// NewTopWorkspaceWithDefaults instantiates a new TopWorkspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopWorkspaceWithDefaults() *TopWorkspace {
	this := TopWorkspace{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *TopWorkspace) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopWorkspace) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *TopWorkspace) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *TopWorkspace) SetID(v string) {
	o.ID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TopWorkspace) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopWorkspace) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TopWorkspace) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TopWorkspace) SetName(v string) {
	o.Name = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TopWorkspace) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TopWorkspace) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TopWorkspace) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *TopWorkspace) SetCount(v int32) {
	o.Count = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TopWorkspace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *TopWorkspace) UnmarshalJSON(bytes []byte) (err error) {
	varTopWorkspace := _TopWorkspace{}

	if err = json.Unmarshal(bytes, &varTopWorkspace); err == nil {
		*o = TopWorkspace(varTopWorkspace)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTopWorkspace is a helper abstraction for handling nullable topworkspace types.
type NullableTopWorkspace struct {
	value *TopWorkspace
	isSet bool
}

// Get returns the value.
func (v NullableTopWorkspace) Get() *TopWorkspace {
	return v.value
}

// Set modifies the value.
func (v *NullableTopWorkspace) Set(val *TopWorkspace) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTopWorkspace) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTopWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTopWorkspace returns a pointer to a new instance of NullableTopWorkspace.
func NewNullableTopWorkspace(val *TopWorkspace) *NullableTopWorkspace {
	return &NullableTopWorkspace{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTopWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableTopWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
