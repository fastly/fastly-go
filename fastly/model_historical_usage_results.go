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

// HistoricalUsageResults struct for HistoricalUsageResults
type HistoricalUsageResults struct {
	Bandwidth *float32 `json:"bandwidth,omitempty"`
	Requests *float32 `json:"requests,omitempty"`
	ComputeRequests *float32 `json:"compute_requests,omitempty"`
	AdditionalProperties map[string]any
}

type _HistoricalUsageResults HistoricalUsageResults

// NewHistoricalUsageResults instantiates a new HistoricalUsageResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalUsageResults() *HistoricalUsageResults {
	this := HistoricalUsageResults{}
	return &this
}

// NewHistoricalUsageResultsWithDefaults instantiates a new HistoricalUsageResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalUsageResultsWithDefaults() *HistoricalUsageResults {
	this := HistoricalUsageResults{}
	return &this
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *HistoricalUsageResults) GetBandwidth() float32 {
	if o == nil || o.Bandwidth == nil {
		var ret float32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageResults) GetBandwidthOk() (*float32, bool) {
	if o == nil || o.Bandwidth == nil {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *HistoricalUsageResults) HasBandwidth() bool {
	if o != nil && o.Bandwidth != nil {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given float32 and assigns it to the Bandwidth field.
func (o *HistoricalUsageResults) SetBandwidth(v float32) {
	o.Bandwidth = &v
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *HistoricalUsageResults) GetRequests() float32 {
	if o == nil || o.Requests == nil {
		var ret float32
		return ret
	}
	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageResults) GetRequestsOk() (*float32, bool) {
	if o == nil || o.Requests == nil {
		return nil, false
	}
	return o.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *HistoricalUsageResults) HasRequests() bool {
	if o != nil && o.Requests != nil {
		return true
	}

	return false
}

// SetRequests gets a reference to the given float32 and assigns it to the Requests field.
func (o *HistoricalUsageResults) SetRequests(v float32) {
	o.Requests = &v
}

// GetComputeRequests returns the ComputeRequests field value if set, zero value otherwise.
func (o *HistoricalUsageResults) GetComputeRequests() float32 {
	if o == nil || o.ComputeRequests == nil {
		var ret float32
		return ret
	}
	return *o.ComputeRequests
}

// GetComputeRequestsOk returns a tuple with the ComputeRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HistoricalUsageResults) GetComputeRequestsOk() (*float32, bool) {
	if o == nil || o.ComputeRequests == nil {
		return nil, false
	}
	return o.ComputeRequests, true
}

// HasComputeRequests returns a boolean if a field has been set.
func (o *HistoricalUsageResults) HasComputeRequests() bool {
	if o != nil && o.ComputeRequests != nil {
		return true
	}

	return false
}

// SetComputeRequests gets a reference to the given float32 and assigns it to the ComputeRequests field.
func (o *HistoricalUsageResults) SetComputeRequests(v float32) {
	o.ComputeRequests = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o HistoricalUsageResults) MarshalJSON() ([]byte, error) {
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
func (o *HistoricalUsageResults) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalUsageResults := _HistoricalUsageResults{}

	if err = json.Unmarshal(bytes, &varHistoricalUsageResults); err == nil {
		*o = HistoricalUsageResults(varHistoricalUsageResults)
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

// NullableHistoricalUsageResults is a helper abstraction for handling nullable historicalusageresults types. 
type NullableHistoricalUsageResults struct {
	value *HistoricalUsageResults
	isSet bool
}

// Get returns the value.
func (v NullableHistoricalUsageResults) Get() *HistoricalUsageResults {
	return v.value
}

// Set modifies the value.
func (v *NullableHistoricalUsageResults) Set(val *HistoricalUsageResults) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableHistoricalUsageResults) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableHistoricalUsageResults) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableHistoricalUsageResults returns a pointer to a new instance of NullableHistoricalUsageResults.
func NewNullableHistoricalUsageResults(val *HistoricalUsageResults) *NullableHistoricalUsageResults {
	return &NullableHistoricalUsageResults{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableHistoricalUsageResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableHistoricalUsageResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
