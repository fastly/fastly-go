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

// RelationshipTLSConfigurationForTLSSubscription The unique identifier for the set of TLS configuration options that apply to the enabled domains on this subscription. Write-only on create.
type RelationshipTLSConfigurationForTLSSubscription struct {
	TLSConfiguration     *RelationshipTLSConfigurationTLSConfiguration `json:"tls_configuration,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSConfigurationForTLSSubscription RelationshipTLSConfigurationForTLSSubscription

// NewRelationshipTLSConfigurationForTLSSubscription instantiates a new RelationshipTLSConfigurationForTLSSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSConfigurationForTLSSubscription() *RelationshipTLSConfigurationForTLSSubscription {
	this := RelationshipTLSConfigurationForTLSSubscription{}
	return &this
}

// NewRelationshipTLSConfigurationForTLSSubscriptionWithDefaults instantiates a new RelationshipTLSConfigurationForTLSSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSConfigurationForTLSSubscriptionWithDefaults() *RelationshipTLSConfigurationForTLSSubscription {
	this := RelationshipTLSConfigurationForTLSSubscription{}
	return &this
}

// GetTLSConfiguration returns the TLSConfiguration field value if set, zero value otherwise.
func (o *RelationshipTLSConfigurationForTLSSubscription) GetTLSConfiguration() RelationshipTLSConfigurationTLSConfiguration {
	if o == nil || o.TLSConfiguration == nil {
		var ret RelationshipTLSConfigurationTLSConfiguration
		return ret
	}
	return *o.TLSConfiguration
}

// GetTLSConfigurationOk returns a tuple with the TLSConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSConfigurationForTLSSubscription) GetTLSConfigurationOk() (*RelationshipTLSConfigurationTLSConfiguration, bool) {
	if o == nil || o.TLSConfiguration == nil {
		return nil, false
	}
	return o.TLSConfiguration, true
}

// HasTLSConfiguration returns a boolean if a field has been set.
func (o *RelationshipTLSConfigurationForTLSSubscription) HasTLSConfiguration() bool {
	if o != nil && o.TLSConfiguration != nil {
		return true
	}

	return false
}

// SetTLSConfiguration gets a reference to the given RelationshipTLSConfigurationTLSConfiguration and assigns it to the TLSConfiguration field.
func (o *RelationshipTLSConfigurationForTLSSubscription) SetTLSConfiguration(v RelationshipTLSConfigurationTLSConfiguration) {
	o.TLSConfiguration = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSConfigurationForTLSSubscription) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipTLSConfigurationForTLSSubscription) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSConfigurationForTLSSubscription := _RelationshipTLSConfigurationForTLSSubscription{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSConfigurationForTLSSubscription); err == nil {
		*o = RelationshipTLSConfigurationForTLSSubscription(varRelationshipTLSConfigurationForTLSSubscription)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_configuration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSConfigurationForTLSSubscription is a helper abstraction for handling nullable relationshiptlsconfigurationfortlssubscription types.
type NullableRelationshipTLSConfigurationForTLSSubscription struct {
	value *RelationshipTLSConfigurationForTLSSubscription
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSConfigurationForTLSSubscription) Get() *RelationshipTLSConfigurationForTLSSubscription {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSConfigurationForTLSSubscription) Set(val *RelationshipTLSConfigurationForTLSSubscription) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSConfigurationForTLSSubscription) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSConfigurationForTLSSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSConfigurationForTLSSubscription returns a pointer to a new instance of NullableRelationshipTLSConfigurationForTLSSubscription.
func NewNullableRelationshipTLSConfigurationForTLSSubscription(val *RelationshipTLSConfigurationForTLSSubscription) *NullableRelationshipTLSConfigurationForTLSSubscription {
	return &NullableRelationshipTLSConfigurationForTLSSubscription{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSConfigurationForTLSSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTLSConfigurationForTLSSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
