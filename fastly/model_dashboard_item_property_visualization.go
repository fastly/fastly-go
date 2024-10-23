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

// DashboardItemPropertyVisualization [An object](#visualization) which describes the data visualization to display.
type DashboardItemPropertyVisualization struct {
	// The type of visualization to display.
	Type                 string                                           `json:"type"`
	Config               DashboardItemPropertyVisualizationPropertyConfig `json:"config"`
	AdditionalProperties map[string]any
}

type _DashboardItemPropertyVisualization DashboardItemPropertyVisualization

// NewDashboardItemPropertyVisualization instantiates a new DashboardItemPropertyVisualization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardItemPropertyVisualization(resourceType string, config DashboardItemPropertyVisualizationPropertyConfig) *DashboardItemPropertyVisualization {
	this := DashboardItemPropertyVisualization{}
	this.Type = resourceType
	this.Config = config
	return &this
}

// NewDashboardItemPropertyVisualizationWithDefaults instantiates a new DashboardItemPropertyVisualization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardItemPropertyVisualizationWithDefaults() *DashboardItemPropertyVisualization {
	this := DashboardItemPropertyVisualization{}
	return &this
}

// GetType returns the Type field value
func (o *DashboardItemPropertyVisualization) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DashboardItemPropertyVisualization) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DashboardItemPropertyVisualization) SetType(v string) {
	o.Type = v
}

// GetConfig returns the Config field value
func (o *DashboardItemPropertyVisualization) GetConfig() DashboardItemPropertyVisualizationPropertyConfig {
	if o == nil {
		var ret DashboardItemPropertyVisualizationPropertyConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *DashboardItemPropertyVisualization) GetConfigOk() (*DashboardItemPropertyVisualizationPropertyConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *DashboardItemPropertyVisualization) SetConfig(v DashboardItemPropertyVisualizationPropertyConfig) {
	o.Config = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DashboardItemPropertyVisualization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["config"] = o.Config
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DashboardItemPropertyVisualization) UnmarshalJSON(bytes []byte) (err error) {
	varDashboardItemPropertyVisualization := _DashboardItemPropertyVisualization{}

	if err = json.Unmarshal(bytes, &varDashboardItemPropertyVisualization); err == nil {
		*o = DashboardItemPropertyVisualization(varDashboardItemPropertyVisualization)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDashboardItemPropertyVisualization is a helper abstraction for handling nullable dashboarditempropertyvisualization types.
type NullableDashboardItemPropertyVisualization struct {
	value *DashboardItemPropertyVisualization
	isSet bool
}

// Get returns the value.
func (v NullableDashboardItemPropertyVisualization) Get() *DashboardItemPropertyVisualization {
	return v.value
}

// Set modifies the value.
func (v *NullableDashboardItemPropertyVisualization) Set(val *DashboardItemPropertyVisualization) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDashboardItemPropertyVisualization) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDashboardItemPropertyVisualization) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardItemPropertyVisualization returns a pointer to a new instance of NullableDashboardItemPropertyVisualization.
func NewNullableDashboardItemPropertyVisualization(val *DashboardItemPropertyVisualization) *NullableDashboardItemPropertyVisualization {
	return &NullableDashboardItemPropertyVisualization{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDashboardItemPropertyVisualization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDashboardItemPropertyVisualization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
