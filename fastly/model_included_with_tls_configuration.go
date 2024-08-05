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

// IncludedWithTLSConfiguration struct for IncludedWithTLSConfiguration
type IncludedWithTLSConfiguration struct {
	Included []IncludedWithTLSConfigurationItem `json:"included,omitempty"`
	AdditionalProperties map[string]any
}

type _IncludedWithTLSConfiguration IncludedWithTLSConfiguration

// NewIncludedWithTLSConfiguration instantiates a new IncludedWithTLSConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncludedWithTLSConfiguration() *IncludedWithTLSConfiguration {
	this := IncludedWithTLSConfiguration{}
	return &this
}

// NewIncludedWithTLSConfigurationWithDefaults instantiates a new IncludedWithTLSConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncludedWithTLSConfigurationWithDefaults() *IncludedWithTLSConfiguration {
	this := IncludedWithTLSConfiguration{}
	return &this
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *IncludedWithTLSConfiguration) GetIncluded() []IncludedWithTLSConfigurationItem {
	if o == nil || o.Included == nil {
		var ret []IncludedWithTLSConfigurationItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncludedWithTLSConfiguration) GetIncludedOk() ([]IncludedWithTLSConfigurationItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *IncludedWithTLSConfiguration) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []IncludedWithTLSConfigurationItem and assigns it to the Included field.
func (o *IncludedWithTLSConfiguration) SetIncluded(v []IncludedWithTLSConfigurationItem) {
	o.Included = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o IncludedWithTLSConfiguration) MarshalJSON() ([]byte, error) {
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
func (o *IncludedWithTLSConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varIncludedWithTLSConfiguration := _IncludedWithTLSConfiguration{}

	if err = json.Unmarshal(bytes, &varIncludedWithTLSConfiguration); err == nil {
		*o = IncludedWithTLSConfiguration(varIncludedWithTLSConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "included")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableIncludedWithTLSConfiguration is a helper abstraction for handling nullable includedwithtlsconfiguration types. 
type NullableIncludedWithTLSConfiguration struct {
	value *IncludedWithTLSConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableIncludedWithTLSConfiguration) Get() *IncludedWithTLSConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableIncludedWithTLSConfiguration) Set(val *IncludedWithTLSConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIncludedWithTLSConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIncludedWithTLSConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncludedWithTLSConfiguration returns a pointer to a new instance of NullableIncludedWithTLSConfiguration.
func NewNullableIncludedWithTLSConfiguration(val *IncludedWithTLSConfiguration) *NullableIncludedWithTLSConfiguration {
	return &NullableIncludedWithTLSConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIncludedWithTLSConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableIncludedWithTLSConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
