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

// WafActiveRuleData struct for WafActiveRuleData
type WafActiveRuleData struct {
	Type *TypeWafActiveRule `json:"type,omitempty"`
	Attributes *WafActiveRuleDataAttributes `json:"attributes,omitempty"`
	Relationships *RelationshipsForWafActiveRule `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafActiveRuleData WafActiveRuleData

// NewWafActiveRuleData instantiates a new WafActiveRuleData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafActiveRuleData() *WafActiveRuleData {
	this := WafActiveRuleData{}
	var resourceType TypeWafActiveRule = TYPEWAFACTIVERULE_WAF_ACTIVE_RULE
	this.Type = &resourceType
	return &this
}

// NewWafActiveRuleDataWithDefaults instantiates a new WafActiveRuleData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafActiveRuleDataWithDefaults() *WafActiveRuleData {
	this := WafActiveRuleData{}
	var resourceType TypeWafActiveRule = TYPEWAFACTIVERULE_WAF_ACTIVE_RULE
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafActiveRuleData) GetType() TypeWafActiveRule {
	if o == nil || o.Type == nil {
		var ret TypeWafActiveRule
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleData) GetTypeOk() (*TypeWafActiveRule, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafActiveRuleData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafActiveRule and assigns it to the Type field.
func (o *WafActiveRuleData) SetType(v TypeWafActiveRule) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafActiveRuleData) GetAttributes() WafActiveRuleDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafActiveRuleDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleData) GetAttributesOk() (*WafActiveRuleDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafActiveRuleData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafActiveRuleDataAttributes and assigns it to the Attributes field.
func (o *WafActiveRuleData) SetAttributes(v WafActiveRuleDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafActiveRuleData) GetRelationships() RelationshipsForWafActiveRule {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForWafActiveRule
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleData) GetRelationshipsOk() (*RelationshipsForWafActiveRule, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafActiveRuleData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForWafActiveRule and assigns it to the Relationships field.
func (o *WafActiveRuleData) SetRelationships(v RelationshipsForWafActiveRule) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafActiveRuleData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
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
func (o *WafActiveRuleData) UnmarshalJSON(bytes []byte) (err error) {
	varWafActiveRuleData := _WafActiveRuleData{}

	if err = json.Unmarshal(bytes, &varWafActiveRuleData); err == nil {
		*o = WafActiveRuleData(varWafActiveRuleData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafActiveRuleData is a helper abstraction for handling nullable wafactiveruledata types. 
type NullableWafActiveRuleData struct {
	value *WafActiveRuleData
	isSet bool
}

// Get returns the value.
func (v NullableWafActiveRuleData) Get() *WafActiveRuleData {
	return v.value
}

// Set modifies the value.
func (v *NullableWafActiveRuleData) Set(val *WafActiveRuleData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafActiveRuleData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafActiveRuleData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafActiveRuleData returns a pointer to a new instance of NullableWafActiveRuleData.
func NewNullableWafActiveRuleData(val *WafActiveRuleData) *NullableWafActiveRuleData {
	return &NullableWafActiveRuleData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafActiveRuleData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafActiveRuleData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
