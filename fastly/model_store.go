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

// Store struct for Store
type Store struct {
	// A human-readable name for the store.
	Name                 *string `json:"name,omitempty"`
	AdditionalProperties map[string]any
}

type _Store Store

// NewStore instantiates a new Store object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStore() *Store {
	this := Store{}
	return &this
}

// NewStoreWithDefaults instantiates a new Store object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreWithDefaults() *Store {
	this := Store{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Store) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Store) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Store) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Store) SetName(v string) {
	o.Name = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o Store) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *Store) UnmarshalJSON(bytes []byte) (err error) {
	varStore := _Store{}

	if err = json.Unmarshal(bytes, &varStore); err == nil {
		*o = Store(varStore)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableStore is a helper abstraction for handling nullable store types.
type NullableStore struct {
	value *Store
	isSet bool
}

// Get returns the value.
func (v NullableStore) Get() *Store {
	return v.value
}

// Set modifies the value.
func (v *NullableStore) Set(val *Store) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableStore) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableStore) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableStore returns a pointer to a new instance of NullableStore.
func NewNullableStore(val *Store) *NullableStore {
	return &NullableStore{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableStore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
