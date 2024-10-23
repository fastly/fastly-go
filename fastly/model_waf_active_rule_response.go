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

// WafActiveRuleResponse struct for WafActiveRuleResponse
type WafActiveRuleResponse struct {
	Data                 *WafActiveRuleResponseData `json:"data,omitempty"`
	AdditionalProperties map[string]any
}

type _WafActiveRuleResponse WafActiveRuleResponse

// NewWafActiveRuleResponse instantiates a new WafActiveRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafActiveRuleResponse() *WafActiveRuleResponse {
	this := WafActiveRuleResponse{}
	return &this
}

// NewWafActiveRuleResponseWithDefaults instantiates a new WafActiveRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafActiveRuleResponseWithDefaults() *WafActiveRuleResponse {
	this := WafActiveRuleResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *WafActiveRuleResponse) GetData() WafActiveRuleResponseData {
	if o == nil || o.Data == nil {
		var ret WafActiveRuleResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WafActiveRuleResponse) GetDataOk() (*WafActiveRuleResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *WafActiveRuleResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given WafActiveRuleResponseData and assigns it to the Data field.
func (o *WafActiveRuleResponse) SetData(v WafActiveRuleResponseData) {
	o.Data = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafActiveRuleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WafActiveRuleResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWafActiveRuleResponse := _WafActiveRuleResponse{}

	if err = json.Unmarshal(bytes, &varWafActiveRuleResponse); err == nil {
		*o = WafActiveRuleResponse(varWafActiveRuleResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafActiveRuleResponse is a helper abstraction for handling nullable wafactiveruleresponse types.
type NullableWafActiveRuleResponse struct {
	value *WafActiveRuleResponse
	isSet bool
}

// Get returns the value.
func (v NullableWafActiveRuleResponse) Get() *WafActiveRuleResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableWafActiveRuleResponse) Set(val *WafActiveRuleResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafActiveRuleResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafActiveRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafActiveRuleResponse returns a pointer to a new instance of NullableWafActiveRuleResponse.
func NewNullableWafActiveRuleResponse(val *WafActiveRuleResponse) *NullableWafActiveRuleResponse {
	return &NullableWafActiveRuleResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafActiveRuleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafActiveRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
