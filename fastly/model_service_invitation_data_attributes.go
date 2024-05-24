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

// ServiceInvitationDataAttributes struct for ServiceInvitationDataAttributes
type ServiceInvitationDataAttributes struct {
	// The permission the accepting user will have in relation to the service.
	Permission *string `json:"permission,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceInvitationDataAttributes ServiceInvitationDataAttributes

// NewServiceInvitationDataAttributes instantiates a new ServiceInvitationDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceInvitationDataAttributes() *ServiceInvitationDataAttributes {
	this := ServiceInvitationDataAttributes{}
	var permission string = "full"
	this.Permission = &permission
	return &this
}

// NewServiceInvitationDataAttributesWithDefaults instantiates a new ServiceInvitationDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceInvitationDataAttributesWithDefaults() *ServiceInvitationDataAttributes {
	this := ServiceInvitationDataAttributes{}
	var permission string = "full"
	this.Permission = &permission
	return &this
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *ServiceInvitationDataAttributes) GetPermission() string {
	if o == nil || o.Permission == nil {
		var ret string
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInvitationDataAttributes) GetPermissionOk() (*string, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *ServiceInvitationDataAttributes) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given string and assigns it to the Permission field.
func (o *ServiceInvitationDataAttributes) SetPermission(v string) {
	o.Permission = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceInvitationDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Permission != nil {
		toSerialize["permission"] = o.Permission
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ServiceInvitationDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varServiceInvitationDataAttributes := _ServiceInvitationDataAttributes{}

	if err = json.Unmarshal(bytes, &varServiceInvitationDataAttributes); err == nil {
		*o = ServiceInvitationDataAttributes(varServiceInvitationDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "permission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceInvitationDataAttributes is a helper abstraction for handling nullable serviceinvitationdataattributes types. 
type NullableServiceInvitationDataAttributes struct {
	value *ServiceInvitationDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableServiceInvitationDataAttributes) Get() *ServiceInvitationDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceInvitationDataAttributes) Set(val *ServiceInvitationDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceInvitationDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceInvitationDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceInvitationDataAttributes returns a pointer to a new instance of NullableServiceInvitationDataAttributes.
func NewNullableServiceInvitationDataAttributes(val *ServiceInvitationDataAttributes) *NullableServiceInvitationDataAttributes {
	return &NullableServiceInvitationDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceInvitationDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceInvitationDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
