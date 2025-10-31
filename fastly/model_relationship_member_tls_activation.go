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

// RelationshipMemberTlsActivation struct for RelationshipMemberTlsActivation
type RelationshipMemberTlsActivation struct {
	Type                 *TypeTlsActivation `json:"type,omitempty"`
	Id                   *string            `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMemberTlsActivation RelationshipMemberTlsActivation

// NewRelationshipMemberTlsActivation instantiates a new RelationshipMemberTlsActivation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMemberTlsActivation() *RelationshipMemberTlsActivation {
	this := RelationshipMemberTlsActivation{}
	var type_ TypeTlsActivation = TYPETLSACTIVATION_TLS_ACTIVATION
	this.Type = &type_
	return &this
}

// NewRelationshipMemberTlsActivationWithDefaults instantiates a new RelationshipMemberTlsActivation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMemberTlsActivationWithDefaults() *RelationshipMemberTlsActivation {
	this := RelationshipMemberTlsActivation{}
	var type_ TypeTlsActivation = TYPETLSACTIVATION_TLS_ACTIVATION
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipMemberTlsActivation) GetType() TypeTlsActivation {
	if o == nil || o.Type == nil {
		var ret TypeTlsActivation
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTlsActivation) GetTypeOk() (*TypeTlsActivation, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipMemberTlsActivation) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsActivation and assigns it to the Type field.
func (o *RelationshipMemberTlsActivation) SetType(v TypeTlsActivation) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RelationshipMemberTlsActivation) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTlsActivation) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RelationshipMemberTlsActivation) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RelationshipMemberTlsActivation) SetId(v string) {
	o.Id = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMemberTlsActivation) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipMemberTlsActivation) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMemberTlsActivation := _RelationshipMemberTlsActivation{}

	if err = json.Unmarshal(bytes, &varRelationshipMemberTlsActivation); err == nil {
		*o = RelationshipMemberTlsActivation(varRelationshipMemberTlsActivation)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMemberTlsActivation is a helper abstraction for handling nullable relationshipmembertlsactivation types.
type NullableRelationshipMemberTlsActivation struct {
	value *RelationshipMemberTlsActivation
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMemberTlsActivation) Get() *RelationshipMemberTlsActivation {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMemberTlsActivation) Set(val *RelationshipMemberTlsActivation) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMemberTlsActivation) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMemberTlsActivation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMemberTlsActivation returns a pointer to a new instance of NullableRelationshipMemberTlsActivation.
func NewNullableRelationshipMemberTlsActivation(val *RelationshipMemberTlsActivation) *NullableRelationshipMemberTlsActivation {
	return &NullableRelationshipMemberTlsActivation{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMemberTlsActivation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipMemberTlsActivation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
