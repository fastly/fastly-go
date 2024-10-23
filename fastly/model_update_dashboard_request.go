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

// UpdateDashboardRequest struct for UpdateDashboardRequest
type UpdateDashboardRequest struct {
	// A human-readable name
	Name *string `json:"name,omitempty"`
	// A short description of the dashboard
	Description *string `json:"description,omitempty"`
	// A list of [dashboard items](#dashboard-item).
	Items                []DashboardItem `json:"items,omitempty"`
	AdditionalProperties map[string]any
}

type _UpdateDashboardRequest UpdateDashboardRequest

// NewUpdateDashboardRequest instantiates a new UpdateDashboardRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDashboardRequest() *UpdateDashboardRequest {
	this := UpdateDashboardRequest{}
	return &this
}

// NewUpdateDashboardRequestWithDefaults instantiates a new UpdateDashboardRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDashboardRequestWithDefaults() *UpdateDashboardRequest {
	this := UpdateDashboardRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateDashboardRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDashboardRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateDashboardRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateDashboardRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateDashboardRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDashboardRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateDashboardRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateDashboardRequest) SetDescription(v string) {
	o.Description = &v
}

// GetItems returns the Items field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateDashboardRequest) GetItems() []DashboardItem {
	if o == nil {
		var ret []DashboardItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateDashboardRequest) GetItemsOk() ([]DashboardItem, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UpdateDashboardRequest) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []DashboardItem and assigns it to the Items field.
func (o *UpdateDashboardRequest) SetItems(v []DashboardItem) {
	o.Items = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o UpdateDashboardRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *UpdateDashboardRequest) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateDashboardRequest := _UpdateDashboardRequest{}

	if err = json.Unmarshal(bytes, &varUpdateDashboardRequest); err == nil {
		*o = UpdateDashboardRequest(varUpdateDashboardRequest)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "items")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableUpdateDashboardRequest is a helper abstraction for handling nullable updatedashboardrequest types.
type NullableUpdateDashboardRequest struct {
	value *UpdateDashboardRequest
	isSet bool
}

// Get returns the value.
func (v NullableUpdateDashboardRequest) Get() *UpdateDashboardRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableUpdateDashboardRequest) Set(val *UpdateDashboardRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableUpdateDashboardRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableUpdateDashboardRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableUpdateDashboardRequest returns a pointer to a new instance of NullableUpdateDashboardRequest.
func NewNullableUpdateDashboardRequest(val *UpdateDashboardRequest) *NullableUpdateDashboardRequest {
	return &NullableUpdateDashboardRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableUpdateDashboardRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableUpdateDashboardRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
