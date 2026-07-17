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

// WafSimulateResponse Response from the WAF simulation containing the WAF response code and detected signals.
type WafSimulateResponse struct {
	// The HTTP status code the WAF would return for the simulated request (e.g., `200` for allowed, `406` for blocked).
	WafResponse int32 `json:"waf_response"`
	// List of signals detected by the WAF during simulation. Empty array when no signals are detected.
	Signals              []WafSimulateSignal `json:"signals"`
	AdditionalProperties map[string]any
}

type _WafSimulateResponse WafSimulateResponse

// NewWafSimulateResponse instantiates a new WafSimulateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWafSimulateResponse(wafResponse int32, signals []WafSimulateSignal) *WafSimulateResponse {
	this := WafSimulateResponse{}
	this.WafResponse = wafResponse
	this.Signals = signals
	return &this
}

// NewWafSimulateResponseWithDefaults instantiates a new WafSimulateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWafSimulateResponseWithDefaults() *WafSimulateResponse {
	this := WafSimulateResponse{}
	return &this
}

// GetWafResponse returns the WafResponse field value
func (o *WafSimulateResponse) GetWafResponse() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WafResponse
}

// GetWafResponseOk returns a tuple with the WafResponse field value
// and a boolean to check if the value has been set.
func (o *WafSimulateResponse) GetWafResponseOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WafResponse, true
}

// SetWafResponse sets field value
func (o *WafSimulateResponse) SetWafResponse(v int32) {
	o.WafResponse = v
}

// GetSignals returns the Signals field value
func (o *WafSimulateResponse) GetSignals() []WafSimulateSignal {
	if o == nil {
		var ret []WafSimulateSignal
		return ret
	}

	return o.Signals
}

// GetSignalsOk returns a tuple with the Signals field value
// and a boolean to check if the value has been set.
func (o *WafSimulateResponse) GetSignalsOk() ([]WafSimulateSignal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Signals, true
}

// SetSignals sets field value
func (o *WafSimulateResponse) SetSignals(v []WafSimulateSignal) {
	o.Signals = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o WafSimulateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["waf_response"] = o.WafResponse
	}
	if true {
		toSerialize["signals"] = o.Signals
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *WafSimulateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varWafSimulateResponse := _WafSimulateResponse{}

	if err = json.Unmarshal(bytes, &varWafSimulateResponse); err == nil {
		*o = WafSimulateResponse(varWafSimulateResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "waf_response")
		delete(additionalProperties, "signals")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableWafSimulateResponse is a helper abstraction for handling nullable wafsimulateresponse types.
type NullableWafSimulateResponse struct {
	value *WafSimulateResponse
	isSet bool
}

// Get returns the value.
func (v NullableWafSimulateResponse) Get() *WafSimulateResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableWafSimulateResponse) Set(val *WafSimulateResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableWafSimulateResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableWafSimulateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableWafSimulateResponse returns a pointer to a new instance of NullableWafSimulateResponse.
func NewNullableWafSimulateResponse(val *WafSimulateResponse) *NullableWafSimulateResponse {
	return &NullableWafSimulateResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableWafSimulateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableWafSimulateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
