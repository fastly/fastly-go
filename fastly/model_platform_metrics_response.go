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

// PlatformMetricsResponse struct for PlatformMetricsResponse
type PlatformMetricsResponse struct {
	Meta *PlatformMetadata `json:"meta,omitempty"`
	// An array of values representing the metric values at each point in time. Note that this dataset is sparse: only the keys with non-zero values will be included in the record.
	Data                 []PlatformValues `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _PlatformMetricsResponse PlatformMetricsResponse

// NewPlatformMetricsResponse instantiates a new PlatformMetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformMetricsResponse() *PlatformMetricsResponse {
	this := PlatformMetricsResponse{}
	return &this
}

// NewPlatformMetricsResponseWithDefaults instantiates a new PlatformMetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformMetricsResponseWithDefaults() *PlatformMetricsResponse {
	this := PlatformMetricsResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *PlatformMetricsResponse) GetMeta() PlatformMetadata {
	if o == nil || o.Meta == nil {
		var ret PlatformMetadata
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformMetricsResponse) GetMetaOk() (*PlatformMetadata, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *PlatformMetricsResponse) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PlatformMetadata and assigns it to the Meta field.
func (o *PlatformMetricsResponse) SetMeta(v PlatformMetadata) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PlatformMetricsResponse) GetData() []PlatformValues {
	if o == nil || o.Data == nil {
		var ret []PlatformValues
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformMetricsResponse) GetDataOk() ([]PlatformValues, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PlatformMetricsResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []PlatformValues and assigns it to the Data field.
func (o *PlatformMetricsResponse) SetData(v []PlatformValues) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PlatformMetricsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PlatformMetricsResponse) UnmarshalJSON(bytes []byte) (err error) {
	varPlatformMetricsResponse := _PlatformMetricsResponse{}

	if err = json.Unmarshal(bytes, &varPlatformMetricsResponse); err == nil {
		*o = PlatformMetricsResponse(varPlatformMetricsResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePlatformMetricsResponse is a helper abstraction for handling nullable platformmetricsresponse types.
type NullablePlatformMetricsResponse struct {
	value *PlatformMetricsResponse
	isSet bool
}

// Get returns the value.
func (v NullablePlatformMetricsResponse) Get() *PlatformMetricsResponse {
	return v.value
}

// Set modifies the value.
func (v *NullablePlatformMetricsResponse) Set(val *PlatformMetricsResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePlatformMetricsResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePlatformMetricsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePlatformMetricsResponse returns a pointer to a new instance of NullablePlatformMetricsResponse.
func NewNullablePlatformMetricsResponse(val *PlatformMetricsResponse) *NullablePlatformMetricsResponse {
	return &NullablePlatformMetricsResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePlatformMetricsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePlatformMetricsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
