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

// IncludedWithWafRuleItem struct for IncludedWithWafRuleItem
type IncludedWithWafRuleItem struct {
	WafRuleRevision *WafRuleRevision
	WafTag *WafTag
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *IncludedWithWafRuleItem) UnmarshalJSON(data []byte) error {
	var err error
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

	// try to unmarshal JSON data into WafTag
	err = json.Unmarshal(data, &o.WafTag);
	if err == nil {
		jsonWafTag, _ := json.Marshal(o.WafTag)
		if string(jsonWafTag) != "{}" { // empty struct
			return nil // data stored in o.WafTag, return on the first match
		}
    o.WafTag = nil
	} else {
		o.WafTag = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(IncludedWithWafRuleItem)")
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o *IncludedWithWafRuleItem) MarshalJSON() ([]byte, error) {
	if o.WafRuleRevision != nil {
		return json.Marshal(&o.WafRuleRevision)
	}

	if o.WafTag != nil {
		return json.Marshal(&o.WafTag)
	}

	return nil, nil // no data in anyOf schemas
}

// NullableIncludedWithWafRuleItem is a helper abstraction for handling nullable includedwithwafruleitem types. 
type NullableIncludedWithWafRuleItem struct {
	value *IncludedWithWafRuleItem
	isSet bool
}

// Get returns the value.
func (v NullableIncludedWithWafRuleItem) Get() *IncludedWithWafRuleItem {
	return v.value
}

// Set modifies the value.
func (v *NullableIncludedWithWafRuleItem) Set(val *IncludedWithWafRuleItem) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableIncludedWithWafRuleItem) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableIncludedWithWafRuleItem) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableIncludedWithWafRuleItem returns a pointer to a new instance of NullableIncludedWithWafRuleItem.
func NewNullableIncludedWithWafRuleItem(val *IncludedWithWafRuleItem) *NullableIncludedWithWafRuleItem {
	return &NullableIncludedWithWafRuleItem{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableIncludedWithWafRuleItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableIncludedWithWafRuleItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
