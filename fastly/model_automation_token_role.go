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

// AutomationTokenRole struct for AutomationTokenRole
type AutomationTokenRole struct {
	// The unique identifier of the role.
	Id *string `json:"id,omitempty"`
	// The name of the role.
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _AutomationTokenRole AutomationTokenRole

// NewAutomationTokenRole instantiates a new AutomationTokenRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutomationTokenRole() *AutomationTokenRole {
	this := AutomationTokenRole{}
	return &this
}

// NewAutomationTokenRoleWithDefaults instantiates a new AutomationTokenRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutomationTokenRoleWithDefaults() *AutomationTokenRole {
	this := AutomationTokenRole{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AutomationTokenRole) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenRole) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AutomationTokenRole) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AutomationTokenRole) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AutomationTokenRole) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutomationTokenRole) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AutomationTokenRole) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AutomationTokenRole) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AutomationTokenRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AutomationTokenRole) UnmarshalJSON(bytes []byte) (err error) {
	varAutomationTokenRole := _AutomationTokenRole{}

	if err = json.Unmarshal(bytes, &varAutomationTokenRole); err == nil {
		*o = AutomationTokenRole(varAutomationTokenRole)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAutomationTokenRole is a helper abstraction for handling nullable automationtokenrole types.
type NullableAutomationTokenRole struct {
	value *AutomationTokenRole
	isSet bool
}

// Get returns the value.
func (v NullableAutomationTokenRole) Get() *AutomationTokenRole {
	return v.value
}

// Set modifies the value.
func (v *NullableAutomationTokenRole) Set(val *AutomationTokenRole) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAutomationTokenRole) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAutomationTokenRole) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAutomationTokenRole returns a pointer to a new instance of NullableAutomationTokenRole.
func NewNullableAutomationTokenRole(val *AutomationTokenRole) *NullableAutomationTokenRole {
	return &NullableAutomationTokenRole{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAutomationTokenRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAutomationTokenRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
