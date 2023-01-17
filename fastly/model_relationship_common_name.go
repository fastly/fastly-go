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

// RelationshipCommonName struct for RelationshipCommonName
type RelationshipCommonName struct {
	CommonName *RelationshipMemberTLSDomain `json:"common_name,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipCommonName RelationshipCommonName

// NewRelationshipCommonName instantiates a new RelationshipCommonName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipCommonName() *RelationshipCommonName {
	this := RelationshipCommonName{}
	return &this
}

// NewRelationshipCommonNameWithDefaults instantiates a new RelationshipCommonName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipCommonNameWithDefaults() *RelationshipCommonName {
	this := RelationshipCommonName{}
	return &this
}

// GetCommonName returns the CommonName field value if set, zero value otherwise.
func (o *RelationshipCommonName) GetCommonName() RelationshipMemberTLSDomain {
	if o == nil || o.CommonName == nil {
		var ret RelationshipMemberTLSDomain
		return ret
	}
	return *o.CommonName
}

// GetCommonNameOk returns a tuple with the CommonName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipCommonName) GetCommonNameOk() (*RelationshipMemberTLSDomain, bool) {
	if o == nil || o.CommonName == nil {
		return nil, false
	}
	return o.CommonName, true
}

// HasCommonName returns a boolean if a field has been set.
func (o *RelationshipCommonName) HasCommonName() bool {
	if o != nil && o.CommonName != nil {
		return true
	}

	return false
}

// SetCommonName gets a reference to the given RelationshipMemberTLSDomain and assigns it to the CommonName field.
func (o *RelationshipCommonName) SetCommonName(v RelationshipMemberTLSDomain) {
	o.CommonName = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipCommonName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.CommonName != nil {
		toSerialize["common_name"] = o.CommonName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipCommonName) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipCommonName := _RelationshipCommonName{}

	if err = json.Unmarshal(bytes, &varRelationshipCommonName); err == nil {
		*o = RelationshipCommonName(varRelationshipCommonName)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "common_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipCommonName is a helper abstraction for handling nullable relationshipcommonname types. 
type NullableRelationshipCommonName struct {
	value *RelationshipCommonName
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipCommonName) Get() *RelationshipCommonName {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipCommonName) Set(val *RelationshipCommonName) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipCommonName) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipCommonName) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipCommonName returns a pointer to a new instance of NullableRelationshipCommonName.
func NewNullableRelationshipCommonName(val *RelationshipCommonName) *NullableRelationshipCommonName {
	return &NullableRelationshipCommonName{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipCommonName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipCommonName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
