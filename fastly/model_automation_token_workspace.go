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

// AutomationTokenWorkspace struct for AutomationTokenWorkspace
type AutomationTokenWorkspace struct {
	// Unique identifier for the workspace.
	Id string `json:"id"`
	// Human-readable workspace name.
	Name                 string `json:"name"`
	AdditionalProperties map[string]any
}

type _AutomationTokenWorkspace AutomationTokenWorkspace

// NewAutomationTokenWorkspace instantiates a new AutomationTokenWorkspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenWorkspace(id string, name string) *AutomationTokenWorkspace {
	this := AutomationTokenWorkspace{}
	this.Id = id
	this.Name = name
	return &this
}

// NewAutomationTokenWorkspaceWithDefaults instantiates a new AutomationTokenWorkspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenWorkspaceWithDefaults() *AutomationTokenWorkspace {
	this := AutomationTokenWorkspace{}
	return &this
}

// GetId returns the Id field value
func (o *AutomationTokenWorkspace) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AutomationTokenWorkspace) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AutomationTokenWorkspace) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AutomationTokenWorkspace) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AutomationTokenWorkspace) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AutomationTokenWorkspace) SetName(v string) {
	o.Name = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenWorkspace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AutomationTokenWorkspace) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenWorkspace := _AutomationTokenWorkspace{}

	if err = json.Unmarshal(bytes, &varAutomationTokenWorkspace); err == nil {
		*o = AutomationTokenWorkspace(varAutomationTokenWorkspace)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenWorkspace is a helper abstraction for handling nullable automationtokenworkspace types.
type NullableAutomationTokenWorkspace struct {
	value *AutomationTokenWorkspace
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenWorkspace) Get() *AutomationTokenWorkspace {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenWorkspace) Set(val *AutomationTokenWorkspace) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenWorkspace) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenWorkspace returns a pointer to a new instance of NullableAutomationTokenWorkspace.
func NewNullableAutomationTokenWorkspace(val *AutomationTokenWorkspace) *NullableAutomationTokenWorkspace {
	return &NullableAutomationTokenWorkspace{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAutomationTokenWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
