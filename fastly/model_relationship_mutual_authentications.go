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

// RelationshipMutualAuthentications struct for RelationshipMutualAuthentications
type RelationshipMutualAuthentications struct {
	MutualAuthentications *RelationshipMutualAuthenticationsMutualAuthentications `json:"mutual_authentications,omitempty"`
	AdditionalProperties  map[string]any
}

type _RelationshipMutualAuthentications RelationshipMutualAuthentications

// NewRelationshipMutualAuthentications instantiates a new RelationshipMutualAuthentications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMutualAuthentications() *RelationshipMutualAuthentications {
	this := RelationshipMutualAuthentications{}
	return &this
}

// NewRelationshipMutualAuthenticationsWithDefaults instantiates a new RelationshipMutualAuthentications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMutualAuthenticationsWithDefaults() *RelationshipMutualAuthentications {
	this := RelationshipMutualAuthentications{}
	return &this
}

// GetMutualAuthentications returns the MutualAuthentications field value if set, zero value otherwise.
func (o *RelationshipMutualAuthentications) GetMutualAuthentications() RelationshipMutualAuthenticationsMutualAuthentications {
	if o == nil || o.MutualAuthentications == nil {
		var ret RelationshipMutualAuthenticationsMutualAuthentications
		return ret
	}
	return *o.MutualAuthentications
}

// GetMutualAuthenticationsOk returns a tuple with the MutualAuthentications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMutualAuthentications) GetMutualAuthenticationsOk() (*RelationshipMutualAuthenticationsMutualAuthentications, bool) {
	if o == nil || o.MutualAuthentications == nil {
		return nil, false
	}
	return o.MutualAuthentications, true
}

// HasMutualAuthentications returns a boolean if a field has been set.
func (o *RelationshipMutualAuthentications) HasMutualAuthentications() bool {
	if o != nil && o.MutualAuthentications != nil {
		return true
	}

	return false
}

// SetMutualAuthentications gets a reference to the given RelationshipMutualAuthenticationsMutualAuthentications and assigns it to the MutualAuthentications field.
func (o *RelationshipMutualAuthentications) SetMutualAuthentications(v RelationshipMutualAuthenticationsMutualAuthentications) {
	o.MutualAuthentications = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMutualAuthentications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.MutualAuthentications != nil {
		toSerialize["mutual_authentications"] = o.MutualAuthentications
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *RelationshipMutualAuthentications) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMutualAuthentications := _RelationshipMutualAuthentications{}

	if err = json.Unmarshal(bytes, &varRelationshipMutualAuthentications); err == nil {
		*o = RelationshipMutualAuthentications(varRelationshipMutualAuthentications)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "mutual_authentications")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMutualAuthentications is a helper abstraction for handling nullable relationshipmutualauthentications types.
type NullableRelationshipMutualAuthentications struct {
	value *RelationshipMutualAuthentications
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMutualAuthentications) Get() *RelationshipMutualAuthentications {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMutualAuthentications) Set(val *RelationshipMutualAuthentications) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMutualAuthentications) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMutualAuthentications) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMutualAuthentications returns a pointer to a new instance of NullableRelationshipMutualAuthentications.
func NewNullableRelationshipMutualAuthentications(val *RelationshipMutualAuthentications) *NullableRelationshipMutualAuthentications {
	return &NullableRelationshipMutualAuthentications{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMutualAuthentications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipMutualAuthentications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
