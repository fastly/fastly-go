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

// RelationshipTlsSubscriptions struct for RelationshipTlsSubscriptions
type RelationshipTlsSubscriptions struct {
	TlsSubscriptions     *RelationshipTlsSubscriptionTlsSubscription `json:"tls_subscriptions,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsSubscriptions RelationshipTlsSubscriptions

// NewRelationshipTlsSubscriptions instantiates a new RelationshipTlsSubscriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsSubscriptions() *RelationshipTlsSubscriptions {
	this := RelationshipTlsSubscriptions{}
	return &this
}

// NewRelationshipTlsSubscriptionsWithDefaults instantiates a new RelationshipTlsSubscriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsSubscriptionsWithDefaults() *RelationshipTlsSubscriptions {
	this := RelationshipTlsSubscriptions{}
	return &this
}

// GetTlsSubscriptions returns the TlsSubscriptions field value if set, zero value otherwise.
func (o *RelationshipTlsSubscriptions) GetTlsSubscriptions() RelationshipTlsSubscriptionTlsSubscription {
	if o == nil || o.TlsSubscriptions == nil {
		var ret RelationshipTlsSubscriptionTlsSubscription
		return ret
	}
	return *o.TlsSubscriptions
}

// GetTlsSubscriptionsOk returns a tuple with the TlsSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsSubscriptions) GetTlsSubscriptionsOk() (*RelationshipTlsSubscriptionTlsSubscription, bool) {
	if o == nil || o.TlsSubscriptions == nil {
		return nil, false
	}
	return o.TlsSubscriptions, true
}

// HasTlsSubscriptions returns a boolean if a field has been set.
func (o *RelationshipTlsSubscriptions) HasTlsSubscriptions() bool {
	if o != nil && o.TlsSubscriptions != nil {
		return true
	}

	return false
}

// SetTlsSubscriptions gets a reference to the given RelationshipTlsSubscriptionTlsSubscription and assigns it to the TlsSubscriptions field.
func (o *RelationshipTlsSubscriptions) SetTlsSubscriptions(v RelationshipTlsSubscriptionTlsSubscription) {
	o.TlsSubscriptions = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsSubscriptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsSubscriptions != nil {
		toSerialize["tls_subscriptions"] = o.TlsSubscriptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsSubscriptions) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsSubscriptions := _RelationshipTlsSubscriptions{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsSubscriptions); err == nil {
		*o = RelationshipTlsSubscriptions(varRelationshipTlsSubscriptions)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_subscriptions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsSubscriptions is a helper abstraction for handling nullable relationshiptlssubscriptions types.
type NullableRelationshipTlsSubscriptions struct {
	value *RelationshipTlsSubscriptions
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsSubscriptions) Get() *RelationshipTlsSubscriptions {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsSubscriptions) Set(val *RelationshipTlsSubscriptions) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsSubscriptions) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsSubscriptions returns a pointer to a new instance of NullableRelationshipTlsSubscriptions.
func NewNullableRelationshipTlsSubscriptions(val *RelationshipTlsSubscriptions) *NullableRelationshipTlsSubscriptions {
	return &NullableRelationshipTlsSubscriptions{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
