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

// WafRule struct for WafRule
type WafRule struct {
	Type                 *TypeWafRule       `json:"type,omitempty"`
	ID                   *string            `json:"id,omitempty"`
	Attributes           *WafRuleAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _WafRule WafRule

// NewWafRule instantiates a new WafRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRule() *WafRule {
	this := WafRule{}
	var resourceType TypeWafRule = TYPEWAFRULE_WAF_RULE
	this.Type = &resourceType
	return &this
}

// NewWafRuleWithDefaults instantiates a new WafRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleWithDefaults() *WafRule {
	this := WafRule{}
	var resourceType TypeWafRule = TYPEWAFRULE_WAF_RULE
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafRule) GetType() TypeWafRule {
	if o == nil || o.Type == nil {
		var ret TypeWafRule
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetTypeOk() (*TypeWafRule, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafRule) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafRule and assigns it to the Type field.
func (o *WafRule) SetType(v TypeWafRule) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WafRule) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WafRule) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WafRule) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafRule) GetAttributes() WafRuleAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafRuleAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRule) GetAttributesOk() (*WafRuleAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafRule) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafRuleAttributes and assigns it to the Attributes field.
func (o *WafRule) SetAttributes(v WafRuleAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WafRule) UnmarshalJSON(bytes []byte) (err error) {
	varWafRule := _WafRule{}

	if err = json.Unmarshal(bytes, &varWafRule); err == nil {
		*o = WafRule(varWafRule)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafRule is a helper abstraction for handling nullable wafrule types.
type NullableWafRule struct {
	value *WafRule
	isSet bool
}

// Get returns the value.
func (v NullableWafRule) Get() *WafRule {
	return v.value
}

// Set modifies the value.
func (v *NullableWafRule) Set(val *WafRule) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafRule) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafRule) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafRule returns a pointer to a new instance of NullableWafRule.
func NewNullableWafRule(val *WafRule) *NullableWafRule {
	return &NullableWafRule{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
