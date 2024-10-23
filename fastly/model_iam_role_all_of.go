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

// IamRoleAllOf struct for IamRoleAllOf
type IamRoleAllOf struct {
	// Alphanumeric string identifying the role.
	ID *string `json:"id,omitempty"`
	// The type of the object.
	Object *string `json:"object,omitempty"`
	// Name of the role.
	Name *string `json:"name,omitempty"`
	// Description of the role.
	Description *string `json:"description,omitempty"`
	// This attribute is set to `true` if the role is managed by the customer. It is set to `false` if the role was created by Fastly.
	Custom *bool `json:"custom,omitempty"`
	// Number of permissions assigned to the role.
	PermissionsCount     *int32 `json:"permissions_count,omitempty"`
	AdditionalProperties map[string]any
}

type _IamRoleAllOf IamRoleAllOf

// NewIamRoleAllOf instantiates a new IamRoleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamRoleAllOf() *IamRoleAllOf {
	this := IamRoleAllOf{}
	return &this
}

// NewIamRoleAllOfWithDefaults instantiates a new IamRoleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamRoleAllOfWithDefaults() *IamRoleAllOf {
	this := IamRoleAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *IamRoleAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *IamRoleAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *IamRoleAllOf) SetID(v string) {
	o.ID = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *IamRoleAllOf) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleAllOf) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *IamRoleAllOf) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *IamRoleAllOf) SetObject(v string) {
	o.Object = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamRoleAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamRoleAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamRoleAllOf) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamRoleAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamRoleAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamRoleAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *IamRoleAllOf) GetCustom() bool {
	if o == nil || o.Custom == nil {
		var ret bool
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleAllOf) GetCustomOk() (*bool, bool) {
	if o == nil || o.Custom == nil {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *IamRoleAllOf) HasCustom() bool {
	if o != nil && o.Custom != nil {
		return true
	}

	return false
}

// SetCustom gets a reference to the given bool and assigns it to the Custom field.
func (o *IamRoleAllOf) SetCustom(v bool) {
	o.Custom = &v
}

// GetPermissionsCount returns the PermissionsCount field value if set, zero value otherwise.
func (o *IamRoleAllOf) GetPermissionsCount() int32 {
	if o == nil || o.PermissionsCount == nil {
		var ret int32
		return ret
	}
	return *o.PermissionsCount
}

// GetPermissionsCountOk returns a tuple with the PermissionsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamRoleAllOf) GetPermissionsCountOk() (*int32, bool) {
	if o == nil || o.PermissionsCount == nil {
		return nil, false
	}
	return o.PermissionsCount, true
}

// HasPermissionsCount returns a boolean if a field has been set.
func (o *IamRoleAllOf) HasPermissionsCount() bool {
	if o != nil && o.PermissionsCount != nil {
		return true
	}

	return false
}

// SetPermissionsCount gets a reference to the given int32 and assigns it to the PermissionsCount field.
func (o *IamRoleAllOf) SetPermissionsCount(v int32) {
	o.PermissionsCount = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o IamRoleAllOf) MarshalJSON() ([]byte, error) {
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
	if o.Custom != nil {
		toSerialize["custom"] = o.Custom
	}
	if o.PermissionsCount != nil {
		toSerialize["permissions_count"] = o.PermissionsCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *IamRoleAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamRoleAllOf := _IamRoleAllOf{}

	if err = json.Unmarshal(bytes, &varIamRoleAllOf); err == nil {
		*o = IamRoleAllOf(varIamRoleAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "custom")
		delete(additionalProperties, "permissions_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableIamRoleAllOf is a helper abstraction for handling nullable iamroleallof types.
type NullableIamRoleAllOf struct {
	value *IamRoleAllOf
	isSet bool
}

// Get returns the value.
func (v NullableIamRoleAllOf) Get() *IamRoleAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableIamRoleAllOf) Set(val *IamRoleAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIamRoleAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIamRoleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIamRoleAllOf returns a pointer to a new instance of NullableIamRoleAllOf.
func NewNullableIamRoleAllOf(val *IamRoleAllOf) *NullableIamRoleAllOf {
	return &NullableIamRoleAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIamRoleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableIamRoleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
