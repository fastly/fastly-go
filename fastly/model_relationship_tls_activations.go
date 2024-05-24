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

// RelationshipTLSActivations struct for RelationshipTLSActivations
type RelationshipTLSActivations struct {
	TLSActivations *RelationshipTLSActivationTLSActivation `json:"tls_activations,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSActivations RelationshipTLSActivations

// NewRelationshipTLSActivations instantiates a new RelationshipTLSActivations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSActivations() *RelationshipTLSActivations {
	this := RelationshipTLSActivations{}
	return &this
}

// NewRelationshipTLSActivationsWithDefaults instantiates a new RelationshipTLSActivations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSActivationsWithDefaults() *RelationshipTLSActivations {
	this := RelationshipTLSActivations{}
	return &this
}

// GetTLSActivations returns the TLSActivations field value if set, zero value otherwise.
func (o *RelationshipTLSActivations) GetTLSActivations() RelationshipTLSActivationTLSActivation {
	if o == nil || o.TLSActivations == nil {
		var ret RelationshipTLSActivationTLSActivation
		return ret
	}
	return *o.TLSActivations
}

// GetTLSActivationsOk returns a tuple with the TLSActivations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSActivations) GetTLSActivationsOk() (*RelationshipTLSActivationTLSActivation, bool) {
	if o == nil || o.TLSActivations == nil {
		return nil, false
	}
	return o.TLSActivations, true
}

// HasTLSActivations returns a boolean if a field has been set.
func (o *RelationshipTLSActivations) HasTLSActivations() bool {
	if o != nil && o.TLSActivations != nil {
		return true
	}

	return false
}

// SetTLSActivations gets a reference to the given RelationshipTLSActivationTLSActivation and assigns it to the TLSActivations field.
func (o *RelationshipTLSActivations) SetTLSActivations(v RelationshipTLSActivationTLSActivation) {
	o.TLSActivations = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSActivations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSActivations != nil {
		toSerialize["tls_activations"] = o.TLSActivations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSActivations) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSActivations := _RelationshipTLSActivations{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSActivations); err == nil {
		*o = RelationshipTLSActivations(varRelationshipTLSActivations)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_activations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSActivations is a helper abstraction for handling nullable relationshiptlsactivations types. 
type NullableRelationshipTLSActivations struct {
	value *RelationshipTLSActivations
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSActivations) Get() *RelationshipTLSActivations {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSActivations) Set(val *RelationshipTLSActivations) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSActivations) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSActivations) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSActivations returns a pointer to a new instance of NullableRelationshipTLSActivations.
func NewNullableRelationshipTLSActivations(val *RelationshipTLSActivations) *NullableRelationshipTLSActivations {
	return &NullableRelationshipTLSActivations{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSActivations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSActivations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
