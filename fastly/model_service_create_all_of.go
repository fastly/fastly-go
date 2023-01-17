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

// ServiceCreateAllOf struct for ServiceCreateAllOf
type ServiceCreateAllOf struct {
	// The type of this service.
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceCreateAllOf ServiceCreateAllOf

// NewServiceCreateAllOf instantiates a new ServiceCreateAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceCreateAllOf() *ServiceCreateAllOf {
	this := ServiceCreateAllOf{}
	return &this
}

// NewServiceCreateAllOfWithDefaults instantiates a new ServiceCreateAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceCreateAllOfWithDefaults() *ServiceCreateAllOf {
	this := ServiceCreateAllOf{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceCreateAllOf) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCreateAllOf) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceCreateAllOf) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceCreateAllOf) SetType(v string) {
	o.Type = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceCreateAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ServiceCreateAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServiceCreateAllOf := _ServiceCreateAllOf{}

	if err = json.Unmarshal(bytes, &varServiceCreateAllOf); err == nil {
		*o = ServiceCreateAllOf(varServiceCreateAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceCreateAllOf is a helper abstraction for handling nullable servicecreateallof types. 
type NullableServiceCreateAllOf struct {
	value *ServiceCreateAllOf
	isSet bool
}

// Get returns the value.
func (v NullableServiceCreateAllOf) Get() *ServiceCreateAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceCreateAllOf) Set(val *ServiceCreateAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceCreateAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceCreateAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceCreateAllOf returns a pointer to a new instance of NullableServiceCreateAllOf.
func NewNullableServiceCreateAllOf(val *ServiceCreateAllOf) *NullableServiceCreateAllOf {
	return &NullableServiceCreateAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceCreateAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableServiceCreateAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
