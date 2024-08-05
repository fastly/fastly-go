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

// DashboardItemPropertyDataSource [An object](#data-source) which describes the data to display.
type DashboardItemPropertyDataSource struct {
	// The source of the data to display.
	Type string `json:"type"`
	Config DashboardItemPropertyDataSourcePropertyConfig `json:"config"`
	AdditionalProperties map[string]any
}

type _DashboardItemPropertyDataSource DashboardItemPropertyDataSource

// NewDashboardItemPropertyDataSource instantiates a new DashboardItemPropertyDataSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardItemPropertyDataSource(resourceType string, config DashboardItemPropertyDataSourcePropertyConfig) *DashboardItemPropertyDataSource {
	this := DashboardItemPropertyDataSource{}
	this.Type = resourceType
	this.Config = config
	return &this
}

// NewDashboardItemPropertyDataSourceWithDefaults instantiates a new DashboardItemPropertyDataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardItemPropertyDataSourceWithDefaults() *DashboardItemPropertyDataSource {
	this := DashboardItemPropertyDataSource{}
	return &this
}

// GetType returns the Type field value
func (o *DashboardItemPropertyDataSource) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DashboardItemPropertyDataSource) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DashboardItemPropertyDataSource) SetType(v string) {
	o.Type = v
}

// GetConfig returns the Config field value
func (o *DashboardItemPropertyDataSource) GetConfig() DashboardItemPropertyDataSourcePropertyConfig {
	if o == nil {
		var ret DashboardItemPropertyDataSourcePropertyConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *DashboardItemPropertyDataSource) GetConfigOk() (*DashboardItemPropertyDataSourcePropertyConfig, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *DashboardItemPropertyDataSource) SetConfig(v DashboardItemPropertyDataSourcePropertyConfig) {
	o.Config = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DashboardItemPropertyDataSource) MarshalJSON() ([]byte, error) {
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
func (o *DashboardItemPropertyDataSource) UnmarshalJSON(bytes []byte) (err error) {
	varDashboardItemPropertyDataSource := _DashboardItemPropertyDataSource{}

	if err = json.Unmarshal(bytes, &varDashboardItemPropertyDataSource); err == nil {
		*o = DashboardItemPropertyDataSource(varDashboardItemPropertyDataSource)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "config")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDashboardItemPropertyDataSource is a helper abstraction for handling nullable dashboarditempropertydatasource types. 
type NullableDashboardItemPropertyDataSource struct {
	value *DashboardItemPropertyDataSource
	isSet bool
}

// Get returns the value.
func (v NullableDashboardItemPropertyDataSource) Get() *DashboardItemPropertyDataSource {
	return v.value
}

// Set modifies the value.
func (v *NullableDashboardItemPropertyDataSource) Set(val *DashboardItemPropertyDataSource) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDashboardItemPropertyDataSource) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDashboardItemPropertyDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardItemPropertyDataSource returns a pointer to a new instance of NullableDashboardItemPropertyDataSource.
func NewNullableDashboardItemPropertyDataSource(val *DashboardItemPropertyDataSource) *NullableDashboardItemPropertyDataSource {
	return &NullableDashboardItemPropertyDataSource{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDashboardItemPropertyDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableDashboardItemPropertyDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
