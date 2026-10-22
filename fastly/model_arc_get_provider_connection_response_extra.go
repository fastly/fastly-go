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

// ArcGetProviderConnectionResponseExtra struct for ArcGetProviderConnectionResponseExtra
type ArcGetProviderConnectionResponseExtra struct {
	// The IDs of the active virtual keys whose failover policy references this connection. Empty if no virtual key references it.
	VirtualKeyIds        []string `json:"virtual_key_ids"`
	AdditionalProperties map[string]any
}

type _ArcGetProviderConnectionResponseExtra ArcGetProviderConnectionResponseExtra

// NewArcGetProviderConnectionResponseExtra instantiates a new ArcGetProviderConnectionResponseExtra object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArcGetProviderConnectionResponseExtra(virtualKeyIds []string) *ArcGetProviderConnectionResponseExtra {
	this := ArcGetProviderConnectionResponseExtra{}
	this.VirtualKeyIds = virtualKeyIds
	return &this
}

// NewArcGetProviderConnectionResponseExtraWithDefaults instantiates a new ArcGetProviderConnectionResponseExtra object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArcGetProviderConnectionResponseExtraWithDefaults() *ArcGetProviderConnectionResponseExtra {
	this := ArcGetProviderConnectionResponseExtra{}
	return &this
}

// GetVirtualKeyIds returns the VirtualKeyIds field value
func (o *ArcGetProviderConnectionResponseExtra) GetVirtualKeyIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.VirtualKeyIds
}

// GetVirtualKeyIdsOk returns a tuple with the VirtualKeyIds field value
// and a boolean to check if the value has been set.
func (o *ArcGetProviderConnectionResponseExtra) GetVirtualKeyIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VirtualKeyIds, true
}

// SetVirtualKeyIds sets field value
func (o *ArcGetProviderConnectionResponseExtra) SetVirtualKeyIds(v []string) {
	o.VirtualKeyIds = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ArcGetProviderConnectionResponseExtra) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["virtual_key_ids"] = o.VirtualKeyIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ArcGetProviderConnectionResponseExtra) UnmarshalJSON(bytes []byte) (err error) {
	varArcGetProviderConnectionResponseExtra := _ArcGetProviderConnectionResponseExtra{}

	if err = json.Unmarshal(bytes, &varArcGetProviderConnectionResponseExtra); err == nil {
		*o = ArcGetProviderConnectionResponseExtra(varArcGetProviderConnectionResponseExtra)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "virtual_key_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableArcGetProviderConnectionResponseExtra is a helper abstraction for handling nullable arcgetproviderconnectionresponseextra types.
type NullableArcGetProviderConnectionResponseExtra struct {
	value *ArcGetProviderConnectionResponseExtra
	isSet bool
}

// Get returns the value.
func (v NullableArcGetProviderConnectionResponseExtra) Get() *ArcGetProviderConnectionResponseExtra {
	return v.value
}

// Set modifies the value.
func (v *NullableArcGetProviderConnectionResponseExtra) Set(val *ArcGetProviderConnectionResponseExtra) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableArcGetProviderConnectionResponseExtra) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableArcGetProviderConnectionResponseExtra) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableArcGetProviderConnectionResponseExtra returns a pointer to a new instance of NullableArcGetProviderConnectionResponseExtra.
func NewNullableArcGetProviderConnectionResponseExtra(val *ArcGetProviderConnectionResponseExtra) *NullableArcGetProviderConnectionResponseExtra {
	return &NullableArcGetProviderConnectionResponseExtra{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableArcGetProviderConnectionResponseExtra) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableArcGetProviderConnectionResponseExtra) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
