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

// NetworkProtectEndpointPolicy The endpoints serving one network region. Replacing `endpoint_policies` preserves attachment timestamps for the same endpoint and region; reordering does not alter them, while changing preference updates the attachment's `updated_at` value.
type NetworkProtectEndpointPolicy struct {
	Region *NetworkProtectRegion `json:"region,omitempty"`
	// The endpoints serving a policy region.
	Endpoints            []NetworkProtectPolicyEndpoint `json:"endpoints"`
	AdditionalProperties map[string]any
}

type _NetworkProtectEndpointPolicy NetworkProtectEndpointPolicy

// NewNetworkProtectEndpointPolicy instantiates a new NetworkProtectEndpointPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectEndpointPolicy(endpoints []NetworkProtectPolicyEndpoint) *NetworkProtectEndpointPolicy {
	this := NetworkProtectEndpointPolicy{}
	var region NetworkProtectRegion = NETWORKPROTECTREGION_DEFAULT
	this.Region = &region
	this.Endpoints = endpoints
	return &this
}

// NewNetworkProtectEndpointPolicyWithDefaults instantiates a new NetworkProtectEndpointPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectEndpointPolicyWithDefaults() *NetworkProtectEndpointPolicy {
	this := NetworkProtectEndpointPolicy{}
	var region NetworkProtectRegion = NETWORKPROTECTREGION_DEFAULT
	this.Region = &region
	return &this
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *NetworkProtectEndpointPolicy) GetRegion() NetworkProtectRegion {
	if o == nil || o.Region == nil {
		var ret NetworkProtectRegion
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpointPolicy) GetRegionOk() (*NetworkProtectRegion, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *NetworkProtectEndpointPolicy) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NetworkProtectRegion and assigns it to the Region field.
func (o *NetworkProtectEndpointPolicy) SetRegion(v NetworkProtectRegion) {
	o.Region = &v
}

// GetEndpoints returns the Endpoints field value
func (o *NetworkProtectEndpointPolicy) GetEndpoints() []NetworkProtectPolicyEndpoint {
	if o == nil {
		var ret []NetworkProtectPolicyEndpoint
		return ret
	}

	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpointPolicy) GetEndpointsOk() ([]NetworkProtectPolicyEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Endpoints, true
}

// SetEndpoints sets field value
func (o *NetworkProtectEndpointPolicy) SetEndpoints(v []NetworkProtectPolicyEndpoint) {
	o.Endpoints = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectEndpointPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["endpoints"] = o.Endpoints
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectEndpointPolicy) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectEndpointPolicy := _NetworkProtectEndpointPolicy{}

	if err = json.Unmarshal(bytes, &varNetworkProtectEndpointPolicy); err == nil {
		*o = NetworkProtectEndpointPolicy(varNetworkProtectEndpointPolicy)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "region")
		delete(additionalProperties, "endpoints")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectEndpointPolicy is a helper abstraction for handling nullable networkprotectendpointpolicy types.
type NullableNetworkProtectEndpointPolicy struct {
	value *NetworkProtectEndpointPolicy
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectEndpointPolicy) Get() *NetworkProtectEndpointPolicy {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectEndpointPolicy) Set(val *NetworkProtectEndpointPolicy) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectEndpointPolicy) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectEndpointPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectEndpointPolicy returns a pointer to a new instance of NullableNetworkProtectEndpointPolicy.
func NewNullableNetworkProtectEndpointPolicy(val *NetworkProtectEndpointPolicy) *NullableNetworkProtectEndpointPolicy {
	return &NullableNetworkProtectEndpointPolicy{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectEndpointPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectEndpointPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
