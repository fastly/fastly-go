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

// RelationshipTLSDomain struct for RelationshipTLSDomain
type RelationshipTLSDomain struct {
	TLSDomain *RelationshipTLSDomainTLSDomain `json:"tls_domain,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipTLSDomain RelationshipTLSDomain

// NewRelationshipTLSDomain instantiates a new RelationshipTLSDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipTLSDomain() *RelationshipTLSDomain {
	this := RelationshipTLSDomain{}
	return &this
}

// NewRelationshipTLSDomainWithDefaults instantiates a new RelationshipTLSDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipTLSDomainWithDefaults() *RelationshipTLSDomain {
	this := RelationshipTLSDomain{}
	return &this
}

// GetTLSDomain returns the TLSDomain field value if set, zero value otherwise.
func (o *RelationshipTLSDomain) GetTLSDomain() RelationshipTLSDomainTLSDomain {
	if o == nil || o.TLSDomain == nil {
		var ret RelationshipTLSDomainTLSDomain
		return ret
	}
	return *o.TLSDomain
}

// GetTLSDomainOk returns a tuple with the TLSDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipTLSDomain) GetTLSDomainOk() (*RelationshipTLSDomainTLSDomain, bool) {
	if o == nil || o.TLSDomain == nil {
		return nil, false
	}
	return o.TLSDomain, true
}

// HasTLSDomain returns a boolean if a field has been set.
func (o *RelationshipTLSDomain) HasTLSDomain() bool {
	if o != nil && o.TLSDomain != nil {
		return true
	}

	return false
}

// SetTLSDomain gets a reference to the given RelationshipTLSDomainTLSDomain and assigns it to the TLSDomain field.
func (o *RelationshipTLSDomain) SetTLSDomain(v RelationshipTLSDomainTLSDomain) {
	o.TLSDomain = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipTLSDomain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.TLSDomain != nil {
		toSerialize["tls_domain"] = o.TLSDomain
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipTLSDomain) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipTLSDomain := _RelationshipTLSDomain{}

	if err = json.Unmarshal(bytes, &varRelationshipTLSDomain); err == nil {
		*o = RelationshipTLSDomain(varRelationshipTLSDomain)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tls_domain")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipTLSDomain is a helper abstraction for handling nullable relationshiptlsdomain types. 
type NullableRelationshipTLSDomain struct {
	value *RelationshipTLSDomain
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipTLSDomain) Get() *RelationshipTLSDomain {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipTLSDomain) Set(val *RelationshipTLSDomain) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipTLSDomain) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipTLSDomain) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipTLSDomain returns a pointer to a new instance of NullableRelationshipTLSDomain.
func NewNullableRelationshipTLSDomain(val *RelationshipTLSDomain) *NullableRelationshipTLSDomain {
	return &NullableRelationshipTLSDomain{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipTLSDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipTLSDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
