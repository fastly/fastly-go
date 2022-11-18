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

// RelationshipService struct for RelationshipService
type RelationshipService struct {
	Service *RelationshipMemberService `json:"service,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipService RelationshipService

// NewRelationshipService instantiates a new RelationshipService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipService() *RelationshipService {
	this := RelationshipService{}
	return &this
}

// NewRelationshipServiceWithDefaults instantiates a new RelationshipService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipServiceWithDefaults() *RelationshipService {
	this := RelationshipService{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *RelationshipService) GetService() RelationshipMemberService {
	if o == nil || o.Service == nil {
		var ret RelationshipMemberService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipService) GetServiceOk() (*RelationshipMemberService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *RelationshipService) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given RelationshipMemberService and assigns it to the Service field.
func (o *RelationshipService) SetService(v RelationshipMemberService) {
	o.Service = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipService) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipService := _RelationshipService{}

	if err = json.Unmarshal(bytes, &varRelationshipService); err == nil {
		*o = RelationshipService(varRelationshipService)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipService is a helper abstraction for handling nullable relationshipservice types. 
type NullableRelationshipService struct {
	value *RelationshipService
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipService) Get() *RelationshipService {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipService) Set(val *RelationshipService) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipService) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipService) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipService returns a pointer to a new instance of NullableRelationshipService.
func NewNullableRelationshipService(val *RelationshipService) *NullableRelationshipService {
	return &NullableRelationshipService{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
