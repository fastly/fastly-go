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

// RelationshipMutualAuthentication The Mutual Authentication for client-to-server authentication. Optional.
type RelationshipMutualAuthentication struct {
	MutualAuthentication *RelationshipMutualAuthenticationMutualAuthentication `json:"mutual_authentication,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMutualAuthentication RelationshipMutualAuthentication

// NewRelationshipMutualAuthentication instantiates a new RelationshipMutualAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMutualAuthentication() *RelationshipMutualAuthentication {
	this := RelationshipMutualAuthentication{}
	return &this
}

// NewRelationshipMutualAuthenticationWithDefaults instantiates a new RelationshipMutualAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMutualAuthenticationWithDefaults() *RelationshipMutualAuthentication {
	this := RelationshipMutualAuthentication{}
	return &this
}

// GetMutualAuthentication returns the MutualAuthentication field value if set, zero value otherwise.
func (o *RelationshipMutualAuthentication) GetMutualAuthentication() RelationshipMutualAuthenticationMutualAuthentication {
	if o == nil || o.MutualAuthentication == nil {
		var ret RelationshipMutualAuthenticationMutualAuthentication
		return ret
	}
	return *o.MutualAuthentication
}

// GetMutualAuthenticationOk returns a tuple with the MutualAuthentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMutualAuthentication) GetMutualAuthenticationOk() (*RelationshipMutualAuthenticationMutualAuthentication, bool) {
	if o == nil || o.MutualAuthentication == nil {
		return nil, false
	}
	return o.MutualAuthentication, true
}

// HasMutualAuthentication returns a boolean if a field has been set.
func (o *RelationshipMutualAuthentication) HasMutualAuthentication() bool {
	if o != nil && o.MutualAuthentication != nil {
		return true
	}

	return false
}

// SetMutualAuthentication gets a reference to the given RelationshipMutualAuthenticationMutualAuthentication and assigns it to the MutualAuthentication field.
func (o *RelationshipMutualAuthentication) SetMutualAuthentication(v RelationshipMutualAuthenticationMutualAuthentication) {
	o.MutualAuthentication = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMutualAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.MutualAuthentication != nil {
		toSerialize["mutual_authentication"] = o.MutualAuthentication
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipMutualAuthentication) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMutualAuthentication := _RelationshipMutualAuthentication{}

	if err = json.Unmarshal(bytes, &varRelationshipMutualAuthentication); err == nil {
		*o = RelationshipMutualAuthentication(varRelationshipMutualAuthentication)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "mutual_authentication")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMutualAuthentication is a helper abstraction for handling nullable relationshipmutualauthentication types. 
type NullableRelationshipMutualAuthentication struct {
	value *RelationshipMutualAuthentication
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMutualAuthentication) Get() *RelationshipMutualAuthentication {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMutualAuthentication) Set(val *RelationshipMutualAuthentication) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMutualAuthentication) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMutualAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMutualAuthentication returns a pointer to a new instance of NullableRelationshipMutualAuthentication.
func NewNullableRelationshipMutualAuthentication(val *RelationshipMutualAuthentication) *NullableRelationshipMutualAuthentication {
	return &NullableRelationshipMutualAuthentication{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMutualAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipMutualAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
