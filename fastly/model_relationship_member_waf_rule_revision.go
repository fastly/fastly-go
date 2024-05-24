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

// RelationshipMemberWafRuleRevision struct for RelationshipMemberWafRuleRevision
type RelationshipMemberWafRuleRevision struct {
	Type *TypeWafRuleRevision `json:"type,omitempty"`
	// Alphanumeric string identifying a WAF rule revision.
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipMemberWafRuleRevision RelationshipMemberWafRuleRevision

// NewRelationshipMemberWafRuleRevision instantiates a new RelationshipMemberWafRuleRevision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipMemberWafRuleRevision() *RelationshipMemberWafRuleRevision {
	this := RelationshipMemberWafRuleRevision{}
	var resourceType TypeWafRuleRevision = TYPEWAFRULEREVISION_WAF_RULE_REVISION
	this.Type = &resourceType
	return &this
}

// NewRelationshipMemberWafRuleRevisionWithDefaults instantiates a new RelationshipMemberWafRuleRevision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipMemberWafRuleRevisionWithDefaults() *RelationshipMemberWafRuleRevision {
	this := RelationshipMemberWafRuleRevision{}
	var resourceType TypeWafRuleRevision = TYPEWAFRULEREVISION_WAF_RULE_REVISION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RelationshipMemberWafRuleRevision) GetType() TypeWafRuleRevision {
	if o == nil || o.Type == nil {
		var ret TypeWafRuleRevision
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberWafRuleRevision) GetTypeOk() (*TypeWafRuleRevision, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RelationshipMemberWafRuleRevision) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafRuleRevision and assigns it to the Type field.
func (o *RelationshipMemberWafRuleRevision) SetType(v TypeWafRuleRevision) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *RelationshipMemberWafRuleRevision) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipMemberWafRuleRevision) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *RelationshipMemberWafRuleRevision) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *RelationshipMemberWafRuleRevision) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipMemberWafRuleRevision) MarshalJSON() ([]byte, error) {
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
func (o *RelationshipMemberWafRuleRevision) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipMemberWafRuleRevision := _RelationshipMemberWafRuleRevision{}

	if err = json.Unmarshal(bytes, &varRelationshipMemberWafRuleRevision); err == nil {
		*o = RelationshipMemberWafRuleRevision(varRelationshipMemberWafRuleRevision)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipMemberWafRuleRevision is a helper abstraction for handling nullable relationshipmemberwafrulerevision types. 
type NullableRelationshipMemberWafRuleRevision struct {
	value *RelationshipMemberWafRuleRevision
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipMemberWafRuleRevision) Get() *RelationshipMemberWafRuleRevision {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipMemberWafRuleRevision) Set(val *RelationshipMemberWafRuleRevision) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipMemberWafRuleRevision) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipMemberWafRuleRevision) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipMemberWafRuleRevision returns a pointer to a new instance of NullableRelationshipMemberWafRuleRevision.
func NewNullableRelationshipMemberWafRuleRevision(val *RelationshipMemberWafRuleRevision) *NullableRelationshipMemberWafRuleRevision {
	return &NullableRelationshipMemberWafRuleRevision{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipMemberWafRuleRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipMemberWafRuleRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
