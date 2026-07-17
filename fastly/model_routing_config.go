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

// RoutingConfig All attributes for creating a routing config.
type RoutingConfig struct {
	// The user-defined name for the routing config.
	Name                 string          `json:"name"`
	InitialVersion       *InitialVersion `json:"initial_version,omitempty"`
	AdditionalProperties map[string]any
}

type _RoutingConfig RoutingConfig

// NewRoutingConfig instantiates a new RoutingConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingConfig(name string) *RoutingConfig {
	this := RoutingConfig{}
	this.Name = name
	return &this
}

// NewRoutingConfigWithDefaults instantiates a new RoutingConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingConfigWithDefaults() *RoutingConfig {
	this := RoutingConfig{}
	return &this
}

// GetName returns the Name field value
func (o *RoutingConfig) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoutingConfig) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoutingConfig) SetName(v string) {
	o.Name = v
}

// GetInitialVersion returns the InitialVersion field value if set, zero value otherwise.
func (o *RoutingConfig) GetInitialVersion() InitialVersion {
	if o == nil || o.InitialVersion == nil {
		var ret InitialVersion
		return ret
	}
	return *o.InitialVersion
}

// GetInitialVersionOk returns a tuple with the InitialVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingConfig) GetInitialVersionOk() (*InitialVersion, bool) {
	if o == nil || o.InitialVersion == nil {
		return nil, false
	}
	return o.InitialVersion, true
}

// HasInitialVersion returns a boolean if a field has been set.
func (o *RoutingConfig) HasInitialVersion() bool {
	if o != nil && o.InitialVersion != nil {
		return true
	}

	return false
}

// SetInitialVersion gets a reference to the given InitialVersion and assigns it to the InitialVersion field.
func (o *RoutingConfig) SetInitialVersion(v InitialVersion) {
	o.InitialVersion = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RoutingConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.InitialVersion != nil {
		toSerialize["initial_version"] = o.InitialVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RoutingConfig) UnmarshalJSON(bytes []byte) (err error) {
	varRoutingConfig := _RoutingConfig{}

	if err = json.Unmarshal(bytes, &varRoutingConfig); err == nil {
		*o = RoutingConfig(varRoutingConfig)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "initial_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRoutingConfig is a helper abstraction for handling nullable routingconfig types.
type NullableRoutingConfig struct {
	value *RoutingConfig
	isSet bool
}

// Get returns the value.
func (v NullableRoutingConfig) Get() *RoutingConfig {
	return v.value
}

// Set modifies the value.
func (v *NullableRoutingConfig) Set(val *RoutingConfig) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRoutingConfig) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRoutingConfig) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRoutingConfig returns a pointer to a new instance of NullableRoutingConfig.
func NewNullableRoutingConfig(val *RoutingConfig) *NullableRoutingConfig {
	return &NullableRoutingConfig{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRoutingConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRoutingConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
