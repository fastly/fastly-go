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

// RelationshipMutualAuthenticationMutualAuthentication struct for RelationshipMutualAuthenticationMutualAuthentication
type RelationshipMutualAuthenticationMutualAuthentication struct {
	Data *RelationshipMemberMutualAuthentication `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMutualAuthenticationMutualAuthentication RelationshipMutualAuthenticationMutualAuthentication

// NewRelationshipMutualAuthenticationMutualAuthentication instantiates a new RelationshipMutualAuthenticationMutualAuthentication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMutualAuthenticationMutualAuthentication() *RelationshipMutualAuthenticationMutualAuthentication {
	this := RelationshipMutualAuthenticationMutualAuthentication{}
	return &this
}

// NewRelationshipMutualAuthenticationMutualAuthenticationWithDefaults instantiates a new RelationshipMutualAuthenticationMutualAuthentication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMutualAuthenticationMutualAuthenticationWithDefaults() *RelationshipMutualAuthenticationMutualAuthentication {
	this := RelationshipMutualAuthenticationMutualAuthentication{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RelationshipMutualAuthenticationMutualAuthentication) GetData() RelationshipMemberMutualAuthentication {
	if o == nil || o.Data == nil {
		var ret RelationshipMemberMutualAuthentication
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMutualAuthenticationMutualAuthentication) GetDataOk() (*RelationshipMemberMutualAuthentication, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RelationshipMutualAuthenticationMutualAuthentication) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given RelationshipMemberMutualAuthentication and assigns it to the Data field.
func (o *RelationshipMutualAuthenticationMutualAuthentication) SetData(v RelationshipMemberMutualAuthentication) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMutualAuthenticationMutualAuthentication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipMutualAuthenticationMutualAuthentication) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMutualAuthenticationMutualAuthentication := _RelationshipMutualAuthenticationMutualAuthentication{}

	if err = json.Unmarshal(bytes, &varRelationshipMutualAuthenticationMutualAuthentication); err == nil {
		*o = RelationshipMutualAuthenticationMutualAuthentication(varRelationshipMutualAuthenticationMutualAuthentication)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMutualAuthenticationMutualAuthentication is a helper abstraction for handling nullable relationshipmutualauthenticationmutualauthentication types. 
type NullableRelationshipMutualAuthenticationMutualAuthentication struct {
	value *RelationshipMutualAuthenticationMutualAuthentication
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMutualAuthenticationMutualAuthentication) Get() *RelationshipMutualAuthenticationMutualAuthentication {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMutualAuthenticationMutualAuthentication) Set(val *RelationshipMutualAuthenticationMutualAuthentication) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMutualAuthenticationMutualAuthentication) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMutualAuthenticationMutualAuthentication) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMutualAuthenticationMutualAuthentication returns a pointer to a new instance of NullableRelationshipMutualAuthenticationMutualAuthentication.
func NewNullableRelationshipMutualAuthenticationMutualAuthentication(val *RelationshipMutualAuthenticationMutualAuthentication) *NullableRelationshipMutualAuthenticationMutualAuthentication {
	return &NullableRelationshipMutualAuthenticationMutualAuthentication{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMutualAuthenticationMutualAuthentication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipMutualAuthenticationMutualAuthentication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
