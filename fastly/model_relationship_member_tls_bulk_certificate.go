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

// RelationshipMemberTLSBulkCertificate struct for RelationshipMemberTLSBulkCertificate
type RelationshipMemberTLSBulkCertificate struct {
	Type *TypeTLSBulkCertificate `json:"type,omitempty"`
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMemberTLSBulkCertificate RelationshipMemberTLSBulkCertificate

// NewRelationshipMemberTLSBulkCertificate instantiates a new RelationshipMemberTLSBulkCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMemberTLSBulkCertificate() *RelationshipMemberTLSBulkCertificate {
	this := RelationshipMemberTLSBulkCertificate{}
	var resourceType TypeTLSBulkCertificate = TYPETLSBULKCERTIFICATE_TLS_BULK_CERTIFICATE
	this.Type = &resourceType
	return &this
}

// NewRelationshipMemberTLSBulkCertificateWithDefaults instantiates a new RelationshipMemberTLSBulkCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMemberTLSBulkCertificateWithDefaults() *RelationshipMemberTLSBulkCertificate {
	this := RelationshipMemberTLSBulkCertificate{}
	var resourceType TypeTLSBulkCertificate = TYPETLSBULKCERTIFICATE_TLS_BULK_CERTIFICATE
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipMemberTLSBulkCertificate) GetType() TypeTLSBulkCertificate {
	if o == nil || o.Type == nil {
		var ret TypeTLSBulkCertificate
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTLSBulkCertificate) GetTypeOk() (*TypeTLSBulkCertificate, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipMemberTLSBulkCertificate) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeTLSBulkCertificate and assigns it to the Type field.
func (o *RelationshipMemberTLSBulkCertificate) SetType(v TypeTLSBulkCertificate) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *RelationshipMemberTLSBulkCertificate) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberTLSBulkCertificate) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *RelationshipMemberTLSBulkCertificate) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *RelationshipMemberTLSBulkCertificate) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMemberTLSBulkCertificate) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipMemberTLSBulkCertificate) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMemberTLSBulkCertificate := _RelationshipMemberTLSBulkCertificate{}

	if err = json.Unmarshal(bytes, &varRelationshipMemberTLSBulkCertificate); err == nil {
		*o = RelationshipMemberTLSBulkCertificate(varRelationshipMemberTLSBulkCertificate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMemberTLSBulkCertificate is a helper abstraction for handling nullable relationshipmembertlsbulkcertificate types. 
type NullableRelationshipMemberTLSBulkCertificate struct {
	value *RelationshipMemberTLSBulkCertificate
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMemberTLSBulkCertificate) Get() *RelationshipMemberTLSBulkCertificate {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMemberTLSBulkCertificate) Set(val *RelationshipMemberTLSBulkCertificate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMemberTLSBulkCertificate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMemberTLSBulkCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMemberTLSBulkCertificate returns a pointer to a new instance of NullableRelationshipMemberTLSBulkCertificate.
func NewNullableRelationshipMemberTLSBulkCertificate(val *RelationshipMemberTLSBulkCertificate) *NullableRelationshipMemberTLSBulkCertificate {
	return &NullableRelationshipMemberTLSBulkCertificate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMemberTLSBulkCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipMemberTLSBulkCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
