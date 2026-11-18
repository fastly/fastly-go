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

// NetworkProtectRuleMatchSourcePrefixList Requires a source prefix-list match field.
type NetworkProtectRuleMatchSourcePrefixList struct {
	// Match traffic whose source appears in this prefix list.
	SourcePrefixListID   string `json:"source_prefix_list_id"`
	AdditionalProperties map[string]any
}

type _NetworkProtectRuleMatchSourcePrefixList NetworkProtectRuleMatchSourcePrefixList

// NewNetworkProtectRuleMatchSourcePrefixList instantiates a new NetworkProtectRuleMatchSourcePrefixList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkProtectRuleMatchSourcePrefixList(sourcePrefixListID string) *NetworkProtectRuleMatchSourcePrefixList {
	this := NetworkProtectRuleMatchSourcePrefixList{}
	this.SourcePrefixListID = sourcePrefixListID
	return &this
}

// NewNetworkProtectRuleMatchSourcePrefixListWithDefaults instantiates a new NetworkProtectRuleMatchSourcePrefixList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkProtectRuleMatchSourcePrefixListWithDefaults() *NetworkProtectRuleMatchSourcePrefixList {
	this := NetworkProtectRuleMatchSourcePrefixList{}
	return &this
}

// GetSourcePrefixListID returns the SourcePrefixListID field value
func (o *NetworkProtectRuleMatchSourcePrefixList) GetSourcePrefixListID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourcePrefixListID
}

// GetSourcePrefixListIDOk returns a tuple with the SourcePrefixListID field value
// and a boolean to check if the value has been set.
func (o *NetworkProtectRuleMatchSourcePrefixList) GetSourcePrefixListIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourcePrefixListID, true
}

// SetSourcePrefixListID sets field value
func (o *NetworkProtectRuleMatchSourcePrefixList) SetSourcePrefixListID(v string) {
	o.SourcePrefixListID = v
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (o NetworkProtectRuleMatchSourcePrefixList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["source_prefix_list_id"] = o.SourcePrefixListID
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (o *NetworkProtectRuleMatchSourcePrefixList) UnmarshalJSON(bytes []byte) (err error) {
	varNetworkProtectRuleMatchSourcePrefixList := _NetworkProtectRuleMatchSourcePrefixList{}

	if err = json.Unmarshal(bytes, &varNetworkProtectRuleMatchSourcePrefixList); err == nil {
		*o = NetworkProtectRuleMatchSourcePrefixList(varNetworkProtectRuleMatchSourcePrefixList)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "source_prefix_list_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// NullableNetworkProtectRuleMatchSourcePrefixList is a helper abstraction for handling nullable networkprotectrulematchsourceprefixlist types.
type NullableNetworkProtectRuleMatchSourcePrefixList struct {
	value *NetworkProtectRuleMatchSourcePrefixList
	isSet bool
}

// Get returns the value.
func (v NullableNetworkProtectRuleMatchSourcePrefixList) Get() *NetworkProtectRuleMatchSourcePrefixList {
	return v.value
}

// Set modifies the value.
func (v *NullableNetworkProtectRuleMatchSourcePrefixList) Set(val *NetworkProtectRuleMatchSourcePrefixList) {
	v.value = val
	v.isSet = true
}

// IsSet indicates if the value was set.
func (v NullableNetworkProtectRuleMatchSourcePrefixList) IsSet() bool {
	return v.isSet
}

// Unset removes the value.
func (v *NullableNetworkProtectRuleMatchSourcePrefixList) Unset() {
	v.value = nil
	v.isSet = false
}

// NewNullableNetworkProtectRuleMatchSourcePrefixList returns a pointer to a new instance of NullableNetworkProtectRuleMatchSourcePrefixList.
func NewNullableNetworkProtectRuleMatchSourcePrefixList(val *NetworkProtectRuleMatchSourcePrefixList) *NullableNetworkProtectRuleMatchSourcePrefixList {
	return &NullableNetworkProtectRuleMatchSourcePrefixList{value: val, isSet: true}
}

// MarshalJSON implements the json.Marshaler interface.
// Marshaler is the interface implemented by types that can marshal themselves into valid JSON.
func (v NullableNetworkProtectRuleMatchSourcePrefixList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

// UnmarshalJSON implements the Unmarshaler interface.
// Unmarshaler is the interface implemented by types that can unmarshal a JSON description of themselves.
func (v *NullableNetworkProtectRuleMatchSourcePrefixList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
