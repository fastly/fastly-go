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

// BotManagementWorkspace A container for bot management settings and policy.
type BotManagementWorkspace struct {
	// The display name of the workspace.
	Name *string `json:"name,omitempty"`
	// A description of the workspace.
	Description          *string                      `json:"description,omitempty"`
	ProtectionMode       *BotManagementProtectionMode `json:"protection_mode,omitempty"`
	AdditionalProperties map[string]any
}

type _BotManagementWorkspace BotManagementWorkspace

// NewBotManagementWorkspace instantiates a new BotManagementWorkspace object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementWorkspace() *BotManagementWorkspace {
	this := BotManagementWorkspace{}
	return &this
}

// NewBotManagementWorkspaceWithDefaults instantiates a new BotManagementWorkspace object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementWorkspaceWithDefaults() *BotManagementWorkspace {
	this := BotManagementWorkspace{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BotManagementWorkspace) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspace) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BotManagementWorkspace) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BotManagementWorkspace) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BotManagementWorkspace) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspace) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BotManagementWorkspace) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BotManagementWorkspace) SetDescription(v string) {
	o.Description = &v
}

// GetProtectionMode returns the ProtectionMode field value if set, zero value otherwise.
func (o *BotManagementWorkspace) GetProtectionMode() BotManagementProtectionMode {
	if o == nil || o.ProtectionMode == nil {
		var ret BotManagementProtectionMode
		return ret
	}
	return *o.ProtectionMode
}

// GetProtectionModeOk returns a tuple with the ProtectionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspace) GetProtectionModeOk() (*BotManagementProtectionMode, bool) {
	if o == nil || o.ProtectionMode == nil {
		return nil, false
	}
	return o.ProtectionMode, true
}

// HasProtectionMode returns a boolean if a field has been set.
func (o *BotManagementWorkspace) HasProtectionMode() bool {
	if o != nil && o.ProtectionMode != nil {
		return true
	}

	return false
}

// SetProtectionMode gets a reference to the given BotManagementProtectionMode and assigns it to the ProtectionMode field.
func (o *BotManagementWorkspace) SetProtectionMode(v BotManagementProtectionMode) {
	o.ProtectionMode = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementWorkspace) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ProtectionMode != nil {
		toSerialize["protection_mode"] = o.ProtectionMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BotManagementWorkspace) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementWorkspace := _BotManagementWorkspace{}

	if err = json.Unmarshal(bytes, &varBotManagementWorkspace); err == nil {
		*o = BotManagementWorkspace(varBotManagementWorkspace)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "protection_mode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementWorkspace is a helper abstraction for handling nullable botmanagementworkspace types.
type NullableBotManagementWorkspace struct {
	value *BotManagementWorkspace
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementWorkspace) Get() *BotManagementWorkspace {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementWorkspace) Set(val *BotManagementWorkspace) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementWorkspace) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementWorkspace) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementWorkspace returns a pointer to a new instance of NullableBotManagementWorkspace.
func NewNullableBotManagementWorkspace(val *BotManagementWorkspace) *NullableBotManagementWorkspace {
	return &NullableBotManagementWorkspace{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementWorkspace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementWorkspace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
