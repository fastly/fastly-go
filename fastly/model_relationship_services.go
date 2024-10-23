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

// RelationshipServices struct for RelationshipServices
type RelationshipServices struct {
	Services             *RelationshipServicesServices `json:"services,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipServices RelationshipServices

// NewRelationshipServices instantiates a new RelationshipServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipServices() *RelationshipServices {
	this := RelationshipServices{}
	return &this
}

// NewRelationshipServicesWithDefaults instantiates a new RelationshipServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipServicesWithDefaults() *RelationshipServices {
	this := RelationshipServices{}
	return &this
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *RelationshipServices) GetServices() RelationshipServicesServices {
	if o == nil || o.Services == nil {
		var ret RelationshipServicesServices
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipServices) GetServicesOk() (*RelationshipServicesServices, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *RelationshipServices) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given RelationshipServicesServices and assigns it to the Services field.
func (o *RelationshipServices) SetServices(v RelationshipServicesServices) {
	o.Services = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipServices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipServices) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipServices := _RelationshipServices{}

	if err = json.Unmarshal(bytes, &varRelationshipServices); err == nil {
		*o = RelationshipServices(varRelationshipServices)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "services")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipServices is a helper abstraction for handling nullable relationshipservices types.
type NullableRelationshipServices struct {
	value *RelationshipServices
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipServices) Get() *RelationshipServices {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipServices) Set(val *RelationshipServices) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipServices) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipServices) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipServices returns a pointer to a new instance of NullableRelationshipServices.
func NewNullableRelationshipServices(val *RelationshipServices) *NullableRelationshipServices {
	return &NullableRelationshipServices{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
