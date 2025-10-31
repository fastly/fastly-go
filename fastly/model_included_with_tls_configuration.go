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

// IncludedWithTlsConfiguration struct for IncludedWithTlsConfiguration
type IncludedWithTlsConfiguration struct {
	Included             []IncludedWithTlsConfigurationItem `json:"included,omitempty"`
	AdditionalProperties map[string]any
}

type _IncludedWithTlsConfiguration IncludedWithTlsConfiguration

// NewIncludedWithTlsConfiguration instantiates a new IncludedWithTlsConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncludedWithTlsConfiguration() *IncludedWithTlsConfiguration {
	this := IncludedWithTlsConfiguration{}
	return &this
}

// NewIncludedWithTlsConfigurationWithDefaults instantiates a new IncludedWithTlsConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncludedWithTlsConfigurationWithDefaults() *IncludedWithTlsConfiguration {
	this := IncludedWithTlsConfiguration{}
	return &this
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncludedWithTlsConfiguration) GetIncluded() []IncludedWithTlsConfigurationItem {
	if o == nil || o.Included == nil {
		var ret []IncludedWithTlsConfigurationItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncludedWithTlsConfiguration) GetIncludedOk() ([]IncludedWithTlsConfigurationItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncludedWithTlsConfiguration) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []IncludedWithTlsConfigurationItem and assigns it to the Included field.
func (o *IncludedWithTlsConfiguration) SetIncluded(v []IncludedWithTlsConfigurationItem) {
	o.Included = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o IncludedWithTlsConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *IncludedWithTlsConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varIncludedWithTlsConfiguration := _IncludedWithTlsConfiguration{}

	if err = json.Unmarshal(bytes, &varIncludedWithTlsConfiguration); err == nil {
		*o = IncludedWithTlsConfiguration(varIncludedWithTlsConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "included")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableIncludedWithTlsConfiguration is a helper abstraction for handling nullable includedwithtlsconfiguration types.
type NullableIncludedWithTlsConfiguration struct {
	value *IncludedWithTlsConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableIncludedWithTlsConfiguration) Get() *IncludedWithTlsConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableIncludedWithTlsConfiguration) Set(val *IncludedWithTlsConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIncludedWithTlsConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIncludedWithTlsConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncludedWithTlsConfiguration returns a pointer to a new instance of NullableIncludedWithTlsConfiguration.
func NewNullableIncludedWithTlsConfiguration(val *IncludedWithTlsConfiguration) *NullableIncludedWithTlsConfiguration {
	return &NullableIncludedWithTlsConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIncludedWithTlsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableIncludedWithTlsConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
