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

// DictionaryItemResponseAllOf struct for DictionaryItemResponseAllOf
type DictionaryItemResponseAllOf struct {
	DictionaryId         *string `json:"dictionary_id,omitempty"`
	ServiceId            *string `json:"service_id,omitempty"`
	AdditionalProperties map[string]any
}

type _DictionaryItemResponseAllOf DictionaryItemResponseAllOf

// NewDictionaryItemResponseAllOf instantiates a new DictionaryItemResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDictionaryItemResponseAllOf() *DictionaryItemResponseAllOf {
	this := DictionaryItemResponseAllOf{}
	return &this
}

// NewDictionaryItemResponseAllOfWithDefaults instantiates a new DictionaryItemResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDictionaryItemResponseAllOfWithDefaults() *DictionaryItemResponseAllOf {
	this := DictionaryItemResponseAllOf{}
	return &this
}

// GetDictionaryId returns the DictionaryId field value if set, zero value otherwise.
func (o *DictionaryItemResponseAllOf) GetDictionaryId() string {
	if o == nil || o.DictionaryId == nil {
		var ret string
		return ret
	}
	return *o.DictionaryId
}

// GetDictionaryIdOk returns a tuple with the DictionaryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryItemResponseAllOf) GetDictionaryIdOk() (*string, bool) {
	if o == nil || o.DictionaryId == nil {
		return nil, false
	}
	return o.DictionaryId, true
}

// HasDictionaryId returns a boolean if a field has been set.
func (o *DictionaryItemResponseAllOf) HasDictionaryId() bool {
	if o != nil && o.DictionaryId != nil {
		return true
	}

	return false
}

// SetDictionaryId gets a reference to the given string and assigns it to the DictionaryId field.
func (o *DictionaryItemResponseAllOf) SetDictionaryId(v string) {
	o.DictionaryId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *DictionaryItemResponseAllOf) GetServiceId() string {
	if o == nil || o.ServiceId == nil {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DictionaryItemResponseAllOf) GetServiceIdOk() (*string, bool) {
	if o == nil || o.ServiceId == nil {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *DictionaryItemResponseAllOf) HasServiceId() bool {
	if o != nil && o.ServiceId != nil {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *DictionaryItemResponseAllOf) SetServiceId(v string) {
	o.ServiceId = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o DictionaryItemResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.DictionaryId != nil {
		toSerialize["dictionary_id"] = o.DictionaryId
	}
	if o.ServiceId != nil {
		toSerialize["service_id"] = o.ServiceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *DictionaryItemResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varDictionaryItemResponseAllOf := _DictionaryItemResponseAllOf{}

	if err = json.Unmarshal(bytes, &varDictionaryItemResponseAllOf); err == nil {
		*o = DictionaryItemResponseAllOf(varDictionaryItemResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dictionary_id")
		delete(additionalProperties, "service_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableDictionaryItemResponseAllOf is a helper abstraction for handling nullable dictionaryitemresponseallof types.
type NullableDictionaryItemResponseAllOf struct {
	value *DictionaryItemResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableDictionaryItemResponseAllOf) Get() *DictionaryItemResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableDictionaryItemResponseAllOf) Set(val *DictionaryItemResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableDictionaryItemResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableDictionaryItemResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableDictionaryItemResponseAllOf returns a pointer to a new instance of NullableDictionaryItemResponseAllOf.
func NewNullableDictionaryItemResponseAllOf(val *DictionaryItemResponseAllOf) *NullableDictionaryItemResponseAllOf {
	return &NullableDictionaryItemResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableDictionaryItemResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableDictionaryItemResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
