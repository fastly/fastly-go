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

// LogExplorerInsightsResponseService struct for LogExplorerInsightsResponseService
type LogExplorerInsightsResponseService struct {
	Service              *ApiDiscoveryResponseServiceService `json:"service,omitempty"`
	AdditionalProperties map[string]any
}

type _LogExplorerInsightsResponseService LogExplorerInsightsResponseService

// NewLogExplorerInsightsResponseService instantiates a new LogExplorerInsightsResponseService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogExplorerInsightsResponseService() *LogExplorerInsightsResponseService {
	this := LogExplorerInsightsResponseService{}
	return &this
}

// NewLogExplorerInsightsResponseServiceWithDefaults instantiates a new LogExplorerInsightsResponseService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogExplorerInsightsResponseServiceWithDefaults() *LogExplorerInsightsResponseService {
	this := LogExplorerInsightsResponseService{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *LogExplorerInsightsResponseService) GetService() ApiDiscoveryResponseServiceService {
	if o == nil || o.Service == nil {
		var ret ApiDiscoveryResponseServiceService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogExplorerInsightsResponseService) GetServiceOk() (*ApiDiscoveryResponseServiceService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *LogExplorerInsightsResponseService) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given ApiDiscoveryResponseServiceService and assigns it to the Service field.
func (o *LogExplorerInsightsResponseService) SetService(v ApiDiscoveryResponseServiceService) {
	o.Service = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o LogExplorerInsightsResponseService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *LogExplorerInsightsResponseService) UnmarshalJSON(bytes []byte) (err error) {
	varLogExplorerInsightsResponseService := _LogExplorerInsightsResponseService{}

	if err = json.Unmarshal(bytes, &varLogExplorerInsightsResponseService); err == nil {
		*o = LogExplorerInsightsResponseService(varLogExplorerInsightsResponseService)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableLogExplorerInsightsResponseService is a helper abstraction for handling nullable logexplorerinsightsresponseservice types.
type NullableLogExplorerInsightsResponseService struct {
	value *LogExplorerInsightsResponseService
	isSet bool
}

// Get returns the value.
func (v NullableLogExplorerInsightsResponseService) Get() *LogExplorerInsightsResponseService {
	return v.value
}

// Set modifies the value.
func (v *NullableLogExplorerInsightsResponseService) Set(val *LogExplorerInsightsResponseService) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableLogExplorerInsightsResponseService) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableLogExplorerInsightsResponseService) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableLogExplorerInsightsResponseService returns a pointer to a new instance of NullableLogExplorerInsightsResponseService.
func NewNullableLogExplorerInsightsResponseService(val *LogExplorerInsightsResponseService) *NullableLogExplorerInsightsResponseService {
	return &NullableLogExplorerInsightsResponseService{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableLogExplorerInsightsResponseService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableLogExplorerInsightsResponseService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
