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

// ComputeACLList An example of an ACL list response.
type ComputeACLList struct {
	Data                 []ComputeACLCreateAclsResponse `json:"data,omitempty"`
	Meta                 *ComputeACLListMeta            `json:"meta,omitempty"`
	AdditionalProperties map[string]any
}

type _ComputeACLList ComputeACLList

// NewComputeACLList instantiates a new ComputeACLList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComputeACLList() *ComputeACLList {
	this := ComputeACLList{}
	return &this
}

// NewComputeACLListWithDefaults instantiates a new ComputeACLList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComputeACLListWithDefaults() *ComputeACLList {
	this := ComputeACLList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ComputeACLList) GetData() []ComputeACLCreateAclsResponse {
	if o == nil || o.Data == nil {
		var ret []ComputeACLCreateAclsResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeACLList) GetDataOk() ([]ComputeACLCreateAclsResponse, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ComputeACLList) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []ComputeACLCreateAclsResponse and assigns it to the Data field.
func (o *ComputeACLList) SetData(v []ComputeACLCreateAclsResponse) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ComputeACLList) GetMeta() ComputeACLListMeta {
	if o == nil || o.Meta == nil {
		var ret ComputeACLListMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComputeACLList) GetMetaOk() (*ComputeACLListMeta, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ComputeACLList) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ComputeACLListMeta and assigns it to the Meta field.
func (o *ComputeACLList) SetMeta(v ComputeACLListMeta) {
	o.Meta = &v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o ComputeACLList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Meta != nil {
		toSerialize["meta"] = o.Meta
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *ComputeACLList) UnmarshalJSON(bytes []byte) (err error) {
	varComputeACLList := _ComputeACLList{}

	if err = json.Unmarshal(bytes, &varComputeACLList); err == nil {
		*o = ComputeACLList(varComputeACLList)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "meta")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableComputeACLList is a helper abstraction for handling nullable computeacllist types.
type NullableComputeACLList struct {
	value *ComputeACLList
	isSet bool
}

// Get returns the value.
func (v NullableComputeACLList) Get() *ComputeACLList {
	return v.value
}

// Set modifies the value.
func (v *NullableComputeACLList) Set(val *ComputeACLList) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableComputeACLList) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableComputeACLList) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableComputeACLList returns a pointer to a new instance of NullableComputeACLList.
func NewNullableComputeACLList(val *ComputeACLList) *NullableComputeACLList {
	return &NullableComputeACLList{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableComputeACLList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableComputeACLList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
