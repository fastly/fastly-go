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

// WafActiveRuleCreationResponse - struct for WafActiveRuleCreationResponse
type WafActiveRuleCreationResponse struct {
	WafActiveRuleResponse  *WafActiveRuleResponse
	WafActiveRulesResponse *WafActiveRulesResponse
}

// WafActiveRuleResponseAsWafActiveRuleCreationResponse is a convenience function that returns WafActiveRuleResponse wrapped in WafActiveRuleCreationResponse
func WafActiveRuleResponseAsWafActiveRuleCreationResponse(v *WafActiveRuleResponse) WafActiveRuleCreationResponse {
	return WafActiveRuleCreationResponse{
		WafActiveRuleResponse: v,
	}
}

// WafActiveRulesResponseAsWafActiveRuleCreationResponse is a convenience function that returns WafActiveRulesResponse wrapped in WafActiveRuleCreationResponse
func WafActiveRulesResponseAsWafActiveRuleCreationResponse(v *WafActiveRulesResponse) WafActiveRuleCreationResponse {
	return WafActiveRuleCreationResponse{
		WafActiveRulesResponse: v,
	}
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WafActiveRuleCreationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into WafActiveRuleResponse
	err = newStrictDecoder(data).Decode(&o.WafActiveRuleResponse)
	if err == nil {
		jsonWafActiveRuleResponse, _ := json.Marshal(o.WafActiveRuleResponse)
		if string(jsonWafActiveRuleResponse) == "{}" { // empty struct
			o.WafActiveRuleResponse = nil
		} else {
			match++
		}
	} else {
		o.WafActiveRuleResponse = nil
	}

	// try to unmarshal data into WafActiveRulesResponse
	err = newStrictDecoder(data).Decode(&o.WafActiveRulesResponse)
	if err == nil {
		jsonWafActiveRulesResponse, _ := json.Marshal(o.WafActiveRulesResponse)
		if string(jsonWafActiveRulesResponse) == "{}" { // empty struct
			o.WafActiveRulesResponse = nil
		} else {
			match++
		}
	} else {
		o.WafActiveRulesResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		o.WafActiveRuleResponse = nil
		o.WafActiveRulesResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(WafActiveRuleCreationResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(WafActiveRuleCreationResponse)")
	}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafActiveRuleCreationResponse) MarshalJSON() ([]byte, error) {
	if o.WafActiveRuleResponse != nil {
		return json.Marshal(&o.WafActiveRuleResponse)
	}

	if o.WafActiveRulesResponse != nil {
		return json.Marshal(&o.WafActiveRulesResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// GetActualInstance returns a specific instance of the type.
func (o *WafActiveRuleCreationResponse) GetActualInstance() any {
	if o == nil {
		return nil
	}
	if o.WafActiveRuleResponse != nil {
		return o.WafActiveRuleResponse
	}

	if o.WafActiveRulesResponse != nil {
		return o.WafActiveRulesResponse
	}

	// all schemas are nil
	return nil
}

// NullableWafActiveRuleCreationResponse is a helper abstraction for handling nullable wafactiverulecreationresponse types.
type NullableWafActiveRuleCreationResponse struct {
	value *WafActiveRuleCreationResponse
	isSet bool
}

// Get returns the value.
func (v NullableWafActiveRuleCreationResponse) Get() *WafActiveRuleCreationResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableWafActiveRuleCreationResponse) Set(val *WafActiveRuleCreationResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafActiveRuleCreationResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafActiveRuleCreationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafActiveRuleCreationResponse returns a pointer to a new instance of NullableWafActiveRuleCreationResponse.
func NewNullableWafActiveRuleCreationResponse(val *WafActiveRuleCreationResponse) *NullableWafActiveRuleCreationResponse {
	return &NullableWafActiveRuleCreationResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafActiveRuleCreationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafActiveRuleCreationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
