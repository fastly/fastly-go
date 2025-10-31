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

// RelationshipTlsSubscriptionTlsSubscription struct for RelationshipTlsSubscriptionTlsSubscription
type RelationshipTlsSubscriptionTlsSubscription struct {
	Data                 []RelationshipMemberTlsSubscription `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTlsSubscriptionTlsSubscription RelationshipTlsSubscriptionTlsSubscription

// NewRelationshipTlsSubscriptionTlsSubscription instantiates a new RelationshipTlsSubscriptionTlsSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTlsSubscriptionTlsSubscription() *RelationshipTlsSubscriptionTlsSubscription {
	this := RelationshipTlsSubscriptionTlsSubscription{}
	return &this
}

// NewRelationshipTlsSubscriptionTlsSubscriptionWithDefaults instantiates a new RelationshipTlsSubscriptionTlsSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTlsSubscriptionTlsSubscriptionWithDefaults() *RelationshipTlsSubscriptionTlsSubscription {
	this := RelationshipTlsSubscriptionTlsSubscription{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RelationshipTlsSubscriptionTlsSubscription) GetData() []RelationshipMemberTlsSubscription {
	if o == nil || o.Data == nil {
		var ret []RelationshipMemberTlsSubscription
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTlsSubscriptionTlsSubscription) GetDataOk() ([]RelationshipMemberTlsSubscription, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RelationshipTlsSubscriptionTlsSubscription) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []RelationshipMemberTlsSubscription and assigns it to the Data field.
func (o *RelationshipTlsSubscriptionTlsSubscription) SetData(v []RelationshipMemberTlsSubscription) {
	o.Data = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTlsSubscriptionTlsSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipTlsSubscriptionTlsSubscription) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTlsSubscriptionTlsSubscription := _RelationshipTlsSubscriptionTlsSubscription{}

	if err = json.Unmarshal(bytes, &varRelationshipTlsSubscriptionTlsSubscription); err == nil {
		*o = RelationshipTlsSubscriptionTlsSubscription(varRelationshipTlsSubscriptionTlsSubscription)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTlsSubscriptionTlsSubscription is a helper abstraction for handling nullable relationshiptlssubscriptiontlssubscription types.
type NullableRelationshipTlsSubscriptionTlsSubscription struct {
	value *RelationshipTlsSubscriptionTlsSubscription
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTlsSubscriptionTlsSubscription) Get() *RelationshipTlsSubscriptionTlsSubscription {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTlsSubscriptionTlsSubscription) Set(val *RelationshipTlsSubscriptionTlsSubscription) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTlsSubscriptionTlsSubscription) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTlsSubscriptionTlsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTlsSubscriptionTlsSubscription returns a pointer to a new instance of NullableRelationshipTlsSubscriptionTlsSubscription.
func NewNullableRelationshipTlsSubscriptionTlsSubscription(val *RelationshipTlsSubscriptionTlsSubscription) *NullableRelationshipTlsSubscriptionTlsSubscription {
	return &NullableRelationshipTlsSubscriptionTlsSubscription{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTlsSubscriptionTlsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipTlsSubscriptionTlsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
