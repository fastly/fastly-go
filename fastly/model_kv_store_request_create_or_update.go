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

// KvStoreRequestCreateOrUpdate struct for KvStoreRequestCreateOrUpdate
type KvStoreRequestCreateOrUpdate struct {
	// A human-readable name for the store. Refer to https://docs.fastly.com/products/compute-resource-limits#kv-store for limitations on the KV store name.
	Name                 string `json:"name"`
	AdditionalProperties map[string]any
}

type _KvStoreRequestCreateOrUpdate KvStoreRequestCreateOrUpdate

// NewKvStoreRequestCreateOrUpdate instantiates a new KvStoreRequestCreateOrUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKvStoreRequestCreateOrUpdate(name string) *KvStoreRequestCreateOrUpdate {
	this := KvStoreRequestCreateOrUpdate{}
	this.Name = name
	return &this
}

// NewKvStoreRequestCreateOrUpdateWithDefaults instantiates a new KvStoreRequestCreateOrUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKvStoreRequestCreateOrUpdateWithDefaults() *KvStoreRequestCreateOrUpdate {
	this := KvStoreRequestCreateOrUpdate{}
	return &this
}

// GetName returns the Name field value
func (o *KvStoreRequestCreateOrUpdate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KvStoreRequestCreateOrUpdate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KvStoreRequestCreateOrUpdate) SetName(v string) {
	o.Name = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o KvStoreRequestCreateOrUpdate) MarshalJSON() ([]byte, error) {
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
func (o *KvStoreRequestCreateOrUpdate) UnmarshalJSON(bytes []byte) (err error) {
	varKvStoreRequestCreateOrUpdate := _KvStoreRequestCreateOrUpdate{}

	if err = json.Unmarshal(bytes, &varKvStoreRequestCreateOrUpdate); err == nil {
		*o = KvStoreRequestCreateOrUpdate(varKvStoreRequestCreateOrUpdate)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableKvStoreRequestCreateOrUpdate is a helper abstraction for handling nullable kvstorerequestcreateorupdate types.
type NullableKvStoreRequestCreateOrUpdate struct {
	value *KvStoreRequestCreateOrUpdate
	isSet bool
}

// Get returns the value.
func (v NullableKvStoreRequestCreateOrUpdate) Get() *KvStoreRequestCreateOrUpdate {
	return v.value
}

// Set modifies the value.
func (v *NullableKvStoreRequestCreateOrUpdate) Set(val *KvStoreRequestCreateOrUpdate) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableKvStoreRequestCreateOrUpdate) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableKvStoreRequestCreateOrUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableKvStoreRequestCreateOrUpdate returns a pointer to a new instance of NullableKvStoreRequestCreateOrUpdate.
func NewNullableKvStoreRequestCreateOrUpdate(val *KvStoreRequestCreateOrUpdate) *NullableKvStoreRequestCreateOrUpdate {
	return &NullableKvStoreRequestCreateOrUpdate{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableKvStoreRequestCreateOrUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableKvStoreRequestCreateOrUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
