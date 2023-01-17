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

// WafExclusionResponseDataAllOf struct for WafExclusionResponseDataAllOf
type WafExclusionResponseDataAllOf struct {
	// Alphanumeric string identifying a WAF exclusion.
	ID *string `json:"id,omitempty"`
	Attributes *WafExclusionResponseDataAttributes `json:"attributes,omitempty"`
	Relationships *WafExclusionResponseDataRelationships `json:"relationships,omitempty"`
	AdditionalProperties map[string]any
}

type _WafExclusionResponseDataAllOf WafExclusionResponseDataAllOf

// NewWafExclusionResponseDataAllOf instantiates a new WafExclusionResponseDataAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafExclusionResponseDataAllOf() *WafExclusionResponseDataAllOf {
	this := WafExclusionResponseDataAllOf{}
	return &this
}

// NewWafExclusionResponseDataAllOfWithDefaults instantiates a new WafExclusionResponseDataAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafExclusionResponseDataAllOfWithDefaults() *WafExclusionResponseDataAllOf {
	this := WafExclusionResponseDataAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WafExclusionResponseDataAllOf) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAllOf) GetAttributes() WafExclusionResponseDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafExclusionResponseDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAllOf) GetAttributesOk() (*WafExclusionResponseDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAllOf) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafExclusionResponseDataAttributes and assigns it to the Attributes field.
func (o *WafExclusionResponseDataAllOf) SetAttributes(v WafExclusionResponseDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *WafExclusionResponseDataAllOf) GetRelationships() WafExclusionResponseDataRelationships {
	if o == nil || o.Relationships == nil {
		var ret WafExclusionResponseDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionResponseDataAllOf) GetRelationshipsOk() (*WafExclusionResponseDataRelationships, bool) {
	if o == nil || o.Relationships == nil {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *WafExclusionResponseDataAllOf) HasRelationships() bool {
	if o != nil && o.Relationships != nil {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given WafExclusionResponseDataRelationships and assigns it to the Relationships field.
func (o *WafExclusionResponseDataAllOf) SetRelationships(v WafExclusionResponseDataRelationships) {
	o.Relationships = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafExclusionResponseDataAllOf) MarshalJSON() ([]byte, error) {
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
func (o *WafExclusionResponseDataAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafExclusionResponseDataAllOf := _WafExclusionResponseDataAllOf{}

	if err = json.Unmarshal(bytes, &varWafExclusionResponseDataAllOf); err == nil {
		*o = WafExclusionResponseDataAllOf(varWafExclusionResponseDataAllOf)
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

// NullableWafExclusionResponseDataAllOf is a helper abstraction for handling nullable wafexclusionresponsedataallof types. 
type NullableWafExclusionResponseDataAllOf struct {
	value *WafExclusionResponseDataAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafExclusionResponseDataAllOf) Get() *WafExclusionResponseDataAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafExclusionResponseDataAllOf) Set(val *WafExclusionResponseDataAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafExclusionResponseDataAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafExclusionResponseDataAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafExclusionResponseDataAllOf returns a pointer to a new instance of NullableWafExclusionResponseDataAllOf.
func NewNullableWafExclusionResponseDataAllOf(val *WafExclusionResponseDataAllOf) *NullableWafExclusionResponseDataAllOf {
	return &NullableWafExclusionResponseDataAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafExclusionResponseDataAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafExclusionResponseDataAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
