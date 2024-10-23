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

// RelationshipMemberTLSSubscription struct for RelationshipMemberTLSSubscription
type RelationshipMemberTLSSubscription struct {
	Type                 *TypeTLSSubscription `json:"type,omitempty"`
	ID                   *string              `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMemberTLSSubscription RelationshipMemberTLSSubscription

// NewRelationshipMemberTLSSubscription instantiates a new RelationshipMemberTLSSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMemberTLSSubscription() *RelationshipMemberTLSSubscription {
	this := RelationshipMemberTLSSubscription{}
	var resourceType TypeTLSSubscription = TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION
	this.Type = &resourceType
	return &this
}

// NewRelationshipMemberTLSSubscriptionWithDefaults instantiates a new RelationshipMemberTLSSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMemberTLSSubscriptionWithDefaults() *RelationshipMemberTLSSubscription {
	this := RelationshipMemberTLSSubscription{}
	var resourceType TypeTLSSubscription = TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipMemberTLSSubscription) GetType() TypeTLSSubscription {
	if o == nil || o.Type == nil {
		var ret TypeTLSSubscription
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTLSSubscription) GetTypeOk() (*TypeTLSSubscription, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipMemberTLSSubscription) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSSubscription and assigns it to the Type field.
func (o *RelationshipMemberTLSSubscription) SetType(v TypeTLSSubscription) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *RelationshipMemberTLSSubscription) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTLSSubscription) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *RelationshipMemberTLSSubscription) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *RelationshipMemberTLSSubscription) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMemberTLSSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipMemberTLSSubscription) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMemberTLSSubscription := _RelationshipMemberTLSSubscription{}

	if err = json.Unmarshal(bytes, &varRelationshipMemberTLSSubscription); err == nil {
		*o = RelationshipMemberTLSSubscription(varRelationshipMemberTLSSubscription)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMemberTLSSubscription is a helper abstraction for handling nullable relationshipmembertlssubscription types.
type NullableRelationshipMemberTLSSubscription struct {
	value *RelationshipMemberTLSSubscription
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMemberTLSSubscription) Get() *RelationshipMemberTLSSubscription {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMemberTLSSubscription) Set(val *RelationshipMemberTLSSubscription) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMemberTLSSubscription) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMemberTLSSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMemberTLSSubscription returns a pointer to a new instance of NullableRelationshipMemberTLSSubscription.
func NewNullableRelationshipMemberTLSSubscription(val *RelationshipMemberTLSSubscription) *NullableRelationshipMemberTLSSubscription {
	return &NullableRelationshipMemberTLSSubscription{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMemberTLSSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipMemberTLSSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
