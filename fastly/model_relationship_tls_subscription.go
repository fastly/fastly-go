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

// RelationshipTLSSubscription struct for RelationshipTLSSubscription
type RelationshipTLSSubscription struct {
	TLSSubscription *RelationshipTLSSubscriptionTLSSubscription `json:"tls_subscription,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSSubscription RelationshipTLSSubscription

// NewRelationshipTLSSubscription instantiates a new RelationshipTLSSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSSubscription() *RelationshipTLSSubscription {
	this := RelationshipTLSSubscription{}
	return &this
}

// NewRelationshipTLSSubscriptionWithDefaults instantiates a new RelationshipTLSSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSSubscriptionWithDefaults() *RelationshipTLSSubscription {
	this := RelationshipTLSSubscription{}
	return &this
}

// GetTLSSubscription returns the TLSSubscription field value if set, zero value otherwise.
func (o *RelationshipTLSSubscription) GetTLSSubscription() RelationshipTLSSubscriptionTLSSubscription {
	if o == nil || o.TLSSubscription == nil {
		var ret RelationshipTLSSubscriptionTLSSubscription
		return ret
	}
	return *o.TLSSubscription
}

// GetTLSSubscriptionOk returns a tuple with the TLSSubscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSSubscription) GetTLSSubscriptionOk() (*RelationshipTLSSubscriptionTLSSubscription, bool) {
	if o == nil || o.TLSSubscription == nil {
		return nil, false
	}
	return o.TLSSubscription, true
}

// HasTLSSubscription returns a boolean if a field has been set.
func (o *RelationshipTLSSubscription) HasTLSSubscription() bool {
	if o != nil && o.TLSSubscription != nil {
		return true
	}

	return false
}

// SetTLSSubscription gets a reference to the given RelationshipTLSSubscriptionTLSSubscription and assigns it to the TLSSubscription field.
func (o *RelationshipTLSSubscription) SetTLSSubscription(v RelationshipTLSSubscriptionTLSSubscription) {
	o.TLSSubscription = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSSubscription != nil {
		toSerialize["tls_subscription"] = o.TLSSubscription
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSSubscription) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSSubscription := _RelationshipTLSSubscription{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSSubscription); err == nil {
		*o = RelationshipTLSSubscription(varRelationshipTLSSubscription)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_subscription")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSSubscription is a helper abstraction for handling nullable relationshiptlssubscription types. 
type NullableRelationshipTLSSubscription struct {
	value *RelationshipTLSSubscription
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSSubscription) Get() *RelationshipTLSSubscription {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSSubscription) Set(val *RelationshipTLSSubscription) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSSubscription) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSSubscription returns a pointer to a new instance of NullableRelationshipTLSSubscription.
func NewNullableRelationshipTLSSubscription(val *RelationshipTLSSubscription) *NullableRelationshipTLSSubscription {
	return &NullableRelationshipTLSSubscription{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
