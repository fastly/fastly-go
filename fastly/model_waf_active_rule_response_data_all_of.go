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

// WafActiveRuleResponseDataAllOf struct for WafActiveRuleResponseDataAllOf
type WafActiveRuleResponseDataAllOf struct {
	ID *string `json:"id,omitempty"`
	Attributes *WafActiveRuleResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *WafActiveRuleResponseDataRelationships `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafActiveRuleResponseDataAllOf WafActiveRuleResponseDataAllOf

// NewWafActiveRuleResponseDataAllOf instantiates a new WafActiveRuleResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafActiveRuleResponseDataAllOf() *WafActiveRuleResponseDataAllOf {
	this := WafActiveRuleResponseDataAllOf{}
	return &this
}

// NewWafActiveRuleResponseDataAllOfWithDefaults instantiates a new WafActiveRuleResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafActiveRuleResponseDataAllOfWithDefaults() *WafActiveRuleResponseDataAllOf {
	this := WafActiveRuleResponseDataAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WafActiveRuleResponseDataAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleResponseDataAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WafActiveRuleResponseDataAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WafActiveRuleResponseDataAllOf) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafActiveRuleResponseDataAllOf) GetAttributes() WafActiveRuleResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafActiveRuleResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleResponseDataAllOf) GetAttributesOk() (*WafActiveRuleResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafActiveRuleResponseDataAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafActiveRuleResponseDataAttributes and assigns it to the Attributes field.
func (o *WafActiveRuleResponseDataAllOf) SetAttributes(v WafActiveRuleResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafActiveRuleResponseDataAllOf) GetRelationships() WafActiveRuleResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret WafActiveRuleResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleResponseDataAllOf) GetRelationshipsOk() (*WafActiveRuleResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafActiveRuleResponseDataAllOf) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given WafActiveRuleResponseDataRelationships and assigns it to the Relationships field.
func (o *WafActiveRuleResponseDataAllOf) SetRelationships(v WafActiveRuleResponseDataRelationships) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafActiveRuleResponseDataAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
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
func (o *WafActiveRuleResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafActiveRuleResponseDataAllOf := _WafActiveRuleResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varWafActiveRuleResponseDataAllOf); err == nil {
		*o = WafActiveRuleResponseDataAllOf(varWafActiveRuleResponseDataAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafActiveRuleResponseDataAllOf is a helper abstraction for handling nullable wafactiveruleresponsedataallof types. 
type NullableWafActiveRuleResponseDataAllOf struct {
	value *WafActiveRuleResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafActiveRuleResponseDataAllOf) Get() *WafActiveRuleResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafActiveRuleResponseDataAllOf) Set(val *WafActiveRuleResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafActiveRuleResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafActiveRuleResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafActiveRuleResponseDataAllOf returns a pointer to a new instance of NullableWafActiveRuleResponseDataAllOf.
func NewNullableWafActiveRuleResponseDataAllOf(val *WafActiveRuleResponseDataAllOf) *NullableWafActiveRuleResponseDataAllOf {
	return &NullableWafActiveRuleResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafActiveRuleResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafActiveRuleResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
