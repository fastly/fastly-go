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

// ResourceCreateAllOf struct for ResourceCreateAllOf
type ResourceCreateAllOf struct {
	// The ID of the linked resource.
	ResourceID *string `json:"resource_id,omitempty"`
	AdditionalProperties map[string]any
}

type _ResourceCreateAllOf ResourceCreateAllOf

// NewResourceCreateAllOf instantiates a new ResourceCreateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceCreateAllOf() *ResourceCreateAllOf {
	this := ResourceCreateAllOf{}
	return &this
}

// NewResourceCreateAllOfWithDefaults instantiates a new ResourceCreateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceCreateAllOfWithDefaults() *ResourceCreateAllOf {
	this := ResourceCreateAllOf{}
	return &this
}

// GetResourceID returns the ResourceID field value if set, zero value otherwise.
func (o *ResourceCreateAllOf) GetResourceID() string {
	if o == nil || o.ResourceID == nil {
		var ret string
		return ret
	}
	return *o.ResourceID
}

// GetResourceIDOk returns a tuple with the ResourceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceCreateAllOf) GetResourceIDOk() (*string, bool) {
	if o == nil || o.ResourceID == nil {
		return nil, false
	}
	return o.ResourceID, true
}

// HasResourceID returns a boolean if a field has been set.
func (o *ResourceCreateAllOf) HasResourceID() bool {
	if o != nil && o.ResourceID != nil {
		return true
	}

	return false
}

// SetResourceID gets a reference to the given string and assigns it to the ResourceID field.
func (o *ResourceCreateAllOf) SetResourceID(v string) {
	o.ResourceID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ResourceCreateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *ResourceCreateAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varResourceCreateAllOf := _ResourceCreateAllOf{}

	if err = json.Unmarshal(bytes, &varResourceCreateAllOf); err == nil {
		*o = ResourceCreateAllOf(varResourceCreateAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "resource_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableResourceCreateAllOf is a helper abstraction for handling nullable resourcecreateallof types. 
type NullableResourceCreateAllOf struct {
	value *ResourceCreateAllOf
	isSet bool
}

// Get returns the value.
func (v NullableResourceCreateAllOf) Get() *ResourceCreateAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableResourceCreateAllOf) Set(val *ResourceCreateAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableResourceCreateAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableResourceCreateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableResourceCreateAllOf returns a pointer to a new instance of NullableResourceCreateAllOf.
func NewNullableResourceCreateAllOf(val *ResourceCreateAllOf) *NullableResourceCreateAllOf {
	return &NullableResourceCreateAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableResourceCreateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableResourceCreateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
