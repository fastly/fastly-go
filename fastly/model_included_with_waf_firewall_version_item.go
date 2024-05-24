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
	"fmt"
)

// IncludedWithWafFirewallVersionItem struct for IncludedWithWafFirewallVersionItem
type IncludedWithWafFirewallVersionItem struct {
	SchemasWafFirewallVersion *SchemasWafFirewallVersion
	WafActiveRule *WafActiveRule
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *IncludedWithWafFirewallVersionItem) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SchemasWafFirewallVersion
	err = json.Unmarshal(data, &o.SchemasWafFirewallVersion);
	if err == nil {
		jsonSchemasWafFirewallVersion, _ := json.Marshal(o.SchemasWafFirewallVersion)
		if string(jsonSchemasWafFirewallVersion) != "{}" { // empty struct
			return nil // data stored in o.SchemasWafFirewallVersion, return on the first match
		}
    o.SchemasWafFirewallVersion = nil
	} else {
		o.SchemasWafFirewallVersion = nil
	}

	// try to unmarshal JSON data into WafActiveRule
	err = json.Unmarshal(data, &o.WafActiveRule);
	if err == nil {
		jsonWafActiveRule, _ := json.Marshal(o.WafActiveRule)
		if string(jsonWafActiveRule) != "{}" { // empty struct
			return nil // data stored in o.WafActiveRule, return on the first match
		}
    o.WafActiveRule = nil
	} else {
		o.WafActiveRule = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(IncludedWithWafFirewallVersionItem)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *IncludedWithWafFirewallVersionItem) MarshalJSON() ([]byte, error) {
	if o.SchemasWafFirewallVersion != nil {
		return json.Marshal(&o.SchemasWafFirewallVersion)
	}

	if o.WafActiveRule != nil {
		return json.Marshal(&o.WafActiveRule)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableIncludedWithWafFirewallVersionItem is a helper abstraction for handling nullable includedwithwaffirewallversionitem types. 
type NullableIncludedWithWafFirewallVersionItem struct {
	value *IncludedWithWafFirewallVersionItem
	isSet bool
}

// Get returns the value.
func (v NullableIncludedWithWafFirewallVersionItem) Get() *IncludedWithWafFirewallVersionItem {
	return v.value
}

// Set modifies the value.
func (v *NullableIncludedWithWafFirewallVersionItem) Set(val *IncludedWithWafFirewallVersionItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIncludedWithWafFirewallVersionItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIncludedWithWafFirewallVersionItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncludedWithWafFirewallVersionItem returns a pointer to a new instance of NullableIncludedWithWafFirewallVersionItem.
func NewNullableIncludedWithWafFirewallVersionItem(val *IncludedWithWafFirewallVersionItem) *NullableIncludedWithWafFirewallVersionItem {
	return &NullableIncludedWithWafFirewallVersionItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIncludedWithWafFirewallVersionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableIncludedWithWafFirewallVersionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
