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

// RelationshipTLSConfiguration The [TLS configuration](https://www.fastly.com/documentation/reference/api/tls/custom-certs/configuration/) being used to terminate TLS traffic. Optional.
type RelationshipTLSConfiguration struct {
	TLSConfiguration *RelationshipTLSConfigurationTLSConfiguration `json:"tls_configuration,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSConfiguration RelationshipTLSConfiguration

// NewRelationshipTLSConfiguration instantiates a new RelationshipTLSConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSConfiguration() *RelationshipTLSConfiguration {
	this := RelationshipTLSConfiguration{}
	return &this
}

// NewRelationshipTLSConfigurationWithDefaults instantiates a new RelationshipTLSConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSConfigurationWithDefaults() *RelationshipTLSConfiguration {
	this := RelationshipTLSConfiguration{}
	return &this
}

// GetTLSConfiguration returns the TLSConfiguration field value if set, zero value otherwise.
func (o *RelationshipTLSConfiguration) GetTLSConfiguration() RelationshipTLSConfigurationTLSConfiguration {
	if o == nil || o.TLSConfiguration == nil {
		var ret RelationshipTLSConfigurationTLSConfiguration
		return ret
	}
	return *o.TLSConfiguration
}

// GetTLSConfigurationOk returns a tuple with the TLSConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSConfiguration) GetTLSConfigurationOk() (*RelationshipTLSConfigurationTLSConfiguration, bool) {
	if o == nil || o.TLSConfiguration == nil {
		return nil, false
	}
	return o.TLSConfiguration, true
}

// HasTLSConfiguration returns a boolean if a field has been set.
func (o *RelationshipTLSConfiguration) HasTLSConfiguration() bool {
	if o != nil && o.TLSConfiguration != nil {
		return true
	}

	return false
}

// SetTLSConfiguration gets a reference to the given RelationshipTLSConfigurationTLSConfiguration and assigns it to the TLSConfiguration field.
func (o *RelationshipTLSConfiguration) SetTLSConfiguration(v RelationshipTLSConfigurationTLSConfiguration) {
	o.TLSConfiguration = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSConfiguration != nil {
		toSerialize["tls_configuration"] = o.TLSConfiguration
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSConfiguration) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSConfiguration := _RelationshipTLSConfiguration{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSConfiguration); err == nil {
		*o = RelationshipTLSConfiguration(varRelationshipTLSConfiguration)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_configuration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSConfiguration is a helper abstraction for handling nullable relationshiptlsconfiguration types. 
type NullableRelationshipTLSConfiguration struct {
	value *RelationshipTLSConfiguration
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSConfiguration) Get() *RelationshipTLSConfiguration {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSConfiguration) Set(val *RelationshipTLSConfiguration) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSConfiguration) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSConfiguration returns a pointer to a new instance of NullableRelationshipTLSConfiguration.
func NewNullableRelationshipTLSConfiguration(val *RelationshipTLSConfiguration) *NullableRelationshipTLSConfiguration {
	return &NullableRelationshipTLSConfiguration{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
