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

// ArcFailoverTarget A provider connection to retry a request against.
type ArcFailoverTarget struct {
	// The ID of the provider connection to retry against.
	ProviderConnectionId string `json:"provider_connection_id"`
	AdditionalProperties map[string]any
}

type _ArcFailoverTarget ArcFailoverTarget

// NewArcFailoverTarget instantiates a new ArcFailoverTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcFailoverTarget(providerConnectionId string) *ArcFailoverTarget {
	this := ArcFailoverTarget{}
	this.ProviderConnectionId = providerConnectionId
	return &this
}

// NewArcFailoverTargetWithDefaults instantiates a new ArcFailoverTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcFailoverTargetWithDefaults() *ArcFailoverTarget {
	this := ArcFailoverTarget{}
	return &this
}

// GetProviderConnectionId returns the ProviderConnectionId field value
func (o *ArcFailoverTarget) GetProviderConnectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderConnectionId
}

// GetProviderConnectionIdOk returns a tuple with the ProviderConnectionId field value
// and a boolean to check if the value has been set.
func (o *ArcFailoverTarget) GetProviderConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderConnectionId, true
}

// SetProviderConnectionId sets field value
func (o *ArcFailoverTarget) SetProviderConnectionId(v string) {
	o.ProviderConnectionId = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcFailoverTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["provider_connection_id"] = o.ProviderConnectionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcFailoverTarget) UnmarshalJSON(bytes []byte) (err error) {
	varArcFailoverTarget := _ArcFailoverTarget{}

	if err = json.Unmarshal(bytes, &varArcFailoverTarget); err == nil {
		*o = ArcFailoverTarget(varArcFailoverTarget)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "provider_connection_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcFailoverTarget is a helper abstraction for handling nullable arcfailovertarget types.
type NullableArcFailoverTarget struct {
	value *ArcFailoverTarget
	isSet bool
}

// Get returns the value.
func (v NullableArcFailoverTarget) Get() *ArcFailoverTarget {
	return v.value
}

// Set modifies the value.
func (v *NullableArcFailoverTarget) Set(val *ArcFailoverTarget) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcFailoverTarget) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcFailoverTarget) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcFailoverTarget returns a pointer to a new instance of NullableArcFailoverTarget.
func NewNullableArcFailoverTarget(val *ArcFailoverTarget) *NullableArcFailoverTarget {
	return &NullableArcFailoverTarget{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcFailoverTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcFailoverTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
