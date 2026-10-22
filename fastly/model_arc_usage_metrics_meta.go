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

// ArcUsageMetricsMeta Metadata about the request.
type ArcUsageMetricsMeta struct {
	// The `sort` value used when making the request.
	Sort *string `json:"sort,omitempty"`
	// The count of results matching the request.
	Total                *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]any
}

type _ArcUsageMetricsMeta ArcUsageMetricsMeta

// NewArcUsageMetricsMeta instantiates a new ArcUsageMetricsMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcUsageMetricsMeta() *ArcUsageMetricsMeta {
	this := ArcUsageMetricsMeta{}
	return &this
}

// NewArcUsageMetricsMetaWithDefaults instantiates a new ArcUsageMetricsMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcUsageMetricsMetaWithDefaults() *ArcUsageMetricsMeta {
	this := ArcUsageMetricsMeta{}
	return &this
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *ArcUsageMetricsMeta) GetSort() string {
	if o == nil || o.Sort == nil {
		var ret string
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUsageMetricsMeta) GetSortOk() (*string, bool) {
	if o == nil || o.Sort == nil {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *ArcUsageMetricsMeta) HasSort() bool {
	if o != nil && o.Sort != nil {
		return true
	}

	return false
}

// SetSort gets a reference to the given string and assigns it to the Sort field.
func (o *ArcUsageMetricsMeta) SetSort(v string) {
	o.Sort = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ArcUsageMetricsMeta) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcUsageMetricsMeta) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ArcUsageMetricsMeta) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ArcUsageMetricsMeta) SetTotal(v int32) {
	o.Total = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcUsageMetricsMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Sort != nil {
		toSerialize["sort"] = o.Sort
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcUsageMetricsMeta) UnmarshalJSON(bytes []byte) (err error) {
	varArcUsageMetricsMeta := _ArcUsageMetricsMeta{}

	if err = json.Unmarshal(bytes, &varArcUsageMetricsMeta); err == nil {
		*o = ArcUsageMetricsMeta(varArcUsageMetricsMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "sort")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcUsageMetricsMeta is a helper abstraction for handling nullable arcusagemetricsmeta types.
type NullableArcUsageMetricsMeta struct {
	value *ArcUsageMetricsMeta
	isSet bool
}

// Get returns the value.
func (v NullableArcUsageMetricsMeta) Get() *ArcUsageMetricsMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableArcUsageMetricsMeta) Set(val *ArcUsageMetricsMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcUsageMetricsMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcUsageMetricsMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcUsageMetricsMeta returns a pointer to a new instance of NullableArcUsageMetricsMeta.
func NewNullableArcUsageMetricsMeta(val *ArcUsageMetricsMeta) *NullableArcUsageMetricsMeta {
	return &NullableArcUsageMetricsMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcUsageMetricsMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcUsageMetricsMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
