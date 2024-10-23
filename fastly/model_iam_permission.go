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

// IamPermission struct for IamPermission
type IamPermission struct {
	// Alphanumeric string identifying the permission.
	ID *string `json:"id,omitempty"`
	// The type of the object.
	Object *string `json:"object,omitempty"`
	// Name of the permission.
	Name *string `json:"name,omitempty"`
	// The description of the permission.
	Description *string `json:"description,omitempty"`
	// The name of the resource the operation will be performed on.
	ResourceName *string `json:"resource_name,omitempty"`
	// The description of the resource.
	ResourceDescription *string `json:"resource_description,omitempty"`
	// Permissions are either \"service\" level or \"account\" level.
	Scope                *string `json:"scope,omitempty"`
	AdditionalProperties map[string]any
}

type _IamPermission IamPermission

// NewIamPermission instantiates a new IamPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamPermission() *IamPermission {
	this := IamPermission{}
	return &this
}

// NewIamPermissionWithDefaults instantiates a new IamPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamPermissionWithDefaults() *IamPermission {
	this := IamPermission{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *IamPermission) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermission) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *IamPermission) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *IamPermission) SetID(v string) {
	o.ID = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *IamPermission) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermission) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *IamPermission) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *IamPermission) SetObject(v string) {
	o.Object = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamPermission) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermission) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamPermission) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamPermission) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamPermission) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermission) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamPermission) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamPermission) SetDescription(v string) {
	o.Description = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *IamPermission) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermission) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *IamPermission) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *IamPermission) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourceDescription returns the ResourceDescription field value if set, zero value otherwise.
func (o *IamPermission) GetResourceDescription() string {
	if o == nil || o.ResourceDescription == nil {
		var ret string
		return ret
	}
	return *o.ResourceDescription
}

// GetResourceDescriptionOk returns a tuple with the ResourceDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermission) GetResourceDescriptionOk() (*string, bool) {
	if o == nil || o.ResourceDescription == nil {
		return nil, false
	}
	return o.ResourceDescription, true
}

// HasResourceDescription returns a boolean if a field has been set.
func (o *IamPermission) HasResourceDescription() bool {
	if o != nil && o.ResourceDescription != nil {
		return true
	}

	return false
}

// SetResourceDescription gets a reference to the given string and assigns it to the ResourceDescription field.
func (o *IamPermission) SetResourceDescription(v string) {
	o.ResourceDescription = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *IamPermission) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamPermission) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *IamPermission) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *IamPermission) SetScope(v string) {
	o.Scope = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o IamPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Object != nil {
		toSerialize["object"] = o.Object
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.ResourceDescription != nil {
		toSerialize["resource_description"] = o.ResourceDescription
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *IamPermission) UnmarshalJSON(bytes []byte) (err error) {
	varIamPermission := _IamPermission{}

	if err = json.Unmarshal(bytes, &varIamPermission); err == nil {
		*o = IamPermission(varIamPermission)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "resource_name")
		delete(additionalProperties, "resource_description")
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableIamPermission is a helper abstraction for handling nullable iampermission types.
type NullableIamPermission struct {
	value *IamPermission
	isSet bool
}

// Get returns the value.
func (v NullableIamPermission) Get() *IamPermission {
	return v.value
}

// Set modifies the value.
func (v *NullableIamPermission) Set(val *IamPermission) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIamPermission) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIamPermission) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIamPermission returns a pointer to a new instance of NullableIamPermission.
func NewNullableIamPermission(val *IamPermission) *NullableIamPermission {
	return &NullableIamPermission{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIamPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableIamPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
