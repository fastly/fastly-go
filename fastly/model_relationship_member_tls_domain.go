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

// RelationshipMemberTlsDomain struct for RelationshipMemberTlsDomain
type RelationshipMemberTlsDomain struct {
	Type *TypeTlsDomain `json:"type,omitempty"`
	// The domain name.
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMemberTlsDomain RelationshipMemberTlsDomain

// NewRelationshipMemberTlsDomain instantiates a new RelationshipMemberTlsDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMemberTlsDomain() *RelationshipMemberTlsDomain {
	this := RelationshipMemberTlsDomain{}
	var type_ TypeTlsDomain = TYPETLSDOMAIN_TLS_DOMAIN
	this.Type = &type_
	return &this
}

// NewRelationshipMemberTlsDomainWithDefaults instantiates a new RelationshipMemberTlsDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMemberTlsDomainWithDefaults() *RelationshipMemberTlsDomain {
	this := RelationshipMemberTlsDomain{}
	var type_ TypeTlsDomain = TYPETLSDOMAIN_TLS_DOMAIN
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipMemberTlsDomain) GetType() TypeTlsDomain {
	if o == nil || o.Type == nil {
		var ret TypeTlsDomain
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTlsDomain) GetTypeOk() (*TypeTlsDomain, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipMemberTlsDomain) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTlsDomain and assigns it to the Type field.
func (o *RelationshipMemberTlsDomain) SetType(v TypeTlsDomain) {
	o.Type = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RelationshipMemberTlsDomain) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTlsDomain) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RelationshipMemberTlsDomain) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RelationshipMemberTlsDomain) SetId(v string) {
	o.Id = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMemberTlsDomain) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipMemberTlsDomain) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMemberTlsDomain := _RelationshipMemberTlsDomain{}

	if err = json.Unmarshal(bytes, &varRelationshipMemberTlsDomain); err == nil {
		*o = RelationshipMemberTlsDomain(varRelationshipMemberTlsDomain)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMemberTlsDomain is a helper abstraction for handling nullable relationshipmembertlsdomain types.
type NullableRelationshipMemberTlsDomain struct {
	value *RelationshipMemberTlsDomain
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMemberTlsDomain) Get() *RelationshipMemberTlsDomain {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMemberTlsDomain) Set(val *RelationshipMemberTlsDomain) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMemberTlsDomain) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMemberTlsDomain) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMemberTlsDomain returns a pointer to a new instance of NullableRelationshipMemberTlsDomain.
func NewNullableRelationshipMemberTlsDomain(val *RelationshipMemberTlsDomain) *NullableRelationshipMemberTlsDomain {
	return &NullableRelationshipMemberTlsDomain{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMemberTlsDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipMemberTlsDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
