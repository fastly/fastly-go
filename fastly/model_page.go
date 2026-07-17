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

// Page struct for Page
type Page struct {
	// Unique page identifier
	Id *string `json:"id,omitempty"`
	// Parent website ID
	WebsiteId *string `json:"website_id,omitempty"`
	// Page name
	Name *string `json:"name,omitempty"`
	// Page description
	Description *string `json:"description,omitempty"`
	// Notification configurations for this page
	Notifications []Notification `json:"notifications,omitempty"`
	// URL paths to monitor
	Paths                []string   `json:"paths,omitempty"`
	CreatedAt            *time.Time `json:"created_at,omitempty"`
	UpdatedAt            *time.Time `json:"updated_at,omitempty"`
	AdditionalProperties map[string]any
}

type _Page Page

// NewPage instantiates a new Page object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPage() *Page {
	this := Page{}
	return &this
}

// NewPageWithDefaults instantiates a new Page object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageWithDefaults() *Page {
	this := Page{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Page) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Page) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Page) SetId(v string) {
	o.Id = &v
}

// GetWebsiteId returns the WebsiteId field value if set, zero value otherwise.
func (o *Page) GetWebsiteId() string {
	if o == nil || o.WebsiteId == nil {
		var ret string
		return ret
	}
	return *o.WebsiteId
}

// GetWebsiteIdOk returns a tuple with the WebsiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetWebsiteIdOk() (*string, bool) {
	if o == nil || o.WebsiteId == nil {
		return nil, false
	}
	return o.WebsiteId, true
}

// HasWebsiteId returns a boolean if a field has been set.
func (o *Page) HasWebsiteId() bool {
	if o != nil && o.WebsiteId != nil {
		return true
	}

	return false
}

// SetWebsiteId gets a reference to the given string and assigns it to the WebsiteId field.
func (o *Page) SetWebsiteId(v string) {
	o.WebsiteId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Page) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Page) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Page) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Page) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Page) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Page) SetDescription(v string) {
	o.Description = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *Page) GetNotifications() []Notification {
	if o == nil || o.Notifications == nil {
		var ret []Notification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetNotificationsOk() ([]Notification, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *Page) HasNotifications() bool {
	if o != nil && o.Notifications != nil {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []Notification and assigns it to the Notifications field.
func (o *Page) SetNotifications(v []Notification) {
	o.Notifications = v
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *Page) GetPaths() []string {
	if o == nil || o.Paths == nil {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetPathsOk() ([]string, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *Page) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *Page) SetPaths(v []string) {
	o.Paths = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Page) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Page) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Page) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Page) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Page) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Page) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Page) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Page) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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
func (o *Page) UnmarshalJSON(bytes []byte) (err error) {
	varPage := _Page{}

	if err = json.Unmarshal(bytes, &varPage); err == nil {
		*o = Page(varPage)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "website_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "notifications")
		delete(additionalProperties, "paths")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "updated_at")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePage is a helper abstraction for handling nullable page types.
type NullablePage struct {
	value *Page
	isSet bool
}

// Get returns the value.
func (v NullablePage) Get() *Page {
	return v.value
}

// Set modifies the value.
func (v *NullablePage) Set(val *Page) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePage) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePage) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePage returns a pointer to a new instance of NullablePage.
func NewNullablePage(val *Page) *NullablePage {
	return &NullablePage{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
