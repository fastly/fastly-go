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

// ComputeAclListEntries Defines the structure of what the ACL List endpoint returns.
type ComputeAclListEntries struct {
	Meta                 *ComputeAclListEntriesMeta  `json:"meta,omitempty"`
	Entries              []ComputeAclListEntriesItem `json:"entries,omitempty"`
	AdditionalProperties map[string]any
}

type _ComputeAclListEntries ComputeAclListEntries

// NewComputeAclListEntries instantiates a new ComputeAclListEntries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeAclListEntries() *ComputeAclListEntries {
	this := ComputeAclListEntries{}
	return &this
}

// NewComputeAclListEntriesWithDefaults instantiates a new ComputeAclListEntries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeAclListEntriesWithDefaults() *ComputeAclListEntries {
	this := ComputeAclListEntries{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ComputeAclListEntries) GetMeta() ComputeAclListEntriesMeta {
	if o == nil || o.Meta == nil {
		var ret ComputeAclListEntriesMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeAclListEntries) GetMetaOk() (*ComputeAclListEntriesMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ComputeAclListEntries) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ComputeAclListEntriesMeta and assigns it to the Meta field.
func (o *ComputeAclListEntries) SetMeta(v ComputeAclListEntriesMeta) {
	o.Meta = &v
}

// GetEntries returns the Entries field value if set, zero value otherwise.
func (o *ComputeAclListEntries) GetEntries() []ComputeAclListEntriesItem {
	if o == nil || o.Entries == nil {
		var ret []ComputeAclListEntriesItem
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeAclListEntries) GetEntriesOk() ([]ComputeAclListEntriesItem, bool) {
	if o == nil || o.Entries == nil {
		return nil, false
	}
	return o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *ComputeAclListEntries) HasEntries() bool {
	if o != nil && o.Entries != nil {
		return true
	}

	return false
}

// SetEntries gets a reference to the given []ComputeAclListEntriesItem and assigns it to the Entries field.
func (o *ComputeAclListEntries) SetEntries(v []ComputeAclListEntriesItem) {
	o.Entries = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ComputeAclListEntries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ComputeAclListEntries) UnmarshalJSON(bytes []byte) (err error) {
	varComputeAclListEntries := _ComputeAclListEntries{}

	if err = json.Unmarshal(bytes, &varComputeAclListEntries); err == nil {
		*o = ComputeAclListEntries(varComputeAclListEntries)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "meta")
		delete(additionalProperties, "entries")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableComputeAclListEntries is a helper abstraction for handling nullable computeacllistentries types.
type NullableComputeAclListEntries struct {
	value *ComputeAclListEntries
	isSet bool
}

// Get returns the value.
func (v NullableComputeAclListEntries) Get() *ComputeAclListEntries {
	return v.value
}

// Set modifies the value.
func (v *NullableComputeAclListEntries) Set(val *ComputeAclListEntries) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableComputeAclListEntries) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableComputeAclListEntries) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableComputeAclListEntries returns a pointer to a new instance of NullableComputeAclListEntries.
func NewNullableComputeAclListEntries(val *ComputeAclListEntries) *NullableComputeAclListEntries {
	return &NullableComputeAclListEntries{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableComputeAclListEntries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableComputeAclListEntries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
