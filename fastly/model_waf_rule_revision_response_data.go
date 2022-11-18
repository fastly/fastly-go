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

// WafRuleRevisionResponseData struct for WafRuleRevisionResponseData
type WafRuleRevisionResponseData struct {
	Type *TypeWafRuleRevision `json:"type,omitempty"`
	// Alphanumeric string identifying a WAF rule revision.
	ID *string `json:"id,omitempty"`
	Attributes *WafRuleRevisionAttributes `json:"attributes,omitempty"`
	Relationships *RelationshipWafRule `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafRuleRevisionResponseData WafRuleRevisionResponseData

// NewWafRuleRevisionResponseData instantiates a new WafRuleRevisionResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRuleRevisionResponseData() *WafRuleRevisionResponseData {
	this := WafRuleRevisionResponseData{}
	var resourceType TypeWafRuleRevision = TYPEWAFRULEREVISION_WAF_RULE_REVISION
	this.Type = &resourceType
	return &this
}

// NewWafRuleRevisionResponseDataWithDefaults instantiates a new WafRuleRevisionResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleRevisionResponseDataWithDefaults() *WafRuleRevisionResponseData {
	this := WafRuleRevisionResponseData{}
	var resourceType TypeWafRuleRevision = TYPEWAFRULEREVISION_WAF_RULE_REVISION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafRuleRevisionResponseData) GetType() TypeWafRuleRevision {
	if o == nil || o.Type == nil {
		var ret TypeWafRuleRevision
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionResponseData) GetTypeOk() (*TypeWafRuleRevision, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafRuleRevisionResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafRuleRevision and assigns it to the Type field.
func (o *WafRuleRevisionResponseData) SetType(v TypeWafRuleRevision) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WafRuleRevisionResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WafRuleRevisionResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WafRuleRevisionResponseData) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafRuleRevisionResponseData) GetAttributes() WafRuleRevisionAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafRuleRevisionAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionResponseData) GetAttributesOk() (*WafRuleRevisionAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafRuleRevisionResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafRuleRevisionAttributes and assigns it to the Attributes field.
func (o *WafRuleRevisionResponseData) SetAttributes(v WafRuleRevisionAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafRuleRevisionResponseData) GetRelationships() RelationshipWafRule {
	if o == nil || o.Relationships == nil {
		var ret RelationshipWafRule
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleRevisionResponseData) GetRelationshipsOk() (*RelationshipWafRule, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafRuleRevisionResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipWafRule and assigns it to the Relationships field.
func (o *WafRuleRevisionResponseData) SetRelationships(v RelationshipWafRule) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafRuleRevisionResponseData) MarshalJSON() ([]byte, error) {
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
func (o *WafRuleRevisionResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varWafRuleRevisionResponseData := _WafRuleRevisionResponseData{}

	if err = json.Unmarshal(bytes, &varWafRuleRevisionResponseData); err == nil {
		*o = WafRuleRevisionResponseData(varWafRuleRevisionResponseData)
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

// NullableWafRuleRevisionResponseData is a helper abstraction for handling nullable wafrulerevisionresponsedata types. 
type NullableWafRuleRevisionResponseData struct {
	value *WafRuleRevisionResponseData
	isSet bool
}

// Get returns the value.
func (v NullableWafRuleRevisionResponseData) Get() *WafRuleRevisionResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableWafRuleRevisionResponseData) Set(val *WafRuleRevisionResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafRuleRevisionResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafRuleRevisionResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafRuleRevisionResponseData returns a pointer to a new instance of NullableWafRuleRevisionResponseData.
func NewNullableWafRuleRevisionResponseData(val *WafRuleRevisionResponseData) *NullableWafRuleRevisionResponseData {
	return &NullableWafRuleRevisionResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafRuleRevisionResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafRuleRevisionResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
