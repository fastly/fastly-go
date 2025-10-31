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

// RelationshipTlsConfigurations struct for RelationshipTlsConfigurations
type RelationshipTlsConfigurations struct {
	TlsConfigurations    *RelationshipTlsConfigurationsTlsConfigurations `json:"tls_configurations,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsConfigurations RelationshipTlsConfigurations

// NewRelationshipTlsConfigurations instantiates a new RelationshipTlsConfigurations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsConfigurations() *RelationshipTlsConfigurations {
	this := RelationshipTlsConfigurations{}
	return &this
}

// NewRelationshipTlsConfigurationsWithDefaults instantiates a new RelationshipTlsConfigurations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsConfigurationsWithDefaults() *RelationshipTlsConfigurations {
	this := RelationshipTlsConfigurations{}
	return &this
}

// GetTlsConfigurations returns the TlsConfigurations field value if set, zero value otherwise.
func (o *RelationshipTlsConfigurations) GetTlsConfigurations() RelationshipTlsConfigurationsTlsConfigurations {
	if o == nil || o.TlsConfigurations == nil {
		var ret RelationshipTlsConfigurationsTlsConfigurations
		return ret
	}
	return *o.TlsConfigurations
}

// GetTlsConfigurationsOk returns a tuple with the TlsConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsConfigurations) GetTlsConfigurationsOk() (*RelationshipTlsConfigurationsTlsConfigurations, bool) {
	if o == nil || o.TlsConfigurations == nil {
		return nil, false
	}
	return o.TlsConfigurations, true
}

// HasTlsConfigurations returns a boolean if a field has been set.
func (o *RelationshipTlsConfigurations) HasTlsConfigurations() bool {
	if o != nil && o.TlsConfigurations != nil {
		return true
	}

	return false
}

// SetTlsConfigurations gets a reference to the given RelationshipTlsConfigurationsTlsConfigurations and assigns it to the TlsConfigurations field.
func (o *RelationshipTlsConfigurations) SetTlsConfigurations(v RelationshipTlsConfigurationsTlsConfigurations) {
	o.TlsConfigurations = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsConfigurations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsConfigurations != nil {
		toSerialize["tls_configurations"] = o.TlsConfigurations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsConfigurations) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsConfigurations := _RelationshipTlsConfigurations{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsConfigurations); err == nil {
		*o = RelationshipTlsConfigurations(varRelationshipTlsConfigurations)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_configurations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsConfigurations is a helper abstraction for handling nullable relationshiptlsconfigurations types.
type NullableRelationshipTlsConfigurations struct {
	value *RelationshipTlsConfigurations
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsConfigurations) Get() *RelationshipTlsConfigurations {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsConfigurations) Set(val *RelationshipTlsConfigurations) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsConfigurations) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsConfigurations) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsConfigurations returns a pointer to a new instance of NullableRelationshipTlsConfigurations.
func NewNullableRelationshipTlsConfigurations(val *RelationshipTlsConfigurations) *NullableRelationshipTlsConfigurations {
	return &NullableRelationshipTlsConfigurations{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsConfigurations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsConfigurations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
