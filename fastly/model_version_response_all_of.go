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

// VersionResponseAllOf struct for VersionResponseAllOf
type VersionResponseAllOf struct {
	ServiceId *string `json:"service_id,omitempty"`
	// A list of environments where the service has been deployed.
	Environments         []Environment `json:"environments,omitempty"`
	AdditionalProperties map[string]any
}

type _VersionResponseAllOf VersionResponseAllOf

// NewVersionResponseAllOf instantiates a new VersionResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionResponseAllOf() *VersionResponseAllOf {
	this := VersionResponseAllOf{}
	return &this
}

// NewVersionResponseAllOfWithDefaults instantiates a new VersionResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionResponseAllOfWithDefaults() *VersionResponseAllOf {
	this := VersionResponseAllOf{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *VersionResponseAllOf) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionResponseAllOf) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *VersionResponseAllOf) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *VersionResponseAllOf) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetEnvironments returns the Environments field value if set, zero value otherwise.
func (o *VersionResponseAllOf) GetEnvironments() []Environment {
	if o == nil || o.Environments == nil {
		var ret []Environment
		return ret
	}
	return o.Environments
}

// GetEnvironmentsOk returns a tuple with the Environments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionResponseAllOf) GetEnvironmentsOk() ([]Environment, bool) {
	if o == nil || o.Environments == nil {
		return nil, false
	}
	return o.Environments, true
}

// HasEnvironments returns a boolean if a field has been set.
func (o *VersionResponseAllOf) HasEnvironments() bool {
	if o != nil && o.Environments != nil {
		return true
	}

	return false
}

// SetEnvironments gets a reference to the given []Environment and assigns it to the Environments field.
func (o *VersionResponseAllOf) SetEnvironments(v []Environment) {
	o.Environments = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o VersionResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}
	if o.Environments != nil {
		toSerialize["environments"] = o.Environments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *VersionResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varVersionResponseAllOf := _VersionResponseAllOf{}

	if err = json.Unmarshal(bytes, &varVersionResponseAllOf); err == nil {
		*o = VersionResponseAllOf(varVersionResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service_id")
		delete(additionalProperties, "environments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableVersionResponseAllOf is a helper abstraction for handling nullable versionresponseallof types.
type NullableVersionResponseAllOf struct {
	value *VersionResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableVersionResponseAllOf) Get() *VersionResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableVersionResponseAllOf) Set(val *VersionResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableVersionResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableVersionResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableVersionResponseAllOf returns a pointer to a new instance of NullableVersionResponseAllOf.
func NewNullableVersionResponseAllOf(val *VersionResponseAllOf) *NullableVersionResponseAllOf {
	return &NullableVersionResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableVersionResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableVersionResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
