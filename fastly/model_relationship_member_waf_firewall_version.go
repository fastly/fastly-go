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

// RelationshipMemberWafFirewallVersion struct for RelationshipMemberWafFirewallVersion
type RelationshipMemberWafFirewallVersion struct {
	Type *TypeWafFirewallVersion `json:"type,omitempty"`
	// Alphanumeric string identifying a Firewall version.
	ID                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMemberWafFirewallVersion RelationshipMemberWafFirewallVersion

// NewRelationshipMemberWafFirewallVersion instantiates a new RelationshipMemberWafFirewallVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMemberWafFirewallVersion() *RelationshipMemberWafFirewallVersion {
	this := RelationshipMemberWafFirewallVersion{}
	var resourceType TypeWafFirewallVersion = TYPEWAFFIREWALLVERSION_WAF_FIREWALL_VERSION
	this.Type = &resourceType
	return &this
}

// NewRelationshipMemberWafFirewallVersionWithDefaults instantiates a new RelationshipMemberWafFirewallVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMemberWafFirewallVersionWithDefaults() *RelationshipMemberWafFirewallVersion {
	this := RelationshipMemberWafFirewallVersion{}
	var resourceType TypeWafFirewallVersion = TYPEWAFFIREWALLVERSION_WAF_FIREWALL_VERSION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipMemberWafFirewallVersion) GetType() TypeWafFirewallVersion {
	if o == nil || o.Type == nil {
		var ret TypeWafFirewallVersion
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberWafFirewallVersion) GetTypeOk() (*TypeWafFirewallVersion, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipMemberWafFirewallVersion) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafFirewallVersion and assigns it to the Type field.
func (o *RelationshipMemberWafFirewallVersion) SetType(v TypeWafFirewallVersion) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *RelationshipMemberWafFirewallVersion) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberWafFirewallVersion) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *RelationshipMemberWafFirewallVersion) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *RelationshipMemberWafFirewallVersion) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMemberWafFirewallVersion) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipMemberWafFirewallVersion) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMemberWafFirewallVersion := _RelationshipMemberWafFirewallVersion{}

	if err = json.Unmarshal(bytes, &varRelationshipMemberWafFirewallVersion); err == nil {
		*o = RelationshipMemberWafFirewallVersion(varRelationshipMemberWafFirewallVersion)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMemberWafFirewallVersion is a helper abstraction for handling nullable relationshipmemberwaffirewallversion types.
type NullableRelationshipMemberWafFirewallVersion struct {
	value *RelationshipMemberWafFirewallVersion
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMemberWafFirewallVersion) Get() *RelationshipMemberWafFirewallVersion {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMemberWafFirewallVersion) Set(val *RelationshipMemberWafFirewallVersion) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMemberWafFirewallVersion) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMemberWafFirewallVersion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMemberWafFirewallVersion returns a pointer to a new instance of NullableRelationshipMemberWafFirewallVersion.
func NewNullableRelationshipMemberWafFirewallVersion(val *RelationshipMemberWafFirewallVersion) *NullableRelationshipMemberWafFirewallVersion {
	return &NullableRelationshipMemberWafFirewallVersion{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMemberWafFirewallVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableRelationshipMemberWafFirewallVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
