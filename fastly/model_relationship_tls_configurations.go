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

// RelationshipTLSConfigurations struct for RelationshipTLSConfigurations
type RelationshipTLSConfigurations struct {
	TLSConfigurations *RelationshipTLSConfigurationsTLSConfigurations `json:"tls_configurations,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSConfigurations RelationshipTLSConfigurations

// NewRelationshipTLSConfigurations instantiates a new RelationshipTLSConfigurations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSConfigurations() *RelationshipTLSConfigurations {
	this := RelationshipTLSConfigurations{}
	return &this
}

// NewRelationshipTLSConfigurationsWithDefaults instantiates a new RelationshipTLSConfigurations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSConfigurationsWithDefaults() *RelationshipTLSConfigurations {
	this := RelationshipTLSConfigurations{}
	return &this
}

// GetTLSConfigurations returns the TLSConfigurations field value if set, zero value otherwise.
func (o *RelationshipTLSConfigurations) GetTLSConfigurations() RelationshipTLSConfigurationsTLSConfigurations {
	if o == nil || o.TLSConfigurations == nil {
		var ret RelationshipTLSConfigurationsTLSConfigurations
		return ret
	}
	return *o.TLSConfigurations
}

// GetTLSConfigurationsOk returns a tuple with the TLSConfigurations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSConfigurations) GetTLSConfigurationsOk() (*RelationshipTLSConfigurationsTLSConfigurations, bool) {
	if o == nil || o.TLSConfigurations == nil {
		return nil, false
	}
	return o.TLSConfigurations, true
}

// HasTLSConfigurations returns a boolean if a field has been set.
func (o *RelationshipTLSConfigurations) HasTLSConfigurations() bool {
	if o != nil && o.TLSConfigurations != nil {
		return true
	}

	return false
}

// SetTLSConfigurations gets a reference to the given RelationshipTLSConfigurationsTLSConfigurations and assigns it to the TLSConfigurations field.
func (o *RelationshipTLSConfigurations) SetTLSConfigurations(v RelationshipTLSConfigurationsTLSConfigurations) {
	o.TLSConfigurations = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSConfigurations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSConfigurations != nil {
		toSerialize["tls_configurations"] = o.TLSConfigurations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSConfigurations) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSConfigurations := _RelationshipTLSConfigurations{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSConfigurations); err == nil {
		*o = RelationshipTLSConfigurations(varRelationshipTLSConfigurations)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_configurations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSConfigurations is a helper abstraction for handling nullable relationshiptlsconfigurations types. 
type NullableRelationshipTLSConfigurations struct {
	value *RelationshipTLSConfigurations
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSConfigurations) Get() *RelationshipTLSConfigurations {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSConfigurations) Set(val *RelationshipTLSConfigurations) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSConfigurations) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSConfigurations) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSConfigurations returns a pointer to a new instance of NullableRelationshipTLSConfigurations.
func NewNullableRelationshipTLSConfigurations(val *RelationshipTLSConfigurations) *NullableRelationshipTLSConfigurations {
	return &NullableRelationshipTLSConfigurations{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSConfigurations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSConfigurations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
