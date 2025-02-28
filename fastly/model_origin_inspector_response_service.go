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

// OriginInspectorResponseService struct for OriginInspectorResponseService
type OriginInspectorResponseService struct {
	Service              *BotManagementResponseServiceService `json:"service,omitempty"`
	AdditionalProperties map[string]any
}

type _OriginInspectorResponseService OriginInspectorResponseService

// NewOriginInspectorResponseService instantiates a new OriginInspectorResponseService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspectorResponseService() *OriginInspectorResponseService {
	this := OriginInspectorResponseService{}
	return &this
}

// NewOriginInspectorResponseServiceWithDefaults instantiates a new OriginInspectorResponseService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorResponseServiceWithDefaults() *OriginInspectorResponseService {
	this := OriginInspectorResponseService{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *OriginInspectorResponseService) GetService() BotManagementResponseServiceService {
	if o == nil || o.Service == nil {
		var ret BotManagementResponseServiceService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorResponseService) GetServiceOk() (*BotManagementResponseServiceService, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *OriginInspectorResponseService) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given BotManagementResponseServiceService and assigns it to the Service field.
func (o *OriginInspectorResponseService) SetService(v BotManagementResponseServiceService) {
	o.Service = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspectorResponseService) MarshalJSON() ([]byte, error) {
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
func (o *OriginInspectorResponseService) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspectorResponseService := _OriginInspectorResponseService{}

	if err = json.Unmarshal(bytes, &varOriginInspectorResponseService); err == nil {
		*o = OriginInspectorResponseService(varOriginInspectorResponseService)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "service")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOriginInspectorResponseService is a helper abstraction for handling nullable origininspectorresponseservice types.
type NullableOriginInspectorResponseService struct {
	value *OriginInspectorResponseService
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspectorResponseService) Get() *OriginInspectorResponseService {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspectorResponseService) Set(val *OriginInspectorResponseService) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspectorResponseService) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspectorResponseService) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspectorResponseService returns a pointer to a new instance of NullableOriginInspectorResponseService.
func NewNullableOriginInspectorResponseService(val *OriginInspectorResponseService) *NullableOriginInspectorResponseService {
	return &NullableOriginInspectorResponseService{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspectorResponseService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOriginInspectorResponseService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
