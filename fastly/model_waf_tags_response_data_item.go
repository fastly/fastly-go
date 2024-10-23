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

// WafTagsResponseDataItem struct for WafTagsResponseDataItem
type WafTagsResponseDataItem struct {
	Type *TypeWafTag `json:"type,omitempty"`
	// Alphanumeric string identifying a WAF tag.
	ID                   *string              `json:"id,omitempty"`
	Attributes           *WafTagAttributes    `json:"attributes,omitempty"`
	Relationships        *RelationshipWafRule `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafTagsResponseDataItem WafTagsResponseDataItem

// NewWafTagsResponseDataItem instantiates a new WafTagsResponseDataItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafTagsResponseDataItem() *WafTagsResponseDataItem {
	this := WafTagsResponseDataItem{}
	var resourceType TypeWafTag = TYPEWAFTAG_WAF_TAG
	this.Type = &resourceType
	return &this
}

// NewWafTagsResponseDataItemWithDefaults instantiates a new WafTagsResponseDataItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafTagsResponseDataItemWithDefaults() *WafTagsResponseDataItem {
	this := WafTagsResponseDataItem{}
	var resourceType TypeWafTag = TYPEWAFTAG_WAF_TAG
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafTagsResponseDataItem) GetType() TypeWafTag {
	if o == nil || o.Type == nil {
		var ret TypeWafTag
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTagsResponseDataItem) GetTypeOk() (*TypeWafTag, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafTagsResponseDataItem) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafTag and assigns it to the Type field.
func (o *WafTagsResponseDataItem) SetType(v TypeWafTag) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WafTagsResponseDataItem) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTagsResponseDataItem) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WafTagsResponseDataItem) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WafTagsResponseDataItem) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafTagsResponseDataItem) GetAttributes() WafTagAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafTagAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTagsResponseDataItem) GetAttributesOk() (*WafTagAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafTagsResponseDataItem) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafTagAttributes and assigns it to the Attributes field.
func (o *WafTagsResponseDataItem) SetAttributes(v WafTagAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafTagsResponseDataItem) GetRelationships() RelationshipWafRule {
	if o == nil || o.Relationships == nil {
		var ret RelationshipWafRule
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTagsResponseDataItem) GetRelationshipsOk() (*RelationshipWafRule, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafTagsResponseDataItem) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipWafRule and assigns it to the Relationships field.
func (o *WafTagsResponseDataItem) SetRelationships(v RelationshipWafRule) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafTagsResponseDataItem) MarshalJSON() ([]byte, error) {
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
func (o *WafTagsResponseDataItem) UnmarshalJSON(bytes []byte) (err error) {
	varWafTagsResponseDataItem := _WafTagsResponseDataItem{}

	if err = json.Unmarshal(bytes, &varWafTagsResponseDataItem); err == nil {
		*o = WafTagsResponseDataItem(varWafTagsResponseDataItem)
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

// NullableWafTagsResponseDataItem is a helper abstraction for handling nullable waftagsresponsedataitem types.
type NullableWafTagsResponseDataItem struct {
	value *WafTagsResponseDataItem
	isSet bool
}

// Get returns the value.
func (v NullableWafTagsResponseDataItem) Get() *WafTagsResponseDataItem {
	return v.value
}

// Set modifies the value.
func (v *NullableWafTagsResponseDataItem) Set(val *WafTagsResponseDataItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafTagsResponseDataItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafTagsResponseDataItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafTagsResponseDataItem returns a pointer to a new instance of NullableWafTagsResponseDataItem.
func NewNullableWafTagsResponseDataItem(val *WafTagsResponseDataItem) *NullableWafTagsResponseDataItem {
	return &NullableWafTagsResponseDataItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafTagsResponseDataItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafTagsResponseDataItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
