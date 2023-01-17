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
	"fmt"
)

// IncludedWithWafActiveRuleItem struct for IncludedWithWafActiveRuleItem
type IncludedWithWafActiveRuleItem struct {
	SchemasWafFirewallVersion *SchemasWafFirewallVersion
	WafRuleRevision *WafRuleRevision
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *IncludedWithWafActiveRuleItem) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into WafRuleRevision
	err = json.Unmarshal(data, &o.WafRuleRevision);
	if err == nil {
		jsonWafRuleRevision, _ := json.Marshal(o.WafRuleRevision)
		if string(jsonWafRuleRevision) != "{}" { // empty struct
			return nil // data stored in o.WafRuleRevision, return on the first match
		}
    o.WafRuleRevision = nil
	} else {
		o.WafRuleRevision = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(IncludedWithWafActiveRuleItem)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *IncludedWithWafActiveRuleItem) MarshalJSON() ([]byte, error) {
	if o.SchemasWafFirewallVersion != nil {
		return json.Marshal(&o.SchemasWafFirewallVersion)
	}

	if o.WafRuleRevision != nil {
		return json.Marshal(&o.WafRuleRevision)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableIncludedWithWafActiveRuleItem is a helper abstraction for handling nullable includedwithwafactiveruleitem types. 
type NullableIncludedWithWafActiveRuleItem struct {
	value *IncludedWithWafActiveRuleItem
	isSet bool
}

// Get returns the value.
func (v NullableIncludedWithWafActiveRuleItem) Get() *IncludedWithWafActiveRuleItem {
	return v.value
}

// Set modifies the value.
func (v *NullableIncludedWithWafActiveRuleItem) Set(val *IncludedWithWafActiveRuleItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIncludedWithWafActiveRuleItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIncludedWithWafActiveRuleItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncludedWithWafActiveRuleItem returns a pointer to a new instance of NullableIncludedWithWafActiveRuleItem.
func NewNullableIncludedWithWafActiveRuleItem(val *IncludedWithWafActiveRuleItem) *NullableIncludedWithWafActiveRuleItem {
	return &NullableIncludedWithWafActiveRuleItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIncludedWithWafActiveRuleItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableIncludedWithWafActiveRuleItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
