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

// WafExclusionResponseData struct for WafExclusionResponseData
type WafExclusionResponseData struct {
	Type *TypeWafExclusion `json:"type,omitempty"`
	Attributes *WafExclusionResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *WafExclusionResponseDataRelationships `json:"relationships,omitempty"`
	// Alphanumeric string identifying a WAF exclusion.
	ID *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _WafExclusionResponseData WafExclusionResponseData

// NewWafExclusionResponseData instantiates a new WafExclusionResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafExclusionResponseData() *WafExclusionResponseData {
	this := WafExclusionResponseData{}
	var resourceType TypeWafExclusion = TYPEWAFEXCLUSION_WAF_EXCLUSION
	this.Type = &resourceType
	return &this
}

// NewWafExclusionResponseDataWithDefaults instantiates a new WafExclusionResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafExclusionResponseDataWithDefaults() *WafExclusionResponseData {
	this := WafExclusionResponseData{}
	var resourceType TypeWafExclusion = TYPEWAFEXCLUSION_WAF_EXCLUSION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafExclusionResponseData) GetType() TypeWafExclusion {
	if o == nil || o.Type == nil {
		var ret TypeWafExclusion
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseData) GetTypeOk() (*TypeWafExclusion, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafExclusionResponseData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafExclusion and assigns it to the Type field.
func (o *WafExclusionResponseData) SetType(v TypeWafExclusion) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafExclusionResponseData) GetAttributes() WafExclusionResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafExclusionResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseData) GetAttributesOk() (*WafExclusionResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafExclusionResponseData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafExclusionResponseDataAttributes and assigns it to the Attributes field.
func (o *WafExclusionResponseData) SetAttributes(v WafExclusionResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafExclusionResponseData) GetRelationships() WafExclusionResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret WafExclusionResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseData) GetRelationshipsOk() (*WafExclusionResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafExclusionResponseData) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given WafExclusionResponseDataRelationships and assigns it to the Relationships field.
func (o *WafExclusionResponseData) SetRelationships(v WafExclusionResponseDataRelationships) {
	o.Relationships = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WafExclusionResponseData) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseData) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WafExclusionResponseData) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WafExclusionResponseData) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafExclusionResponseData) MarshalJSON() ([]byte, error) {
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
func (o *WafExclusionResponseData) UnmarshalJSON(bytes []byte) (err error) {
	varWafExclusionResponseData := _WafExclusionResponseData{}

	if err = json.Unmarshal(bytes, &varWafExclusionResponseData); err == nil {
		*o = WafExclusionResponseData(varWafExclusionResponseData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "relationships")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafExclusionResponseData is a helper abstraction for handling nullable wafexclusionresponsedata types. 
type NullableWafExclusionResponseData struct {
	value *WafExclusionResponseData
	isSet bool
}

// Get returns the value.
func (v NullableWafExclusionResponseData) Get() *WafExclusionResponseData {
	return v.value
}

// Set modifies the value.
func (v *NullableWafExclusionResponseData) Set(val *WafExclusionResponseData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafExclusionResponseData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafExclusionResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafExclusionResponseData returns a pointer to a new instance of NullableWafExclusionResponseData.
func NewNullableWafExclusionResponseData(val *WafExclusionResponseData) *NullableWafExclusionResponseData {
	return &NullableWafExclusionResponseData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafExclusionResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafExclusionResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
