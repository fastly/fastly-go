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

// NgwafRequestUpdateConfiguration struct for NgwafRequestUpdateConfiguration
type NgwafRequestUpdateConfiguration struct {
	// The workspace to link.
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// The percentage of traffic to inspect.
	TrafficRamp          *string `json:"traffic_ramp,omitempty"`
	AdditionalProperties map[string]any
}

type _NgwafRequestUpdateConfiguration NgwafRequestUpdateConfiguration

// NewNgwafRequestUpdateConfiguration instantiates a new NgwafRequestUpdateConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgwafRequestUpdateConfiguration() *NgwafRequestUpdateConfiguration {
	this := NgwafRequestUpdateConfiguration{}
	return &this
}

// NewNgwafRequestUpdateConfigurationWithDefaults instantiates a new NgwafRequestUpdateConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgwafRequestUpdateConfigurationWithDefaults() *NgwafRequestUpdateConfiguration {
	this := NgwafRequestUpdateConfiguration{}
	return &this
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *NgwafRequestUpdateConfiguration) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgwafRequestUpdateConfiguration) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *NgwafRequestUpdateConfiguration) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *NgwafRequestUpdateConfiguration) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

// GetTrafficRamp returns the TrafficRamp field value if set, zero value otherwise.
func (o *NgwafRequestUpdateConfiguration) GetTrafficRamp() string {
	if o == nil || o.TrafficRamp == nil {
		var ret string
		return ret
	}
	return *o.TrafficRamp
}

// GetTrafficRampOk returns a tuple with the TrafficRamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NgwafRequestUpdateConfiguration) GetTrafficRampOk() (*string, bool) {
	if o == nil || o.TrafficRamp == nil {
		return nil, false
	}
	return o.TrafficRamp, true
}

// HasTrafficRamp returns a boolean if a field has been set.
func (o *NgwafRequestUpdateConfiguration) HasTrafficRamp() bool {
	if o != nil && o.TrafficRamp != nil {
		return true
	}

	return false
}

// SetTrafficRamp gets a reference to the given string and assigns it to the TrafficRamp field.
func (o *NgwafRequestUpdateConfiguration) SetTrafficRamp(v string) {
	o.TrafficRamp = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NgwafRequestUpdateConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.WorkspaceId != nil {
		toSerialize["workspace_id"] = o.WorkspaceId
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
func (o *NgwafRequestUpdateConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varNgwafRequestUpdateConfiguration := _NgwafRequestUpdateConfiguration{}

	if err = json.Unmarshal(bytes, &varNgwafRequestUpdateConfiguration); err == nil {
		*o = NgwafRequestUpdateConfiguration(varNgwafRequestUpdateConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "workspace_id")
		delete(additionalProperties, "traffic_ramp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNgwafRequestUpdateConfiguration is a helper abstraction for handling nullable ngwafrequestupdateconfiguration types.
type NullableNgwafRequestUpdateConfiguration struct {
	value *NgwafRequestUpdateConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableNgwafRequestUpdateConfiguration) Get() *NgwafRequestUpdateConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableNgwafRequestUpdateConfiguration) Set(val *NgwafRequestUpdateConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNgwafRequestUpdateConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNgwafRequestUpdateConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNgwafRequestUpdateConfiguration returns a pointer to a new instance of NullableNgwafRequestUpdateConfiguration.
func NewNullableNgwafRequestUpdateConfiguration(val *NgwafRequestUpdateConfiguration) *NullableNgwafRequestUpdateConfiguration {
	return &NullableNgwafRequestUpdateConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNgwafRequestUpdateConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNgwafRequestUpdateConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
