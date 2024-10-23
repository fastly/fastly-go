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

// WafExclusionsResponseAllOf struct for WafExclusionsResponseAllOf
type WafExclusionsResponseAllOf struct {
	Data                 []WafExclusionResponseData     `json:"data,omitempty"`
	Included             []IncludedWithWafExclusionItem `json:"included,omitempty"`
	AdditionalProperties map[string]any
}

type _WafExclusionsResponseAllOf WafExclusionsResponseAllOf

// NewWafExclusionsResponseAllOf instantiates a new WafExclusionsResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafExclusionsResponseAllOf() *WafExclusionsResponseAllOf {
	this := WafExclusionsResponseAllOf{}
	return &this
}

// NewWafExclusionsResponseAllOfWithDefaults instantiates a new WafExclusionsResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafExclusionsResponseAllOfWithDefaults() *WafExclusionsResponseAllOf {
	this := WafExclusionsResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafExclusionsResponseAllOf) GetData() []WafExclusionResponseData {
	if o == nil || o.Data == nil {
		var ret []WafExclusionResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionsResponseAllOf) GetDataOk() ([]WafExclusionResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafExclusionsResponseAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []WafExclusionResponseData and assigns it to the Data field.
func (o *WafExclusionsResponseAllOf) SetData(v []WafExclusionResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *WafExclusionsResponseAllOf) GetIncluded() []IncludedWithWafExclusionItem {
	if o == nil || o.Included == nil {
		var ret []IncludedWithWafExclusionItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafExclusionsResponseAllOf) GetIncludedOk() ([]IncludedWithWafExclusionItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *WafExclusionsResponseAllOf) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []IncludedWithWafExclusionItem and assigns it to the Included field.
func (o *WafExclusionsResponseAllOf) SetIncluded(v []IncludedWithWafExclusionItem) {
	o.Included = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafExclusionsResponseAllOf) MarshalJSON() ([]byte, error) {
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
func (o *WafExclusionsResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafExclusionsResponseAllOf := _WafExclusionsResponseAllOf{}

	if err = json.Unmarshal(bytes, &varWafExclusionsResponseAllOf); err == nil {
		*o = WafExclusionsResponseAllOf(varWafExclusionsResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafExclusionsResponseAllOf is a helper abstraction for handling nullable wafexclusionsresponseallof types.
type NullableWafExclusionsResponseAllOf struct {
	value *WafExclusionsResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafExclusionsResponseAllOf) Get() *WafExclusionsResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafExclusionsResponseAllOf) Set(val *WafExclusionsResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafExclusionsResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafExclusionsResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafExclusionsResponseAllOf returns a pointer to a new instance of NullableWafExclusionsResponseAllOf.
func NewNullableWafExclusionsResponseAllOf(val *WafExclusionsResponseAllOf) *NullableWafExclusionsResponseAllOf {
	return &NullableWafExclusionsResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafExclusionsResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafExclusionsResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
