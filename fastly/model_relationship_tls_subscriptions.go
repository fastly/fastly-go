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

// RelationshipTLSSubscriptions struct for RelationshipTLSSubscriptions
type RelationshipTLSSubscriptions struct {
	TLSSubscriptions *RelationshipTLSSubscriptionTLSSubscription `json:"tls_subscriptions,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSSubscriptions RelationshipTLSSubscriptions

// NewRelationshipTLSSubscriptions instantiates a new RelationshipTLSSubscriptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSSubscriptions() *RelationshipTLSSubscriptions {
	this := RelationshipTLSSubscriptions{}
	return &this
}

// NewRelationshipTLSSubscriptionsWithDefaults instantiates a new RelationshipTLSSubscriptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSSubscriptionsWithDefaults() *RelationshipTLSSubscriptions {
	this := RelationshipTLSSubscriptions{}
	return &this
}

// GetTLSSubscriptions returns the TLSSubscriptions field value if set, zero value otherwise.
func (o *RelationshipTLSSubscriptions) GetTLSSubscriptions() RelationshipTLSSubscriptionTLSSubscription {
	if o == nil || o.TLSSubscriptions == nil {
		var ret RelationshipTLSSubscriptionTLSSubscription
		return ret
	}
	return *o.TLSSubscriptions
}

// GetTLSSubscriptionsOk returns a tuple with the TLSSubscriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSSubscriptions) GetTLSSubscriptionsOk() (*RelationshipTLSSubscriptionTLSSubscription, bool) {
	if o == nil || o.TLSSubscriptions == nil {
		return nil, false
	}
	return o.TLSSubscriptions, true
}

// HasTLSSubscriptions returns a boolean if a field has been set.
func (o *RelationshipTLSSubscriptions) HasTLSSubscriptions() bool {
	if o != nil && o.TLSSubscriptions != nil {
		return true
	}

	return false
}

// SetTLSSubscriptions gets a reference to the given RelationshipTLSSubscriptionTLSSubscription and assigns it to the TLSSubscriptions field.
func (o *RelationshipTLSSubscriptions) SetTLSSubscriptions(v RelationshipTLSSubscriptionTLSSubscription) {
	o.TLSSubscriptions = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSSubscriptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSSubscriptions != nil {
		toSerialize["tls_subscriptions"] = o.TLSSubscriptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSSubscriptions) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSSubscriptions := _RelationshipTLSSubscriptions{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSSubscriptions); err == nil {
		*o = RelationshipTLSSubscriptions(varRelationshipTLSSubscriptions)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_subscriptions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSSubscriptions is a helper abstraction for handling nullable relationshiptlssubscriptions types. 
type NullableRelationshipTLSSubscriptions struct {
	value *RelationshipTLSSubscriptions
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSSubscriptions) Get() *RelationshipTLSSubscriptions {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSSubscriptions) Set(val *RelationshipTLSSubscriptions) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSSubscriptions) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSSubscriptions) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSSubscriptions returns a pointer to a new instance of NullableRelationshipTLSSubscriptions.
func NewNullableRelationshipTLSSubscriptions(val *RelationshipTLSSubscriptions) *NullableRelationshipTLSSubscriptions {
	return &NullableRelationshipTLSSubscriptions{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSSubscriptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSSubscriptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
