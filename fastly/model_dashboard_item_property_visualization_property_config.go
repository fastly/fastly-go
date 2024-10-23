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

// DashboardItemPropertyVisualizationPropertyConfig [Configuration options](#visualization-config) for the given visualization.
type DashboardItemPropertyVisualizationPropertyConfig struct {
	// The type of chart to display.
	PlotType string `json:"plot_type"`
	// (Optional) The units to use to format the data.
	Format *string `json:"format,omitempty"`
	// (Optional) The aggregation function to apply to the dataset.
	CalculationMethod    *string `json:"calculation_method,omitempty"`
	AdditionalProperties map[string]any
}

type _DashboardItemPropertyVisualizationPropertyConfig DashboardItemPropertyVisualizationPropertyConfig

// NewDashboardItemPropertyVisualizationPropertyConfig instantiates a new DashboardItemPropertyVisualizationPropertyConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardItemPropertyVisualizationPropertyConfig(plotType string) *DashboardItemPropertyVisualizationPropertyConfig {
	this := DashboardItemPropertyVisualizationPropertyConfig{}
	this.PlotType = plotType
	var format string = "number"
	this.Format = &format
	return &this
}

// NewDashboardItemPropertyVisualizationPropertyConfigWithDefaults instantiates a new DashboardItemPropertyVisualizationPropertyConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardItemPropertyVisualizationPropertyConfigWithDefaults() *DashboardItemPropertyVisualizationPropertyConfig {
	this := DashboardItemPropertyVisualizationPropertyConfig{}
	var format string = "number"
	this.Format = &format
	return &this
}

// GetPlotType returns the PlotType field value
func (o *DashboardItemPropertyVisualizationPropertyConfig) GetPlotType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlotType
}

// GetPlotTypeOk returns a tuple with the PlotType field value
// and a boolean to check if the value has been set.
func (o *DashboardItemPropertyVisualizationPropertyConfig) GetPlotTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlotType, true
}

// SetPlotType sets field value
func (o *DashboardItemPropertyVisualizationPropertyConfig) SetPlotType(v string) {
	o.PlotType = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *DashboardItemPropertyVisualizationPropertyConfig) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardItemPropertyVisualizationPropertyConfig) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *DashboardItemPropertyVisualizationPropertyConfig) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *DashboardItemPropertyVisualizationPropertyConfig) SetFormat(v string) {
	o.Format = &v
}

// GetCalculationMethod returns the CalculationMethod field value if set, zero value otherwise.
func (o *DashboardItemPropertyVisualizationPropertyConfig) GetCalculationMethod() string {
	if o == nil || o.CalculationMethod == nil {
		var ret string
		return ret
	}
	return *o.CalculationMethod
}

// GetCalculationMethodOk returns a tuple with the CalculationMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardItemPropertyVisualizationPropertyConfig) GetCalculationMethodOk() (*string, bool) {
	if o == nil || o.CalculationMethod == nil {
		return nil, false
	}
	return o.CalculationMethod, true
}

// HasCalculationMethod returns a boolean if a field has been set.
func (o *DashboardItemPropertyVisualizationPropertyConfig) HasCalculationMethod() bool {
	if o != nil && o.CalculationMethod != nil {
		return true
	}

	return false
}

// SetCalculationMethod gets a reference to the given string and assigns it to the CalculationMethod field.
func (o *DashboardItemPropertyVisualizationPropertyConfig) SetCalculationMethod(v string) {
	o.CalculationMethod = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DashboardItemPropertyVisualizationPropertyConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["plot_type"] = o.PlotType
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.CalculationMethod != nil {
		toSerialize["calculation_method"] = o.CalculationMethod
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DashboardItemPropertyVisualizationPropertyConfig) UnmarshalJSON(bytes []byte) (err error) {
	varDashboardItemPropertyVisualizationPropertyConfig := _DashboardItemPropertyVisualizationPropertyConfig{}

	if err = json.Unmarshal(bytes, &varDashboardItemPropertyVisualizationPropertyConfig); err == nil {
		*o = DashboardItemPropertyVisualizationPropertyConfig(varDashboardItemPropertyVisualizationPropertyConfig)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "plot_type")
		delete(additionalProperties, "format")
		delete(additionalProperties, "calculation_method")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDashboardItemPropertyVisualizationPropertyConfig is a helper abstraction for handling nullable dashboarditempropertyvisualizationpropertyconfig types.
type NullableDashboardItemPropertyVisualizationPropertyConfig struct {
	value *DashboardItemPropertyVisualizationPropertyConfig
	isSet bool
}

// Get returns the value.
func (v NullableDashboardItemPropertyVisualizationPropertyConfig) Get() *DashboardItemPropertyVisualizationPropertyConfig {
	return v.value
}

// Set modifies the value.
func (v *NullableDashboardItemPropertyVisualizationPropertyConfig) Set(val *DashboardItemPropertyVisualizationPropertyConfig) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDashboardItemPropertyVisualizationPropertyConfig) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDashboardItemPropertyVisualizationPropertyConfig) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardItemPropertyVisualizationPropertyConfig returns a pointer to a new instance of NullableDashboardItemPropertyVisualizationPropertyConfig.
func NewNullableDashboardItemPropertyVisualizationPropertyConfig(val *DashboardItemPropertyVisualizationPropertyConfig) *NullableDashboardItemPropertyVisualizationPropertyConfig {
	return &NullableDashboardItemPropertyVisualizationPropertyConfig{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDashboardItemPropertyVisualizationPropertyConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDashboardItemPropertyVisualizationPropertyConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
