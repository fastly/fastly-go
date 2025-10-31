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

// ServiceIdAndVersion struct for ServiceIdAndVersion
type ServiceIdAndVersion struct {
	ServiceId            *string `json:"service_id,omitempty"`
	Version              *int32  `json:"version,omitempty"`
	AdditionalProperties map[string]any
}

type _ServiceIdAndVersion ServiceIdAndVersion

// NewServiceIdAndVersion instantiates a new ServiceIdAndVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceIdAndVersion() *ServiceIdAndVersion {
	this := ServiceIdAndVersion{}
	return &this
}

// NewServiceIdAndVersionWithDefaults instantiates a new ServiceIdAndVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceIdAndVersionWithDefaults() *ServiceIdAndVersion {
	this := ServiceIdAndVersion{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ServiceIdAndVersion) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceIdAndVersion) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ServiceIdAndVersion) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *ServiceIdAndVersion) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ServiceIdAndVersion) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceIdAndVersion) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ServiceIdAndVersion) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ServiceIdAndVersion) SetVersion(v int32) {
	o.Version = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ServiceIdAndVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ServiceIdAndVersion) UnmarshalJSON(bytes []byte) (err error) {
	varServiceIdAndVersion := _ServiceIdAndVersion{}

	if err = json.Unmarshal(bytes, &varServiceIdAndVersion); err == nil {
		*o = ServiceIdAndVersion(varServiceIdAndVersion)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableServiceIdAndVersion is a helper abstraction for handling nullable serviceidandversion types.
type NullableServiceIdAndVersion struct {
	value *ServiceIdAndVersion
	isSet bool
}

// Get returns the value.
func (v NullableServiceIdAndVersion) Get() *ServiceIdAndVersion {
	return v.value
}

// Set modifies the value.
func (v *NullableServiceIdAndVersion) Set(val *ServiceIdAndVersion) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableServiceIdAndVersion) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableServiceIdAndVersion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableServiceIdAndVersion returns a pointer to a new instance of NullableServiceIdAndVersion.
func NewNullableServiceIdAndVersion(val *ServiceIdAndVersion) *NullableServiceIdAndVersion {
	return &NullableServiceIdAndVersion{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableServiceIdAndVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableServiceIdAndVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
