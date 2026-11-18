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

// DashboardItem A dashboard item. Typically a data visualization like a chart.
type DashboardItem struct {
	// Dashboard item identifier (UUID)
	ID *string `json:"id,omitempty"`
	// A human-readable title for the dashboard item
	Title string `json:"title"`
	// A human-readable subtitle for the dashboard item. Often a description of the visualization.
	Subtitle      string                             `json:"subtitle"`
	DataSource    DashboardItemPropertyDataSource    `json:"data_source"`
	Visualization DashboardItemPropertyVisualization `json:"visualization"`
	// The number of columns for the dashboard item to span. Dashboards are rendered on a 12-column grid on \"desktop\" screen sizes.
	Span                 *int32 `json:"span,omitempty"`
	AdditionalProperties map[string]any
}

type _DashboardItem DashboardItem

// NewDashboardItem instantiates a new DashboardItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardItem(title string, subtitle string, dataSource DashboardItemPropertyDataSource, visualization DashboardItemPropertyVisualization) *DashboardItem {
	this := DashboardItem{}
	this.Title = title
	this.Subtitle = subtitle
	this.DataSource = dataSource
	this.Visualization = visualization
	var span int32 = 4
	this.Span = &span
	return &this
}

// NewDashboardItemWithDefaults instantiates a new DashboardItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardItemWithDefaults() *DashboardItem {
	this := DashboardItem{}
	var span int32 = 4
	this.Span = &span
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *DashboardItem) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardItem) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *DashboardItem) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *DashboardItem) SetID(v string) {
	o.ID = &v
}

// GetTitle returns the Title field value
func (o *DashboardItem) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *DashboardItem) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *DashboardItem) SetTitle(v string) {
	o.Title = v
}

// GetSubtitle returns the Subtitle field value
func (o *DashboardItem) GetSubtitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value
// and a boolean to check if the value has been set.
func (o *DashboardItem) GetSubtitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtitle, true
}

// SetSubtitle sets field value
func (o *DashboardItem) SetSubtitle(v string) {
	o.Subtitle = v
}

// GetDataSource returns the DataSource field value
func (o *DashboardItem) GetDataSource() DashboardItemPropertyDataSource {
	if o == nil {
		var ret DashboardItemPropertyDataSource
		return ret
	}

	return o.DataSource
}

// GetDataSourceOk returns a tuple with the DataSource field value
// and a boolean to check if the value has been set.
func (o *DashboardItem) GetDataSourceOk() (*DashboardItemPropertyDataSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataSource, true
}

// SetDataSource sets field value
func (o *DashboardItem) SetDataSource(v DashboardItemPropertyDataSource) {
	o.DataSource = v
}

// GetVisualization returns the Visualization field value
func (o *DashboardItem) GetVisualization() DashboardItemPropertyVisualization {
	if o == nil {
		var ret DashboardItemPropertyVisualization
		return ret
	}

	return o.Visualization
}

// GetVisualizationOk returns a tuple with the Visualization field value
// and a boolean to check if the value has been set.
func (o *DashboardItem) GetVisualizationOk() (*DashboardItemPropertyVisualization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Visualization, true
}

// SetVisualization sets field value
func (o *DashboardItem) SetVisualization(v DashboardItemPropertyVisualization) {
	o.Visualization = v
}

// GetSpan returns the Span field value if set, zero value otherwise.
func (o *DashboardItem) GetSpan() int32 {
	if o == nil || o.Span == nil {
		var ret int32
		return ret
	}
	return *o.Span
}

// GetSpanOk returns a tuple with the Span field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardItem) GetSpanOk() (*int32, bool) {
	if o == nil || o.Span == nil {
		return nil, false
	}
	return o.Span, true
}

// HasSpan returns a boolean if a field has been set.
func (o *DashboardItem) HasSpan() bool {
	if o != nil && o.Span != nil {
		return true
	}

	return false
}

// SetSpan gets a reference to the given int32 and assigns it to the Span field.
func (o *DashboardItem) SetSpan(v int32) {
	o.Span = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DashboardItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["subtitle"] = o.Subtitle
	}
	if true {
		toSerialize["data_source"] = o.DataSource
	}
	if true {
		toSerialize["visualization"] = o.Visualization
	}
	if o.Span != nil {
		toSerialize["span"] = o.Span
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DashboardItem) UnmarshalJSON(bytes []byte) (err error) {
	varDashboardItem := _DashboardItem{}

	if err = json.Unmarshal(bytes, &varDashboardItem); err == nil {
		*o = DashboardItem(varDashboardItem)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "title")
		delete(additionalProperties, "subtitle")
		delete(additionalProperties, "data_source")
		delete(additionalProperties, "visualization")
		delete(additionalProperties, "span")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDashboardItem is a helper abstraction for handling nullable dashboarditem types.
type NullableDashboardItem struct {
	value *DashboardItem
	isSet bool
}

// Get returns the value.
func (v NullableDashboardItem) Get() *DashboardItem {
	return v.value
}

// Set modifies the value.
func (v *NullableDashboardItem) Set(val *DashboardItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDashboardItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDashboardItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDashboardItem returns a pointer to a new instance of NullableDashboardItem.
func NewNullableDashboardItem(val *DashboardItem) *NullableDashboardItem {
	return &NullableDashboardItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDashboardItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDashboardItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
