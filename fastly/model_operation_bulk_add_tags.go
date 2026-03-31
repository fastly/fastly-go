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

// OperationBulkAddTags struct for OperationBulkAddTags
type OperationBulkAddTags struct {
	// List of operation IDs to add tags to.
	OperationIds []string `json:"operation_ids"`
	// List of tag IDs to add to the operations.
	TagIds               []string `json:"tag_ids"`
	AdditionalProperties map[string]any
}

type _OperationBulkAddTags OperationBulkAddTags

// NewOperationBulkAddTags instantiates a new OperationBulkAddTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOperationBulkAddTags(operationIds []string, tagIds []string) *OperationBulkAddTags {
	this := OperationBulkAddTags{}
	this.OperationIds = operationIds
	this.TagIds = tagIds
	return &this
}

// NewOperationBulkAddTagsWithDefaults instantiates a new OperationBulkAddTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOperationBulkAddTagsWithDefaults() *OperationBulkAddTags {
	this := OperationBulkAddTags{}
	return &this
}

// GetOperationIds returns the OperationIds field value
func (o *OperationBulkAddTags) GetOperationIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OperationIds
}

// GetOperationIdsOk returns a tuple with the OperationIds field value
// and a boolean to check if the value has been set.
func (o *OperationBulkAddTags) GetOperationIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperationIds, true
}

// SetOperationIds sets field value
func (o *OperationBulkAddTags) SetOperationIds(v []string) {
	o.OperationIds = v
}

// GetTagIds returns the TagIds field value
func (o *OperationBulkAddTags) GetTagIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value
// and a boolean to check if the value has been set.
func (o *OperationBulkAddTags) GetTagIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TagIds, true
}

// SetTagIds sets field value
func (o *OperationBulkAddTags) SetTagIds(v []string) {
	o.TagIds = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o OperationBulkAddTags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["operation_ids"] = o.OperationIds
	}
	if true {
		toSerialize["tag_ids"] = o.TagIds
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *OperationBulkAddTags) UnmarshalJSON(bytes []byte) (err error) {
	varOperationBulkAddTags := _OperationBulkAddTags{}

	if err = json.Unmarshal(bytes, &varOperationBulkAddTags); err == nil {
		*o = OperationBulkAddTags(varOperationBulkAddTags)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "operation_ids")
		delete(additionalProperties, "tag_ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableOperationBulkAddTags is a helper abstraction for handling nullable operationbulkaddtags types.
type NullableOperationBulkAddTags struct {
	value *OperationBulkAddTags
	isSet bool
}

// Get returns the value.
func (v NullableOperationBulkAddTags) Get() *OperationBulkAddTags {
	return v.value
}

// Set modifies the value.
func (v *NullableOperationBulkAddTags) Set(val *OperationBulkAddTags) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableOperationBulkAddTags) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableOperationBulkAddTags) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableOperationBulkAddTags returns a pointer to a new instance of NullableOperationBulkAddTags.
func NewNullableOperationBulkAddTags(val *OperationBulkAddTags) *NullableOperationBulkAddTags {
	return &NullableOperationBulkAddTags{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableOperationBulkAddTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableOperationBulkAddTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
