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

// WafExclusionData struct for WafExclusionData
type WafExclusionData struct {
	Type *TypeWafExclusion `json:"type,omitempty"`
	Attributes *WafExclusionDataAttributes `json:"attributes,omitempty"`
	Relationships *RelationshipsForWafExclusion `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafExclusionData WafExclusionData

// NewWafExclusionData instantiates a new WafExclusionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafExclusionData() *WafExclusionData {
	this := WafExclusionData{}
	var resourceType TypeWafExclusion = TYPEWAFEXCLUSION_WAF_EXCLUSION
	this.Type = &resourceType
	return &this
}

// NewWafExclusionDataWithDefaults instantiates a new WafExclusionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafExclusionDataWithDefaults() *WafExclusionData {
	this := WafExclusionData{}
	var resourceType TypeWafExclusion = TYPEWAFEXCLUSION_WAF_EXCLUSION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafExclusionData) GetType() TypeWafExclusion {
	if o == nil || o.Type == nil {
		var ret TypeWafExclusion
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionData) GetTypeOk() (*TypeWafExclusion, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafExclusionData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafExclusion and assigns it to the Type field.
func (o *WafExclusionData) SetType(v TypeWafExclusion) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafExclusionData) GetAttributes() WafExclusionDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafExclusionDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionData) GetAttributesOk() (*WafExclusionDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafExclusionData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafExclusionDataAttributes and assigns it to the Attributes field.
func (o *WafExclusionData) SetAttributes(v WafExclusionDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafExclusionData) GetRelationships() RelationshipsForWafExclusion {
	if o == nil || o.Relationships == nil {
		var ret RelationshipsForWafExclusion
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionData) GetRelationshipsOk() (*RelationshipsForWafExclusion, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafExclusionData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given RelationshipsForWafExclusion and assigns it to the Relationships field.
func (o *WafExclusionData) SetRelationships(v RelationshipsForWafExclusion) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafExclusionData) MarshalJSON() ([]byte, error) {
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
func (o *WafExclusionData) UnmarshalJSON(bytes []byte) (err error) {
	varWafExclusionData := _WafExclusionData{}

	if err = json.Unmarshal(bytes, &varWafExclusionData); err == nil {
		*o = WafExclusionData(varWafExclusionData)
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

// NullableWafExclusionData is a helper abstraction for handling nullable wafexclusiondata types. 
type NullableWafExclusionData struct {
	value *WafExclusionData
	isSet bool
}

// Get returns the value.
func (v NullableWafExclusionData) Get() *WafExclusionData {
	return v.value
}

// Set modifies the value.
func (v *NullableWafExclusionData) Set(val *WafExclusionData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafExclusionData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafExclusionData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafExclusionData returns a pointer to a new instance of NullableWafExclusionData.
func NewNullableWafExclusionData(val *WafExclusionData) *NullableWafExclusionData {
	return &NullableWafExclusionData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafExclusionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafExclusionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
