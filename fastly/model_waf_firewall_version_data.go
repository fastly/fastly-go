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

// WafFirewallVersionData struct for WafFirewallVersionData
type WafFirewallVersionData struct {
	Type *TypeWafFirewallVersion `json:"type,omitempty"`
	Attributes *WafFirewallVersionDataAttributes `json:"attributes,omitempty"`
	AdditionalProperties map[string]any
}

type _WafFirewallVersionData WafFirewallVersionData

// NewWafFirewallVersionData instantiates a new WafFirewallVersionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafFirewallVersionData() *WafFirewallVersionData {
	this := WafFirewallVersionData{}
	var resourceType TypeWafFirewallVersion = TYPEWAFFIREWALLVERSION_WAF_FIREWALL_VERSION
	this.Type = &resourceType
	return &this
}

// NewWafFirewallVersionDataWithDefaults instantiates a new WafFirewallVersionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafFirewallVersionDataWithDefaults() *WafFirewallVersionData {
	this := WafFirewallVersionData{}
	var resourceType TypeWafFirewallVersion = TYPEWAFFIREWALLVERSION_WAF_FIREWALL_VERSION
	this.Type = &resourceType
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WafFirewallVersionData) GetType() TypeWafFirewallVersion {
	if o == nil || o.Type == nil {
		var ret TypeWafFirewallVersion
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionData) GetTypeOk() (*TypeWafFirewallVersion, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WafFirewallVersionData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given TypeWafFirewallVersion and assigns it to the Type field.
func (o *WafFirewallVersionData) SetType(v TypeWafFirewallVersion) {
	o.Type = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *WafFirewallVersionData) GetAttributes() WafFirewallVersionDataAttributes {
	if o == nil || o.Attributes == nil {
		var ret WafFirewallVersionDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafFirewallVersionData) GetAttributesOk() (*WafFirewallVersionDataAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *WafFirewallVersionData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given WafFirewallVersionDataAttributes and assigns it to the Attributes field.
func (o *WafFirewallVersionData) SetAttributes(v WafFirewallVersionDataAttributes) {
	o.Attributes = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafFirewallVersionData) MarshalJSON() ([]byte, error) {
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
func (o *WafFirewallVersionData) UnmarshalJSON(bytes []byte) (err error) {
	varWafFirewallVersionData := _WafFirewallVersionData{}

	if err = json.Unmarshal(bytes, &varWafFirewallVersionData); err == nil {
		*o = WafFirewallVersionData(varWafFirewallVersionData)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafFirewallVersionData is a helper abstraction for handling nullable waffirewallversiondata types. 
type NullableWafFirewallVersionData struct {
	value *WafFirewallVersionData
	isSet bool
}

// Get returns the value.
func (v NullableWafFirewallVersionData) Get() *WafFirewallVersionData {
	return v.value
}

// Set modifies the value.
func (v *NullableWafFirewallVersionData) Set(val *WafFirewallVersionData) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafFirewallVersionData) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafFirewallVersionData) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafFirewallVersionData returns a pointer to a new instance of NullableWafFirewallVersionData.
func NewNullableWafFirewallVersionData(val *WafFirewallVersionData) *NullableWafFirewallVersionData {
	return &NullableWafFirewallVersionData{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafFirewallVersionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafFirewallVersionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
