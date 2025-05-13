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

// OriginInspectorResponseEnabledServices struct for OriginInspectorResponseEnabledServices
type OriginInspectorResponseEnabledServices struct {
	// A list of services with Origin Inspector enabled.
	Services             []string `json:"services,omitempty"`
	AdditionalProperties map[string]any
}

type _OriginInspectorResponseEnabledServices OriginInspectorResponseEnabledServices

// NewOriginInspectorResponseEnabledServices instantiates a new OriginInspectorResponseEnabledServices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginInspectorResponseEnabledServices() *OriginInspectorResponseEnabledServices {
	this := OriginInspectorResponseEnabledServices{}
	return &this
}

// NewOriginInspectorResponseEnabledServicesWithDefaults instantiates a new OriginInspectorResponseEnabledServices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginInspectorResponseEnabledServicesWithDefaults() *OriginInspectorResponseEnabledServices {
	this := OriginInspectorResponseEnabledServices{}
	return &this
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *OriginInspectorResponseEnabledServices) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginInspectorResponseEnabledServices) GetServicesOk() ([]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *OriginInspectorResponseEnabledServices) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *OriginInspectorResponseEnabledServices) SetServices(v []string) {
	o.Services = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OriginInspectorResponseEnabledServices) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *OriginInspectorResponseEnabledServices) UnmarshalJSON(bytes []byte) (err error) {
	varOriginInspectorResponseEnabledServices := _OriginInspectorResponseEnabledServices{}

	if err = json.Unmarshal(bytes, &varOriginInspectorResponseEnabledServices); err == nil {
		*o = OriginInspectorResponseEnabledServices(varOriginInspectorResponseEnabledServices)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "services")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOriginInspectorResponseEnabledServices is a helper abstraction for handling nullable origininspectorresponseenabledservices types.
type NullableOriginInspectorResponseEnabledServices struct {
	value *OriginInspectorResponseEnabledServices
	isSet bool
}

// Get returns the value.
func (v NullableOriginInspectorResponseEnabledServices) Get() *OriginInspectorResponseEnabledServices {
	return v.value
}

// Set modifies the value.
func (v *NullableOriginInspectorResponseEnabledServices) Set(val *OriginInspectorResponseEnabledServices) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOriginInspectorResponseEnabledServices) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOriginInspectorResponseEnabledServices) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOriginInspectorResponseEnabledServices returns a pointer to a new instance of NullableOriginInspectorResponseEnabledServices.
func NewNullableOriginInspectorResponseEnabledServices(val *OriginInspectorResponseEnabledServices) *NullableOriginInspectorResponseEnabledServices {
	return &NullableOriginInspectorResponseEnabledServices{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOriginInspectorResponseEnabledServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOriginInspectorResponseEnabledServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
