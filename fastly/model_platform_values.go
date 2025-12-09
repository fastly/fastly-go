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
	"time"
)

// PlatformValues The results of the query, optionally filtered and grouped over the requested timespan.
type PlatformValues struct {
	// Timestamp of the metrics data point.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// 25th percentile of time to first byte from origin, in microseconds.
	TtfbOriginP25Us *float32 `json:"ttfb_origin_p25_us,omitempty"`
	// 50th percentile of time to first byte from origin, in microseconds.
	TtfbOriginP50Us *float32 `json:"ttfb_origin_p50_us,omitempty"`
	// 75th percentile of time to first byte from origin, in microseconds.
	TtfbOriginP75Us *float32 `json:"ttfb_origin_p75_us,omitempty"`
	// 95th percentile of time to first byte from origin, in microseconds.
	TtfbOriginP95Us *float32 `json:"ttfb_origin_p95_us,omitempty"`
	// 99th percentile of time to first byte from origin, in microseconds.
	TtfbOriginP99Us *float32 `json:"ttfb_origin_p99_us,omitempty"`
	// 25th percentile of time to first byte from shield, in microseconds.
	TtfbShieldP25Us *float32 `json:"ttfb_shield_p25_us,omitempty"`
	// 50th percentile of time to first byte from shield, in microseconds.
	TtfbShieldP50Us *float32 `json:"ttfb_shield_p50_us,omitempty"`
	// 75th percentile of time to first byte from shield, in microseconds.
	TtfbShieldP75Us *float32 `json:"ttfb_shield_p75_us,omitempty"`
	// 95th percentile of time to first byte from shield, in microseconds.
	TtfbShieldP95Us *float32 `json:"ttfb_shield_p95_us,omitempty"`
	// 99th percentile of time to first byte from shield, in microseconds.
	TtfbShieldP99Us *float32 `json:"ttfb_shield_p99_us,omitempty"`
	// 25th percentile of time to first byte from edge, in microseconds.
	TtfbEdgeP25Us *float32 `json:"ttfb_edge_p25_us,omitempty"`
	// 50th percentile of time to first byte from edge, in microseconds.
	TtfbEdgeP50Us *float32 `json:"ttfb_edge_p50_us,omitempty"`
	// 75th percentile of time to first byte from edge, in microseconds.
	TtfbEdgeP75Us *float32 `json:"ttfb_edge_p75_us,omitempty"`
	// 95th percentile of time to first byte from edge, in microseconds.
	TtfbEdgeP95Us *float32 `json:"ttfb_edge_p95_us,omitempty"`
	// 99th percentile of time to first byte from edge, in microseconds.
	TtfbEdgeP99Us        *float32 `json:"ttfb_edge_p99_us,omitempty"`
	AdditionalProperties map[string]any
}

type _PlatformValues PlatformValues

// NewPlatformValues instantiates a new PlatformValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformValues() *PlatformValues {
	this := PlatformValues{}
	return &this
}

// NewPlatformValuesWithDefaults instantiates a new PlatformValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformValuesWithDefaults() *PlatformValues {
	this := PlatformValues{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *PlatformValues) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *PlatformValues) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *PlatformValues) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTtfbOriginP25Us returns the TtfbOriginP25Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbOriginP25Us() float32 {
	if o == nil || o.TtfbOriginP25Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbOriginP25Us
}

// GetTtfbOriginP25UsOk returns a tuple with the TtfbOriginP25Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbOriginP25UsOk() (*float32, bool) {
	if o == nil || o.TtfbOriginP25Us == nil {
		return nil, false
	}
	return o.TtfbOriginP25Us, true
}

// HasTtfbOriginP25Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbOriginP25Us() bool {
	if o != nil && o.TtfbOriginP25Us != nil {
		return true
	}

	return false
}

// SetTtfbOriginP25Us gets a reference to the given float32 and assigns it to the TtfbOriginP25Us field.
func (o *PlatformValues) SetTtfbOriginP25Us(v float32) {
	o.TtfbOriginP25Us = &v
}

// GetTtfbOriginP50Us returns the TtfbOriginP50Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbOriginP50Us() float32 {
	if o == nil || o.TtfbOriginP50Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbOriginP50Us
}

// GetTtfbOriginP50UsOk returns a tuple with the TtfbOriginP50Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbOriginP50UsOk() (*float32, bool) {
	if o == nil || o.TtfbOriginP50Us == nil {
		return nil, false
	}
	return o.TtfbOriginP50Us, true
}

// HasTtfbOriginP50Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbOriginP50Us() bool {
	if o != nil && o.TtfbOriginP50Us != nil {
		return true
	}

	return false
}

// SetTtfbOriginP50Us gets a reference to the given float32 and assigns it to the TtfbOriginP50Us field.
func (o *PlatformValues) SetTtfbOriginP50Us(v float32) {
	o.TtfbOriginP50Us = &v
}

// GetTtfbOriginP75Us returns the TtfbOriginP75Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbOriginP75Us() float32 {
	if o == nil || o.TtfbOriginP75Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbOriginP75Us
}

// GetTtfbOriginP75UsOk returns a tuple with the TtfbOriginP75Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbOriginP75UsOk() (*float32, bool) {
	if o == nil || o.TtfbOriginP75Us == nil {
		return nil, false
	}
	return o.TtfbOriginP75Us, true
}

// HasTtfbOriginP75Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbOriginP75Us() bool {
	if o != nil && o.TtfbOriginP75Us != nil {
		return true
	}

	return false
}

// SetTtfbOriginP75Us gets a reference to the given float32 and assigns it to the TtfbOriginP75Us field.
func (o *PlatformValues) SetTtfbOriginP75Us(v float32) {
	o.TtfbOriginP75Us = &v
}

// GetTtfbOriginP95Us returns the TtfbOriginP95Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbOriginP95Us() float32 {
	if o == nil || o.TtfbOriginP95Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbOriginP95Us
}

// GetTtfbOriginP95UsOk returns a tuple with the TtfbOriginP95Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbOriginP95UsOk() (*float32, bool) {
	if o == nil || o.TtfbOriginP95Us == nil {
		return nil, false
	}
	return o.TtfbOriginP95Us, true
}

// HasTtfbOriginP95Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbOriginP95Us() bool {
	if o != nil && o.TtfbOriginP95Us != nil {
		return true
	}

	return false
}

// SetTtfbOriginP95Us gets a reference to the given float32 and assigns it to the TtfbOriginP95Us field.
func (o *PlatformValues) SetTtfbOriginP95Us(v float32) {
	o.TtfbOriginP95Us = &v
}

// GetTtfbOriginP99Us returns the TtfbOriginP99Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbOriginP99Us() float32 {
	if o == nil || o.TtfbOriginP99Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbOriginP99Us
}

// GetTtfbOriginP99UsOk returns a tuple with the TtfbOriginP99Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbOriginP99UsOk() (*float32, bool) {
	if o == nil || o.TtfbOriginP99Us == nil {
		return nil, false
	}
	return o.TtfbOriginP99Us, true
}

// HasTtfbOriginP99Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbOriginP99Us() bool {
	if o != nil && o.TtfbOriginP99Us != nil {
		return true
	}

	return false
}

// SetTtfbOriginP99Us gets a reference to the given float32 and assigns it to the TtfbOriginP99Us field.
func (o *PlatformValues) SetTtfbOriginP99Us(v float32) {
	o.TtfbOriginP99Us = &v
}

// GetTtfbShieldP25Us returns the TtfbShieldP25Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbShieldP25Us() float32 {
	if o == nil || o.TtfbShieldP25Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbShieldP25Us
}

// GetTtfbShieldP25UsOk returns a tuple with the TtfbShieldP25Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbShieldP25UsOk() (*float32, bool) {
	if o == nil || o.TtfbShieldP25Us == nil {
		return nil, false
	}
	return o.TtfbShieldP25Us, true
}

// HasTtfbShieldP25Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbShieldP25Us() bool {
	if o != nil && o.TtfbShieldP25Us != nil {
		return true
	}

	return false
}

// SetTtfbShieldP25Us gets a reference to the given float32 and assigns it to the TtfbShieldP25Us field.
func (o *PlatformValues) SetTtfbShieldP25Us(v float32) {
	o.TtfbShieldP25Us = &v
}

// GetTtfbShieldP50Us returns the TtfbShieldP50Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbShieldP50Us() float32 {
	if o == nil || o.TtfbShieldP50Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbShieldP50Us
}

// GetTtfbShieldP50UsOk returns a tuple with the TtfbShieldP50Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbShieldP50UsOk() (*float32, bool) {
	if o == nil || o.TtfbShieldP50Us == nil {
		return nil, false
	}
	return o.TtfbShieldP50Us, true
}

// HasTtfbShieldP50Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbShieldP50Us() bool {
	if o != nil && o.TtfbShieldP50Us != nil {
		return true
	}

	return false
}

// SetTtfbShieldP50Us gets a reference to the given float32 and assigns it to the TtfbShieldP50Us field.
func (o *PlatformValues) SetTtfbShieldP50Us(v float32) {
	o.TtfbShieldP50Us = &v
}

// GetTtfbShieldP75Us returns the TtfbShieldP75Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbShieldP75Us() float32 {
	if o == nil || o.TtfbShieldP75Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbShieldP75Us
}

// GetTtfbShieldP75UsOk returns a tuple with the TtfbShieldP75Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbShieldP75UsOk() (*float32, bool) {
	if o == nil || o.TtfbShieldP75Us == nil {
		return nil, false
	}
	return o.TtfbShieldP75Us, true
}

// HasTtfbShieldP75Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbShieldP75Us() bool {
	if o != nil && o.TtfbShieldP75Us != nil {
		return true
	}

	return false
}

// SetTtfbShieldP75Us gets a reference to the given float32 and assigns it to the TtfbShieldP75Us field.
func (o *PlatformValues) SetTtfbShieldP75Us(v float32) {
	o.TtfbShieldP75Us = &v
}

// GetTtfbShieldP95Us returns the TtfbShieldP95Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbShieldP95Us() float32 {
	if o == nil || o.TtfbShieldP95Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbShieldP95Us
}

// GetTtfbShieldP95UsOk returns a tuple with the TtfbShieldP95Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbShieldP95UsOk() (*float32, bool) {
	if o == nil || o.TtfbShieldP95Us == nil {
		return nil, false
	}
	return o.TtfbShieldP95Us, true
}

// HasTtfbShieldP95Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbShieldP95Us() bool {
	if o != nil && o.TtfbShieldP95Us != nil {
		return true
	}

	return false
}

// SetTtfbShieldP95Us gets a reference to the given float32 and assigns it to the TtfbShieldP95Us field.
func (o *PlatformValues) SetTtfbShieldP95Us(v float32) {
	o.TtfbShieldP95Us = &v
}

// GetTtfbShieldP99Us returns the TtfbShieldP99Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbShieldP99Us() float32 {
	if o == nil || o.TtfbShieldP99Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbShieldP99Us
}

// GetTtfbShieldP99UsOk returns a tuple with the TtfbShieldP99Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbShieldP99UsOk() (*float32, bool) {
	if o == nil || o.TtfbShieldP99Us == nil {
		return nil, false
	}
	return o.TtfbShieldP99Us, true
}

// HasTtfbShieldP99Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbShieldP99Us() bool {
	if o != nil && o.TtfbShieldP99Us != nil {
		return true
	}

	return false
}

// SetTtfbShieldP99Us gets a reference to the given float32 and assigns it to the TtfbShieldP99Us field.
func (o *PlatformValues) SetTtfbShieldP99Us(v float32) {
	o.TtfbShieldP99Us = &v
}

// GetTtfbEdgeP25Us returns the TtfbEdgeP25Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbEdgeP25Us() float32 {
	if o == nil || o.TtfbEdgeP25Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbEdgeP25Us
}

// GetTtfbEdgeP25UsOk returns a tuple with the TtfbEdgeP25Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbEdgeP25UsOk() (*float32, bool) {
	if o == nil || o.TtfbEdgeP25Us == nil {
		return nil, false
	}
	return o.TtfbEdgeP25Us, true
}

// HasTtfbEdgeP25Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbEdgeP25Us() bool {
	if o != nil && o.TtfbEdgeP25Us != nil {
		return true
	}

	return false
}

// SetTtfbEdgeP25Us gets a reference to the given float32 and assigns it to the TtfbEdgeP25Us field.
func (o *PlatformValues) SetTtfbEdgeP25Us(v float32) {
	o.TtfbEdgeP25Us = &v
}

// GetTtfbEdgeP50Us returns the TtfbEdgeP50Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbEdgeP50Us() float32 {
	if o == nil || o.TtfbEdgeP50Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbEdgeP50Us
}

// GetTtfbEdgeP50UsOk returns a tuple with the TtfbEdgeP50Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbEdgeP50UsOk() (*float32, bool) {
	if o == nil || o.TtfbEdgeP50Us == nil {
		return nil, false
	}
	return o.TtfbEdgeP50Us, true
}

// HasTtfbEdgeP50Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbEdgeP50Us() bool {
	if o != nil && o.TtfbEdgeP50Us != nil {
		return true
	}

	return false
}

// SetTtfbEdgeP50Us gets a reference to the given float32 and assigns it to the TtfbEdgeP50Us field.
func (o *PlatformValues) SetTtfbEdgeP50Us(v float32) {
	o.TtfbEdgeP50Us = &v
}

// GetTtfbEdgeP75Us returns the TtfbEdgeP75Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbEdgeP75Us() float32 {
	if o == nil || o.TtfbEdgeP75Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbEdgeP75Us
}

// GetTtfbEdgeP75UsOk returns a tuple with the TtfbEdgeP75Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbEdgeP75UsOk() (*float32, bool) {
	if o == nil || o.TtfbEdgeP75Us == nil {
		return nil, false
	}
	return o.TtfbEdgeP75Us, true
}

// HasTtfbEdgeP75Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbEdgeP75Us() bool {
	if o != nil && o.TtfbEdgeP75Us != nil {
		return true
	}

	return false
}

// SetTtfbEdgeP75Us gets a reference to the given float32 and assigns it to the TtfbEdgeP75Us field.
func (o *PlatformValues) SetTtfbEdgeP75Us(v float32) {
	o.TtfbEdgeP75Us = &v
}

// GetTtfbEdgeP95Us returns the TtfbEdgeP95Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbEdgeP95Us() float32 {
	if o == nil || o.TtfbEdgeP95Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbEdgeP95Us
}

// GetTtfbEdgeP95UsOk returns a tuple with the TtfbEdgeP95Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbEdgeP95UsOk() (*float32, bool) {
	if o == nil || o.TtfbEdgeP95Us == nil {
		return nil, false
	}
	return o.TtfbEdgeP95Us, true
}

// HasTtfbEdgeP95Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbEdgeP95Us() bool {
	if o != nil && o.TtfbEdgeP95Us != nil {
		return true
	}

	return false
}

// SetTtfbEdgeP95Us gets a reference to the given float32 and assigns it to the TtfbEdgeP95Us field.
func (o *PlatformValues) SetTtfbEdgeP95Us(v float32) {
	o.TtfbEdgeP95Us = &v
}

// GetTtfbEdgeP99Us returns the TtfbEdgeP99Us field value if set, zero value otherwise.
func (o *PlatformValues) GetTtfbEdgeP99Us() float32 {
	if o == nil || o.TtfbEdgeP99Us == nil {
		var ret float32
		return ret
	}
	return *o.TtfbEdgeP99Us
}

// GetTtfbEdgeP99UsOk returns a tuple with the TtfbEdgeP99Us field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformValues) GetTtfbEdgeP99UsOk() (*float32, bool) {
	if o == nil || o.TtfbEdgeP99Us == nil {
		return nil, false
	}
	return o.TtfbEdgeP99Us, true
}

// HasTtfbEdgeP99Us returns a boolean if a field has been set.
func (o *PlatformValues) HasTtfbEdgeP99Us() bool {
	if o != nil && o.TtfbEdgeP99Us != nil {
		return true
	}

	return false
}

// SetTtfbEdgeP99Us gets a reference to the given float32 and assigns it to the TtfbEdgeP99Us field.
func (o *PlatformValues) SetTtfbEdgeP99Us(v float32) {
	o.TtfbEdgeP99Us = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PlatformValues) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.TtfbOriginP25Us != nil {
		toSerialize["ttfb_origin_p25_us"] = o.TtfbOriginP25Us
	}
	if o.TtfbOriginP50Us != nil {
		toSerialize["ttfb_origin_p50_us"] = o.TtfbOriginP50Us
	}
	if o.TtfbOriginP75Us != nil {
		toSerialize["ttfb_origin_p75_us"] = o.TtfbOriginP75Us
	}
	if o.TtfbOriginP95Us != nil {
		toSerialize["ttfb_origin_p95_us"] = o.TtfbOriginP95Us
	}
	if o.TtfbOriginP99Us != nil {
		toSerialize["ttfb_origin_p99_us"] = o.TtfbOriginP99Us
	}
	if o.TtfbShieldP25Us != nil {
		toSerialize["ttfb_shield_p25_us"] = o.TtfbShieldP25Us
	}
	if o.TtfbShieldP50Us != nil {
		toSerialize["ttfb_shield_p50_us"] = o.TtfbShieldP50Us
	}
	if o.TtfbShieldP75Us != nil {
		toSerialize["ttfb_shield_p75_us"] = o.TtfbShieldP75Us
	}
	if o.TtfbShieldP95Us != nil {
		toSerialize["ttfb_shield_p95_us"] = o.TtfbShieldP95Us
	}
	if o.TtfbShieldP99Us != nil {
		toSerialize["ttfb_shield_p99_us"] = o.TtfbShieldP99Us
	}
	if o.TtfbEdgeP25Us != nil {
		toSerialize["ttfb_edge_p25_us"] = o.TtfbEdgeP25Us
	}
	if o.TtfbEdgeP50Us != nil {
		toSerialize["ttfb_edge_p50_us"] = o.TtfbEdgeP50Us
	}
	if o.TtfbEdgeP75Us != nil {
		toSerialize["ttfb_edge_p75_us"] = o.TtfbEdgeP75Us
	}
	if o.TtfbEdgeP95Us != nil {
		toSerialize["ttfb_edge_p95_us"] = o.TtfbEdgeP95Us
	}
	if o.TtfbEdgeP99Us != nil {
		toSerialize["ttfb_edge_p99_us"] = o.TtfbEdgeP99Us
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PlatformValues) UnmarshalJSON(bytes []byte) (err error) {
	varPlatformValues := _PlatformValues{}

	if err = json.Unmarshal(bytes, &varPlatformValues); err == nil {
		*o = PlatformValues(varPlatformValues)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "ttfb_origin_p25_us")
		delete(additionalProperties, "ttfb_origin_p50_us")
		delete(additionalProperties, "ttfb_origin_p75_us")
		delete(additionalProperties, "ttfb_origin_p95_us")
		delete(additionalProperties, "ttfb_origin_p99_us")
		delete(additionalProperties, "ttfb_shield_p25_us")
		delete(additionalProperties, "ttfb_shield_p50_us")
		delete(additionalProperties, "ttfb_shield_p75_us")
		delete(additionalProperties, "ttfb_shield_p95_us")
		delete(additionalProperties, "ttfb_shield_p99_us")
		delete(additionalProperties, "ttfb_edge_p25_us")
		delete(additionalProperties, "ttfb_edge_p50_us")
		delete(additionalProperties, "ttfb_edge_p75_us")
		delete(additionalProperties, "ttfb_edge_p95_us")
		delete(additionalProperties, "ttfb_edge_p99_us")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePlatformValues is a helper abstraction for handling nullable platformvalues types.
type NullablePlatformValues struct {
	value *PlatformValues
	isSet bool
}

// Get returns the value.
func (v NullablePlatformValues) Get() *PlatformValues {
	return v.value
}

// Set modifies the value.
func (v *NullablePlatformValues) Set(val *PlatformValues) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePlatformValues) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePlatformValues) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePlatformValues returns a pointer to a new instance of NullablePlatformValues.
func NewNullablePlatformValues(val *PlatformValues) *NullablePlatformValues {
	return &NullablePlatformValues{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePlatformValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePlatformValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
