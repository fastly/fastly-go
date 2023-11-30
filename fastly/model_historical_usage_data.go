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

// HistoricalUsageData The results of usage related queries, grouped by service and/or region depending on endpoint, and aggregated over the appropriate time span.
type HistoricalUsageData struct {
	Bandwidth *float32 `json:"bandwidth,omitempty"`
	Requests *float32 `json:"requests,omitempty"`
	ComputeRequests *float32 `json:"compute_requests,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalUsageData HistoricalUsageData

// NewHistoricalUsageData instantiates a new HistoricalUsageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalUsageData() *HistoricalUsageData {
	this := HistoricalUsageData{}
	return &this
}

// NewHistoricalUsageDataWithDefaults instantiates a new HistoricalUsageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalUsageDataWithDefaults() *HistoricalUsageData {
	this := HistoricalUsageData{}
	return &this
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *HistoricalUsageData) GetBandwidth() float32 {
	if o == nil || o.Bandwidth == nil {
		var ret float32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageData) GetBandwidthOk() (*float32, bool) {
	if o == nil || o.Bandwidth == nil {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *HistoricalUsageData) HasBandwidth() bool {
	if o != nil && o.Bandwidth != nil {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given float32 and assigns it to the Bandwidth field.
func (o *HistoricalUsageData) SetBandwidth(v float32) {
	o.Bandwidth = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *HistoricalUsageData) GetRequests() float32 {
	if o == nil || o.Requests == nil {
		var ret float32
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageData) GetRequestsOk() (*float32, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *HistoricalUsageData) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given float32 and assigns it to the Requests field.
func (o *HistoricalUsageData) SetRequests(v float32) {
	o.Requests = &v
}

// GetComputeRequests returns the ComputeRequests field value if set, zero value otherwise.
func (o *HistoricalUsageData) GetComputeRequests() float32 {
	if o == nil || o.ComputeRequests == nil {
		var ret float32
		return ret
	}
	return *o.ComputeRequests
}

// GetComputeRequestsOk returns a tuple with the ComputeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageData) GetComputeRequestsOk() (*float32, bool) {
	if o == nil || o.ComputeRequests == nil {
		return nil, false
	}
	return o.ComputeRequests, true
}

// HasComputeRequests returns a boolean if a field has been set.
func (o *HistoricalUsageData) HasComputeRequests() bool {
	if o != nil && o.ComputeRequests != nil {
		return true
	}

	return false
}

// SetComputeRequests gets a reference to the given float32 and assigns it to the ComputeRequests field.
func (o *HistoricalUsageData) SetComputeRequests(v float32) {
	o.ComputeRequests = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalUsageData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Bandwidth != nil {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if o.Requests != nil {
		toSerialize["requests"] = o.Requests
	}
	if o.ComputeRequests != nil {
		toSerialize["compute_requests"] = o.ComputeRequests
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *HistoricalUsageData) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalUsageData := _HistoricalUsageData{}

	if err = json.Unmarshal(bytes, &varHistoricalUsageData); err == nil {
		*o = HistoricalUsageData(varHistoricalUsageData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "bandwidth")
		delete(additionalProperties, "requests")
		delete(additionalProperties, "compute_requests")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableHistoricalUsageData is a helper abstraction for handling nullable historicalusagedata types. 
type NullableHistoricalUsageData struct {
	value *HistoricalUsageData
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalUsageData) Get() *HistoricalUsageData {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalUsageData) Set(val *HistoricalUsageData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalUsageData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalUsageData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalUsageData returns a pointer to a new instance of NullableHistoricalUsageData.
func NewNullableHistoricalUsageData(val *HistoricalUsageData) *NullableHistoricalUsageData {
	return &NullableHistoricalUsageData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalUsageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHistoricalUsageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
