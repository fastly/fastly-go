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

// ImageOptimizerResponseService struct for ImageOptimizerResponseService
type ImageOptimizerResponseService struct {
	Service              *ApiDiscoveryResponseServiceService `json:"service,omitempty"`
	AdditionalProperties map[string]any
}

type _ImageOptimizerResponseService ImageOptimizerResponseService

// NewImageOptimizerResponseService instantiates a new ImageOptimizerResponseService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageOptimizerResponseService() *ImageOptimizerResponseService {
	this := ImageOptimizerResponseService{}
	return &this
}

// NewImageOptimizerResponseServiceWithDefaults instantiates a new ImageOptimizerResponseService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageOptimizerResponseServiceWithDefaults() *ImageOptimizerResponseService {
	this := ImageOptimizerResponseService{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ImageOptimizerResponseService) GetService() ApiDiscoveryResponseServiceService {
	if o == nil || o.Service == nil {
		var ret ApiDiscoveryResponseServiceService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageOptimizerResponseService) GetServiceOk() (*ApiDiscoveryResponseServiceService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ImageOptimizerResponseService) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given ApiDiscoveryResponseServiceService and assigns it to the Service field.
func (o *ImageOptimizerResponseService) SetService(v ApiDiscoveryResponseServiceService) {
	o.Service = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ImageOptimizerResponseService) MarshalJSON() ([]byte, error) {
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
func (o *ImageOptimizerResponseService) UnmarshalJSON(bytes []byte) (err error) {
	varImageOptimizerResponseService := _ImageOptimizerResponseService{}

	if err = json.Unmarshal(bytes, &varImageOptimizerResponseService); err == nil {
		*o = ImageOptimizerResponseService(varImageOptimizerResponseService)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableImageOptimizerResponseService is a helper abstraction for handling nullable imageoptimizerresponseservice types.
type NullableImageOptimizerResponseService struct {
	value *ImageOptimizerResponseService
	isSet bool
}

// Get returns the value.
func (v NullableImageOptimizerResponseService) Get() *ImageOptimizerResponseService {
	return v.value
}

// Set modifies the value.
func (v *NullableImageOptimizerResponseService) Set(val *ImageOptimizerResponseService) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableImageOptimizerResponseService) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableImageOptimizerResponseService) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableImageOptimizerResponseService returns a pointer to a new instance of NullableImageOptimizerResponseService.
func NewNullableImageOptimizerResponseService(val *ImageOptimizerResponseService) *NullableImageOptimizerResponseService {
	return &NullableImageOptimizerResponseService{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableImageOptimizerResponseService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableImageOptimizerResponseService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
