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

// ArcVirtualKeyFailoverRequest A failover policy for a virtual key. `targets` are tried in order after the primary request fails and one of the enabled error conditions is met.
type ArcVirtualKeyFailoverRequest struct {
	// The provider connections to retry against, in order.
	Targets []ArcFailoverTarget `json:"targets"`
	// Whether a request timeout triggers failover.
	TimeoutEnabled *bool `json:"timeout_enabled,omitempty"`
	// Whether a `429 Too Many Requests` response from the provider triggers failover.
	RateLimit429Enabled *bool `json:"rate_limit_429_enabled,omitempty"`
	// Whether a `5xx` server error response from the provider triggers failover.
	ServerError5xxEnabled *bool `json:"server_error_5xx_enabled,omitempty"`
	AdditionalProperties  map[string]any
}

type _ArcVirtualKeyFailoverRequest ArcVirtualKeyFailoverRequest

// NewArcVirtualKeyFailoverRequest instantiates a new ArcVirtualKeyFailoverRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcVirtualKeyFailoverRequest(targets []ArcFailoverTarget) *ArcVirtualKeyFailoverRequest {
	this := ArcVirtualKeyFailoverRequest{}
	this.Targets = targets
	return &this
}

// NewArcVirtualKeyFailoverRequestWithDefaults instantiates a new ArcVirtualKeyFailoverRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcVirtualKeyFailoverRequestWithDefaults() *ArcVirtualKeyFailoverRequest {
	this := ArcVirtualKeyFailoverRequest{}
	return &this
}

// GetTargets returns the Targets field value
func (o *ArcVirtualKeyFailoverRequest) GetTargets() []ArcFailoverTarget {
	if o == nil {
		var ret []ArcFailoverTarget
		return ret
	}

	return o.Targets
}

// GetTargetsOk returns a tuple with the Targets field value
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyFailoverRequest) GetTargetsOk() ([]ArcFailoverTarget, bool) {
	if o == nil {
		return nil, false
	}
	return o.Targets, true
}

// SetTargets sets field value
func (o *ArcVirtualKeyFailoverRequest) SetTargets(v []ArcFailoverTarget) {
	o.Targets = v
}

// GetTimeoutEnabled returns the TimeoutEnabled field value if set, zero value otherwise.
func (o *ArcVirtualKeyFailoverRequest) GetTimeoutEnabled() bool {
	if o == nil || o.TimeoutEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TimeoutEnabled
}

// GetTimeoutEnabledOk returns a tuple with the TimeoutEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyFailoverRequest) GetTimeoutEnabledOk() (*bool, bool) {
	if o == nil || o.TimeoutEnabled == nil {
		return nil, false
	}
	return o.TimeoutEnabled, true
}

// HasTimeoutEnabled returns a boolean if a field has been set.
func (o *ArcVirtualKeyFailoverRequest) HasTimeoutEnabled() bool {
	if o != nil && o.TimeoutEnabled != nil {
		return true
	}

	return false
}

// SetTimeoutEnabled gets a reference to the given bool and assigns it to the TimeoutEnabled field.
func (o *ArcVirtualKeyFailoverRequest) SetTimeoutEnabled(v bool) {
	o.TimeoutEnabled = &v
}

// GetRateLimit429Enabled returns the RateLimit429Enabled field value if set, zero value otherwise.
func (o *ArcVirtualKeyFailoverRequest) GetRateLimit429Enabled() bool {
	if o == nil || o.RateLimit429Enabled == nil {
		var ret bool
		return ret
	}
	return *o.RateLimit429Enabled
}

// GetRateLimit429EnabledOk returns a tuple with the RateLimit429Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyFailoverRequest) GetRateLimit429EnabledOk() (*bool, bool) {
	if o == nil || o.RateLimit429Enabled == nil {
		return nil, false
	}
	return o.RateLimit429Enabled, true
}

// HasRateLimit429Enabled returns a boolean if a field has been set.
func (o *ArcVirtualKeyFailoverRequest) HasRateLimit429Enabled() bool {
	if o != nil && o.RateLimit429Enabled != nil {
		return true
	}

	return false
}

// SetRateLimit429Enabled gets a reference to the given bool and assigns it to the RateLimit429Enabled field.
func (o *ArcVirtualKeyFailoverRequest) SetRateLimit429Enabled(v bool) {
	o.RateLimit429Enabled = &v
}

// GetServerError5xxEnabled returns the ServerError5xxEnabled field value if set, zero value otherwise.
func (o *ArcVirtualKeyFailoverRequest) GetServerError5xxEnabled() bool {
	if o == nil || o.ServerError5xxEnabled == nil {
		var ret bool
		return ret
	}
	return *o.ServerError5xxEnabled
}

// GetServerError5xxEnabledOk returns a tuple with the ServerError5xxEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArcVirtualKeyFailoverRequest) GetServerError5xxEnabledOk() (*bool, bool) {
	if o == nil || o.ServerError5xxEnabled == nil {
		return nil, false
	}
	return o.ServerError5xxEnabled, true
}

// HasServerError5xxEnabled returns a boolean if a field has been set.
func (o *ArcVirtualKeyFailoverRequest) HasServerError5xxEnabled() bool {
	if o != nil && o.ServerError5xxEnabled != nil {
		return true
	}

	return false
}

// SetServerError5xxEnabled gets a reference to the given bool and assigns it to the ServerError5xxEnabled field.
func (o *ArcVirtualKeyFailoverRequest) SetServerError5xxEnabled(v bool) {
	o.ServerError5xxEnabled = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcVirtualKeyFailoverRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["targets"] = o.Targets
	}
	if o.TimeoutEnabled != nil {
		toSerialize["timeout_enabled"] = o.TimeoutEnabled
	}
	if o.RateLimit429Enabled != nil {
		toSerialize["rate_limit_429_enabled"] = o.RateLimit429Enabled
	}
	if o.ServerError5xxEnabled != nil {
		toSerialize["server_error_5xx_enabled"] = o.ServerError5xxEnabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcVirtualKeyFailoverRequest) UnmarshalJSON(bytes []byte) (err error) {
	varArcVirtualKeyFailoverRequest := _ArcVirtualKeyFailoverRequest{}

	if err = json.Unmarshal(bytes, &varArcVirtualKeyFailoverRequest); err == nil {
		*o = ArcVirtualKeyFailoverRequest(varArcVirtualKeyFailoverRequest)
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

// NullableArcVirtualKeyFailoverRequest is a helper abstraction for handling nullable arcvirtualkeyfailoverrequest types.
type NullableArcVirtualKeyFailoverRequest struct {
	value *ArcVirtualKeyFailoverRequest
	isSet bool
}

// Get returns the value.
func (v NullableArcVirtualKeyFailoverRequest) Get() *ArcVirtualKeyFailoverRequest {
	return v.value
}

// Set modifies the value.
func (v *NullableArcVirtualKeyFailoverRequest) Set(val *ArcVirtualKeyFailoverRequest) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcVirtualKeyFailoverRequest) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcVirtualKeyFailoverRequest) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcVirtualKeyFailoverRequest returns a pointer to a new instance of NullableArcVirtualKeyFailoverRequest.
func NewNullableArcVirtualKeyFailoverRequest(val *ArcVirtualKeyFailoverRequest) *NullableArcVirtualKeyFailoverRequest {
	return &NullableArcVirtualKeyFailoverRequest{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcVirtualKeyFailoverRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcVirtualKeyFailoverRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
