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

// ConfigStoreResponseAllOf struct for ConfigStoreResponseAllOf
type ConfigStoreResponseAllOf struct {
	// An alphanumeric string identifying the config store.
	ID                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]any
}

type _ConfigStoreResponseAllOf ConfigStoreResponseAllOf

// NewConfigStoreResponseAllOf instantiates a new ConfigStoreResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigStoreResponseAllOf() *ConfigStoreResponseAllOf {
	this := ConfigStoreResponseAllOf{}
	return &this
}

// NewConfigStoreResponseAllOfWithDefaults instantiates a new ConfigStoreResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigStoreResponseAllOfWithDefaults() *ConfigStoreResponseAllOf {
	this := ConfigStoreResponseAllOf{}
	return &this
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ConfigStoreResponseAllOf) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigStoreResponseAllOf) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ConfigStoreResponseAllOf) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ConfigStoreResponseAllOf) SetID(v string) {
	o.ID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ConfigStoreResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.ID != nil {
		toSerialize["id"] = o.ID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ConfigStoreResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConfigStoreResponseAllOf := _ConfigStoreResponseAllOf{}

	if err = json.Unmarshal(bytes, &varConfigStoreResponseAllOf); err == nil {
		*o = ConfigStoreResponseAllOf(varConfigStoreResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableConfigStoreResponseAllOf is a helper abstraction for handling nullable configstoreresponseallof types.
type NullableConfigStoreResponseAllOf struct {
	value *ConfigStoreResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableConfigStoreResponseAllOf) Get() *ConfigStoreResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableConfigStoreResponseAllOf) Set(val *ConfigStoreResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableConfigStoreResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableConfigStoreResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableConfigStoreResponseAllOf returns a pointer to a new instance of NullableConfigStoreResponseAllOf.
func NewNullableConfigStoreResponseAllOf(val *ConfigStoreResponseAllOf) *NullableConfigStoreResponseAllOf {
	return &NullableConfigStoreResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableConfigStoreResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableConfigStoreResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
