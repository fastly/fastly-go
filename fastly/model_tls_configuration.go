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

// TLSConfiguration struct for TLSConfiguration
type TLSConfiguration struct {
	Data *TLSConfigurationData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _TLSConfiguration TLSConfiguration

// NewTLSConfiguration instantiates a new TLSConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTLSConfiguration() *TLSConfiguration {
	this := TLSConfiguration{}
	return &this
}

// NewTLSConfigurationWithDefaults instantiates a new TLSConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTLSConfigurationWithDefaults() *TLSConfiguration {
	this := TLSConfiguration{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *TLSConfiguration) GetData() TLSConfigurationData {
	if o == nil || o.Data == nil {
		var ret TLSConfigurationData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TLSConfiguration) GetDataOk() (*TLSConfigurationData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *TLSConfiguration) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given TLSConfigurationData and assigns it to the Data field.
func (o *TLSConfiguration) SetData(v TLSConfigurationData) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o TLSConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *TLSConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varTLSConfiguration := _TLSConfiguration{}

	if err = json.Unmarshal(bytes, &varTLSConfiguration); err == nil {
		*o = TLSConfiguration(varTLSConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableTLSConfiguration is a helper abstraction for handling nullable tlsconfiguration types. 
type NullableTLSConfiguration struct {
	value *TLSConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableTLSConfiguration) Get() *TLSConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableTLSConfiguration) Set(val *TLSConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableTLSConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableTLSConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableTLSConfiguration returns a pointer to a new instance of NullableTLSConfiguration.
func NewNullableTLSConfiguration(val *TLSConfiguration) *NullableTLSConfiguration {
	return &NullableTLSConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableTLSConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableTLSConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
