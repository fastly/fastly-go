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

// LogInsightsMeta Echoes the filters that were supplied in the request.
type LogInsightsMeta struct {
	Filters              *LogInsightsMetaFilter `json:"filters,omitempty"`
	AdditionalProperties map[string]any
}

type _LogInsightsMeta LogInsightsMeta

// NewLogInsightsMeta instantiates a new LogInsightsMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogInsightsMeta() *LogInsightsMeta {
	this := LogInsightsMeta{}
	return &this
}

// NewLogInsightsMetaWithDefaults instantiates a new LogInsightsMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogInsightsMetaWithDefaults() *LogInsightsMeta {
	this := LogInsightsMeta{}
	return &this
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *LogInsightsMeta) GetFilters() LogInsightsMetaFilter {
	if o == nil || o.Filters == nil {
		var ret LogInsightsMetaFilter
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogInsightsMeta) GetFiltersOk() (*LogInsightsMetaFilter, bool) {
	if o == nil || o.Filters == nil {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *LogInsightsMeta) HasFilters() bool {
	if o != nil && o.Filters != nil {
		return true
	}

	return false
}

// SetFilters gets a reference to the given LogInsightsMetaFilter and assigns it to the Filters field.
func (o *LogInsightsMeta) SetFilters(v LogInsightsMetaFilter) {
	o.Filters = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogInsightsMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Filters != nil {
		toSerialize["filters"] = o.Filters
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LogInsightsMeta) UnmarshalJSON(bytes []byte) (err error) {
	varLogInsightsMeta := _LogInsightsMeta{}

	if err = json.Unmarshal(bytes, &varLogInsightsMeta); err == nil {
		*o = LogInsightsMeta(varLogInsightsMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "filters")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLogInsightsMeta is a helper abstraction for handling nullable loginsightsmeta types.
type NullableLogInsightsMeta struct {
	value *LogInsightsMeta
	isSet bool
}

// Get returns the value.
func (v NullableLogInsightsMeta) Get() *LogInsightsMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableLogInsightsMeta) Set(val *LogInsightsMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogInsightsMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogInsightsMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogInsightsMeta returns a pointer to a new instance of NullableLogInsightsMeta.
func NewNullableLogInsightsMeta(val *LogInsightsMeta) *NullableLogInsightsMeta {
	return &NullableLogInsightsMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogInsightsMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogInsightsMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
