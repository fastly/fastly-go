// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// ResourceCreate struct for ResourceCreate
type ResourceCreate struct {
	// The name of the resource.
	Name *string `json:"name,omitempty"`
	// The ID of the linked resource.
	ResourceID *string `json:"resource_id,omitempty"`
	AdditionalProperties map[string]any
}

type _ResourceCreate ResourceCreate

// NewResourceCreate instantiates a new ResourceCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceCreate() *ResourceCreate {
	this := ResourceCreate{}
	return &this
}

// NewResourceCreateWithDefaults instantiates a new ResourceCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceCreateWithDefaults() *ResourceCreate {
	this := ResourceCreate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceCreate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceCreate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceCreate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceCreate) SetName(v string) {
	o.Name = &v
}

// GetResourceID returns the ResourceID field value if set, zero value otherwise.
func (o *ResourceCreate) GetResourceID() string {
	if o == nil || o.ResourceID == nil {
		var ret string
		return ret
	}
	return *o.ResourceID
}

// GetResourceIDOk returns a tuple with the ResourceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceCreate) GetResourceIDOk() (*string, bool) {
	if o == nil || o.ResourceID == nil {
		return nil, false
	}
	return o.ResourceID, true
}

// HasResourceID returns a boolean if a field has been set.
func (o *ResourceCreate) HasResourceID() bool {
	if o != nil && o.ResourceID != nil {
		return true
	}

	return false
}

// SetResourceID gets a reference to the given string and assigns it to the ResourceID field.
func (o *ResourceCreate) SetResourceID(v string) {
	o.ResourceID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ResourceCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ResourceID != nil {
		toSerialize["resource_id"] = o.ResourceID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ResourceCreate) UnmarshalJSON(bytes []byte) (err error) {
	varResourceCreate := _ResourceCreate{}

	if err = json.Unmarshal(bytes, &varResourceCreate); err == nil {
		*o = ResourceCreate(varResourceCreate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "resource_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableResourceCreate is a helper abstraction for handling nullable resourcecreate types. 
type NullableResourceCreate struct {
	value *ResourceCreate
	isSet bool
}

// Get returns the value.
func (v NullableResourceCreate) Get() *ResourceCreate {
	return v.value
}

// Set modifies the value.
func (v *NullableResourceCreate) Set(val *ResourceCreate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableResourceCreate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableResourceCreate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableResourceCreate returns a pointer to a new instance of NullableResourceCreate.
func NewNullableResourceCreate(val *ResourceCreate) *NullableResourceCreate {
	return &NullableResourceCreate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableResourceCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableResourceCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
