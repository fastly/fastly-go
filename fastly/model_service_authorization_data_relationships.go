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

// ServiceAuthorizationDataRelationships struct for ServiceAuthorizationDataRelationships
type ServiceAuthorizationDataRelationships struct {
	Service *RelationshipMemberService `json:"service,omitempty"`
	User *ServiceAuthorizationDataRelationshipsUser `json:"user,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceAuthorizationDataRelationships ServiceAuthorizationDataRelationships

// NewServiceAuthorizationDataRelationships instantiates a new ServiceAuthorizationDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAuthorizationDataRelationships() *ServiceAuthorizationDataRelationships {
	this := ServiceAuthorizationDataRelationships{}
	return &this
}

// NewServiceAuthorizationDataRelationshipsWithDefaults instantiates a new ServiceAuthorizationDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAuthorizationDataRelationshipsWithDefaults() *ServiceAuthorizationDataRelationships {
	this := ServiceAuthorizationDataRelationships{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ServiceAuthorizationDataRelationships) GetService() RelationshipMemberService {
	if o == nil || o.Service == nil {
		var ret RelationshipMemberService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAuthorizationDataRelationships) GetServiceOk() (*RelationshipMemberService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ServiceAuthorizationDataRelationships) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given RelationshipMemberService and assigns it to the Service field.
func (o *ServiceAuthorizationDataRelationships) SetService(v RelationshipMemberService) {
	o.Service = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ServiceAuthorizationDataRelationships) GetUser() ServiceAuthorizationDataRelationshipsUser {
	if o == nil || o.User == nil {
		var ret ServiceAuthorizationDataRelationshipsUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAuthorizationDataRelationships) GetUserOk() (*ServiceAuthorizationDataRelationshipsUser, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ServiceAuthorizationDataRelationships) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given ServiceAuthorizationDataRelationshipsUser and assigns it to the User field.
func (o *ServiceAuthorizationDataRelationships) SetUser(v ServiceAuthorizationDataRelationshipsUser) {
	o.User = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceAuthorizationDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ServiceAuthorizationDataRelationships) UnmarshalJSON(bytes []byte) (err error) {
	varServiceAuthorizationDataRelationships := _ServiceAuthorizationDataRelationships{}

	if err = json.Unmarshal(bytes, &varServiceAuthorizationDataRelationships); err == nil {
		*o = ServiceAuthorizationDataRelationships(varServiceAuthorizationDataRelationships)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service")
		delete(additionalProperties, "user")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceAuthorizationDataRelationships is a helper abstraction for handling nullable serviceauthorizationdatarelationships types. 
type NullableServiceAuthorizationDataRelationships struct {
	value *ServiceAuthorizationDataRelationships
	isSet bool
}

// Get returns the value.
func (v NullableServiceAuthorizationDataRelationships) Get() *ServiceAuthorizationDataRelationships {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceAuthorizationDataRelationships) Set(val *ServiceAuthorizationDataRelationships) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceAuthorizationDataRelationships) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceAuthorizationDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceAuthorizationDataRelationships returns a pointer to a new instance of NullableServiceAuthorizationDataRelationships.
func NewNullableServiceAuthorizationDataRelationships(val *ServiceAuthorizationDataRelationships) *NullableServiceAuthorizationDataRelationships {
	return &NullableServiceAuthorizationDataRelationships{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceAuthorizationDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceAuthorizationDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
