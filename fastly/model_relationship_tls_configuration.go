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

// RelationshipTlsConfiguration The [TLS configuration](https://www.fastly.com/documentation/reference/api/tls/custom-certs/configuration/) being used to terminate TLS traffic. Optional.
type RelationshipTlsConfiguration struct {
	TlsConfiguration     *RelationshipTlsConfigurationTlsConfiguration `json:"tls_configuration,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsConfiguration RelationshipTlsConfiguration

// NewRelationshipTlsConfiguration instantiates a new RelationshipTlsConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsConfiguration() *RelationshipTlsConfiguration {
	this := RelationshipTlsConfiguration{}
	return &this
}

// NewRelationshipTlsConfigurationWithDefaults instantiates a new RelationshipTlsConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsConfigurationWithDefaults() *RelationshipTlsConfiguration {
	this := RelationshipTlsConfiguration{}
	return &this
}

// GetTlsConfiguration returns the TlsConfiguration field value if set, zero value otherwise.
func (o *RelationshipTlsConfiguration) GetTlsConfiguration() RelationshipTlsConfigurationTlsConfiguration {
	if o == nil || o.TlsConfiguration == nil {
		var ret RelationshipTlsConfigurationTlsConfiguration
		return ret
	}
	return *o.TlsConfiguration
}

// GetTlsConfigurationOk returns a tuple with the TlsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsConfiguration) GetTlsConfigurationOk() (*RelationshipTlsConfigurationTlsConfiguration, bool) {
	if o == nil || o.TlsConfiguration == nil {
		return nil, false
	}
	return o.TlsConfiguration, true
}

// HasTlsConfiguration returns a boolean if a field has been set.
func (o *RelationshipTlsConfiguration) HasTlsConfiguration() bool {
	if o != nil && o.TlsConfiguration != nil {
		return true
	}

	return false
}

// SetTlsConfiguration gets a reference to the given RelationshipTlsConfigurationTlsConfiguration and assigns it to the TlsConfiguration field.
func (o *RelationshipTlsConfiguration) SetTlsConfiguration(v RelationshipTlsConfigurationTlsConfiguration) {
	o.TlsConfiguration = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsConfiguration != nil {
		toSerialize["tls_configuration"] = o.TlsConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsConfiguration := _RelationshipTlsConfiguration{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsConfiguration); err == nil {
		*o = RelationshipTlsConfiguration(varRelationshipTlsConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_configuration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsConfiguration is a helper abstraction for handling nullable relationshiptlsconfiguration types.
type NullableRelationshipTlsConfiguration struct {
	value *RelationshipTlsConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsConfiguration) Get() *RelationshipTlsConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsConfiguration) Set(val *RelationshipTlsConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsConfiguration returns a pointer to a new instance of NullableRelationshipTlsConfiguration.
func NewNullableRelationshipTlsConfiguration(val *RelationshipTlsConfiguration) *NullableRelationshipTlsConfiguration {
	return &NullableRelationshipTlsConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
