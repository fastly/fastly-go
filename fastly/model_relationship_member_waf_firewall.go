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

// RelationshipMemberWafFirewall struct for RelationshipMemberWafFirewall
type RelationshipMemberWafFirewall struct {
	Type *TypeWafFirewall `json:"type,omitempty"`
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMemberWafFirewall RelationshipMemberWafFirewall

// NewRelationshipMemberWafFirewall instantiates a new RelationshipMemberWafFirewall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMemberWafFirewall() *RelationshipMemberWafFirewall {
	this := RelationshipMemberWafFirewall{}
	var resourceType TypeWafFirewall = TYPEWAFFIREWALL_WAF_FIREWALL
	this.Type = &resourceType
	return &this
}

// NewRelationshipMemberWafFirewallWithDefaults instantiates a new RelationshipMemberWafFirewall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMemberWafFirewallWithDefaults() *RelationshipMemberWafFirewall {
	this := RelationshipMemberWafFirewall{}
	var resourceType TypeWafFirewall = TYPEWAFFIREWALL_WAF_FIREWALL
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipMemberWafFirewall) GetType() TypeWafFirewall {
	if o == nil || o.Type == nil {
		var ret TypeWafFirewall
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberWafFirewall) GetTypeOk() (*TypeWafFirewall, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipMemberWafFirewall) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafFirewall and assigns it to the Type field.
func (o *RelationshipMemberWafFirewall) SetType(v TypeWafFirewall) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *RelationshipMemberWafFirewall) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberWafFirewall) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *RelationshipMemberWafFirewall) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *RelationshipMemberWafFirewall) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMemberWafFirewall) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipMemberWafFirewall) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMemberWafFirewall := _RelationshipMemberWafFirewall{}

	if err = json.Unmarshal(bytes, &varRelationshipMemberWafFirewall); err == nil {
		*o = RelationshipMemberWafFirewall(varRelationshipMemberWafFirewall)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMemberWafFirewall is a helper abstraction for handling nullable relationshipmemberwaffirewall types. 
type NullableRelationshipMemberWafFirewall struct {
	value *RelationshipMemberWafFirewall
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMemberWafFirewall) Get() *RelationshipMemberWafFirewall {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMemberWafFirewall) Set(val *RelationshipMemberWafFirewall) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMemberWafFirewall) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMemberWafFirewall) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMemberWafFirewall returns a pointer to a new instance of NullableRelationshipMemberWafFirewall.
func NewNullableRelationshipMemberWafFirewall(val *RelationshipMemberWafFirewall) *NullableRelationshipMemberWafFirewall {
	return &NullableRelationshipMemberWafFirewall{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMemberWafFirewall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipMemberWafFirewall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
