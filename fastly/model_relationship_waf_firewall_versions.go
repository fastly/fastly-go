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

// RelationshipWafFirewallVersions struct for RelationshipWafFirewallVersions
type RelationshipWafFirewallVersions struct {
	WafFirewallVersions *RelationshipWafFirewallVersionWafFirewallVersion `json:"waf_firewall_versions,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipWafFirewallVersions RelationshipWafFirewallVersions

// NewRelationshipWafFirewallVersions instantiates a new RelationshipWafFirewallVersions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipWafFirewallVersions() *RelationshipWafFirewallVersions {
	this := RelationshipWafFirewallVersions{}
	return &this
}

// NewRelationshipWafFirewallVersionsWithDefaults instantiates a new RelationshipWafFirewallVersions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWafFirewallVersionsWithDefaults() *RelationshipWafFirewallVersions {
	this := RelationshipWafFirewallVersions{}
	return &this
}

// GetWafFirewallVersions returns the WafFirewallVersions field value if set, zero value otherwise.
func (o *RelationshipWafFirewallVersions) GetWafFirewallVersions() RelationshipWafFirewallVersionWafFirewallVersion {
	if o == nil || o.WafFirewallVersions == nil {
		var ret RelationshipWafFirewallVersionWafFirewallVersion
		return ret
	}
	return *o.WafFirewallVersions
}

// GetWafFirewallVersionsOk returns a tuple with the WafFirewallVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipWafFirewallVersions) GetWafFirewallVersionsOk() (*RelationshipWafFirewallVersionWafFirewallVersion, bool) {
	if o == nil || o.WafFirewallVersions == nil {
		return nil, false
	}
	return o.WafFirewallVersions, true
}

// HasWafFirewallVersions returns a boolean if a field has been set.
func (o *RelationshipWafFirewallVersions) HasWafFirewallVersions() bool {
	if o != nil && o.WafFirewallVersions != nil {
		return true
	}

	return false
}

// SetWafFirewallVersions gets a reference to the given RelationshipWafFirewallVersionWafFirewallVersion and assigns it to the WafFirewallVersions field.
func (o *RelationshipWafFirewallVersions) SetWafFirewallVersions(v RelationshipWafFirewallVersionWafFirewallVersion) {
	o.WafFirewallVersions = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipWafFirewallVersions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.WafFirewallVersions != nil {
		toSerialize["waf_firewall_versions"] = o.WafFirewallVersions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipWafFirewallVersions) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipWafFirewallVersions := _RelationshipWafFirewallVersions{}

	if err = json.Unmarshal(bytes, &varRelationshipWafFirewallVersions); err == nil {
		*o = RelationshipWafFirewallVersions(varRelationshipWafFirewallVersions)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "waf_firewall_versions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipWafFirewallVersions is a helper abstraction for handling nullable relationshipwaffirewallversions types. 
type NullableRelationshipWafFirewallVersions struct {
	value *RelationshipWafFirewallVersions
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipWafFirewallVersions) Get() *RelationshipWafFirewallVersions {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipWafFirewallVersions) Set(val *RelationshipWafFirewallVersions) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipWafFirewallVersions) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipWafFirewallVersions) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipWafFirewallVersions returns a pointer to a new instance of NullableRelationshipWafFirewallVersions.
func NewNullableRelationshipWafFirewallVersions(val *RelationshipWafFirewallVersions) *NullableRelationshipWafFirewallVersions {
	return &NullableRelationshipWafFirewallVersions{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipWafFirewallVersions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipWafFirewallVersions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
