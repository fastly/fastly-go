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

// DdosProtectionRequestUpdateConfiguration struct for DdosProtectionRequestUpdateConfiguration
type DdosProtectionRequestUpdateConfiguration struct {
	// Operation mode
	Mode                 string `json:"mode"`
	AdditionalProperties map[string]any
}

type _DdosProtectionRequestUpdateConfiguration DdosProtectionRequestUpdateConfiguration

// NewDdosProtectionRequestUpdateConfiguration instantiates a new DdosProtectionRequestUpdateConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionRequestUpdateConfiguration(mode string) *DdosProtectionRequestUpdateConfiguration {
	this := DdosProtectionRequestUpdateConfiguration{}
	this.Mode = mode
	return &this
}

// NewDdosProtectionRequestUpdateConfigurationWithDefaults instantiates a new DdosProtectionRequestUpdateConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionRequestUpdateConfigurationWithDefaults() *DdosProtectionRequestUpdateConfiguration {
	this := DdosProtectionRequestUpdateConfiguration{}
	return &this
}

// GetMode returns the Mode field value
func (o *DdosProtectionRequestUpdateConfiguration) GetMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *DdosProtectionRequestUpdateConfiguration) GetModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *DdosProtectionRequestUpdateConfiguration) SetMode(v string) {
	o.Mode = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionRequestUpdateConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["mode"] = o.Mode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionRequestUpdateConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionRequestUpdateConfiguration := _DdosProtectionRequestUpdateConfiguration{}

	if err = json.Unmarshal(bytes, &varDdosProtectionRequestUpdateConfiguration); err == nil {
		*o = DdosProtectionRequestUpdateConfiguration(varDdosProtectionRequestUpdateConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "mode")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionRequestUpdateConfiguration is a helper abstraction for handling nullable ddosprotectionrequestupdateconfiguration types.
type NullableDdosProtectionRequestUpdateConfiguration struct {
	value *DdosProtectionRequestUpdateConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionRequestUpdateConfiguration) Get() *DdosProtectionRequestUpdateConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionRequestUpdateConfiguration) Set(val *DdosProtectionRequestUpdateConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionRequestUpdateConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionRequestUpdateConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionRequestUpdateConfiguration returns a pointer to a new instance of NullableDdosProtectionRequestUpdateConfiguration.
func NewNullableDdosProtectionRequestUpdateConfiguration(val *DdosProtectionRequestUpdateConfiguration) *NullableDdosProtectionRequestUpdateConfiguration {
	return &NullableDdosProtectionRequestUpdateConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionRequestUpdateConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionRequestUpdateConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
