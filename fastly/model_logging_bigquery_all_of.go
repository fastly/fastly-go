// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// LoggingBigqueryAllOf struct for LoggingBigqueryAllOf
type LoggingBigqueryAllOf struct {
	// The name of the BigQuery logging object. Used as a primary key for API access.
	Name *string `json:"name,omitempty"`
	// A Fastly [log format string](https://docs.fastly.com/en/guides/custom-log-formats). Must produce JSON that matches the schema of your BigQuery table.
	Format *string `json:"format,omitempty"`
	// Your BigQuery dataset.
	Dataset *string `json:"dataset,omitempty"`
	// Your BigQuery table.
	Table *string `json:"table,omitempty"`
	// BigQuery table name suffix template. Optional.
	TemplateSuffix NullableString `json:"template_suffix,omitempty"`
	// Your Google Cloud Platform project ID. Required
	ProjectID *string `json:"project_id,omitempty"`
	AdditionalProperties map[string]any
}

type _LoggingBigqueryAllOf LoggingBigqueryAllOf

// NewLoggingBigqueryAllOf instantiates a new LoggingBigqueryAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoggingBigqueryAllOf() *LoggingBigqueryAllOf {
	this := LoggingBigqueryAllOf{}
	return &this
}

// NewLoggingBigqueryAllOfWithDefaults instantiates a new LoggingBigqueryAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoggingBigqueryAllOfWithDefaults() *LoggingBigqueryAllOf {
	this := LoggingBigqueryAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoggingBigqueryAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingBigqueryAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoggingBigqueryAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoggingBigqueryAllOf) SetName(v string) {
	o.Name = &v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *LoggingBigqueryAllOf) GetFormat() string {
	if o == nil || o.Format == nil {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingBigqueryAllOf) GetFormatOk() (*string, bool) {
	if o == nil || o.Format == nil {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *LoggingBigqueryAllOf) HasFormat() bool {
	if o != nil && o.Format != nil {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *LoggingBigqueryAllOf) SetFormat(v string) {
	o.Format = &v
}

// GetDataset returns the Dataset field value if set, zero value otherwise.
func (o *LoggingBigqueryAllOf) GetDataset() string {
	if o == nil || o.Dataset == nil {
		var ret string
		return ret
	}
	return *o.Dataset
}

// GetDatasetOk returns a tuple with the Dataset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingBigqueryAllOf) GetDatasetOk() (*string, bool) {
	if o == nil || o.Dataset == nil {
		return nil, false
	}
	return o.Dataset, true
}

// HasDataset returns a boolean if a field has been set.
func (o *LoggingBigqueryAllOf) HasDataset() bool {
	if o != nil && o.Dataset != nil {
		return true
	}

	return false
}

// SetDataset gets a reference to the given string and assigns it to the Dataset field.
func (o *LoggingBigqueryAllOf) SetDataset(v string) {
	o.Dataset = &v
}

// GetTable returns the Table field value if set, zero value otherwise.
func (o *LoggingBigqueryAllOf) GetTable() string {
	if o == nil || o.Table == nil {
		var ret string
		return ret
	}
	return *o.Table
}

// GetTableOk returns a tuple with the Table field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingBigqueryAllOf) GetTableOk() (*string, bool) {
	if o == nil || o.Table == nil {
		return nil, false
	}
	return o.Table, true
}

// HasTable returns a boolean if a field has been set.
func (o *LoggingBigqueryAllOf) HasTable() bool {
	if o != nil && o.Table != nil {
		return true
	}

	return false
}

// SetTable gets a reference to the given string and assigns it to the Table field.
func (o *LoggingBigqueryAllOf) SetTable(v string) {
	o.Table = &v
}

// GetTemplateSuffix returns the TemplateSuffix field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoggingBigqueryAllOf) GetTemplateSuffix() string {
	if o == nil || o.TemplateSuffix.Get() == nil {
		var ret string
		return ret
	}
	return *o.TemplateSuffix.Get()
}

// GetTemplateSuffixOk returns a tuple with the TemplateSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoggingBigqueryAllOf) GetTemplateSuffixOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TemplateSuffix.Get(), o.TemplateSuffix.IsSet()
}

// HasTemplateSuffix returns a boolean if a field has been set.
func (o *LoggingBigqueryAllOf) HasTemplateSuffix() bool {
	if o != nil && o.TemplateSuffix.IsSet() {
		return true
	}

	return false
}

// SetTemplateSuffix gets a reference to the given NullableString and assigns it to the TemplateSuffix field.
func (o *LoggingBigqueryAllOf) SetTemplateSuffix(v string) {
	o.TemplateSuffix.Set(&v)
}
// SetTemplateSuffixNil sets the value for TemplateSuffix to be an explicit nil
func (o *LoggingBigqueryAllOf) SetTemplateSuffixNil() {
	o.TemplateSuffix.Set(nil)
}

// UnsetTemplateSuffix ensures that no value is present for TemplateSuffix, not even an explicit nil
func (o *LoggingBigqueryAllOf) UnsetTemplateSuffix() {
	o.TemplateSuffix.Unset()
}

// GetProjectID returns the ProjectID field value if set, zero value otherwise.
func (o *LoggingBigqueryAllOf) GetProjectID() string {
	if o == nil || o.ProjectID == nil {
		var ret string
		return ret
	}
	return *o.ProjectID
}

// GetProjectIDOk returns a tuple with the ProjectID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoggingBigqueryAllOf) GetProjectIDOk() (*string, bool) {
	if o == nil || o.ProjectID == nil {
		return nil, false
	}
	return o.ProjectID, true
}

// HasProjectID returns a boolean if a field has been set.
func (o *LoggingBigqueryAllOf) HasProjectID() bool {
	if o != nil && o.ProjectID != nil {
		return true
	}

	return false
}

// SetProjectID gets a reference to the given string and assigns it to the ProjectID field.
func (o *LoggingBigqueryAllOf) SetProjectID(v string) {
	o.ProjectID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LoggingBigqueryAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Format != nil {
		toSerialize["format"] = o.Format
	}
	if o.Dataset != nil {
		toSerialize["dataset"] = o.Dataset
	}
	if o.Table != nil {
		toSerialize["table"] = o.Table
	}
	if o.TemplateSuffix.IsSet() {
		toSerialize["template_suffix"] = o.TemplateSuffix.Get()
	}
	if o.ProjectID != nil {
		toSerialize["project_id"] = o.ProjectID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *LoggingBigqueryAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varLoggingBigqueryAllOf := _LoggingBigqueryAllOf{}

	if err = json.Unmarshal(bytes, &varLoggingBigqueryAllOf); err == nil {
		*o = LoggingBigqueryAllOf(varLoggingBigqueryAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "format")
		delete(additionalProperties, "dataset")
		delete(additionalProperties, "table")
		delete(additionalProperties, "template_suffix")
		delete(additionalProperties, "project_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLoggingBigqueryAllOf is a helper abstraction for handling nullable loggingbigqueryallof types. 
type NullableLoggingBigqueryAllOf struct {
	value *LoggingBigqueryAllOf
	isSet bool
}

// Get returns the value.
func (v NullableLoggingBigqueryAllOf) Get() *LoggingBigqueryAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableLoggingBigqueryAllOf) Set(val *LoggingBigqueryAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLoggingBigqueryAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLoggingBigqueryAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLoggingBigqueryAllOf returns a pointer to a new instance of NullableLoggingBigqueryAllOf.
func NewNullableLoggingBigqueryAllOf(val *LoggingBigqueryAllOf) *NullableLoggingBigqueryAllOf {
	return &NullableLoggingBigqueryAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLoggingBigqueryAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableLoggingBigqueryAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
