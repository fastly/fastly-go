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

// WafActiveRulesResponseAllOf struct for WafActiveRulesResponseAllOf
type WafActiveRulesResponseAllOf struct {
	Data []WafActiveRuleResponseData `json:"data,omitempty"`
	Included []IncludedWithWafActiveRuleItem `json:"included,omitempty"`
	AdditionalProperties map[string]any
}

type _WafActiveRulesResponseAllOf WafActiveRulesResponseAllOf

// NewWafActiveRulesResponseAllOf instantiates a new WafActiveRulesResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafActiveRulesResponseAllOf() *WafActiveRulesResponseAllOf {
	this := WafActiveRulesResponseAllOf{}
	return &this
}

// NewWafActiveRulesResponseAllOfWithDefaults instantiates a new WafActiveRulesResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafActiveRulesResponseAllOfWithDefaults() *WafActiveRulesResponseAllOf {
	this := WafActiveRulesResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafActiveRulesResponseAllOf) GetData() []WafActiveRuleResponseData {
	if o == nil || o.Data == nil {
		var ret []WafActiveRuleResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRulesResponseAllOf) GetDataOk() ([]WafActiveRuleResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafActiveRulesResponseAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []WafActiveRuleResponseData and assigns it to the Data field.
func (o *WafActiveRulesResponseAllOf) SetData(v []WafActiveRuleResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *WafActiveRulesResponseAllOf) GetIncluded() []IncludedWithWafActiveRuleItem {
	if o == nil || o.Included == nil {
		var ret []IncludedWithWafActiveRuleItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRulesResponseAllOf) GetIncludedOk() ([]IncludedWithWafActiveRuleItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *WafActiveRulesResponseAllOf) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []IncludedWithWafActiveRuleItem and assigns it to the Included field.
func (o *WafActiveRulesResponseAllOf) SetIncluded(v []IncludedWithWafActiveRuleItem) {
	o.Included = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafActiveRulesResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Included != nil {
		toSerialize["included"] = o.Included
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *WafActiveRulesResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafActiveRulesResponseAllOf := _WafActiveRulesResponseAllOf{}

	if err = json.Unmarshal(bytes, &varWafActiveRulesResponseAllOf); err == nil {
		*o = WafActiveRulesResponseAllOf(varWafActiveRulesResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafActiveRulesResponseAllOf is a helper abstraction for handling nullable wafactiverulesresponseallof types. 
type NullableWafActiveRulesResponseAllOf struct {
	value *WafActiveRulesResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafActiveRulesResponseAllOf) Get() *WafActiveRulesResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafActiveRulesResponseAllOf) Set(val *WafActiveRulesResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafActiveRulesResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafActiveRulesResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafActiveRulesResponseAllOf returns a pointer to a new instance of NullableWafActiveRulesResponseAllOf.
func NewNullableWafActiveRulesResponseAllOf(val *WafActiveRulesResponseAllOf) *NullableWafActiveRulesResponseAllOf {
	return &NullableWafActiveRulesResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafActiveRulesResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafActiveRulesResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
