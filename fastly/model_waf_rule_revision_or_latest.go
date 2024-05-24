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

// WafRuleRevisionOrLatest - struct for WafRuleRevisionOrLatest
type WafRuleRevisionOrLatest struct {
	Int32 *int32
	String *string
}

// Int32AsWafRuleRevisionOrLatest is a convenience function that returns int32 wrapped in WafRuleRevisionOrLatest
func Int32AsWafRuleRevisionOrLatest(v *int32) WafRuleRevisionOrLatest {
	return WafRuleRevisionOrLatest{
		Int32: v,
	}
}

// StringAsWafRuleRevisionOrLatest is a convenience function that returns string wrapped in WafRuleRevisionOrLatest
func StringAsWafRuleRevisionOrLatest(v *string) WafRuleRevisionOrLatest {
	return WafRuleRevisionOrLatest{
		String: v,
	}
}


// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *WafRuleRevisionOrLatest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&o.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(o.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			o.Int32 = nil
		} else {
			match++
		}
	} else {
		o.Int32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&o.String)
	if err == nil {
		jsonString, _ := json.Marshal(o.String)
		if string(jsonString) == "{}" { // empty struct
			o.String = nil
		} else {
			match++
		}
	} else {
		o.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		o.Int32 = nil
		o.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(WafRuleRevisionOrLatest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(WafRuleRevisionOrLatest)")
	}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafRuleRevisionOrLatest) MarshalJSON() ([]byte, error) {
	if o.Int32 != nil {
		return json.Marshal(&o.Int32)
	}

	if o.String != nil {
		return json.Marshal(&o.String)
	}

	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns a specific instance of the type.
func (o *WafRuleRevisionOrLatest) GetActualInstance() (any) {
	if o == nil {
		return nil
	}
	if o.Int32 != nil {
		return o.Int32
	}

	if o.String != nil {
		return o.String
	}

	// all schemas are nil
	return nil
}

// NullableWafRuleRevisionOrLatest is a helper abstraction for handling nullable wafrulerevisionorlatest types. 
type NullableWafRuleRevisionOrLatest struct {
	value *WafRuleRevisionOrLatest
	isSet bool
}

// Get returns the value.
func (v NullableWafRuleRevisionOrLatest) Get() *WafRuleRevisionOrLatest {
	return v.value
}

// Set modifies the value.
func (v *NullableWafRuleRevisionOrLatest) Set(val *WafRuleRevisionOrLatest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafRuleRevisionOrLatest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafRuleRevisionOrLatest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafRuleRevisionOrLatest returns a pointer to a new instance of NullableWafRuleRevisionOrLatest.
func NewNullableWafRuleRevisionOrLatest(val *WafRuleRevisionOrLatest) *NullableWafRuleRevisionOrLatest {
	return &NullableWafRuleRevisionOrLatest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafRuleRevisionOrLatest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableWafRuleRevisionOrLatest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
