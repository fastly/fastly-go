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

// WafRuleRevisionResponseDataAllOf struct for WafRuleRevisionResponseDataAllOf
type WafRuleRevisionResponseDataAllOf struct {
	Relationships        *RelationshipWafRule `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafRuleRevisionResponseDataAllOf WafRuleRevisionResponseDataAllOf

// NewWafRuleRevisionResponseDataAllOf instantiates a new WafRuleRevisionResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRuleRevisionResponseDataAllOf() *WafRuleRevisionResponseDataAllOf {
	this := WafRuleRevisionResponseDataAllOf{}
	return &this
}

// NewWafRuleRevisionResponseDataAllOfWithDefaults instantiates a new WafRuleRevisionResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleRevisionResponseDataAllOfWithDefaults() *WafRuleRevisionResponseDataAllOf {
	this := WafRuleRevisionResponseDataAllOf{}
	return &this
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafRuleRevisionResponseDataAllOf) GetRelationships() RelationshipWafRule {
	if o == nil || o.Relationships == nil {
		var ret RelationshipWafRule
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionResponseDataAllOf) GetRelationshipsOk() (*RelationshipWafRule, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafRuleRevisionResponseDataAllOf) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipWafRule and assigns it to the Relationships field.
func (o *WafRuleRevisionResponseDataAllOf) SetRelationships(v RelationshipWafRule) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafRuleRevisionResponseDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Relationships != nil {
		toSerialize["relationships"] = o.Relationships
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WafRuleRevisionResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafRuleRevisionResponseDataAllOf := _WafRuleRevisionResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varWafRuleRevisionResponseDataAllOf); err == nil {
		*o = WafRuleRevisionResponseDataAllOf(varWafRuleRevisionResponseDataAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafRuleRevisionResponseDataAllOf is a helper abstraction for handling nullable wafrulerevisionresponsedataallof types.
type NullableWafRuleRevisionResponseDataAllOf struct {
	value *WafRuleRevisionResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafRuleRevisionResponseDataAllOf) Get() *WafRuleRevisionResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafRuleRevisionResponseDataAllOf) Set(val *WafRuleRevisionResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafRuleRevisionResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafRuleRevisionResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafRuleRevisionResponseDataAllOf returns a pointer to a new instance of NullableWafRuleRevisionResponseDataAllOf.
func NewNullableWafRuleRevisionResponseDataAllOf(val *WafRuleRevisionResponseDataAllOf) *NullableWafRuleRevisionResponseDataAllOf {
	return &NullableWafRuleRevisionResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafRuleRevisionResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafRuleRevisionResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
