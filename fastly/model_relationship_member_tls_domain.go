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

// RelationshipMemberTLSDomain struct for RelationshipMemberTLSDomain
type RelationshipMemberTLSDomain struct {
	Type *TypeTLSDomain `json:"type,omitempty"`
	// The domain name.
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMemberTLSDomain RelationshipMemberTLSDomain

// NewRelationshipMemberTLSDomain instantiates a new RelationshipMemberTLSDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMemberTLSDomain() *RelationshipMemberTLSDomain {
	this := RelationshipMemberTLSDomain{}
	var resourceType TypeTLSDomain = TYPETLSDOMAIN_TLS_DOMAIN
	this.Type = &resourceType
	return &this
}

// NewRelationshipMemberTLSDomainWithDefaults instantiates a new RelationshipMemberTLSDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMemberTLSDomainWithDefaults() *RelationshipMemberTLSDomain {
	this := RelationshipMemberTLSDomain{}
	var resourceType TypeTLSDomain = TYPETLSDOMAIN_TLS_DOMAIN
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipMemberTLSDomain) GetType() TypeTLSDomain {
	if o == nil || o.Type == nil {
		var ret TypeTLSDomain
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTLSDomain) GetTypeOk() (*TypeTLSDomain, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipMemberTLSDomain) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSDomain and assigns it to the Type field.
func (o *RelationshipMemberTLSDomain) SetType(v TypeTLSDomain) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *RelationshipMemberTLSDomain) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTLSDomain) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *RelationshipMemberTLSDomain) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *RelationshipMemberTLSDomain) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMemberTLSDomain) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipMemberTLSDomain) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMemberTLSDomain := _RelationshipMemberTLSDomain{}

	if err = json.Unmarshal(bytes, &varRelationshipMemberTLSDomain); err == nil {
		*o = RelationshipMemberTLSDomain(varRelationshipMemberTLSDomain)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMemberTLSDomain is a helper abstraction for handling nullable relationshipmembertlsdomain types. 
type NullableRelationshipMemberTLSDomain struct {
	value *RelationshipMemberTLSDomain
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMemberTLSDomain) Get() *RelationshipMemberTLSDomain {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMemberTLSDomain) Set(val *RelationshipMemberTLSDomain) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMemberTLSDomain) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMemberTLSDomain) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMemberTLSDomain returns a pointer to a new instance of NullableRelationshipMemberTLSDomain.
func NewNullableRelationshipMemberTLSDomain(val *RelationshipMemberTLSDomain) *NullableRelationshipMemberTLSDomain {
	return &NullableRelationshipMemberTLSDomain{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMemberTLSDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipMemberTLSDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
