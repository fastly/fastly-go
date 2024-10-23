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

// ConfiguredProductResponseConfiguration struct for ConfiguredProductResponseConfiguration
type ConfiguredProductResponseConfiguration struct {
	// Workspace ID
	WorkspaceID *string `json:"workspace_id,omitempty"`
	// Traffic ramp
	TrafficRamp          *string `json:"traffic_ramp,omitempty"`
	AdditionalProperties map[string]any
}

type _ConfiguredProductResponseConfiguration ConfiguredProductResponseConfiguration

// NewConfiguredProductResponseConfiguration instantiates a new ConfiguredProductResponseConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfiguredProductResponseConfiguration() *ConfiguredProductResponseConfiguration {
	this := ConfiguredProductResponseConfiguration{}
	return &this
}

// NewConfiguredProductResponseConfigurationWithDefaults instantiates a new ConfiguredProductResponseConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfiguredProductResponseConfigurationWithDefaults() *ConfiguredProductResponseConfiguration {
	this := ConfiguredProductResponseConfiguration{}
	return &this
}

// GetWorkspaceID returns the WorkspaceID field value if set, zero value otherwise.
func (o *ConfiguredProductResponseConfiguration) GetWorkspaceID() string {
	if o == nil || o.WorkspaceID == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceID
}

// GetWorkspaceIDOk returns a tuple with the WorkspaceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguredProductResponseConfiguration) GetWorkspaceIDOk() (*string, bool) {
	if o == nil || o.WorkspaceID == nil {
		return nil, false
	}
	return o.WorkspaceID, true
}

// HasWorkspaceID returns a boolean if a field has been set.
func (o *ConfiguredProductResponseConfiguration) HasWorkspaceID() bool {
	if o != nil && o.WorkspaceID != nil {
		return true
	}

	return false
}

// SetWorkspaceID gets a reference to the given string and assigns it to the WorkspaceID field.
func (o *ConfiguredProductResponseConfiguration) SetWorkspaceID(v string) {
	o.WorkspaceID = &v
}

// GetTrafficRamp returns the TrafficRamp field value if set, zero value otherwise.
func (o *ConfiguredProductResponseConfiguration) GetTrafficRamp() string {
	if o == nil || o.TrafficRamp == nil {
		var ret string
		return ret
	}
	return *o.TrafficRamp
}

// GetTrafficRampOk returns a tuple with the TrafficRamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfiguredProductResponseConfiguration) GetTrafficRampOk() (*string, bool) {
	if o == nil || o.TrafficRamp == nil {
		return nil, false
	}
	return o.TrafficRamp, true
}

// HasTrafficRamp returns a boolean if a field has been set.
func (o *ConfiguredProductResponseConfiguration) HasTrafficRamp() bool {
	if o != nil && o.TrafficRamp != nil {
		return true
	}

	return false
}

// SetTrafficRamp gets a reference to the given string and assigns it to the TrafficRamp field.
func (o *ConfiguredProductResponseConfiguration) SetTrafficRamp(v string) {
	o.TrafficRamp = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ConfiguredProductResponseConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.WorkspaceID != nil {
		toSerialize["workspace_id"] = o.WorkspaceID
	}
	if o.TrafficRamp != nil {
		toSerialize["traffic_ramp"] = o.TrafficRamp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ConfiguredProductResponseConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varConfiguredProductResponseConfiguration := _ConfiguredProductResponseConfiguration{}

	if err = json.Unmarshal(bytes, &varConfiguredProductResponseConfiguration); err == nil {
		*o = ConfiguredProductResponseConfiguration(varConfiguredProductResponseConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "workspace_id")
		delete(additionalProperties, "traffic_ramp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableConfiguredProductResponseConfiguration is a helper abstraction for handling nullable configuredproductresponseconfiguration types.
type NullableConfiguredProductResponseConfiguration struct {
	value *ConfiguredProductResponseConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableConfiguredProductResponseConfiguration) Get() *ConfiguredProductResponseConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableConfiguredProductResponseConfiguration) Set(val *ConfiguredProductResponseConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableConfiguredProductResponseConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableConfiguredProductResponseConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableConfiguredProductResponseConfiguration returns a pointer to a new instance of NullableConfiguredProductResponseConfiguration.
func NewNullableConfiguredProductResponseConfiguration(val *ConfiguredProductResponseConfiguration) *NullableConfiguredProductResponseConfiguration {
	return &NullableConfiguredProductResponseConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableConfiguredProductResponseConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableConfiguredProductResponseConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
