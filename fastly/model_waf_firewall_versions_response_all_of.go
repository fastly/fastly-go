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

// WafFirewallVersionsResponseAllOf struct for WafFirewallVersionsResponseAllOf
type WafFirewallVersionsResponseAllOf struct {
	Data []WafFirewallVersionResponseData `json:"data,omitempty"`
	Included []IncludedWithWafFirewallVersionItem `json:"included,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallVersionsResponseAllOf WafFirewallVersionsResponseAllOf

// NewWafFirewallVersionsResponseAllOf instantiates a new WafFirewallVersionsResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallVersionsResponseAllOf() *WafFirewallVersionsResponseAllOf {
	this := WafFirewallVersionsResponseAllOf{}
	return &this
}

// NewWafFirewallVersionsResponseAllOfWithDefaults instantiates a new WafFirewallVersionsResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallVersionsResponseAllOfWithDefaults() *WafFirewallVersionsResponseAllOf {
	this := WafFirewallVersionsResponseAllOf{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafFirewallVersionsResponseAllOf) GetData() []WafFirewallVersionResponseData {
	if o == nil || o.Data == nil {
		var ret []WafFirewallVersionResponseData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionsResponseAllOf) GetDataOk() ([]WafFirewallVersionResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafFirewallVersionsResponseAllOf) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []WafFirewallVersionResponseData and assigns it to the Data field.
func (o *WafFirewallVersionsResponseAllOf) SetData(v []WafFirewallVersionResponseData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *WafFirewallVersionsResponseAllOf) GetIncluded() []IncludedWithWafFirewallVersionItem {
	if o == nil || o.Included == nil {
		var ret []IncludedWithWafFirewallVersionItem
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionsResponseAllOf) GetIncludedOk() ([]IncludedWithWafFirewallVersionItem, bool) {
	if o == nil || o.Included == nil {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *WafFirewallVersionsResponseAllOf) HasIncluded() bool {
	if o != nil && o.Included != nil {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []IncludedWithWafFirewallVersionItem and assigns it to the Included field.
func (o *WafFirewallVersionsResponseAllOf) SetIncluded(v []IncludedWithWafFirewallVersionItem) {
	o.Included = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallVersionsResponseAllOf) MarshalJSON() ([]byte, error) {
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
func (o *WafFirewallVersionsResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallVersionsResponseAllOf := _WafFirewallVersionsResponseAllOf{}

	if err = json.Unmarshal(bytes, &varWafFirewallVersionsResponseAllOf); err == nil {
		*o = WafFirewallVersionsResponseAllOf(varWafFirewallVersionsResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "included")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallVersionsResponseAllOf is a helper abstraction for handling nullable waffirewallversionsresponseallof types. 
type NullableWafFirewallVersionsResponseAllOf struct {
	value *WafFirewallVersionsResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallVersionsResponseAllOf) Get() *WafFirewallVersionsResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallVersionsResponseAllOf) Set(val *WafFirewallVersionsResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallVersionsResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallVersionsResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallVersionsResponseAllOf returns a pointer to a new instance of NullableWafFirewallVersionsResponseAllOf.
func NewNullableWafFirewallVersionsResponseAllOf(val *WafFirewallVersionsResponseAllOf) *NullableWafFirewallVersionsResponseAllOf {
	return &NullableWafFirewallVersionsResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallVersionsResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafFirewallVersionsResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
