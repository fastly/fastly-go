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

// ComputeAclListMeta Meta for the total of ACLs.
type ComputeAclListMeta struct {
	// Total of ACLs.
	Total                *int32 `json:"total,omitempty"`
	AdditionalProperties map[string]any
}

type _ComputeAclListMeta ComputeAclListMeta

// NewComputeAclListMeta instantiates a new ComputeAclListMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeAclListMeta() *ComputeAclListMeta {
	this := ComputeAclListMeta{}
	return &this
}

// NewComputeAclListMetaWithDefaults instantiates a new ComputeAclListMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeAclListMetaWithDefaults() *ComputeAclListMeta {
	this := ComputeAclListMeta{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ComputeAclListMeta) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeAclListMeta) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ComputeAclListMeta) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ComputeAclListMeta) SetTotal(v int32) {
	o.Total = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ComputeAclListMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ComputeAclListMeta) UnmarshalJSON(bytes []byte) (err error) {
	varComputeAclListMeta := _ComputeAclListMeta{}

	if err = json.Unmarshal(bytes, &varComputeAclListMeta); err == nil {
		*o = ComputeAclListMeta(varComputeAclListMeta)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableComputeAclListMeta is a helper abstraction for handling nullable computeacllistmeta types.
type NullableComputeAclListMeta struct {
	value *ComputeAclListMeta
	isSet bool
}

// Get returns the value.
func (v NullableComputeAclListMeta) Get() *ComputeAclListMeta {
	return v.value
}

// Set modifies the value.
func (v *NullableComputeAclListMeta) Set(val *ComputeAclListMeta) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableComputeAclListMeta) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableComputeAclListMeta) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableComputeAclListMeta returns a pointer to a new instance of NullableComputeAclListMeta.
func NewNullableComputeAclListMeta(val *ComputeAclListMeta) *NullableComputeAclListMeta {
	return &NullableComputeAclListMeta{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableComputeAclListMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableComputeAclListMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
