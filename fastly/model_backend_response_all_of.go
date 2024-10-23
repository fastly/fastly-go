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

// BackendResponseAllOf struct for BackendResponseAllOf
type BackendResponseAllOf struct {
	// Indicates whether the version of the service this backend is attached to accepts edits.
	Locked               *bool `json:"locked,omitempty"`
	AdditionalProperties map[string]any
}

type _BackendResponseAllOf BackendResponseAllOf

// NewBackendResponseAllOf instantiates a new BackendResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackendResponseAllOf() *BackendResponseAllOf {
	this := BackendResponseAllOf{}
	return &this
}

// NewBackendResponseAllOfWithDefaults instantiates a new BackendResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackendResponseAllOfWithDefaults() *BackendResponseAllOf {
	this := BackendResponseAllOf{}
	return &this
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *BackendResponseAllOf) GetLocked() bool {
	if o == nil || o.Locked == nil {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendResponseAllOf) GetLockedOk() (*bool, bool) {
	if o == nil || o.Locked == nil {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *BackendResponseAllOf) HasLocked() bool {
	if o != nil && o.Locked != nil {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *BackendResponseAllOf) SetLocked(v bool) {
	o.Locked = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o BackendResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Locked != nil {
		toSerialize["locked"] = o.Locked
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *BackendResponseAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varBackendResponseAllOf := _BackendResponseAllOf{}

	if err = json.Unmarshal(bytes, &varBackendResponseAllOf); err == nil {
		*o = BackendResponseAllOf(varBackendResponseAllOf)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "locked")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableBackendResponseAllOf is a helper abstraction for handling nullable backendresponseallof types.
type NullableBackendResponseAllOf struct {
	value *BackendResponseAllOf
	isSet bool
}

// Get returns the value.
func (v NullableBackendResponseAllOf) Get() *BackendResponseAllOf {
	return v.value
}

// Set modifies the value.
func (v *NullableBackendResponseAllOf) Set(val *BackendResponseAllOf) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableBackendResponseAllOf) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableBackendResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableBackendResponseAllOf returns a pointer to a new instance of NullableBackendResponseAllOf.
func NewNullableBackendResponseAllOf(val *BackendResponseAllOf) *NullableBackendResponseAllOf {
	return &NullableBackendResponseAllOf{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableBackendResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableBackendResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
