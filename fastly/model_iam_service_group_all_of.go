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

// IamServiceGroupAllOf struct for IamServiceGroupAllOf
type IamServiceGroupAllOf struct {
	// Alphanumeric string identifying the service group.
	ID *string `json:"id,omitempty"`
	// The type of the object.
	Object *string `json:"object,omitempty"`
	// Name of the service group.
	Name *string `json:"name,omitempty"`
	// Description of the service group.
	Description *string `json:"description,omitempty"`
	// Number of services in the service group.
	ServicesCount        *int32 `json:"services_count,omitempty"`
	AdditionalProperties map[string]any
}

type _IamServiceGroupAllOf IamServiceGroupAllOf

// NewIamServiceGroupAllOf instantiates a new IamServiceGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamServiceGroupAllOf() *IamServiceGroupAllOf {
	this := IamServiceGroupAllOf{}
	return &this
}

// NewIamServiceGroupAllOfWithDefaults instantiates a new IamServiceGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamServiceGroupAllOfWithDefaults() *IamServiceGroupAllOf {
	this := IamServiceGroupAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *IamServiceGroupAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamServiceGroupAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *IamServiceGroupAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *IamServiceGroupAllOf) SetID(v string) {
	o.ID = &v
}

// GetObject returns the Object field value if set, zero value otherwise.
func (o *IamServiceGroupAllOf) GetObject() string {
	if o == nil || o.Object == nil {
		var ret string
		return ret
	}
	return *o.Object
}

// GetObjectOk returns a tuple with the Object field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamServiceGroupAllOf) GetObjectOk() (*string, bool) {
	if o == nil || o.Object == nil {
		return nil, false
	}
	return o.Object, true
}

// HasObject returns a boolean if a field has been set.
func (o *IamServiceGroupAllOf) HasObject() bool {
	if o != nil && o.Object != nil {
		return true
	}

	return false
}

// SetObject gets a reference to the given string and assigns it to the Object field.
func (o *IamServiceGroupAllOf) SetObject(v string) {
	o.Object = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamServiceGroupAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamServiceGroupAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamServiceGroupAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamServiceGroupAllOf) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IamServiceGroupAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamServiceGroupAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IamServiceGroupAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IamServiceGroupAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetServicesCount returns the ServicesCount field value if set, zero value otherwise.
func (o *IamServiceGroupAllOf) GetServicesCount() int32 {
	if o == nil || o.ServicesCount == nil {
		var ret int32
		return ret
	}
	return *o.ServicesCount
}

// GetServicesCountOk returns a tuple with the ServicesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamServiceGroupAllOf) GetServicesCountOk() (*int32, bool) {
	if o == nil || o.ServicesCount == nil {
		return nil, false
	}
	return o.ServicesCount, true
}

// HasServicesCount returns a boolean if a field has been set.
func (o *IamServiceGroupAllOf) HasServicesCount() bool {
	if o != nil && o.ServicesCount != nil {
		return true
	}

	return false
}

// SetServicesCount gets a reference to the given int32 and assigns it to the ServicesCount field.
func (o *IamServiceGroupAllOf) SetServicesCount(v int32) {
	o.ServicesCount = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o IamServiceGroupAllOf) MarshalJSON() ([]byte, error) {
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
	if o.ServicesCount != nil {
		toSerialize["services_count"] = o.ServicesCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *IamServiceGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamServiceGroupAllOf := _IamServiceGroupAllOf{}

	if err = json.Unmarshal(bytes, &varIamServiceGroupAllOf); err == nil {
		*o = IamServiceGroupAllOf(varIamServiceGroupAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "object")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "services_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableIamServiceGroupAllOf is a helper abstraction for handling nullable iamservicegroupallof types.
type NullableIamServiceGroupAllOf struct {
	value *IamServiceGroupAllOf
	isSet bool
}

// Get returns the value.
func (v NullableIamServiceGroupAllOf) Get() *IamServiceGroupAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableIamServiceGroupAllOf) Set(val *IamServiceGroupAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIamServiceGroupAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIamServiceGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIamServiceGroupAllOf returns a pointer to a new instance of NullableIamServiceGroupAllOf.
func NewNullableIamServiceGroupAllOf(val *IamServiceGroupAllOf) *NullableIamServiceGroupAllOf {
	return &NullableIamServiceGroupAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIamServiceGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableIamServiceGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
