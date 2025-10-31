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

// RelationshipTlsSubscription struct for RelationshipTlsSubscription
type RelationshipTlsSubscription struct {
	TlsSubscription      *RelationshipTlsSubscriptionTlsSubscription `json:"tls_subscription,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsSubscription RelationshipTlsSubscription

// NewRelationshipTlsSubscription instantiates a new RelationshipTlsSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsSubscription() *RelationshipTlsSubscription {
	this := RelationshipTlsSubscription{}
	return &this
}

// NewRelationshipTlsSubscriptionWithDefaults instantiates a new RelationshipTlsSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsSubscriptionWithDefaults() *RelationshipTlsSubscription {
	this := RelationshipTlsSubscription{}
	return &this
}

// GetTlsSubscription returns the TlsSubscription field value if set, zero value otherwise.
func (o *RelationshipTlsSubscription) GetTlsSubscription() RelationshipTlsSubscriptionTlsSubscription {
	if o == nil || o.TlsSubscription == nil {
		var ret RelationshipTlsSubscriptionTlsSubscription
		return ret
	}
	return *o.TlsSubscription
}

// GetTlsSubscriptionOk returns a tuple with the TlsSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsSubscription) GetTlsSubscriptionOk() (*RelationshipTlsSubscriptionTlsSubscription, bool) {
	if o == nil || o.TlsSubscription == nil {
		return nil, false
	}
	return o.TlsSubscription, true
}

// HasTlsSubscription returns a boolean if a field has been set.
func (o *RelationshipTlsSubscription) HasTlsSubscription() bool {
	if o != nil && o.TlsSubscription != nil {
		return true
	}

	return false
}

// SetTlsSubscription gets a reference to the given RelationshipTlsSubscriptionTlsSubscription and assigns it to the TlsSubscription field.
func (o *RelationshipTlsSubscription) SetTlsSubscription(v RelationshipTlsSubscriptionTlsSubscription) {
	o.TlsSubscription = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TlsSubscription != nil {
		toSerialize["tls_subscription"] = o.TlsSubscription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsSubscription) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsSubscription := _RelationshipTlsSubscription{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsSubscription); err == nil {
		*o = RelationshipTlsSubscription(varRelationshipTlsSubscription)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_subscription")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsSubscription is a helper abstraction for handling nullable relationshiptlssubscription types.
type NullableRelationshipTlsSubscription struct {
	value *RelationshipTlsSubscription
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsSubscription) Get() *RelationshipTlsSubscription {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsSubscription) Set(val *RelationshipTlsSubscription) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsSubscription) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsSubscription returns a pointer to a new instance of NullableRelationshipTlsSubscription.
func NewNullableRelationshipTlsSubscription(val *RelationshipTlsSubscription) *NullableRelationshipTlsSubscription {
	return &NullableRelationshipTlsSubscription{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
