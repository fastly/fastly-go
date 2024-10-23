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

// IncludedWithWafExclusionItem struct for IncludedWithWafExclusionItem
type IncludedWithWafExclusionItem struct {
	WafRule         *WafRule
	WafRuleRevision *WafRuleRevision
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *IncludedWithWafExclusionItem) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into WafRule
	err = json.Unmarshal(data, &o.WafRule)
	if err == nil {
		jsonWafRule, _ := json.Marshal(o.WafRule)
		if string(jsonWafRule) != "{}" { // empty struct
			return nil // data stored in o.WafRule, return on the first match
		}
		o.WafRule = nil
	} else {
		o.WafRule = nil
	}

	// try to unmarshal JSON data into WafRuleRevision
	err = json.Unmarshal(data, &o.WafRuleRevision)
	if err == nil {
		jsonWafRuleRevision, _ := json.Marshal(o.WafRuleRevision)
		if string(jsonWafRuleRevision) != "{}" { // empty struct
			return nil // data stored in o.WafRuleRevision, return on the first match
		}
		o.WafRuleRevision = nil
	} else {
		o.WafRuleRevision = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(IncludedWithWafExclusionItem)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *IncludedWithWafExclusionItem) MarshalJSON() ([]byte, error) {
	if o.WafRule != nil {
		return json.Marshal(&o.WafRule)
	}

	if o.WafRuleRevision != nil {
		return json.Marshal(&o.WafRuleRevision)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableIncludedWithWafExclusionItem is a helper abstraction for handling nullable includedwithwafexclusionitem types.
type NullableIncludedWithWafExclusionItem struct {
	value *IncludedWithWafExclusionItem
	isSet bool
}

// Get returns the value.
func (v NullableIncludedWithWafExclusionItem) Get() *IncludedWithWafExclusionItem {
	return v.value
}

// Set modifies the value.
func (v *NullableIncludedWithWafExclusionItem) Set(val *IncludedWithWafExclusionItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIncludedWithWafExclusionItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIncludedWithWafExclusionItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncludedWithWafExclusionItem returns a pointer to a new instance of NullableIncludedWithWafExclusionItem.
func NewNullableIncludedWithWafExclusionItem(val *IncludedWithWafExclusionItem) *NullableIncludedWithWafExclusionItem {
	return &NullableIncludedWithWafExclusionItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIncludedWithWafExclusionItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableIncludedWithWafExclusionItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
