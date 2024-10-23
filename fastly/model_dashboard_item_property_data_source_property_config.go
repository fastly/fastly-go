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

// DashboardItemPropertyDataSourcePropertyConfig [Configuration options](#data-source-config) for the selected data source.
type DashboardItemPropertyDataSourcePropertyConfig struct {
	// The metrics to visualize. Valid options are defined by the selected [data source](#field_data_source).
	Metrics              []string `json:"metrics"`
	AdditionalProperties map[string]any
}

type _DashboardItemPropertyDataSourcePropertyConfig DashboardItemPropertyDataSourcePropertyConfig

// NewDashboardItemPropertyDataSourcePropertyConfig instantiates a new DashboardItemPropertyDataSourcePropertyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardItemPropertyDataSourcePropertyConfig(metrics []string) *DashboardItemPropertyDataSourcePropertyConfig {
	this := DashboardItemPropertyDataSourcePropertyConfig{}
	this.Metrics = metrics
	return &this
}

// NewDashboardItemPropertyDataSourcePropertyConfigWithDefaults instantiates a new DashboardItemPropertyDataSourcePropertyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardItemPropertyDataSourcePropertyConfigWithDefaults() *DashboardItemPropertyDataSourcePropertyConfig {
	this := DashboardItemPropertyDataSourcePropertyConfig{}
	return &this
}

// GetMetrics returns the Metrics field value
func (o *DashboardItemPropertyDataSourcePropertyConfig) GetMetrics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *DashboardItemPropertyDataSourcePropertyConfig) GetMetricsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics, true
}

// SetMetrics sets field value
func (o *DashboardItemPropertyDataSourcePropertyConfig) SetMetrics(v []string) {
	o.Metrics = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DashboardItemPropertyDataSourcePropertyConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["metrics"] = o.Metrics
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DashboardItemPropertyDataSourcePropertyConfig) UnmarshalJSON(bytes []byte) (err error) {
	varDashboardItemPropertyDataSourcePropertyConfig := _DashboardItemPropertyDataSourcePropertyConfig{}

	if err = json.Unmarshal(bytes, &varDashboardItemPropertyDataSourcePropertyConfig); err == nil {
		*o = DashboardItemPropertyDataSourcePropertyConfig(varDashboardItemPropertyDataSourcePropertyConfig)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "metrics")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDashboardItemPropertyDataSourcePropertyConfig is a helper abstraction for handling nullable dashboarditempropertydatasourcepropertyconfig types.
type NullableDashboardItemPropertyDataSourcePropertyConfig struct {
	value *DashboardItemPropertyDataSourcePropertyConfig
	isSet bool
}

// Get returns the value.
func (v NullableDashboardItemPropertyDataSourcePropertyConfig) Get() *DashboardItemPropertyDataSourcePropertyConfig {
	return v.value
}

// Set modifies the value.
func (v *NullableDashboardItemPropertyDataSourcePropertyConfig) Set(val *DashboardItemPropertyDataSourcePropertyConfig) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDashboardItemPropertyDataSourcePropertyConfig) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDashboardItemPropertyDataSourcePropertyConfig) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardItemPropertyDataSourcePropertyConfig returns a pointer to a new instance of NullableDashboardItemPropertyDataSourcePropertyConfig.
func NewNullableDashboardItemPropertyDataSourcePropertyConfig(val *DashboardItemPropertyDataSourcePropertyConfig) *NullableDashboardItemPropertyDataSourcePropertyConfig {
	return &NullableDashboardItemPropertyDataSourcePropertyConfig{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDashboardItemPropertyDataSourcePropertyConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDashboardItemPropertyDataSourcePropertyConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
