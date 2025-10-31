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

// RelationshipTlsActivation struct for RelationshipTlsActivation
type RelationshipTlsActivation struct {
	TlsActivation        *RelationshipTlsActivationTlsActivation `json:"tls_activation,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsActivation RelationshipTlsActivation

// NewRelationshipTlsActivation instantiates a new RelationshipTlsActivation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsActivation() *RelationshipTlsActivation {
	this := RelationshipTlsActivation{}
	return &this
}

// NewRelationshipTlsActivationWithDefaults instantiates a new RelationshipTlsActivation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsActivationWithDefaults() *RelationshipTlsActivation {
	this := RelationshipTlsActivation{}
	return &this
}

// GetTlsActivation returns the TlsActivation field value if set, zero value otherwise.
func (o *RelationshipTlsActivation) GetTlsActivation() RelationshipTlsActivationTlsActivation {
	if o == nil || o.TlsActivation == nil {
		var ret RelationshipTlsActivationTlsActivation
		return ret
	}
	return *o.TlsActivation
}

// GetTlsActivationOk returns a tuple with the TlsActivation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsActivation) GetTlsActivationOk() (*RelationshipTlsActivationTlsActivation, bool) {
	if o == nil || o.TlsActivation == nil {
		return nil, false
	}
	return o.TlsActivation, true
}

// HasTlsActivation returns a boolean if a field has been set.
func (o *RelationshipTlsActivation) HasTlsActivation() bool {
	if o != nil && o.TlsActivation != nil {
		return true
	}

	return false
}

// SetTlsActivation gets a reference to the given RelationshipTlsActivationTlsActivation and assigns it to the TlsActivation field.
func (o *RelationshipTlsActivation) SetTlsActivation(v RelationshipTlsActivationTlsActivation) {
	o.TlsActivation = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsActivation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsActivation != nil {
		toSerialize["tls_activation"] = o.TlsActivation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsActivation) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsActivation := _RelationshipTlsActivation{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsActivation); err == nil {
		*o = RelationshipTlsActivation(varRelationshipTlsActivation)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_activation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsActivation is a helper abstraction for handling nullable relationshiptlsactivation types.
type NullableRelationshipTlsActivation struct {
	value *RelationshipTlsActivation
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsActivation) Get() *RelationshipTlsActivation {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsActivation) Set(val *RelationshipTlsActivation) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsActivation) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsActivation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsActivation returns a pointer to a new instance of NullableRelationshipTlsActivation.
func NewNullableRelationshipTlsActivation(val *RelationshipTlsActivation) *NullableRelationshipTlsActivation {
	return &NullableRelationshipTlsActivation{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsActivation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsActivation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
