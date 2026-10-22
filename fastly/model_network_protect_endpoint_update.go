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

// NetworkProtectEndpointUpdate Attributes accepted when updating an endpoint.
type NetworkProtectEndpointUpdate struct {
	// A human-readable name, unique within the customer.
	Name                 string `json:"name"`
	AdditionalProperties map[string]any
}

type _NetworkProtectEndpointUpdate NetworkProtectEndpointUpdate

// NewNetworkProtectEndpointUpdate instantiates a new NetworkProtectEndpointUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectEndpointUpdate(name string) *NetworkProtectEndpointUpdate {
	this := NetworkProtectEndpointUpdate{}
	this.Name = name
	return &this
}

// NewNetworkProtectEndpointUpdateWithDefaults instantiates a new NetworkProtectEndpointUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectEndpointUpdateWithDefaults() *NetworkProtectEndpointUpdate {
	this := NetworkProtectEndpointUpdate{}
	return &this
}

// GetName returns the Name field value
func (o *NetworkProtectEndpointUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectEndpointUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NetworkProtectEndpointUpdate) SetName(v string) {
	o.Name = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectEndpointUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectEndpointUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectEndpointUpdate := _NetworkProtectEndpointUpdate{}

	if err = json.Unmarshal(bytes, &varNetworkProtectEndpointUpdate); err == nil {
		*o = NetworkProtectEndpointUpdate(varNetworkProtectEndpointUpdate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectEndpointUpdate is a helper abstraction for handling nullable networkprotectendpointupdate types.
type NullableNetworkProtectEndpointUpdate struct {
	value *NetworkProtectEndpointUpdate
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectEndpointUpdate) Get() *NetworkProtectEndpointUpdate {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectEndpointUpdate) Set(val *NetworkProtectEndpointUpdate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectEndpointUpdate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectEndpointUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectEndpointUpdate returns a pointer to a new instance of NullableNetworkProtectEndpointUpdate.
func NewNullableNetworkProtectEndpointUpdate(val *NetworkProtectEndpointUpdate) *NullableNetworkProtectEndpointUpdate {
	return &NullableNetworkProtectEndpointUpdate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectEndpointUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectEndpointUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
