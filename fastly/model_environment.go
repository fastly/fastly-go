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

// Environment struct for Environment
type Environment struct {
	Name *string `json:"name,omitempty"`
	// Alphanumeric string identifying the service.
	ServiceId            *string `json:"service_id,omitempty"`
	ActiveVersion        *int32  `json:"active_version,omitempty"`
	AdditionalProperties map[string]any
}

type _Environment Environment

// NewEnvironment instantiates a new Environment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironment() *Environment {
	this := Environment{}
	return &this
}

// NewEnvironmentWithDefaults instantiates a new Environment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentWithDefaults() *Environment {
	this := Environment{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Environment) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Environment) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Environment) SetName(v string) {
	o.Name = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *Environment) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *Environment) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *Environment) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetActiveVersion returns the ActiveVersion field value if set, zero value otherwise.
func (o *Environment) GetActiveVersion() int32 {
	if o == nil || o.ActiveVersion == nil {
		var ret int32
		return ret
	}
	return *o.ActiveVersion
}

// GetActiveVersionOk returns a tuple with the ActiveVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetActiveVersionOk() (*int32, bool) {
	if o == nil || o.ActiveVersion == nil {
		return nil, false
	}
	return o.ActiveVersion, true
}

// HasActiveVersion returns a boolean if a field has been set.
func (o *Environment) HasActiveVersion() bool {
	if o != nil && o.ActiveVersion != nil {
		return true
	}

	return false
}

// SetActiveVersion gets a reference to the given int32 and assigns it to the ActiveVersion field.
func (o *Environment) SetActiveVersion(v int32) {
	o.ActiveVersion = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Environment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.ActiveVersion != nil {
		toSerialize["active_version"] = o.ActiveVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Environment) UnmarshalJSON(bytes []byte) (err error) {
	varEnvironment := _Environment{}

	if err = json.Unmarshal(bytes, &varEnvironment); err == nil {
		*o = Environment(varEnvironment)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "active_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableEnvironment is a helper abstraction for handling nullable environment types.
type NullableEnvironment struct {
	value *Environment
	isSet bool
}

// Get returns the value.
func (v NullableEnvironment) Get() *Environment {
	return v.value
}

// Set modifies the value.
func (v *NullableEnvironment) Set(val *Environment) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableEnvironment) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableEnvironment returns a pointer to a new instance of NullableEnvironment.
func NewNullableEnvironment(val *Environment) *NullableEnvironment {
	return &NullableEnvironment{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
