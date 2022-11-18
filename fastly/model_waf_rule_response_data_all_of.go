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

// WafRuleResponseDataAllOf struct for WafRuleResponseDataAllOf
type WafRuleResponseDataAllOf struct {
	Relationships *RelationshipsForWafRule `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafRuleResponseDataAllOf WafRuleResponseDataAllOf

// NewWafRuleResponseDataAllOf instantiates a new WafRuleResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRuleResponseDataAllOf() *WafRuleResponseDataAllOf {
	this := WafRuleResponseDataAllOf{}
	return &this
}

// NewWafRuleResponseDataAllOfWithDefaults instantiates a new WafRuleResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleResponseDataAllOfWithDefaults() *WafRuleResponseDataAllOf {
	this := WafRuleResponseDataAllOf{}
	return &this
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafRuleResponseDataAllOf) GetRelationships() RelationshipsForWafRule {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForWafRule
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleResponseDataAllOf) GetRelationshipsOk() (*RelationshipsForWafRule, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafRuleResponseDataAllOf) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForWafRule and assigns it to the Relationships field.
func (o *WafRuleResponseDataAllOf) SetRelationships(v RelationshipsForWafRule) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafRuleResponseDataAllOf) MarshalJSON() ([]byte, error) {
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
func (o *WafRuleResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafRuleResponseDataAllOf := _WafRuleResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varWafRuleResponseDataAllOf); err == nil {
		*o = WafRuleResponseDataAllOf(varWafRuleResponseDataAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafRuleResponseDataAllOf is a helper abstraction for handling nullable wafruleresponsedataallof types. 
type NullableWafRuleResponseDataAllOf struct {
	value *WafRuleResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafRuleResponseDataAllOf) Get() *WafRuleResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafRuleResponseDataAllOf) Set(val *WafRuleResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafRuleResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafRuleResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafRuleResponseDataAllOf returns a pointer to a new instance of NullableWafRuleResponseDataAllOf.
func NewNullableWafRuleResponseDataAllOf(val *WafRuleResponseDataAllOf) *NullableWafRuleResponseDataAllOf {
	return &NullableWafRuleResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafRuleResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafRuleResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
