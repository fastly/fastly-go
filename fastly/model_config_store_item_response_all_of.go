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

// ConfigStoreItemResponseAllOf struct for ConfigStoreItemResponseAllOf
type ConfigStoreItemResponseAllOf struct {
	StoreID *string `json:"store_id,omitempty"`
	AdditionalProperties map[string]any
}

type _ConfigStoreItemResponseAllOf ConfigStoreItemResponseAllOf

// NewConfigStoreItemResponseAllOf instantiates a new ConfigStoreItemResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigStoreItemResponseAllOf() *ConfigStoreItemResponseAllOf {
	this := ConfigStoreItemResponseAllOf{}
	return &this
}

// NewConfigStoreItemResponseAllOfWithDefaults instantiates a new ConfigStoreItemResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigStoreItemResponseAllOfWithDefaults() *ConfigStoreItemResponseAllOf {
	this := ConfigStoreItemResponseAllOf{}
	return &this
}

// GetStoreID returns the StoreID field value if set, zero value otherwise.
func (o *ConfigStoreItemResponseAllOf) GetStoreID() string {
	if o == nil || o.StoreID == nil {
		var ret string
		return ret
	}
	return *o.StoreID
}

// GetStoreIDOk returns a tuple with the StoreID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigStoreItemResponseAllOf) GetStoreIDOk() (*string, bool) {
	if o == nil || o.StoreID == nil {
		return nil, false
	}
	return o.StoreID, true
}

// HasStoreID returns a boolean if a field has been set.
func (o *ConfigStoreItemResponseAllOf) HasStoreID() bool {
	if o != nil && o.StoreID != nil {
		return true
	}

	return false
}

// SetStoreID gets a reference to the given string and assigns it to the StoreID field.
func (o *ConfigStoreItemResponseAllOf) SetStoreID(v string) {
	o.StoreID = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ConfigStoreItemResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.StoreID != nil {
		toSerialize["store_id"] = o.StoreID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (o *ConfigStoreItemResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varConfigStoreItemResponseAllOf := _ConfigStoreItemResponseAllOf{}

	if err = json.Unmarshal(bytes, &varConfigStoreItemResponseAllOf); err == nil {
		*o = ConfigStoreItemResponseAllOf(varConfigStoreItemResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "store_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableConfigStoreItemResponseAllOf is a helper abstraction for handling nullable configstoreitemresponseallof types. 
type NullableConfigStoreItemResponseAllOf struct {
	value *ConfigStoreItemResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableConfigStoreItemResponseAllOf) Get() *ConfigStoreItemResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableConfigStoreItemResponseAllOf) Set(val *ConfigStoreItemResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableConfigStoreItemResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableConfigStoreItemResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableConfigStoreItemResponseAllOf returns a pointer to a new instance of NullableConfigStoreItemResponseAllOf.
func NewNullableConfigStoreItemResponseAllOf(val *ConfigStoreItemResponseAllOf) *NullableConfigStoreItemResponseAllOf {
	return &NullableConfigStoreItemResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableConfigStoreItemResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves. 
func (v *NullableConfigStoreItemResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
