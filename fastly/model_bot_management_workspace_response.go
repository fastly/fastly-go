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

// BotManagementWorkspaceResponse struct for BotManagementWorkspaceResponse
type BotManagementWorkspaceResponse struct {
	// The display name of the workspace.
	Name string `json:"name"`
	// A description of the workspace.
	Description    *string                     `json:"description,omitempty"`
	ProtectionMode BotManagementProtectionMode `json:"protection_mode"`
	// A UUID used to uniquely identify the workspace.
	Id string `json:"id"`
	// The IDs of the services attached to this workspace.
	Services []string `json:"services,omitempty"`
	// The date and time the workspace was created, in RFC 3339 format.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the workspace was last updated, in RFC 3339 format.
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]any
}

type _BotManagementWorkspaceResponse BotManagementWorkspaceResponse

// NewBotManagementWorkspaceResponse instantiates a new BotManagementWorkspaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementWorkspaceResponse(name string, protectionMode BotManagementProtectionMode, id string, createdAt time.Time, updatedAt time.Time) *BotManagementWorkspaceResponse {
	this := BotManagementWorkspaceResponse{}
	this.Name = name
	this.ProtectionMode = protectionMode
	this.Id = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewBotManagementWorkspaceResponseWithDefaults instantiates a new BotManagementWorkspaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementWorkspaceResponseWithDefaults() *BotManagementWorkspaceResponse {
	this := BotManagementWorkspaceResponse{}
	return &this
}

// GetName returns the Name field value
func (o *BotManagementWorkspaceResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BotManagementWorkspaceResponse) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BotManagementWorkspaceResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BotManagementWorkspaceResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BotManagementWorkspaceResponse) SetDescription(v string) {
	o.Description = &v
}

// GetProtectionMode returns the ProtectionMode field value
func (o *BotManagementWorkspaceResponse) GetProtectionMode() BotManagementProtectionMode {
	if o == nil {
		var ret BotManagementProtectionMode
		return ret
	}

	return o.ProtectionMode
}

// GetProtectionModeOk returns a tuple with the ProtectionMode field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceResponse) GetProtectionModeOk() (*BotManagementProtectionMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtectionMode, true
}

// SetProtectionMode sets field value
func (o *BotManagementWorkspaceResponse) SetProtectionMode(v BotManagementProtectionMode) {
	o.ProtectionMode = v
}

// GetId returns the Id field value
func (o *BotManagementWorkspaceResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BotManagementWorkspaceResponse) SetId(v string) {
	o.Id = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *BotManagementWorkspaceResponse) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceResponse) GetServicesOk() ([]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *BotManagementWorkspaceResponse) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *BotManagementWorkspaceResponse) SetServices(v []string) {
	o.Services = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BotManagementWorkspaceResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BotManagementWorkspaceResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *BotManagementWorkspaceResponse) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *BotManagementWorkspaceResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementWorkspaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["protection_mode"] = o.ProtectionMode
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
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
func (o *BotManagementWorkspaceResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementWorkspaceResponse := _BotManagementWorkspaceResponse{}

	if err = json.Unmarshal(bytes, &varBotManagementWorkspaceResponse); err == nil {
		*o = BotManagementWorkspaceResponse(varBotManagementWorkspaceResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "protection_mode")
		delete(additionalProperties, "id")
		delete(additionalProperties, "services")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementWorkspaceResponse is a helper abstraction for handling nullable botmanagementworkspaceresponse types.
type NullableBotManagementWorkspaceResponse struct {
	value *BotManagementWorkspaceResponse
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementWorkspaceResponse) Get() *BotManagementWorkspaceResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementWorkspaceResponse) Set(val *BotManagementWorkspaceResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementWorkspaceResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementWorkspaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementWorkspaceResponse returns a pointer to a new instance of NullableBotManagementWorkspaceResponse.
func NewNullableBotManagementWorkspaceResponse(val *BotManagementWorkspaceResponse) *NullableBotManagementWorkspaceResponse {
	return &NullableBotManagementWorkspaceResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementWorkspaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementWorkspaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
