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

// Resource struct for Resource
type Resource struct {
	// The ID of the underlying linked resource.
	ResourceID *string `json:"resource_id,omitempty"`
	// The name of the resource link.
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _Resource Resource

// NewResource instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource() *Resource {
	this := Resource{}
	return &this
}

// NewResourceWithDefaults instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	this := Resource{}
	return &this
}

// GetResourceID returns the ResourceID field value if set, zero value otherwise.
func (o *Resource) GetResourceID() string {
	if o == nil || o.ResourceID == nil {
		var ret string
		return ret
	}
	return *o.ResourceID
}

// GetResourceIDOk returns a tuple with the ResourceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetResourceIDOk() (*string, bool) {
	if o == nil || o.ResourceID == nil {
		return nil, false
	}
	return o.ResourceID, true
}

// HasResourceID returns a boolean if a field has been set.
func (o *Resource) HasResourceID() bool {
	if o != nil && o.ResourceID != nil {
		return true
	}

	return false
}

// SetResourceID gets a reference to the given string and assigns it to the ResourceID field.
func (o *Resource) SetResourceID(v string) {
	o.ResourceID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Resource) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Resource) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Resource) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ResourceID != nil {
		toSerialize["resource_id"] = o.ResourceID
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
func (o *Resource) UnmarshalJSON(bytes []byte) (err error) {
	varResource := _Resource{}

	if err = json.Unmarshal(bytes, &varResource); err == nil {
		*o = Resource(varResource)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "resource_id")
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableResource is a helper abstraction for handling nullable resource types.
type NullableResource struct {
	value *Resource
	isSet bool
}

// Get returns the value.
func (v NullableResource) Get() *Resource {
	return v.value
}

// Set modifies the value.
func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableResource) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableResource returns a pointer to a new instance of NullableResource.
func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
