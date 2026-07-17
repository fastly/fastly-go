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

// PageUpdate struct for PageUpdate
type PageUpdate struct {
	WebsiteId            *string        `json:"website_id,omitempty"`
	Name                 *string        `json:"name,omitempty"`
	Description          *string        `json:"description,omitempty"`
	Notifications        []Notification `json:"notifications,omitempty"`
	Paths                []string       `json:"paths,omitempty"`
	AdditionalProperties map[string]any
}

type _PageUpdate PageUpdate

// NewPageUpdate instantiates a new PageUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageUpdate() *PageUpdate {
	this := PageUpdate{}
	return &this
}

// NewPageUpdateWithDefaults instantiates a new PageUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageUpdateWithDefaults() *PageUpdate {
	this := PageUpdate{}
	return &this
}

// GetWebsiteId returns the WebsiteId field value if set, zero value otherwise.
func (o *PageUpdate) GetWebsiteId() string {
	if o == nil || o.WebsiteId == nil {
		var ret string
		return ret
	}
	return *o.WebsiteId
}

// GetWebsiteIdOk returns a tuple with the WebsiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageUpdate) GetWebsiteIdOk() (*string, bool) {
	if o == nil || o.WebsiteId == nil {
		return nil, false
	}
	return o.WebsiteId, true
}

// HasWebsiteId returns a boolean if a field has been set.
func (o *PageUpdate) HasWebsiteId() bool {
	if o != nil && o.WebsiteId != nil {
		return true
	}

	return false
}

// SetWebsiteId gets a reference to the given string and assigns it to the WebsiteId field.
func (o *PageUpdate) SetWebsiteId(v string) {
	o.WebsiteId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PageUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PageUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PageUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PageUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PageUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PageUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *PageUpdate) GetNotifications() []Notification {
	if o == nil || o.Notifications == nil {
		var ret []Notification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageUpdate) GetNotificationsOk() ([]Notification, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *PageUpdate) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []Notification and assigns it to the Notifications field.
func (o *PageUpdate) SetNotifications(v []Notification) {
	o.Notifications = v
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *PageUpdate) GetPaths() []string {
	if o == nil || o.Paths == nil {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageUpdate) GetPathsOk() ([]string, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *PageUpdate) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *PageUpdate) SetPaths(v []string) {
	o.Paths = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PageUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.WebsiteId != nil {
		toSerialize["website_id"] = o.WebsiteId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PageUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varPageUpdate := _PageUpdate{}

	if err = json.Unmarshal(bytes, &varPageUpdate); err == nil {
		*o = PageUpdate(varPageUpdate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "website_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "notifications")
		delete(additionalProperties, "paths")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePageUpdate is a helper abstraction for handling nullable pageupdate types.
type NullablePageUpdate struct {
	value *PageUpdate
	isSet bool
}

// Get returns the value.
func (v NullablePageUpdate) Get() *PageUpdate {
	return v.value
}

// Set modifies the value.
func (v *NullablePageUpdate) Set(val *PageUpdate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePageUpdate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePageUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePageUpdate returns a pointer to a new instance of NullablePageUpdate.
func NewNullablePageUpdate(val *PageUpdate) *NullablePageUpdate {
	return &NullablePageUpdate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePageUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePageUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
