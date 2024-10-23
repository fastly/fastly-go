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

// SetConfiguration struct for SetConfiguration
type SetConfiguration struct {
	// The new workspace_id. Required in the `PUT` request body when `product_id` is `ngwaf`. Optional in the `PATCH` request body for `ngwaf`.
	WorkspaceID *string `json:"workspace_id,omitempty"`
	// The new traffic ramp. Optional in the `PATCH` request body for `ngwaf`.
	TrafficRamp          *string `json:"traffic_ramp,omitempty"`
	AdditionalProperties map[string]any
}

type _SetConfiguration SetConfiguration

// NewSetConfiguration instantiates a new SetConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetConfiguration() *SetConfiguration {
	this := SetConfiguration{}
	return &this
}

// NewSetConfigurationWithDefaults instantiates a new SetConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetConfigurationWithDefaults() *SetConfiguration {
	this := SetConfiguration{}
	return &this
}

// GetWorkspaceID returns the WorkspaceID field value if set, zero value otherwise.
func (o *SetConfiguration) GetWorkspaceID() string {
	if o == nil || o.WorkspaceID == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceID
}

// GetWorkspaceIDOk returns a tuple with the WorkspaceID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetConfiguration) GetWorkspaceIDOk() (*string, bool) {
	if o == nil || o.WorkspaceID == nil {
		return nil, false
	}
	return o.WorkspaceID, true
}

// HasWorkspaceID returns a boolean if a field has been set.
func (o *SetConfiguration) HasWorkspaceID() bool {
	if o != nil && o.WorkspaceID != nil {
		return true
	}

	return false
}

// SetWorkspaceID gets a reference to the given string and assigns it to the WorkspaceID field.
func (o *SetConfiguration) SetWorkspaceID(v string) {
	o.WorkspaceID = &v
}

// GetTrafficRamp returns the TrafficRamp field value if set, zero value otherwise.
func (o *SetConfiguration) GetTrafficRamp() string {
	if o == nil || o.TrafficRamp == nil {
		var ret string
		return ret
	}
	return *o.TrafficRamp
}

// GetTrafficRampOk returns a tuple with the TrafficRamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetConfiguration) GetTrafficRampOk() (*string, bool) {
	if o == nil || o.TrafficRamp == nil {
		return nil, false
	}
	return o.TrafficRamp, true
}

// HasTrafficRamp returns a boolean if a field has been set.
func (o *SetConfiguration) HasTrafficRamp() bool {
	if o != nil && o.TrafficRamp != nil {
		return true
	}

	return false
}

// SetTrafficRamp gets a reference to the given string and assigns it to the TrafficRamp field.
func (o *SetConfiguration) SetTrafficRamp(v string) {
	o.TrafficRamp = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o SetConfiguration) MarshalJSON() ([]byte, error) {
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
func (o *SetConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varSetConfiguration := _SetConfiguration{}

	if err = json.Unmarshal(bytes, &varSetConfiguration); err == nil {
		*o = SetConfiguration(varSetConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "workspace_id")
		delete(additionalProperties, "traffic_ramp")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableSetConfiguration is a helper abstraction for handling nullable setconfiguration types.
type NullableSetConfiguration struct {
	value *SetConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableSetConfiguration) Get() *SetConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableSetConfiguration) Set(val *SetConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableSetConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableSetConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableSetConfiguration returns a pointer to a new instance of NullableSetConfiguration.
func NewNullableSetConfiguration(val *SetConfiguration) *NullableSetConfiguration {
	return &NullableSetConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableSetConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableSetConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
