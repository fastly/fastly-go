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

// HistoricalDdosMeta Meta information about the scope of the query in a human readable format.
type HistoricalDdosMeta struct {
	// Start time that was used to perform the query as an ISO-8601-formatted date and time.
	Start *string `json:"start,omitempty"`
	// End time that was used to perform the query as an ISO-8601-formatted date and time.
	End *string `json:"end,omitempty"`
	// Downsample that was used to perform the query. One of `hour` or `day`.
	Downsample *string `json:"downsample,omitempty"`
	// A comma-separated list of the metrics that were requested.
	Metric               *string `json:"metric,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalDdosMeta HistoricalDdosMeta

// NewHistoricalDdosMeta instantiates a new HistoricalDdosMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalDdosMeta() *HistoricalDdosMeta {
	this := HistoricalDdosMeta{}
	return &this
}

// NewHistoricalDdosMetaWithDefaults instantiates a new HistoricalDdosMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalDdosMetaWithDefaults() *HistoricalDdosMeta {
	this := HistoricalDdosMeta{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *HistoricalDdosMeta) GetStart() string {
	if o == nil || o.Start == nil {
		var ret string
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDdosMeta) GetStartOk() (*string, bool) {
	if o == nil || o.Start == nil {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *HistoricalDdosMeta) HasStart() bool {
	if o != nil && o.Start != nil {
		return true
	}

	return false
}

// SetStart gets a reference to the given string and assigns it to the Start field.
func (o *HistoricalDdosMeta) SetStart(v string) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *HistoricalDdosMeta) GetEnd() string {
	if o == nil || o.End == nil {
		var ret string
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDdosMeta) GetEndOk() (*string, bool) {
	if o == nil || o.End == nil {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *HistoricalDdosMeta) HasEnd() bool {
	if o != nil && o.End != nil {
		return true
	}

	return false
}

// SetEnd gets a reference to the given string and assigns it to the End field.
func (o *HistoricalDdosMeta) SetEnd(v string) {
	o.End = &v
}

// GetDownsample returns the Downsample field value if set, zero value otherwise.
func (o *HistoricalDdosMeta) GetDownsample() string {
	if o == nil || o.Downsample == nil {
		var ret string
		return ret
	}
	return *o.Downsample
}

// GetDownsampleOk returns a tuple with the Downsample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDdosMeta) GetDownsampleOk() (*string, bool) {
	if o == nil || o.Downsample == nil {
		return nil, false
	}
	return o.Downsample, true
}

// HasDownsample returns a boolean if a field has been set.
func (o *HistoricalDdosMeta) HasDownsample() bool {
	if o != nil && o.Downsample != nil {
		return true
	}

	return false
}

// SetDownsample gets a reference to the given string and assigns it to the Downsample field.
func (o *HistoricalDdosMeta) SetDownsample(v string) {
	o.Downsample = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *HistoricalDdosMeta) GetMetric() string {
	if o == nil || o.Metric == nil {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalDdosMeta) GetMetricOk() (*string, bool) {
	if o == nil || o.Metric == nil {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *HistoricalDdosMeta) HasMetric() bool {
	if o != nil && o.Metric != nil {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *HistoricalDdosMeta) SetMetric(v string) {
	o.Metric = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalDdosMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Start != nil {
		toSerialize["start"] = o.Start
	}
	if o.End != nil {
		toSerialize["end"] = o.End
	}
	if o.Downsample != nil {
		toSerialize["downsample"] = o.Downsample
	}
	if o.Metric != nil {
		toSerialize["metric"] = o.Metric
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *HistoricalDdosMeta) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalDdosMeta := _HistoricalDdosMeta{}

	if err = json.Unmarshal(bytes, &varHistoricalDdosMeta); err == nil {
		*o = HistoricalDdosMeta(varHistoricalDdosMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "start")
		delete(additionalProperties, "end")
		delete(additionalProperties, "downsample")
		delete(additionalProperties, "metric")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHistoricalDdosMeta is a helper abstraction for handling nullable historicalddosmeta types.
type NullableHistoricalDdosMeta struct {
	value *HistoricalDdosMeta
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalDdosMeta) Get() *HistoricalDdosMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalDdosMeta) Set(val *HistoricalDdosMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalDdosMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalDdosMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalDdosMeta returns a pointer to a new instance of NullableHistoricalDdosMeta.
func NewNullableHistoricalDdosMeta(val *HistoricalDdosMeta) *NullableHistoricalDdosMeta {
	return &NullableHistoricalDdosMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalDdosMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableHistoricalDdosMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
