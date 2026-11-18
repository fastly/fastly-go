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

// BotManagementWorkspaceResponseExtra struct for BotManagementWorkspaceResponseExtra
type BotManagementWorkspaceResponseExtra struct {
	// A UUID used to uniquely identify the workspace.
	ID string `json:"id"`
	// The IDs of the services attached to this workspace.
	Services []string `json:"services,omitempty"`
	// The date and time the workspace was created, in RFC 3339 format.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the workspace was last updated, in RFC 3339 format.
	UpdatedAt            time.Time `json:"updated_at"`
	AdditionalProperties map[string]any
}

type _BotManagementWorkspaceResponseExtra BotManagementWorkspaceResponseExtra

// NewBotManagementWorkspaceResponseExtra instantiates a new BotManagementWorkspaceResponseExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBotManagementWorkspaceResponseExtra(id string, createdAt time.Time, updatedAt time.Time) *BotManagementWorkspaceResponseExtra {
	this := BotManagementWorkspaceResponseExtra{}
	this.ID = id
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	return &this
}

// NewBotManagementWorkspaceResponseExtraWithDefaults instantiates a new BotManagementWorkspaceResponseExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBotManagementWorkspaceResponseExtraWithDefaults() *BotManagementWorkspaceResponseExtra {
	this := BotManagementWorkspaceResponseExtra{}
	return &this
}

// GetID returns the ID field value
func (o *BotManagementWorkspaceResponseExtra) GetID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ID
}

// GetIDOk returns a tuple with the ID field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceResponseExtra) GetIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ID, true
}

// SetID sets field value
func (o *BotManagementWorkspaceResponseExtra) SetID(v string) {
	o.ID = v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *BotManagementWorkspaceResponseExtra) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceResponseExtra) GetServicesOk() ([]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *BotManagementWorkspaceResponseExtra) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *BotManagementWorkspaceResponseExtra) SetServices(v []string) {
	o.Services = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BotManagementWorkspaceResponseExtra) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceResponseExtra) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BotManagementWorkspaceResponseExtra) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *BotManagementWorkspaceResponseExtra) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *BotManagementWorkspaceResponseExtra) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *BotManagementWorkspaceResponseExtra) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BotManagementWorkspaceResponseExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["id"] = o.ID
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
func (o *BotManagementWorkspaceResponseExtra) UnmarshalJSON(bytes []byte) (err error) {
	varBotManagementWorkspaceResponseExtra := _BotManagementWorkspaceResponseExtra{}

	if err = json.Unmarshal(bytes, &varBotManagementWorkspaceResponseExtra); err == nil {
		*o = BotManagementWorkspaceResponseExtra(varBotManagementWorkspaceResponseExtra)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "services")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBotManagementWorkspaceResponseExtra is a helper abstraction for handling nullable botmanagementworkspaceresponseextra types.
type NullableBotManagementWorkspaceResponseExtra struct {
	value *BotManagementWorkspaceResponseExtra
	isSet bool
}

// Get returns the value.
func (v NullableBotManagementWorkspaceResponseExtra) Get() *BotManagementWorkspaceResponseExtra {
	return v.value
}

// Set modifies the value.
func (v *NullableBotManagementWorkspaceResponseExtra) Set(val *BotManagementWorkspaceResponseExtra) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBotManagementWorkspaceResponseExtra) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBotManagementWorkspaceResponseExtra) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBotManagementWorkspaceResponseExtra returns a pointer to a new instance of NullableBotManagementWorkspaceResponseExtra.
func NewNullableBotManagementWorkspaceResponseExtra(val *BotManagementWorkspaceResponseExtra) *NullableBotManagementWorkspaceResponseExtra {
	return &NullableBotManagementWorkspaceResponseExtra{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBotManagementWorkspaceResponseExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBotManagementWorkspaceResponseExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
