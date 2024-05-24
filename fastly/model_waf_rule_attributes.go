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

// WafRuleAttributes struct for WafRuleAttributes
type WafRuleAttributes struct {
	// Corresponding ModSecurity rule ID.
	ModsecRuleID *int32 `json:"modsec_rule_id,omitempty"`
	// Rule publisher.
	Publisher *string `json:"publisher,omitempty"`
	// The rule's [type](https://docs.fastly.com/en/guides/managing-rules-on-the-fastly-waf#understanding-the-types-of-rules).
	Type *string `json:"type,omitempty"`
	AdditionalProperties map[string]any
}

type _WafRuleAttributes WafRuleAttributes

// NewWafRuleAttributes instantiates a new WafRuleAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafRuleAttributes() *WafRuleAttributes {
	this := WafRuleAttributes{}
	return &this
}

// NewWafRuleAttributesWithDefaults instantiates a new WafRuleAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafRuleAttributesWithDefaults() *WafRuleAttributes {
	this := WafRuleAttributes{}
	return &this
}

// GetModsecRuleID returns the ModsecRuleID field value if set, zero value otherwise.
func (o *WafRuleAttributes) GetModsecRuleID() int32 {
	if o == nil || o.ModsecRuleID == nil {
		var ret int32
		return ret
	}
	return *o.ModsecRuleID
}

// GetModsecRuleIDOk returns a tuple with the ModsecRuleID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleAttributes) GetModsecRuleIDOk() (*int32, bool) {
	if o == nil || o.ModsecRuleID == nil {
		return nil, false
	}
	return o.ModsecRuleID, true
}

// HasModsecRuleID returns a boolean if a field has been set.
func (o *WafRuleAttributes) HasModsecRuleID() bool {
	if o != nil && o.ModsecRuleID != nil {
		return true
	}

	return false
}

// SetModsecRuleID gets a reference to the given int32 and assigns it to the ModsecRuleID field.
func (o *WafRuleAttributes) SetModsecRuleID(v int32) {
	o.ModsecRuleID = &v
}

// GetPublisher returns the Publisher field value if set, zero value otherwise.
func (o *WafRuleAttributes) GetPublisher() string {
	if o == nil || o.Publisher == nil {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleAttributes) GetPublisherOk() (*string, bool) {
	if o == nil || o.Publisher == nil {
		return nil, false
	}
	return o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *WafRuleAttributes) HasPublisher() bool {
	if o != nil && o.Publisher != nil {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *WafRuleAttributes) SetPublisher(v string) {
	o.Publisher = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafRuleAttributes) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafRuleAttributes) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafRuleAttributes) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WafRuleAttributes) SetType(v string) {
	o.Type = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ModsecRuleID != nil {
		toSerialize["modsec_rule_id"] = o.ModsecRuleID
	}
	if o.Publisher != nil {
		toSerialize["publisher"] = o.Publisher
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *WafRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	varWafRuleAttributes := _WafRuleAttributes{}

	if err = json.Unmarshal(bytes, &varWafRuleAttributes); err == nil {
		*o = WafRuleAttributes(varWafRuleAttributes)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "modsec_rule_id")
		delete(additionalProperties, "publisher")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafRuleAttributes is a helper abstraction for handling nullable wafruleattributes types. 
type NullableWafRuleAttributes struct {
	value *WafRuleAttributes
	isSet bool
}

// Get returns the value.
func (v NullableWafRuleAttributes) Get() *WafRuleAttributes {
	return v.value
}

// Set modifies the value.
func (v *NullableWafRuleAttributes) Set(val *WafRuleAttributes) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafRuleAttributes) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafRuleAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafRuleAttributes returns a pointer to a new instance of NullableWafRuleAttributes.
func NewNullableWafRuleAttributes(val *WafRuleAttributes) *NullableWafRuleAttributes {
	return &NullableWafRuleAttributes{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafRuleAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafRuleAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
