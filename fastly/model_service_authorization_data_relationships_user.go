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

// ServiceAuthorizationDataRelationshipsUser The ID of the user being given access to the service.
type ServiceAuthorizationDataRelationshipsUser struct {
	Data                 *ServiceAuthorizationDataRelationshipsUserData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceAuthorizationDataRelationshipsUser ServiceAuthorizationDataRelationshipsUser

// NewServiceAuthorizationDataRelationshipsUser instantiates a new ServiceAuthorizationDataRelationshipsUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAuthorizationDataRelationshipsUser() *ServiceAuthorizationDataRelationshipsUser {
	this := ServiceAuthorizationDataRelationshipsUser{}
	return &this
}

// NewServiceAuthorizationDataRelationshipsUserWithDefaults instantiates a new ServiceAuthorizationDataRelationshipsUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAuthorizationDataRelationshipsUserWithDefaults() *ServiceAuthorizationDataRelationshipsUser {
	this := ServiceAuthorizationDataRelationshipsUser{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ServiceAuthorizationDataRelationshipsUser) GetData() ServiceAuthorizationDataRelationshipsUserData {
	if o == nil || o.Data == nil {
		var ret ServiceAuthorizationDataRelationshipsUserData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAuthorizationDataRelationshipsUser) GetDataOk() (*ServiceAuthorizationDataRelationshipsUserData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ServiceAuthorizationDataRelationshipsUser) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given ServiceAuthorizationDataRelationshipsUserData and assigns it to the Data field.
func (o *ServiceAuthorizationDataRelationshipsUser) SetData(v ServiceAuthorizationDataRelationshipsUserData) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceAuthorizationDataRelationshipsUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ServiceAuthorizationDataRelationshipsUser) UnmarshalJSON(bytes []byte) (err error) {
	varServiceAuthorizationDataRelationshipsUser := _ServiceAuthorizationDataRelationshipsUser{}

	if err = json.Unmarshal(bytes, &varServiceAuthorizationDataRelationshipsUser); err == nil {
		*o = ServiceAuthorizationDataRelationshipsUser(varServiceAuthorizationDataRelationshipsUser)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceAuthorizationDataRelationshipsUser is a helper abstraction for handling nullable serviceauthorizationdatarelationshipsuser types.
type NullableServiceAuthorizationDataRelationshipsUser struct {
	value *ServiceAuthorizationDataRelationshipsUser
	isSet bool
}

// Get returns the value.
func (v NullableServiceAuthorizationDataRelationshipsUser) Get() *ServiceAuthorizationDataRelationshipsUser {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceAuthorizationDataRelationshipsUser) Set(val *ServiceAuthorizationDataRelationshipsUser) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceAuthorizationDataRelationshipsUser) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceAuthorizationDataRelationshipsUser) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceAuthorizationDataRelationshipsUser returns a pointer to a new instance of NullableServiceAuthorizationDataRelationshipsUser.
func NewNullableServiceAuthorizationDataRelationshipsUser(val *ServiceAuthorizationDataRelationshipsUser) *NullableServiceAuthorizationDataRelationshipsUser {
	return &NullableServiceAuthorizationDataRelationshipsUser{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceAuthorizationDataRelationshipsUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableServiceAuthorizationDataRelationshipsUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
