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

// RelationshipMemberTlsSubscription struct for RelationshipMemberTlsSubscription
type RelationshipMemberTlsSubscription struct {
	Type                 *TypeTlsSubscription `json:"type,omitempty"`
	Id                   *string              `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMemberTlsSubscription RelationshipMemberTlsSubscription

// NewRelationshipMemberTlsSubscription instantiates a new RelationshipMemberTlsSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMemberTlsSubscription() *RelationshipMemberTlsSubscription {
	this := RelationshipMemberTlsSubscription{}
	var type_ TypeTlsSubscription = TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION
	this.Type = &type_
	return &this
}

// NewRelationshipMemberTlsSubscriptionWithDefaults instantiates a new RelationshipMemberTlsSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMemberTlsSubscriptionWithDefaults() *RelationshipMemberTlsSubscription {
	this := RelationshipMemberTlsSubscription{}
	var type_ TypeTlsSubscription = TYPETLSSUBSCRIPTION_TLS_SUBSCRIPTION
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipMemberTlsSubscription) GetType() TypeTlsSubscription {
	if o == nil || o.Type == nil {
		var ret TypeTlsSubscription
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTlsSubscription) GetTypeOk() (*TypeTlsSubscription, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipMemberTlsSubscription) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsSubscription and assigns it to the Type field.
func (o *RelationshipMemberTlsSubscription) SetType(v TypeTlsSubscription) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RelationshipMemberTlsSubscription) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTlsSubscription) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RelationshipMemberTlsSubscription) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RelationshipMemberTlsSubscription) SetId(v string) {
	o.Id = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMemberTlsSubscription) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipMemberTlsSubscription) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMemberTlsSubscription := _RelationshipMemberTlsSubscription{}

	if err = json.Unmarshal(bytes, &varRelationshipMemberTlsSubscription); err == nil {
		*o = RelationshipMemberTlsSubscription(varRelationshipMemberTlsSubscription)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMemberTlsSubscription is a helper abstraction for handling nullable relationshipmembertlssubscription types.
type NullableRelationshipMemberTlsSubscription struct {
	value *RelationshipMemberTlsSubscription
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMemberTlsSubscription) Get() *RelationshipMemberTlsSubscription {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMemberTlsSubscription) Set(val *RelationshipMemberTlsSubscription) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMemberTlsSubscription) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMemberTlsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMemberTlsSubscription returns a pointer to a new instance of NullableRelationshipMemberTlsSubscription.
func NewNullableRelationshipMemberTlsSubscription(val *RelationshipMemberTlsSubscription) *NullableRelationshipMemberTlsSubscription {
	return &NullableRelationshipMemberTlsSubscription{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMemberTlsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipMemberTlsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
