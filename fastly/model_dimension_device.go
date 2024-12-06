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

// DimensionDevice struct for DimensionDevice
type DimensionDevice struct {
	// The client's device type for this dimension.
	Device               *string `json:"device,omitempty"`
	AdditionalProperties map[string]any
}

type _DimensionDevice DimensionDevice

// NewDimensionDevice instantiates a new DimensionDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDimensionDevice() *DimensionDevice {
	this := DimensionDevice{}
	return &this
}

// NewDimensionDeviceWithDefaults instantiates a new DimensionDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDimensionDeviceWithDefaults() *DimensionDevice {
	this := DimensionDevice{}
	return &this
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *DimensionDevice) GetDevice() string {
	if o == nil || o.Device == nil {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DimensionDevice) GetDeviceOk() (*string, bool) {
	if o == nil || o.Device == nil {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *DimensionDevice) HasDevice() bool {
	if o != nil && o.Device != nil {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *DimensionDevice) SetDevice(v string) {
	o.Device = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DimensionDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Device != nil {
		toSerialize["device"] = o.Device
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DimensionDevice) UnmarshalJSON(bytes []byte) (err error) {
	varDimensionDevice := _DimensionDevice{}

	if err = json.Unmarshal(bytes, &varDimensionDevice); err == nil {
		*o = DimensionDevice(varDimensionDevice)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "device")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDimensionDevice is a helper abstraction for handling nullable dimensiondevice types.
type NullableDimensionDevice struct {
	value *DimensionDevice
	isSet bool
}

// Get returns the value.
func (v NullableDimensionDevice) Get() *DimensionDevice {
	return v.value
}

// Set modifies the value.
func (v *NullableDimensionDevice) Set(val *DimensionDevice) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDimensionDevice) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDimensionDevice) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDimensionDevice returns a pointer to a new instance of NullableDimensionDevice.
func NewNullableDimensionDevice(val *DimensionDevice) *NullableDimensionDevice {
	return &NullableDimensionDevice{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDimensionDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDimensionDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
