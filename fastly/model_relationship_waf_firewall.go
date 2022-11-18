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

// RelationshipWafFirewall struct for RelationshipWafFirewall
type RelationshipWafFirewall struct {
	WafFirewall *RelationshipWafFirewallWafFirewall `json:"waf_firewall,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipWafFirewall RelationshipWafFirewall

// NewRelationshipWafFirewall instantiates a new RelationshipWafFirewall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipWafFirewall() *RelationshipWafFirewall {
	this := RelationshipWafFirewall{}
	return &this
}

// NewRelationshipWafFirewallWithDefaults instantiates a new RelationshipWafFirewall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWafFirewallWithDefaults() *RelationshipWafFirewall {
	this := RelationshipWafFirewall{}
	return &this
}

// GetWafFirewall returns the WafFirewall field value if set, zero value otherwise.
func (o *RelationshipWafFirewall) GetWafFirewall() RelationshipWafFirewallWafFirewall {
	if o == nil || o.WafFirewall == nil {
		var ret RelationshipWafFirewallWafFirewall
		return ret
	}
	return *o.WafFirewall
}

// GetWafFirewallOk returns a tuple with the WafFirewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipWafFirewall) GetWafFirewallOk() (*RelationshipWafFirewallWafFirewall, bool) {
	if o == nil || o.WafFirewall == nil {
		return nil, false
	}
	return o.WafFirewall, true
}

// HasWafFirewall returns a boolean if a field has been set.
func (o *RelationshipWafFirewall) HasWafFirewall() bool {
	if o != nil && o.WafFirewall != nil {
		return true
	}

	return false
}

// SetWafFirewall gets a reference to the given RelationshipWafFirewallWafFirewall and assigns it to the WafFirewall field.
func (o *RelationshipWafFirewall) SetWafFirewall(v RelationshipWafFirewallWafFirewall) {
	o.WafFirewall = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipWafFirewall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.WafFirewall != nil {
		toSerialize["waf_firewall"] = o.WafFirewall
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipWafFirewall) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipWafFirewall := _RelationshipWafFirewall{}

	if err = json.Unmarshal(bytes, &varRelationshipWafFirewall); err == nil {
		*o = RelationshipWafFirewall(varRelationshipWafFirewall)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "waf_firewall")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipWafFirewall is a helper abstraction for handling nullable relationshipwaffirewall types. 
type NullableRelationshipWafFirewall struct {
	value *RelationshipWafFirewall
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipWafFirewall) Get() *RelationshipWafFirewall {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipWafFirewall) Set(val *RelationshipWafFirewall) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipWafFirewall) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipWafFirewall) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipWafFirewall returns a pointer to a new instance of NullableRelationshipWafFirewall.
func NewNullableRelationshipWafFirewall(val *RelationshipWafFirewall) *NullableRelationshipWafFirewall {
	return &NullableRelationshipWafFirewall{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipWafFirewall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipWafFirewall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
