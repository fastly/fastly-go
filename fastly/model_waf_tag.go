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

// WafTag struct for WafTag
type WafTag struct {
	Type *TypeWafTag `json:"type,omitempty"`
	// Alphanumeric string identifying a WAF tag.
	ID *string `json:"id,omitempty"`
	Attributes *WafTagAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _WafTag WafTag

// NewWafTag instantiates a new WafTag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafTag() *WafTag {
	this := WafTag{}
	var resourceType TypeWafTag = TYPEWAFTAG_WAF_TAG
	this.Type = &resourceType
	return &this
}

// NewWafTagWithDefaults instantiates a new WafTag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafTagWithDefaults() *WafTag {
	this := WafTag{}
	var resourceType TypeWafTag = TYPEWAFTAG_WAF_TAG
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafTag) GetType() TypeWafTag {
	if o == nil || o.Type == nil {
		var ret TypeWafTag
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTag) GetTypeOk() (*TypeWafTag, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafTag) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafTag and assigns it to the Type field.
func (o *WafTag) SetType(v TypeWafTag) {
	o.Type = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *WafTag) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTag) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *WafTag) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *WafTag) SetID(v string) {
	o.ID = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafTag) GetAttributes() WafTagAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafTagAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafTag) GetAttributesOk() (*WafTagAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafTag) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafTagAttributes and assigns it to the Attributes field.
func (o *WafTag) SetAttributes(v WafTagAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafTag) MarshalJSON() ([]byte, error) {
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
func (o *WafTag) UnmarshalJSON(bytes []byte) (err error) {
	varWafTag := _WafTag{}

	if err = json.Unmarshal(bytes, &varWafTag); err == nil {
		*o = WafTag(varWafTag)
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

// NullableWafTag is a helper abstraction for handling nullable waftag types. 
type NullableWafTag struct {
	value *WafTag
	isSet bool
}

// Get returns the value.
func (v NullableWafTag) Get() *WafTag {
	return v.value
}

// Set modifies the value.
func (v *NullableWafTag) Set(val *WafTag) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafTag) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafTag) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafTag returns a pointer to a new instance of NullableWafTag.
func NewNullableWafTag(val *WafTag) *NullableWafTag {
	return &NullableWafTag{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
