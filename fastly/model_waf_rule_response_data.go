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

// WafRuleResponseData struct for WafRuleResponseData
type WafRuleResponseData struct {
	Type *TypeWafRule `json:"type,omitempty"`
	ID *string `json:"id,omitempty"`
	Attributes *WafRuleAttributes `json:"attributes,omitempty"`
	Relationships *RelationshipsForWafRule `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafRuleResponseData WafRuleResponseData

// NewWafRuleResponseData instantiates a new WafRuleResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRuleResponseData() *WafRuleResponseData {
	this := WafRuleResponseData{}
	var resourceType TypeWafRule = TYPEWAFRULE_WAF_RULE
	this.Type = &resourceType
	return &this
}

// NewWafRuleResponseDataWithDefaults instantiates a new WafRuleResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleResponseDataWithDefaults() *WafRuleResponseData {
	this := WafRuleResponseData{}
	var resourceType TypeWafRule = TYPEWAFRULE_WAF_RULE
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafRuleResponseData) GetType() TypeWafRule {
	if o == nil || o.Type == nil {
		var ret TypeWafRule
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleResponseData) GetTypeOk() (*TypeWafRule, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafRuleResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafRule and assigns it to the Type field.
func (o *WafRuleResponseData) SetType(v TypeWafRule) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WafRuleResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WafRuleResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WafRuleResponseData) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafRuleResponseData) GetAttributes() WafRuleAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafRuleAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleResponseData) GetAttributesOk() (*WafRuleAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafRuleResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafRuleAttributes and assigns it to the Attributes field.
func (o *WafRuleResponseData) SetAttributes(v WafRuleAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafRuleResponseData) GetRelationships() RelationshipsForWafRule {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForWafRule
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleResponseData) GetRelationshipsOk() (*RelationshipsForWafRule, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafRuleResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForWafRule and assigns it to the Relationships field.
func (o *WafRuleResponseData) SetRelationships(v RelationshipsForWafRule) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafRuleResponseData) MarshalJSON() ([]byte, error) {
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
func (o *WafRuleResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varWafRuleResponseData := _WafRuleResponseData{}

	if err = json.Unmarshal(bytes, &varWafRuleResponseData); err == nil {
		*o = WafRuleResponseData(varWafRuleResponseData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafRuleResponseData is a helper abstraction for handling nullable wafruleresponsedata types. 
type NullableWafRuleResponseData struct {
	value *WafRuleResponseData
	isSet bool
}

// Get returns the value.
func (v NullableWafRuleResponseData) Get() *WafRuleResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableWafRuleResponseData) Set(val *WafRuleResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafRuleResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafRuleResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafRuleResponseData returns a pointer to a new instance of NullableWafRuleResponseData.
func NewNullableWafRuleResponseData(val *WafRuleResponseData) *NullableWafRuleResponseData {
	return &NullableWafRuleResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafRuleResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafRuleResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
