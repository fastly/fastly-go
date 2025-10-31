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

// RelationshipTlsActivations struct for RelationshipTlsActivations
type RelationshipTlsActivations struct {
	TlsActivations       *RelationshipTlsActivationTlsActivation `json:"tls_activations,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsActivations RelationshipTlsActivations

// NewRelationshipTlsActivations instantiates a new RelationshipTlsActivations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsActivations() *RelationshipTlsActivations {
	this := RelationshipTlsActivations{}
	return &this
}

// NewRelationshipTlsActivationsWithDefaults instantiates a new RelationshipTlsActivations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsActivationsWithDefaults() *RelationshipTlsActivations {
	this := RelationshipTlsActivations{}
	return &this
}

// GetTlsActivations returns the TlsActivations field value if set, zero value otherwise.
func (o *RelationshipTlsActivations) GetTlsActivations() RelationshipTlsActivationTlsActivation {
	if o == nil || o.TlsActivations == nil {
		var ret RelationshipTlsActivationTlsActivation
		return ret
	}
	return *o.TlsActivations
}

// GetTlsActivationsOk returns a tuple with the TlsActivations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsActivations) GetTlsActivationsOk() (*RelationshipTlsActivationTlsActivation, bool) {
	if o == nil || o.TlsActivations == nil {
		return nil, false
	}
	return o.TlsActivations, true
}

// HasTlsActivations returns a boolean if a field has been set.
func (o *RelationshipTlsActivations) HasTlsActivations() bool {
	if o != nil && o.TlsActivations != nil {
		return true
	}

	return false
}

// SetTlsActivations gets a reference to the given RelationshipTlsActivationTlsActivation and assigns it to the TlsActivations field.
func (o *RelationshipTlsActivations) SetTlsActivations(v RelationshipTlsActivationTlsActivation) {
	o.TlsActivations = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsActivations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsActivations != nil {
		toSerialize["tls_activations"] = o.TlsActivations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsActivations) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsActivations := _RelationshipTlsActivations{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsActivations); err == nil {
		*o = RelationshipTlsActivations(varRelationshipTlsActivations)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_activations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsActivations is a helper abstraction for handling nullable relationshiptlsactivations types.
type NullableRelationshipTlsActivations struct {
	value *RelationshipTlsActivations
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsActivations) Get() *RelationshipTlsActivations {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsActivations) Set(val *RelationshipTlsActivations) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsActivations) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsActivations) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsActivations returns a pointer to a new instance of NullableRelationshipTlsActivations.
func NewNullableRelationshipTlsActivations(val *RelationshipTlsActivations) *NullableRelationshipTlsActivations {
	return &NullableRelationshipTlsActivations{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsActivations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsActivations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
