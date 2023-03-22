// Package fastly is an API client library for interacting with most facets of the Fastly API.
package fastly

/*
Fastly API

Via the Fastly API you can perform any of the operations that are possible within the management console,  including creating services, domains, and backends, configuring rules or uploading your own application code, as well as account operations such as user administration and billing reports. The API is organized into collections of endpoints that allow manipulation of objects related to Fastly services and accounts. For the most accurate and up-to-date API reference content, visit our [Developer Hub](https://developer.fastly.com/reference/api/) 

API version: 1.0.0
Contact: oss@fastly.com
*/

// This code is auto-generated; DO NOT EDIT.


import (
	"encoding/json"
)

// RelationshipMemberTLSActivation struct for RelationshipMemberTLSActivation
type RelationshipMemberTLSActivation struct {
	Type *TypeTLSActivation `json:"type,omitempty"`
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMemberTLSActivation RelationshipMemberTLSActivation

// NewRelationshipMemberTLSActivation instantiates a new RelationshipMemberTLSActivation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMemberTLSActivation() *RelationshipMemberTLSActivation {
	this := RelationshipMemberTLSActivation{}
	var resourceType TypeTLSActivation = TYPETLSACTIVATION_TLS_ACTIVATION
	this.Type = &resourceType
	return &this
}

// NewRelationshipMemberTLSActivationWithDefaults instantiates a new RelationshipMemberTLSActivation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMemberTLSActivationWithDefaults() *RelationshipMemberTLSActivation {
	this := RelationshipMemberTLSActivation{}
	var resourceType TypeTLSActivation = TYPETLSACTIVATION_TLS_ACTIVATION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipMemberTLSActivation) GetType() TypeTLSActivation {
	if o == nil || o.Type == nil {
		var ret TypeTLSActivation
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTLSActivation) GetTypeOk() (*TypeTLSActivation, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipMemberTLSActivation) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSActivation and assigns it to the Type field.
func (o *RelationshipMemberTLSActivation) SetType(v TypeTLSActivation) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *RelationshipMemberTLSActivation) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTLSActivation) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *RelationshipMemberTLSActivation) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *RelationshipMemberTLSActivation) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMemberTLSActivation) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipMemberTLSActivation) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMemberTLSActivation := _RelationshipMemberTLSActivation{}

	if err = json.Unmarshal(bytes, &varRelationshipMemberTLSActivation); err == nil {
		*o = RelationshipMemberTLSActivation(varRelationshipMemberTLSActivation)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMemberTLSActivation is a helper abstraction for handling nullable relationshipmembertlsactivation types. 
type NullableRelationshipMemberTLSActivation struct {
	value *RelationshipMemberTLSActivation
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMemberTLSActivation) Get() *RelationshipMemberTLSActivation {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMemberTLSActivation) Set(val *RelationshipMemberTLSActivation) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMemberTLSActivation) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMemberTLSActivation) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMemberTLSActivation returns a pointer to a new instance of NullableRelationshipMemberTLSActivation.
func NewNullableRelationshipMemberTLSActivation(val *RelationshipMemberTLSActivation) *NullableRelationshipMemberTLSActivation {
	return &NullableRelationshipMemberTLSActivation{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMemberTLSActivation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipMemberTLSActivation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
