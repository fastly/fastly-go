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

// AgentKeyMeta Metadata about the agent keys request.
type AgentKeyMeta struct {
	// Total number of agent keys.
	Total                int32 `json:"total"`
	AdditionalProperties map[string]any
}

type _AgentKeyMeta AgentKeyMeta

// NewAgentKeyMeta instantiates a new AgentKeyMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentKeyMeta(total int32) *AgentKeyMeta {
	this := AgentKeyMeta{}
	this.Total = total
	return &this
}

// NewAgentKeyMetaWithDefaults instantiates a new AgentKeyMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentKeyMetaWithDefaults() *AgentKeyMeta {
	this := AgentKeyMeta{}
	return &this
}

// GetTotal returns the Total field value
func (o *AgentKeyMeta) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AgentKeyMeta) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AgentKeyMeta) SetTotal(v int32) {
	o.Total = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AgentKeyMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AgentKeyMeta) UnmarshalJSON(bytes []byte) (err error) {
	varAgentKeyMeta := _AgentKeyMeta{}

	if err = json.Unmarshal(bytes, &varAgentKeyMeta); err == nil {
		*o = AgentKeyMeta(varAgentKeyMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAgentKeyMeta is a helper abstraction for handling nullable agentkeymeta types.
type NullableAgentKeyMeta struct {
	value *AgentKeyMeta
	isSet bool
}

// Get returns the value.
func (v NullableAgentKeyMeta) Get() *AgentKeyMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableAgentKeyMeta) Set(val *AgentKeyMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAgentKeyMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAgentKeyMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAgentKeyMeta returns a pointer to a new instance of NullableAgentKeyMeta.
func NewNullableAgentKeyMeta(val *AgentKeyMeta) *NullableAgentKeyMeta {
	return &NullableAgentKeyMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAgentKeyMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAgentKeyMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
