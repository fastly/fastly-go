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

// RelationshipWafFirewallVersion struct for RelationshipWafFirewallVersion
type RelationshipWafFirewallVersion struct {
	WafFirewallVersion *RelationshipWafFirewallVersionWafFirewallVersion `json:"waf_firewall_version,omitempty"`
	AdditionalProperties map[string]any
}

type _RelationshipWafFirewallVersion RelationshipWafFirewallVersion

// NewRelationshipWafFirewallVersion instantiates a new RelationshipWafFirewallVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelationshipWafFirewallVersion() *RelationshipWafFirewallVersion {
	this := RelationshipWafFirewallVersion{}
	return &this
}

// NewRelationshipWafFirewallVersionWithDefaults instantiates a new RelationshipWafFirewallVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelationshipWafFirewallVersionWithDefaults() *RelationshipWafFirewallVersion {
	this := RelationshipWafFirewallVersion{}
	return &this
}

// GetWafFirewallVersion returns the WafFirewallVersion field value if set, zero value otherwise.
func (o *RelationshipWafFirewallVersion) GetWafFirewallVersion() RelationshipWafFirewallVersionWafFirewallVersion {
	if o == nil || o.WafFirewallVersion == nil {
		var ret RelationshipWafFirewallVersionWafFirewallVersion
		return ret
	}
	return *o.WafFirewallVersion
}

// GetWafFirewallVersionOk returns a tuple with the WafFirewallVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RelationshipWafFirewallVersion) GetWafFirewallVersionOk() (*RelationshipWafFirewallVersionWafFirewallVersion, bool) {
	if o == nil || o.WafFirewallVersion == nil {
		return nil, false
	}
	return o.WafFirewallVersion, true
}

// HasWafFirewallVersion returns a boolean if a field has been set.
func (o *RelationshipWafFirewallVersion) HasWafFirewallVersion() bool {
	if o != nil && o.WafFirewallVersion != nil {
		return true
	}

	return false
}

// SetWafFirewallVersion gets a reference to the given RelationshipWafFirewallVersionWafFirewallVersion and assigns it to the WafFirewallVersion field.
func (o *RelationshipWafFirewallVersion) SetWafFirewallVersion(v RelationshipWafFirewallVersionWafFirewallVersion) {
	o.WafFirewallVersion = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o RelationshipWafFirewallVersion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.WafFirewallVersion != nil {
		toSerialize["waf_firewall_version"] = o.WafFirewallVersion
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *RelationshipWafFirewallVersion) UnmarshalJSON(bytes []byte) (err error) {
	varRelationshipWafFirewallVersion := _RelationshipWafFirewallVersion{}

	if err = json.Unmarshal(bytes, &varRelationshipWafFirewallVersion); err == nil {
		*o = RelationshipWafFirewallVersion(varRelationshipWafFirewallVersion)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "waf_firewall_version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableRelationshipWafFirewallVersion is a helper abstraction for handling nullable relationshipwaffirewallversion types. 
type NullableRelationshipWafFirewallVersion struct {
	value *RelationshipWafFirewallVersion
	isSet bool
}

// Get returns the value.
func (v NullableRelationshipWafFirewallVersion) Get() *RelationshipWafFirewallVersion {
	return v.value
}

// Set modifies the value.
func (v *NullableRelationshipWafFirewallVersion) Set(val *RelationshipWafFirewallVersion) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableRelationshipWafFirewallVersion) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableRelationshipWafFirewallVersion) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableRelationshipWafFirewallVersion returns a pointer to a new instance of NullableRelationshipWafFirewallVersion.
func NewNullableRelationshipWafFirewallVersion(val *RelationshipWafFirewallVersion) *NullableRelationshipWafFirewallVersion {
	return &NullableRelationshipWafFirewallVersion{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableRelationshipWafFirewallVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableRelationshipWafFirewallVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
