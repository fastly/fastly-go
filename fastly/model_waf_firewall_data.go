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

// WafFirewallData struct for WafFirewallData
type WafFirewallData struct {
	Type *TypeWafFirewall `json:"type,omitempty"`
	Attributes *WafFirewallDataAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallData WafFirewallData

// NewWafFirewallData instantiates a new WafFirewallData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallData() *WafFirewallData {
	this := WafFirewallData{}
	var resourceType TypeWafFirewall = TYPEWAFFIREWALL_WAF_FIREWALL
	this.Type = &resourceType
	return &this
}

// NewWafFirewallDataWithDefaults instantiates a new WafFirewallData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallDataWithDefaults() *WafFirewallData {
	this := WafFirewallData{}
	var resourceType TypeWafFirewall = TYPEWAFFIREWALL_WAF_FIREWALL
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafFirewallData) GetType() TypeWafFirewall {
	if o == nil || o.Type == nil {
		var ret TypeWafFirewall
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallData) GetTypeOk() (*TypeWafFirewall, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafFirewallData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafFirewall and assigns it to the Type field.
func (o *WafFirewallData) SetType(v TypeWafFirewall) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafFirewallData) GetAttributes() WafFirewallDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafFirewallDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallData) GetAttributesOk() (*WafFirewallDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafFirewallData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafFirewallDataAttributes and assigns it to the Attributes field.
func (o *WafFirewallData) SetAttributes(v WafFirewallDataAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
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
func (o *WafFirewallData) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallData := _WafFirewallData{}

	if err = json.Unmarshal(bytes, &varWafFirewallData); err == nil {
		*o = WafFirewallData(varWafFirewallData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallData is a helper abstraction for handling nullable waffirewalldata types. 
type NullableWafFirewallData struct {
	value *WafFirewallData
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallData) Get() *WafFirewallData {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallData) Set(val *WafFirewallData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallData returns a pointer to a new instance of NullableWafFirewallData.
func NewNullableWafFirewallData(val *WafFirewallData) *NullableWafFirewallData {
	return &NullableWafFirewallData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafFirewallData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
