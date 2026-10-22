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

// ArcVirtualKeyFailoverResponse A virtual key's failover policy. `targets` are tried in order after the primary request fails and one of the enabled error conditions is met.
type ArcVirtualKeyFailoverResponse struct {
	// The provider connections to retry against, in order.
	Targets []ArcFailoverTarget `json:"targets"`
	// Whether a request timeout triggers failover.
	TimeoutEnabled bool `json:"timeout_enabled"`
	// Whether a `429 Too Many Requests` response from the provider triggers failover.
	RateLimit429Enabled bool `json:"rate_limit_429_enabled"`
	// Whether a `5xx` server error response from the provider triggers failover.
	ServerError5xxEnabled bool `json:"server_error_5xx_enabled"`
	AdditionalProperties  map[string]any
}

type _ArcVirtualKeyFailoverResponse ArcVirtualKeyFailoverResponse

// NewArcVirtualKeyFailoverResponse instantiates a new ArcVirtualKeyFailoverResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcVirtualKeyFailoverResponse(targets []ArcFailoverTarget, timeoutEnabled bool, rateLimit429Enabled bool, serverError5xxEnabled bool) *ArcVirtualKeyFailoverResponse {
	this := ArcVirtualKeyFailoverResponse{}
	this.Targets = targets
	this.TimeoutEnabled = timeoutEnabled
	this.RateLimit429Enabled = rateLimit429Enabled
	this.ServerError5xxEnabled = serverError5xxEnabled
	return &this
}

// NewArcVirtualKeyFailoverResponseWithDefaults instantiates a new ArcVirtualKeyFailoverResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcVirtualKeyFailoverResponseWithDefaults() *ArcVirtualKeyFailoverResponse {
	this := ArcVirtualKeyFailoverResponse{}
	return &this
}

// GetTargets returns the Targets field value
func (o *ArcVirtualKeyFailoverResponse) GetTargets() []ArcFailoverTarget {
	if o == nil {
		var ret []ArcFailoverTarget
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyFailoverResponse) GetTargetsOk() ([]ArcFailoverTarget, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *ArcVirtualKeyFailoverResponse) SetTargets(v []ArcFailoverTarget) {
	o.Targets = v
}

// GetTimeoutEnabled returns the TimeoutEnabled field value
func (o *ArcVirtualKeyFailoverResponse) GetTimeoutEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TimeoutEnabled
}

// GetTimeoutEnabledOk returns a tuple with the TimeoutEnabled field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyFailoverResponse) GetTimeoutEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeoutEnabled, true
}

// SetTimeoutEnabled sets field value
func (o *ArcVirtualKeyFailoverResponse) SetTimeoutEnabled(v bool) {
	o.TimeoutEnabled = v
}

// GetRateLimit429Enabled returns the RateLimit429Enabled field value
func (o *ArcVirtualKeyFailoverResponse) GetRateLimit429Enabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RateLimit429Enabled
}

// GetRateLimit429EnabledOk returns a tuple with the RateLimit429Enabled field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyFailoverResponse) GetRateLimit429EnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RateLimit429Enabled, true
}

// SetRateLimit429Enabled sets field value
func (o *ArcVirtualKeyFailoverResponse) SetRateLimit429Enabled(v bool) {
	o.RateLimit429Enabled = v
}

// GetServerError5xxEnabled returns the ServerError5xxEnabled field value
func (o *ArcVirtualKeyFailoverResponse) GetServerError5xxEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ServerError5xxEnabled
}

// GetServerError5xxEnabledOk returns a tuple with the ServerError5xxEnabled field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyFailoverResponse) GetServerError5xxEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerError5xxEnabled, true
}

// SetServerError5xxEnabled sets field value
func (o *ArcVirtualKeyFailoverResponse) SetServerError5xxEnabled(v bool) {
	o.ServerError5xxEnabled = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcVirtualKeyFailoverResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["targets"] = o.Targets
	}
	if true {
		toSerialize["timeout_enabled"] = o.TimeoutEnabled
	}
	if true {
		toSerialize["rate_limit_429_enabled"] = o.RateLimit429Enabled
	}
	if true {
		toSerialize["server_error_5xx_enabled"] = o.ServerError5xxEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcVirtualKeyFailoverResponse) UnmarshalJSON(bytes []byte) (err error) {
	varArcVirtualKeyFailoverResponse := _ArcVirtualKeyFailoverResponse{}

	if err = json.Unmarshal(bytes, &varArcVirtualKeyFailoverResponse); err == nil {
		*o = ArcVirtualKeyFailoverResponse(varArcVirtualKeyFailoverResponse)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "targets")
		delete(additionalProperties, "timeout_enabled")
		delete(additionalProperties, "rate_limit_429_enabled")
		delete(additionalProperties, "server_error_5xx_enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcVirtualKeyFailoverResponse is a helper abstraction for handling nullable arcvirtualkeyfailoverresponse types.
type NullableArcVirtualKeyFailoverResponse struct {
	value *ArcVirtualKeyFailoverResponse
	isSet bool
}

// Get returns the value.
func (v NullableArcVirtualKeyFailoverResponse) Get() *ArcVirtualKeyFailoverResponse {
	return v.value
}

// Set modifies the value.
func (v *NullableArcVirtualKeyFailoverResponse) Set(val *ArcVirtualKeyFailoverResponse) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcVirtualKeyFailoverResponse) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcVirtualKeyFailoverResponse) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcVirtualKeyFailoverResponse returns a pointer to a new instance of NullableArcVirtualKeyFailoverResponse.
func NewNullableArcVirtualKeyFailoverResponse(val *ArcVirtualKeyFailoverResponse) *NullableArcVirtualKeyFailoverResponse {
	return &NullableArcVirtualKeyFailoverResponse{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcVirtualKeyFailoverResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcVirtualKeyFailoverResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
