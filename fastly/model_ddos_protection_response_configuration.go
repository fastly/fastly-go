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

// DdosProtectionResponseConfiguration struct for DdosProtectionResponseConfiguration
type DdosProtectionResponseConfiguration struct {
	Configuration        *DdosProtectionResponseConfigurationConfiguration `json:"configuration,omitempty"`
	AdditionalProperties map[string]any
}

type _DdosProtectionResponseConfiguration DdosProtectionResponseConfiguration

// NewDdosProtectionResponseConfiguration instantiates a new DdosProtectionResponseConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosProtectionResponseConfiguration() *DdosProtectionResponseConfiguration {
	this := DdosProtectionResponseConfiguration{}
	return &this
}

// NewDdosProtectionResponseConfigurationWithDefaults instantiates a new DdosProtectionResponseConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosProtectionResponseConfigurationWithDefaults() *DdosProtectionResponseConfiguration {
	this := DdosProtectionResponseConfiguration{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *DdosProtectionResponseConfiguration) GetConfiguration() DdosProtectionResponseConfigurationConfiguration {
	if o == nil || o.Configuration == nil {
		var ret DdosProtectionResponseConfigurationConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosProtectionResponseConfiguration) GetConfigurationOk() (*DdosProtectionResponseConfigurationConfiguration, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *DdosProtectionResponseConfiguration) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given DdosProtectionResponseConfigurationConfiguration and assigns it to the Configuration field.
func (o *DdosProtectionResponseConfiguration) SetConfiguration(v DdosProtectionResponseConfigurationConfiguration) {
	o.Configuration = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DdosProtectionResponseConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DdosProtectionResponseConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varDdosProtectionResponseConfiguration := _DdosProtectionResponseConfiguration{}

	if err = json.Unmarshal(bytes, &varDdosProtectionResponseConfiguration); err == nil {
		*o = DdosProtectionResponseConfiguration(varDdosProtectionResponseConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "configuration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDdosProtectionResponseConfiguration is a helper abstraction for handling nullable ddosprotectionresponseconfiguration types.
type NullableDdosProtectionResponseConfiguration struct {
	value *DdosProtectionResponseConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableDdosProtectionResponseConfiguration) Get() *DdosProtectionResponseConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableDdosProtectionResponseConfiguration) Set(val *DdosProtectionResponseConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDdosProtectionResponseConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDdosProtectionResponseConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDdosProtectionResponseConfiguration returns a pointer to a new instance of NullableDdosProtectionResponseConfiguration.
func NewNullableDdosProtectionResponseConfiguration(val *DdosProtectionResponseConfiguration) *NullableDdosProtectionResponseConfiguration {
	return &NullableDdosProtectionResponseConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDdosProtectionResponseConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDdosProtectionResponseConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
