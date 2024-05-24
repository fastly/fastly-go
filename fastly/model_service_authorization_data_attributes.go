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

// ServiceAuthorizationDataAttributes struct for ServiceAuthorizationDataAttributes
type ServiceAuthorizationDataAttributes struct {
	Permission *Permission `json:"permission,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceAuthorizationDataAttributes ServiceAuthorizationDataAttributes

// NewServiceAuthorizationDataAttributes instantiates a new ServiceAuthorizationDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAuthorizationDataAttributes() *ServiceAuthorizationDataAttributes {
	this := ServiceAuthorizationDataAttributes{}
	var permission Permission = PERMISSION_FULL
	this.Permission = &permission
	return &this
}

// NewServiceAuthorizationDataAttributesWithDefaults instantiates a new ServiceAuthorizationDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAuthorizationDataAttributesWithDefaults() *ServiceAuthorizationDataAttributes {
	this := ServiceAuthorizationDataAttributes{}
	var permission Permission = PERMISSION_FULL
	this.Permission = &permission
	return &this
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *ServiceAuthorizationDataAttributes) GetPermission() Permission {
	if o == nil || o.Permission == nil {
		var ret Permission
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAuthorizationDataAttributes) GetPermissionOk() (*Permission, bool) {
	if o == nil || o.Permission == nil {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *ServiceAuthorizationDataAttributes) HasPermission() bool {
	if o != nil && o.Permission != nil {
		return true
	}

	return false
}

// SetPermission gets a reference to the given Permission and assigns it to the Permission field.
func (o *ServiceAuthorizationDataAttributes) SetPermission(v Permission) {
	o.Permission = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceAuthorizationDataAttributes) MarshalJSON() ([]byte, error) {
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
func (o *ServiceAuthorizationDataAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varServiceAuthorizationDataAttributes := _ServiceAuthorizationDataAttributes{}

	if err = json.Unmarshal(bytes, &varServiceAuthorizationDataAttributes); err == nil {
		*o = ServiceAuthorizationDataAttributes(varServiceAuthorizationDataAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "permission")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceAuthorizationDataAttributes is a helper abstraction for handling nullable serviceauthorizationdataattributes types. 
type NullableServiceAuthorizationDataAttributes struct {
	value *ServiceAuthorizationDataAttributes
	isSet bool
}

// Get returns the value.
func (v NullableServiceAuthorizationDataAttributes) Get() *ServiceAuthorizationDataAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceAuthorizationDataAttributes) Set(val *ServiceAuthorizationDataAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceAuthorizationDataAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceAuthorizationDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceAuthorizationDataAttributes returns a pointer to a new instance of NullableServiceAuthorizationDataAttributes.
func NewNullableServiceAuthorizationDataAttributes(val *ServiceAuthorizationDataAttributes) *NullableServiceAuthorizationDataAttributes {
	return &NullableServiceAuthorizationDataAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceAuthorizationDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceAuthorizationDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
