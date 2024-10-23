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

// PlatformDdosEntry struct for PlatformDdosEntry
type PlatformDdosEntry struct {
	// An array of values representing the metric values at each point in time. Note that this dataset is sparse: only the keys with non-zero values will be included in the record.
	Values               []ValuesDdos `json:"values,omitempty"`
	AdditionalProperties map[string]any
}

type _PlatformDdosEntry PlatformDdosEntry

// NewPlatformDdosEntry instantiates a new PlatformDdosEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlatformDdosEntry() *PlatformDdosEntry {
	this := PlatformDdosEntry{}
	return &this
}

// NewPlatformDdosEntryWithDefaults instantiates a new PlatformDdosEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlatformDdosEntryWithDefaults() *PlatformDdosEntry {
	this := PlatformDdosEntry{}
	return &this
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *PlatformDdosEntry) GetValues() []ValuesDdos {
	if o == nil || o.Values == nil {
		var ret []ValuesDdos
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlatformDdosEntry) GetValuesOk() ([]ValuesDdos, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *PlatformDdosEntry) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []ValuesDdos and assigns it to the Values field.
func (o *PlatformDdosEntry) SetValues(v []ValuesDdos) {
	o.Values = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o PlatformDdosEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *PlatformDdosEntry) UnmarshalJSON(bytes []byte) (err error) {
	varPlatformDdosEntry := _PlatformDdosEntry{}

	if err = json.Unmarshal(bytes, &varPlatformDdosEntry); err == nil {
		*o = PlatformDdosEntry(varPlatformDdosEntry)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullablePlatformDdosEntry is a helper abstraction for handling nullable platformddosentry types.
type NullablePlatformDdosEntry struct {
	value *PlatformDdosEntry
	isSet bool
}

// Get returns the value.
func (v NullablePlatformDdosEntry) Get() *PlatformDdosEntry {
	return v.value
}

// Set modifies the value.
func (v *NullablePlatformDdosEntry) Set(val *PlatformDdosEntry) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullablePlatformDdosEntry) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullablePlatformDdosEntry) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullablePlatformDdosEntry returns a pointer to a new instance of NullablePlatformDdosEntry.
func NewNullablePlatformDdosEntry(val *PlatformDdosEntry) *NullablePlatformDdosEntry {
	return &NullablePlatformDdosEntry{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullablePlatformDdosEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullablePlatformDdosEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
