// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// RelationshipTLSActivation struct for RelationshipTLSActivation
type RelationshipTLSActivation struct {
	TLSActivation *RelationshipTLSActivationTLSActivation `json:"tls_activation,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSActivation RelationshipTLSActivation

// NewRelationshipTLSActivation instantiates a new RelationshipTLSActivation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSActivation() *RelationshipTLSActivation {
	this := RelationshipTLSActivation{}
	return &this
}

// NewRelationshipTLSActivationWithDefaults instantiates a new RelationshipTLSActivation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSActivationWithDefaults() *RelationshipTLSActivation {
	this := RelationshipTLSActivation{}
	return &this
}

// GetTLSActivation returns the TLSActivation field value if set, zero value otherwise.
func (o *RelationshipTLSActivation) GetTLSActivation() RelationshipTLSActivationTLSActivation {
	if o == nil || o.TLSActivation == nil {
		var ret RelationshipTLSActivationTLSActivation
		return ret
	}
	return *o.TLSActivation
}

// GetTLSActivationOk returns a tuple with the TLSActivation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSActivation) GetTLSActivationOk() (*RelationshipTLSActivationTLSActivation, bool) {
	if o == nil || o.TLSActivation == nil {
		return nil, false
	}
	return o.TLSActivation, true
}

// HasTLSActivation returns a boolean if a field has been set.
func (o *RelationshipTLSActivation) HasTLSActivation() bool {
	if o != nil && o.TLSActivation != nil {
		return true
	}

	return false
}

// SetTLSActivation gets a reference to the given RelationshipTLSActivationTLSActivation and assigns it to the TLSActivation field.
func (o *RelationshipTLSActivation) SetTLSActivation(v RelationshipTLSActivationTLSActivation) {
	o.TLSActivation = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSActivation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSActivation != nil {
		toSerialize["tls_activation"] = o.TLSActivation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSActivation) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSActivation := _RelationshipTLSActivation{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSActivation); err == nil {
		*o = RelationshipTLSActivation(varRelationshipTLSActivation)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_activation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSActivation is a helper abstraction for handling nullable relationshiptlsactivation types. 
type NullableRelationshipTLSActivation struct {
	value *RelationshipTLSActivation
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSActivation) Get() *RelationshipTLSActivation {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSActivation) Set(val *RelationshipTLSActivation) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSActivation) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSActivation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSActivation returns a pointer to a new instance of NullableRelationshipTLSActivation.
func NewNullableRelationshipTLSActivation(val *RelationshipTLSActivation) *NullableRelationshipTLSActivation {
	return &NullableRelationshipTLSActivation{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSActivation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSActivation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
