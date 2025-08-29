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

// AttackSignal struct for AttackSignal
type AttackSignal struct {
	// Name of the attack signal tag
	TagName string `json:"tag_name"`
	// Count of requests with this attack signal
	TagCount int32 `json:"tag_count"`
	// Total number of attacks considered
	TotalCount           int32 `json:"total_count"`
	AdditionalProperties map[string]any
}

type _AttackSignal AttackSignal

// NewAttackSignal instantiates a new AttackSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttackSignal(tagName string, tagCount int32, totalCount int32) *AttackSignal {
	this := AttackSignal{}
	this.TagName = tagName
	this.TagCount = tagCount
	this.TotalCount = totalCount
	return &this
}

// NewAttackSignalWithDefaults instantiates a new AttackSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttackSignalWithDefaults() *AttackSignal {
	this := AttackSignal{}
	return &this
}

// GetTagName returns the TagName field value
func (o *AttackSignal) GetTagName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TagName
}

// GetTagNameOk returns a tuple with the TagName field value
// and a boolean to check if the value has been set.
func (o *AttackSignal) GetTagNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagName, true
}

// SetTagName sets field value
func (o *AttackSignal) SetTagName(v string) {
	o.TagName = v
}

// GetTagCount returns the TagCount field value
func (o *AttackSignal) GetTagCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TagCount
}

// GetTagCountOk returns a tuple with the TagCount field value
// and a boolean to check if the value has been set.
func (o *AttackSignal) GetTagCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TagCount, true
}

// SetTagCount sets field value
func (o *AttackSignal) SetTagCount(v int32) {
	o.TagCount = v
}

// GetTotalCount returns the TotalCount field value
func (o *AttackSignal) GetTotalCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *AttackSignal) GetTotalCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *AttackSignal) SetTotalCount(v int32) {
	o.TotalCount = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o AttackSignal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["tag_name"] = o.TagName
	}
	if true {
		toSerialize["tag_count"] = o.TagCount
	}
	if true {
		toSerialize["total_count"] = o.TotalCount
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *AttackSignal) UnmarshalJSON(bytes []byte) (err error) {
	varAttackSignal := _AttackSignal{}

	if err = json.Unmarshal(bytes, &varAttackSignal); err == nil {
		*o = AttackSignal(varAttackSignal)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "tag_name")
		delete(additionalProperties, "tag_count")
		delete(additionalProperties, "total_count")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableAttackSignal is a helper abstraction for handling nullable attacksignal types.
type NullableAttackSignal struct {
	value *AttackSignal
	isSet bool
}

// Get returns the value.
func (v NullableAttackSignal) Get() *AttackSignal {
	return v.value
}

// Set modifies the value.
func (v *NullableAttackSignal) Set(val *AttackSignal) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableAttackSignal) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableAttackSignal) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableAttackSignal returns a pointer to a new instance of NullableAttackSignal.
func NewNullableAttackSignal(val *AttackSignal) *NullableAttackSignal {
	return &NullableAttackSignal{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableAttackSignal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableAttackSignal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
