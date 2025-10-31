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

// RelationshipTlsConfigurationForTlsSubscription The unique identifier for the set of TLS configuration options that apply to the enabled domains on this subscription. Write-only on create.
type RelationshipTlsConfigurationForTlsSubscription struct {
	TlsConfiguration     *RelationshipTlsConfigurationTlsConfiguration `json:"tls_configuration,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsConfigurationForTlsSubscription RelationshipTlsConfigurationForTlsSubscription

// NewRelationshipTlsConfigurationForTlsSubscription instantiates a new RelationshipTlsConfigurationForTlsSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsConfigurationForTlsSubscription() *RelationshipTlsConfigurationForTlsSubscription {
	this := RelationshipTlsConfigurationForTlsSubscription{}
	return &this
}

// NewRelationshipTlsConfigurationForTlsSubscriptionWithDefaults instantiates a new RelationshipTlsConfigurationForTlsSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsConfigurationForTlsSubscriptionWithDefaults() *RelationshipTlsConfigurationForTlsSubscription {
	this := RelationshipTlsConfigurationForTlsSubscription{}
	return &this
}

// GetTlsConfiguration returns the TlsConfiguration field value if set, zero value otherwise.
func (o *RelationshipTlsConfigurationForTlsSubscription) GetTlsConfiguration() RelationshipTlsConfigurationTlsConfiguration {
	if o == nil || o.TlsConfiguration == nil {
		var ret RelationshipTlsConfigurationTlsConfiguration
		return ret
	}
	return *o.TlsConfiguration
}

// GetTlsConfigurationOk returns a tuple with the TlsConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsConfigurationForTlsSubscription) GetTlsConfigurationOk() (*RelationshipTlsConfigurationTlsConfiguration, bool) {
	if o == nil || o.TlsConfiguration == nil {
		return nil, false
	}
	return o.TlsConfiguration, true
}

// HasTlsConfiguration returns a boolean if a field has been set.
func (o *RelationshipTlsConfigurationForTlsSubscription) HasTlsConfiguration() bool {
	if o != nil && o.TlsConfiguration != nil {
		return true
	}

	return false
}

// SetTlsConfiguration gets a reference to the given RelationshipTlsConfigurationTlsConfiguration and assigns it to the TlsConfiguration field.
func (o *RelationshipTlsConfigurationForTlsSubscription) SetTlsConfiguration(v RelationshipTlsConfigurationTlsConfiguration) {
	o.TlsConfiguration = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsConfigurationForTlsSubscription) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipTlsConfigurationForTlsSubscription) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsConfigurationForTlsSubscription := _RelationshipTlsConfigurationForTlsSubscription{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsConfigurationForTlsSubscription); err == nil {
		*o = RelationshipTlsConfigurationForTlsSubscription(varRelationshipTlsConfigurationForTlsSubscription)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_configuration")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsConfigurationForTlsSubscription is a helper abstraction for handling nullable relationshiptlsconfigurationfortlssubscription types.
type NullableRelationshipTlsConfigurationForTlsSubscription struct {
	value *RelationshipTlsConfigurationForTlsSubscription
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsConfigurationForTlsSubscription) Get() *RelationshipTlsConfigurationForTlsSubscription {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsConfigurationForTlsSubscription) Set(val *RelationshipTlsConfigurationForTlsSubscription) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsConfigurationForTlsSubscription) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsConfigurationForTlsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsConfigurationForTlsSubscription returns a pointer to a new instance of NullableRelationshipTlsConfigurationForTlsSubscription.
func NewNullableRelationshipTlsConfigurationForTlsSubscription(val *RelationshipTlsConfigurationForTlsSubscription) *NullableRelationshipTlsConfigurationForTlsSubscription {
	return &NullableRelationshipTlsConfigurationForTlsSubscription{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsConfigurationForTlsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsConfigurationForTlsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
