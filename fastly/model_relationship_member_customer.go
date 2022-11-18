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

// RelationshipMemberCustomer struct for RelationshipMemberCustomer
type RelationshipMemberCustomer struct {
	Type *TypeCustomer `json:"type,omitempty"`
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMemberCustomer RelationshipMemberCustomer

// NewRelationshipMemberCustomer instantiates a new RelationshipMemberCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMemberCustomer() *RelationshipMemberCustomer {
	this := RelationshipMemberCustomer{}
	var resourceType TypeCustomer = TYPECUSTOMER_CUSTOMER
	this.Type = &resourceType
	return &this
}

// NewRelationshipMemberCustomerWithDefaults instantiates a new RelationshipMemberCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMemberCustomerWithDefaults() *RelationshipMemberCustomer {
	this := RelationshipMemberCustomer{}
	var resourceType TypeCustomer = TYPECUSTOMER_CUSTOMER
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipMemberCustomer) GetType() TypeCustomer {
	if o == nil || o.Type == nil {
		var ret TypeCustomer
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberCustomer) GetTypeOk() (*TypeCustomer, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipMemberCustomer) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeCustomer and assigns it to the Type field.
func (o *RelationshipMemberCustomer) SetType(v TypeCustomer) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *RelationshipMemberCustomer) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberCustomer) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *RelationshipMemberCustomer) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *RelationshipMemberCustomer) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMemberCustomer) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipMemberCustomer) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMemberCustomer := _RelationshipMemberCustomer{}

	if err = json.Unmarshal(bytes, &varRelationshipMemberCustomer); err == nil {
		*o = RelationshipMemberCustomer(varRelationshipMemberCustomer)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMemberCustomer is a helper abstraction for handling nullable relationshipmembercustomer types. 
type NullableRelationshipMemberCustomer struct {
	value *RelationshipMemberCustomer
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMemberCustomer) Get() *RelationshipMemberCustomer {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMemberCustomer) Set(val *RelationshipMemberCustomer) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMemberCustomer) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMemberCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMemberCustomer returns a pointer to a new instance of NullableRelationshipMemberCustomer.
func NewNullableRelationshipMemberCustomer(val *RelationshipMemberCustomer) *NullableRelationshipMemberCustomer {
	return &NullableRelationshipMemberCustomer{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMemberCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipMemberCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
